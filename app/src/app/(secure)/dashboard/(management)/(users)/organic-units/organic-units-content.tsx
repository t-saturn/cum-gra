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
import { Search, Plus, Filter, Download, MoreHorizontal, Edit, Trash2, Eye, Users, Building2, AlertCircle, ChevronLeft, ChevronRight, Upload } from 'lucide-react';
import { toast } from 'sonner';
import { fn_get_organic_units } from '@/actions/units/fn_get_organic_units';
import { fn_delete_organic_unit } from '@/actions/units/fn_delete_organic_unit';
import { fn_restore_organic_unit } from '@/actions/units/fn_restore_organic_unit';
import type { OrganicUnitItemDTO } from '@/types/units';
import { OrganicUnitsStatsCards } from '@/components/custom/card/organic-units-stats-cards';
import OrganicUnitModal from './organic-unit-modal';
import BulkUploadOrganicUnitsModal from './bulk-upload-modal';

export default function OrganicUnitsContent() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const [isPending, startTransition] = useTransition();

  // Obtener params de URL
  const page = Number(searchParams.get('page')) || 1;
  const pageSize = Number(searchParams.get('page_size')) || 10;
  const showDeleted = searchParams.get('deleted') === 'true';

  const [units, setUnits] = useState<OrganicUnitItemDTO[]>([]);
  const [searchTerm, setSearchTerm] = useState('');
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [totalItems, setTotalItems] = useState(0);

  // Modals
  const [isCreateModalOpen, setIsCreateModalOpen] = useState(false);
  const [isEditModalOpen, setIsEditModalOpen] = useState(false);
  const [isDetailModalOpen, setIsDetailModalOpen] = useState(false);
  const [isDeleteDialogOpen, setIsDeleteDialogOpen] = useState(false);
  const [selectedUnit, setSelectedUnit] = useState<OrganicUnitItemDTO | null>(null);
  const [isBulkUploadModalOpen, setIsBulkUploadModalOpen] = useState(false);

  const loadUnits = async () => {
    try {
      setLoading(true);
      setError(null);
      const response = await fn_get_organic_units(page, pageSize, showDeleted);
      setUnits(response.data);
      setTotalItems(response.total);
    } catch (err: any) {
      setError(err.message ?? 'Error desconocido');
      toast.error('Error al cargar unidades orgánicas');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    loadUnits();
  }, [page, pageSize, showDeleted]);

  const filteredUnits = units.filter(
    (u) =>
      u.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
      u.acronym.toLowerCase().includes(searchTerm.toLowerCase()) ||
      (u.description ?? '').toLowerCase().includes(searchTerm.toLowerCase())
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

  const getStatusBadge = (active: boolean) =>
    active ? (
      <Badge className="bg-chart-4/20 border-chart-4/30 text-chart-4">Activa</Badge>
    ) : (
      <Badge className="bg-muted border-muted-foreground/30 text-muted-foreground">Inactiva</Badge>
    );

  const handleEdit = (u: OrganicUnitItemDTO) => {
    setSelectedUnit(u);
    setIsEditModalOpen(true);
  };

  const handleViewDetails = (u: OrganicUnitItemDTO) => {
    setSelectedUnit(u);
    setIsDetailModalOpen(true);
  };

  const handleDeleteClick = (u: OrganicUnitItemDTO) => {
    setSelectedUnit(u);
    setIsDeleteDialogOpen(true);
  };

  const handleDelete = async () => {
    if (!selectedUnit) return;

    try {
      await fn_delete_organic_unit(selectedUnit.id);
      toast.success('Unidad orgánica eliminada correctamente');
      setIsDeleteDialogOpen(false);
      setSelectedUnit(null);
      loadUnits();
    } catch (err: any) {
      toast.error(err.message || 'Error al eliminar unidad orgánica');
    }
  };

  const handleRestore = async (u: OrganicUnitItemDTO) => {
    try {
      await fn_restore_organic_unit(u.id);
      toast.success('Unidad orgánica restaurada correctamente');
      loadUnits();
    } catch (err: any) {
      toast.error(err.message || 'Error al restaurar unidad orgánica');
    }
  };

  if (error) return <p className="py-10 text-destructive text-center">{error}</p>;

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <div>
          <h1 className="font-bold text-foreground text-3xl">Unidades Orgánicas</h1>
          <p className="mt-1 text-muted-foreground">Gestiona la estructura organizacional de la entidad</p>
        </div>
        <div className="flex gap-3">
          <Button variant="outline" onClick={() => setIsBulkUploadModalOpen(true)}>
            <Upload className="mr-2 w-4 h-4" />
            Carga Masiva
          </Button>
          <Button className="bg-linear-to-r from-primary to-chart-1" onClick={() => setIsCreateModalOpen(true)}>
            <Plus className="mr-2 w-4 h-4" />
            Nueva Unidad
          </Button>
        </div>
      </div>

      <OrganicUnitsStatsCards />

      <Card className="bg-card/50 border-border">
        <CardHeader>
          <div className="flex justify-between items-center">
            <div>
              <CardTitle>Lista de Unidades Orgánicas</CardTitle>
              <CardDescription>
                Mostrando {(page - 1) * pageSize + 1} - {Math.min(page * pageSize, totalItems)} de {totalItems} unidades
              </CardDescription>
            </div>
            <div className="flex gap-2">
              <div className="relative">
                <Search className="top-1/2 left-3 absolute w-4 h-4 text-muted-foreground -translate-y-1/2 transform" />
                <Input placeholder="Buscar unidades..." value={searchTerm} onChange={(e) => setSearchTerm(e.target.value)} className="bg-background/50 pl-10 w-80" />
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
                  <TableHead>Unidad Orgánica</TableHead>
                  <TableHead>Acrónimo</TableHead>
                  <TableHead>Empleados</TableHead>
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
                ) : filteredUnits.length === 0 ? (
                  <TableRow>
                    <TableCell colSpan={5} className="text-center text-muted-foreground py-8">
                      No se encontraron unidades orgánicas
                    </TableCell>
                  </TableRow>
                ) : (
                  filteredUnits.map((u) => (
                    <TableRow key={u.id} className={u.is_deleted ? 'opacity-60' : ''}>
                      <TableCell>
                        <div className="space-y-1">
                          <div className="flex items-center gap-2">
                            <Building2 className="w-4 h-4 text-primary" />
                            <p className="font-medium text-foreground">{u.name}</p>
                          </div>
                          {u.description && <p className="text-muted-foreground text-sm line-clamp-2">{u.description}</p>}
                        </div>
                      </TableCell>
                      <TableCell>
                        <code className="text-xs bg-muted px-2 py-1 rounded">{u.acronym}</code>
                      </TableCell>
                      <TableCell>
                        <div className="flex items-center gap-2">
                          <Users className="w-4 h-4 text-chart-2" />
                          {u.users_count}
                        </div>
                      </TableCell>
                      <TableCell>{u.is_deleted ? <Badge variant="destructive">Eliminada</Badge> : getStatusBadge(u.is_active)}</TableCell>
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
                            <DropdownMenuItem onClick={() => handleViewDetails(u)}>
                              <Eye className="mr-2 w-4 h-4" />
                              Ver Detalles
                            </DropdownMenuItem>
                            {!u.is_deleted && (
                              <>
                                <DropdownMenuItem onClick={() => handleEdit(u)}>
                                  <Edit className="mr-2 w-4 h-4" />
                                  Editar
                                </DropdownMenuItem>
                                <DropdownMenuItem className="text-destructive" onClick={() => handleDeleteClick(u)}>
                                  <Trash2 className="mr-2 w-4 h-4" />
                                  Eliminar
                                </DropdownMenuItem>
                              </>
                            )}
                            {u.is_deleted && (
                              <DropdownMenuItem onClick={() => handleRestore(u)}>
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
      <OrganicUnitModal open={isCreateModalOpen} onOpenChange={setIsCreateModalOpen} onSuccess={loadUnits} />

      {/* Modal Editar */}
      <OrganicUnitModal open={isEditModalOpen} onOpenChange={setIsEditModalOpen} unit={selectedUnit} onSuccess={loadUnits} />

      {/* Modal Subir Masivo */}
      <BulkUploadOrganicUnitsModal
        open={isBulkUploadModalOpen}
        onOpenChange={setIsBulkUploadModalOpen}
        onSuccess={loadUnits}
      />

      {/* Modal Detalles */}
      <Dialog open={isDetailModalOpen} onOpenChange={setIsDetailModalOpen}>
        <DialogContent className="sm:max-w-[600px]">
          <DialogHeader>
            <DialogTitle>Detalles de la Unidad Orgánica</DialogTitle>
            <DialogDescription>Información completa de la unidad organizacional</DialogDescription>
          </DialogHeader>
          {selectedUnit && (
            <div className="space-y-4">
              <div className="grid grid-cols-2 gap-4">
                <div>
                  <p className="text-sm text-muted-foreground">Nombre</p>
                  <p className="font-medium">{selectedUnit.name}</p>
                </div>
                <div>
                  <p className="text-sm text-muted-foreground">Acrónimo</p>
                  <code className="text-sm bg-muted px-2 py-1 rounded">{selectedUnit.acronym}</code>
                </div>
              </div>

              <div className="grid grid-cols-2 gap-4">
                <div>
                  <p className="text-sm text-muted-foreground">Marca</p>
                  {selectedUnit.brand ? (
                    selectedUnit.brand.startsWith('http') ? (
                      <a
                        href={selectedUnit.brand}
                        target="_blank"
                        rel="noopener noreferrer"
                        className="text-sm text-primary hover:underline break-all"
                      >
                        {selectedUnit.brand}
                      </a>
                    ) : (
                      <p className="font-medium break-all">{selectedUnit.brand}</p>
                    )
                  ) : (
                    <p className="font-medium">N/A</p>
                  )}
                </div>
                <div>
                  <p className="text-sm text-muted-foreground">Código SGD</p>
                  <p className="font-medium">{selectedUnit.cod_dep_sgd || 'N/A'}</p>
                </div>
              </div>

              <div>
                <p className="text-sm text-muted-foreground">Descripción</p>
                <p className="font-medium">{selectedUnit.description || 'Sin descripción'}</p>
              </div>

              <div className="grid grid-cols-2 gap-4">
                <div>
                  <p className="text-sm text-muted-foreground">Empleados Asignados</p>
                  <p className="font-medium flex items-center gap-2">
                    <Users className="w-4 h-4" />
                    {selectedUnit.users_count}
                  </p>
                </div>
                <div>
                  <p className="text-sm text-muted-foreground">Estado</p>
                  <div className="mt-1">{getStatusBadge(selectedUnit.is_active)}</div>
                </div>
              </div>

              <div className="grid grid-cols-2 gap-4 pt-2 border-t">
                <div>
                  <p className="text-sm text-muted-foreground">Creado</p>
                  <p className="text-sm">{new Date(selectedUnit.created_at).toLocaleString('es-PE')}</p>
                </div>
                <div>
                  <p className="text-sm text-muted-foreground">Actualizado</p>
                  <p className="text-sm">{new Date(selectedUnit.updated_at).toLocaleString('es-PE')}</p>
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
            <DialogTitle>¿Eliminar unidad orgánica?</DialogTitle>
            <DialogDescription>
              Esta acción eliminará la unidad <strong>{selectedUnit?.name}</strong>.{' '}
              {selectedUnit?.users_count ? 'Esta unidad tiene empleados asignados.' : 'Esta acción se puede revertir.'}
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