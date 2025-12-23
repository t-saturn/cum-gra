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
import { Textarea } from '@/components/ui/textarea';
import { Search, Plus, Filter, MoreHorizontal, Eye, XCircle, CheckCircle, Trash2, UserCheck, ChevronLeft, ChevronRight, AlertCircle, Layers, Users2, ShieldAlert } from 'lucide-react';
import { toast } from 'sonner';
import { fn_get_user_application_roles } from '@/actions/user-application-roles/fn_get_user_application_roles';
import { fn_get_applications } from '@/actions/applications/fn_get_applications';
import { fn_revoke_role } from '@/actions/user-application-roles/fn_revoke_role';
import { fn_restore_role } from '@/actions/user-application-roles/fn_restore_role';
import { fn_delete_assignment } from '@/actions/user-application-roles/fn_delete_assignment';
import { fn_undelete_assignment } from '@/actions/user-application-roles/fn_undelete_assignment';
import type { UserApplicationRoleItem } from '@/types/user-application-roles';
import type { ApplicationItem } from '@/types/applications';
import { UserApplicationRolesStatsCards } from '@/components/custom/card/user-application-roles-stats-card';
import AssignRoleModal from './assign-role-modal';
import BulkAssignModal from './bulk-assign-modal';

export default function UserRolesContent() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const [isPending, startTransition] = useTransition();

  // Obtener params de URL
  const page = Number(searchParams.get('page')) || 1;
  const pageSize = Number(searchParams.get('page_size')) || 10;
  const showDeleted = searchParams.get('deleted') === 'true';
  const showRevoked = searchParams.get('revoked') === 'true';

  const [assignments, setAssignments] = useState<UserApplicationRoleItem[]>([]);
  const [applications, setApplications] = useState<ApplicationItem[]>([]);
  const [selectedApplication, setSelectedApplication] = useState<string>('');
  const [searchTerm, setSearchTerm] = useState('');
  const [loading, setLoading] = useState(false);
  const [loadingApps, setLoadingApps] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [totalItems, setTotalItems] = useState(0);

  // Modals
  const [isAssignModalOpen, setIsAssignModalOpen] = useState(false);
  const [isBulkAssignModalOpen, setIsBulkAssignModalOpen] = useState(false);
  const [isDetailModalOpen, setIsDetailModalOpen] = useState(false);
  const [isRevokeDialogOpen, setIsRevokeDialogOpen] = useState(false);
  const [isDeleteDialogOpen, setIsDeleteDialogOpen] = useState(false);
  const [selectedAssignment, setSelectedAssignment] = useState<UserApplicationRoleItem | null>(null);
  const [revokeReason, setRevokeReason] = useState('');

  // Cargar aplicaciones al montar
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

  const loadAssignments = async () => {
    if (!selectedApplication) {
      setAssignments([]);
      setTotalItems(0);
      return;
    }

    try {
      setLoading(true);
      setError(null);
      const filters: any = { application_id: selectedApplication };
      if (showDeleted) filters.is_deleted = true;
      if (showRevoked) filters.is_revoked = true;

      const response = await fn_get_user_application_roles(page, pageSize, filters);
      setAssignments(response.data);
      setTotalItems(response.total);
    } catch (err: any) {
      setError(err.message ?? 'Error desconocido');
      toast.error('Error al cargar asignaciones');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    if (selectedApplication) {
      loadAssignments();
    }
  }, [page, pageSize, showDeleted, showRevoked, selectedApplication]);

  const filteredAssignments = assignments.filter(
    (a) =>
      (a.user_email || '').toLowerCase().includes(searchTerm.toLowerCase()) ||
      (a.user_full_name || '').toLowerCase().includes(searchTerm.toLowerCase()) ||
      (a.role_name || '').toLowerCase().includes(searchTerm.toLowerCase())
  );

  const totalPages = Math.ceil(totalItems / pageSize);

  // Funciones de navegación
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

  const handleToggleRevoked = () => {
    updateSearchParams({ page: 1, revoked: !showRevoked });
  };

  const handleApplicationChange = (appId: string) => {
    setSelectedApplication(appId);
    updateSearchParams({ page: 1 });
  };

  const handleViewDetails = (assignment: UserApplicationRoleItem) => {
    setSelectedAssignment(assignment);
    setIsDetailModalOpen(true);
  };

  const handleRevokeClick = (assignment: UserApplicationRoleItem) => {
    setSelectedAssignment(assignment);
    setRevokeReason('');
    setIsRevokeDialogOpen(true);
  };

  const handleRevoke = async () => {
    if (!selectedAssignment) return;

    try {
      await fn_revoke_role(selectedAssignment.id, { reason: revokeReason || undefined });
      toast.success('Rol revocado correctamente');
      setIsRevokeDialogOpen(false);
      setSelectedAssignment(null);
      setRevokeReason('');
      loadAssignments();
    } catch (err: any) {
      toast.error(err.message || 'Error al revocar rol');
    }
  };

  const handleRestore = async (assignment: UserApplicationRoleItem) => {
    try {
      await fn_restore_role(assignment.id);
      toast.success('Rol restaurado correctamente');
      loadAssignments();
    } catch (err: any) {
      toast.error(err.message || 'Error al restaurar rol');
    }
  };

  const handleDeleteClick = (assignment: UserApplicationRoleItem) => {
    setSelectedAssignment(assignment);
    setIsDeleteDialogOpen(true);
  };

  const handleDelete = async () => {
    if (!selectedAssignment) return;

    try {
      await fn_delete_assignment(selectedAssignment.id);
      toast.success('Asignación eliminada correctamente');
      setIsDeleteDialogOpen(false);
      setSelectedAssignment(null);
      loadAssignments();
    } catch (err: any) {
      toast.error(err.message || 'Error al eliminar asignación');
    }
  };

  const handleUndelete = async (assignment: UserApplicationRoleItem) => {
    try {
      await fn_undelete_assignment(assignment.id);
      toast.success('Asignación recuperada correctamente');
      loadAssignments();
    } catch (err: any) {
      toast.error(err.message || 'Error al recuperar asignación');
    }
  };

  const getStatusBadge = (assignment: UserApplicationRoleItem) => {
    if (assignment.is_deleted) {
      return <Badge variant="destructive">Eliminado</Badge>;
    }
    if (assignment.revoked_at) {
      return <Badge className="bg-orange-500/20 border-orange-500/30 text-orange-500">Revocado</Badge>;
    }
    return <Badge className="bg-green-500/20 border-green-500/30 text-green-500">Activo</Badge>;
  };

  const selectedApp = applications.find(app => app.id === selectedApplication);

  if (error) return <p className="py-10 text-destructive text-center">{error}</p>;

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <div>
          <h1 className="font-bold text-foreground text-3xl">Asignación de Roles</h1>
          <p className="mt-1 text-muted-foreground">Gestiona los roles asignados a los usuarios por aplicación</p>
        </div>
        <div className="flex gap-3">
           <Button 
            variant="outline" 
            onClick={() => setIsBulkAssignModalOpen(true)}
            disabled={!selectedApplication}
          >
            <Users2 className="mr-2 w-4 h-4" />
            Asignación Masiva
          </Button>
          <Button 
            className="bg-linear-to-r from-primary to-chart-1" 
            onClick={() => setIsAssignModalOpen(true)}
            disabled={!selectedApplication}
          >
            <Plus className="mr-2 w-4 h-4" />
            Nueva Asignación
          </Button>
        </div>
      </div>

      <UserApplicationRolesStatsCards />

      <Card className="bg-card/50 border-border">
        <CardHeader>
          <div className="space-y-4">
            <div className="flex justify-between items-center">
              <div>
                <CardTitle>Lista de Asignaciones</CardTitle>
                <CardDescription>
                  {selectedApplication 
                    ? `Mostrando ${(page - 1) * pageSize + 1} - ${Math.min(page * pageSize, totalItems)} de ${totalItems} asignaciones`
                    : 'Selecciona una aplicación para ver sus asignaciones'
                  }
                </CardDescription>
              </div>
              <div className="flex gap-2">
                <Button variant="ghost" size="sm" onClick={handleToggleRevoked} className={showRevoked ? 'bg-accent' : ''}>
                   <ShieldAlert className="w-4 h-4 mr-2" />
                   {showRevoked ? 'Ocultar Revocados' : 'Ver Revocados'}
                </Button>
                <Button variant="ghost" size="sm" onClick={handleToggleDeleted} className={showDeleted ? 'bg-destructive/10 text-destructive' : ''}>
                   <Trash2 className="w-4 h-4 mr-2" />
                   {showDeleted ? 'Ocultar Eliminados' : 'Ver Eliminados'}
                </Button>
              </div>
            </div>

            {/* Filtros y Buscador */}
            <div className="flex items-center gap-4">
              <div className="flex-1">
                <div className="relative">
                  <Search className="top-1/2 left-3 absolute w-4 h-4 text-muted-foreground -translate-y-1/2 transform" />
                  <Input 
                    placeholder="Buscar por usuario, email o rol..." 
                    value={searchTerm} 
                    onChange={(e) => setSearchTerm(e.target.value)} 
                    className="bg-background/50 pl-10 max-w-sm"
                    disabled={!selectedApplication}
                  />
                </div>
              </div>
              
              <div className="flex items-center gap-2 bg-muted/50 p-2 rounded-lg border border-border min-w-[300px]">
                <Layers className="w-4 h-4 text-primary ml-2" />
                <Select value={selectedApplication} onValueChange={handleApplicationChange} disabled={loadingApps}>
                  <SelectTrigger className="bg-background border-none shadow-none focus:ring-0">
                    <SelectValue placeholder="Selecciona una aplicación" />
                  </SelectTrigger>
                  <SelectContent>
                    {applications.map((app) => (
                      <SelectItem key={app.id} value={app.id}>
                        {app.name}
                      </SelectItem>
                    ))}
                  </SelectContent>
                </Select>
              </div>
            </div>
          </div>
        </CardHeader>

        <CardContent>
          {!selectedApplication ? (
            <div className="flex flex-col items-center justify-center py-16 text-center">
              <Layers className="w-16 h-16 text-muted-foreground/50 mb-4" />
              <h3 className="text-lg font-semibold mb-2">Selecciona una Aplicación</h3>
              <p className="text-sm text-muted-foreground max-w-md">
                Para gestionar las asignaciones de roles, primero debes seleccionar una aplicación del selector.
              </p>
            </div>
          ) : (
            <>
              <div className="border border-border rounded-lg">
                <Table>
                  <TableHeader>
                    <TableRow className="bg-accent/50">
                      <TableHead>Usuario</TableHead>
                      <TableHead>Rol Asignado</TableHead>
                      <TableHead>Otorgado El</TableHead>
                      <TableHead>Estado</TableHead>
                      <TableHead className="text-right">Acciones</TableHead>
                    </TableRow>
                  </TableHeader>
                  <TableBody>
                    {loading || isPending ? (
                      <TableRow>
                        <TableCell colSpan={5} className="text-center text-muted-foreground py-8">
                          Cargando asignaciones...
                        </TableCell>
                      </TableRow>
                    ) : filteredAssignments.length === 0 ? (
                      <TableRow>
                        <TableCell colSpan={5} className="text-center text-muted-foreground py-8">
                          No se encontraron asignaciones que coincidan con los filtros
                        </TableCell>
                      </TableRow>
                    ) : (
                      filteredAssignments.map((assignment) => (
                        <TableRow key={assignment.id} className={assignment.is_deleted ? 'opacity-60' : ''}>
                          <TableCell>
                            <div className="flex flex-col">
                              <span className="font-medium">{assignment.user_full_name || 'Desconocido'}</span>
                              <span className="text-xs text-muted-foreground">{assignment.user_email}</span>
                            </div>
                          </TableCell>
                          <TableCell>
                            <Badge variant="outline" className="font-medium">
                              {assignment.role_name || 'Sin nombre'}
                            </Badge>
                          </TableCell>
                          <TableCell>
                             <span className="text-sm text-muted-foreground">
                               {new Date(assignment.granted_at).toLocaleDateString('es-PE')}
                             </span>
                          </TableCell>
                          <TableCell>
                            {getStatusBadge(assignment)}
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
                                <DropdownMenuItem onClick={() => handleViewDetails(assignment)}>
                                  <Eye className="mr-2 w-4 h-4" />
                                  Detalles
                                </DropdownMenuItem>
                                
                                {/* Acciones si NO está eliminado ni revocado */}
                                {!assignment.is_deleted && !assignment.revoked_at && (
                                  <DropdownMenuItem onClick={() => handleRevokeClick(assignment)} className="text-orange-500 focus:text-orange-500">
                                    <XCircle className="mr-2 w-4 h-4" />
                                    Revocar Acceso
                                  </DropdownMenuItem>
                                )}

                                {/* Acciones si está revocado pero no eliminado */}
                                {assignment.revoked_at && !assignment.is_deleted && (
                                  <DropdownMenuItem onClick={() => handleAssignRoleAgain(assignment)}>
                                     {/* Nota: Re-asignar podría ser simplemente crear uno nuevo, o restaurar este si la lógica lo permite */}
                                     <UserCheck className="mr-2 w-4 h-4" />
                                     Reasignar
                                  </DropdownMenuItem>
                                )}

                                <DropdownMenuSeparator />
                                
                                {!assignment.is_deleted ? (
                                  <DropdownMenuItem onClick={() => handleDeleteClick(assignment)} className="text-destructive focus:text-destructive">
                                    <Trash2 className="mr-2 w-4 h-4" />
                                    Eliminar
                                  </DropdownMenuItem>
                                ) : (
                                  <DropdownMenuItem onClick={() => handleUndelete(assignment)}>
                                    <CheckCircle className="mr-2 w-4 h-4" />
                                    Restaurar (Undelete)
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
                        <SelectItem value="50">50</SelectItem>
                      </SelectContent>
                    </Select>
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

      {/* Modals */}
      <AssignRoleModal 
        open={isAssignModalOpen} 
        onOpenChange={setIsAssignModalOpen} 
        onSuccess={loadAssignments}
        defaultApplicationId={selectedApplication}
      />

      <BulkAssignModal 
        open={isBulkAssignModalOpen} 
        onOpenChange={setIsBulkAssignModalOpen} 
        onSuccess={loadAssignments}
        defaultApplicationId={selectedApplication}
      />

      {/* Detalle Dialog */}
      <Dialog open={isDetailModalOpen} onOpenChange={setIsDetailModalOpen}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Detalle de Asignación</DialogTitle>
          </DialogHeader>
          {selectedAssignment && (
            <div className="space-y-4">
              <div className="grid grid-cols-2 gap-4">
                <div>
                  <Label className="text-muted-foreground">Usuario</Label>
                  <p className="font-medium">{selectedAssignment.user_full_name}</p>
                  <p className="text-sm text-muted-foreground">{selectedAssignment.user_email}</p>
                </div>
                <div>
                  <Label className="text-muted-foreground">Rol</Label>
                  <p className="font-medium">{selectedAssignment.role_name}</p>
                </div>
              </div>
              <div>
                  <Label className="text-muted-foreground">Aplicación</Label>
                  <p>{selectedAssignment.application_name}</p>
              </div>
              {selectedAssignment.revoked_at && (
                <div className="bg-orange-500/10 p-3 rounded-md border border-orange-500/20">
                  <Label className="text-orange-600 font-medium">Revocado</Label>
                  <p className="text-sm">Fecha: {new Date(selectedAssignment.revoked_at).toLocaleString()}</p>
                  {selectedAssignment.revoke_reason && (
                    <p className="text-sm mt-1">Motivo: {selectedAssignment.revoke_reason}</p>
                  )}
                </div>
              )}
            </div>
          )}
        </DialogContent>
      </Dialog>

      {/* Revocar Dialog */}
      <Dialog open={isRevokeDialogOpen} onOpenChange={setIsRevokeDialogOpen}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Revocar Acceso</DialogTitle>
            <DialogDescription>
              ¿Estás seguro de que deseas revocar este rol? El usuario perderá los permisos asociados inmediatamente.
            </DialogDescription>
          </DialogHeader>
          <div className="space-y-2">
            <Label>Motivo (Opcional)</Label>
            <Textarea 
              value={revokeReason}
              onChange={(e) => setRevokeReason(e.target.value)}
              placeholder="Ej: Cambio de puesto, baja de usuario..."
            />
          </div>
          <DialogFooter>
            <Button variant="outline" onClick={() => setIsRevokeDialogOpen(false)}>Cancelar</Button>
            <Button variant="destructive" onClick={handleRevoke}>Revocar Acceso</Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>

      {/* Eliminar Dialog */}
      <Dialog open={isDeleteDialogOpen} onOpenChange={setIsDeleteDialogOpen}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Eliminar Asignación</DialogTitle>
            <DialogDescription>
              Esta acción eliminará el registro de asignación. Si el rol estaba activo, el usuario perderá acceso.
            </DialogDescription>
          </DialogHeader>
          <DialogFooter>
            <Button variant="outline" onClick={() => setIsDeleteDialogOpen(false)}>Cancelar</Button>
            <Button variant="destructive" onClick={handleDelete}>Eliminar</Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </div>
  );

  // Helper function para el botón re-asignar en el menú (opcional)
  function handleAssignRoleAgain(assignment: UserApplicationRoleItem) {
    setIsAssignModalOpen(true);
    // Nota: Idealmente pasaríamos props para pre-llenar el modal, pero por simplicidad abrimos el modal genérico
  }
}