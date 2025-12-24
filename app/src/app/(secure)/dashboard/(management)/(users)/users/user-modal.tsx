'use client';

import { useState, useEffect } from 'react';
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle } from '@/components/ui/dialog';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { toast } from 'sonner';
import { Loader2, Eye, EyeOff } from 'lucide-react';

import { fn_create_user, CreateUserInput } from '@/actions/users/fn_create_user';
import { fn_update_user, UpdateUserInput } from '@/actions/users/fn_update_user';
import { fn_get_user_form_options, UserFormOptions } from '@/actions/users/fn_get_user_form_options';
import type { UserItem } from '@/types/users';

import type { DepartmentItem, ProvinceItem, DistrictItem } from '@/types/ubigeos';
import { fn_get_departments } from '@/actions/ubigeos/fn_get_departments';
import { fn_get_provinces } from '@/actions/ubigeos/fn_get_provinces';
import { fn_get_districts } from '@/actions/ubigeos/fn_get_districts';

interface UserModalProps {
    open: boolean;
    onOpenChange: (open: boolean) => void;
    user?: UserItem | null;
    onSuccess: () => void;
}

export default function UserModal({ open, onOpenChange, user, onSuccess }: UserModalProps) {
    const [loading, setLoading] = useState(false);
    const [isHydrating, setIsHydrating] = useState(false);
    const [showPassword, setShowPassword] = useState(false);
    
    const [options, setOptions] = useState<UserFormOptions>({
        positions: [],
        units: [],
    });

    const initialFormState: CreateUserInput = {
        email: '',
        dni: '',
        password: '', // NUEVO campo requerido
        first_name: '',
        last_name: '',
        phone: '',
        status: 'active',
        cod_emp_sgd: '',
        structural_position_id: '',
        organic_unit_id: '',
        ubigeo_id: '',
    };

    const [formData, setFormData] = useState<CreateUserInput>(initialFormState);

    const [departments, setDepartments] = useState<DepartmentItem[]>([]);
    const [provinces, setProvinces] = useState<ProvinceItem[]>([]);
    const [districts, setDistricts] = useState<DistrictItem[]>([]);

    const [selectedDepartment, setSelectedDepartment] = useState('');
    const [selectedProvince, setSelectedProvince] = useState('');

    useEffect(() => {
        if (!open) {
            setFormData(initialFormState);
            setSelectedDepartment('');
            setSelectedProvince('');
            setProvinces([]);
            setDistricts([]);
            setShowPassword(false);
            return;
        }

        const initModal = async () => {
            setIsHydrating(true);
            try {
                const [formOptions, depts] = await Promise.all([
                    fn_get_user_form_options(),
                    fn_get_departments()
                ]);
                
                setOptions(formOptions);
                setDepartments(depts);

                if (user) {
                    const existingUbigeoId = user.ubigeo?.id ? String(user.ubigeo.id) : '';
                    setFormData({
                        email: user.email || '',
                        dni: user.dni || '',
                        password: user.dni || '', // Por defecto usar DNI como contraseña
                        first_name: user.first_name || '',
                        last_name: user.last_name || '',
                        phone: user.phone || '',
                        status: user.status || 'active',
                        cod_emp_sgd: user.cod_emp_sgd || '',
                        structural_position_id: user.structural_position_id || '',
                        organic_unit_id: user.organic_unit_id || '',
                        ubigeo_id: existingUbigeoId || '',
                    });

                    if (user.ubigeo) {
                        const { department, province } = user.ubigeo;

                        setSelectedDepartment(department);
                        
                        const loadedProvinces = await fn_get_provinces(department);
                        setProvinces(loadedProvinces);
                        setSelectedProvince(province);

                        const loadedDistricts = await fn_get_districts(department, province);
                        setDistricts(loadedDistricts);
                    }
                }
            } catch (error) {
                console.error("Error en initModal:", error);
                toast.error("Error al cargar datos iniciales");
            } finally {
                setIsHydrating(false);
            }
        };

        initModal();
    }, [open, user]);

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        setLoading(true);

        try {
            if (user) {
                // Modo edición - NO se envía password
                const { dni, password, ...updatePayload } = formData;
                await fn_update_user(user.id, {
                    ...updatePayload,
                    keycloak_id: user.keycloak_id ?? undefined,
                } as UpdateUserInput);
                toast.success('Usuario actualizado correctamente');
            } else {
                // Modo creación - SE REQUIERE password
                if (!formData.password || formData.password.trim() === '') {
                    toast.error('La contraseña es requerida');
                    setLoading(false);
                    return;
                }
                
                await fn_create_user({
                    email: formData.email,
                    dni: formData.dni,
                    password: formData.password, // NUEVO campo requerido
                    first_name: formData.first_name,
                    last_name: formData.last_name,
                    phone: formData.phone || undefined,
                    status: formData.status,
                    cod_emp_sgd: formData.cod_emp_sgd || undefined,
                    structural_position_id: formData.structural_position_id || undefined,
                    organic_unit_id: formData.organic_unit_id || undefined,
                    ubigeo_id: formData.ubigeo_id || undefined,
                });
                toast.success('Usuario creado correctamente en el sistema y Keycloak');
            }
            onSuccess();
            onOpenChange(false);
        } catch (error: any) {
            toast.error(error.message || 'Error al guardar el usuario');
        } finally {
            setLoading(false);
        }
    };

    const validateDNI = (dni: string) => /^\d{8}$/.test(dni);

    const handleDNIChange = (value: string) => {
        const cleaned = value.replace(/\D/g, '').slice(0, 8);
        setFormData({ ...formData, dni: cleaned });
    };

    const handleDepartmentChange = async (department: string) => {
        setSelectedDepartment(department);
        setSelectedProvince('');
        setFormData({ ...formData, ubigeo_id: '' });
        setProvinces([]);
        setDistricts([]);

        if (department && department !== ' ') {
            try {
                const provs = await fn_get_provinces(department);
                setProvinces(provs);
            } catch (error) {
                toast.error('Error al cargar provincias');
            }
        }
    };

    const handleProvinceChange = async (province: string) => {
        setSelectedProvince(province);
        setFormData({ ...formData, ubigeo_id: '' });
        setDistricts([]);

        if (selectedDepartment && province && province !== ' ') {
            try {
                const dists = await fn_get_districts(selectedDepartment, province);
                setDistricts(dists);
            } catch (error) {
                toast.error('Error al cargar distritos');
            }
        }
    };

    const handleDistrictChange = (ubigeo_id: string) => {
        setFormData({ ...formData, ubigeo_id: ubigeo_id === ' ' ? '' : ubigeo_id });
    };

    const isFormValid = (formData.email || '').trim() !== '' && 
                        validateDNI(formData.dni || '') && 
                        (formData.first_name || '').trim() !== '' && 
                        (formData.last_name || '').trim() !== '' &&
                        (user || (formData.password || '').trim() !== ''); // Password solo requerido en creación

    return (
        <Dialog open={open} onOpenChange={onOpenChange}>
            <DialogContent className="bg-card/80 backdrop-blur-xl border-border sm:max-w-[700px] max-h-[90vh] overflow-y-auto">
                <DialogHeader>
                    <DialogTitle>{user ? 'Editar Usuario' : 'Crear Nuevo Usuario'}</DialogTitle>
                    <DialogDescription>
                        {user ? 'Actualiza la información del usuario.' : 'Registra un nuevo usuario en el sistema y Keycloak automáticamente.'}
                    </DialogDescription>
                </DialogHeader>

                {isHydrating ? (
                    <div className="flex flex-col items-center justify-center py-10 space-y-4">
                        <Loader2 className="h-8 w-8 animate-spin text-primary" />
                        <p className="text-sm text-muted-foreground">Cargando información del usuario...</p>
                    </div>
                ) : (
                    <form onSubmit={handleSubmit} className="space-y-4">
                        <div className="grid grid-cols-2 gap-4">
                            <div className="space-y-2">
                                <Label htmlFor="first_name">Nombre <span className="text-destructive">*</span></Label>
                                <Input id="first_name" value={formData.first_name || ''} onChange={(e) => setFormData({ ...formData, first_name: e.target.value })} placeholder="Ej: Juan Carlos" required />
                            </div>
                            <div className="space-y-2">
                                <Label htmlFor="last_name">Apellidos <span className="text-destructive">*</span></Label>
                                <Input id="last_name" value={formData.last_name || ''} onChange={(e) => setFormData({ ...formData, last_name: e.target.value })} placeholder="Ej: Pérez García" required />
                            </div>
                        </div>

                        <div className="grid grid-cols-2 gap-4">
                            <div className="space-y-2">
                                <Label htmlFor="dni">DNI <span className="text-destructive">*</span></Label>
                                <Input id="dni" value={formData.dni || ''} onChange={(e) => handleDNIChange(e.target.value)} placeholder="12345678" maxLength={8} required disabled={!!user} />
                                <p className="text-xs text-muted-foreground">{(formData.dni || '').length}/8 dígitos</p>
                            </div>
                            <div className="space-y-2">
                                <Label htmlFor="email">Email <span className="text-destructive">*</span></Label>
                                <Input id="email" type="email" value={formData.email || ''} onChange={(e) => setFormData({ ...formData, email: e.target.value })} placeholder="usuario@regionayacucho.gob.pe" required />
                            </div>
                        </div>

                        {/* Campo de contraseña - SOLO en modo creación */}
                        {!user && (
                            <div className="space-y-2">
                                <Label htmlFor="password">Contraseña <span className="text-destructive">*</span></Label>
                                <div className="relative">
                                    <Input 
                                        id="password" 
                                        type={showPassword ? "text" : "password"}
                                        value={formData.password || ''} 
                                        onChange={(e) => setFormData({ ...formData, password: e.target.value })} 
                                        placeholder="Ingrese contraseña" 
                                        required 
                                        className="pr-10"
                                    />
                                    <Button
                                        type="button"
                                        variant="ghost"
                                        size="sm"
                                        className="absolute right-0 top-0 h-full px-3 py-2 hover:bg-transparent"
                                        onClick={() => setShowPassword(!showPassword)}
                                    >
                                        {showPassword ? (
                                            <EyeOff className="h-4 w-4 text-muted-foreground" />
                                        ) : (
                                            <Eye className="h-4 w-4 text-muted-foreground" />
                                        )}
                                    </Button>
                                </div>
                                <p className="text-xs text-muted-foreground">
                                    Mínimo 4 caracteres. Si no se ingresa, se usará el DNI como contraseña.
                                </p>
                            </div>
                        )}

                        <div className="grid grid-cols-2 gap-4">
                            <div className="space-y-2">
                                <Label htmlFor="phone">Teléfono</Label>
                                <Input id="phone" value={formData.phone || ''} onChange={(e) => setFormData({ ...formData, phone: e.target.value })} placeholder="+51987654321" />
                            </div>
                            <div className="space-y-2">
                                <Label htmlFor="status">Estado <span className="text-destructive">*</span></Label>
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
                        </div>

                        <div className="grid grid-cols-2 gap-4">
                            <div className="space-y-2">
                                <Label htmlFor="organic_unit_id">Unidad Orgánica</Label>
                                <Select
                                    value={formData.organic_unit_id || ' '}
                                    onValueChange={(value) => setFormData({ ...formData, organic_unit_id: value === ' ' ? '' : value })}
                                >
                                    <SelectTrigger className="w-full">
                                        <SelectValue placeholder="Seleccionar unidad">
                                            {formData.organic_unit_id && formData.organic_unit_id !== ' ' ? (
                                                <span className="truncate block">
                                                    {options.units.find(u => u.id === formData.organic_unit_id)?.name}
                                                </span>
                                            ) : ('Seleccionar unidad')}
                                        </SelectValue>
                                    </SelectTrigger>
                                    <SelectContent className="max-h-[300px]">
                                        <SelectItem value=" ">Sin asignar</SelectItem>
                                        {options.units.map((unit) => (
                                            <SelectItem key={unit.id} value={unit.id}>
                                                <div className="truncate" title={unit.name}>{unit.name} ({unit.acronym})</div>
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
                                >
                                    <SelectTrigger className="w-full">
                                        <SelectValue placeholder="Seleccionar posición">
                                            {formData.structural_position_id && formData.structural_position_id !== ' ' ? (
                                                <span className="truncate block">
                                                    {options.positions.find(p => p.id === formData.structural_position_id)?.name}
                                                </span>
                                            ) : ('Seleccionar posición')}
                                        </SelectValue>
                                    </SelectTrigger>
                                    <SelectContent className="max-h-[300px]">
                                        <SelectItem value=" ">Sin asignar</SelectItem>
                                        {options.positions.map((position) => (
                                            <SelectItem key={position.id} value={position.id}>
                                                <div className="truncate" title={position.name}>{position.name} - {position.code}</div>
                                            </SelectItem>
                                        ))}
                                    </SelectContent>
                                </Select>
                            </div>
                        </div>

                        <div className="space-y-2">
                            <Label>Ubicación Geográfica</Label>
                            <div className="grid grid-cols-3 gap-2">
                                {/* DEPARTAMENTO */}
                                <div className="space-y-1">
                                    <Label htmlFor="department" className="text-xs text-muted-foreground">Departamento</Label>
                                    <Select 
                                        value={selectedDepartment || ' '} 
                                        onValueChange={handleDepartmentChange}
                                        disabled={isHydrating}
                                    >
                                        <SelectTrigger id="department">
                                            <SelectValue placeholder="Seleccionar" />
                                        </SelectTrigger>
                                        <SelectContent>
                                            <SelectItem value=" ">Sin asignar</SelectItem>
                                            {departments.map((dept) => (
                                                <SelectItem key={dept.name} value={dept.name}>{dept.name}</SelectItem>
                                            ))}
                                        </SelectContent>
                                    </Select>
                                </div>

                                {/* PROVINCIA */}
                                <div className="space-y-1">
                                    <Label htmlFor="province" className="text-xs text-muted-foreground">Provincia</Label>
                                    <Select
                                        key={`prov-select-${provinces.length}`} 
                                        value={selectedProvince || ' '} 
                                        onValueChange={handleProvinceChange} 
                                        disabled={!selectedDepartment || selectedDepartment === ' ' || isHydrating}
                                    >
                                        <SelectTrigger id="province">
                                            <SelectValue placeholder="Seleccionar" />
                                        </SelectTrigger>
                                        <SelectContent>
                                            <SelectItem value=" ">Sin asignar</SelectItem>
                                            {provinces.map((prov) => (
                                                <SelectItem key={prov.name} value={prov.name}>{prov.name}</SelectItem>
                                            ))}
                                        </SelectContent>
                                    </Select>
                                </div>

                                {/* DISTRITO */}
                                <div className="space-y-1">
                                    <Label htmlFor="district" className="text-xs text-muted-foreground">Distrito</Label>
                                    <Select
                                        key={`dist-select-${districts.length}`}
                                        value={formData.ubigeo_id || ' '} 
                                        onValueChange={handleDistrictChange} 
                                        disabled={!selectedProvince || selectedProvince === ' ' || isHydrating}
                                    >
                                        <SelectTrigger id="district">
                                            <SelectValue placeholder="Seleccionar" />
                                        </SelectTrigger>
                                        <SelectContent>
                                            <SelectItem value=" ">Sin asignar</SelectItem>
                                            {districts.map((d) => (
                                                <SelectItem key={d.id} value={d.id}>{d.name}</SelectItem>
                                            ))}
                                        </SelectContent>
                                    </Select>
                                </div>
                            </div>
                            
                            {!isHydrating && selectedDepartment && selectedProvince && formData.ubigeo_id && (
                                <p className="text-xs text-muted-foreground mt-1">
                                    Seleccionado: {selectedDepartment} / {selectedProvince} / {districts.find(d => d.id === formData.ubigeo_id)?.name}
                                </p>
                            )}
                        </div>

                        {!user && (
                            <div className="bg-muted/50 p-4 rounded-lg">
                                <p className="text-sm font-medium mb-2">Creación automática en Keycloak:</p>
                                <ul className="text-sm text-muted-foreground space-y-1">
                                    <li>✓ Usuario se creará automáticamente en Keycloak y en la base de datos</li>
                                    <li>✓ El DNI se usará como username en Keycloak</li>
                                    <li>✓ La contraseña ingresada será permanente (no temporal)</li>
                                    <li>✓ El email se marcará como verificado</li>
                                    <li>✓ El estado se sincronizará automáticamente</li>
                                </ul>
                            </div>
                        )}

                        <DialogFooter>
                            <Button type="button" variant="outline" onClick={() => onOpenChange(false)} disabled={loading}>
                                Cancelar
                            </Button>
                            <Button type="submit" className="bg-linear-to-r from-primary to-chart-1" disabled={loading || !isFormValid || isHydrating}>
                                {loading ? 'Guardando...' : user ? 'Actualizar Usuario' : 'Crear Usuario'}
                            </Button>
                        </DialogFooter>
                    </form>
                )}
            </DialogContent>
        </Dialog>
    );
}