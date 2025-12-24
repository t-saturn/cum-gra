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
import { Search, Plus, Filter, MoreHorizontal, Edit, Trash2, Eye, Key, ChevronLeft, ChevronRight, AlertCircle, Layers, Shield, Package, Upload } from 'lucide-react';
import { toast } from 'sonner';
import { fn_get_module_role_permissions } from '@/actions/module-role-permissions/fn_get_module_role_permissions';
import { fn_get_applications } from '@/actions/applications/fn_get_applications';
import { fn_get_application_roles } from '@/actions/application-roles/fn_get_application_roles';
import { fn_delete_module_role_permission } from '@/actions/module-role-permissions/fn_delete_module_role_permission';
import { fn_restore_module_role_permission } from '@/actions/module-role-permissions/fn_restore_module_role_permission';
import type { ModuleRolePermissionItem, PermissionType } from '@/types/module-role-permissions';
import type { ApplicationItem } from '@/types/applications';
import type { ApplicationRoleItem } from '@/types/application-roles';
import { ModuleRolePermissionsStatsCards } from '@/components/custom/card/module-role-permissions-stats-card';
import PermissionModal from './permission-modal';
import BulkPermissionModal from './bulk-permission-modal';

const PERMISSION_COLORS: Record<PermissionType, string> = {
  read: 'bg-blue-500/20 border-blue-500/30 text-blue-600',
  write: 'bg-green-500/20 border-green-500/30 text-green-600',
  execute: 'bg-purple-500/20 border-purple-500/30 text-purple-600',
  delete: 'bg-red-500/20 border-red-500/30 text-red-600',
  admin: 'bg-orange-500/20 border-orange-500/30 text-orange-600',
};

const PERMISSION_LABELS: Record<PermissionType, string> = {
  read: 'Lectura',
  write: 'Escritura',
  execute: 'Ejecución',
  delete: 'Eliminación',
  admin: 'Admin',
};

