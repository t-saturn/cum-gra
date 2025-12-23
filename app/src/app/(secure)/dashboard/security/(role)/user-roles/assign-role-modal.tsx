'use client';

import { useState, useEffect } from 'react';
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle } from '@/components/ui/dialog';
import { Button } from '@/components/ui/button';
import { Label } from '@/components/ui/label';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { toast } from 'sonner';
import { fn_assign_role, AssignRoleInput } from '@/actions/user-application-roles/fn_assign_role';
import { fn_get_assignment_form_options, AssignmentFormOptions } from '@/actions/user-application-roles/fn_get_assignment_form_options';

interface AssignRoleModalProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  onSuccess: () => void;
  defaultApplicationId?: string;
}

export default function AssignRoleModal({ open, onOpenChange, onSuccess, defaultApplicationId }: AssignRoleModalProps) {
  const [loading, setLoading] = useState(false);
  const [loadingOptions, setLoadingOptions] = useState(false);
  const [options, setOptions] = useState<AssignmentFormOptions>({
    applications: [],
    users: [],
    roles: [],
  });

  const [formData, setFormData] = useState<AssignRoleInput>({
    user_id: '',
    application_id: '',
    application_role_id: '',
  });

  // CORRECCIÓN: Sincronizar el estado cuando se abre el modal con la app seleccionada
  useEffect(() => {
    if (open) {
      if (defaultApplicationId) {
        setFormData(prev => ({ ...prev, application_id: defaultApplicationId }));
        loadOptions(defaultApplicationId);
      } else {
        loadOptions();
      }
    }
  }, [open, defaultApplicationId]);

  // Recargar opciones si el usuario cambia manualmente la aplicación dentro del modal
  const handleApplicationChange = (value: string) => {
    setFormData({ ...formData, application_id: value, application_role_id: '' });
    loadOptions(value);
  };

  const loadOptions = async (appId?: string) => {
    try {
      setLoadingOptions(true);
      const data = await fn_get_assignment_form_options(appId);
      setOptions(prev => ({
        ...data,
        users: data.users.length ? data.users : prev.users,
        applications: data.applications.length ? data.applications : prev.applications,
      }));
    } catch (error) {
      console.error('Error loading options:', error);
      toast.error('Error al cargar opciones');
    } finally {
      setLoadingOptions(false);
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);

    try {
      await fn_assign_role(formData);
      toast.success('Rol asignado correctamente');
      onSuccess();
      onOpenChange(false);
      // Reset parcial (mantener app seleccionada por comodidad)
      setFormData(prev => ({ ...prev, user_id: '', application_role_id: '' }));
    } catch (error: any) {
      toast.error(error.message || 'Error al asignar rol');
    } finally {
      setLoading(false);
    }
  };

  const isFormValid = 
    formData.user_id && 
    formData.application_id && 
    formData.application_role_id;

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="bg-card/80 backdrop-blur-xl border-border sm:max-w-[500px]">
        <DialogHeader>
          <DialogTitle>Asignar Nuevo Rol</DialogTitle>
          <DialogDescription>
            Otorga permisos a un usuario asignándole un rol específico en una aplicación.
          </DialogDescription>
        </DialogHeader>

        <form onSubmit={handleSubmit} className="space-y-4">
          {/* Aplicación */}
          <div className="space-y-2">
            <Label htmlFor="application_id">Aplicación <span className="text-destructive">*</span></Label>
            <Select
              value={formData.application_id}
              onValueChange={handleApplicationChange}
              disabled={loadingOptions || !!defaultApplicationId}
            >
              <SelectTrigger>
                <SelectValue placeholder="Seleccionar aplicación" />
              </SelectTrigger>
              <SelectContent>
                {options.applications.map((app) => (
                  <SelectItem key={app.id} value={app.id}>{app.name}</SelectItem>
                ))}
              </SelectContent>
            </Select>
          </div>

          {/* Usuario */}
          <div className="space-y-2">
            <Label htmlFor="user_id">Usuario <span className="text-destructive">*</span></Label>
            <Select
              value={formData.user_id}
              onValueChange={(value) => setFormData({ ...formData, user_id: value })}
              disabled={loadingOptions}
            >
              <SelectTrigger>
                <SelectValue placeholder="Seleccionar usuario" />
              </SelectTrigger>
              <SelectContent className="max-h-[300px]">
                {options.users.map((user) => (
                  <SelectItem key={user.id} value={user.id}>
                    {user.full_name} ({user.email})
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
          </div>

          {/* Rol */}
          <div className="space-y-2">
            <Label htmlFor="role_id">Rol <span className="text-destructive">*</span></Label>
            <Select
              value={formData.application_role_id}
              onValueChange={(value) => setFormData({ ...formData, application_role_id: value })}
              disabled={loadingOptions || !formData.application_id}
            >
              <SelectTrigger>
                <SelectValue placeholder={formData.application_id ? "Seleccionar rol" : "Seleccione una aplicación primero"} />
              </SelectTrigger>
              <SelectContent>
                {options.roles.length === 0 ? (
                  <SelectItem value=" " disabled>No hay roles disponibles</SelectItem>
                ) : (
                  options.roles.map((role) => (
                    <SelectItem key={role.id} value={role.id}>{role.name}</SelectItem>
                  ))
                )}
              </SelectContent>
            </Select>
          </div>

          <DialogFooter className="mt-6">
            <Button type="button" variant="outline" onClick={() => onOpenChange(false)} disabled={loading}>
              Cancelar
            </Button>
            <Button type="submit" className="bg-linear-to-r from-primary to-chart-1" disabled={loading || !isFormValid}>
              {loading ? 'Asignando...' : 'Asignar Rol'}
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  );
}