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
import { Search, Plus, Filter, MoreHorizontal, Edit, Trash2, Eye, LayoutGrid, ChevronLeft, ChevronRight, AlertCircle, Layers } from 'lucide-react';
import { toast } from 'sonner';
import { fn_get_modules } from '@/actions/modules/fn_get_modules';
import { fn_get_applications } from '@/actions/applications/fn_get_applications';
import { fn_delete_module } from '@/actions/modules/fn_delete_module';
import { fn_restore_module } from '@/actions/modules/fn_restore_module';
import type { ModuleItem } from '@/types/modules';
import type { ApplicationItem } from '@/types/applications';
import { ModulesStatsCards } from '@/components/custom/card/modules-stats-card';
import ModuleModal from './module-modal';
import { Label } from '@/components/ui/label';
import { enrichModulesWithParentInfo } from '@/utils/modules';

export default function ModulesContent() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const [isPending, startTransition] = useTransition();

  // Obtener params de URL
  const page = Number(searchParams.get('page')) || 1;
  const pageSize = Number(searchParams.get('page_size')) || 10;
  const showDeleted = searchParams.get('deleted') === 'true';

  const [modules, setModules] = useState<ModuleItem[]>([]);
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
  const [selectedModule, setSelectedModule] = useState<ModuleItem | null>(null);

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

  const loadModules = async () => {
    if (!selectedApplication) {
      setModules([]);
      setTotalItems(0);
      return;
    }

    try {
      setLoading(true);
      setError(null);
      const filters: any = { application_id: selectedApplication };
      if (showDeleted) filters.is_deleted = true;

      const response = await fn_get_modules(page, pageSize, filters);
      
      // Enriquecer con información del padre
      const enrichedModules = enrichModulesWithParentInfo(response.data);
      
      setModules(enrichedModules);
      setTotalItems(response.total);
    } catch (err: any) {
      setError(err.message ?? 'Error desconocido');
      toast.error('Error al cargar módulos');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    if (selectedApplication) {
      loadModules();
    }
  }, [page, pageSize, showDeleted, selectedApplication]);

  const filteredModules = modules.filter(
    (m) =>
      m.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
      (m.route || '').toLowerCase().includes(searchTerm.toLowerCase()) ||
      (m.item || '').toLowerCase().includes(searchTerm.toLowerCase())
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

  const getStatusBadge = (status: string) => {
    switch (status) {
      case 'active':
        return <Badge className="bg-chart-4/20 border-chart-4/30 text-chart-4">Activo</Badge>;
      case 'inactive':
        return <Badge className="bg-muted border-muted-foreground/30 text-muted-foreground">Inactivo</Badge>;
      default:
        return <Badge variant="outline">{status}</Badge>;
    }
  };

  const handleEdit = (module: ModuleItem) => {
    setSelectedModule(module);
    setIsEditModalOpen(true);
  };

  const handleViewDetails = (module: ModuleItem) => {
    setSelectedModule(module);
    setIsDetailModalOpen(true);
  };

  const handleDeleteClick = (module: ModuleItem) => {
    setSelectedModule(module);
    setIsDeleteDialogOpen(true);
  };

  const handleDelete = async () => {
    if (!selectedModule) return;

    try {
      await fn_delete_module(selectedModule.id);
      toast.success('Módulo eliminado correctamente');
      setIsDeleteDialogOpen(false);
      setSelectedModule(null);
      loadModules();
    } catch (err: any) {
      toast.error(err.message || 'Error al eliminar módulo');
    }
  };

  const handleRestore = async (module: ModuleItem) => {
    try {
      await fn_restore_module(module.id);
      toast.success('Módulo restaurado correctamente');
      loadModules();
    } catch (err: any) {
      toast.error(err.message || 'Error al restaurar módulo');
    }
  };

  const selectedApp = applications.find(app => app.id === selectedApplication);

  if (error) return <p className="py-10 text-destructive text-center">{error}</p>;

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <div>
          <h1 className="font-bold text-foreground text-3xl">Módulos</h1>
          <p className="mt-1 text-muted-foreground">Gestiona los módulos y permisos de las aplicaciones</p>
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
            Nuevo Módulo
          </Button>
        </div>
      </div>

      <ModulesStatsCards />

      <Card className="bg-card/50 border-border">
        <CardHeader>
          <div className="space-y-4">
            <div className="flex justify-between items-center">
              <div>
                <CardTitle>Lista de Módulos</CardTitle>
                <CardDescription>
                  {selectedApplication 
                    ? `Mostrando ${(page - 1) * pageSize + 1} - ${Math.min(page * pageSize, totalItems)} de ${totalItems} módulos`
                    : 'Selecciona una aplicación para ver sus módulos'
                  }
                </CardDescription>
              </div>
              <div className="flex gap-2">
                <div className="relative">
                  <Search className="top-1/2 left-3 absolute w-4 h-4 text-muted-foreground -translate-y-1/2 transform" />
                  <Input 
                    placeholder="Buscar módulos..." 
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
                  <Badge variant="outline">{modules.length} módulos</Badge>
                </div>
              )}
            </div>
          </div>
        </CardHeader>

        <CardContent>
          {!selectedApplication ? (
            <div className="flex flex-col items-center justify-center py-16 text-center">
              <Layers className="w-16 h-16 text-muted-foreground/50 mb-4" />
              <h3 className="text-lg font-semibold mb-2">Selecciona una Aplicación</h3>
              <p className="text-sm text-muted-foreground max-w-md">
                Para ver y gestionar los módulos, primero debes seleccionar una aplicación del selector superior.
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
                      <TableHead>Módulo Padre</TableHead>
                      <TableHead>Orden</TableHead>
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
                    ) : filteredModules.length === 0 ? (
                      <TableRow>
                        <TableCell colSpan={6} className="text-center text-muted-foreground py-8">
                          No se encontraron módulos para esta aplicación
                        </TableCell>
                      </TableRow>
                    ) : (
                      filteredModules.map((module) => (
                        <TableRow key={module.id} className={module.deleted_at ? 'opacity-60' : ''}>
                          <TableCell>
                            <div className="flex items-center gap-2">
                              <LayoutGrid className="w-4 h-4 text-primary" />
                              <div>
                                <p className="font-medium text-foreground">{module.name}</p>
                                {module.item && <p className="text-xs text-muted-foreground">{module.item}</p>}
                              </div>
                            </div>
                          </TableCell>
                          <TableCell>
                            <code className="text-xs bg-muted px-2 py-1 rounded">{module.route}</code>
                          </TableCell>
                          <TableCell>
                            {module.parent ? (
                              <span className="text-sm">{module.parent.name}</span>
                            ) : (
                              <span className="text-muted-foreground text-sm">Raíz</span>
                            )}
                          </TableCell>
                          <TableCell>
                            <Badge variant="outline">{module.sort_order}</Badge>
                          </TableCell>
                          <TableCell>
                            {module.deleted_at ? <Badge variant="destructive">Eliminado</Badge> : getStatusBadge(module.status)}
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
                                <DropdownMenuItem onClick={() => handleViewDetails(module)}>
                                  <Eye className="mr-2 w-4 h-4" />
                                  Ver Detalles
                                </DropdownMenuItem>
                                {!module.deleted_at && (
                                  <>
                                    <DropdownMenuItem onClick={() => handleEdit(module)}>
                                      <Edit className="mr-2 w-4 h-4" />
                                      Editar
                                    </DropdownMenuItem>
                                    <DropdownMenuItem className="text-destructive" onClick={() => handleDeleteClick(module)}>
                                      <Trash2 className="mr-2 w-4 h-4" />
                                      Eliminar
                                    </DropdownMenuItem>
                                  </>
                                )}
                                {module.deleted_at && (
                                  <DropdownMenuItem onClick={() => handleRestore(module)}>
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
      <ModuleModal 
        open={isCreateModalOpen} 
        onOpenChange={setIsCreateModalOpen} 
        onSuccess={loadModules}
        defaultApplicationId={selectedApplication}
      />

      {/* Modal Editar */}
      <ModuleModal 
        open={isEditModalOpen} 
        onOpenChange={setIsEditModalOpen} 
        module={selectedModule} 
        onSuccess={loadModules}
        defaultApplicationId={selectedApplication}
      />

      {/* Modal Detalles */}
      <Dialog open={isDetailModalOpen} onOpenChange={setIsDetailModalOpen}>
        <DialogContent className="sm:max-w-[600px] max-h-[90vh] overflow-y-auto">
          <DialogHeader>
            <DialogTitle>Detalles del Módulo</DialogTitle>
            <DialogDescription>Información completa del módulo del sistema</DialogDescription>
          </DialogHeader>
          {selectedModule && (
            <div className="space-y-4">
              <div className="grid grid-cols-2 gap-4">
                <div>
                  <p className="text-sm text-muted-foreground">Nombre</p>
                  <p className="font-medium text-lg">{selectedModule.name}</p>
                </div>
                <div>
                  <p className="text-sm text-muted-foreground">Estado</p>
                  <div className="mt-1">{getStatusBadge(selectedModule.status)}</div>
                </div>
              </div>

              {selectedModule.item && (
                <div>
                  <p className="text-sm text-muted-foreground">Item</p>
                  <code className="text-sm bg-muted px-2 py-1 rounded">{selectedModule.item}</code>
                </div>
              )}

              <div>
                <p className="text-sm text-muted-foreground">Ruta</p>
                <code className="text-sm bg-muted px-2 py-1 rounded block mt-1">{selectedModule.route}</code>
              </div>

              {selectedModule.icon && (
                <div>
                  <p className="text-sm text-muted-foreground">Icono</p>
                  <p className="font-medium">{selectedModule.icon}</p>
                </div>
              )}

              <div className="grid grid-cols-2 gap-4">
                <div>
                  <p className="text-sm text-muted-foreground">Orden</p>
                  <Badge variant="outline" className="mt-1">{selectedModule.sort_order}</Badge>
                </div>
                <div>
                  <p className="text-sm text-muted-foreground">Usuarios Asignados</p>
                  <p className="font-medium">{selectedModule.users_count || 0}</p>
                </div>
              </div>

              {selectedModule.application_name && (
                <div>
                  <p className="text-sm text-muted-foreground">Aplicación</p>
                  <div className="mt-1">
                    <p className="font-medium">{selectedModule.application_name}</p>
                    <p className="text-xs text-muted-foreground">{selectedModule.application_client_id}</p>
                  </div>
                </div>
              )}

              {selectedModule.parent && (
                <div>
                  <p className="text-sm text-muted-foreground">Módulo Padre</p>
                  <p className="font-medium">{selectedModule.parent.name}</p>
                </div>
              )}

              <div className="grid grid-cols-2 gap-4 pt-2 border-t">
                <div>
                  <p className="text-sm text-muted-foreground">Creado</p>
                  <p className="text-sm">{new Date(selectedModule.created_at).toLocaleString('es-PE')}</p>
                </div>
                <div>
                  <p className="text-sm text-muted-foreground">Actualizado</p>
                  <p className="text-sm">{new Date(selectedModule.updated_at).toLocaleString('es-PE')}</p>
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
            <DialogTitle>¿Eliminar módulo?</DialogTitle>
            <DialogDescription>
              Esta acción eliminará el módulo <strong>{selectedModule?.name}</strong>. Esta acción se puede revertir.
            </DialogDescription>
          </DialogHeader>
          <div className="bg-yellow-500/10 border border-yellow-500/20 rounded-lg p-4 flex gap-3">
            <AlertCircle className="w-5 h-5 text-yellow-500 shrink-0 mt-0.5" />
            <div className="text-sm">
              <p className="font-medium text-yellow-600 dark:text-yellow-500 mb-1">Advertencia</p>
              <p className="text-muted-foreground">Si el módulo tiene submódulos, no podrá ser eliminado.</p>
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