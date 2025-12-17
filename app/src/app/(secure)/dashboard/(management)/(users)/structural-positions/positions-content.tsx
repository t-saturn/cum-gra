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
import { Search, Plus, Filter, Download, MoreHorizontal, Edit, Trash2, Eye, Users, CircleQuestionMark, AlertCircle, ChevronLeft, ChevronRight, Crown, BadgeCheck, Briefcase, Award, Shield } from 'lucide-react';
import { toast } from 'sonner';
import { fn_get_positions } from '@/actions/positions/fn_get_positions';
import { fn_delete_position } from '@/actions/positions/fn_delete_position';
import { fn_restore_position } from '@/actions/positions/fn_restore_position';
import type { StructuralPositionItem } from '@/types/structural_positions';
import { PositionsStatsCards } from '@/components/custom/card/positions-stats-cards';
import PositionModal from './position-modal';

export default function StructuralPositionsContent() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const [isPending, startTransition] = useTransition();

  // Obtener params de URL
  const page = Number(searchParams.get('page')) || 1;
  const pageSize = Number(searchParams.get('page_size')) || 10;
  const showDeleted = searchParams.get('deleted') === 'true';

  const [positions, setPositions] = useState<StructuralPositionItem[]>([]);
  const [searchTerm, setSearchTerm] = useState('');
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [totalItems, setTotalItems] = useState(0);

  // Modals
  const [isCreateModalOpen, setIsCreateModalOpen] = useState(false);
  const [isEditModalOpen, setIsEditModalOpen] = useState(false);
  const [isDetailModalOpen, setIsDetailModalOpen] = useState(false);
  const [isDeleteDialogOpen, setIsDeleteDialogOpen] = useState(false);
  const [selectedPosition, setSelectedPosition] = useState<StructuralPositionItem | null>(null);

  const loadPositions = async () => {
    try {
      setLoading(true);
      setError(null);
      const response = await fn_get_positions(page, pageSize, showDeleted);
      setPositions(response.data);
      setTotalItems(response.total);
    } catch (err: any) {
      setError(err.message ?? 'Error desconocido');
      toast.error('Error al cargar posiciones');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    loadPositions();
  }, [page, pageSize, showDeleted]);

  const filteredPositions = positions.filter(
    (p) =>
      p.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
      p.code.toLowerCase().includes(searchTerm.toLowerCase()) ||
      (p.description ?? '').toLowerCase().includes(searchTerm.toLowerCase())
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

  const getLevelBadge = (level?: number | null) => {
    if (!level) return <Badge variant="outline">-</Badge>;
    return <Badge variant="secondary">Nivel {level}</Badge>;
  };

  const getPositionIcon = (level?: number | null) => {
  if (!level) return <Briefcase className="w-4 h-4 text-primary" />;
  
  switch (level) {
    case 1: return <Crown className="w-4 h-4 text-amber-500" />; // Directivo
    case 2: return <Award className="w-4 h-4 text-blue-500" />; // Ejecutivo
    case 3: return <Briefcase className="w-4 h-4 text-purple-500" />; // Profesional
    case 4: return <Shield className="w-4 h-4 text-green-500" />; // Técnico
    case 5: return <Users className="w-4 h-4 text-gray-500" />; // Apoyo
    default: return <Briefcase className="w-4 h-4 text-primary" />;
  }
};

  const handleEdit = (p: StructuralPositionItem) => {
    setSelectedPosition(p);
    setIsEditModalOpen(true);
  };

  const handleViewDetails = (p: StructuralPositionItem) => {
    setSelectedPosition(p);
    setIsDetailModalOpen(true);
  };

  const handleDeleteClick = (p: StructuralPositionItem) => {
    setSelectedPosition(p);
    setIsDeleteDialogOpen(true);
  };

  const handleDelete = async () => {
    if (!selectedPosition) return;

    try {
      await fn_delete_position(selectedPosition.id);
      toast.success('Posición eliminada correctamente');
      setIsDeleteDialogOpen(false);
      setSelectedPosition(null);
      loadPositions();
    } catch (err: any) {
      toast.error(err.message || 'Error al eliminar posición');
    }
  };

  const handleRestore = async (p: StructuralPositionItem) => {
    try {
      await fn_restore_position(p.id);
      toast.success('Posición restaurada correctamente');
      loadPositions();
    } catch (err: any) {
      toast.error(err.message || 'Error al restaurar posición');
    }
  };

  if (error) return <p className="py-10 text-destructive text-center">{error}</p>;

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <div>
          <h1 className="font-bold text-foreground text-3xl">Posiciones Estructurales</h1>
          <p className="mt-1 text-muted-foreground">Gestiona los cargos y posiciones de la estructura organizacional</p>
        </div>
        <div className="flex gap-3">
          <Button variant="outline">
            <Download className="mr-2 w-4 h-4" />
            Exportar
          </Button>
          <Button className="bg-linear-to-r from-primary to-chart-1" onClick={() => setIsCreateModalOpen(true)}>
            <Plus className="mr-2 w-4 h-4" />
            Nueva Posición
          </Button>
        </div>
      </div>

      <PositionsStatsCards />

      <Card className="bg-card/50 border-border">
        <CardHeader>
          <div className="flex justify-between items-center">
            <div>
              <CardTitle>Lista de Posiciones Estructurales</CardTitle>
              <CardDescription>
                Mostrando {(page - 1) * pageSize + 1} - {Math.min(page * pageSize, totalItems)} de {totalItems} posiciones
              </CardDescription>
            </div>
            <div className="flex gap-2">
              <div className="relative">
                <Search className="top-1/2 left-3 absolute w-4 h-4 text-muted-foreground -translate-y-1/2 transform" />
                <Input placeholder="Buscar posiciones..." value={searchTerm} onChange={(e) => setSearchTerm(e.target.value)} className="bg-background/50 pl-10 w-80" />
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
                  <TableHead>Posición</TableHead>
                  <TableHead>Código</TableHead>
                  <TableHead>Nivel</TableHead>
                  <TableHead>Empleados</TableHead>
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
                ) : filteredPositions.length === 0 ? (
                  <TableRow>
                    <TableCell colSpan={6} className="text-center text-muted-foreground py-8">
                      No se encontraron posiciones
                    </TableCell>
                  </TableRow>
                ) : (
                  filteredPositions.map((p) => (
                    <TableRow key={p.id} className={p.is_deleted ? 'opacity-60' : ''}>
                      <TableCell>
                        <div className="space-y-1">
                          <div className="flex items-center gap-2">
                            {getPositionIcon(p.level)}
                            <p className="font-medium text-foreground">{p.name}</p>
                          </div>
                          {p.description && <p className="text-muted-foreground text-sm line-clamp-2">{p.description}</p>}
                        </div>
                      </TableCell>
                      <TableCell>
                        <code className="text-xs bg-muted px-2 py-1 rounded">{p.code}</code>
                      </TableCell>
                      <TableCell>{getLevelBadge(p.level)}</TableCell>
                      <TableCell>
                        <div className="flex items-center gap-2">
                          <Users className="w-4 h-4 text-chart-2" />
                          {p.users_count}
                        </div>
                      </TableCell>
                      <TableCell>{p.is_deleted ? <Badge variant="destructive">Eliminada</Badge> : getStatusBadge(p.is_active)}</TableCell>
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
                            <DropdownMenuItem onClick={() => handleViewDetails(p)}>
                              <Eye className="mr-2 w-4 h-4" />
                              Ver Detalles
                            </DropdownMenuItem>
                            {!p.is_deleted && (
                              <>
                                <DropdownMenuItem onClick={() => handleEdit(p)}>
                                  <Edit className="mr-2 w-4 h-4" />
                                  Editar
                                </DropdownMenuItem>
                                <DropdownMenuItem className="text-destructive" onClick={() => handleDeleteClick(p)}>
                                  <Trash2 className="mr-2 w-4 h-4" />
                                  Eliminar
                                </DropdownMenuItem>
                              </>
                            )}
                            {p.is_deleted && (
                              <DropdownMenuItem onClick={() => handleRestore(p)}>
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
      <PositionModal open={isCreateModalOpen} onOpenChange={setIsCreateModalOpen} onSuccess={loadPositions} />

      {/* Modal Editar */}
      <PositionModal open={isEditModalOpen} onOpenChange={setIsEditModalOpen} position={selectedPosition} onSuccess={loadPositions} />

      {/* Modal Detalles */}
      <Dialog open={isDetailModalOpen} onOpenChange={setIsDetailModalOpen}>
        <DialogContent className="sm:max-w-[600px]">
          <DialogHeader>
            <DialogTitle>Detalles de la Posición</DialogTitle>
            <DialogDescription>Información completa de la posición estructural</DialogDescription>
          </DialogHeader>
          {selectedPosition && (
            <div className="space-y-4">
              <div className="grid grid-cols-2 gap-4">
                <div>
                  <p className="text-sm text-muted-foreground">Nombre</p>
                  <p className="font-medium">{selectedPosition.name}</p>
                </div>
                <div>
                  <p className="text-sm text-muted-foreground">Código</p>
                  <code className="text-sm bg-muted px-2 py-1 rounded">{selectedPosition.code}</code>
                </div>
              </div>

              <div className="grid grid-cols-2 gap-4">
                <div>
                  <p className="text-sm text-muted-foreground">Nivel</p>
                  <div className="mt-1">{getLevelBadge(selectedPosition.level)}</div>
                </div>
                <div>
                  <p className="text-sm text-muted-foreground">Código SGD</p>
                  <p className="font-medium">{selectedPosition.cod_car_sgd || 'N/A'}</p>
                </div>
              </div>

              <div>
                <p className="text-sm text-muted-foreground">Descripción</p>
                <p className="font-medium">{selectedPosition.description || 'Sin descripción'}</p>
              </div>

              <div className="grid grid-cols-2 gap-4">
                <div>
                  <p className="text-sm text-muted-foreground">Empleados Asignados</p>
                  <p className="font-medium flex items-center gap-2">
                    <Users className="w-4 h-4" />
                    {selectedPosition.users_count}
                  </p>
                </div>
                <div>
                  <p className="text-sm text-muted-foreground">Estado</p>
                  <div className="mt-1">{getStatusBadge(selectedPosition.is_active)}</div>
                </div>
              </div>

              <div className="grid grid-cols-2 gap-4 pt-2 border-t">
                <div>
                  <p className="text-sm text-muted-foreground">Creado</p>
                  <p className="text-sm">{new Date(selectedPosition.created_at).toLocaleString('es-PE')}</p>
                </div>
                <div>
                  <p className="text-sm text-muted-foreground">Actualizado</p>
                  <p className="text-sm">{new Date(selectedPosition.updated_at).toLocaleString('es-PE')}</p>
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
            <DialogTitle>¿Eliminar posición?</DialogTitle>
            <DialogDescription>
              Esta acción eliminará la posición <strong>{selectedPosition?.name}</strong>. {selectedPosition?.users_count ? 'Esta posición tiene empleados asignados.' : 'Esta acción se puede revertir.'}
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