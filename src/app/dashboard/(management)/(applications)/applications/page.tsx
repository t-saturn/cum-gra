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
import { Switch } from "@/components/ui/switch";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import {
  Search,
  Plus,
  Filter,
  Download,
  MoreHorizontal,
  Edit,
  Trash2,
  Eye,
  Building2,
  Users,
  Shield,
  Globe,
  Key,
  Settings,
  Activity,
  CheckCircle,
  XCircle,
  Clock,
  ExternalLink,
  Copy,
  RefreshCw,
} from "lucide-react";

// Mock data
const applications = [
  {
    id: "1",
    name: "Sistema de Inventario",
    description: "Aplicación para gestión de inventarios y almacenes",
    clientId: "inv_client_12345",
    clientSecret: "inv_secret_abcdef123456",
    redirectUris: [
      "https://inventario.empresa.com/callback",
      "https://inventario.empresa.com/auth",
    ],
    allowedOrigins: ["https://inventario.empresa.com"],
    tokenExpiration: 3600,
    refreshTokenExpiration: 86400,
    status: "active",
    type: "web",
    owner: {
      name: "Carlos López",
      email: "clopez@empresa.com",
    },
    statistics: {
      totalUsers: 45,
      activeUsers: 38,
      totalLogins: 1250,
      lastAccess: "2024-01-15 14:30:00",
    },
    modules: [
      { id: "inv_products", name: "Gestión de Productos", enabled: true },
      { id: "inv_warehouse", name: "Control de Almacén", enabled: true },
      { id: "inv_reports", name: "Reportes", enabled: true },
      { id: "inv_analytics", name: "Analíticas", enabled: false },
    ],
    scopes: ["read:inventory", "write:inventory", "read:reports"],
    createdAt: "2024-01-01",
    updatedAt: "2024-01-15",
  },
  {
    id: "2",
    name: "Portal de Recursos Humanos",
    description: "Sistema integral para gestión de recursos humanos y nómina",
    clientId: "hr_client_67890",
    clientSecret: "hr_secret_xyz789456",
    redirectUris: ["https://rrhh.empresa.com/oauth/callback"],
    allowedOrigins: [
      "https://rrhh.empresa.com",
      "https://admin.rrhh.empresa.com",
    ],
    tokenExpiration: 7200,
    refreshTokenExpiration: 172800,
    status: "active",
    type: "web",
    owner: {
      name: "María García",
      email: "mgarcia@empresa.com",
    },
    statistics: {
      totalUsers: 120,
      activeUsers: 95,
      totalLogins: 3450,
      lastAccess: "2024-01-15 16:45:00",
    },
    modules: [
      { id: "hr_employees", name: "Gestión de Empleados", enabled: true },
      { id: "hr_payroll", name: "Nómina", enabled: true },
      { id: "hr_attendance", name: "Control de Asistencia", enabled: true },
      { id: "hr_performance", name: "Evaluación de Desempeño", enabled: true },
      { id: "hr_recruitment", name: "Reclutamiento", enabled: false },
    ],
    scopes: [
      "read:employees",
      "write:employees",
      "read:payroll",
      "write:payroll",
    ],
    createdAt: "2024-01-02",
    updatedAt: "2024-01-14",
  },
  {
    id: "3",
    name: "App Móvil Ventas",
    description: "Aplicación móvil para el equipo de ventas en campo",
    clientId: "mobile_sales_11111",
    clientSecret: "mobile_secret_qwerty123",
    redirectUris: ["com.empresa.ventas://oauth/callback"],
    allowedOrigins: ["*"],
    tokenExpiration: 1800,
    refreshTokenExpiration: 604800,
    status: "development",
    type: "mobile",
    owner: {
      name: "Ana Martínez",
      email: "amartinez@empresa.com",
    },
    statistics: {
      totalUsers: 25,
      activeUsers: 18,
      totalLogins: 890,
      lastAccess: "2024-01-15 12:15:00",
    },
    modules: [
      { id: "sales_clients", name: "Gestión de Clientes", enabled: true },
      { id: "sales_orders", name: "Pedidos", enabled: true },
      { id: "sales_catalog", name: "Catálogo de Productos", enabled: true },
      { id: "sales_reports", name: "Reportes de Ventas", enabled: false },
    ],
    scopes: ["read:clients", "write:orders", "read:products"],
    createdAt: "2024-01-03",
    updatedAt: "2024-01-13",
  },
  {
    id: "4",
    name: "Sistema Financiero",
    description: "Plataforma para gestión financiera y contable",
    clientId: "finance_client_22222",
    clientSecret: "finance_secret_asdf456",
    redirectUris: ["https://finanzas.empresa.com/auth/callback"],
    allowedOrigins: ["https://finanzas.empresa.com"],
    tokenExpiration: 5400,
    refreshTokenExpiration: 259200,
    status: "suspended",
    type: "web",
    owner: {
      name: "Luis Torres",
      email: "ltorres@empresa.com",
    },
    statistics: {
      totalUsers: 15,
      activeUsers: 0,
      totalLogins: 2100,
      lastAccess: "2024-01-10 09:30:00",
    },
    modules: [
      { id: "fin_accounting", name: "Contabilidad", enabled: true },
      { id: "fin_invoicing", name: "Facturación", enabled: true },
      { id: "fin_treasury", name: "Tesorería", enabled: false },
      { id: "fin_budget", name: "Presupuestos", enabled: false },
    ],
    scopes: ["read:accounting", "write:invoices", "read:treasury"],
    createdAt: "2024-01-04",
    updatedAt: "2024-01-12",
  },
];

