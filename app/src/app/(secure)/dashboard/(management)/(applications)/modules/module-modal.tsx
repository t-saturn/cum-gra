'use client';

import { useState, useEffect } from 'react';
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle } from '@/components/ui/dialog';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { toast } from 'sonner';
import { fn_create_module, CreateModuleInput } from '@/actions/modules/fn_create_module';
import { fn_update_module, UpdateModuleInput } from '@/actions/modules/fn_update_module';
import { fn_get_module_form_options, ModuleFormOptions } from '@/actions/modules/fn_get_module_form_options';
import type { ModuleItem } from '@/types/modules';

interface ModuleModalProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  module?: ModuleItem | null;
  onSuccess: () => void;
  defaultApplicationId?: string;
}

export default function ModuleModal({ open, onOpenChange, module, onSuccess, defaultApplicationId }: ModuleModalProps) {
  const [loading, setLoading] = useState(false);
  const [loadingOptions, setLoadingOptions] = useState(false);
  const [loadingModules, setLoadingModules] = useState(false);
  const [options, setOptions] = useState<ModuleFormOptions>({
    applications: [],
    modules: [],
  });
  const [formData, setFormData] = useState<CreateModuleInput>({
    item: '',
    name: '',
    route: '',
    icon: '',
    parent_id: '',
    application_id: '',
    sort_order: 1,
    status: 'active',
  });

  useEffect(() => {
    if (open) {
      loadFormOptions();
    }
  }, [open]);

  useEffect(() => {
    if (module) {
      setFormData({
        item: module.item || '',
        name: module.name || '',
        route: module.route || '',
        icon: module.icon || '',
        parent_id: module.parent_id || '',
        application_id: module.application_id || '',
        sort_order: module.sort_order || 1,
        status: module.status || 'active',
      });
    } else {
      setFormData({
        item: '',
        name: '',
        route: '',
        icon: '',
        parent_id: '',
        application_id: defaultApplicationId || '',
        sort_order: 1,
        status: 'active',
      });
    }
  }, [module, open, defaultApplicationId]);

  // Cargar módulos cuando cambia la aplicación
  useEffect(() => {
    if (formData.application_id && open) {
      loadModulesForApplication(formData.application_id);
    }
  }, [formData.application_id, open]);

  const loadFormOptions = async () => {
    try {
      setLoadingOptions(true);
      // Solo cargar aplicaciones inicialmente
      const data = await fn_get_module_form_options();
      setOptions({
        applications: data.applications,
        modules: [], // Inicialmente vacío
      });

      // Si hay una aplicación por defecto, cargar sus módulos
      if (defaultApplicationId) {
        await loadModulesForApplication(defaultApplicationId);
      } else if (module?.application_id) {
        await loadModulesForApplication(module.application_id);
      }
    } catch (error) {
      console.error('Error loading form options:', error);
      toast.error('Error al cargar opciones del formulario');
    } finally {
      setLoadingOptions(false);
    }
  };

  const loadModulesForApplication = async (applicationId: string) => {
    try {
      setLoadingModules(true);
      const data = await fn_get_module_form_options(applicationId);
      setOptions(prev => ({
        ...prev,
        modules: data.modules,
      }));
    } catch (error) {
      console.error('Error loading modules for application:', error);
      toast.error('Error al cargar módulos de la aplicación');
    } finally {
      setLoadingModules(false);
    }
  };

  const handleApplicationChange = (appId: string) => {
    setFormData({ 
      ...formData, 
      application_id: appId === ' ' ? '' : appId,
      parent_id: '' // Limpiar el padre cuando cambia la aplicación
    });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);

    try {
      const payload = {
        item: formData.item || undefined,
        name: formData.name,
        route: formData.route,
        icon: formData.icon || undefined,
        parent_id: formData.parent_id || undefined,
        application_id: formData.application_id || undefined,
        sort_order: formData.sort_order,
        status: formData.status,
      };

      if (module) {
        await fn_update_module(module.id, payload as UpdateModuleInput);
        toast.success('Módulo actualizado correctamente');
      } else {
        await fn_create_module(payload);
        toast.success('Módulo creado correctamente');
      }
      onSuccess();
      onOpenChange(false);
    } catch (error: any) {
      toast.error(error.message || 'Error al guardar el módulo');
    } finally {
      setLoading(false);
    }
  };

  const isFormValid = 
    (formData.name || '').trim() !== '' && 
    (formData.route || '').trim() !== '' &&
    (formData.application_id || '').trim() !== '';

  // Filtrar módulos para evitar que un módulo sea su propio padre
  const availableParentModules = options.modules.filter(m => m.id !== module?.id);

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="bg-card/80 backdrop-blur-xl border-border sm:max-w-[600px] max-h-[90vh] overflow-y-auto">
        <DialogHeader>
          <DialogTitle>{module ? 'Editar Módulo' : 'Crear Nuevo Módulo'}</DialogTitle>
          <DialogDescription>
            {module ? 'Actualiza la información del módulo.' : 'Registra un nuevo módulo en el sistema.'}
          </DialogDescription>
        </DialogHeader>

        <form onSubmit={handleSubmit} className="space-y-4">
          {/* Aplicación (requerido y bloqueado si viene por defecto) */}
          <div className="space-y-2">
            <Label htmlFor="application_id">
              Aplicación <span className="text-destructive">*</span>
            </Label>
            <Select
              value={formData.application_id || ' '}
              onValueChange={handleApplicationChange}
              disabled={loadingOptions || !!defaultApplicationId || !!module}
            >
              <SelectTrigger className="[&>span]:truncate">
                <SelectValue placeholder="Seleccionar aplicación" />
              </SelectTrigger>
              <SelectContent position="popper" sideOffset={5} className="max-h-[300px]">
                <SelectItem value=" " disabled>Selecciona una aplicación</SelectItem>
                {options.applications.map((app) => (
                  <SelectItem key={app.id} value={app.id} title={`${app.name} (${app.client_id})`}>
                    {app.name}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
            {(defaultApplicationId || module) && (
              <p className="text-xs text-muted-foreground">
                La aplicación no se puede cambiar
              </p>
            )}
          </div>

          <div className="grid grid-cols-2 gap-4">
            <div className="space-y-2">
              <Label htmlFor="name">
                Nombre <span className="text-destructive">*</span>
              </Label>
              <Input
                id="name"
                value={formData.name || ''}
                onChange={(e) => setFormData({ ...formData, name: e.target.value })}
                placeholder="Ej: Dashboard Principal"
                required
              />
            </div>

            <div className="space-y-2">
              <Label htmlFor="item">Item</Label>
              <Input
                id="item"
                value={formData.item || ''}
                onChange={(e) => setFormData({ ...formData, item: e.target.value })}
                placeholder="Ej: dashboard"
              />
              <p className="text-xs text-muted-foreground">Identificador técnico</p>
            </div>
          </div>

          <div className="space-y-2">
            <Label htmlFor="route">
              Ruta <span className="text-destructive">*</span>
            </Label>
            <Input
              id="route"
              value={formData.route || ''}
              onChange={(e) => setFormData({ ...formData, route: e.target.value })}
              placeholder="Ej: /dashboard"
              required
            />
            <p className="text-xs text-muted-foreground">Ruta de navegación en la aplicación</p>
          </div>

          <div className="space-y-2">
            <Label htmlFor="icon">Icono</Label>
            <Input
              id="icon"
              value={formData.icon || ''}
              onChange={(e) => setFormData({ ...formData, icon: e.target.value })}
              placeholder="Ej: dashboard-icon, Home, Settings"
            />
            <p className="text-xs text-muted-foreground">Nombre del icono (Lucide React)</p>
          </div>

          <div className="space-y-2">
            <Label htmlFor="parent_id">Módulo Padre</Label>
            <Select
              value={formData.parent_id || ' '}
              onValueChange={(value) => setFormData({ ...formData, parent_id: value === ' ' ? '' : value })}
              disabled={loadingOptions || loadingModules || !formData.application_id}
            >
              <SelectTrigger className="[&>span]:truncate">
                <SelectValue placeholder={loadingModules ? "Cargando módulos..." : "Seleccionar padre"} />
              </SelectTrigger>
              <SelectContent position="popper" sideOffset={5} className="max-h-[300px]">
                <SelectItem value=" ">Sin padre (Raíz)</SelectItem>
                {availableParentModules.length === 0 && formData.application_id && !loadingModules ? (
                  <SelectItem value="__empty__" disabled>
                    No hay módulos disponibles
                  </SelectItem>
                ) : (
                  availableParentModules.map((mod) => (
                    <SelectItem key={mod.id} value={mod.id} title={`${mod.name} - ${mod.route}`}>
                      {mod.name}
                    </SelectItem>
                  ))
                )}
              </SelectContent>
            </Select>
            <p className="text-xs text-muted-foreground">
              {!formData.application_id 
                ? "Selecciona primero una aplicación"
                : "Para crear submódulos, selecciona el módulo padre"
              }
            </p>
          </div>

          <div className="grid grid-cols-2 gap-4">
            <div className="space-y-2">
              <Label htmlFor="sort_order">
                Orden <span className="text-destructive">*</span>
              </Label>
              <Input
                id="sort_order"
                type="number"
                min="1"
                value={formData.sort_order}
                onChange={(e) => setFormData({ ...formData, sort_order: parseInt(e.target.value) || 1 })}
                required
              />
              <p className="text-xs text-muted-foreground">Posición en el menú</p>
            </div>

            <div className="space-y-2">
              <Label htmlFor="status">
                Estado <span className="text-destructive">*</span>
              </Label>
              <Select value={formData.status} onValueChange={(value: any) => setFormData({ ...formData, status: value })}>
                <SelectTrigger>
                  <SelectValue placeholder="Seleccionar estado" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="active">Activo</SelectItem>
                  <SelectItem value="inactive">Inactivo</SelectItem>
                </SelectContent>
              </Select>
            </div>
          </div>

          <div className="bg-muted/50 p-4 rounded-lg">
            <p className="text-sm font-medium mb-2">Estructura de Módulos:</p>
            <ul className="text-sm text-muted-foreground space-y-1">
              <li>• Cada módulo pertenece a una aplicación específica</li>
              <li>• Los módulos pueden organizarse jerárquicamente</li>
              <li>• Un módulo sin padre es un módulo raíz del menú</li>
              <li>• El orden determina la posición en el menú de navegación</li>
            </ul>
          </div>

          <DialogFooter>
            <Button type="button" variant="outline" onClick={() => onOpenChange(false)} disabled={loading}>
              Cancelar
            </Button>
            <Button type="submit" className="bg-linear-to-r from-primary to-chart-1" disabled={loading || !isFormValid || loadingOptions}>
              {loading ? 'Guardando...' : module ? 'Actualizar Módulo' : 'Crear Módulo'}
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  );
}