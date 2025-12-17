'use client';

import { useState, useEffect } from 'react';
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle } from '@/components/ui/dialog';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { toast } from 'sonner';
import { fn_create_ubigeo, CreateUbigeoInput } from '@/actions/ubigeos/fn_create_ubigeo';
import { fn_update_ubigeo, UpdateUbigeoInput } from '@/actions/ubigeos/fn_update_ubigeo';
import type { UbigeoItem } from '@/types/ubigeos';

interface UbigeoModalProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  ubigeo?: UbigeoItem | null;
  onSuccess: () => void;
}

export default function UbigeoModal({ open, onOpenChange, ubigeo, onSuccess }: UbigeoModalProps) {
  const [loading, setLoading] = useState(false);
  const [formData, setFormData] = useState<CreateUbigeoInput>({
    ubigeo_code: '',
    inei_code: '',
    department: '',
    province: '',
    district: '',
  });

  useEffect(() => {
    if (ubigeo) {
      setFormData({
        ubigeo_code: ubigeo.ubigeo_code,
        inei_code: ubigeo.inei_code,
        department: ubigeo.department,
        province: ubigeo.province,
        district: ubigeo.district,
      });
    } else {
      setFormData({
        ubigeo_code: '',
        inei_code: '',
        department: '',
        province: '',
        district: '',
      });
    }
  }, [ubigeo, open]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);

    try {
      if (ubigeo) {
        await fn_update_ubigeo(ubigeo.id, formData as UpdateUbigeoInput);
        toast.success('Ubigeo actualizado correctamente');
      } else {
        await fn_create_ubigeo(formData);
        toast.success('Ubigeo creado correctamente');
      }
      onSuccess();
      onOpenChange(false);
    } catch (error: any) {
      toast.error(error.message || 'Error al guardar el ubigeo');
    } finally {
      setLoading(false);
    }
  };

  const validateUbigeoCode = (code: string) => {
    // Validar que sea exactamente 6 dígitos
    const regex = /^\d{6}$/;
    return regex.test(code);
  };

  const handleUbigeoCodeChange = (value: string) => {
    // Solo permitir números y máximo 6 caracteres
    const cleaned = value.replace(/\D/g, '').slice(0, 6);
    setFormData({ ...formData, ubigeo_code: cleaned });
  };

  const handleIneiCodeChange = (value: string) => {
    // Solo permitir números y máximo 6 caracteres
    const cleaned = value.replace(/\D/g, '').slice(0, 6);
    setFormData({ ...formData, inei_code: cleaned });
  };

  const isFormValid = 
    validateUbigeoCode(formData.ubigeo_code) &&
    validateUbigeoCode(formData.inei_code) &&
    formData.department.trim() !== '' &&
    formData.province.trim() !== '' &&
    formData.district.trim() !== '';

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="bg-card/80 backdrop-blur-xl border-border sm:max-w-[600px] max-h-[90vh] overflow-y-auto">
        <DialogHeader>
          <DialogTitle>{ubigeo ? 'Editar Ubigeo' : 'Crear Nuevo Ubigeo'}</DialogTitle>
          <DialogDescription>
            {ubigeo
              ? 'Actualiza la información del código de ubicación geográfica.'
              : 'Registra un nuevo código de ubicación geográfica del Perú.'}
          </DialogDescription>
        </DialogHeader>

        <form onSubmit={handleSubmit} className="space-y-4">
          <div className="grid grid-cols-2 gap-4">
            <div className="space-y-2">
              <Label htmlFor="ubigeo_code">
                Código Ubigeo <span className="text-destructive">*</span>
              </Label>
              <Input
                id="ubigeo_code"
                value={formData.ubigeo_code}
                onChange={(e) => handleUbigeoCodeChange(e.target.value)}
                placeholder="050101"
                maxLength={6}
                required
                disabled={!!ubigeo}
              />
              <p className="text-xs text-muted-foreground">
                {formData.ubigeo_code.length}/6 dígitos
              </p>
              {ubigeo && (
                <p className="text-xs text-muted-foreground">El código ubigeo no se puede modificar</p>
              )}
            </div>

            <div className="space-y-2">
              <Label htmlFor="inei_code">
                Código INEI <span className="text-destructive">*</span>
              </Label>
              <Input
                id="inei_code"
                value={formData.inei_code}
                onChange={(e) => handleIneiCodeChange(e.target.value)}
                placeholder="050101"
                maxLength={6}
                required
              />
              <p className="text-xs text-muted-foreground">
                {formData.inei_code.length}/6 dígitos
              </p>
            </div>
          </div>

          <div className="space-y-2">
            <Label htmlFor="department">
              Departamento <span className="text-destructive">*</span>
            </Label>
            <Input
              id="department"
              value={formData.department}
              onChange={(e) => setFormData({ ...formData, department: e.target.value })}
              placeholder="Ej: Ayacucho"
              required
            />
          </div>

          <div className="space-y-2">
            <Label htmlFor="province">
              Provincia <span className="text-destructive">*</span>
            </Label>
            <Input
              id="province"
              value={formData.province}
              onChange={(e) => setFormData({ ...formData, province: e.target.value })}
              placeholder="Ej: Huamanga"
              required
            />
          </div>

          <div className="space-y-2">
            <Label htmlFor="district">
              Distrito <span className="text-destructive">*</span>
            </Label>
            <Input
              id="district"
              value={formData.district}
              onChange={(e) => setFormData({ ...formData, district: e.target.value })}
              placeholder="Ej: Ayacucho"
              required
            />
          </div>

          <div className="bg-muted/50 p-4 rounded-lg">
            <p className="text-sm font-medium mb-2">Formato del Código Ubigeo:</p>
            <ul className="text-sm text-muted-foreground space-y-1">
              <li>• 6 dígitos en formato DDPPDD</li>
              <li>• DD (primeros 2): Código del departamento</li>
              <li>• PP (siguientes 2): Código de la provincia</li>
              <li>• DD (últimos 2): Código del distrito</li>
            </ul>
            <p className="text-xs text-muted-foreground mt-2">
              Ejemplo: <code className="bg-background px-1 rounded">050101</code> = Ayacucho (05) - Huamanga (01) - Ayacucho (01)
            </p>
          </div>

          <DialogFooter>
            <Button type="button" variant="outline" onClick={() => onOpenChange(false)} disabled={loading}>
              Cancelar
            </Button>
            <Button 
              type="submit" 
              className="bg-linear-to-r from-primary to-chart-1" 
              disabled={loading || !isFormValid}
            >
              {loading ? 'Guardando...' : ubigeo ? 'Actualizar Ubigeo' : 'Crear Ubigeo'}
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  );
}