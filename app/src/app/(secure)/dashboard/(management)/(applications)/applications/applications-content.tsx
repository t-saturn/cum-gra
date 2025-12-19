'use client';

import { useState, useEffect, useTransition } from 'react';
import { useRouter, useSearchParams } from 'next/navigation';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Badge } from '@/components/ui/badge';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table';
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuLabel, DropdownMenuSeparator, DropdownMenuTrigger } from '@/components/ui/dropdown-menu';
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle } from '@/components/ui/dialog';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { Search, Plus, Filter, Download, MoreHorizontal, Edit, Trash2, Eye, Globe, RefreshCw, AlertCircle, ChevronLeft, ChevronRight, Boxes } from 'lucide-react';
import { toast } from 'sonner';
import { fn_get_applications } from '@/actions/applications/fn_get_applications';
import { fn_delete_application } from '@/actions/applications/fn_delete_application';
import { fn_restore_application } from '@/actions/applications/fn_restore_application';
import { fn_get_keycloak_clients, KeycloakClientSimple } from '@/actions/applications/fn_get_keycloak_clients';
import type { ApplicationItem } from '@/types/applications';
import { ApplicationsStatsCards } from '@/components/custom/card/application-stats-card';
import ApplicationModal from './application-modal';
import SyncKeycloakModal from './sync-keycloak-modal';

