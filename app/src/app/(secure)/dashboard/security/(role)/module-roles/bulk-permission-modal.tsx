'use client';

import { useState, useEffect } from 'react';
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle } from '@/components/ui/dialog';
import { Button } from '@/components/ui/button';
import { Label } from '@/components/ui/label';
import { Checkbox } from '@/components/ui/checkbox';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { ScrollArea } from '@/components/ui/scroll-area';
import { toast } from 'sonner';
import { fn_bulk_assign_permissions } from '@/actions/module-role-permissions/fn_bulk_assign_permissions';
import { fn_get_modules } from '@/actions/modules/fn_get_modules';
import type { PermissionType } from '@/types/module-role-permissions';
import type { ModuleItem } from '@/types/modules';
import { Package } from 'lucide-react';

interface BulkPermissionModalProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  onSuccess: () => void;
  applicationId: string;
  roleId: string;
  roleName: string;
}

const PERMISSION_TYPES: { value: PermissionType; label: string }[] = [
  { value: 'read', label: 'Lectura' },
  { value: 'write', label: 'Escritura' },
  { value: 'execute', label: 'Ejecución' },
  { value: 'delete', label: 'Eliminación' },
  { value: 'admin', label: 'Administrador' },
];

export default function BulkPermissionModal({ 
  open, 
  onOpenChange, 
  onSuccess, 
  applicationId,
  roleId,
  roleName
}: BulkPermissionModalProps) {
  const [loading, setLoading] = useState(false);
  const [loadingModules, setLoadingModules] = useState(false);
  const [modules, setModules] = useState<ModuleItem[]>([]);
  const [selectedModules, setSelectedModules] = useState<string[]>([]);
  const [permissionType, setPermissionType] = useState<PermissionType>('read');

  useEffect(() => {
    if (open && applicationId) {
      loadModules();
      setSelectedModules([]);
      setPermissionType('read');
    }
  }, [open, applicationId]);

  const loadModules = async () => {
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

  const handleToggleModule = (moduleId: string) => {
    setSelectedModules(prev => 
      prev.includes(moduleId) 
        ? prev.filter(id => id !== moduleId)
        : [...prev, moduleId]
    );
  };

  const handleSelectAll = () => {
    if (selectedModules.length === modules.length) {
      setSelectedModules([]);
    } else {
      setSelectedModules(modules.map(m => m.id));
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    
    if (selectedModules.length === 0) {
      toast.error('Selecciona al menos un módulo');
      return;
    }

    setLoading(true);

    try {
      const result = await fn_bulk_assign_permissions({
        application_role_id: roleId,
        module_ids: selectedModules,
        permission_type: permissionType,
      });
      
      toast.success(`${result.created} permisos creados, ${result.skipped} omitidos`);
      onSuccess();
      onOpenChange(false);
    } catch (error: any) {
      toast.error(error.message || 'Error en asignación masiva');
    } finally {
      setLoading(false);
    }
  };

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="bg-card/80 backdrop-blur-xl border-border sm:max-w-[600px]">
        <DialogHeader>
          <DialogTitle>Asignación Masiva de Permisos</DialogTitle>
          <DialogDescription>
            Asigna permisos a múltiples módulos para el rol <strong>{roleName}</strong>
          </DialogDescription>
        </DialogHeader>

        <form onSubmit={handleSubmit} className="space-y-4">
          {/* Tipo de Permiso */}
          <div className="space-y-2">
            <Label>Tipo de Permiso</Label>
            <Select value={permissionType} onValueChange={(v) => setPermissionType(v as PermissionType)}>
              <SelectTrigger>
                <SelectValue />
              </SelectTrigger>
              <SelectContent>
                {PERMISSION_TYPES.map((type) => (
                  <SelectItem key={type.value} value={type.value}>
                    {type.label}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
          </div>

          {/* Lista de Módulos */}
          <div className="space-y-2">
            <div className="flex items-center justify-between">
              <Label>Módulos ({selectedModules.length} seleccionados)</Label>
              <Button type="button" variant="ghost" size="sm" onClick={handleSelectAll}>
                {selectedModules.length === modules.length ? 'Deseleccionar todos' : 'Seleccionar todos'}
              </Button>
            </div>
            
            <ScrollArea className="h-[300px] border rounded-lg p-4">
              {loadingModules ? (
                <div className="text-center text-muted-foreground py-8">Cargando módulos...</div>
              ) : modules.length === 0 ? (
                <div className="text-center text-muted-foreground py-8">No hay módulos disponibles</div>
              ) : (
                <div className="space-y-2">
                  {modules.map((module) => (
                    <div 
                      key={module.id} 
                      className="flex items-center gap-3 p-3 rounded-lg hover:bg-muted/50 cursor-pointer"
                      onClick={() => handleToggleModule(module.id)}
                    >
                      <Checkbox 
                        checked={selectedModules.includes(module.id)}
                        onCheckedChange={() => handleToggleModule(module.id)}
                      />
                      <Package className="w-4 h-4 text-primary" />
                      <div className="flex-1">
                        <p className="font-medium text-sm">{module.name}</p>
                        <p className="text-xs text-muted-foreground">{module.route}</p>
                      </div>
                    </div>
                  ))}
                </div>
              )}
            </ScrollArea>
          </div>

          <div className="bg-muted/50 p-4 rounded-lg">
            <p className="text-sm text-muted-foreground">
              Los permisos duplicados serán omitidos automáticamente.
            </p>
          </div>

          <DialogFooter>
            <Button type="button" variant="outline" onClick={() => onOpenChange(false)} disabled={loading}>
              Cancelar
            </Button>
            <Button 
              type="submit" 
              className="bg-linear-to-r from-primary to-chart-1" 
              disabled={loading || selectedModules.length === 0}
            >
              {loading ? 'Asignando...' : `Asignar a ${selectedModules.length} módulos`}
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  );
}