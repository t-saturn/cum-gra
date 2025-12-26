'use client';

import { useState, useEffect } from 'react';
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle } from '@/components/ui/dialog';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { Textarea } from '@/components/ui/textarea';
import { Switch } from '@/components/ui/switch';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { toast } from 'sonner';
import { fn_create_organic_unit, CreateOrganicUnitInput } from '@/actions/units/fn_create_organic_unit';
import { fn_update_organic_unit, UpdateOrganicUnitInput } from '@/actions/units/fn_update_organic_unit';
import { fn_get_all_organic_units, OrganicUnitSelectItem } from '@/actions/units/fn_get_all_organic_units';
import type { OrganicUnitItemDTO } from '@/types/units';

interface OrganicUnitModalProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  unit?: OrganicUnitItemDTO | null;
  onSuccess: () => void;
}

export default function OrganicUnitModal({ open, onOpenChange, unit, onSuccess }: OrganicUnitModalProps) {
  const [loading, setLoading] = useState(false);
  const [loadingParents, setLoadingParents] = useState(false);
  const [parentUnits, setParentUnits] = useState<OrganicUnitSelectItem[]>([]);
  const [formData, setFormData] = useState<CreateOrganicUnitInput>({
    name: '',
    acronym: '',
    brand: '',
    description: '',
    parent_id: '',
    is_active: true,
    cod_dep_sgd: '',
  });

  // Cargar unidades padres disponibles
  useEffect(() => {
    const loadParentUnits = async () => {
      try {
        setLoadingParents(true);
        const units = await fn_get_all_organic_units(false);
        // Filtrar la unidad actual si estamos editando
        const filtered = unit ? units.filter((u) => u.id !== unit.id) : units;
        setParentUnits(filtered);
      } catch (err) {
        console.error('Error al cargar unidades padre:', err);
      } finally {
        setLoadingParents(false);
      }
    };

    if (open) {
      loadParentUnits();
    }
  }, [open, unit]);

  useEffect(() => {
    if (unit) {
      setFormData({
        name: unit.name,
        acronym: unit.acronym,
        brand: unit.brand ?? '',
        description: unit.description ?? '',
        parent_id: unit.parent_id ?? '',
        is_active: unit.is_active,
        cod_dep_sgd: unit.cod_dep_sgd ?? '',
      });
    } else {
      setFormData({
        name: '',
        acronym: '',
        brand: '',
        description: '',
        parent_id: '',
        is_active: true,
        cod_dep_sgd: '',
      });
    }
  }, [unit, open]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);

    try {
      // Limpiar parent_id vacío
      const payload = {
        ...formData,
        parent_id: formData.parent_id || undefined,
        brand: formData.brand || undefined,
        description: formData.description || undefined,
        cod_dep_sgd: formData.cod_dep_sgd || undefined,
      };

      if (unit) {
        await fn_update_organic_unit(unit.id, payload as UpdateOrganicUnitInput);
        toast.success('Unidad orgánica actualizada correctamente');
      } else {
        await fn_create_organic_unit(payload);
        toast.success('Unidad orgánica creada correctamente');
      }
      onSuccess();
      onOpenChange(false);
    } catch (error: any) {
      toast.error(error.message || 'Error al guardar la unidad orgánica');
    } finally {
      setLoading(false);
    }
  };

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="bg-card/80 backdrop-blur-xl border-border sm:max-w-[700px] max-h-[90vh] overflow-y-auto">
        <DialogHeader>
          <DialogTitle>{unit ? 'Editar Unidad Orgánica' : 'Crear Nueva Unidad Orgánica'}</DialogTitle>
          <DialogDescription>
            {unit ? 'Actualiza la información de la unidad organizacional.' : 'Define una nueva unidad en la estructura organizacional.'}
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
                placeholder="Ej: Gerencia de Tecnología"
                required
              />
            </div>

            <div className="space-y-2">
              <Label htmlFor="acronym">
                Acrónimo <span className="text-destructive">*</span>
              </Label>
              <Input
                id="acronym"
                value={formData.acronym}
                onChange={(e) => setFormData({ ...formData, acronym: e.target.value })}
                placeholder="Ej: GTI"
                required
              />
            </div>
          </div>

          <div className="grid grid-cols-2 gap-4">
            <div className="space-y-2">
              <Label htmlFor="brand">Marca</Label>
              <Input
                id="brand"
                value={formData.brand}
                onChange={(e) => setFormData({ ...formData, brand: e.target.value })}
                placeholder="Ej: TI Ayacucho"
              />
            </div>

            <div className="space-y-2">
              <Label htmlFor="cod_dep_sgd">Código SGD</Label>
              <Input
                id="cod_dep_sgd"
                value={formData.cod_dep_sgd}
                onChange={(e) => setFormData({ ...formData, cod_dep_sgd: e.target.value })}
                placeholder="Ej: GTI01"
              />
            </div>
          </div>

          <div className="space-y-2">
            <Label htmlFor="parent_id">Unidad Padre (Opcional)</Label>
            <Select
              value={formData.parent_id || 'none'}
              onValueChange={(value) => setFormData({ ...formData, parent_id: value === 'none' ? '' : value })}
              disabled={loadingParents}
            >
              <SelectTrigger>
                <SelectValue placeholder={loadingParents ? 'Cargando...' : 'Seleccionar unidad padre'} />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="none">Sin unidad padre (raíz)</SelectItem>
                {parentUnits.map((u) => (
                  <SelectItem key={u.id} value={u.id}>
                    {u.name} ({u.acronym})
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
            <p className="text-xs text-muted-foreground">Deja vacío para crear una unidad raíz</p>
          </div>

          <div className="space-y-2">
            <Label htmlFor="description">Descripción</Label>
            <Textarea
              id="description"
              value={formData.description}
              onChange={(e) => setFormData({ ...formData, description: e.target.value })}
              placeholder="Describe las funciones y responsabilidades"
              rows={3}
            />
          </div>

          <div className="flex items-center justify-between rounded-lg border p-4">
            <div className="space-y-0.5">
              <Label htmlFor="is_active">Estado Activo</Label>
              <p className="text-muted-foreground text-sm">La unidad estará disponible para asignación</p>
            </div>
            <Switch id="is_active" checked={formData.is_active} onCheckedChange={(checked) => setFormData({ ...formData, is_active: checked })} />
          </div>

          <DialogFooter>
            <Button type="button" variant="outline" onClick={() => onOpenChange(false)} disabled={loading}>
              Cancelar
            </Button>
            <Button type="submit" className="bg-linear-to-r from-primary to-chart-1" disabled={loading}>
              {loading ? 'Guardando...' : unit ? 'Actualizar Unidad' : 'Crear Unidad'}
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  );
}