"use client";

import { useState } from "react";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Badge } from "@/components/ui/badge";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Label } from "@/components/ui/label";
import { Textarea } from "@/components/ui/textarea";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import {
  Search,
  Plus,
  Filter,
  Download,
  MoreHorizontal,
  Edit,
  Trash2,
  Eye,
  Briefcase,
  Users,
  Award,
  DollarSign,
  Building,
  Crown,
  Star,
} from "lucide-react";

// Mock data
const structuralPositions = [
  {
    id: "1",
    title: "Gerente General",
    description:
      "Máxima autoridad ejecutiva responsable de la dirección estratégica de la organización",
    level: 1,
    department: "Gerencia General",
    reportsTo: null,
    salaryRange: {
      min: 15000,
      max: 25000,
    },
    requirements: [
      "Título profesional en Administración o carreras afines",
      "Mínimo 10 años de experiencia en cargos gerenciales",
      "MBA o estudios de postgrado",
      "Liderazgo y visión estratégica",
    ],
    responsibilities: [
      "Definir la estrategia corporativa",
      "Supervisar todas las operaciones",
      "Representar a la empresa ante stakeholders",
      "Tomar decisiones estratégicas",
    ],
    status: "active",
    employeeCount: 1,
    createdAt: "2024-01-01",
    updatedAt: "2024-01-15",
  },
  {
    id: "2",
    title: "Gerente de Sistemas",
    description:
      "Responsable de liderar el área de tecnología y sistemas de información",
    level: 2,
    department: "Gerencia de Sistemas",
    reportsTo: "Gerente General",
    salaryRange: {
      min: 8000,
      max: 12000,
    },
    requirements: [
      "Título en Ingeniería de Sistemas o afines",
      "Mínimo 7 años de experiencia en TI",
      "Certificaciones en gestión de proyectos",
      "Conocimiento en arquitectura de software",
    ],
    responsibilities: [
      "Dirigir el área de sistemas",
      "Planificar proyectos tecnológicos",
      "Gestionar el equipo de desarrollo",
      "Definir arquitectura tecnológica",
    ],
    status: "active",
    employeeCount: 1,
    createdAt: "2024-01-02",
    updatedAt: "2024-01-14",
  },
  {
    id: "3",
    title: "Desarrollador Senior",
    description:
      "Especialista en desarrollo de software con experiencia avanzada",
    level: 3,
    department: "Desarrollo de Software",
    reportsTo: "Gerente de Sistemas",
    salaryRange: {
      min: 5000,
      max: 8000,
    },
    requirements: [
      "Título en Ingeniería de Software o afines",
      "Mínimo 5 años de experiencia en desarrollo",
      "Dominio de múltiples lenguajes de programación",
      "Experiencia en metodologías ágiles",
    ],
    responsibilities: [
      "Desarrollar aplicaciones complejas",
      "Mentorizar desarrolladores junior",
      "Revisar código y arquitectura",
      "Participar en diseño de soluciones",
    ],
    status: "active",
    employeeCount: 8,
    createdAt: "2024-01-03",
    updatedAt: "2024-01-13",
  },
  {
    id: "4",
    title: "Analista de Sistemas",
    description: "Especialista en análisis y diseño de sistemas de información",
    level: 3,
    department: "Desarrollo de Software",
    reportsTo: "Gerente de Sistemas",
    salaryRange: {
      min: 4000,
      max: 6000,
    },
    requirements: [
      "Título en Ingeniería de Sistemas",
      "Mínimo 3 años de experiencia",
      "Conocimiento en UML y metodologías",
      "Habilidades analíticas",
    ],
    responsibilities: [
      "Analizar requerimientos de usuario",
      "Diseñar soluciones tecnológicas",
      "Documentar procesos y sistemas",
      "Coordinar con equipos de desarrollo",
    ],
    status: "active",
    employeeCount: 5,
    createdAt: "2024-01-04",
    updatedAt: "2024-01-12",
  },
  {
    id: "5",
    title: "Especialista en Infraestructura",
    description:
      "Responsable de la gestión y mantenimiento de la infraestructura tecnológica",
    level: 3,
    department: "Infraestructura TI",
    reportsTo: "Gerente de Sistemas",
    salaryRange: {
      min: 4500,
      max: 7000,
    },
    requirements: [
      "Título en Ingeniería de Redes o afines",
      "Certificaciones en tecnologías de red",
      "Experiencia en administración de servidores",
      "Conocimiento en seguridad informática",
    ],
    responsibilities: [
      "Administrar infraestructura de red",
      "Mantener servidores y servicios",
      "Implementar medidas de seguridad",
      "Monitorear rendimiento del sistema",
    ],
    status: "active",
    employeeCount: 3,
    createdAt: "2024-01-05",
    updatedAt: "2024-01-11",
  },
  {
    id: "6",
    title: "Gerente de Recursos Humanos",
    description: "Líder del área de gestión del talento humano",
    level: 2,
    department: "Recursos Humanos",
    reportsTo: "Gerente General",
    salaryRange: {
      min: 7000,
      max: 10000,
    },
    requirements: [
      "Título en Psicología o Administración",
      "Especialización en Recursos Humanos",
      "Mínimo 6 años de experiencia",
      "Conocimiento en legislación laboral",
    ],
    responsibilities: [
      "Dirigir políticas de RRHH",
      "Gestionar procesos de selección",
      "Desarrollar programas de capacitación",
      "Administrar relaciones laborales",
    ],
    status: "suspended",
    employeeCount: 1,
    createdAt: "2024-01-06",
    updatedAt: "2024-01-10",
  },
];