export default function ApplicationsContent() {
    const router = useRouter();
    const searchParams = useSearchParams();
    const [isPending, startTransition] = useTransition();

    // Obtener params de URL
    const page = Number(searchParams.get('page')) || 1;
    const pageSize = Number(searchParams.get('page_size')) || 10;
    const showDeleted = searchParams.get('deleted') === 'true';

    const [applications, setApplications] = useState<ApplicationItem[]>([]);
    const [searchTerm, setSearchTerm] = useState('');
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);
    const [totalItems, setTotalItems] = useState(0);

    // Keycloak sync
    const [keycloakClients, setKeycloakClients] = useState<KeycloakClientSimple[]>([]);
    const [unsyncedClients, setUnsyncedClients] = useState<KeycloakClientSimple[]>([]);
    const [checkingSync, setCheckingSync] = useState(false);

    // Modals
    const [isCreateModalOpen, setIsCreateModalOpen] = useState(false);
    const [isEditModalOpen, setIsEditModalOpen] = useState(false);
    const [isDetailModalOpen, setIsDetailModalOpen] = useState(false);
    const [isDeleteDialogOpen, setIsDeleteDialogOpen] = useState(false);
    const [isSyncModalOpen, setIsSyncModalOpen] = useState(false);
    const [selectedApp, setSelectedApp] = useState<ApplicationItem | null>(null);

    const loadApplications = async () => {
        try {
            setLoading(true);
            setError(null);
            const response = await fn_get_applications(page, pageSize, showDeleted);
            setApplications(response.data);
            setTotalItems(response.total);
        } catch (err: any) {
            setError(err.message ?? 'Error desconocido');
            toast.error('Error al cargar aplicaciones');
        } finally {
            setLoading(false);
        }
    };

    const checkKeycloakSync = async () => {
        try {
            setCheckingSync(true);
            const clients = await fn_get_keycloak_clients();
            setKeycloakClients(clients);

            // Comparar con apps del backend
            const backendClientIds = new Set(applications.map((app) => app.client_id));
            const unsynced = clients.filter((kc) => !backendClientIds.has(kc.client_id));
            setUnsyncedClients(unsynced);

            if (unsynced.length > 0) {
                toast.info(`${unsynced.length} cliente(s) de Keycloak sin sincronizar`);
            }
        } catch (err: any) {
            console.error('Error checking Keycloak sync:', err);
            toast.error('Error al verificar sincronización con Keycloak');
        } finally {
            setCheckingSync(false);
        }
    };

    useEffect(() => {
        loadApplications();
    }, [page, pageSize, showDeleted]);

    useEffect(() => {
        if (applications.length > 0 && !showDeleted) {
            checkKeycloakSync();
        }
    }, [applications, showDeleted]);

    const filteredApps = applications.filter(
        (app) =>
            app.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
            app.client_id.toLowerCase().includes(searchTerm.toLowerCase()) ||
            (app.description ?? '').toLowerCase().includes(searchTerm.toLowerCase()) ||
            app.domain.toLowerCase().includes(searchTerm.toLowerCase())
    );

    const totalPages = Math.ceil(totalItems / pageSize);

    // Funciones de navegación con search params
    const updateSearchParams = (updates: Record<string, string | number | boolean>) => {
        const params = new URLSearchParams(searchParams.toString());

        Object.entries(updates).forEach(([key, value]) => {
            if (value === '' || value === false || value === null || value === undefined) {
                params.delete(key);
            } else {
                params.set(key, String(value));
            }
        });

        startTransition(() => {
            router.push(`?${params.toString()}`, { scroll: false });
        });
    };

    const handlePageChange = (newPage: number) => {
        updateSearchParams({ page: newPage });
    };

    const handlePageSizeChange = (newPageSize: string) => {
        updateSearchParams({ page: 1, page_size: newPageSize });
    };

    const handleToggleDeleted = () => {
        updateSearchParams({ page: 1, deleted: !showDeleted });
    };

    const getStatusBadge = (status: string) => {
        switch (status) {
            case 'active':
                return <Badge className="bg-chart-4/20 border-chart-4/30 text-chart-4">Activa</Badge>;
            case 'development':
                return <Badge className="bg-blue-500/20 border-blue-500/30 text-blue-500">Desarrollo</Badge>;
            case 'inactive':
                return <Badge className="bg-muted border-muted-foreground/30 text-muted-foreground">Inactiva</Badge>;
            default:
                return <Badge variant="outline">{status}</Badge>;
        }
    };

    const handleEdit = (app: ApplicationItem) => {
        setSelectedApp(app);
        setIsEditModalOpen(true);
    };

    const handleViewDetails = (app: ApplicationItem) => {
        setSelectedApp(app);
        setIsDetailModalOpen(true);
    };

    const handleDeleteClick = (app: ApplicationItem) => {
        setSelectedApp(app);
        setIsDeleteDialogOpen(true);
    };

    const handleDelete = async () => {
        if (!selectedApp) return;

        try {
            await fn_delete_application(selectedApp.id, selectedApp.keycloak_id);
            toast.success('Aplicación eliminada en backend y deshabilitada en Keycloak');
            setIsDeleteDialogOpen(false);
            setSelectedApp(null);
            loadApplications();
        } catch (err: any) {
            toast.error(err.message || 'Error al eliminar aplicación');
        }
    };

    const handleRestore = async (app: ApplicationItem) => {
        try {
            await fn_restore_application(app.id);
            toast.success('Aplicación restaurada correctamente');
            loadApplications();
        } catch (err: any) {
            toast.error(err.message || 'Error al restaurar aplicación');
        }
    };

    if (error) return <p className="py-10 text-destructive text-center">{error}</p>;

    return (
        <div className="space-y-6">
            <div className="flex justify-between items-center">
                <div>
                    <h1 className="font-bold text-foreground text-3xl">Aplicaciones</h1>
                    <p className="mt-1 text-muted-foreground">Gestiona las aplicaciones integradas con el sistema SSO</p>
                </div>
                <div className="flex gap-3">
                    {!showDeleted && unsyncedClients.length > 0 && (
                        <Button variant="outline" className="relative" onClick={() => setIsSyncModalOpen(true)}>
                            <RefreshCw className="mr-2 w-4 h-4" />
                            Sincronizar Keycloak
                            <Badge variant="destructive" className="absolute -top-2 -right-2 h-5 w-5 p-0 flex items-center justify-center">
                                {unsyncedClients.length}
                            </Badge>
                        </Button>
                    )}
                    <Button variant="outline">
                        <Download className="mr-2 w-4 h-4" />
                        Exportar
                    </Button>
                    <Button className="bg-linear-to-r from-primary to-chart-1" onClick={() => setIsCreateModalOpen(true)}>
                        <Plus className="mr-2 w-4 h-4" />
                        Nueva Aplicación
                    </Button>
                </div>
            </div>

            <ApplicationsStatsCards />

            <Card className="bg-card/50 border-border">
                <CardHeader>
                    <div className="flex justify-between items-center">
                        <div>
                            <CardTitle>Lista de Aplicaciones</CardTitle>
                            <CardDescription>
                                Mostrando {(page - 1) * pageSize + 1} - {Math.min(page * pageSize, totalItems)} de {totalItems} aplicaciones
                            </CardDescription>
                        </div>
                        <div className="flex gap-2">
                            <div className="relative">
                                <Search className="top-1/2 left-3 absolute w-4 h-4 text-muted-foreground -translate-y-1/2 transform" />
                                <Input placeholder="Buscar aplicaciones..." value={searchTerm} onChange={(e) => setSearchTerm(e.target.value)} className="bg-background/50 pl-10 w-80" />
                            </div>
                            <Button variant="outline" onClick={handleToggleDeleted}>
                                <Filter className="mr-2 w-4 h-4" />
                                {showDeleted ? 'Ver Activas' : 'Ver Eliminadas'}
                            </Button>
                        </div>
                    </div>
                </CardHeader>

                <CardContent>
                    <div className="border border-border rounded-lg">
                        <Table>
                            <TableHeader>
                                <TableRow className="bg-accent/50">
                                    <TableHead>Aplicación</TableHead>
                                    <TableHead>Client ID</TableHead>
                                    <TableHead>Dominio</TableHead>
                                    <TableHead>Usuarios</TableHead>
                                    <TableHead>Estado</TableHead>
                                    <TableHead className="text-right">Acciones</TableHead>
                                </TableRow>
                            </TableHeader>
                            <TableBody>
                                {loading || isPending ? (
                                    <TableRow>
                                        <TableCell colSpan={6} className="text-center text-muted-foreground py-8">
                                            Cargando...
                                        </TableCell>
                                    </TableRow>
                                ) : filteredApps.length === 0 ? (
                                    <TableRow>
                                        <TableCell colSpan={6} className="text-center text-muted-foreground py-8">
                                            No se encontraron aplicaciones
                                        </TableCell>
                                    </TableRow>
                                ) : (
                                    filteredApps.map((app) => (
                                        <TableRow key={app.id} className={app.is_deleted ? 'opacity-60' : ''}>
                                            <TableCell>
                                                <div className="space-y-1">
                                                    <div className="flex items-center gap-2">
                                                        <Boxes className="w-4 h-4 text-primary" />
                                                        <p className="font-medium text-foreground">{app.name}</p>
                                                    </div>
                                                    {app.description && <p className="text-muted-foreground text-sm line-clamp-2">{app.description}</p>}
                                                </div>
                                            </TableCell>
                                            <TableCell>
                                                <code className="text-xs bg-muted px-2 py-1 rounded">{app.client_id}</code>
                                            </TableCell>
                                            <TableCell>
                                                <div className="flex items-center gap-2">
                                                    <Globe className="w-4 h-4 text-chart-4" />
                                                    <span className="text-sm truncate max-w-[200px]" title={app.domain}>
                                                        {app.domain}
                                                    </span>
                                                </div>
                                            </TableCell>
                                            <TableCell>
                                                <span className="font-medium">{app.users_count}</span>
                                            </TableCell>
                                            <TableCell>{app.is_deleted ? <Badge variant="destructive">Eliminada</Badge> : getStatusBadge(app.status)}</TableCell>
                                            <TableCell className="text-right">
                                                <DropdownMenu>
                                                    <DropdownMenuTrigger asChild>
                                                        <Button variant="ghost" size="sm">
                                                            <MoreHorizontal className="w-4 h-4" />
                                                        </Button>
                                                    </DropdownMenuTrigger>
                                                    <DropdownMenuContent align="end">
                                                        <DropdownMenuLabel>Acciones</DropdownMenuLabel>
                                                        <DropdownMenuSeparator />
                                                        <DropdownMenuItem onClick={() => handleViewDetails(app)}>
                                                            <Eye className="mr-2 w-4 h-4" />
                                                            Ver Detalles
                                                        </DropdownMenuItem>
                                                        {!app.is_deleted && (
                                                            <>
                                                                <DropdownMenuItem onClick={() => handleEdit(app)}>
                                                                    <Edit className="mr-2 w-4 h-4" />
                                                                    Editar
                                                                </DropdownMenuItem>
                                                                <DropdownMenuItem className="text-destructive" onClick={() => handleDeleteClick(app)}>
                                                                    <Trash2 className="mr-2 w-4 h-4" />
                                                                    Eliminar
                                                                </DropdownMenuItem>
                                                            </>
                                                        )}
                                                        {app.is_deleted && (
                                                            <DropdownMenuItem onClick={() => handleRestore(app)}>
                                                                <AlertCircle className="mr-2 w-4 h-4" />
                                                                Restaurar
                                                            </DropdownMenuItem>
                                                        )}
                                                    </DropdownMenuContent>
                                                </DropdownMenu>
                                            </TableCell>
                                        </TableRow>
                                    ))
                                )}
                            </TableBody>
                        </Table>
                    </div>

                    {/* Paginación */}
                    <div className="flex items-center justify-between mt-4">
                        <div className="flex items-center gap-2">
                            <span className="text-sm text-muted-foreground">Mostrar</span>
                            <Select value={String(pageSize)} onValueChange={handlePageSizeChange}>
                                <SelectTrigger className="w-[100px]">
                                    <SelectValue />
                                </SelectTrigger>
                                <SelectContent>
                                    <SelectItem value="5">5</SelectItem>
                                    <SelectItem value="10">10</SelectItem>
                                    <SelectItem value="20">20</SelectItem>
                                    <SelectItem value="30">30</SelectItem>
                                    <SelectItem value="50">50</SelectItem>
                                </SelectContent>
                            </Select>
                            <span className="text-sm text-muted-foreground">por página</span>
                        </div>

                        <div className="flex items-center gap-4">
                            <div className="text-sm text-muted-foreground">
                                Página {page} de {totalPages}
                            </div>
                            <div className="flex gap-2">
                                <Button variant="outline" size="sm" onClick={() => handlePageChange(page - 1)} disabled={page === 1 || loading || isPending}>
                                    <ChevronLeft className="w-4 h-4 mr-1" />
                                    Anterior
                                </Button>
                                <Button variant="outline" size="sm" onClick={() => handlePageChange(page + 1)} disabled={page === totalPages || loading || isPending}>
                                    Siguiente
                                    <ChevronRight className="w-4 h-4 ml-1" />
                                </Button>
                            </div>
                        </div>
                    </div>
                </CardContent>
            </Card>

            {/* Modal Crear */}
            <ApplicationModal open={isCreateModalOpen} onOpenChange={setIsCreateModalOpen} onSuccess={loadApplications} />

            {/* Modal Editar */}
            <ApplicationModal open={isEditModalOpen} onOpenChange={setIsEditModalOpen} application={selectedApp} onSuccess={loadApplications} />

            {/* Modal Sincronizar Keycloak */}
            <SyncKeycloakModal open={isSyncModalOpen} onOpenChange={setIsSyncModalOpen} unsyncedClients={unsyncedClients} onSuccess={loadApplications} />

            {/* Modal Detalles */}
            <Dialog open={isDetailModalOpen} onOpenChange={setIsDetailModalOpen}>
                <DialogContent className="sm:max-w-[600px]">
                    <DialogHeader>
                        <DialogTitle>Detalles de la Aplicación</DialogTitle>
                        <DialogDescription>Información completa de la aplicación</DialogDescription>
                    </DialogHeader>
                    {selectedApp && (
                        <div className="space-y-4">
                            <div className="grid grid-cols-2 gap-4">
                                <div>
                                    <p className="text-sm text-muted-foreground">Nombre</p>
                                    <p className="font-medium">{selectedApp.name}</p>
                                </div>
                                <div>
                                    <p className="text-sm text-muted-foreground">Client ID</p>
                                    <code className="text-sm bg-muted px-2 py-1 rounded">{selectedApp.client_id}</code>
                                </div>
                            </div>

                            <div className="grid grid-cols-2 gap-4">
                                <div className="col-span-2">
                                    <p className="text-sm text-muted-foreground">Dominio</p>
                                    <p className="font-medium text-sm break-all">{selectedApp.domain}</p>
                                </div>
                            </div>

                            <div>
                                <p className="text-sm text-muted-foreground">Descripción</p>
                                <p className="font-medium">{selectedApp.description || 'Sin descripción'}</p>
                            </div>

                            <div className="grid grid-cols-2 gap-4">
                                <div>
                                    <p className="text-sm text-muted-foreground">Usuarios</p>
                                    <p className="font-medium">{selectedApp.users_count}</p>
                                </div>
                                <div>
                                    <p className="text-sm text-muted-foreground">Estado</p>
                                    <div className="mt-1">{getStatusBadge(selectedApp.status)}</div>
                                </div>
                            </div>

                            {selectedApp.admins && selectedApp.admins.length > 0 && (
                                <div>
                                    <p className="text-sm text-muted-foreground mb-2">Administradores</p>
                                    <div className="space-y-2">
                                        {selectedApp.admins.map((admin, idx) => (
                                            <div key={idx} className="flex items-center gap-2 text-sm bg-muted p-2 rounded">
                                                <span className="font-medium">{admin.full_name}</span>
                                                <span className="text-muted-foreground">({admin.email})</span>
                                            </div>
                                        ))}
                                    </div>
                                </div>
                            )}

                            <div className="grid grid-cols-2 gap-4 pt-2 border-t">
                                <div>
                                    <p className="text-sm text-muted-foreground">Creado</p>
                                    <p className="text-sm">{new Date(selectedApp.created_at).toLocaleString('es-PE')}</p>
                                </div>
                                <div>
                                    <p className="text-sm text-muted-foreground">Actualizado</p>
                                    <p className="text-sm">{new Date(selectedApp.updated_at).toLocaleString('es-PE')}</p>
                                </div>
                            </div>
                        </div>
                    )}
                    <DialogFooter>
                        <Button variant="outline" onClick={() => setIsDetailModalOpen(false)}>
                            Cerrar
                        </Button>
                    </DialogFooter>
                </DialogContent>
            </Dialog>

            {/* Dialog Confirmar Eliminación */}
            <Dialog open={isDeleteDialogOpen} onOpenChange={setIsDeleteDialogOpen}>
                <DialogContent>
                    <DialogHeader>
                        <DialogTitle>¿Eliminar aplicación?</DialogTitle>
                        <DialogDescription>
                            Esta acción eliminará la aplicación <strong>{selectedApp?.name}</strong>. Esta acción se puede revertir.
                        </DialogDescription>
                    </DialogHeader>
                    <DialogFooter>
                        <Button variant="outline" onClick={() => setIsDeleteDialogOpen(false)}>
                            Cancelar
                        </Button>
                        <Button variant="destructive" onClick={handleDelete}>
                            Eliminar
                        </Button>
                    </DialogFooter>
                </DialogContent>
            </Dialog>
        </div>
    );
}