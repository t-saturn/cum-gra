'use client';

import { useState } from 'react';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Badge } from '@/components/ui/badge';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table';
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuLabel, DropdownMenuSeparator, DropdownMenuTrigger } from '@/components/ui/dropdown-menu';
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle, DialogTrigger } from '@/components/ui/dialog';
import { Label } from '@/components/ui/label';
import { Textarea } from '@/components/ui/textarea';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { Search, Plus, Filter, Download, MoreHorizontal, Edit, Trash2, Eye, Users, MapPin, Phone, Mail, User } from 'lucide-react';
import CardStatsContain from '@/components/custom/card/card-stats-contain';
import { statsOrganicUnits } from '@/mocks/stats-mocks';

// Mock data
const organicUnits = [
  {
    id: '1',
    name: 'Gerencia General',
    description: 'Unidad orgánica principal encargada de la dirección estratégica de la empresa',
    parentUnit: null,
    manager: {
      name: 'Juan Carlos Pérez',
      email: 'jperez@empresa.com',
      phone: '+51 999 888 777',
    },
    location: 'Piso 10 - Torre Principal',
    status: 'active',
    employeeCount: 5,
    budget: 2500000,
    createdAt: '2024-01-01',
    updatedAt: '2024-01-15',
  },
  {
    id: '2',
    name: 'Gerencia de Sistemas',
    description: 'Responsable de la gestión y desarrollo de sistemas tecnológicos',
    parentUnit: 'Gerencia General',
    manager: {
      name: 'María García López',
      email: 'mgarcia@empresa.com',
      phone: '+51 888 777 666',
    },
    location: 'Piso 8 - Torre Principal',
    status: 'active',
    employeeCount: 25,
    budget: 800000,
    createdAt: '2024-01-02',
    updatedAt: '2024-01-14',
  },
  {
    id: '3',
    name: 'Desarrollo de Software',
    description: 'Equipo encargado del desarrollo y mantenimiento de aplicaciones',
    parentUnit: 'Gerencia de Sistemas',
    manager: {
      name: 'Carlos López Ruiz',
      email: 'clopez@empresa.com',
      phone: '+51 777 666 555',
    },
    location: 'Piso 7 - Torre Principal',
    status: 'active',
    employeeCount: 15,
    budget: 450000,
    createdAt: '2024-01-03',
    updatedAt: '2024-01-13',
  },
  {
    id: '4',
    name: 'Infraestructura TI',
    description: 'Gestión de infraestructura tecnológica y redes',
    parentUnit: 'Gerencia de Sistemas',
    manager: {
      name: 'Ana Martínez Silva',
      email: 'amartinez@empresa.com',
      phone: '+51 666 555 444',
    },
    location: 'Piso 6 - Torre Principal',
    status: 'active',
    employeeCount: 10,
    budget: 350000,
    createdAt: '2024-01-04',
    updatedAt: '2024-01-12',
  },
  {
    id: '5',
    name: 'Gerencia de Recursos Humanos',
    description: 'Gestión del talento humano y desarrollo organizacional',
    parentUnit: 'Gerencia General',
    manager: {
      name: 'Luis Fernando Torres',
      email: 'ltorres@empresa.com',
      phone: '+51 555 444 333',
    },
    location: 'Piso 9 - Torre Principal',
    status: 'suspended',
    employeeCount: 8,
    budget: 300000,
    createdAt: '2024-01-05',
    updatedAt: '2024-01-11',
  },
];