export default function StructuralPositionsManagement() {
  const [searchTerm, setSearchTerm] = useState("");
  const [isCreateDialogOpen, setIsCreateDialogOpen] = useState(false);
  const [selectedPosition, setSelectedPosition] = useState<{
    id: string;
    title: string;
    description: string;
    level: number;
    department: string;
    reportsTo: string | null;
    salaryRange: {
      min: number;
      max: number;
    };
    requirements: string[];
    responsibilities: string[];
    status: string;
    employeeCount: number;
    createdAt: string;
    updatedAt: string;
  } | null>(null);
  const [isEditDialogOpen, setIsEditDialogOpen] = useState(false);
  const [isDetailDialogOpen, setIsDetailDialogOpen] = useState(false);

  const filteredPositions = structuralPositions.filter(
    (position) =>
      position.title.toLowerCase().includes(searchTerm.toLowerCase()) ||
      position.description.toLowerCase().includes(searchTerm.toLowerCase()) ||
      position.department.toLowerCase().includes(searchTerm.toLowerCase())
  );

  const getStatusBadge = (status: string) => {
    switch (status) {
      case "active":
        return (
          <Badge className="bg-chart-4/20 text-chart-4 border-chart-4/30">
            Activa
          </Badge>
        );
      case "suspended":
        return (
          <Badge className="bg-chart-5/20 text-chart-5 border-chart-5/30">
            Suspendida
          </Badge>
        );
      case "inactive":
        return (
          <Badge className="bg-muted text-muted-foreground border-muted-foreground/30">
            Inactiva
          </Badge>
        );
      default:
        return <Badge variant="secondary">Desconocido</Badge>;
    }
  };

  const getLevelBadge = (level: number) => {
    const configs = [
      { color: "bg-primary/20 text-primary border-primary/30", icon: Crown },
      { color: "bg-chart-2/20 text-chart-2 border-chart-2/30", icon: Star },
      { color: "bg-chart-3/20 text-chart-3 border-chart-3/30", icon: Award },
      {
        color: "bg-chart-4/20 text-chart-4 border-chart-4/30",
        icon: Briefcase,
      },
    ];
    const config = configs[level - 1] || configs[3];
    const Icon = config.icon;
    return (
      <Badge className={config.color}>
        <Icon className="w-3 h-3 mr-1" />
        Nivel {level}
      </Badge>
    );
  };

  const formatCurrency = (amount: number) => {
    return new Intl.NumberFormat("es-PE", {
      style: "currency",
      currency: "PEN",
    }).format(amount);
  };

  const handleEdit = (position: {
    id: string;
    title: string;
    description: string;
    level: number;
    department: string;
    reportsTo: string | null;
    salaryRange: {
      min: number;
      max: number;
    };
    requirements: string[];
    responsibilities: string[];
    status: string;
    employeeCount: number;
    createdAt: string;
    updatedAt: string;
  }) => {
    setSelectedPosition(position);
    setIsEditDialogOpen(true);
  };

  const handleViewDetails = (position: {
    id: string;
    title: string;
    description: string;
    level: number;
    department: string;
    reportsTo: string | null;
    salaryRange: {
      min: number;
      max: number;
    };
    requirements: string[];
    responsibilities: string[];
    status: string;
    employeeCount: number;
    createdAt: string;
    updatedAt: string;
  }) => {
    setSelectedPosition(position);
    setIsDetailDialogOpen(true);
  };

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold text-foreground">
            Posiciones Estructurales
          </h1>
          <p className="text-muted-foreground mt-1">
            Gestiona los cargos y posiciones de la estructura organizacional
          </p>
        </div>
        <div className="flex gap-3">
          <Button variant="outline">
            <Download className="w-4 h-4 mr-2" />
            Exportar
          </Button>
          <Dialog
            open={isCreateDialogOpen}
            onOpenChange={setIsCreateDialogOpen}
          >
            <DialogTrigger asChild>
              <Button className="bg-gradient-to-r from-primary to-chart-1 hover:from-primary/90 hover:to-chart-1/90 shadow-lg shadow-primary/25">
                <Plus className="w-4 h-4 mr-2" />
                Nueva Posición
              </Button>
            </DialogTrigger>
            <DialogContent className="sm:max-w-[700px] bg-card/80 backdrop-blur-xl border-border max-h-[90vh] overflow-y-auto">
              <DialogHeader>
                <DialogTitle>Crear Nueva Posición Estructural</DialogTitle>
                <DialogDescription>
                  Define una nueva posición en la estructura organizacional de
                  la empresa.
                </DialogDescription>
              </DialogHeader>
              <div className="grid gap-4 py-4">
                <div className="grid grid-cols-2 gap-4">
                  <div className="space-y-2">
                    <Label htmlFor="title">Título del Cargo</Label>
                    <Input id="title" placeholder="Ej: Analista Senior" />
                  </div>
                  <div className="space-y-2">
                    <Label htmlFor="department">Departamento</Label>
                    <Select>
                      <SelectTrigger>
                        <SelectValue placeholder="Seleccionar departamento" />
                      </SelectTrigger>
                      <SelectContent>
                        <SelectItem value="gerencia">
                          Gerencia General
                        </SelectItem>
                        <SelectItem value="sistemas">
                          Gerencia de Sistemas
                        </SelectItem>
                        <SelectItem value="desarrollo">
                          Desarrollo de Software
                        </SelectItem>
                        <SelectItem value="infraestructura">
                          Infraestructura TI
                        </SelectItem>
                        <SelectItem value="rrhh">Recursos Humanos</SelectItem>
                      </SelectContent>
                    </Select>
                  </div>
                </div>
                <div className="space-y-2">
                  <Label htmlFor="description">Descripción</Label>
                  <Textarea
                    id="description"
                    placeholder="Describe las funciones principales del cargo..."
                  />
                </div>
                <div className="grid grid-cols-3 gap-4">
                  <div className="space-y-2">
                    <Label htmlFor="level">Nivel Jerárquico</Label>
                    <Select>
                      <SelectTrigger>
                        <SelectValue placeholder="Nivel" />
                      </SelectTrigger>
                      <SelectContent>
                        <SelectItem value="1">Nivel 1 - Ejecutivo</SelectItem>
                        <SelectItem value="2">Nivel 2 - Gerencial</SelectItem>
                        <SelectItem value="3">
                          Nivel 3 - Especialista
                        </SelectItem>
                        <SelectItem value="4">Nivel 4 - Operativo</SelectItem>
                      </SelectContent>
                    </Select>
                  </div>
                  <div className="space-y-2">
                    <Label htmlFor="salary-min">Salario Mínimo (PEN)</Label>
                    <Input id="salary-min" type="number" placeholder="0.00" />
                  </div>
                  <div className="space-y-2">
                    <Label htmlFor="salary-max">Salario Máximo (PEN)</Label>
                    <Input id="salary-max" type="number" placeholder="0.00" />
                  </div>
                </div>
                <div className="space-y-2">
                  <Label htmlFor="reports-to">Reporta a</Label>
                  <Select>
                    <SelectTrigger>
                      <SelectValue placeholder="Seleccionar supervisor" />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value="none">
                        Sin supervisor directo
                      </SelectItem>
                      <SelectItem value="1">Gerente General</SelectItem>
                      <SelectItem value="2">Gerente de Sistemas</SelectItem>
                      <SelectItem value="6">
                        Gerente de Recursos Humanos
                      </SelectItem>
                    </SelectContent>
                  </Select>
                </div>
                <div className="space-y-2">
                  <Label htmlFor="requirements">
                    Requisitos (uno por línea)
                  </Label>
                  <Textarea
                    id="requirements"
                    placeholder="Título profesional en...&#10;Mínimo X años de experiencia...&#10;Conocimientos en..."
                    rows={4}
                  />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="responsibilities">
                    Responsabilidades (una por línea)
                  </Label>
                  <Textarea
                    id="responsibilities"
                    placeholder="Gestionar el equipo de...&#10;Desarrollar estrategias de...&#10;Supervisar procesos de..."
                    rows={4}
                  />
                </div>
              </div>
              <DialogFooter>
                <Button
                  variant="outline"
                  onClick={() => setIsCreateDialogOpen(false)}
                >
                  Cancelar
                </Button>
                <Button className="bg-gradient-to-r from-primary to-chart-1">
                  Crear Posición
                </Button>
              </DialogFooter>
            </DialogContent>
          </Dialog>
        </div>
      </div>

      {/* Stats Cards */}
      <div className="grid grid-cols-1 md:grid-cols-4 gap-4">
        <Card className="border-border bg-card/50">
          <CardContent className="p-4">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm text-muted-foreground">
                  Total Posiciones
                </p>
                <p className="text-2xl font-bold text-foreground">45</p>
              </div>
              <Briefcase className="w-8 h-8 text-chart-2" />
            </div>
          </CardContent>
        </Card>
        <Card className="border-border bg-card/50">
          <CardContent className="p-4">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm text-muted-foreground">
                  Posiciones Activas
                </p>
                <p className="text-2xl font-bold text-chart-4">42</p>
              </div>
              <Briefcase className="w-8 h-8 text-chart-4" />
            </div>
          </CardContent>
        </Card>
        <Card className="border-border bg-card/50">
          <CardContent className="p-4">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm text-muted-foreground">
                  Empleados Asignados
                </p>
                <p className="text-2xl font-bold text-primary">18</p>
              </div>
              <Users className="w-8 h-8 text-primary" />
            </div>
          </CardContent>
        </Card>
        <Card className="border-border bg-card/50">
          <CardContent className="p-4">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm text-muted-foreground">
                  Salario Promedio
                </p>
                <p className="text-2xl font-bold text-chart-5">7.2K</p>
              </div>
              <DollarSign className="w-8 h-8 text-chart-5" />
            </div>
          </CardContent>
        </Card>
      </div>

      {/* Filters and Search */}
      <Card className="border-border bg-card/50">
        <CardHeader>
          <div className="flex items-center justify-between">
            <div>
              <CardTitle>Lista de Posiciones Estructurales</CardTitle>
              <CardDescription>
                {filteredPositions.length} de {structuralPositions.length}{" "}
                posiciones
              </CardDescription>
            </div>
            <div className="flex gap-2">
              <div className="relative">
                <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 text-muted-foreground w-4 h-4" />
                <Input
                  placeholder="Buscar posiciones..."
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
                  <TableHead>Posición</TableHead>
                  <TableHead>Nivel</TableHead>
                  <TableHead>Departamento</TableHead>
                  <TableHead>Reporta a</TableHead>
                  <TableHead>Rango Salarial</TableHead>
                  <TableHead>Empleados</TableHead>
                  <TableHead>Estado</TableHead>
                  <TableHead className="text-right">Acciones</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {filteredPositions.map((position) => (
                  <TableRow key={position.id} className="hover:bg-accent/30">
                    <TableCell>
                      <div className="space-y-1">
                        <p className="font-medium text-foreground">
                          {position.title}
                        </p>
                        <p className="text-sm text-muted-foreground line-clamp-2">
                          {position.description}
                        </p>
                      </div>
                    </TableCell>
                    <TableCell>{getLevelBadge(position.level)}</TableCell>
                    <TableCell>
                      <div className="flex items-center gap-2">
                        <Building className="w-4 h-4 text-muted-foreground" />
                        <span className="text-sm">{position.department}</span>
                      </div>
                    </TableCell>
                    <TableCell>
                      <span className="text-sm">
                        {position.reportsTo || "Sin supervisor"}
                      </span>
                    </TableCell>
                    <TableCell className="font-mono text-sm">
                      <div className="space-y-1">
                        <div>{formatCurrency(position.salaryRange.min)}</div>
                        <div className="text-muted-foreground">
                          a {formatCurrency(position.salaryRange.max)}
                        </div>
                      </div>
                    </TableCell>
                    <TableCell>
                      <div className="flex items-center gap-2">
                        <Users className="w-4 h-4 text-chart-2" />
                        <span className="font-medium">
                          {position.employeeCount}
                        </span>
                      </div>
                    </TableCell>
                    <TableCell>{getStatusBadge(position.status)}</TableCell>
                    <TableCell className="text-right">
                      <DropdownMenu>
                        <DropdownMenuTrigger asChild>
                          <Button variant="ghost" size="sm">
                            <MoreHorizontal className="w-4 h-4" />
                          </Button>
                        </DropdownMenuTrigger>
                        <DropdownMenuContent
                          align="end"
                          className="bg-card/80 backdrop-blur-xl border-border"
                        >
                          <DropdownMenuLabel>Acciones</DropdownMenuLabel>
                          <DropdownMenuSeparator />
                          <DropdownMenuItem
                            onClick={() => handleViewDetails(position)}
                          >
                            <Eye className="w-4 h-4 mr-2" />
                            Ver Detalles
                          </DropdownMenuItem>
                          <DropdownMenuItem
                            onClick={() => handleEdit(position)}
                          >
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

      {/* Detail Dialog */}
      <Dialog open={isDetailDialogOpen} onOpenChange={setIsDetailDialogOpen}>
        <DialogContent className="sm:max-w-[800px] bg-card/80 backdrop-blur-xl border-border max-h-[90vh] overflow-y-auto">
          <DialogHeader>
            <DialogTitle className="flex items-center gap-2">
              <Briefcase className="w-5 h-5" />
              {selectedPosition?.title}
            </DialogTitle>
            <DialogDescription>
              {selectedPosition?.description}
            </DialogDescription>
          </DialogHeader>
          {selectedPosition && (
            <div className="grid gap-6 py-4">
              <div className="grid grid-cols-2 gap-4">
                <Card className="border-border bg-accent/20">
                  <CardContent className="p-4">
                    <h4 className="font-semibold mb-2">Información General</h4>
                    <div className="space-y-2 text-sm">
                      <div className="flex justify-between">
                        <span className="text-muted-foreground">
                          Departamento:
                        </span>
                        <span>{selectedPosition.department}</span>
                      </div>
                      <div className="flex justify-between">
                        <span className="text-muted-foreground">Nivel:</span>
                        {getLevelBadge(selectedPosition.level)}
                      </div>
                      <div className="flex justify-between">
                        <span className="text-muted-foreground">
                          Reporta a:
                        </span>
                        <span>
                          {selectedPosition.reportsTo || "Sin supervisor"}
                        </span>
                      </div>
                      <div className="flex justify-between">
                        <span className="text-muted-foreground">Estado:</span>
                        {getStatusBadge(selectedPosition.status)}
                      </div>
                    </div>
                  </CardContent>
                </Card>
                <Card className="border-border bg-accent/20">
                  <CardContent className="p-4">
                    <h4 className="font-semibold mb-2">Información Salarial</h4>
                    <div className="space-y-2 text-sm">
                      <div className="flex justify-between">
                        <span className="text-muted-foreground">
                          Salario Mínimo:
                        </span>
                        <span className="font-mono">
                          {formatCurrency(selectedPosition.salaryRange.min)}
                        </span>
                      </div>
                      <div className="flex justify-between">
                        <span className="text-muted-foreground">
                          Salario Máximo:
                        </span>
                        <span className="font-mono">
                          {formatCurrency(selectedPosition.salaryRange.max)}
                        </span>
                      </div>
                      <div className="flex justify-between">
                        <span className="text-muted-foreground">
                          Empleados:
                        </span>
                        <span>{selectedPosition.employeeCount}</span>
                      </div>
                    </div>
                  </CardContent>
                </Card>
              </div>
              <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                <Card className="border-border bg-accent/20">
                  <CardContent className="p-4">
                    <h4 className="font-semibold mb-3">Requisitos</h4>
                    <ul className="space-y-2 text-sm">
                      {selectedPosition.requirements.map(
                        (req: string, index: number) => (
                          <li key={index} className="flex items-start gap-2">
                            <div className="w-1.5 h-1.5 bg-primary rounded-full mt-2 flex-shrink-0" />
                            <span>{req}</span>
                          </li>
                        )
                      )}
                    </ul>
                  </CardContent>
                </Card>
                <Card className="border-border bg-accent/20">
                  <CardContent className="p-4">
                    <h4 className="font-semibold mb-3">Responsabilidades</h4>
                    <ul className="space-y-2 text-sm">
                      {selectedPosition.responsibilities.map(
                        (resp: string, index: number) => (
                          <li key={index} className="flex items-start gap-2">
                            <div className="w-1.5 h-1.5 bg-chart-2 rounded-full mt-2 flex-shrink-0" />
                            <span>{resp}</span>
                          </li>
                        )
                      )}
                    </ul>
                  </CardContent>
                </Card>
              </div>
            </div>
          )}
          <DialogFooter>
            <Button
              variant="outline"
              onClick={() => setIsDetailDialogOpen(false)}
            >
              Cerrar
            </Button>
            <Button
              className="bg-gradient-to-r from-primary to-chart-1"
              onClick={() => {
                setIsDetailDialogOpen(false);
                if (selectedPosition) {
                  handleEdit(selectedPosition);
                }
              }}
            >
              Editar Posición
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>

      {/* Edit Dialog */}
      <Dialog open={isEditDialogOpen} onOpenChange={setIsEditDialogOpen}>
        <DialogContent className="sm:max-w-[700px] bg-card/80 backdrop-blur-xl border-border max-h-[90vh] overflow-y-auto">
          <DialogHeader>
            <DialogTitle>Editar Posición Estructural</DialogTitle>
            <DialogDescription>
              Modifica la información de la posición seleccionada.
            </DialogDescription>
          </DialogHeader>
          {selectedPosition && (
            <div className="grid gap-4 py-4">
              <div className="grid grid-cols-2 gap-4">
                <div className="space-y-2">
                  <Label htmlFor="edit-title">Título del Cargo</Label>
                  <Input
                    id="edit-title"
                    defaultValue={selectedPosition.title}
                  />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="edit-department">Departamento</Label>
                  <Select defaultValue={selectedPosition.department}>
                    <SelectTrigger>
                      <SelectValue />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value="Gerencia General">
                        Gerencia General
                      </SelectItem>
                      <SelectItem value="Gerencia de Sistemas">
                        Gerencia de Sistemas
                      </SelectItem>
                      <SelectItem value="Desarrollo de Software">
                        Desarrollo de Software
                      </SelectItem>
                      <SelectItem value="Infraestructura TI">
                        Infraestructura TI
                      </SelectItem>
                      <SelectItem value="Recursos Humanos">
                        Recursos Humanos
                      </SelectItem>
                    </SelectContent>
                  </Select>
                </div>
              </div>
              <div className="space-y-2">
                <Label htmlFor="edit-description">Descripción</Label>
                <Textarea
                  id="edit-description"
                  defaultValue={selectedPosition.description}
                />
              </div>
              <div className="grid grid-cols-3 gap-4">
                <div className="space-y-2">
                  <Label htmlFor="edit-level">Nivel Jerárquico</Label>
                  <Select defaultValue={selectedPosition.level.toString()}>
                    <SelectTrigger>
                      <SelectValue />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value="1">Nivel 1 - Ejecutivo</SelectItem>
                      <SelectItem value="2">Nivel 2 - Gerencial</SelectItem>
                      <SelectItem value="3">Nivel 3 - Especialista</SelectItem>
                      <SelectItem value="4">Nivel 4 - Operativo</SelectItem>
                    </SelectContent>
                  </Select>
                </div>
                <div className="space-y-2">
                  <Label htmlFor="edit-salary-min">Salario Mínimo (PEN)</Label>
                  <Input
                    id="edit-salary-min"
                    type="number"
                    defaultValue={selectedPosition.salaryRange.min}
                  />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="edit-salary-max">Salario Máximo (PEN)</Label>
                  <Input
                    id="edit-salary-max"
                    type="number"
                    defaultValue={selectedPosition.salaryRange.max}
                  />
                </div>
              </div>
              <div className="space-y-2">
                <Label htmlFor="edit-status">Estado</Label>
                <Select defaultValue={selectedPosition.status}>
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
            <Button
              variant="outline"
              onClick={() => setIsEditDialogOpen(false)}
            >
              Cancelar
            </Button>
            <Button className="bg-gradient-to-r from-primary to-chart-1">
              Guardar Cambios
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </div>
  );
}
