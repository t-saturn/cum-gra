'use client';

import { useEffect, useState } from 'react';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Badge } from '@/components/ui/badge';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table';
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuLabel, DropdownMenuSeparator, DropdownMenuTrigger } from '@/components/ui/dropdown-menu';
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle, DialogTrigger } from '@/components/ui/dialog';
import { Label } from '@/components/ui/label';
import { Textarea } from '@/components/ui/textarea';
import { Search, Plus, Filter, Download, MoreHorizontal, Edit, Trash2, Eye, Loader2, Users } from 'lucide-react';
import { fn_get_organic_units } from '@/actions/units/fn_get_organic_units';
import { OrganicUnitItemDTO } from '@/types/units';
import { OrganicUnitsStatsCards } from '@/components/custom/card/organic-units-stats-cards';

export default function OrganicUnitsManagement() {
  const [searchTerm, setSearchTerm] = useState('');
  const [isCreateDialogOpen, setIsCreateDialogOpen] = useState(false);
  const [selectedUnit, setSelectedUnit] = useState<OrganicUnitItemDTO | null>(null);
  const [isEditDialogOpen, setIsEditDialogOpen] = useState(false);
  const [units, setUnits] = useState<OrganicUnitItemDTO[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const loadUnits = async () => {
      try {
        setLoading(true);
        const data = await fn_get_organic_units(1, 40, false);
        setUnits(data.data);
      } catch (err: any) {
        console.error('Error al cargar unidades:', err);
        setError('No se pudieron cargar las unidades orgánicas.');
      } finally {
        setLoading(false);
      }
    };
    loadUnits();
  }, []);

  const filteredUnits = units.filter(
    (unit) => unit.name.toLowerCase().includes(searchTerm.toLowerCase()) || (unit.description ?? '').toLowerCase().includes(searchTerm.toLowerCase()),
  );

  const getStatusBadge = (active: boolean) => {
    return active ? (
      <Badge className="bg-chart-4/20 border-chart-4/30 text-chart-4">Activa</Badge>
    ) : (
      <Badge className="bg-muted border-muted-foreground/30 text-muted-foreground">Inactiva</Badge>
    );
  };

  if (loading) {
    return (
      <div className="flex justify-center items-center h-64">
        <Loader2 className="w-8 h-8 text-primary animate-spin" />
      </div>
    );
  }

  if (error) {
    return <div className="py-12 text-muted-foreground text-center">{error}</div>;
  }

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <div>
          <h1 className="font-bold text-foreground text-3xl">Unidades Orgánicas</h1>
          <p className="mt-1 text-muted-foreground">Gestiona la estructura organizacional de la entidad</p>
        </div>
        <div className="flex gap-3">
          <Button variant="outline">
            <Download className="mr-2 w-4 h-4" />
            Exportar
          </Button>
          <Dialog open={isCreateDialogOpen} onOpenChange={setIsCreateDialogOpen}>
            <DialogTrigger asChild>
              <Button className="bg-gradient-to-r from-primary hover:from-primary/90 to-chart-1 hover:to-chart-1/90 shadow-lg shadow-primary/25">
                <Plus className="mr-2 w-4 h-4" />
                Nueva Unidad
              </Button>
            </DialogTrigger>
            <DialogContent className="bg-card/80 backdrop-blur-xl border-border sm:max-w-[600px]">
              <DialogHeader>
                <DialogTitle>Crear Nueva Unidad Orgánica</DialogTitle>
                <DialogDescription>Completa la información para crear una nueva unidad orgánica en el sistema.</DialogDescription>
              </DialogHeader>
              <div className="gap-4 grid py-4">
                <div className="space-y-2">
                  <Label htmlFor="name">Nombre de la Unidad</Label>
                  <Input id="name" placeholder="Ej: Gerencia de Ventas" />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="description">Descripción</Label>
                  <Textarea id="description" placeholder="Describe las funciones y responsabilidades..." />
                </div>
              </div>
              <DialogFooter>
                <Button variant="outline" onClick={() => setIsCreateDialogOpen(false)}>
                  Cancelar
                </Button>
                <Button className="bg-gradient-to-r from-primary to-chart-1">Crear Unidad</Button>
              </DialogFooter>
            </DialogContent>
          </Dialog>
        </div>
      </div>

      <OrganicUnitsStatsCards />

      <Card className="bg-card/50 border-border">
        <CardHeader>
          <div className="flex justify-between items-center">
            <div>
              <CardTitle>Lista de Unidades Orgánicas</CardTitle>
              <CardDescription>
                {filteredUnits.length} de {units.length} unidades
              </CardDescription>
            </div>
            <div className="flex gap-2">
              <div className="relative">
                <Search className="top-1/2 left-3 absolute w-4 h-4 text-muted-foreground -translate-y-1/2 transform" />
                <Input
                  placeholder="Buscar unidades..."
                  value={searchTerm}
                  onChange={(e) => setSearchTerm(e.target.value)}
                  className="bg-background/50 pl-10 focus:border-primary border-border focus:ring-ring w-80"
                />
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
                  <TableHead>Unidad Orgánica</TableHead>
                  <TableHead>Acrónimo</TableHead>
                  <TableHead>Empleados</TableHead>
                  <TableHead>Estado</TableHead>
                  <TableHead className="text-right">Acciones</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {filteredUnits.map((unit) => (
                  <TableRow key={unit.id} className="hover:bg-accent/30">
                    <TableCell>
                      <div className="space-y-1">
                        <p className="font-medium text-foreground">{unit.name}</p>
                        {unit.description && <p className="text-muted-foreground text-sm line-clamp-2">{unit.description}</p>}
                      </div>
                    </TableCell>
                    <TableCell className="font-mono text-muted-foreground text-sm">{unit.acronym}</TableCell>
                    <TableCell className="font-medium">
                      <div className="space-y-1">
                        <div className="flex items-center gap-2">
                          <Users className="w-4 h-4 text-chart-2" />
                          <span className="font-medium">{unit.users_count}</span>
                        </div>
                      </div>
                    </TableCell>
                    <TableCell>{getStatusBadge(unit.is_active)}</TableCell>
                    <TableCell className="text-right">
                      <DropdownMenu>
                        <DropdownMenuTrigger asChild>
                          <Button variant="ghost" size="sm">
                            <MoreHorizontal className="w-4 h-4" />
                          </Button>
                        </DropdownMenuTrigger>
                        <DropdownMenuContent align="end" className="bg-card/80 backdrop-blur-xl border-border">
                          <DropdownMenuLabel>Acciones</DropdownMenuLabel>
                          <DropdownMenuSeparator />
                          <DropdownMenuItem>
                            <Eye className="mr-2 w-4 h-4" />
                            Ver Detalles
                          </DropdownMenuItem>
                          <DropdownMenuItem onClick={() => setSelectedUnit(unit)}>
                            <Edit className="mr-2 w-4 h-4" />
                            Editar
                          </DropdownMenuItem>
                          <DropdownMenuSeparator />
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
