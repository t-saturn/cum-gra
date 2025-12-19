'use client';

import { useState, useEffect } from 'react';
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle } from '@/components/ui/dialog';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { Textarea } from '@/components/ui/textarea';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { Switch } from '@/components/ui/switch';
import { toast } from 'sonner';
import { fn_create_position, CreatePositionInput } from '@/actions/positions/fn_create_position';
import { fn_update_position, UpdatePositionInput } from '@/actions/positions/fn_update_position';
import type { StructuralPositionItem } from '@/types/structural_positions';

interface PositionModalProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  position?: StructuralPositionItem | null;
  onSuccess: () => void;
}

export default function PositionModal({ open, onOpenChange, position, onSuccess }: PositionModalProps) {
  const [loading, setLoading] = useState(false);
  const [formData, setFormData] = useState<CreatePositionInput>({
    name: '',
    code: '',
    level: undefined,
    description: '',
    is_active: true,
    cod_car_sgd: '',
  });

  useEffect(() => {
    if (position) {
      setFormData({
        name: position.name,
        code: position.code,
        level: position.level ?? undefined,
        description: position.description ?? '',
        is_active: position.is_active,
        cod_car_sgd: position.cod_car_sgd ?? '',
      });
    } else {
      setFormData({
        name: '',
        code: '',
        level: undefined,
        description: '',
        is_active: true,
        cod_car_sgd: '',
      });
    }
  }, [position, open]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);

    try {
      if (position) {
        await fn_update_position(position.id, formData as UpdatePositionInput);
        toast.success('Posición actualizada correctamente');
      } else {
        await fn_create_position(formData);
        toast.success('Posición creada correctamente');
      }
      onSuccess();
      onOpenChange(false);
    } catch (error: any) {
      toast.error(error.message || 'Error al guardar la posición');
    } finally {
      setLoading(false);
    }
  };

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="bg-card/80 backdrop-blur-xl border-border sm:max-w-[700px] max-h-[90vh] overflow-y-auto">
        <DialogHeader>
          <DialogTitle>{position ? 'Editar Posición' : 'Crear Nueva Posición'}</DialogTitle>
          <DialogDescription>
            {position ? 'Actualiza la información de la posición estructural.' : 'Define una nueva posición en la estructura organizacional.'}
          </DialogDescription>
        </DialogHeader>

        <form onSubmit={handleSubmit} className="space-y-4">
          <div className="grid grid-cols-2 gap-4">
            <div className="space-y-2">
              <Label htmlFor="name">
                Nombre <span className="text-destructive">*</span>
              </Label>
              <Input
                id="name"
                value={formData.name}
                onChange={(e) => setFormData({ ...formData, name: e.target.value })}
                placeholder="Ej: Director Regional"
                required
              />
            </div>

            <div className="space-y-2">
              <Label htmlFor="code">
                Código <span className="text-destructive">*</span>
              </Label>
              <Input
                id="code"
                value={formData.code}
                onChange={(e) => setFormData({ ...formData, code: e.target.value })}
                placeholder="Ej: DIR-REG-001"
                required
              />
            </div>
          </div>

          <div className="grid grid-cols-2 gap-4">
            <div className="space-y-2">
              <Label htmlFor="level">Nivel Jerárquico</Label>
              <Select value={formData.level?.toString() || ''} onValueChange={(value) => setFormData({ ...formData, level: value ? parseInt(value) : undefined })}>
                <SelectTrigger>
                  <SelectValue placeholder="Seleccionar nivel" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="1">Nivel 1 - Directivo</SelectItem>
                  <SelectItem value="2">Nivel 2 - Ejecutivo</SelectItem>
                  <SelectItem value="3">Nivel 3 - Profesional</SelectItem>
                  <SelectItem value="4">Nivel 4 - Técnico</SelectItem>
                  <SelectItem value="5">Nivel 5 - Apoyo</SelectItem>
                </SelectContent>
              </Select>
            </div>

            <div className="space-y-2">
              <Label htmlFor="cod_car_sgd">Código SGD</Label>
              <Input
                id="cod_car_sgd"
                value={formData.cod_car_sgd}
                onChange={(e) => setFormData({ ...formData, cod_car_sgd: e.target.value })}
                placeholder="Ej: DR01"
              />
            </div>
          </div>

          <div className="space-y-2">
            <Label htmlFor="description">Descripción</Label>
            <Textarea
              id="description"
              value={formData.description}
              onChange={(e) => setFormData({ ...formData, description: e.target.value })}
              placeholder="Describe las responsabilidades y funciones"
              rows={3}
            />
          </div>

          <div className="flex items-center justify-between rounded-lg border p-4">
            <div className="space-y-0.5">
              <Label htmlFor="is_active">Estado Activo</Label>
              <p className="text-muted-foreground text-sm">La posición estará disponible para asignación</p>
            </div>
            <Switch id="is_active" checked={formData.is_active} onCheckedChange={(checked) => setFormData({ ...formData, is_active: checked })} />
          </div>

          <DialogFooter>
            <Button type="button" variant="outline" onClick={() => onOpenChange(false)} disabled={loading}>
              Cancelar
            </Button>
            <Button type="submit" className="bg-linear-to-r from-primary to-chart-1" disabled={loading}>
              {loading ? 'Guardando...' : position ? 'Actualizar Posición' : 'Crear Posición'}
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  );
}