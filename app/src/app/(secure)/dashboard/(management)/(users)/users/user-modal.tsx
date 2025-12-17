'use client';

import { useState, useEffect } from 'react';
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle } from '@/components/ui/dialog';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { toast } from 'sonner';
import { fn_create_user, CreateUserInput } from '@/actions/users/fn_create_user';
import { fn_update_user, UpdateUserInput } from '@/actions/users/fn_update_user';
import { fn_get_user_form_options, UserFormOptions } from '@/actions/users/fn_get_user_form_options';
import type { UserItem } from '@/types/users';

interface UserModalProps {
    open: boolean;
    onOpenChange: (open: boolean) => void;
    user?: UserItem | null;
    onSuccess: () => void;
}

export default function UserModal({ open, onOpenChange, user, onSuccess }: UserModalProps) {
    const [loading, setLoading] = useState(false);
    const [loadingOptions, setLoadingOptions] = useState(false);
    const [options, setOptions] = useState<UserFormOptions>({
        positions: [],
        units: [],
        ubigeos: [],
    });
    const [formData, setFormData] = useState<CreateUserInput>({
        email: '',
        dni: '',
        first_name: '',
        last_name: '',
        phone: '',
        status: 'active',
        cod_emp_sgd: '',
        structural_position_id: '',
        organic_unit_id: '',
        ubigeo_id: '',
        sync_to_keycloak: true,
    });

    // Estados para ubigeo en cascada
    const [selectedDepartment, setSelectedDepartment] = useState('');
    const [selectedProvince, setSelectedProvince] = useState('');
    const [filteredProvinces, setFilteredProvinces] = useState<typeof options.ubigeos>([]);
    const [filteredDistricts, setFilteredDistricts] = useState<typeof options.ubigeos>([]);

    useEffect(() => {
        if (open) {
            loadFormOptions();
        }
    }, [open]);

    useEffect(() => {
        if (user) {
            setFormData({
                email: user.email || '',
                dni: user.dni || '',
                first_name: user.first_name || '',
                last_name: user.last_name || '',
                phone: user.phone || '',
                status: user.status || 'active',
                cod_emp_sgd: user.cod_emp_sgd || '',
                structural_position_id: user.structural_position_id || '',
                organic_unit_id: user.organic_unit_id || '',
                ubigeo_id: user.ubigeo_id || '',
                sync_to_keycloak: true,
            });
        } else {
            setFormData({
                email: '',
                dni: '',
                first_name: '',
                last_name: '',
                phone: '',
                status: 'active',
                cod_emp_sgd: '',
                structural_position_id: '',
                organic_unit_id: '',
                ubigeo_id: '',
                sync_to_keycloak: true,
            });
            setSelectedDepartment('');
            setSelectedProvince('');
            setFilteredProvinces([]);
            setFilteredDistricts([]);
        }
    }, [user, open]);

    // Cargar ubigeo cuando editas
    useEffect(() => {
        if (user && user.ubigeo_id && options.ubigeos.length > 0) {
            const ubigeo = options.ubigeos.find((u) => u.id === user.ubigeo_id);
            if (ubigeo) {
                setSelectedDepartment(ubigeo.department);
                setSelectedProvince(ubigeo.province);

                const provinces = options.ubigeos.filter((u) => u.department === ubigeo.department);
                setFilteredProvinces(provinces);

                const districts = options.ubigeos.filter((u) => u.department === ubigeo.department && u.province === ubigeo.province);
                setFilteredDistricts(districts);
            }
        }
    }, [user, options.ubigeos]);

    const loadFormOptions = async () => {
        try {
            setLoadingOptions(true);
            const data = await fn_get_user_form_options();
            setOptions(data);
        } catch (error) {
            console.error('Error loading form options:', error);
            toast.error('Error al cargar opciones del formulario');
        } finally {
            setLoadingOptions(false);
        }
    };

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        setLoading(true);

        try {
            const payload = {
                email: formData.email,
                dni: formData.dni,
                first_name: formData.first_name,
                last_name: formData.last_name,
                phone: formData.phone || undefined,
                status: formData.status,
                cod_emp_sgd: formData.cod_emp_sgd || undefined,
                structural_position_id: formData.structural_position_id || undefined,
                organic_unit_id: formData.organic_unit_id || undefined,
                ubigeo_id: formData.ubigeo_id || undefined,
                sync_to_keycloak: true,
            };

            if (user) {
                const { dni, ...updatePayload } = payload;
                await fn_update_user(user.id, {
                    ...updatePayload,
                    keycloak_id: user.keycloak_id ?? undefined,
                } as UpdateUserInput);
                toast.success('Usuario actualizado en backend y Keycloak');
            } else {
                await fn_create_user(payload);
                toast.success('Usuario creado en backend y Keycloak');
            }
            onSuccess();
            onOpenChange(false);
        } catch (error: any) {
            toast.error(error.message || 'Error al guardar el usuario');
        } finally {
            setLoading(false);
        }
    };

    const validateDNI = (dni: string) => {
        return /^\d{8}$/.test(dni);
    };

    const handleDNIChange = (value: string) => {
        const cleaned = value.replace(/\D/g, '').slice(0, 8);
        setFormData({ ...formData, dni: cleaned });
    };

    const handleDepartmentChange = (department: string) => {
        if (department === ' ') {
            setSelectedDepartment('');
            setSelectedProvince('');
            setFormData({ ...formData, ubigeo_id: '' });
            setFilteredProvinces([]);
            setFilteredDistricts([]);
            return;
        }

        setSelectedDepartment(department);
        setSelectedProvince('');
        setFormData({ ...formData, ubigeo_id: '' });

        const provinces = options.ubigeos.filter((u) => u.department === department);
        setFilteredProvinces(provinces);
        setFilteredDistricts([]);
    };

    const handleProvinceChange = (province: string) => {
        if (province === ' ') {
            setSelectedProvince('');
            setFormData({ ...formData, ubigeo_id: '' });
            setFilteredDistricts([]);
            return;
        }

        setSelectedProvince(province);
        setFormData({ ...formData, ubigeo_id: '' });

        if (selectedDepartment) {
            const districts = options.ubigeos.filter((u) => u.department === selectedDepartment && u.province === province);
            setFilteredDistricts(districts);
        }
    };

    const handleDistrictChange = (ubigeoId: string) => {
        setFormData({ ...formData, ubigeo_id: ubigeoId === ' ' ? '' : ubigeoId });
    };

    const departments = Array.from(new Set(options.ubigeos.map((u) => u.department))).sort();
    const provinces = Array.from(new Set(filteredProvinces.map((u) => u.province))).sort();

    const isFormValid = (formData.email || '').trim() !== '' && validateDNI(formData.dni || '') && (formData.first_name || '').trim() !== '' && (formData.last_name || '').trim() !== '';

    return (
        <Dialog open={open} onOpenChange={onOpenChange}>
            <DialogContent className="bg-card/80 backdrop-blur-xl border-border sm:max-w-[700px] max-h-[90vh] overflow-y-auto">
                <DialogHeader>
                    <DialogTitle>{user ? 'Editar Usuario' : 'Crear Nuevo Usuario'}</DialogTitle>
                    <DialogDescription>{user ? 'Actualiza la información del usuario.' : 'Registra un nuevo usuario en el sistema y Keycloak.'}</DialogDescription>
                </DialogHeader>

                <form onSubmit={handleSubmit} className="space-y-4">
                    <div className="grid grid-cols-2 gap-4">
                        <div className="space-y-2">
                            <Label htmlFor="first_name">
                                Nombre <span className="text-destructive">*</span>
                            </Label>
                            <Input id="first_name" value={formData.first_name || ''} onChange={(e) => setFormData({ ...formData, first_name: e.target.value })} placeholder="Ej: Juan Carlos" required />
                        </div>

                        <div className="space-y-2">
                            <Label htmlFor="last_name">
                                Apellidos <span className="text-destructive">*</span>
                            </Label>
                            <Input id="last_name" value={formData.last_name || ''} onChange={(e) => setFormData({ ...formData, last_name: e.target.value })} placeholder="Ej: Pérez García" required />
                        </div>
                    </div>

                    <div className="grid grid-cols-2 gap-4">
                        <div className="space-y-2">
                            <Label htmlFor="dni">
                                DNI <span className="text-destructive">*</span>
                            </Label>
                            <Input id="dni" value={formData.dni || ''} onChange={(e) => handleDNIChange(e.target.value)} placeholder="12345678" maxLength={8} required disabled={!!user} />
                            <p className="text-xs text-muted-foreground">{(formData.dni || '').length}/8 dígitos</p>
                            {user && <p className="text-xs text-muted-foreground">El DNI no se puede modificar</p>}
                        </div>

                        <div className="space-y-2">
                            <Label htmlFor="email">
                                Email <span className="text-destructive">*</span>
                            </Label>
                            <Input id="email" type="email" value={formData.email || ''} onChange={(e) => setFormData({ ...formData, email: e.target.value })} placeholder="usuario@regionayacucho.gob.pe" required />
                        </div>
                    </div>

                    <div className="grid grid-cols-2 gap-4">
                        <div className="space-y-2">
                            <Label htmlFor="phone">Teléfono</Label>
                            <Input id="phone" value={formData.phone || ''} onChange={(e) => setFormData({ ...formData, phone: e.target.value })} placeholder="+51987654321" />
                        </div>

                        <div className="space-y-2">
                            <Label htmlFor="status">
                                Estado <span className="text-destructive">*</span>
                            </Label>
                            <Select value={formData.status} onValueChange={(value: any) => setFormData({ ...formData, status: value })}>
                                <SelectTrigger>
                                    <SelectValue placeholder="Seleccionar estado" />
                                </SelectTrigger>
                                <SelectContent>
                                    <SelectItem value="active">Activo</SelectItem>
                                    <SelectItem value="suspended">Suspendido</SelectItem>
                                    <SelectItem value="inactive">Inactivo</SelectItem>
                                </SelectContent>
                            </Select>
                        </div>
                    </div>

                    <div className="space-y-2">
                        <Label htmlFor="cod_emp_sgd">Código Empleado SGD</Label>
                        <Input id="cod_emp_sgd" value={formData.cod_emp_sgd || ''} onChange={(e) => setFormData({ ...formData, cod_emp_sgd: e.target.value.slice(0, 5) })} placeholder="00001" maxLength={5} />
                        <p className="text-xs text-muted-foreground">{(formData.cod_emp_sgd || '').length}/5 caracteres (Sistema legacy)</p>
                    </div>

                    <div className="grid grid-cols-2 gap-4">
                        <div className="space-y-2">
                            <Label htmlFor="organic_unit_id">Unidad Orgánica</Label>
                            <Select
                                value={formData.organic_unit_id || ' '}
                                onValueChange={(value) => setFormData({ ...formData, organic_unit_id: value === ' ' ? '' : value })}
                                disabled={loadingOptions}
                            >
                                <SelectTrigger className="w-full">
                                    <SelectValue placeholder="Seleccionar unidad">
                                        {formData.organic_unit_id && formData.organic_unit_id !== ' ' ? (
                                            <span className="truncate block">
                                                {options.units.find(u => u.id === formData.organic_unit_id)?.name} (
                                                {options.units.find(u => u.id === formData.organic_unit_id)?.acronym})
                                            </span>
                                        ) : (
                                            'Seleccionar unidad'
                                        )}
                                    </SelectValue>
                                </SelectTrigger>
                                <SelectContent position="popper" sideOffset={5} className="max-h-[300px] w-(--radix-select-trigger-width)">
                                    <SelectItem value=" ">Sin asignar</SelectItem>
                                    {options.units.map((unit) => (
                                        <SelectItem key={unit.id} value={unit.id}>
                                            <div className="truncate" title={`${unit.name} (${unit.acronym})`}>
                                                {unit.name} ({unit.acronym})
                                            </div>
                                        </SelectItem>
                                    ))}
                                </SelectContent>
                            </Select>
                        </div>

                        <div className="space-y-2">
                            <Label htmlFor="structural_position_id">Posición Estructural</Label>
                            <Select
                                value={formData.structural_position_id || ' '}
                                onValueChange={(value) => setFormData({ ...formData, structural_position_id: value === ' ' ? '' : value })}
                                disabled={loadingOptions}
                            >
                                <SelectTrigger className="w-full">
                                    <SelectValue placeholder="Seleccionar posición">
                                        {formData.structural_position_id && formData.structural_position_id !== ' ' ? (
                                            <span className="truncate block">
                                                {options.positions.find(p => p.id === formData.structural_position_id)?.name} - {options.positions.find(p => p.id === formData.structural_position_id)?.code}
                                            </span>
                                        ) : (
                                            'Seleccionar posición'
                                        )}
                                    </SelectValue>
                                </SelectTrigger>
                                <SelectContent position="popper" sideOffset={5} className="max-h-[300px] w-(--radix-select-trigger-width)">
                                    <SelectItem value=" ">Sin asignar</SelectItem>
                                    {options.positions.map((position) => (
                                        <SelectItem key={position.id} value={position.id}>
                                            <div className="truncate" title={`${position.name} - ${position.code}`}>
                                                {position.name} - {position.code}
                                            </div>
                                        </SelectItem>
                                    ))}
                                </SelectContent>
                            </Select>
                        </div>
                    </div>

                    <div className="space-y-2">
                        <Label>Ubicación Geográfica</Label>
                        <div className="grid grid-cols-3 gap-2">
                            <div className="space-y-1">
                                <Label htmlFor="department" className="text-xs text-muted-foreground">
                                    Departamento
                                </Label>
                                <Select value={selectedDepartment || ' '} onValueChange={handleDepartmentChange} disabled={loadingOptions}>
                                    <SelectTrigger id="department">
                                        <SelectValue placeholder="Seleccionar" />
                                    </SelectTrigger>
                                    <SelectContent position="popper" sideOffset={5}>
                                        <SelectItem value=" ">Sin asignar</SelectItem>
                                        {departments.map((dept) => (
                                            <SelectItem key={dept} value={dept}>
                                                {dept}
                                            </SelectItem>
                                        ))}
                                    </SelectContent>
                                </Select>
                            </div>

                            <div className="space-y-1">
                                <Label htmlFor="province" className="text-xs text-muted-foreground">
                                    Provincia
                                </Label>
                                <Select value={selectedProvince || ' '} onValueChange={handleProvinceChange} disabled={loadingOptions || !selectedDepartment}>
                                    <SelectTrigger id="province">
                                        <SelectValue placeholder="Seleccionar" />
                                    </SelectTrigger>
                                    <SelectContent position="popper" sideOffset={5}>
                                        <SelectItem value=" ">Sin asignar</SelectItem>
                                        {provinces.map((prov) => (
                                            <SelectItem key={prov} value={prov}>
                                                {prov}
                                            </SelectItem>
                                        ))}
                                    </SelectContent>
                                </Select>
                            </div>

                            <div className="space-y-1">
                                <Label htmlFor="district" className="text-xs text-muted-foreground">
                                    Distrito
                                </Label>
                                <Select value={formData.ubigeo_id || ' '} onValueChange={handleDistrictChange} disabled={loadingOptions || !selectedProvince}>
                                    <SelectTrigger id="district">
                                        <SelectValue placeholder="Seleccionar" />
                                    </SelectTrigger>
                                    <SelectContent position="popper" sideOffset={5}>
                                        <SelectItem value=" ">Sin asignar</SelectItem>
                                        {filteredDistricts.map((ubigeo) => (
                                            <SelectItem key={ubigeo.id} value={ubigeo.id}>
                                                {ubigeo.district}
                                            </SelectItem>
                                        ))}
                                    </SelectContent>
                                </Select>
                            </div>
                        </div>
                        {selectedDepartment && selectedProvince && formData.ubigeo_id && (
                            <p className="text-xs text-muted-foreground">
                                Seleccionado: {selectedDepartment} / {selectedProvince} / {filteredDistricts.find((u) => u.id === formData.ubigeo_id)?.district}
                            </p>
                        )}
                    </div>

                    <div className="bg-muted/50 p-4 rounded-lg">
                        <p className="text-sm font-medium mb-2">Sincronización con Keycloak:</p>
                        <ul className="text-sm text-muted-foreground space-y-1">
                            <li>✓ Se creará/actualizará automáticamente en Keycloak</li>
                            <li>✓ El DNI se usará como username en Keycloak</li>
                            <li>✓ Se generará una contraseña temporal (cambio obligatorio en primer login)</li>
                            <li>✓ El estado se sincronizará (activo/suspendido/inactivo)</li>
                        </ul>
                    </div>

                    <DialogFooter>
                        <Button type="button" variant="outline" onClick={() => onOpenChange(false)} disabled={loading}>
                            Cancelar
                        </Button>
                        <Button type="submit" className="bg-linear-to-r from-primary to-chart-1" disabled={loading || !isFormValid || loadingOptions}>
                            {loading ? 'Guardando...' : user ? 'Actualizar Usuario' : 'Crear Usuario'}
                        </Button>
                    </DialogFooter>
                </form>
            </DialogContent>
        </Dialog>
    );
}