export default function OrganicUnitsManagement() {
  const [searchTerm, setSearchTerm] = useState('');
  const [isCreateDialogOpen, setIsCreateDialogOpen] = useState(false);
  const [selectedUnit, setSelectedUnit] = useState<(typeof organicUnits)[0] | null>(null);
  const [isEditDialogOpen, setIsEditDialogOpen] = useState(false);

  const filteredUnits = organicUnits.filter(
    (unit) =>
      unit.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
      unit.description.toLowerCase().includes(searchTerm.toLowerCase()) ||
      unit.manager.name.toLowerCase().includes(searchTerm.toLowerCase()),
  );

  const getStatusBadge = (status: string) => {
    switch (status) {
      case 'active':
        return <Badge className="bg-chart-4/20 text-chart-4 border-chart-4/30">Activa</Badge>;
      case 'suspended':
        return <Badge className="bg-chart-5/20 text-chart-5 border-chart-5/30">Suspendida</Badge>;
      case 'inactive':
        return <Badge className="bg-muted text-muted-foreground border-muted-foreground/30">Inactiva</Badge>;
      default:
        return <Badge variant="secondary">Desconocido</Badge>;
    }
  };

  const formatCurrency = (amount: number) => {
    return new Intl.NumberFormat('es-PE', {
      style: 'currency',
      currency: 'PEN',
    }).format(amount);
  };

  const handleEdit = (unit: (typeof organicUnits)[0]) => {
    setSelectedUnit(unit);
    setIsEditDialogOpen(true);
  };

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold text-foreground">Unidades Orgánicas</h1>
          <p className="text-muted-foreground mt-1">Gestiona la estructura organizacional de la empresa</p>
        </div>
        <div className="flex gap-3">
          <Button variant="outline">
            <Download className="w-4 h-4 mr-2" />
            Exportar
          </Button>
          <Dialog open={isCreateDialogOpen} onOpenChange={setIsCreateDialogOpen}>
            <DialogTrigger asChild>
              <Button className="bg-gradient-to-r from-primary to-chart-1 hover:from-primary/90 hover:to-chart-1/90 shadow-lg shadow-primary/25">
                <Plus className="w-4 h-4 mr-2" />
                Nueva Unidad
              </Button>
            </DialogTrigger>
            <DialogContent className="sm:max-w-[600px] bg-card/80 backdrop-blur-xl border-border">
              <DialogHeader>
                <DialogTitle>Crear Nueva Unidad Orgánica</DialogTitle>
                <DialogDescription>Completa la información para crear una nueva unidad orgánica en el sistema.</DialogDescription>
              </DialogHeader>
              <div className="grid gap-4 py-4">
                <div className="grid grid-cols-2 gap-4">
                  <div className="space-y-2">
                    <Label htmlFor="name">Nombre de la Unidad</Label>
                    <Input id="name" placeholder="Ej: Gerencia de Ventas" />
                  </div>
                  <div className="space-y-2">
                    <Label htmlFor="parent">Unidad Padre</Label>
                    <Select>
                      <SelectTrigger>
                        <SelectValue placeholder="Seleccionar unidad padre" />
                      </SelectTrigger>
                      <SelectContent>
                        <SelectItem value="none">Sin unidad padre</SelectItem>
                        <SelectItem value="1">Gerencia General</SelectItem>
                        <SelectItem value="2">Gerencia de Sistemas</SelectItem>
                      </SelectContent>
                    </Select>
                  </div>
                </div>
                <div className="space-y-2">
                  <Label htmlFor="description">Descripción</Label>
                  <Textarea id="description" placeholder="Describe las funciones y responsabilidades..." />
                </div>
                <div className="grid grid-cols-2 gap-4">
                  <div className="space-y-2">
                    <Label htmlFor="manager">Responsable</Label>
                    <Select>
                      <SelectTrigger>
                        <SelectValue placeholder="Seleccionar responsable" />
                      </SelectTrigger>
                      <SelectContent>
                        <SelectItem value="1">Juan Carlos Pérez</SelectItem>
                        <SelectItem value="2">María García López</SelectItem>
                        <SelectItem value="3">Carlos López Ruiz</SelectItem>
                      </SelectContent>
                    </Select>
                  </div>
                  <div className="space-y-2">
                    <Label htmlFor="location">Ubicación</Label>
                    <Input id="location" placeholder="Ej: Piso 5 - Torre Principal" />
                  </div>
                </div>
                <div className="grid grid-cols-2 gap-4">
                  <div className="space-y-2">
                    <Label htmlFor="budget">Presupuesto (PEN)</Label>
                    <Input id="budget" type="number" placeholder="0.00" />
                  </div>
                  <div className="space-y-2">
                    <Label htmlFor="status">Estado</Label>
                    <Select defaultValue="active">
                      <SelectTrigger>
                        <SelectValue />
                      </SelectTrigger>
                      <SelectContent>
                        <SelectItem value="active">Activa</SelectItem>
                        <SelectItem value="suspended">Suspendida</SelectItem>
                        <SelectItem value="inactive">Inactiva</SelectItem>
                      </SelectContent>
                    </Select>
                  </div>
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

      {/* Stats Cards */}
      <CardStatsContain stats={statsOrganicUnits} />

      {/* Filters and Search */}
      <Card className="border-border bg-card/50">
        <CardHeader>
          <div className="flex items-center justify-between">
            <div>
              <CardTitle>Lista de Unidades Orgánicas</CardTitle>
              <CardDescription>
                {filteredUnits.length} de {organicUnits.length} unidades
              </CardDescription>
            </div>
            <div className="flex gap-2">
              <div className="relative">
                <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 text-muted-foreground w-4 h-4" />
                <Input
                  placeholder="Buscar unidades..."
                  value={searchTerm}
                  onChange={(e) => setSearchTerm(e.target.value)}
                  className="pl-10 w-80 bg-background/50 border-border focus:border-primary focus:ring-ring"
                />
              </div>
              <Button variant="outline">
                <Filter className="w-4 h-4 mr-2" />
                Filtros
              </Button>
            </div>
          </div>
        </CardHeader>
        <CardContent>
          <div className="rounded-lg border border-border">
            <Table>
              <TableHeader>
                <TableRow className="bg-accent/50">
                  <TableHead>Unidad Orgánica</TableHead>
                  <TableHead>Responsable</TableHead>
                  <TableHead>Ubicación</TableHead>
                  <TableHead>Empleados</TableHead>
                  <TableHead>Presupuesto</TableHead>
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
                        <p className="text-sm text-muted-foreground line-clamp-2">{unit.description}</p>
                        {unit.parentUnit && <p className="text-xs text-muted-foreground/80">Depende de: {unit.parentUnit}</p>}
                      </div>
                    </TableCell>
                    <TableCell>
                      <div className="space-y-1">
                        <div className="flex items-center gap-2">
                          <User className="w-4 h-4 text-muted-foreground" />
                          <span className="font-medium text-sm">{unit.manager.name}</span>
                        </div>
                        <div className="flex items-center gap-2">
                          <Mail className="w-3 h-3 text-muted-foreground" />
                          <span className="text-xs text-muted-foreground">{unit.manager.email}</span>
                        </div>
                        <div className="flex items-center gap-2">
                          <Phone className="w-3 h-3 text-muted-foreground" />
                          <span className="text-xs text-muted-foreground">{unit.manager.phone}</span>
                        </div>
                      </div>
                    </TableCell>
                    <TableCell>
                      <div className="flex items-center gap-2">
                        <MapPin className="w-4 h-4 text-muted-foreground" />
                        <span className="text-sm">{unit.location}</span>
                      </div>
                    </TableCell>
                    <TableCell>
                      <div className="flex items-center gap-2">
                        <Users className="w-4 h-4 text-chart-2" />
                        <span className="font-medium">{unit.employeeCount}</span>
                      </div>
                    </TableCell>
                    <TableCell className="font-mono text-sm">{formatCurrency(unit.budget)}</TableCell>
                    <TableCell>{getStatusBadge(unit.status)}</TableCell>
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
                            <Eye className="w-4 h-4 mr-2" />
                            Ver Detalles
                          </DropdownMenuItem>
                          <DropdownMenuItem onClick={() => handleEdit(unit)}>
                            <Edit className="w-4 h-4 mr-2" />
                            Editar
                          </DropdownMenuItem>
                          <DropdownMenuItem>
                            <Users className="w-4 h-4 mr-2" />
                            Ver Empleados
                          </DropdownMenuItem>
                          <DropdownMenuSeparator />
                          <DropdownMenuItem className="text-destructive">
                            <Trash2 className="w-4 h-4 mr-2" />
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

      {/* Edit Dialog */}
      <Dialog open={isEditDialogOpen} onOpenChange={setIsEditDialogOpen}>
        <DialogContent className="sm:max-w-[600px] bg-card/80 backdrop-blur-xl border-border">
          <DialogHeader>
            <DialogTitle>Editar Unidad Orgánica</DialogTitle>
            <DialogDescription>Modifica la información de la unidad orgánica seleccionada.</DialogDescription>
          </DialogHeader>
          {selectedUnit && (
            <div className="grid gap-4 py-4">
              <div className="grid grid-cols-2 gap-4">
                <div className="space-y-2">
                  <Label htmlFor="edit-name">Nombre de la Unidad</Label>
                  <Input id="edit-name" defaultValue={selectedUnit.name} />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="edit-parent">Unidad Padre</Label>
                  <Select defaultValue={selectedUnit.parentUnit || 'none'}>
                    <SelectTrigger>
                      <SelectValue />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value="none">Sin unidad padre</SelectItem>
                      <SelectItem value="1">Gerencia General</SelectItem>
                      <SelectItem value="2">Gerencia de Sistemas</SelectItem>
                    </SelectContent>
                  </Select>
                </div>
              </div>
              <div className="space-y-2">
                <Label htmlFor="edit-description">Descripción</Label>
                <Textarea id="edit-description" defaultValue={selectedUnit.description} />
              </div>
              <div className="grid grid-cols-2 gap-4">
                <div className="space-y-2">
                  <Label htmlFor="edit-location">Ubicación</Label>
                  <Input id="edit-location" defaultValue={selectedUnit.location} />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="edit-budget">Presupuesto (PEN)</Label>
                  <Input id="edit-budget" type="number" defaultValue={selectedUnit.budget} />
                </div>
              </div>
              <div className="space-y-2">
                <Label htmlFor="edit-status">Estado</Label>
                <Select defaultValue={selectedUnit.status}>
                  <SelectTrigger>
                    <SelectValue />
                  </SelectTrigger>
                  <SelectContent>
                    <SelectItem value="active">Activa</SelectItem>
                    <SelectItem value="suspended">Suspendida</SelectItem>
                    <SelectItem value="inactive">Inactiva</SelectItem>
                  </SelectContent>
                </Select>
              </div>
            </div>
          )}
          <DialogFooter>
            <Button variant="outline" onClick={() => setIsEditDialogOpen(false)}>
              Cancelar
            </Button>
            <Button className="bg-gradient-to-r from-primary to-chart-1">Guardar Cambios</Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </div>
  );
}
