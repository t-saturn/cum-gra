'use client';

import { useState, useEffect } from 'react';
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle } from '@/components/ui/dialog';
import { Button } from '@/components/ui/button';
import { Label } from '@/components/ui/label';
import { Input } from '@/components/ui/input';
import { Textarea } from '@/components/ui/textarea';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { toast } from 'sonner';
import { fn_get_modules } from '@/actions/modules/fn_get_modules';
import type { UserRestrictionItem, RestrictionType, PermissionLevel } from '@/types/user-restrictions';
import type { ModuleItem } from '@/types/modules';
import { fn_get_restriction_form_options } from '@/actions/users-restrictions/fn_get_restriction_form_options';
import { fn_update_user_restriction } from '@/actions/users-restrictions/fn_update_user_restriction';
import { fn_create_user_restriction } from '@/actions/users-restrictions/fn_create_user_restriction';

interface RestrictionModalProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  restriction?: UserRestrictionItem | null;
  onSuccess: () => void;
  defaultApplicationId?: string;
}

export default function RestrictionModal({ open, onOpenChange, restriction, onSuccess, defaultApplicationId }: RestrictionModalProps) {
  const [loading, setLoading] = useState(false);
  const [loadingModules, setLoadingModules] = useState(false);

  // Estados de opciones
  const [users, setUsers] = useState<Array<{ id: string; full_name: string; email: string; dni: string }>>([]);
  const [apps, setApps] = useState<Array<{ id: string; name: string }>>([]);
  const [modules, setModules] = useState<ModuleItem[]>([]);
  
  const [formData, setFormData] = useState({
    user_id: '',
    application_id: '',
    module_id: '',
    restriction_type: 'block' as RestrictionType,
    max_permission_level: '' as PermissionLevel | '',
    reason: '',
    expires_at: '',
  });

  // 1. Carga Inicial de Opciones Globales (Usuarios y Apps)
  useEffect(() => {
    if (open) {
      fn_get_restriction_form_options()
        .then(res => {
          setUsers(res.users);
          setApps(res.applications);
        })
        .catch(() => toast.error('Error cargando listas de opciones'));

      // Configurar estado inicial del formulario
      if (restriction) {
        setFormData({
          user_id: restriction.user_id,
          application_id: restriction.application_id,
          module_id: restriction.module_id,
          restriction_type: restriction.restriction_type,
          max_permission_level: restriction.max_permission_level || '',
          reason: restriction.reason || '',
          expires_at: restriction.expires_at ? new Date(restriction.expires_at).toISOString().slice(0, 16) : '',
        });
        // Cargar módulos de la app de la restricción
        loadModules(restriction.application_id);
      } else {
        setFormData({
          user_id: '',
          application_id: defaultApplicationId || '',
          module_id: '',
          restriction_type: 'block',
          max_permission_level: '',
          reason: '',
          expires_at: '',
        });
        
        // Si hay una app por defecto, cargar sus módulos
        if (defaultApplicationId) {
          loadModules(defaultApplicationId);
        } else {
          setModules([]);
        }
      }
    }
  }, [open, restriction, defaultApplicationId]);

  // Función para cargar módulos replicando tu referencia
  const loadModules = async (appId: string) => {
    if (!appId) {
      setModules([]);
      return;
    }
    try {
      setLoadingModules(true);
      // CORRECCIÓN: Usamos page_size 200 y el filtro exacto { application_id }
      const res = await fn_get_modules(1, 200, { application_id: appId });
      setModules(res.data);
    } catch (error) {
      console.error(error);
      toast.error('Error al cargar módulos');
    } finally {
      setLoadingModules(false);
    }
  };

  const handleAppChange = (val: string) => {
    setFormData(prev => ({ ...prev, application_id: val, module_id: '' }));
    loadModules(val);
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);

    try {
      const payload: any = {
        restriction_type: formData.restriction_type,
        reason: formData.reason,
        expires_at: formData.expires_at ? new Date(formData.expires_at).toISOString() : undefined,
      };

      if (formData.restriction_type === 'limit' && formData.max_permission_level) {
        payload.max_permission_level = formData.max_permission_level;
      }

      if (restriction) {
        await fn_update_user_restriction(restriction.id, payload);
        toast.success('Restricción actualizada');
      } else {
        await fn_create_user_restriction({
          ...payload,
          user_id: formData.user_id,
          application_id: formData.application_id,
          module_id: formData.module_id,
        });
        toast.success('Restricción creada');
      }
      onSuccess();
      onOpenChange(false);
    } catch (error: any) {
      toast.error(error.message || 'Error al guardar');
    } finally {
      setLoading(false);
    }
  };
  

  const isFormValid = formData.user_id && formData.application_id && formData.module_id;

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="sm:max-w-[500px] bg-card/80 backdrop-blur-xl border-border">
        <DialogHeader>
          <DialogTitle>{restriction ? 'Editar Restricción' : 'Nueva Restricción'}</DialogTitle>
          <DialogDescription>Define el nivel de acceso y bloqueo para el usuario.</DialogDescription>
        </DialogHeader>
        <form onSubmit={handleSubmit} className="space-y-4">
          
          <div className="grid grid-cols-2 gap-4">
            <div className="space-y-2">
              <Label>Aplicación <span className="text-destructive">*</span></Label>
              <Select 
                value={formData.application_id} 
                onValueChange={handleAppChange} 
                disabled={!!restriction || !!defaultApplicationId}
              >
                <SelectTrigger>
                  <SelectValue placeholder="Seleccionar..." />
                </SelectTrigger>
                <SelectContent>
                  {apps.map(a => <SelectItem key={a.id} value={a.id}>{a.name}</SelectItem>)}
                </SelectContent>
              </Select>
            </div>
            <div className="space-y-2">
              <Label>Usuario <span className="text-destructive">*</span></Label>
              <Select 
                value={formData.user_id} 
                onValueChange={v => setFormData({...formData, user_id: v})} 
                disabled={!!restriction}
              >
                <SelectTrigger>
                  <SelectValue placeholder={users.length ? "Seleccionar..." : "Cargando..."} />
                </SelectTrigger>
                <SelectContent>
                  {users.map(u => (
                    <SelectItem key={u.id} value={u.id}>
                      {u.full_name}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
            </div>
          </div>

          <div className="space-y-2">
            <Label>Módulo <span className="text-destructive">*</span></Label>
            <Select 
              value={formData.module_id} 
              onValueChange={v => setFormData({...formData, module_id: v})} 
              disabled={!!restriction || !formData.application_id || loadingModules}
            >
              <SelectTrigger>
                <SelectValue placeholder={loadingModules ? "Cargando módulos..." : "Seleccionar módulo..."} />
              </SelectTrigger>
              <SelectContent>
                {modules.length === 0 && !loadingModules ? (
                  <SelectItem value="none" disabled>No hay módulos disponibles</SelectItem>
                ) : (
                  modules.map(m => (
                    <SelectItem key={m.id} value={m.id}>
                      {/* CORRECCIÓN: Mostramos Nombre y Ruta (más identificable) */}
                      {m.name} ({m.route})
                    </SelectItem>
                  ))
                )}
              </SelectContent>
            </Select>
          </div>

          <div className="grid grid-cols-2 gap-4">
            <div className="space-y-2">
              <Label>Tipo de Restricción <span className="text-destructive">*</span></Label>
              <Select value={formData.restriction_type} onValueChange={(v: RestrictionType) => setFormData({...formData, restriction_type: v})}>
                <SelectTrigger><SelectValue /></SelectTrigger>
                <SelectContent>
                  <SelectItem value="block">Bloqueo Total (Block)</SelectItem>
                  <SelectItem value="read_only">Solo Lectura</SelectItem>
                  <SelectItem value="limit">Limitar Permisos</SelectItem>
                </SelectContent>
              </Select>
            </div>
            
            {formData.restriction_type === 'limit' && (
              <div className="space-y-2">
                <Label>Nivel Máximo <span className="text-destructive">*</span></Label>
                <Select value={formData.max_permission_level} onValueChange={(v: PermissionLevel) => setFormData({...formData, max_permission_level: v})}>
                  <SelectTrigger><SelectValue placeholder="Nivel..." /></SelectTrigger>
                  <SelectContent>
                    <SelectItem value="read">Leer</SelectItem>
                    <SelectItem value="write">Escribir</SelectItem>
                    <SelectItem value="execute">Ejecutar</SelectItem>
                  </SelectContent>
                </Select>
              </div>
            )}
          </div>

          <div className="space-y-2">
            <Label>Expiración (Opcional)</Label>
            <Input 
              type="datetime-local" 
              value={formData.expires_at} 
              onChange={e => setFormData({...formData, expires_at: e.target.value})} 
            />
          </div>

          <div className="space-y-2">
            <Label>Razón / Motivo</Label>
            <Textarea 
              value={formData.reason} 
              onChange={e => setFormData({...formData, reason: e.target.value})} 
              placeholder="Explica por qué se aplica esta restricción..."
            />
          </div>

          <DialogFooter>
            <Button type="button" variant="outline" onClick={() => onOpenChange(false)} disabled={loading}>Cancelar</Button>
            <Button type="submit" className="bg-linear-to-r from-primary to-chart-1" disabled={loading || !isFormValid}>
              {loading ? 'Guardando...' : 'Guardar'}
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  );
}