'use client';

import { useState, useEffect } from 'react';
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle } from '@/components/ui/dialog';
import { Button } from '@/components/ui/button';
import { Label } from '@/components/ui/label';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { toast } from 'sonner';
import { fn_create_module_role_permission } from '@/actions/module-role-permissions/fn_create_module_role_permission';
import { fn_update_module_role_permission } from '@/actions/module-role-permissions/fn_update_module_role_permission';
import { fn_get_modules } from '@/actions/modules/fn_get_modules';
import { fn_get_application_roles } from '@/actions/application-roles/fn_get_application_roles';
import type { ModuleRolePermissionItem, PermissionType, CreateModuleRolePermissionInput } from '@/types/module-role-permissions';
import type { ModuleItem } from '@/types/modules';
import type { ApplicationRoleItem } from '@/types/application-roles';

interface PermissionModalProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  permission?: ModuleRolePermissionItem | null;
  onSuccess: () => void;
  defaultApplicationId?: string;
  defaultRoleId?: string;
}

const PERMISSION_TYPES: { value: PermissionType; label: string; description: string }[] = [
  { value: 'read', label: 'Lectura', description: 'Permite ver contenido' },
  { value: 'write', label: 'Escritura', description: 'Permite crear y editar' },
  { value: 'execute', label: 'Ejecución', description: 'Permite ejecutar acciones' },
  { value: 'delete', label: 'Eliminación', description: 'Permite eliminar contenido' },
  { value: 'admin', label: 'Administrador', description: 'Acceso completo' },
];

