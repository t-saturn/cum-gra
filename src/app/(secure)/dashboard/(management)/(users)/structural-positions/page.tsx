'use client';

import { useState, useEffect } from 'react';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Badge } from '@/components/ui/badge';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table';
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuLabel, DropdownMenuSeparator, DropdownMenuTrigger } from '@/components/ui/dropdown-menu';
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle, DialogTrigger } from '@/components/ui/dialog';
import { Search, Plus, Filter, Download, MoreHorizontal, Edit, Trash2, Eye, Briefcase, Users, Building, Crown, Star, Award } from 'lucide-react';
import { fn_get_positions } from '@/actions/positions/fn_get_positions';
import type { StructuralPositionItem } from '@/types/structural_positions';
import { PositionsStatsCards } from '@/components/custom/card/positions-stats-cards';

export default function StructuralPositionsManagement() {
  const [positions, setPositions] = useState<StructuralPositionItem[]>([]);
  const [searchTerm, setSearchTerm] = useState('');
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const [isCreateDialogOpen, setIsCreateDialogOpen] = useState(false);
  const [selectedPosition, setSelectedPosition] = useState<StructuralPositionItem | null>(null);
  const [isEditDialogOpen, setIsEditDialogOpen] = useState(false);
  const [isDetailDialogOpen, setIsDetailDialogOpen] = useState(false);

  useEffect(() => {
    const loadPositions = async () => {
      try {
        setLoading(true);
        const response = await fn_get_positions(1, 10, false);
        setPositions(response.data);
      } catch (err: any) {
        setError(err.message ?? 'Error desconocido');
      } finally {
        setLoading(false);
      }
    };
    loadPositions();
  }, []);

  const filteredPositions = positions.filter(
    (p) => p.name.toLowerCase().includes(searchTerm.toLowerCase()) || (p.description ?? '').toLowerCase().includes(searchTerm.toLowerCase()),
  );

  const getStatusBadge = (active: boolean) =>
    active ? (
      <Badge className="bg-chart-4/20 border-chart-4/30 text-chart-4">Activa</Badge>
    ) : (
      <Badge className="bg-muted border-muted-foreground/30 text-muted-foreground">Inactiva</Badge>
    );

  const handleEdit = (p: StructuralPositionItem) => {
    setSelectedPosition(p);
    setIsEditDialogOpen(true);
  };

  const handleViewDetails = (p: StructuralPositionItem) => {
    setSelectedPosition(p);
    setIsDetailDialogOpen(true);
  };

  if (loading) return <p className="py-10 text-center">Cargando posiciones...</p>;
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
          <Dialog open={isCreateDialogOpen} onOpenChange={setIsCreateDialogOpen}>
            <DialogTrigger asChild>
              <Button className="bg-gradient-to-r from-primary to-chart-1">
                <Plus className="mr-2 w-4 h-4" />
                Nueva Posición
              </Button>
            </DialogTrigger>
            <DialogContent className="bg-card/80 backdrop-blur-xl border-border sm:max-w-[700px] max-h-[90vh] overflow-y-auto">
              <DialogHeader>
                <DialogTitle>Crear Nueva Posición</DialogTitle>
                <DialogDescription>Define una nueva posición en la estructura organizacional.</DialogDescription>
              </DialogHeader>
              <DialogFooter>
                <Button variant="outline" onClick={() => setIsCreateDialogOpen(false)}>
                  Cancelar
                </Button>
                <Button className="bg-gradient-to-r from-primary to-chart-1">Crear Posición</Button>
              </DialogFooter>
            </DialogContent>
          </Dialog>
        </div>
      </div>

      <PositionsStatsCards />

      <Card className="bg-card/50 border-border">
        <CardHeader>
          <div className="flex justify-between items-center">
            <div>
              <CardTitle>Lista de Posiciones Estructurales</CardTitle>
              <CardDescription>
                {filteredPositions.length} de {positions.length} posiciones
              </CardDescription>
            </div>
            <div className="flex gap-2">
              <div className="relative">
                <Search className="top-1/2 left-3 absolute w-4 h-4 text-muted-foreground -translate-y-1/2 transform" />
                <Input placeholder="Buscar posiciones..." value={searchTerm} onChange={(e) => setSearchTerm(e.target.value)} className="bg-background/50 pl-10 w-80" />
              </div>
              <Button variant="outline">
                <Filter className="mr-2 w-4 h-4" />
                Filtros
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
                  <TableHead>Empleados</TableHead>
                  <TableHead>Estado</TableHead>
                  <TableHead className="text-right">Acciones</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {filteredPositions.map((p) => (
                  <TableRow key={p.id}>
                    <TableCell>
                      <div className="space-y-1">
                        <p className="font-medium">{p.name}</p>
                        <p className="text-muted-foreground text-sm line-clamp-2">{p.description}</p>
                      </div>
                    </TableCell>
                    <TableCell>
                      <div className="flex items-center gap-2">
                        <Users className="w-4 h-4 text-chart-2" />
                        {p.users_count}
                      </div>
                    </TableCell>
                    <TableCell>{getStatusBadge(p.is_active)}</TableCell>
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
                          <DropdownMenuItem onClick={() => handleEdit(p)}>
                            <Edit className="mr-2 w-4 h-4" />
                            Editar
                          </DropdownMenuItem>
                          <DropdownMenuItem className="text-destructive">
                            <Trash2 className="mr-2 w-4 h-4" />
                            Eliminar
                          </DropdownMenuItem>
                        </DropdownMenuContent>
                      </DropdownMenu>
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}