export default function ApplicationsManagement() {
  const [searchTerm, setSearchTerm] = useState("");
  const [isCreateDialogOpen, setIsCreateDialogOpen] = useState(false);
  const [selectedApp, setSelectedApp] = useState<(typeof applications)[0]>();
  const [, setIsEditDialogOpen] = useState(false);
  const [isDetailDialogOpen, setIsDetailDialogOpen] = useState(false);

  const filteredApps = applications.filter(
    (app) =>
      app.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
      app.description.toLowerCase().includes(searchTerm.toLowerCase()) ||
      app.owner.name.toLowerCase().includes(searchTerm.toLowerCase())
  );

  const getStatusBadge = (status: string) => {
    switch (status) {
      case "active":
        return (
          <Badge className="bg-chart-4/20 text-chart-4 border-chart-4/30">
            <CheckCircle className="w-3 h-3 mr-1" />
            Activa
          </Badge>
        );
      case "development":
        return (
          <Badge className="bg-chart-5/20 text-chart-5 border-chart-5/30">
            <Clock className="w-3 h-3 mr-1" />
            Desarrollo
          </Badge>
        );
      case "suspended":
        return (
          <Badge className="bg-destructive/20 text-destructive border-destructive/30">
            <XCircle className="w-3 h-3 mr-1" />
            Suspendida
          </Badge>
        );
      case "inactive":
        return (
          <Badge className="bg-muted text-muted-foreground border-muted-foreground/30">
            <XCircle className="w-3 h-3 mr-1" />
            Inactiva
          </Badge>
        );
      default:
        return <Badge variant="secondary">Desconocido</Badge>;
    }
  };

  const getTypeBadge = (type: string) => {
    switch (type) {
      case "web":
        return (
          <Badge className="bg-primary/20 text-primary border-primary/30">
            <Globe className="w-3 h-3 mr-1" />
            Web
          </Badge>
        );
      case "mobile":
        return (
          <Badge className="bg-chart-2/20 text-chart-2 border-chart-2/30">
            <Activity className="w-3 h-3 mr-1" />
            Móvil
          </Badge>
        );
      case "desktop":
        return (
          <Badge className="bg-chart-3/20 text-chart-3 border-chart-3/30">
            <Settings className="w-3 h-3 mr-1" />
            Escritorio
          </Badge>
        );
      default:
        return <Badge variant="secondary">Otro</Badge>;
    }
  };

  const handleEdit = (app: (typeof applications)[0]) => {
    setSelectedApp(app);
    setIsEditDialogOpen(true);
  };

  const handleViewDetails = (app: (typeof applications)[0]) => {
    setSelectedApp(app);
    setIsDetailDialogOpen(true);
  };

  const copyToClipboard = (text: string) => {
    navigator.clipboard.writeText(text);
  };

  const formatDateTime = (dateString: string) => {
    return new Date(dateString).toLocaleString("es-ES");
  };

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold text-foreground">Aplicaciones</h1>
          <p className="text-muted-foreground mt-1">
            Gestiona las aplicaciones OAuth y sus configuraciones
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
                Nueva Aplicación
              </Button>
            </DialogTrigger>
            <DialogContent className="sm:max-w-[700px] bg-card/80 backdrop-blur-xl border-border max-h-[90vh] overflow-y-auto">
              <DialogHeader>
                <DialogTitle>Crear Nueva Aplicación</DialogTitle>
                <DialogDescription>
                  Registra una nueva aplicación OAuth en el sistema para
                  permitir autenticación centralizada.
                </DialogDescription>
              </DialogHeader>
              <Tabs defaultValue="basic" className="w-full">
                <TabsList className="grid w-full grid-cols-3">
                  <TabsTrigger value="basic">Información Básica</TabsTrigger>
                  <TabsTrigger value="oauth">Configuración OAuth</TabsTrigger>
                  <TabsTrigger value="permissions">Permisos</TabsTrigger>
                </TabsList>
                <TabsContent value="basic" className="space-y-4">
                  <div className="grid grid-cols-2 gap-4">
                    <div className="space-y-2">
                      <Label htmlFor="name">Nombre de la Aplicación</Label>
                      <Input id="name" placeholder="Ej: Sistema de Ventas" />
                    </div>
                    <div className="space-y-2">
                      <Label htmlFor="type">Tipo de Aplicación</Label>
                      <Select>
                        <SelectTrigger>
                          <SelectValue placeholder="Seleccionar tipo" />
                        </SelectTrigger>
                        <SelectContent>
                          <SelectItem value="web">Aplicación Web</SelectItem>
                          <SelectItem value="mobile">
                            Aplicación Móvil
                          </SelectItem>
                          <SelectItem value="desktop">
                            Aplicación de Escritorio
                          </SelectItem>
                        </SelectContent>
                      </Select>
                    </div>
                  </div>
                  <div className="space-y-2">
                    <Label htmlFor="description">Descripción</Label>
                    <Textarea
                      id="description"
                      placeholder="Describe la funcionalidad de la aplicación..."
                    />
                  </div>
                  <div className="grid grid-cols-2 gap-4">
                    <div className="space-y-2">
                      <Label htmlFor="owner">Propietario</Label>
                      <Select>
                        <SelectTrigger>
                          <SelectValue placeholder="Seleccionar propietario" />
                        </SelectTrigger>
                        <SelectContent>
                          <SelectItem value="1">Carlos López</SelectItem>
                          <SelectItem value="2">María García</SelectItem>
                          <SelectItem value="3">Ana Martínez</SelectItem>
                        </SelectContent>
                      </Select>
                    </div>
                    <div className="space-y-2">
                      <Label htmlFor="status">Estado</Label>
                      <Select defaultValue="development">
                        <SelectTrigger>
                          <SelectValue />
                        </SelectTrigger>
                        <SelectContent>
                          <SelectItem value="development">
                            Desarrollo
                          </SelectItem>
                          <SelectItem value="active">Activa</SelectItem>
                          <SelectItem value="suspended">Suspendida</SelectItem>
                          <SelectItem value="inactive">Inactiva</SelectItem>
                        </SelectContent>
                      </Select>
                    </div>
                  </div>
                </TabsContent>
                <TabsContent value="oauth" className="space-y-4">
                  <div className="space-y-2">
                    <Label htmlFor="redirect-uris">
                      URLs de Redirección (una por línea)
                    </Label>
                    <Textarea
                      id="redirect-uris"
                      placeholder="https://miapp.com/callback&#10;https://miapp.com/auth"
                      rows={3}
                    />
                  </div>
                  <div className="space-y-2">
                    <Label htmlFor="allowed-origins">
                      Orígenes Permitidos (uno por línea)
                    </Label>
                    <Textarea
                      id="allowed-origins"
                      placeholder="https://miapp.com&#10;https://admin.miapp.com"
                      rows={3}
                    />
                  </div>
                  <div className="grid grid-cols-2 gap-4">
                    <div className="space-y-2">
                      <Label htmlFor="token-expiration">
                        Expiración Token (segundos)
                      </Label>
                      <Input
                        id="token-expiration"
                        type="number"
                        defaultValue="3600"
                      />
                    </div>
                    <div className="space-y-2">
                      <Label htmlFor="refresh-expiration">
                        Expiración Refresh Token (segundos)
                      </Label>
                      <Input
                        id="refresh-expiration"
                        type="number"
                        defaultValue="86400"
                      />
                    </div>
                  </div>
                </TabsContent>
                <TabsContent value="permissions" className="space-y-4">
                  <div className="space-y-2">
                    <Label>Scopes Disponibles</Label>
                    <div className="grid grid-cols-2 gap-4">
                      <div className="space-y-2">
                        <div className="flex items-center space-x-2">
                          <Switch id="read-users" />
                          <Label htmlFor="read-users">read:users</Label>
                        </div>
                        <div className="flex items-center space-x-2">
                          <Switch id="write-users" />
                          <Label htmlFor="write-users">write:users</Label>
                        </div>
                        <div className="flex items-center space-x-2">
                          <Switch id="read-reports" />
                          <Label htmlFor="read-reports">read:reports</Label>
                        </div>
                      </div>
                      <div className="space-y-2">
                        <div className="flex items-center space-x-2">
                          <Switch id="write-reports" />
                          <Label htmlFor="write-reports">write:reports</Label>
                        </div>
                        <div className="flex items-center space-x-2">
                          <Switch id="admin" />
                          <Label htmlFor="admin">admin</Label>
                        </div>
                        <div className="flex items-center space-x-2">
                          <Switch id="read-analytics" />
                          <Label htmlFor="read-analytics">read:analytics</Label>
                        </div>
                      </div>
                    </div>
                  </div>
                </TabsContent>
              </Tabs>
              <DialogFooter>
                <Button
                  variant="outline"
                  onClick={() => setIsCreateDialogOpen(false)}
                >
                  Cancelar
                </Button>
                <Button className="bg-gradient-to-r from-primary to-chart-1">
                  Crear Aplicación
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
                  Total Aplicaciones
                </p>
                <p className="text-2xl font-bold text-foreground">24</p>
              </div>
              <Building2 className="w-8 h-8 text-chart-2" />
            </div>
          </CardContent>
        </Card>
        <Card className="border-border bg-card/50">
          <CardContent className="p-4">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm text-muted-foreground">
                  Aplicaciones Activas
                </p>
                <p className="text-2xl font-bold text-chart-4">18</p>
              </div>
              <CheckCircle className="w-8 h-8 text-chart-4" />
            </div>
          </CardContent>
        </Card>
        <Card className="border-border bg-card/50">
          <CardContent className="p-4">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm text-muted-foreground">
                  Usuarios Totales
                </p>
                <p className="text-2xl font-bold text-primary">205</p>
              </div>
              <Users className="w-8 h-8 text-primary" />
            </div>
          </CardContent>
        </Card>
        <Card className="border-border bg-card/50">
          <CardContent className="p-4">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm text-muted-foreground">Logins Totales</p>
                <p className="text-2xl font-bold text-chart-5">7.7K</p>
              </div>
              <Activity className="w-8 h-8 text-chart-5" />
            </div>
          </CardContent>
        </Card>
      </div>

      {/* Applications Table */}
      <Card className="border-border bg-card/50">
        <CardHeader>
          <div className="flex items-center justify-between">
            <div>
              <CardTitle>Lista de Aplicaciones</CardTitle>
              <CardDescription>
                {filteredApps.length} de {applications.length} aplicaciones
              </CardDescription>
            </div>
            <div className="flex gap-2">
              <div className="relative">
                <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 text-muted-foreground w-4 h-4" />
                <Input
                  placeholder="Buscar aplicaciones..."
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
                  <TableHead>Aplicación</TableHead>
                  <TableHead>Tipo</TableHead>
                  <TableHead>Propietario</TableHead>
                  <TableHead>Usuarios</TableHead>
                  <TableHead>Último Acceso</TableHead>
                  <TableHead>Estado</TableHead>
                  <TableHead className="text-right">Acciones</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {filteredApps.map((app) => (
                  <TableRow key={app.id} className="hover:bg-accent/30">
                    <TableCell>
                      <div className="space-y-1">
                        <div className="flex items-center gap-2">
                          <Building2 className="w-4 h-4 text-primary" />
                          <p className="font-medium text-foreground">
                            {app.name}
                          </p>
                        </div>
                        <p className="text-sm text-muted-foreground line-clamp-2">
                          {app.description}
                        </p>
                        <div className="flex items-center gap-2 text-xs text-muted-foreground">
                          <Key className="w-3 h-3" />
                          <span className="font-mono">{app.clientId}</span>
                        </div>
                      </div>
                    </TableCell>
                    <TableCell>{getTypeBadge(app.type)}</TableCell>
                    <TableCell>
                      <div className="space-y-1">
                        <p className="font-medium text-sm">{app.owner.name}</p>
                        <p className="text-xs text-muted-foreground">
                          {app.owner.email}
                        </p>
                      </div>
                    </TableCell>
                    <TableCell>
                      <div className="space-y-1">
                        <div className="flex items-center gap-2">
                          <Users className="w-4 h-4 text-chart-2" />
                          <span className="font-medium">
                            {app.statistics.totalUsers}
                          </span>
                        </div>
                        <p className="text-xs text-muted-foreground">
                          {app.statistics.activeUsers} activos
                        </p>
                      </div>
                    </TableCell>
                    <TableCell className="text-sm">
                      {formatDateTime(app.statistics.lastAccess)}
                    </TableCell>
                    <TableCell>{getStatusBadge(app.status)}</TableCell>
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
                            onClick={() => handleViewDetails(app)}
                          >
                            <Eye className="w-4 h-4 mr-2" />
                            Ver Detalles
                          </DropdownMenuItem>
                          <DropdownMenuItem onClick={() => handleEdit(app)}>
                            <Edit className="w-4 h-4 mr-2" />
                            Editar
                          </DropdownMenuItem>
                          <DropdownMenuItem>
                            <Shield className="w-4 h-4 mr-2" />
                            Gestionar Módulos
                          </DropdownMenuItem>
                          <DropdownMenuItem>
                            <RefreshCw className="w-4 h-4 mr-2" />
                            Regenerar Secreto
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
        <DialogContent className="sm:max-w-[900px] bg-card/80 backdrop-blur-xl border-border max-h-[90vh] overflow-y-auto">
          <DialogHeader>
            <DialogTitle className="flex items-center gap-2">
              <Building2 className="w-5 h-5" />
              {selectedApp?.name}
            </DialogTitle>
            <DialogDescription>{selectedApp?.description}</DialogDescription>
          </DialogHeader>
          {selectedApp && (
            <Tabs defaultValue="overview" className="w-full">
              <TabsList className="grid w-full grid-cols-4">
                <TabsTrigger value="overview">Resumen</TabsTrigger>
                <TabsTrigger value="oauth">OAuth</TabsTrigger>
                <TabsTrigger value="modules">Módulos</TabsTrigger>
                <TabsTrigger value="stats">Estadísticas</TabsTrigger>
              </TabsList>
              <TabsContent value="overview" className="space-y-4">
                <div className="grid grid-cols-2 gap-4">
                  <Card className="border-border bg-accent/20">
                    <CardContent className="p-4">
                      <h4 className="font-semibold mb-2">
                        Información General
                      </h4>
                      <div className="space-y-2 text-sm">
                        <div className="flex justify-between">
                          <span className="text-muted-foreground">Tipo:</span>
                          {getTypeBadge(selectedApp.type)}
                        </div>
                        <div className="flex justify-between">
                          <span className="text-muted-foreground">Estado:</span>
                          {getStatusBadge(selectedApp.status)}
                        </div>
                        <div className="flex justify-between">
                          <span className="text-muted-foreground">
                            Propietario:
                          </span>
                          <span>{selectedApp.owner.name}</span>
                        </div>
                        <div className="flex justify-between">
                          <span className="text-muted-foreground">Creada:</span>
                          <span>
                            {new Date(selectedApp.createdAt).toLocaleDateString(
                              "es-ES"
                            )}
                          </span>
                        </div>
                      </div>
                    </CardContent>
                  </Card>
                  <Card className="border-border bg-accent/20">
                    <CardContent className="p-4">
                      <h4 className="font-semibold mb-2">
                        Estadísticas de Uso
                      </h4>
                      <div className="space-y-2 text-sm">
                        <div className="flex justify-between">
                          <span className="text-muted-foreground">
                            Usuarios Totales:
                          </span>
                          <span className="font-medium">
                            {selectedApp.statistics.totalUsers}
                          </span>
                        </div>
                        <div className="flex justify-between">
                          <span className="text-muted-foreground">
                            Usuarios Activos:
                          </span>
                          <span className="font-medium text-chart-4">
                            {selectedApp.statistics.activeUsers}
                          </span>
                        </div>
                        <div className="flex justify-between">
                          <span className="text-muted-foreground">
                            Total Logins:
                          </span>
                          <span className="font-medium">
                            {selectedApp.statistics.totalLogins.toLocaleString()}
                          </span>
                        </div>
                        <div className="flex justify-between">
                          <span className="text-muted-foreground">
                            Último Acceso:
                          </span>
                          <span className="text-xs">
                            {formatDateTime(selectedApp.statistics.lastAccess)}
                          </span>
                        </div>
                      </div>
                    </CardContent>
                  </Card>
                </div>
              </TabsContent>
              <TabsContent value="oauth" className="space-y-4">
                <div className="space-y-4">
                  <Card className="border-border bg-accent/20">
                    <CardContent className="p-4">
                      <h4 className="font-semibold mb-3">Credenciales OAuth</h4>
                      <div className="space-y-3">
                        <div>
                          <Label className="text-sm font-medium">
                            Client ID
                          </Label>
                          <div className="flex items-center gap-2 mt-1">
                            <Input
                              value={selectedApp.clientId}
                              readOnly
                              className="font-mono text-sm"
                            />
                            <Button
                              variant="outline"
                              size="sm"
                              onClick={() =>
                                copyToClipboard(selectedApp.clientId)
                              }
                            >
                              <Copy className="w-4 h-4" />
                            </Button>
                          </div>
                        </div>
                        <div>
                          <Label className="text-sm font-medium">
                            Client Secret
                          </Label>
                          <div className="flex items-center gap-2 mt-1">
                            <Input
                              value={selectedApp.clientSecret}
                              type="password"
                              readOnly
                              className="font-mono text-sm"
                            />
                            <Button
                              variant="outline"
                              size="sm"
                              onClick={() =>
                                copyToClipboard(selectedApp.clientSecret)
                              }
                            >
                              <Copy className="w-4 h-4" />
                            </Button>
                          </div>
                        </div>
                      </div>
                    </CardContent>
                  </Card>
                  <div className="grid grid-cols-2 gap-4">
                    <Card className="border-border bg-accent/20">
                      <CardContent className="p-4">
                        <h4 className="font-semibold mb-3">
                          URLs de Redirección
                        </h4>
                        <div className="space-y-2">
                          {selectedApp.redirectUris.map(
                            (uri: string, index: number) => (
                              <div
                                key={index}
                                className="flex items-center gap-2"
                              >
                                <ExternalLink className="w-4 h-4 text-muted-foreground" />
                                <span className="text-sm font-mono">{uri}</span>
                              </div>
                            )
                          )}
                        </div>
                      </CardContent>
                    </Card>
                    <Card className="border-border bg-accent/20">
                      <CardContent className="p-4">
                        <h4 className="font-semibold mb-3">
                          Orígenes Permitidos
                        </h4>
                        <div className="space-y-2">
                          {selectedApp.allowedOrigins.map(
                            (origin: string, index: number) => (
                              <div
                                key={index}
                                className="flex items-center gap-2"
                              >
                                <Globe className="w-4 h-4 text-muted-foreground" />
                                <span className="text-sm font-mono">
                                  {origin}
                                </span>
                              </div>
                            )
                          )}
                        </div>
                      </CardContent>
                    </Card>
                  </div>
                </div>
              </TabsContent>
              <TabsContent value="modules" className="space-y-4">
                <Card className="border-border bg-accent/20">
                  <CardContent className="p-4">
                    <h4 className="font-semibold mb-3">
                      Módulos de la Aplicación
                    </h4>
                    <div className="space-y-3">
                      {selectedApp.modules.map(
                        (module: (typeof applications)[0]["modules"][0]) => (
                          <div
                            key={module.id}
                            className="flex items-center justify-between p-3 bg-background/50 rounded-lg"
                          >
                            <div className="flex items-center gap-3">
                              <div
                                className={`w-2 h-2 rounded-full ${
                                  module.enabled
                                    ? "bg-chart-4"
                                    : "bg-muted-foreground"
                                }`}
                              />
                              <div>
                                <p className="font-medium text-sm">
                                  {module.name}
                                </p>
                                <p className="text-xs text-muted-foreground">
                                  ID: {module.id}
                                </p>
                              </div>
                            </div>
                            <Badge
                              variant={module.enabled ? "default" : "secondary"}
                            >
                              {module.enabled ? "Habilitado" : "Deshabilitado"}
                            </Badge>
                          </div>
                        )
                      )}
                    </div>
                  </CardContent>
                </Card>
              </TabsContent>
              <TabsContent value="stats" className="space-y-4">
                <div className="grid grid-cols-2 gap-4">
                  <Card className="border-border bg-accent/20">
                    <CardContent className="p-4">
                      <h4 className="font-semibold mb-3">
                        Configuración de Tokens
                      </h4>
                      <div className="space-y-2 text-sm">
                        <div className="flex justify-between">
                          <span className="text-muted-foreground">
                            Expiración Token:
                          </span>
                          <span>{selectedApp.tokenExpiration}s</span>
                        </div>
                        <div className="flex justify-between">
                          <span className="text-muted-foreground">
                            Expiración Refresh:
                          </span>
                          <span>{selectedApp.refreshTokenExpiration}s</span>
                        </div>
                      </div>
                    </CardContent>
                  </Card>
                  <Card className="border-border bg-accent/20">
                    <CardContent className="p-4">
                      <h4 className="font-semibold mb-3">Scopes Asignados</h4>
                      <div className="flex flex-wrap gap-2">
                        {selectedApp.scopes.map(
                          (scope: string, index: number) => (
                            <Badge
                              key={index}
                              variant="outline"
                              className="text-xs"
                            >
                              {scope}
                            </Badge>
                          )
                        )}
                      </div>
                    </CardContent>
                  </Card>
                </div>
              </TabsContent>
            </Tabs>
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
                if (selectedApp) {
                  handleEdit(selectedApp);
                }
              }}
            >
              Editar Aplicación
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </div>
  );
}
