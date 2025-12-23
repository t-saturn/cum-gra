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
import { Search, Plus, Filter, Download, MoreHorizontal, Edit, Trash2, Eye, User, RefreshCw, AlertCircle, ChevronLeft, ChevronRight, X, Upload } from 'lucide-react';
import { toast } from 'sonner';
import { fn_get_users } from '@/actions/users/fn_get_users';
import { fn_delete_user } from '@/actions/users/fn_delete_user';
import { fn_restore_user } from '@/actions/users/fn_restore_user';
import { fn_download_users_template } from '@/actions/users/fn_download_template';
import { fn_get_keycloak_users, KeycloakUserSimple } from '@/actions/keycloak/users/fn_get_keycloak_users';
import type { UserItem } from '@/types/users';
import { UsersStatsCards } from '@/components/custom/card/users-stats-cards';
import UserModal from './user-modal';
import SyncKeycloakUsersModal from './sync-keycloak-users-modal';
import BulkUploadModal from './bulk-upload-modal';

export default function UsersContent() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const [isPending, startTransition] = useTransition();

  // Obtener params de URL
  const page = Number(searchParams.get('page')) || 1;
  const pageSize = Number(searchParams.get('page_size')) || 10;
  const showDeleted = searchParams.get('deleted') === 'true';

  const [users, setUsers] = useState<UserItem[]>([]);
  const [searchTerm, setSearchTerm] = useState('');
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [totalItems, setTotalItems] = useState(0);
  const [downloadingTemplate, setDownloadingTemplate] = useState(false);

  // Filtros
  const [statusFilter, setStatusFilter] = useState<string>('');
  const [unitFilter, setUnitFilter] = useState<string>('');
  const [positionFilter, setPositionFilter] = useState<string>('');

  // Keycloak sync
  const [keycloakUsers, setKeycloakUsers] = useState<KeycloakUserSimple[]>([]);
  const [unsyncedUsers, setUnsyncedUsers] = useState<KeycloakUserSimple[]>([]);
  const [checkingSync, setCheckingSync] = useState(false);

  // Modals
  const [isCreateModalOpen, setIsCreateModalOpen] = useState(false);
  const [isEditModalOpen, setIsEditModalOpen] = useState(false);
  const [isDetailModalOpen, setIsDetailModalOpen] = useState(false);
  const [isDeleteDialogOpen, setIsDeleteDialogOpen] = useState(false);
  const [isSyncModalOpen, setIsSyncModalOpen] = useState(false);
  const [isBulkUploadModalOpen, setIsBulkUploadModalOpen] = useState(false);
  const [selectedUser, setSelectedUser] = useState<UserItem | null>(null);

  const loadUsers = async () => {
    try {
      setLoading(true);
      setError(null);
      const filters: any = {};
      if (statusFilter) filters.status = statusFilter;
      if (unitFilter) filters.organic_unit_id = unitFilter;
      if (positionFilter) filters.position_id = positionFilter;
      if (showDeleted) filters.is_deleted = true;

      const response = await fn_get_users(page, pageSize, filters);
      setUsers(response.data);
      setTotalItems(response.total);
    } catch (err: any) {
      setError(err.message ?? 'Error desconocido');
      toast.error('Error al cargar usuarios');
    } finally {
      setLoading(false);
    }
  };

  const checkKeycloakSync = async () => {
    try {
      setCheckingSync(true);
      const kcUsers = await fn_get_keycloak_users();
      setKeycloakUsers(kcUsers);

      // Comparar con usuarios del backend
      const backendUsernames = new Set(users.map((u) => u.dni));
      const unsynced = kcUsers.filter((kc) => !backendUsernames.has(kc.username));
      setUnsyncedUsers(unsynced);

      if (unsynced.length > 0) {
        toast.info(`${unsynced.length} usuario(s) de Keycloak sin sincronizar`);
      }
    } catch (err: any) {
      console.error('Error checking Keycloak sync:', err);
      toast.error('Error al verificar sincronización con Keycloak');
    } finally {
      setCheckingSync(false);
    }
  };

  const handleDownloadTemplate = async () => {
    try {
      setDownloadingTemplate(true);
      const blob = await fn_download_users_template();
      
      // Crear URL temporal y descargar
      const url = window.URL.createObjectURL(blob);
      const link = document.createElement('a');
      link.href = url;
      link.download = `plantilla_usuarios_${new Date().toISOString().split('T')[0]}.xlsx`;
      document.body.appendChild(link);
      link.click();
      document.body.removeChild(link);
      window.URL.revokeObjectURL(url);
      
      toast.success('Plantilla descargada correctamente');
    } catch (err: any) {
      toast.error(err.message || 'Error al descargar plantilla');
    } finally {
      setDownloadingTemplate(false);
    }
  };

  useEffect(() => {
    loadUsers();
  }, [page, pageSize, showDeleted, statusFilter, unitFilter, positionFilter]);

  useEffect(() => {
    if (users.length > 0 && !showDeleted) {
      checkKeycloakSync();
    }
  }, [users, showDeleted]);

  const filteredUsers = users.filter(
    (u) =>
      u.email.toLowerCase().includes(searchTerm.toLowerCase()) ||
      u.dni.includes(searchTerm) ||
      u.first_name.toLowerCase().includes(searchTerm.toLowerCase()) ||
      u.last_name.toLowerCase().includes(searchTerm.toLowerCase()) ||
      (u.phone ?? '').includes(searchTerm)
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

  const clearFilters = () => {
    setStatusFilter('');
    setUnitFilter('');
    setPositionFilter('');
  };

  const hasActiveFilters = statusFilter || unitFilter || positionFilter;

  const getStatusBadge = (status: string) => {
    switch (status) {
      case 'active':
        return <Badge className="bg-chart-4/20 border-chart-4/30 text-chart-4">Activo</Badge>;
      case 'suspended':
        return <Badge className="bg-yellow-500/20 border-yellow-500/30 text-yellow-500">Suspendido</Badge>;
      case 'inactive':
        return <Badge className="bg-muted border-muted-foreground/30 text-muted-foreground">Inactivo</Badge>;
      default:
        return <Badge variant="outline">{status}</Badge>;
    }
  };

  const handleEdit = (user: UserItem) => {
    setSelectedUser(user);
    setIsEditModalOpen(true);
  };

  const handleViewDetails = (user: UserItem) => {
    setSelectedUser(user);
    setIsDetailModalOpen(true);
  };

  const handleDeleteClick = (user: UserItem) => {
    setSelectedUser(user);
    setIsDeleteDialogOpen(true);
  };

  const handleDelete = async () => {
    if (!selectedUser) return;

    try {
      await fn_delete_user(selectedUser.id, selectedUser.keycloak_id ?? undefined);
      toast.success('Usuario eliminado y deshabilitado en Keycloak');
      setIsDeleteDialogOpen(false);
      setSelectedUser(null);
      loadUsers();
    } catch (err: any) {
      toast.error(err.message || 'Error al eliminar usuario');
    }
  };

  const handleRestore = async (user: UserItem) => {
    try {
      await fn_restore_user(user.id, user.keycloak_id ?? undefined);
      toast.success('Usuario restaurado y habilitado en Keycloak');
      loadUsers();
    } catch (err: any) {
      toast.error(err.message || 'Error al restaurar usuario');
    }
  };

  if (error) return <p className="py-10 text-destructive text-center">{error}</p>;

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <div>
          <h1 className="font-bold text-foreground text-3xl">Usuarios</h1>
          <p className="mt-1 text-muted-foreground">Gestiona los usuarios del sistema y sincronización con Keycloak</p>
        </div>
        <div className="flex gap-3">
          {!showDeleted && unsyncedUsers.length > 0 && (
            <Button variant="outline" className="relative" onClick={() => setIsSyncModalOpen(true)}>
              <RefreshCw className="mr-2 w-4 h-4" />
              Sincronizar Keycloak
              <Badge variant="destructive" className="absolute -top-2 -right-2 h-5 w-5 p-0 flex items-center justify-center">
                {unsyncedUsers.length}
              </Badge>
            </Button>
          )}
          <Button variant="outline" onClick={handleToggleDeleted}>
            <Filter className="mr-2 w-4 h-4" />
            {showDeleted ? 'Ver Activos' : 'Ver Eliminados'}
          </Button>
          <Button 
            variant="outline" 
            onClick={handleDownloadTemplate}
            disabled={downloadingTemplate}
          >
            <Download className="mr-2 w-4 h-4" />
            {downloadingTemplate ? 'Descargando...' : 'Descargar Plantilla'}
          </Button>
          <Button 
            variant="outline"
            onClick={() => setIsBulkUploadModalOpen(true)}
          >
            <Upload className="mr-2 w-4 h-4" />
            Carga Masiva
          </Button>
          <Button className="bg-linear-to-r from-primary to-chart-1" onClick={() => setIsCreateModalOpen(true)}>
            <Plus className="mr-2 w-4 h-4" />
            Nuevo Usuario
          </Button>
        </div>
      </div>

      <UsersStatsCards />

      <Card className="bg-card/50 border-border">
        <CardHeader>
          <div className="flex justify-between items-center">
            <div>
              <CardTitle>Lista de Usuarios</CardTitle>
              <CardDescription>
                Mostrando {(page - 1) * pageSize + 1} - {Math.min(page * pageSize, totalItems)} de {totalItems} usuarios
              </CardDescription>
            </div>
            <div className="flex gap-2">
              <div className="relative">
                <Search className="top-1/2 left-3 absolute w-4 h-4 text-muted-foreground -translate-y-1/2 transform" />
                <Input placeholder="Buscar usuarios..." value={searchTerm} onChange={(e) => setSearchTerm(e.target.value)} className="bg-background/50 pl-10 w-80" />
              </div>
            </div>
          </div>

          {/* Filtros */}
          <div className="flex flex-wrap gap-2 mt-4">
            <Select value={statusFilter} onValueChange={setStatusFilter}>
              <SelectTrigger className="w-[180px]">
                <SelectValue placeholder="Estado" />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value=" ">Todos los estados</SelectItem>
                <SelectItem value="active">Activos</SelectItem>
                <SelectItem value="suspended">Suspendidos</SelectItem>
                <SelectItem value="inactive">Inactivos</SelectItem>
              </SelectContent>
            </Select>

            <Input
              placeholder="Filtrar por unidad orgánica"
              value={unitFilter}
              onChange={(e) => setUnitFilter(e.target.value)}
              className="w-[200px]"
            />
            <Input
              placeholder="Filtrar por posición"
              value={positionFilter}
              onChange={(e) => setPositionFilter(e.target.value)}
              className="w-[200px]"
            />
            {hasActiveFilters && (
              <Button variant="ghost" size="sm" onClick={clearFilters}>
                <X className="w-4 h-4 mr-2" />
                Limpiar filtros
              </Button>
            )}
          </div>
        </CardHeader>

        <CardContent>
          <div className="border border-border rounded-lg">
            <Table>
              <TableHeader>
                <TableRow className="bg-accent/50">
                  <TableHead>Usuario</TableHead>
                  <TableHead>DNI</TableHead>
                  <TableHead>Unidad Orgánica</TableHead>
                  <TableHead>Posición</TableHead>
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
                ) : filteredUsers.length === 0 ? (
                  <TableRow>
                    <TableCell colSpan={6} className="text-center text-muted-foreground py-8">
                      No se encontraron usuarios
                    </TableCell>
                  </TableRow>
                ) : (
                  filteredUsers.map((user) => (
                    <TableRow key={user.id} className={user.is_deleted ? 'opacity-60' : ''}>
                      <TableCell>
                        <div className="space-y-1">
                          <div className="flex items-center gap-2">
                            <User className="w-4 h-4 text-primary" />
                            <div>
                              <p className="font-medium text-foreground">
                                {user.first_name} {user.last_name}
                              </p>
                              <p className="text-xs text-muted-foreground">{user.email}</p>
                            </div>
                          </div>
                        </div>
                      </TableCell>
                      <TableCell>
                        <code className="text-xs bg-muted px-2 py-1 rounded">{user.dni}</code>
                      </TableCell>
                      <TableCell>
                        {user.organic_unit ? (
                          <div>
                            <p className="text-sm font-medium">{user.organic_unit.name}</p>
                            <p className="text-xs text-muted-foreground">{user.organic_unit.acronym}</p>
                          </div>
                        ) : (
                          <span className="text-muted-foreground text-sm">Sin asignar</span>
                        )}
                      </TableCell>
                      <TableCell>
                        {user.structural_position ? (
                          <div>
                            <p className="text-sm font-medium">{user.structural_position.name}</p>
                            {user.structural_position.level && (
                              <Badge variant="outline" className="text-xs mt-1">
                                Nivel {user.structural_position.level}
                              </Badge>
                            )}
                          </div>
                        ) : (
                          <span className="text-muted-foreground text-sm">Sin asignar</span>
                        )}
                      </TableCell>
                      <TableCell>{user.is_deleted ? <Badge variant="destructive">Eliminado</Badge> : getStatusBadge(user.status)}</TableCell>
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
                            <DropdownMenuItem onClick={() => handleViewDetails(user)}>
                              <Eye className="mr-2 w-4 h-4" />
                              Ver Detalles
                            </DropdownMenuItem>
                            {!user.is_deleted && (
                              <>
                                <DropdownMenuItem onClick={() => handleEdit(user)}>
                                  <Edit className="mr-2 w-4 h-4" />
                                  Editar
                                </DropdownMenuItem>
                                <DropdownMenuItem className="text-destructive" onClick={() => handleDeleteClick(user)}>
                                  <Trash2 className="mr-2 w-4 h-4" />
                                  Eliminar
                                </DropdownMenuItem>
                              </>
                            )}
                            {user.is_deleted && (
                              <DropdownMenuItem onClick={() => handleRestore(user)}>
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
      <UserModal open={isCreateModalOpen} onOpenChange={setIsCreateModalOpen} onSuccess={loadUsers} />

      {/* Modal Editar */}
      <UserModal open={isEditModalOpen} onOpenChange={setIsEditModalOpen} user={selectedUser} onSuccess={loadUsers} />

      {/* Modal Carga Masiva */}
      <BulkUploadModal open={isBulkUploadModalOpen} onOpenChange={setIsBulkUploadModalOpen} onSuccess={loadUsers} />

      {/* Modal Sincronizar Keycloak */}
      <SyncKeycloakUsersModal open={isSyncModalOpen} onOpenChange={setIsSyncModalOpen} unsyncedUsers={unsyncedUsers} onSuccess={loadUsers} />

      {/* Modal Detalles */}
      <Dialog open={isDetailModalOpen} onOpenChange={setIsDetailModalOpen}>
        <DialogContent className="sm:max-w-[700px] max-h-[90vh] overflow-y-auto">
          <DialogHeader>
            <DialogTitle>Detalles del Usuario</DialogTitle>
            <DialogDescription>Información completa del usuario del sistema</DialogDescription>
          </DialogHeader>
          {selectedUser && (
            <div className="space-y-4">
              <div className="grid grid-cols-2 gap-4">
                <div>
                  <p className="text-sm text-muted-foreground">Nombre Completo</p>
                  <p className="font-medium text-lg">
                    {selectedUser.first_name} {selectedUser.last_name}
                  </p>
                </div>
                <div>
                  <p className="text-sm text-muted-foreground">DNI</p>
                  <code className="text-sm bg-muted px-2 py-1 rounded">{selectedUser.dni}</code>
                </div>
              </div>

              <div className="grid grid-cols-2 gap-4">
                <div>
                  <p className="text-sm text-muted-foreground">Email</p>
                  <p className="font-medium">{selectedUser.email}</p>
                </div>
                <div>
                  <p className="text-sm text-muted-foreground">Teléfono</p>
                  <p className="font-medium">{selectedUser.phone || 'N/A'}</p>
                </div>
              </div>

              <div className="grid grid-cols-2 gap-4">
                <div>
                  <p className="text-sm text-muted-foreground">Estado</p>
                  <div className="mt-1">{getStatusBadge(selectedUser.status)}</div>
                </div>
                <div>
                  <p className="text-sm text-muted-foreground">Código SGD</p>
                  <p className="font-medium">{selectedUser.cod_emp_sgd || 'N/A'}</p>
                </div>
              </div>

              {selectedUser.organic_unit && (
                <div>
                  <p className="text-sm text-muted-foreground">Unidad Orgánica</p>
                  <p className="font-medium">
                    {selectedUser.organic_unit.name} ({selectedUser.organic_unit.acronym})
                  </p>
                </div>
              )}

              {selectedUser.structural_position && (
                <div>
                  <p className="text-sm text-muted-foreground">Posición Estructural</p>
                  <div className="flex items-center gap-2">
                    <p className="font-medium">{selectedUser.structural_position.name}</p>
                    {selectedUser.structural_position.level && (
                      <Badge variant="outline">Nivel {selectedUser.structural_position.level}</Badge>
                    )}
                  </div>
                </div>
              )}

              {selectedUser.ubigeo && (
                <div>
                  <p className="text-sm text-muted-foreground">Ubicación</p>
                  <p className="font-medium">
                    {selectedUser.ubigeo.department} / {selectedUser.ubigeo.province} / {selectedUser.ubigeo.district}
                  </p>
                  <code className="text-xs bg-muted px-2 py-1 rounded">{selectedUser.ubigeo.ubigeo_code}</code>
                </div>
              )}

              <div className="grid grid-cols-2 gap-4 pt-2 border-t">
                <div>
                  <p className="text-sm text-muted-foreground">Creado</p>
                  <p className="text-sm">{new Date(selectedUser.created_at).toLocaleString('es-PE')}</p>
                </div>
                <div>
                  <p className="text-sm text-muted-foreground">Actualizado</p>
                  <p className="text-sm">{new Date(selectedUser.updated_at).toLocaleString('es-PE')}</p>
                </div>
              </div>

              {selectedUser.keycloak_id && (
                <div className="bg-muted/50 p-3 rounded-lg">
                  <p className="text-xs text-muted-foreground">Sincronizado con Keycloak</p>
                  <code className="text-xs">{selectedUser.keycloak_id}</code>
                </div>
              )}
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
            <DialogTitle>¿Eliminar usuario?</DialogTitle>
            <DialogDescription>
              Esta acción eliminará al usuario <strong>{selectedUser?.first_name} {selectedUser?.last_name}</strong> y lo deshabilitará en Keycloak. Esta acción se puede revertir.
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