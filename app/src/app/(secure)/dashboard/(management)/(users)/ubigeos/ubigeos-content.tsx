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
import { Search, Plus, Download, MoreHorizontal, Edit, Trash2, Eye, MapPin, ChevronLeft, ChevronRight, X } from 'lucide-react';
import { toast } from 'sonner';
import { fn_get_ubigeos } from '@/actions/ubigeos/fn_get_ubigeos';
import { fn_delete_ubigeo } from '@/actions/ubigeos/fn_delete_ubigeo';
import type { UbigeoItem } from '@/types/ubigeos';
import { UbigeosStatsCards } from '@/components/custom/card/ubigeos-stats-card';
import UbigeoModal from './ubigeo-modal';

export default function UbigeosContent() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const [isPending, startTransition] = useTransition();

  // Obtener params de URL
  const page = Number(searchParams.get('page')) || 1;
  const pageSize = Number(searchParams.get('page_size')) || 10;

  const [ubigeos, setUbigeos] = useState<UbigeoItem[]>([]);
  const [searchTerm, setSearchTerm] = useState('');
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [totalItems, setTotalItems] = useState(0);

  // Filtros
  const [departmentFilter, setDepartmentFilter] = useState('');
  const [provinceFilter, setProvinceFilter] = useState('');
  const [districtFilter, setDistrictFilter] = useState('');

  // Modals
  const [isCreateModalOpen, setIsCreateModalOpen] = useState(false);
  const [isEditModalOpen, setIsEditModalOpen] = useState(false);
  const [isDetailModalOpen, setIsDetailModalOpen] = useState(false);
  const [isDeleteDialogOpen, setIsDeleteDialogOpen] = useState(false);
  const [selectedUbigeo, setSelectedUbigeo] = useState<UbigeoItem | null>(null);

  const loadUbigeos = async () => {
    try {
      setLoading(true);
      setError(null);
      const filters: any = {};
      if (departmentFilter) filters.department = departmentFilter;
      if (provinceFilter) filters.province = provinceFilter;
      if (districtFilter) filters.district = districtFilter;

      const response = await fn_get_ubigeos(page, pageSize, filters);
      setUbigeos(response.data);
      setTotalItems(response.total);
    } catch (err: any) {
      setError(err.message ?? 'Error desconocido');
      toast.error('Error al cargar ubigeos');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    loadUbigeos();
  }, [page, pageSize, departmentFilter, provinceFilter, districtFilter]);

  const filteredUbigeos = ubigeos.filter(
    (u) =>
      u.ubigeo_code.toLowerCase().includes(searchTerm.toLowerCase()) ||
      u.inei_code.toLowerCase().includes(searchTerm.toLowerCase()) ||
      u.department.toLowerCase().includes(searchTerm.toLowerCase()) ||
      u.province.toLowerCase().includes(searchTerm.toLowerCase()) ||
      u.district.toLowerCase().includes(searchTerm.toLowerCase())
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

  const clearFilters = () => {
    setDepartmentFilter('');
    setProvinceFilter('');
    setDistrictFilter('');
  };

  const hasActiveFilters = departmentFilter || provinceFilter || districtFilter;

  const handleEdit = (ubigeo: UbigeoItem) => {
    setSelectedUbigeo(ubigeo);
    setIsEditModalOpen(true);
  };

  const handleViewDetails = (ubigeo: UbigeoItem) => {
    setSelectedUbigeo(ubigeo);
    setIsDetailModalOpen(true);
  };

  const handleDeleteClick = (ubigeo: UbigeoItem) => {
    setSelectedUbigeo(ubigeo);
    setIsDeleteDialogOpen(true);
  };

  const handleDelete = async () => {
    if (!selectedUbigeo) return;

    try {
      await fn_delete_ubigeo(selectedUbigeo.id);
      toast.success('Ubigeo eliminado correctamente');
      setIsDeleteDialogOpen(false);
      setSelectedUbigeo(null);
      loadUbigeos();
    } catch (err: any) {
      toast.error(err.message || 'Error al eliminar ubigeo');
    }
  };

  if (error) return <p className="py-10 text-destructive text-center">{error}</p>;

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <div>
          <h1 className="font-bold text-foreground text-3xl">Ubigeos</h1>
          <p className="mt-1 text-muted-foreground">Gestiona los códigos de ubicación geográfica del Perú</p>
        </div>
        <div className="flex gap-3">
          <Button variant="outline">
            <Download className="mr-2 w-4 h-4" />
            Exportar
          </Button>
          <Button className="bg-linear-to-r from-primary to-chart-1" onClick={() => setIsCreateModalOpen(true)}>
            <Plus className="mr-2 w-4 h-4" />
            Nuevo Ubigeo
          </Button>
        </div>
      </div>

      <UbigeosStatsCards />

      <Card className="bg-card/50 border-border">
        <CardHeader>
          <div className="flex justify-between items-center">
            <div>
              <CardTitle>Lista de Ubigeos</CardTitle>
              <CardDescription>
                Mostrando {(page - 1) * pageSize + 1} - {Math.min(page * pageSize, totalItems)} de {totalItems} ubigeos
              </CardDescription>
            </div>
            <div className="flex gap-2">
              <div className="relative">
                <Search className="top-1/2 left-3 absolute w-4 h-4 text-muted-foreground -translate-y-1/2 transform" />
                <Input placeholder="Buscar ubigeos..." value={searchTerm} onChange={(e) => setSearchTerm(e.target.value)} className="bg-background/50 pl-10 w-80" />
              </div>
            </div>
          </div>

          {/* Filtros */}
          <div className="flex flex-wrap gap-2 mt-4">
            <Input
              placeholder="Filtrar por departamento"
              value={departmentFilter}
              onChange={(e) => setDepartmentFilter(e.target.value)}
              className="w-[200px]"
            />
            <Input
              placeholder="Filtrar por provincia"
              value={provinceFilter}
              onChange={(e) => setProvinceFilter(e.target.value)}
              className="w-[200px]"
            />
            <Input
              placeholder="Filtrar por distrito"
              value={districtFilter}
              onChange={(e) => setDistrictFilter(e.target.value)}
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
                  <TableHead>Código Ubigeo</TableHead>
                  <TableHead>Código INEI</TableHead>
                  <TableHead>Departamento</TableHead>
                  <TableHead>Provincia</TableHead>
                  <TableHead>Distrito</TableHead>
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
                ) : filteredUbigeos.length === 0 ? (
                  <TableRow>
                    <TableCell colSpan={6} className="text-center text-muted-foreground py-8">
                      No se encontraron ubigeos
                    </TableCell>
                  </TableRow>
                ) : (
                  filteredUbigeos.map((ubigeo) => (
                    <TableRow key={ubigeo.id}>
                      <TableCell>
                        <div className="flex items-center gap-2">
                          <MapPin className="w-4 h-4 text-primary" />
                          <code className="text-sm font-mono bg-muted px-2 py-1 rounded">{ubigeo.ubigeo_code}</code>
                        </div>
                      </TableCell>
                      <TableCell>
                        <code className="text-sm font-mono">{ubigeo.inei_code}</code>
                      </TableCell>
                      <TableCell>
                        <Badge variant="outline">{ubigeo.department}</Badge>
                      </TableCell>
                      <TableCell>{ubigeo.province}</TableCell>
                      <TableCell className="font-medium">{ubigeo.district}</TableCell>
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
                            <DropdownMenuItem onClick={() => handleViewDetails(ubigeo)}>
                              <Eye className="mr-2 w-4 h-4" />
                              Ver Detalles
                            </DropdownMenuItem>
                            <DropdownMenuItem onClick={() => handleEdit(ubigeo)}>
                              <Edit className="mr-2 w-4 h-4" />
                              Editar
                            </DropdownMenuItem>
                            <DropdownMenuItem className="text-destructive" onClick={() => handleDeleteClick(ubigeo)}>
                              <Trash2 className="mr-2 w-4 h-4" />
                              Eliminar
                            </DropdownMenuItem>
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
      <UbigeoModal open={isCreateModalOpen} onOpenChange={setIsCreateModalOpen} onSuccess={loadUbigeos} />

      {/* Modal Editar */}
      <UbigeoModal open={isEditModalOpen} onOpenChange={setIsEditModalOpen} ubigeo={selectedUbigeo} onSuccess={loadUbigeos} />

      {/* Modal Detalles */}
      <Dialog open={isDetailModalOpen} onOpenChange={setIsDetailModalOpen}>
        <DialogContent className="sm:max-w-[600px]">
          <DialogHeader>
            <DialogTitle>Detalles del Ubigeo</DialogTitle>
            <DialogDescription>Información completa del código de ubicación geográfica</DialogDescription>
          </DialogHeader>
          {selectedUbigeo && (
            <div className="space-y-4">
              <div className="grid grid-cols-2 gap-4">
                <div>
                  <p className="text-sm text-muted-foreground">Código Ubigeo</p>
                  <code className="text-sm font-mono bg-muted px-2 py-1 rounded">{selectedUbigeo.ubigeo_code}</code>
                </div>
                <div>
                  <p className="text-sm text-muted-foreground">Código INEI</p>
                  <code className="text-sm font-mono bg-muted px-2 py-1 rounded">{selectedUbigeo.inei_code}</code>
                </div>
              </div>

              <div>
                <p className="text-sm text-muted-foreground">Departamento</p>
                <p className="font-medium text-lg">{selectedUbigeo.department}</p>
              </div>

              <div>
                <p className="text-sm text-muted-foreground">Provincia</p>
                <p className="font-medium text-lg">{selectedUbigeo.province}</p>
              </div>

              <div>
                <p className="text-sm text-muted-foreground">Distrito</p>
                <p className="font-medium text-lg">{selectedUbigeo.district}</p>
              </div>

              <div className="grid grid-cols-2 gap-4 pt-2 border-t">
                <div>
                  <p className="text-sm text-muted-foreground">Creado</p>
                  <p className="text-sm">{new Date(selectedUbigeo.created_at).toLocaleString('es-PE')}</p>
                </div>
                <div>
                  <p className="text-sm text-muted-foreground">Actualizado</p>
                  <p className="text-sm">{new Date(selectedUbigeo.updated_at).toLocaleString('es-PE')}</p>
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
            <DialogTitle>¿Eliminar ubigeo?</DialogTitle>
            <DialogDescription>
              Esta acción eliminará el ubigeo <strong>{selectedUbigeo?.district}</strong> ({selectedUbigeo?.ubigeo_code}).
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