export default function ModuleRolesContent() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const [isPending, startTransition] = useTransition();

  const page = Number(searchParams.get('page')) || 1;
  const pageSize = Number(searchParams.get('page_size')) || 10;
  const showDeleted = searchParams.get('deleted') === 'true';

  const [permissions, setPermissions] = useState<ModuleRolePermissionItem[]>([]);
  const [applications, setApplications] = useState<ApplicationItem[]>([]);
  const [roles, setRoles] = useState<ApplicationRoleItem[]>([]);
  const [selectedApplication, setSelectedApplication] = useState<string>('');
  const [selectedRole, setSelectedRole] = useState<string>('');
  const [searchTerm, setSearchTerm] = useState('');
  const [loading, setLoading] = useState(false);
  const [loadingApps, setLoadingApps] = useState(true);
  const [loadingRoles, setLoadingRoles] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [totalItems, setTotalItems] = useState(0);

  // Modals
  const [isCreateModalOpen, setIsCreateModalOpen] = useState(false);
  const [isEditModalOpen, setIsEditModalOpen] = useState(false);
  const [isBulkModalOpen, setIsBulkModalOpen] = useState(false);
  const [isDetailModalOpen, setIsDetailModalOpen] = useState(false);
  const [isDeleteDialogOpen, setIsDeleteDialogOpen] = useState(false);
  const [selectedPermission, setSelectedPermission] = useState<ModuleRolePermissionItem | null>(null);

  // Cargar aplicaciones
  useEffect(() => {
    const loadApplications = async () => {
      try {
        setLoadingApps(true);
        const response = await fn_get_applications(1, 100, false);
        setApplications(response.data);
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

  // Cargar roles cuando cambia la aplicación
  useEffect(() => {
    const loadRoles = async () => {
      if (!selectedApplication) {
        setRoles([]);
        setSelectedRole('');
        return;
      }
      try {
        setLoadingRoles(true);
        const response = await fn_get_application_roles(1, 100, { application_id: selectedApplication });
        setRoles(response.data);
        if (response.data.length > 0) {
          setSelectedRole(response.data[0].id);
        } else {
          setSelectedRole('');
        }
      } catch (err: any) {
        console.error('Error loading roles:', err);
        toast.error('Error al cargar roles');
      } finally {
        setLoadingRoles(false);
      }
    };
    loadRoles();
  }, [selectedApplication]);

  const loadPermissions = async () => {
    if (!selectedRole) {
      setPermissions([]);
      setTotalItems(0);
      return;
    }

    try {
      setLoading(true);
      setError(null);
      const filters: any = { role_id: selectedRole };
      if (showDeleted) filters.is_deleted = true;

      const response = await fn_get_module_role_permissions(page, pageSize, filters);
      setPermissions(response.data);
      setTotalItems(response.total);
    } catch (err: any) {
      setError(err.message ?? 'Error desconocido');
      toast.error('Error al cargar permisos');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    if (selectedRole) {
      loadPermissions();
    }
  }, [page, pageSize, showDeleted, selectedRole]);

  const filteredPermissions = permissions.filter(
    (p) =>
      p.module_name.toLowerCase().includes(searchTerm.toLowerCase()) ||
      p.role_name.toLowerCase().includes(searchTerm.toLowerCase()) ||
      p.permission_type.toLowerCase().includes(searchTerm.toLowerCase())
  );

  const totalPages = Math.ceil(totalItems / pageSize);

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

  const handlePageChange = (newPage: number) => updateSearchParams({ page: newPage });
  const handlePageSizeChange = (newPageSize: string) => updateSearchParams({ page: 1, page_size: newPageSize });
  const handleToggleDeleted = () => updateSearchParams({ page: 1, deleted: !showDeleted });

  const handleApplicationChange = (appId: string) => {
    setSelectedApplication(appId);
    setSelectedRole('');
    updateSearchParams({ page: 1 });
  };

  const handleRoleChange = (roleId: string) => {
    setSelectedRole(roleId);
    updateSearchParams({ page: 1 });
  };

  const handleEdit = (permission: ModuleRolePermissionItem) => {
    setSelectedPermission(permission);
    setIsEditModalOpen(true);
  };

  const handleViewDetails = (permission: ModuleRolePermissionItem) => {
    setSelectedPermission(permission);
    setIsDetailModalOpen(true);
  };

  const handleDeleteClick = (permission: ModuleRolePermissionItem) => {
    setSelectedPermission(permission);
    setIsDeleteDialogOpen(true);
  };

  const handleDelete = async () => {
    if (!selectedPermission) return;
    try {
      await fn_delete_module_role_permission(selectedPermission.id);
      toast.success('Permiso eliminado correctamente');
      setIsDeleteDialogOpen(false);
      setSelectedPermission(null);
      loadPermissions();
    } catch (err: any) {
      toast.error(err.message || 'Error al eliminar permiso');
    }
  };

  const handleRestore = async (permission: ModuleRolePermissionItem) => {
    try {
      await fn_restore_module_role_permission(permission.id);
      toast.success('Permiso restaurado correctamente');
      loadPermissions();
    } catch (err: any) {
      toast.error(err.message || 'Error al restaurar permiso');
    }
  };

  const selectedApp = applications.find(app => app.id === selectedApplication);
  const selectedRoleData = roles.find(r => r.id === selectedRole);

  if (error) return <p className="py-10 text-destructive text-center">{error}</p>;

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <div>
          <h1 className="font-bold text-foreground text-3xl">Permisos de Módulos</h1>
          <p className="mt-1 text-muted-foreground">Gestiona los permisos de cada rol sobre los módulos</p>
        </div>
        <div className="flex gap-3">
          <Button variant="outline" onClick={handleToggleDeleted} disabled={!selectedRole}>
            <Filter className="mr-2 w-4 h-4" />
            {showDeleted ? 'Ver Activos' : 'Ver Eliminados'}
          </Button>
          <Button 
            variant="outline"
            onClick={() => setIsBulkModalOpen(true)}
            disabled={!selectedRole}
          >
            <Upload className="mr-2 w-4 h-4" />
            Asignación Masiva
          </Button>
          <Button 
            className="bg-linear-to-r from-primary to-chart-1" 
            onClick={() => setIsCreateModalOpen(true)}
            disabled={!selectedRole}
          >
            <Plus className="mr-2 w-4 h-4" />
            Nuevo Permiso
          </Button>
        </div>
      </div>

      <ModuleRolePermissionsStatsCards />

      <Card className="bg-card/50 border-border">
        <CardHeader>
          <div className="space-y-4">
            <div className="flex justify-between items-center">
              <div>
                <CardTitle>Lista de Permisos</CardTitle>
                <CardDescription>
                  {selectedRole 
                    ? `Mostrando ${(page - 1) * pageSize + 1} - ${Math.min(page * pageSize, totalItems)} de ${totalItems} permisos`
                    : 'Selecciona una aplicación y rol para ver sus permisos'
                  }
                </CardDescription>
              </div>
              <div className="flex gap-2">
                <div className="relative">
                  <Search className="top-1/2 left-3 absolute w-4 h-4 text-muted-foreground -translate-y-1/2 transform" />
                  <Input 
                    placeholder="Buscar permisos..." 
                    value={searchTerm} 
                    onChange={(e) => setSearchTerm(e.target.value)} 
                    className="bg-background/50 pl-10 w-80"
                    disabled={!selectedRole}
                  />
                </div>
              </div>
            </div>

            {/* Selectores de Aplicación y Rol */}
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4 p-4 bg-muted/50 rounded-lg border border-border">
              {/* Aplicación */}
              <div className="flex items-center gap-3">
                <Layers className="w-5 h-5 text-primary" />
                <div className="flex-1">
                  <Label className="text-sm font-medium mb-2 block">Aplicación</Label>
                  <Select value={selectedApplication} onValueChange={handleApplicationChange} disabled={loadingApps}>
                    <SelectTrigger className="bg-background">
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
              </div>

              {/* Rol */}
              <div className="flex items-center gap-3">
                <Shield className="w-5 h-5 text-chart-1" />
                <div className="flex-1">
                  <Label className="text-sm font-medium mb-2 block">Rol</Label>
                  <Select 
                    value={selectedRole} 
                    onValueChange={handleRoleChange} 
                    disabled={loadingRoles || !selectedApplication || roles.length === 0}
                  >
                    <SelectTrigger className="bg-background">
                      <SelectValue placeholder={roles.length === 0 ? "No hay roles" : "Selecciona un rol"} />
                    </SelectTrigger>
                    <SelectContent>
                      {roles.map((role) => (
                        <SelectItem key={role.id} value={role.id}>
                          {role.name}
                        </SelectItem>
                      ))}
                    </SelectContent>
                  </Select>
                </div>
                {selectedRoleData && (
                  <Badge variant="outline" className="mt-6">{permissions.length} permisos</Badge>
                )}
              </div>
            </div>
          </div>
        </CardHeader>

        <CardContent>
          {!selectedRole ? (
            <div className="flex flex-col items-center justify-center py-16 text-center">
              <Key className="w-16 h-16 text-muted-foreground/50 mb-4" />
              <h3 className="text-lg font-semibold mb-2">Selecciona un Rol</h3>
              <p className="text-sm text-muted-foreground max-w-md">
                Para ver y gestionar los permisos, primero debes seleccionar una aplicación y un rol.
              </p>
            </div>
          ) : (
            <>
              <div className="border border-border rounded-lg">
                <Table>
                  <TableHeader>
                    <TableRow className="bg-accent/50">
                      <TableHead>Módulo</TableHead>
                      <TableHead>Ruta</TableHead>
                      <TableHead>Tipo de Permiso</TableHead>
                      <TableHead>Estado</TableHead>
                      <TableHead className="text-right">Acciones</TableHead>
                    </TableRow>
                  </TableHeader>
                  <TableBody>
                    {loading || isPending ? (
                      <TableRow>
                        <TableCell colSpan={5} className="text-center text-muted-foreground py-8">
                          Cargando...
                        </TableCell>
                      </TableRow>
                    ) : filteredPermissions.length === 0 ? (
                      <TableRow>
                        <TableCell colSpan={5} className="text-center text-muted-foreground py-8">
                          No se encontraron permisos para este rol
                        </TableCell>
                      </TableRow>
                    ) : (
                      filteredPermissions.map((permission) => (
                        <TableRow key={permission.id} className={permission.is_deleted ? 'opacity-60' : ''}>
                          <TableCell>
                            <div className="flex items-center gap-2">
                              <Package className="w-4 h-4 text-primary" />
                              <span className="font-medium text-foreground">{permission.module_name}</span>
                            </div>
                          </TableCell>
                          <TableCell>
                            <code className="text-xs bg-muted px-2 py-1 rounded">{permission.module_route}</code>
                          </TableCell>
                          <TableCell>
                            <Badge className={PERMISSION_COLORS[permission.permission_type]}>
                              {PERMISSION_LABELS[permission.permission_type]}
                            </Badge>
                          </TableCell>
                          <TableCell>
                            {permission.is_deleted ? (
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
                                <DropdownMenuItem onClick={() => handleViewDetails(permission)}>
                                  <Eye className="mr-2 w-4 h-4" />
                                  Ver Detalles
                                </DropdownMenuItem>
                                {!permission.is_deleted && (
                                  <>
                                    <DropdownMenuItem onClick={() => handleEdit(permission)}>
                                      <Edit className="mr-2 w-4 h-4" />
                                      Editar
                                    </DropdownMenuItem>
                                    <DropdownMenuItem className="text-destructive" onClick={() => handleDeleteClick(permission)}>
                                      <Trash2 className="mr-2 w-4 h-4" />
                                      Eliminar
                                    </DropdownMenuItem>
                                  </>
                                )}
                                {permission.is_deleted && (
                                  <DropdownMenuItem onClick={() => handleRestore(permission)}>
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
      <PermissionModal 
        open={isCreateModalOpen} 
        onOpenChange={setIsCreateModalOpen} 
        onSuccess={loadPermissions}
        defaultApplicationId={selectedApplication}
        defaultRoleId={selectedRole}
      />

      {/* Modal Editar */}
      <PermissionModal 
        open={isEditModalOpen} 
        onOpenChange={setIsEditModalOpen} 
        permission={selectedPermission} 
        onSuccess={loadPermissions}
        defaultApplicationId={selectedApplication}
        defaultRoleId={selectedRole}
      />

      {/* Modal Asignación Masiva */}
      {selectedRoleData && (
        <BulkPermissionModal 
          open={isBulkModalOpen} 
          onOpenChange={setIsBulkModalOpen} 
          onSuccess={loadPermissions}
          applicationId={selectedApplication}
          roleId={selectedRole}
          roleName={selectedRoleData.name}
        />
      )}

      {/* Modal Detalles */}
      <Dialog open={isDetailModalOpen} onOpenChange={setIsDetailModalOpen}>
        <DialogContent className="sm:max-w-[600px] max-h-[90vh] overflow-y-auto">
          <DialogHeader>
            <DialogTitle>Detalles del Permiso</DialogTitle>
            <DialogDescription>Información completa del permiso de módulo</DialogDescription>
          </DialogHeader>
          {selectedPermission && (
            <div className="space-y-4">
              <div className="grid grid-cols-2 gap-4">
                <div>
                  <p className="text-sm text-muted-foreground">Módulo</p>
                  <p className="font-medium text-lg">{selectedPermission.module_name}</p>
                </div>
                <div>
                  <p className="text-sm text-muted-foreground">Ruta</p>
                  <code className="text-sm bg-muted px-2 py-1 rounded">{selectedPermission.module_route}</code>
                </div>
              </div>

              <div className="grid grid-cols-2 gap-4">
                <div>
                  <p className="text-sm text-muted-foreground">Rol</p>
                  <p className="font-medium">{selectedPermission.role_name}</p>
                </div>
                <div>
                  <p className="text-sm text-muted-foreground">Tipo de Permiso</p>
                  <Badge className={`mt-1 ${PERMISSION_COLORS[selectedPermission.permission_type]}`}>
                    {PERMISSION_LABELS[selectedPermission.permission_type]}
                  </Badge>
                </div>
              </div>

              <div className="grid grid-cols-2 gap-4">
                <div>
                  <p className="text-sm text-muted-foreground">Aplicación</p>
                  <p className="font-medium">{selectedPermission.application_name}</p>
                  <p className="text-xs text-muted-foreground">{selectedPermission.application_client_id}</p>
                </div>
                <div>
                  <p className="text-sm text-muted-foreground">Estado</p>
                  <div className="mt-1">
                    {selectedPermission.is_deleted ? (
                      <Badge variant="destructive">Eliminado</Badge>
                    ) : (
                      <Badge className="bg-chart-4/20 border-chart-4/30 text-chart-4">Activo</Badge>
                    )}
                  </div>
                </div>
              </div>

              <div className="grid grid-cols-2 gap-4 pt-2 border-t">
                <div>
                  <p className="text-sm text-muted-foreground">Creado</p>
                  <p className="text-sm">{new Date(selectedPermission.created_at).toLocaleString('es-PE')}</p>
                </div>
                {selectedPermission.deleted_at && (
                  <div>
                    <p className="text-sm text-muted-foreground">Eliminado</p>
                    <p className="text-sm">{new Date(selectedPermission.deleted_at).toLocaleString('es-PE')}</p>
                  </div>
                )}
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
            <DialogTitle>¿Eliminar permiso?</DialogTitle>
            <DialogDescription>
              Esta acción eliminará el permiso <strong>{PERMISSION_LABELS[selectedPermission?.permission_type || 'read']}</strong> del módulo <strong>{selectedPermission?.module_name}</strong>. Esta acción se puede revertir.
            </DialogDescription>
          </DialogHeader>
          <div className="bg-yellow-500/10 border border-yellow-500/20 rounded-lg p-4 flex gap-3">
            <AlertCircle className="w-5 h-5 text-yellow-500 shrink-0 mt-0.5" />
            <div className="text-sm">
              <p className="font-medium text-yellow-600 dark:text-yellow-500 mb-1">Advertencia</p>
              <p className="text-muted-foreground">
                Los usuarios con este rol perderán acceso al módulo especificado.
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