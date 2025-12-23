'use client';

import { useState, useEffect, useTransition } from 'react';
import { useRouter, useSearchParams } from 'next/navigation';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { Badge } from '@/components/ui/badge';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table';
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuLabel, DropdownMenuSeparator, DropdownMenuTrigger } from '@/components/ui/dropdown-menu';
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle } from '@/components/ui/dialog';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { Search, Plus, Filter, MoreHorizontal, Edit, Trash2, Eye, Shield, ChevronLeft, ChevronRight, AlertCircle, Layers, Package, Users } from 'lucide-react';
import { toast } from 'sonner';
import { fn_get_application_roles } from '@/actions/application-roles/fn_get_application_roles';
import { fn_get_applications } from '@/actions/applications/fn_get_applications';
import { fn_delete_application_role } from '@/actions/application-roles/fn_delete_application_role';
import { fn_restore_application_role } from '@/actions/application-roles/fn_restore_application_role';
import type { ApplicationRoleItem } from '@/types/application-roles';
import type { ApplicationItem } from '@/types/applications';
import { ApplicationRolesStatsCards } from '@/components/custom/card/application-roles-stats-card';
import ApplicationRoleModal from './application-role-modal';

export default function ApplicationRolesContent() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const [isPending, startTransition] = useTransition();

  // Obtener params de URL
  const page = Number(searchParams.get('page')) || 1;
  const pageSize = Number(searchParams.get('page_size')) || 10;
  const showDeleted = searchParams.get('deleted') === 'true';

  const [roles, setRoles] = useState<ApplicationRoleItem[]>([]);
  const [applications, setApplications] = useState<ApplicationItem[]>([]);
  const [selectedApplication, setSelectedApplication] = useState<string>('');
  const [searchTerm, setSearchTerm] = useState('');
  const [loading, setLoading] = useState(false);
  const [loadingApps, setLoadingApps] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [totalItems, setTotalItems] = useState(0);

  // Modals
  const [isCreateModalOpen, setIsCreateModalOpen] = useState(false);
  const [isEditModalOpen, setIsEditModalOpen] = useState(false);
  const [isDetailModalOpen, setIsDetailModalOpen] = useState(false);
  const [isDeleteDialogOpen, setIsDeleteDialogOpen] = useState(false);
  const [selectedRole, setSelectedRole] = useState<ApplicationRoleItem | null>(null);

  // Cargar aplicaciones al montar
  useEffect(() => {
    const loadApplications = async () => {
      try {
        setLoadingApps(true);
        const response = await fn_get_applications(1, 100, false);
        setApplications(response.data);
        
        // Seleccionar la primera aplicación por defecto
        if (response.data.length > 0 && !selectedApplication) {
          setSelectedApplication(response.data[0].id);
        }
      } catch (err: any) {
        console.error('Error loading applications:', err);
        toast.error('Error al cargar aplicaciones');
      } finally {
        setLoadingApps(false);
      }
    };
    loadApplications();
  }, []);

  const loadRoles = async () => {
    if (!selectedApplication) {
      setRoles([]);
      setTotalItems(0);
      return;
    }

    try {
      setLoading(true);
      setError(null);
      const filters: any = { application_id: selectedApplication };
      if (showDeleted) filters.is_deleted = true;

      const response = await fn_get_application_roles(page, pageSize, filters);
      setRoles(response.data);
      setTotalItems(response.total);
    } catch (err: any) {
      setError(err.message ?? 'Error desconocido');
      toast.error('Error al cargar roles de aplicación');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    if (selectedApplication) {
      loadRoles();
    }
  }, [page, pageSize, showDeleted, selectedApplication]);

  const filteredRoles = roles.filter(
    (r) =>
      r.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
      (r.description || '').toLowerCase().includes(searchTerm.toLowerCase())
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

  const handleApplicationChange = (appId: string) => {
    setSelectedApplication(appId);
    updateSearchParams({ page: 1 });
  };

  const handleEdit = (role: ApplicationRoleItem) => {
    setSelectedRole(role);
    setIsEditModalOpen(true);
  };

  const handleViewDetails = (role: ApplicationRoleItem) => {
    setSelectedRole(role);
    setIsDetailModalOpen(true);
  };

  const handleDeleteClick = (role: ApplicationRoleItem) => {
    setSelectedRole(role);
    setIsDeleteDialogOpen(true);
  };

  const handleDelete = async () => {
    if (!selectedRole) return;

    try {
      await fn_delete_application_role(selectedRole.id);
      toast.success('Rol eliminado correctamente');
      setIsDeleteDialogOpen(false);
      setSelectedRole(null);
      loadRoles();
    } catch (err: any) {
      toast.error(err.message || 'Error al eliminar rol');
    }
  };

  const handleRestore = async (role: ApplicationRoleItem) => {
    try {
      await fn_restore_application_role(role.id);
      toast.success('Rol restaurado correctamente');
      loadRoles();
    } catch (err: any) {
      toast.error(err.message || 'Error al restaurar rol');
    }
  };

  const selectedApp = applications.find(app => app.id === selectedApplication);

  if (error) return <p className="py-10 text-destructive text-center">{error}</p>;

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <div>
          <h1 className="font-bold text-foreground text-3xl">Roles de Aplicación</h1>
          <p className="mt-1 text-muted-foreground">Gestiona los roles y permisos de cada aplicación</p>
        </div>
        <div className="flex gap-3">
          <Button variant="outline" onClick={handleToggleDeleted} disabled={!selectedApplication}>
            <Filter className="mr-2 w-4 h-4" />
            {showDeleted ? 'Ver Activos' : 'Ver Eliminados'}
          </Button>
          <Button 
            className="bg-linear-to-r from-primary to-chart-1" 
            onClick={() => setIsCreateModalOpen(true)}
            disabled={!selectedApplication}
          >
            <Plus className="mr-2 w-4 h-4" />
            Nuevo Rol
          </Button>
        </div>
      </div>

      <ApplicationRolesStatsCards />

      <Card className="bg-card/50 border-border">
        <CardHeader>
          <div className="space-y-4">
            <div className="flex justify-between items-center">
              <div>
                <CardTitle>Lista de Roles</CardTitle>
                <CardDescription>
                  {selectedApplication 
                    ? `Mostrando ${(page - 1) * pageSize + 1} - ${Math.min(page * pageSize, totalItems)} de ${totalItems} roles`
                    : 'Selecciona una aplicación para ver sus roles'
                  }
                </CardDescription>
              </div>
              <div className="flex gap-2">
                <div className="relative">
                  <Search className="top-1/2 left-3 absolute w-4 h-4 text-muted-foreground -translate-y-1/2 transform" />
                  <Input 
                    placeholder="Buscar roles..." 
                    value={searchTerm} 
                    onChange={(e) => setSearchTerm(e.target.value)} 
                    className="bg-background/50 pl-10 w-80"
                    disabled={!selectedApplication}
                  />
                </div>
              </div>
            </div>

            {/* Selector de Aplicación */}
            <div className="flex items-center gap-3 p-4 bg-muted/50 rounded-lg border border-border">
              <Layers className="w-5 h-5 text-primary" />
              <div className="flex-1">
                <Label className="text-sm font-medium mb-2 block">Aplicación</Label>
                <Select value={selectedApplication} onValueChange={handleApplicationChange} disabled={loadingApps}>
                  <SelectTrigger className="w-full max-w-md bg-background">
                    <SelectValue placeholder="Selecciona una aplicación" />
                  </SelectTrigger>
                  <SelectContent>
                    {applications.map((app) => (
                      <SelectItem key={app.id} value={app.id}>
                        <div className="flex items-center gap-2">
                          <span className="font-medium">{app.name}</span>
                          <span className="text-xs text-muted-foreground">({app.client_id})</span>
                        </div>
                      </SelectItem>
                    ))}
                  </SelectContent>
                </Select>
              </div>
              {selectedApp && (
                <div className="text-sm text-muted-foreground">
                  <Badge variant="outline">{roles.length} roles</Badge>
                </div>
              )}
            </div>
          </div>
        </CardHeader>

        <CardContent>
          {!selectedApplication ? (
            <div className="flex flex-col items-center justify-center py-16 text-center">
              <Shield className="w-16 h-16 text-muted-foreground/50 mb-4" />
              <h3 className="text-lg font-semibold mb-2">Selecciona una Aplicación</h3>
              <p className="text-sm text-muted-foreground max-w-md">
                Para ver y gestionar los roles, primero debes seleccionar una aplicación del selector superior.
              </p>
            </div>
          ) : (
            <>
              <div className="border border-border rounded-lg">
                <Table>
                  <TableHeader>
                    <TableRow className="bg-accent/50">
                      <TableHead>Nombre del Rol</TableHead>
                      <TableHead>Descripción</TableHead>
                      <TableHead>Módulos</TableHead>
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
                    ) : filteredRoles.length === 0 ? (
                      <TableRow>
                        <TableCell colSpan={6} className="text-center text-muted-foreground py-8">
                          No se encontraron roles para esta aplicación
                        </TableCell>
                      </TableRow>
                    ) : (
                      filteredRoles.map((role) => (
                        <TableRow key={role.id} className={role.is_deleted ? 'opacity-60' : ''}>
                          <TableCell>
                            <div className="flex items-center gap-2">
                              <Shield className="w-4 h-4 text-primary" />
                              <span className="font-medium text-foreground">{role.name}</span>
                            </div>
                          </TableCell>
                          <TableCell>
                            <span className="text-sm text-muted-foreground">
                              {role.description || 'Sin descripción'}
                            </span>
                          </TableCell>
                          <TableCell>
                            <div className="flex items-center gap-2">
                              <Package className="w-4 h-4 text-chart-1" />
                              <Badge variant="outline">{role.modules_count || 0}</Badge>
                            </div>
                          </TableCell>
                          <TableCell>
                            <div className="flex items-center gap-2">
                              <Users className="w-4 h-4 text-chart-4" />
                              <Badge variant="outline">{role.users_count || 0}</Badge>
                            </div>
                          </TableCell>
                          <TableCell>
                            {role.is_deleted ? (
                              <Badge variant="destructive">Eliminado</Badge>
                            ) : (
                              <Badge className="bg-chart-4/20 border-chart-4/30 text-chart-4">Activo</Badge>
                            )}
                          </TableCell>
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
                                <DropdownMenuItem onClick={() => handleViewDetails(role)}>
                                  <Eye className="mr-2 w-4 h-4" />
                                  Ver Detalles
                                </DropdownMenuItem>
                                {!role.is_deleted && (
                                  <>
                                    <DropdownMenuItem onClick={() => handleEdit(role)}>
                                      <Edit className="mr-2 w-4 h-4" />
                                      Editar
                                    </DropdownMenuItem>
                                    <DropdownMenuItem className="text-destructive" onClick={() => handleDeleteClick(role)}>
                                      <Trash2 className="mr-2 w-4 h-4" />
                                      Eliminar
                                    </DropdownMenuItem>
                                  </>
                                )}
                                {role.is_deleted && (
                                  <DropdownMenuItem onClick={() => handleRestore(role)}>
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
              {totalItems > 0 && (
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
              )}
            </>
          )}
        </CardContent>
      </Card>

      {/* Modal Crear */}
      <ApplicationRoleModal 
        open={isCreateModalOpen} 
        onOpenChange={setIsCreateModalOpen} 
        onSuccess={loadRoles}
        defaultApplicationId={selectedApplication}
      />

      {/* Modal Editar */}
      <ApplicationRoleModal 
        open={isEditModalOpen} 
        onOpenChange={setIsEditModalOpen} 
        role={selectedRole} 
        onSuccess={loadRoles}
        defaultApplicationId={selectedApplication}
      />

      {/* Modal Detalles continúa en el siguiente mensaje... */}

      {/* Modal Detalles */}
      <Dialog open={isDetailModalOpen} onOpenChange={setIsDetailModalOpen}>
        <DialogContent className="sm:max-w-[600px] max-h-[90vh] overflow-y-auto">
          <DialogHeader>
            <DialogTitle>Detalles del Rol</DialogTitle>
            <DialogDescription>Información completa del rol de aplicación</DialogDescription>
          </DialogHeader>
          {selectedRole && (
            <div className="space-y-4">
              <div className="grid grid-cols-2 gap-4">
                <div>
                  <p className="text-sm text-muted-foreground">Nombre del Rol</p>
                  <p className="font-medium text-lg">{selectedRole.name}</p>
                </div>
                <div>
                  <p className="text-sm text-muted-foreground">Estado</p>
                  <div className="mt-1">
                    {selectedRole.is_deleted ? (
                      <Badge variant="destructive">Eliminado</Badge>
                    ) : (
                      <Badge className="bg-chart-4/20 border-chart-4/30 text-chart-4">Activo</Badge>
                    )}
                  </div>
                </div>
              </div>

              <div>
                <p className="text-sm text-muted-foreground">Descripción</p>
                <p className="text-sm mt-1">{selectedRole.description || 'Sin descripción'}</p>
              </div>

              {selectedRole.application && (
                <div>
                  <p className="text-sm text-muted-foreground">Aplicación</p>
                  <div className="mt-1">
                    <p className="font-medium">{selectedRole.application.name}</p>
                    <p className="text-xs text-muted-foreground">{selectedRole.application.client_id}</p>
                  </div>
                </div>
              )}

              <div className="grid grid-cols-2 gap-4">
                <div className="flex items-center gap-3 p-3 bg-muted/50 rounded-lg">
                  <Package className="w-8 h-8 text-chart-1" />
                  <div>
                    <p className="text-sm text-muted-foreground">Módulos</p>
                    <p className="text-xl font-bold">{selectedRole.modules_count || 0}</p>
                  </div>
                </div>
                <div className="flex items-center gap-3 p-3 bg-muted/50 rounded-lg">
                  <Users className="w-8 h-8 text-chart-4" />
                  <div>
                    <p className="text-sm text-muted-foreground">Usuarios</p>
                    <p className="text-xl font-bold">{selectedRole.users_count || 0}</p>
                  </div>
                </div>
              </div>

              <div className="grid grid-cols-2 gap-4 pt-2 border-t">
                <div>
                  <p className="text-sm text-muted-foreground">Creado</p>
                  <p className="text-sm">{new Date(selectedRole.created_at).toLocaleString('es-PE')}</p>
                </div>
                <div>
                  <p className="text-sm text-muted-foreground">Actualizado</p>
                  <p className="text-sm">{new Date(selectedRole.updated_at).toLocaleString('es-PE')}</p>
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
            <DialogTitle>¿Eliminar rol?</DialogTitle>
            <DialogDescription>
              Esta acción eliminará el rol <strong>{selectedRole?.name}</strong>. Esta acción se puede revertir.
            </DialogDescription>
          </DialogHeader>
          <div className="bg-yellow-500/10 border border-yellow-500/20 rounded-lg p-4 flex gap-3">
            <AlertCircle className="w-5 h-5 text-yellow-500 shrink-0 mt-0.5" />
            <div className="text-sm">
              <p className="font-medium text-yellow-600 dark:text-yellow-500 mb-1">Advertencia</p>
              <p className="text-muted-foreground">
                Los usuarios asignados a este rol perderán sus permisos asociados.
              </p>
            </div>
          </div>
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