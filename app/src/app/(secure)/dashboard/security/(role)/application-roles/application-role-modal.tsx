'use client';

import { useState, useEffect } from 'react';
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle } from '@/components/ui/dialog';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { Textarea } from '@/components/ui/textarea';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { toast } from 'sonner';
import { fn_create_application_role, CreateApplicationRoleInput } from '@/actions/application-roles/fn_create_application_role';
import { fn_update_application_role, UpdateApplicationRoleInput } from '@/actions/application-roles/fn_update_application_role';
import { fn_get_applications } from '@/actions/applications/fn_get_applications';
import type { ApplicationRoleItem } from '@/types/application-roles';
import type { ApplicationItem } from '@/types/applications';

interface ApplicationRoleModalProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  role?: ApplicationRoleItem | null;
  onSuccess: () => void;
  defaultApplicationId?: string;
}

export default function ApplicationRoleModal({ open, onOpenChange, role, onSuccess, defaultApplicationId }: ApplicationRoleModalProps) {
  const [loading, setLoading] = useState(false);
  const [loadingApps, setLoadingApps] = useState(false);
  const [applications, setApplications] = useState<ApplicationItem[]>([]);
  const [formData, setFormData] = useState<CreateApplicationRoleInput>({
    name: '',
    description: '',
    application_id: '',
  });

  useEffect(() => {
    if (open) {
      loadApplications();
    }
  }, [open]);

  useEffect(() => {
    if (role) {
      setFormData({
        name: role.name || '',
        description: role.description || '',
        application_id: role.application_id || '',
      });
    } else {
      setFormData({
        name: '',
        description: '',
        application_id: defaultApplicationId || '',
      });
    }
  }, [role, open, defaultApplicationId]);

  const loadApplications = async () => {
    try {
      setLoadingApps(true);
      const response = await fn_get_applications(1, 100, false);
      setApplications(response.data);
    } catch (error) {
      console.error('Error loading applications:', error);
      toast.error('Error al cargar aplicaciones');
    } finally {
      setLoadingApps(false);
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);

    try {
      const payload = {
        name: formData.name,
        description: formData.description || undefined,
        application_id: formData.application_id,
      };

      if (role) {
        const { application_id, ...updatePayload } = payload;
        await fn_update_application_role(role.id, updatePayload as UpdateApplicationRoleInput);
        toast.success('Rol actualizado correctamente');
      } else {
        await fn_create_application_role(payload);
        toast.success('Rol creado correctamente');
      }
      onSuccess();
      onOpenChange(false);
    } catch (error: any) {
      toast.error(error.message || 'Error al guardar el rol');
    } finally {
      setLoading(false);
    }
  };

  const isFormValid = 
    (formData.name || '').trim() !== '' && 
    (formData.application_id || '').trim() !== '';

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="bg-card/80 backdrop-blur-xl border-border sm:max-w-[550px]">
        <DialogHeader>
          <DialogTitle>{role ? 'Editar Rol' : 'Crear Nuevo Rol'}</DialogTitle>
          <DialogDescription>
            {role ? 'Actualiza la información del rol de aplicación.' : 'Registra un nuevo rol para la aplicación seleccionada.'}
          </DialogDescription>
        </DialogHeader>

        <form onSubmit={handleSubmit} className="space-y-4">
          {/* Aplicación */}
          <div className="space-y-2">
            <Label htmlFor="application_id">
              Aplicación <span className="text-destructive">*</span>
            </Label>
            <Select
              value={formData.application_id || ' '}
              onValueChange={(value) => setFormData({ ...formData, application_id: value === ' ' ? '' : value })}
              disabled={loadingApps || !!defaultApplicationId || !!role}
            >
              <SelectTrigger className="[&>span]:truncate">
                <SelectValue placeholder="Seleccionar aplicación" />
              </SelectTrigger>
              <SelectContent position="popper" sideOffset={5} className="max-h-[300px]">
                <SelectItem value=" " disabled>Selecciona una aplicación</SelectItem>
                {applications.map((app) => (
                  <SelectItem key={app.id} value={app.id} title={`${app.name} (${app.client_id})`}>
                    {app.name}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
            {(defaultApplicationId || role) && (
              <p className="text-xs text-muted-foreground">
                La aplicación no se puede cambiar
              </p>
            )}
          </div>

          {/* Nombre */}
          <div className="space-y-2">
            <Label htmlFor="name">
              Nombre del Rol <span className="text-destructive">*</span>
            </Label>
            <Input
              id="name"
              value={formData.name || ''}
              onChange={(e) => setFormData({ ...formData, name: e.target.value })}
              placeholder="Ej: Administrator, Editor, Viewer"
              required
            />
            <p className="text-xs text-muted-foreground">
              Nombre descriptivo del rol (ej: Administrator, Editor)
            </p>
          </div>

          {/* Descripción */}
          <div className="space-y-2">
            <Label htmlFor="description">Descripción</Label>
            <Textarea
              id="description"
              value={formData.description || ''}
              onChange={(e) => setFormData({ ...formData, description: e.target.value })}
              placeholder="Describe los permisos y responsabilidades de este rol..."
              rows={4}
            />
            <p className="text-xs text-muted-foreground">
              Descripción opcional de las funciones y permisos del rol
            </p>
          </div>

          <div className="bg-muted/50 p-4 rounded-lg">
            <p className="text-sm font-medium mb-2">Roles de Aplicación:</p>
            <ul className="text-sm text-muted-foreground space-y-1">
              <li>• Los roles definen conjuntos de permisos para los usuarios</li>
              <li>• Cada rol pertenece a una aplicación específica</li>
              <li>• Puedes asignar módulos a cada rol para controlar accesos</li>
              <li>• Los usuarios pueden tener múltiples roles</li>
            </ul>
          </div>

          <DialogFooter>
            <Button type="button" variant="outline" onClick={() => onOpenChange(false)} disabled={loading}>
              Cancelar
            </Button>
            <Button type="submit" className="bg-linear-to-r from-primary to-chart-1" disabled={loading || !isFormValid || loadingApps}>
              {loading ? 'Guardando...' : role ? 'Actualizar Rol' : 'Crear Rol'}
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  );
}