export default function PermissionModal({ 
  open, 
  onOpenChange, 
  permission, 
  onSuccess, 
  defaultApplicationId,
  defaultRoleId 
}: PermissionModalProps) {
  const [loading, setLoading] = useState(false);
  const [loadingModules, setLoadingModules] = useState(false);
  const [loadingRoles, setLoadingRoles] = useState(false);
  const [modules, setModules] = useState<ModuleItem[]>([]);
  const [roles, setRoles] = useState<ApplicationRoleItem[]>([]);
  const [formData, setFormData] = useState<CreateModuleRolePermissionInput>({
    module_id: '',
    application_role_id: '',
    permission_type: 'read',
  });

  useEffect(() => {
    if (open) {
      if (defaultApplicationId) {
        loadModules(defaultApplicationId);
        loadRoles(defaultApplicationId);
      }
    }
  }, [open, defaultApplicationId]);

  useEffect(() => {
    if (permission) {
      setFormData({
        module_id: permission.module_id,
        application_role_id: permission.application_role_id,
        permission_type: permission.permission_type,
      });
    } else {
      setFormData({
        module_id: '',
        application_role_id: defaultRoleId || '',
        permission_type: 'read',
      });
    }
  }, [permission, open, defaultRoleId]);

  const loadModules = async (applicationId: string) => {
    try {
      setLoadingModules(true);
      const response = await fn_get_modules(1, 100, { application_id: applicationId, is_deleted: false });
      setModules(response.data);
    } catch (error) {
      console.error('Error loading modules:', error);
      toast.error('Error al cargar módulos');
    } finally {
      setLoadingModules(false);
    }
  };

  const loadRoles = async (applicationId: string) => {
    try {
      setLoadingRoles(true);
      const response = await fn_get_application_roles(1, 100, { application_id: applicationId, is_deleted: false });
      setRoles(response.data);
    } catch (error) {
      console.error('Error loading roles:', error);
      toast.error('Error al cargar roles');
    } finally {
      setLoadingRoles(false);
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);

    try {
      if (permission) {
        await fn_update_module_role_permission(permission.id, {
          permission_type: formData.permission_type,
        });
        toast.success('Permiso actualizado correctamente');
      } else {
        await fn_create_module_role_permission(formData);
        toast.success('Permiso creado correctamente');
      }
      onSuccess();
      onOpenChange(false);
    } catch (error: any) {
      toast.error(error.message || 'Error al guardar el permiso');
    } finally {
      setLoading(false);
    }
  };

  const isFormValid = 
    formData.module_id.trim() !== '' && 
    formData.application_role_id.trim() !== '' &&
    formData.permission_type.trim() !== '';

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="bg-card/80 backdrop-blur-xl border-border sm:max-w-[550px]">
        <DialogHeader>
          <DialogTitle>{permission ? 'Editar Permiso' : 'Crear Nuevo Permiso'}</DialogTitle>
          <DialogDescription>
            {permission ? 'Actualiza el tipo de permiso.' : 'Asigna un permiso a un módulo para un rol específico.'}
          </DialogDescription>
        </DialogHeader>

        <form onSubmit={handleSubmit} className="space-y-4">
          {/* Rol */}
          <div className="space-y-2">
            <Label htmlFor="application_role_id">
              Rol <span className="text-destructive">*</span>
            </Label>
            <Select
              value={formData.application_role_id || ' '}
              onValueChange={(value) => setFormData({ ...formData, application_role_id: value === ' ' ? '' : value })}
              disabled={loadingRoles || !!defaultRoleId || !!permission}
            >
              <SelectTrigger className="[&>span]:truncate">
                <SelectValue placeholder="Seleccionar rol" />
              </SelectTrigger>
              <SelectContent position="popper" sideOffset={5} className="max-h-[300px]">
                <SelectItem value=" " disabled>Selecciona un rol</SelectItem>
                {roles.map((role) => (
                  <SelectItem key={role.id} value={role.id}>
                    {role.name}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
            {(defaultRoleId || permission) && (
              <p className="text-xs text-muted-foreground">
                El rol no se puede cambiar
              </p>
            )}
          </div>

          {/* Módulo */}
          <div className="space-y-2">
            <Label htmlFor="module_id">
              Módulo <span className="text-destructive">*</span>
            </Label>
            <Select
              value={formData.module_id || ' '}
              onValueChange={(value) => setFormData({ ...formData, module_id: value === ' ' ? '' : value })}
              disabled={loadingModules || !!permission}
            >
              <SelectTrigger className="[&>span]:truncate">
                <SelectValue placeholder="Seleccionar módulo" />
              </SelectTrigger>
              <SelectContent position="popper" sideOffset={5} className="max-h-[300px]">
                <SelectItem value=" " disabled>Selecciona un módulo</SelectItem>
                {modules.map((module) => (
                  <SelectItem key={module.id} value={module.id} title={module.route}>
                    {module.name}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
            {permission && (
              <p className="text-xs text-muted-foreground">
                El módulo no se puede cambiar
              </p>
            )}
          </div>

          {/* Tipo de Permiso */}
          <div className="space-y-2">
            <Label htmlFor="permission_type">
              Tipo de Permiso <span className="text-destructive">*</span>
            </Label>
            <Select
              value={formData.permission_type}
              onValueChange={(value) => setFormData({ ...formData, permission_type: value as PermissionType })}
            >
              <SelectTrigger>
                <SelectValue placeholder="Seleccionar tipo" />
              </SelectTrigger>
              <SelectContent>
                {PERMISSION_TYPES.map((type) => (
                  <SelectItem key={type.value} value={type.value}>
                    <div className="flex flex-col">
                      <span className="font-medium">{type.label}</span>
                      <span className="text-xs text-muted-foreground">{type.description}</span>
                    </div>
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
          </div>

          <div className="bg-muted/50 p-4 rounded-lg">
            <p className="text-sm font-medium mb-2">Tipos de Permiso:</p>
            <ul className="text-sm text-muted-foreground space-y-1">
              <li>• <strong>Lectura:</strong> Ver contenido del módulo</li>
              <li>• <strong>Escritura:</strong> Crear y editar contenido</li>
              <li>• <strong>Ejecución:</strong> Ejecutar acciones específicas</li>
              <li>• <strong>Eliminación:</strong> Eliminar contenido</li>
              <li>• <strong>Administrador:</strong> Acceso completo al módulo</li>
            </ul>
          </div>

          <DialogFooter>
            <Button type="button" variant="outline" onClick={() => onOpenChange(false)} disabled={loading}>
              Cancelar
            </Button>
            <Button type="submit" className="bg-linear-to-r from-primary to-chart-1" disabled={loading || !isFormValid}>
              {loading ? 'Guardando...' : permission ? 'Actualizar Permiso' : 'Crear Permiso'}
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  );
}