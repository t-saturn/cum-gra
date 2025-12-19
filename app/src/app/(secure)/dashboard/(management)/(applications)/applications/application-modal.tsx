'use client';

import { useState, useEffect } from 'react';
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle } from '@/components/ui/dialog';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { Textarea } from '@/components/ui/textarea';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { toast } from 'sonner';
import { fn_create_application, CreateApplicationInput } from '@/actions/applications/fn_create_application';
import { fn_update_application, UpdateApplicationInput } from '@/actions/applications/fn_update_application';
import type { ApplicationItem } from '@/types/applications';

interface ApplicationModalProps {
    open: boolean;
    onOpenChange: (open: boolean) => void;
    application?: ApplicationItem | null;
    onSuccess: () => void;
}

export default function ApplicationModal({ open, onOpenChange, application, onSuccess }: ApplicationModalProps) {
    const [loading, setLoading] = useState(false);
    const [formData, setFormData] = useState<CreateApplicationInput>({
        name: '',
        client_id: '',
        client_secret: '',
        domain: '',
        logo: '',
        description: '',
        status: 'development',
    });

    useEffect(() => {
        if (application) {
            setFormData({
                name: application.name,
                client_id: application.client_id,
                client_secret: '',
                domain: application.domain,
                logo: application.logo ?? '',
                description: application.description ?? '',
                status: application.status,
            });
        } else {
            setFormData({
                name: '',
                client_id: '',
                client_secret: '',
                domain: '',
                logo: '',
                description: '',
                status: 'development',
            });
        }
    }, [application, open]);

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        setLoading(true);

        try {
            const payload = {
                ...formData,
                client_secret: formData.client_secret || undefined,
                logo: formData.logo || undefined,
                description: formData.description || undefined,
                sync_to_keycloak: true, // Siempre sincronizar
            };

            if (application) {
                await fn_update_application(application.id, {
                    ...payload,
                    keycloak_id: application.keycloak_id, // Pasar keycloak_id para actualizar
                });
                toast.success('Aplicación actualizada en backend y Keycloak');
            } else {
                await fn_create_application(payload);
                toast.success('Aplicación creada en backend y Keycloak');
            }
            onSuccess();
            onOpenChange(false);
        } catch (error: any) {
            toast.error(error.message || 'Error al guardar la aplicación');
        } finally {
            setLoading(false);
        }
    };

    return (
        <Dialog open={open} onOpenChange={onOpenChange}>
            <DialogContent className="bg-card/80 backdrop-blur-xl border-border sm:max-w-[700px] max-h-[90vh] overflow-y-auto">
                <DialogHeader>
                    <DialogTitle>{application ? 'Editar Aplicación' : 'Crear Nueva Aplicación'}</DialogTitle>
                    <DialogDescription>
                        {application ? 'Actualiza la información de la aplicación.' : 'Registra una nueva aplicación en el sistema SSO.'}
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
                                placeholder="Ej: Sistema de Gestión Documental"
                                required
                            />
                        </div>

                        <div className="space-y-2">
                            <Label htmlFor="client_id">
                                Client ID <span className="text-destructive">*</span>
                            </Label>
                            <Input
                                id="client_id"
                                value={formData.client_id}
                                onChange={(e) => setFormData({ ...formData, client_id: e.target.value })}
                                placeholder="Ej: sgd-client"
                                required
                                disabled={!!application}
                            />
                            {application && (
                                <p className="text-xs text-muted-foreground">El Client ID no se puede modificar</p>
                            )}
                        </div>
                    </div>

                    <div className="space-y-2">
                        <Label htmlFor="domain">
                            Dominio <span className="text-destructive">*</span>
                        </Label>
                        <Input
                            id="domain"
                            value={formData.domain}
                            onChange={(e) => setFormData({ ...formData, domain: e.target.value })}
                            placeholder="https://app.regionayacucho.gob.pe"
                            required
                        />
                    </div>

                    <div className="grid grid-cols-2 gap-4">
                        <div className="space-y-2">
                            <Label htmlFor="client_secret">
                                Client Secret {!application && <span className="text-destructive">*</span>}
                            </Label>
                            <Input
                                id="client_secret"
                                type="password"
                                value={formData.client_secret}
                                onChange={(e) => setFormData({ ...formData, client_secret: e.target.value })}
                                placeholder={application ? 'Dejar vacío para no cambiar' : 'Contraseña secreta'}
                                required={!application}
                            />
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
                                    <SelectItem value="development">Desarrollo</SelectItem>
                                    <SelectItem value="active">Activo</SelectItem>
                                    <SelectItem value="inactive">Inactivo</SelectItem>
                                </SelectContent>
                            </Select>
                        </div>
                    </div>

                    <div className="space-y-2">
                        <Label htmlFor="logo">Logo URL</Label>
                        <Input
                            id="logo"
                            value={formData.logo}
                            onChange={(e) => setFormData({ ...formData, logo: e.target.value })}
                            placeholder="https://app.regionayacucho.gob.pe/logo.png"
                        />
                    </div>

                    <div className="space-y-2">
                        <Label htmlFor="description">Descripción</Label>
                        <Textarea
                            id="description"
                            value={formData.description}
                            onChange={(e) => setFormData({ ...formData, description: e.target.value })}
                            placeholder="Describe la funcionalidad de la aplicación"
                            rows={3}
                        />
                    </div>

                    <DialogFooter>
                        <Button type="button" variant="outline" onClick={() => onOpenChange(false)} disabled={loading}>
                            Cancelar
                        </Button>
                        <Button type="submit" className="bg-linear-to-r from-primary to-chart-1" disabled={loading}>
                            {loading ? 'Guardando...' : application ? 'Actualizar Aplicación' : 'Crear Aplicación'}
                        </Button>
                    </DialogFooter>
                </form>
            </DialogContent>
        </Dialog>
    );
}