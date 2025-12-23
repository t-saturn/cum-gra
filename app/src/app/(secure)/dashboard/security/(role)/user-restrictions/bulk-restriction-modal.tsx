'use client';

import { useState, useEffect } from 'react';
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle } from '@/components/ui/dialog';
import { Button } from '@/components/ui/button';
import { Label } from '@/components/ui/label';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { Checkbox } from '@/components/ui/checkbox';
import { ScrollArea } from '@/components/ui/scroll-area';
import { Textarea } from '@/components/ui/textarea';
import { toast } from 'sonner';
import { fn_get_modules } from '@/actions/modules/fn_get_modules';
import type { RestrictionType } from '@/types/user-restrictions';
import { fn_bulk_create_user_restrictions } from '@/actions/users-restrictions/fn_bulk_create_user_restrictions';
import { fn_get_restriction_form_options } from '@/actions/users-restrictions/fn_get_restriction_form_options';

interface BulkRestrictionModalProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  onSuccess: () => void;
  defaultApplicationId?: string;
}

export default function BulkRestrictionModal({ open, onOpenChange, onSuccess, defaultApplicationId }: BulkRestrictionModalProps) {
  const [loading, setLoading] = useState(false);
  const [apps, setApps] = useState<any[]>([]);
  const [users, setUsers] = useState<any[]>([]);
  const [modules, setModules] = useState<any[]>([]);
  
  const [appId, setAppId] = useState(defaultApplicationId || '');
  const [userId, setUserId] = useState('');
  const [selectedModules, setSelectedModules] = useState<Set<string>>(new Set());
  const [type, setType] = useState<RestrictionType>('block');
  const [reason, setReason] = useState('');

  useEffect(() => {
    if (open) {
      fn_get_restriction_form_options().then(res => {
        setApps(res.applications);
        setUsers(res.users);
      });
      if (defaultApplicationId) {
        setAppId(defaultApplicationId);
        loadModules(defaultApplicationId);
      }
    }
  }, [open, defaultApplicationId]);

  const loadModules = async (id: string) => {
    try {
      const res = await fn_get_modules(1, 200, { application_id: id });
      setModules(res.data);
    } catch {
      toast.error('Error cargando módulos');
    }
  };

  const handleAppChange = (val: string) => {
    setAppId(val);
    setSelectedModules(new Set());
    loadModules(val);
  };

  const toggleModule = (id: string) => {
    const next = new Set(selectedModules);
    if (next.has(id)) next.delete(id);
    else next.add(id);
    setSelectedModules(next);
  };

  const handleSelectAll = () => {
    if (selectedModules.size === modules.length) setSelectedModules(new Set());
    else setSelectedModules(new Set(modules.map(m => m.id)));
  };

  const handleSubmit = async () => {
    if (!appId || !userId || selectedModules.size === 0) return;
    setLoading(true);
    try {
      const res = await fn_bulk_create_user_restrictions({
        application_id: appId,
        user_id: userId,
        module_ids: Array.from(selectedModules),
        restriction_type: type,
        reason,
      });
      toast.success(`Creadas: ${res.created}, Omitidas: ${res.skipped}`);
      onSuccess();
      onOpenChange(false);
    } catch (err: any) {
      toast.error(err.message);
    } finally {
      setLoading(false);
    }
  };

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="sm:max-w-[600px]">
        <DialogHeader>
          <DialogTitle>Restricción Masiva</DialogTitle>
          <DialogDescription>Aplica la misma restricción a múltiples módulos de una aplicación.</DialogDescription>
        </DialogHeader>
        
        <div className="grid grid-cols-2 gap-4 py-4">
          <div className="space-y-2">
            <Label>Aplicación</Label>
            <Select value={appId} onValueChange={handleAppChange} disabled={!!defaultApplicationId}>
              <SelectTrigger><SelectValue placeholder="Seleccionar..." /></SelectTrigger>
              <SelectContent>{apps.map(a => <SelectItem key={a.id} value={a.id}>{a.name}</SelectItem>)}</SelectContent>
            </Select>
          </div>
          <div className="space-y-2">
            <Label>Usuario</Label>
            <Select value={userId} onValueChange={setUserId}>
              <SelectTrigger><SelectValue placeholder="Seleccionar..." /></SelectTrigger>
              <SelectContent>{users.map(u => <SelectItem key={u.id} value={u.id}>{u.full_name}</SelectItem>)}</SelectContent>
            </Select>
          </div>
        </div>

        <div className="space-y-2">
          <div className="flex justify-between items-center">
            <Label>Módulos ({selectedModules.size})</Label>
            <Button variant="ghost" size="sm" onClick={handleSelectAll} disabled={!modules.length}>
              {selectedModules.size === modules.length ? 'Deseleccionar' : 'Todos'}
            </Button>
          </div>
          <ScrollArea className="h-[200px] border rounded-md p-2">
            {modules.length === 0 ? <p className="text-sm text-muted-foreground text-center py-4">Selecciona una app primero</p> : 
              modules.map(m => (
                <div key={m.id} className="flex items-center space-x-2 py-1 hover:bg-accent rounded px-2">
                  <Checkbox id={m.id} checked={selectedModules.has(m.id)} onCheckedChange={() => toggleModule(m.id)} />
                  <Label htmlFor={m.id} className="flex-1 cursor-pointer">{m.name}</Label>
                </div>
              ))
            }
          </ScrollArea>
        </div>

        <div className="grid grid-cols-2 gap-4">
            <div className="space-y-2">
                <Label>Tipo</Label>
                <Select value={type} onValueChange={(v: RestrictionType) => setType(v)}>
                    <SelectTrigger><SelectValue /></SelectTrigger>
                    <SelectContent>
                        <SelectItem value="block">Bloquear</SelectItem>
                        <SelectItem value="read_only">Solo Lectura</SelectItem>
                    </SelectContent>
                </Select>
            </div>
            <div className="space-y-2">
                <Label>Razón</Label>
                <Textarea value={reason} onChange={e => setReason(e.target.value)} placeholder="Motivo..." className="h-10 min-h-[40px] resize-none" />
            </div>
        </div>

        <DialogFooter>
          <Button variant="outline" onClick={() => onOpenChange(false)}>Cancelar</Button>
          <Button onClick={handleSubmit} disabled={loading || !appId || !userId || !selectedModules.size}>
            {loading ? 'Procesando...' : 'Aplicar Restricciones'}
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}