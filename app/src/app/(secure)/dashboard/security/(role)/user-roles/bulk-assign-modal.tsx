'use client';

import { useState, useEffect } from 'react';
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle } from '@/components/ui/dialog';
import { Button } from '@/components/ui/button';
import { Label } from '@/components/ui/label';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { ScrollArea } from '@/components/ui/scroll-area';
import { Checkbox } from '@/components/ui/checkbox';
import { Input } from '@/components/ui/input';
import { toast } from 'sonner';
import { Search } from 'lucide-react';
import { fn_bulk_assign_role_to_users, BulkAssignRoleToUsersInput } from '@/actions/user-application-roles/fn_bulk_assign_role_to_users';
import { fn_get_assignment_form_options, AssignmentFormOptions } from '@/actions/user-application-roles/fn_get_assignment_form_options';

interface BulkAssignModalProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  onSuccess: () => void;
  defaultApplicationId?: string;
}

export default function BulkAssignModal({ open, onOpenChange, onSuccess, defaultApplicationId }: BulkAssignModalProps) {
  const [loading, setLoading] = useState(false);
  const [loadingOptions, setLoadingOptions] = useState(false);
  const [options, setOptions] = useState<AssignmentFormOptions>({
    applications: [],
    users: [],
    roles: [],
  });

  const [applicationId, setApplicationId] = useState('');
  const [roleId, setRoleId] = useState('');
  const [selectedUsers, setSelectedUsers] = useState<Set<string>>(new Set());
  const [userSearch, setUserSearch] = useState('');

  // CORRECCIÓN: Sincronizar estado inicial al abrir
  useEffect(() => {
    if (open) {
      if (defaultApplicationId) {
        setApplicationId(defaultApplicationId);
        loadOptions(defaultApplicationId);
      } else {
        loadOptions();
      }
    }
  }, [open, defaultApplicationId]);

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

  const handleApplicationChange = (val: string) => {
    setApplicationId(val); 
    setRoleId('');
    loadOptions(val);
  };

  const toggleUser = (userId: string) => {
    const newSelected = new Set(selectedUsers);
    if (newSelected.has(userId)) {
      newSelected.delete(userId);
    } else {
      newSelected.add(userId);
    }
    setSelectedUsers(newSelected);
  };

  const handleSelectAllFiltered = () => {
    const newSelected = new Set(selectedUsers);
    filteredUsers.forEach(u => newSelected.add(u.id));
    setSelectedUsers(newSelected);
  };

  const handleDeselectAll = () => {
    setSelectedUsers(new Set());
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);

    try {
      const payload: BulkAssignRoleToUsersInput = {
        application_id: applicationId,
        application_role_id: roleId,
        user_ids: Array.from(selectedUsers),
      };

      const result = await fn_bulk_assign_role_to_users(payload);
      toast.success(`Proceso completado: ${result.created} asignados, ${result.failed} fallidos`);
      onSuccess();
      onOpenChange(false);
      setSelectedUsers(new Set());
    } catch (error: any) {
      toast.error(error.message || 'Error en asignación masiva');
    } finally {
      setLoading(false);
    }
  };

  const filteredUsers = options.users.filter(u => 
    u.full_name.toLowerCase().includes(userSearch.toLowerCase()) || 
    u.email.toLowerCase().includes(userSearch.toLowerCase())
  );

  const isFormValid = applicationId && roleId && selectedUsers.size > 0;

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="bg-card/80 backdrop-blur-xl border-border sm:max-w-[600px]">
        <DialogHeader>
          <DialogTitle>Asignación Masiva de Rol</DialogTitle>
          <DialogDescription>
            Asigna un rol específico a múltiples usuarios simultáneamente.
          </DialogDescription>
        </DialogHeader>

        <form onSubmit={handleSubmit} className="space-y-4">
          <div className="grid grid-cols-2 gap-4">
            {/* Aplicación */}
            <div className="space-y-2">
              <Label>Aplicación <span className="text-destructive">*</span></Label>
              <Select
                value={applicationId}
                onValueChange={handleApplicationChange}
                disabled={loadingOptions || !!defaultApplicationId}
              >
                <SelectTrigger>
                  <SelectValue placeholder="Seleccionar app" />
                </SelectTrigger>
                <SelectContent>
                  {options.applications.map((app) => (
                    <SelectItem key={app.id} value={app.id}>{app.name}</SelectItem>
                  ))}
                </SelectContent>
              </Select>
            </div>

            {/* Rol */}
            <div className="space-y-2">
              <Label>Rol <span className="text-destructive">*</span></Label>
              <Select
                value={roleId}
                onValueChange={setRoleId}
                disabled={loadingOptions || !applicationId}
              >
                <SelectTrigger>
                  <SelectValue placeholder="Seleccionar rol" />
                </SelectTrigger>
                <SelectContent>
                  {options.roles.map((role) => (
                    <SelectItem key={role.id} value={role.id}>{role.name}</SelectItem>
                  ))}
                </SelectContent>
              </Select>
            </div>
          </div>

          {/* Selección de Usuarios */}
          <div className="space-y-2">
            <div className="flex justify-between items-center">
              <Label>Seleccionar Usuarios ({selectedUsers.size}) <span className="text-destructive">*</span></Label>
              <div className="flex gap-2">
                <Button type="button" variant="ghost" size="sm" onClick={handleSelectAllFiltered} className="h-6 text-xs">
                  Todos
                </Button>
                <Button type="button" variant="ghost" size="sm" onClick={handleDeselectAll} className="h-6 text-xs">
                  Ninguno
                </Button>
              </div>
            </div>
            
            <div className="relative">
              <Search className="absolute left-2 top-2.5 h-4 w-4 text-muted-foreground" />
              <Input 
                placeholder="Buscar usuarios..." 
                value={userSearch}
                onChange={(e) => setUserSearch(e.target.value)}
                className="pl-8"
              />
            </div>

            <ScrollArea className="h-[200px] border rounded-md p-2">
              <div className="space-y-2">
                {filteredUsers.length === 0 ? (
                  <p className="text-sm text-center text-muted-foreground py-4">No se encontraron usuarios</p>
                ) : (
                  filteredUsers.map((user) => (
                    <div key={user.id} className="flex items-center space-x-2 p-1 hover:bg-muted/50 rounded">
                      <Checkbox 
                        id={`user-${user.id}`} 
                        checked={selectedUsers.has(user.id)}
                        onCheckedChange={() => toggleUser(user.id)}
                      />
                      <Label 
                        htmlFor={`user-${user.id}`} 
                        className="flex-1 cursor-pointer text-sm font-normal"
                      >
                        <span className="font-medium">{user.full_name}</span>
                        <span className="text-muted-foreground ml-2 text-xs">{user.email}</span>
                      </Label>
                    </div>
                  ))
                )}
              </div>
            </ScrollArea>
          </div>

          <DialogFooter className="mt-4">
            <Button type="button" variant="outline" onClick={() => onOpenChange(false)} disabled={loading}>
              Cancelar
            </Button>
            <Button type="submit" className="bg-linear-to-r from-primary to-chart-1" disabled={loading || !isFormValid}>
              {loading ? 'Procesando...' : `Asignar Rol a ${selectedUsers.size} Usuarios`}
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  );
}