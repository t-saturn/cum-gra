"use client"

import { useState } from "react"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Badge } from "@/components/ui/badge"
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog"
import { Label } from "@/components/ui/label"
import { Textarea } from "@/components/ui/textarea"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"
import { Switch } from "@/components/ui/switch"
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs"
import {
  Search,
  Plus,
  Filter,
  Download,
  MoreHorizontal,
  Edit,
  Trash2,
  Eye,
  Grid3X3,
  Building2,
  Shield,
  Users,
  Settings,
  CheckCircle,
  XCircle,
  AlertTriangle,
  Code,
  Database,
  FileText,
  BarChart3,
  Lock,
  Unlock,
} from "lucide-react"
import CardStatsContain from "@/components/custom/card/card-stats-contain"
import { statsModules } from "@/mocks/stats-mocks"

// Mock data
const modules = [
  {
    id: "inv_products",
    name: "Gestión de Productos",
    description: "Módulo para administrar el catálogo de productos, precios y categorías",
    application: {
      id: "1",
      name: "Sistema de Inventario",
      clientId: "inv_client_12345",
    },
    version: "2.1.4",
    status: "active",
    category: "inventory",
    permissions: [
      { id: "products.read", name: "Leer Productos", enabled: true },
      { id: "products.write", name: "Escribir Productos", enabled: true },
      { id: "products.delete", name: "Eliminar Productos", enabled: false },
      { id: "categories.manage", name: "Gestionar Categorías", enabled: true },
    ],
    endpoints: [
      { method: "GET", path: "/api/products", description: "Listar productos" },
      { method: "POST", path: "/api/products", description: "Crear producto" },
      { method: "PUT", path: "/api/products/{id}", description: "Actualizar producto" },
      { method: "DELETE", path: "/api/products/{id}", description: "Eliminar producto" },
    ],
    statistics: {
      totalUsers: 25,
      activeUsers: 18,
      totalRequests: 15420,
      errorRate: 0.02,
    },
    lastUpdated: "2024-01-15",
    createdAt: "2024-01-01",
  },
  {
    id: "inv_warehouse",
    name: "Control de Almacén",
    description: "Gestión de inventarios, movimientos de stock y ubicaciones de almacén",
    application: {
      id: "1",
      name: "Sistema de Inventario",
      clientId: "inv_client_12345",
    },
    version: "1.8.2",
    status: "active",
    category: "inventory",
    permissions: [
      { id: "warehouse.read", name: "Leer Almacén", enabled: true },
      { id: "warehouse.write", name: "Escribir Almacén", enabled: true },
      { id: "stock.manage", name: "Gestionar Stock", enabled: true },
      { id: "movements.track", name: "Rastrear Movimientos", enabled: true },
    ],
    endpoints: [
      { method: "GET", path: "/api/warehouse/stock", description: "Consultar stock" },
      { method: "POST", path: "/api/warehouse/movements", description: "Registrar movimiento" },
      { method: "GET", path: "/api/warehouse/locations", description: "Listar ubicaciones" },
    ],
    statistics: {
      totalUsers: 15,
      activeUsers: 12,
      totalRequests: 8750,
      errorRate: 0.01,
    },
    lastUpdated: "2024-01-14",
    createdAt: "2024-01-02",
  },
  {
    id: "hr_employees",
    name: "Gestión de Empleados",
    description: "Administración completa de información de empleados y expedientes",
    application: {
      id: "2",
      name: "Portal de Recursos Humanos",
      clientId: "hr_client_67890",
    },
    version: "3.2.1",
    status: "active",
    category: "hr",
    permissions: [
      { id: "employees.read", name: "Leer Empleados", enabled: true },
      { id: "employees.write", name: "Escribir Empleados", enabled: true },
      { id: "employees.delete", name: "Eliminar Empleados", enabled: false },
      { id: "personal.data", name: "Datos Personales", enabled: true },
      { id: "salary.info", name: "Información Salarial", enabled: false },
    ],
    endpoints: [
      { method: "GET", path: "/api/employees", description: "Listar empleados" },
      { method: "POST", path: "/api/employees", description: "Crear empleado" },
      { method: "PUT", path: "/api/employees/{id}", description: "Actualizar empleado" },
      { method: "GET", path: "/api/employees/{id}/profile", description: "Perfil del empleado" },
    ],
    statistics: {
      totalUsers: 45,
      activeUsers: 38,
      totalRequests: 22100,
      errorRate: 0.005,
    },
    lastUpdated: "2024-01-13",
    createdAt: "2024-01-02",
  },
  {
    id: "hr_payroll",
    name: "Nómina",
    description: "Sistema de cálculo y gestión de nóminas, beneficios y deducciones",
    application: {
      id: "2",
      name: "Portal de Recursos Humanos",
      clientId: "hr_client_67890",
    },
    version: "2.5.0",
    status: "maintenance",
    category: "hr",
    permissions: [
      { id: "payroll.read", name: "Leer Nómina", enabled: true },
      { id: "payroll.write", name: "Escribir Nómina", enabled: true },
      { id: "payroll.process", name: "Procesar Nómina", enabled: true },
      { id: "benefits.manage", name: "Gestionar Beneficios", enabled: true },
    ],
    endpoints: [
      { method: "GET", path: "/api/payroll", description: "Consultar nómina" },
      { method: "POST", path: "/api/payroll/process", description: "Procesar nómina" },
      { method: "GET", path: "/api/payroll/reports", description: "Reportes de nómina" },
    ],
    statistics: {
      totalUsers: 8,
      activeUsers: 5,
      totalRequests: 3200,
      errorRate: 0.03,
    },
    lastUpdated: "2024-01-12",
    createdAt: "2024-01-03",
  },
  {
    id: "sales_clients",
    name: "Gestión de Clientes",
    description: "CRM para gestión de clientes, contactos y oportunidades de venta",
    application: {
      id: "3",
      name: "App Móvil Ventas",
      clientId: "mobile_sales_11111",
    },
    version: "1.4.3",
    status: "active",
    category: "sales",
    permissions: [
      { id: "clients.read", name: "Leer Clientes", enabled: true },
      { id: "clients.write", name: "Escribir Clientes", enabled: true },
      { id: "contacts.manage", name: "Gestionar Contactos", enabled: true },
      { id: "opportunities.track", name: "Rastrear Oportunidades", enabled: true },
    ],
    endpoints: [
      { method: "GET", path: "/api/clients", description: "Listar clientes" },
      { method: "POST", path: "/api/clients", description: "Crear cliente" },
      { method: "GET", path: "/api/clients/{id}/opportunities", description: "Oportunidades del cliente" },
    ],
    statistics: {
      totalUsers: 20,
      activeUsers: 16,
      totalRequests: 12800,
      errorRate: 0.015,
    },
    lastUpdated: "2024-01-11",
    createdAt: "2024-01-04",
  },
  {
    id: "fin_accounting",
    name: "Contabilidad",
    description: "Sistema contable con plan de cuentas, asientos y estados financieros",
    application: {
      id: "4",
      name: "Sistema Financiero",
      clientId: "finance_client_22222",
    },
    version: "4.1.0",
    status: "suspended",
    category: "finance",
    permissions: [
      { id: "accounting.read", name: "Leer Contabilidad", enabled: true },
      { id: "accounting.write", name: "Escribir Contabilidad", enabled: false },
      { id: "entries.manage", name: "Gestionar Asientos", enabled: false },
      { id: "reports.generate", name: "Generar Reportes", enabled: true },
    ],
    endpoints: [
      { method: "GET", path: "/api/accounting/entries", description: "Consultar asientos" },
      { method: "POST", path: "/api/accounting/entries", description: "Crear asiento" },
      { method: "GET", path: "/api/accounting/reports", description: "Reportes contables" },
    ],
    statistics: {
      totalUsers: 12,
      activeUsers: 0,
      totalRequests: 5600,
      errorRate: 0.008,
    },
    lastUpdated: "2024-01-10",
    createdAt: "2024-01-05",
  },
]

export default function ModulesManagement() {
  const [searchTerm, setSearchTerm] = useState("")
  const [isCreateDialogOpen, setIsCreateDialogOpen] = useState(false)
  const [selectedModule, setSelectedModule] = useState<(typeof modules)[0]>()
  const [isEditDialogOpen, setIsEditDialogOpen] = useState(false)
  const [isDetailDialogOpen, setIsDetailDialogOpen] = useState(false)

  const filteredModules = modules.filter(
    (module) =>
      module.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
      module.description.toLowerCase().includes(searchTerm.toLowerCase()) ||
      module.application.name.toLowerCase().includes(searchTerm.toLowerCase()),
  )

  const getStatusBadge = (status: string) => {
    switch (status) {
      case "active":
        return (
          <Badge className="bg-chart-4/20 text-chart-4 border-chart-4/30">
            <CheckCircle className="w-3 h-3 mr-1" />
            Activo
          </Badge>
        )
      case "maintenance":
        return (
          <Badge className="bg-chart-5/20 text-chart-5 border-chart-5/30">
            <AlertTriangle className="w-3 h-3 mr-1" />
            Mantenimiento
          </Badge>
        )
      case "suspended":
        return (
          <Badge className="bg-destructive/20 text-destructive border-destructive/30">
            <XCircle className="w-3 h-3 mr-1" />
            Suspendido
          </Badge>
        )
      case "inactive":
        return (
          <Badge className="bg-muted text-muted-foreground border-muted-foreground/30">
            <XCircle className="w-3 h-3 mr-1" />
            Inactivo
          </Badge>
        )
      default:
        return <Badge variant="secondary">Desconocido</Badge>
    }
  }

  const getCategoryBadge = (category: string) => {
    const configs = {
      inventory: { color: "bg-primary/20 text-primary border-primary/30", icon: Database },
      hr: { color: "bg-chart-2/20 text-chart-2 border-chart-2/30", icon: Users },
      sales: { color: "bg-chart-3/20 text-chart-3 border-chart-3/30", icon: BarChart3 },
      finance: { color: "bg-chart-4/20 text-chart-4 border-chart-4/30", icon: FileText },
      system: { color: "bg-chart-5/20 text-chart-5 border-chart-5/30", icon: Settings },
    }
    const config = configs[category as keyof typeof configs] || configs.system
    const Icon = config.icon
    return (
      <Badge className={config.color}>
        <Icon className="w-3 h-3 mr-1" />
        {category.charAt(0).toUpperCase() + category.slice(1)}
      </Badge>
    )
  }

  const getMethodBadge = (method: string) => {
    const configs = {
      GET: "bg-chart-4/20 text-chart-4 border-chart-4/30",
      POST: "bg-primary/20 text-primary border-primary/30",
      PUT: "bg-chart-5/20 text-chart-5 border-chart-5/30",
      DELETE: "bg-destructive/20 text-destructive border-destructive/30",
    }
    return <Badge className={configs[method as keyof typeof configs] || "bg-muted"}>{method}</Badge>
  }

  const handleEdit = (module: (typeof modules)[0]) => {
    setSelectedModule(module)
    setIsEditDialogOpen(true)
  }

  const handleViewDetails = (module: (typeof modules)[0]) => {
    setSelectedModule(module)
    setIsDetailDialogOpen(true)
  }

  const formatErrorRate = (rate: number) => {
    return `${(rate * 100).toFixed(2)}%`
  }

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold text-foreground">Módulos</h1>
          <p className="text-muted-foreground mt-1">Gestiona los módulos y funcionalidades de las aplicaciones</p>
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
                Nuevo Módulo
              </Button>
            </DialogTrigger>
            <DialogContent className="sm:max-w-[700px] bg-card/80 backdrop-blur-xl border-border max-h-[90vh] overflow-y-auto">
              <DialogHeader>
                <DialogTitle>Crear Nuevo Módulo</DialogTitle>
                <DialogDescription>Define un nuevo módulo funcional para una aplicación existente.</DialogDescription>
              </DialogHeader>
              <Tabs defaultValue="basic" className="w-full">
                <TabsList className="grid w-full grid-cols-3">
                  <TabsTrigger value="basic">Información Básica</TabsTrigger>
                  <TabsTrigger value="permissions">Permisos</TabsTrigger>
                  <TabsTrigger value="endpoints">Endpoints</TabsTrigger>
                </TabsList>
                <TabsContent value="basic" className="space-y-4">
                  <div className="grid grid-cols-2 gap-4">
                    <div className="space-y-2">
                      <Label htmlFor="module-name">Nombre del Módulo</Label>
                      <Input id="module-name" placeholder="Ej: Gestión de Ventas" />
                    </div>
                    <div className="space-y-2">
                      <Label htmlFor="module-id">ID del Módulo</Label>
                      <Input id="module-id" placeholder="Ej: sales_management" />
                    </div>
                  </div>
                  <div className="space-y-2">
                    <Label htmlFor="module-description">Descripción</Label>
                    <Textarea id="module-description" placeholder="Describe la funcionalidad del módulo..." />
                  </div>
                  <div className="grid grid-cols-3 gap-4">
                    <div className="space-y-2">
                      <Label htmlFor="application">Aplicación</Label>
                      <Select>
                        <SelectTrigger>
                          <SelectValue placeholder="Seleccionar aplicación" />
                        </SelectTrigger>
                        <SelectContent>
                          <SelectItem value="1">Sistema de Inventario</SelectItem>
                          <SelectItem value="2">Portal de Recursos Humanos</SelectItem>
                          <SelectItem value="3">App Móvil Ventas</SelectItem>
                          <SelectItem value="4">Sistema Financiero</SelectItem>
                        </SelectContent>
                      </Select>
                    </div>
                    <div className="space-y-2">
                      <Label htmlFor="category">Categoría</Label>
                      <Select>
                        <SelectTrigger>
                          <SelectValue placeholder="Seleccionar categoría" />
                        </SelectTrigger>
                        <SelectContent>
                          <SelectItem value="inventory">Inventario</SelectItem>
                          <SelectItem value="hr">Recursos Humanos</SelectItem>
                          <SelectItem value="sales">Ventas</SelectItem>
                          <SelectItem value="finance">Finanzas</SelectItem>
                          <SelectItem value="system">Sistema</SelectItem>
                        </SelectContent>
                      </Select>
                    </div>
                    <div className="space-y-2">
                      <Label htmlFor="version">Versión</Label>
                      <Input id="version" placeholder="1.0.0" />
                    </div>
                  </div>
                </TabsContent>
                <TabsContent value="permissions" className="space-y-4">
                  <div className="space-y-4">
                    <div className="space-y-2">
                      <Label>Permisos del Módulo</Label>
                      <p className="text-sm text-muted-foreground">
                        Define los permisos específicos que este módulo puede otorgar.
                      </p>
                    </div>
                    <div className="space-y-3">
                      <div className="grid grid-cols-2 gap-4">
                        <div className="space-y-2">
                          <Label htmlFor="perm-id">ID del Permiso</Label>
                          <Input id="perm-id" placeholder="module.action" />
                        </div>
                        <div className="space-y-2">
                          <Label htmlFor="perm-name">Nombre del Permiso</Label>
                          <Input id="perm-name" placeholder="Descripción del permiso" />
                        </div>
                      </div>
                      <Button variant="outline" size="sm">
                        <Plus className="w-4 h-4 mr-2" />
                        Agregar Permiso
                      </Button>
                    </div>
                  </div>
                </TabsContent>
                <TabsContent value="endpoints" className="space-y-4">
                  <div className="space-y-4">
                    <div className="space-y-2">
                      <Label>Endpoints de la API</Label>
                      <p className="text-sm text-muted-foreground">Define los endpoints que expone este módulo.</p>
                    </div>
                    <div className="space-y-3">
                      <div className="grid grid-cols-3 gap-4">
                        <div className="space-y-2">
                          <Label htmlFor="method">Método HTTP</Label>
                          <Select>
                            <SelectTrigger>
                              <SelectValue placeholder="Método" />
                            </SelectTrigger>
                            <SelectContent>
                              <SelectItem value="GET">GET</SelectItem>
                              <SelectItem value="POST">POST</SelectItem>
                              <SelectItem value="PUT">PUT</SelectItem>
                              <SelectItem value="DELETE">DELETE</SelectItem>
                            </SelectContent>
                          </Select>
                        </div>
                        <div className="space-y-2">
                          <Label htmlFor="path">Ruta</Label>
                          <Input id="path" placeholder="/api/resource" />
                        </div>
                        <div className="space-y-2">
                          <Label htmlFor="endpoint-desc">Descripción</Label>
                          <Input id="endpoint-desc" placeholder="Descripción del endpoint" />
                        </div>
                      </div>
                      <Button variant="outline" size="sm">
                        <Plus className="w-4 h-4 mr-2" />
                        Agregar Endpoint
                      </Button>
                    </div>
                  </div>
                </TabsContent>
              </Tabs>
              <DialogFooter>
                <Button variant="outline" onClick={() => setIsCreateDialogOpen(false)}>
                  Cancelar
                </Button>
                <Button className="bg-gradient-to-r from-primary to-chart-1">Crear Módulo</Button>
              </DialogFooter>
            </DialogContent>
          </Dialog>
        </div>
      </div>

      {/* Stats Cards */}
      <CardStatsContain stats={statsModules}/>

      {/* Modules Table */}
      <Card className="border-border bg-card/50">
        <CardHeader>
          <div className="flex items-center justify-between">
            <div>
              <CardTitle>Lista de Módulos</CardTitle>
              <CardDescription>
                {filteredModules.length} de {modules.length} módulos
              </CardDescription>
            </div>
            <div className="flex gap-2">
              <div className="relative">
                <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 text-muted-foreground w-4 h-4" />
                <Input
                  placeholder="Buscar módulos..."
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
                  <TableHead>Módulo</TableHead>
                  <TableHead>Aplicación</TableHead>
                  <TableHead>Categoría</TableHead>
                  <TableHead>Versión</TableHead>
                  <TableHead>Usuarios</TableHead>
                  <TableHead>Requests</TableHead>
                  <TableHead>Estado</TableHead>
                  <TableHead className="text-right">Acciones</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {filteredModules.map((module) => (
                  <TableRow key={module.id} className="hover:bg-accent/30">
                    <TableCell>
                      <div className="space-y-1">
                        <div className="flex items-center gap-2">
                          <Grid3X3 className="w-4 h-4 text-primary" />
                          <p className="font-medium text-foreground">{module.name}</p>
                        </div>
                        <p className="text-sm text-muted-foreground line-clamp-2">{module.description}</p>
                        <div className="flex items-center gap-2 text-xs text-muted-foreground">
                          <Code className="w-3 h-3" />
                          <span className="font-mono">{module.id}</span>
                        </div>
                      </div>
                    </TableCell>
                    <TableCell>
                      <div className="space-y-1">
                        <div className="flex items-center gap-2">
                          <Building2 className="w-4 h-4 text-chart-2" />
                          <p className="font-medium text-sm">{module.application.name}</p>
                        </div>
                        <p className="text-xs text-muted-foreground font-mono">{module.application.clientId}</p>
                      </div>
                    </TableCell>
                    <TableCell>{getCategoryBadge(module.category)}</TableCell>
                    <TableCell>
                      <Badge variant="outline" className="font-mono text-xs">
                        v{module.version}
                      </Badge>
                    </TableCell>
                    <TableCell>
                      <div className="space-y-1">
                        <div className="flex items-center gap-2">
                          <Users className="w-4 h-4 text-chart-2" />
                          <span className="font-medium">{module.statistics.totalUsers}</span>
                        </div>
                        <p className="text-xs text-muted-foreground">{module.statistics.activeUsers} activos</p>
                      </div>
                    </TableCell>
                    <TableCell>
                      <div className="space-y-1">
                        <p className="font-medium text-sm">{module.statistics.totalRequests.toLocaleString()}</p>
                        <p className="text-xs text-muted-foreground">
                          Error: {formatErrorRate(module.statistics.errorRate)}
                        </p>
                      </div>
                    </TableCell>
                    <TableCell>{getStatusBadge(module.status)}</TableCell>
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
                          <DropdownMenuItem onClick={() => handleViewDetails(module)}>
                            <Eye className="w-4 h-4 mr-2" />
                            Ver Detalles
                          </DropdownMenuItem>
                          <DropdownMenuItem onClick={() => handleEdit(module)}>
                            <Edit className="w-4 h-4 mr-2" />
                            Editar
                          </DropdownMenuItem>
                          <DropdownMenuItem>
                            <Shield className="w-4 h-4 mr-2" />
                            Gestionar Permisos
                          </DropdownMenuItem>
                          <DropdownMenuItem>
                            <Settings className="w-4 h-4 mr-2" />
                            Configuración
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
              <Grid3X3 className="w-5 h-5" />
              {selectedModule?.name}
            </DialogTitle>
            <DialogDescription>{selectedModule?.description}</DialogDescription>
          </DialogHeader>
          {selectedModule && (
            <Tabs defaultValue="overview" className="w-full">
              <TabsList className="grid w-full grid-cols-4">
                <TabsTrigger value="overview">Resumen</TabsTrigger>
                <TabsTrigger value="permissions">Permisos</TabsTrigger>
                <TabsTrigger value="endpoints">Endpoints</TabsTrigger>
                <TabsTrigger value="stats">Estadísticas</TabsTrigger>
              </TabsList>
              <TabsContent value="overview" className="space-y-4">
                <div className="grid grid-cols-2 gap-4">
                  <Card className="border-border bg-accent/20">
                    <CardContent className="p-4">
                      <h4 className="font-semibold mb-2">Información General</h4>
                      <div className="space-y-2 text-sm">
                        <div className="flex justify-between">
                          <span className="text-muted-foreground">ID del Módulo:</span>
                          <span className="font-mono">{selectedModule.id}</span>
                        </div>
                        <div className="flex justify-between">
                          <span className="text-muted-foreground">Versión:</span>
                          <Badge variant="outline" className="text-xs">
                            v{selectedModule.version}
                          </Badge>
                        </div>
                        <div className="flex justify-between">
                          <span className="text-muted-foreground">Categoría:</span>
                          {getCategoryBadge(selectedModule.category)}
                        </div>
                        <div className="flex justify-between">
                          <span className="text-muted-foreground">Estado:</span>
                          {getStatusBadge(selectedModule.status)}
                        </div>
                        <div className="flex justify-between">
                          <span className="text-muted-foreground">Última Actualización:</span>
                          <span>{new Date(selectedModule.lastUpdated).toLocaleDateString("es-ES")}</span>
                        </div>
                      </div>
                    </CardContent>
                  </Card>
                  <Card className="border-border bg-accent/20">
                    <CardContent className="p-4">
                      <h4 className="font-semibold mb-2">Aplicación Asociada</h4>
                      <div className="space-y-2 text-sm">
                        <div className="flex items-center gap-2">
                          <Building2 className="w-4 h-4 text-chart-2" />
                          <span className="font-medium">{selectedModule.application.name}</span>
                        </div>
                        <div className="flex items-center gap-2">
                          <Code className="w-4 h-4 text-muted-foreground" />
                          <span className="font-mono text-xs">{selectedModule.application.clientId}</span>
                        </div>
                      </div>
                    </CardContent>
                  </Card>
                </div>
              </TabsContent>
              <TabsContent value="permissions" className="space-y-4">
                <Card className="border-border bg-accent/20">
                  <CardContent className="p-4">
                    <h4 className="font-semibold mb-3">Permisos del Módulo</h4>
                    <div className="space-y-3">
                      {selectedModule.permissions.map((permission: (typeof modules)[0]["permissions"][0]) => (
                        <div
                          key={permission.id}
                          className="flex items-center justify-between p-3 bg-background/50 rounded-lg"
                        >
                          <div className="flex items-center gap-3">
                            {permission.enabled ? (
                              <Lock className="w-4 h-4 text-chart-4" />
                            ) : (
                              <Unlock className="w-4 h-4 text-muted-foreground" />
                            )}
                            <div>
                              <p className="font-medium text-sm">{permission.name}</p>
                              <p className="text-xs text-muted-foreground font-mono">{permission.id}</p>
                            </div>
                          </div>
                          <Badge variant={permission.enabled ? "default" : "secondary"}>
                            {permission.enabled ? "Habilitado" : "Deshabilitado"}
                          </Badge>
                        </div>
                      ))}
                    </div>
                  </CardContent>
                </Card>
              </TabsContent>
              <TabsContent value="endpoints" className="space-y-4">
                <Card className="border-border bg-accent/20">
                  <CardContent className="p-4">
                    <h4 className="font-semibold mb-3">Endpoints de la API</h4>
                    <div className="space-y-3">
                      {selectedModule.endpoints.map((endpoint: (typeof modules)[0]["endpoints"][0], index: number) => (
                        <div key={index} className="flex items-center justify-between p-3 bg-background/50 rounded-lg">
                          <div className="flex items-center gap-3">
                            {getMethodBadge(endpoint.method)}
                            <div>
                              <p className="font-mono text-sm">{endpoint.path}</p>
                              <p className="text-xs text-muted-foreground">{endpoint.description}</p>
                            </div>
                          </div>
                        </div>
                      ))}
                    </div>
                  </CardContent>
                </Card>
              </TabsContent>
              <TabsContent value="stats" className="space-y-4">
                <div className="grid grid-cols-2 gap-4">
                  <Card className="border-border bg-accent/20">
                    <CardContent className="p-4">
                      <h4 className="font-semibold mb-3">Estadísticas de Uso</h4>
                      <div className="space-y-2 text-sm">
                        <div className="flex justify-between">
                          <span className="text-muted-foreground">Usuarios Totales:</span>
                          <span className="font-medium">{selectedModule.statistics.totalUsers}</span>
                        </div>
                        <div className="flex justify-between">
                          <span className="text-muted-foreground">Usuarios Activos:</span>
                          <span className="font-medium text-chart-4">{selectedModule.statistics.activeUsers}</span>
                        </div>
                        <div className="flex justify-between">
                          <span className="text-muted-foreground">Total Requests:</span>
                          <span className="font-medium">
                            {selectedModule.statistics.totalRequests.toLocaleString()}
                          </span>
                        </div>
                        <div className="flex justify-between">
                          <span className="text-muted-foreground">Tasa de Error:</span>
                          <span className="font-medium">{formatErrorRate(selectedModule.statistics.errorRate)}</span>
                        </div>
                      </div>
                    </CardContent>
                  </Card>
                  <Card className="border-border bg-accent/20">
                    <CardContent className="p-4">
                      <h4 className="font-semibold mb-3">Información Técnica</h4>
                      <div className="space-y-2 text-sm">
                        <div className="flex justify-between">
                          <span className="text-muted-foreground">Versión:</span>
                          <span className="font-mono">v{selectedModule.version}</span>
                        </div>
                        <div className="flex justify-between">
                          <span className="text-muted-foreground">Endpoints:</span>
                          <span className="font-medium">{selectedModule.endpoints.length}</span>
                        </div>
                        <div className="flex justify-between">
                          <span className="text-muted-foreground">Permisos:</span>
                          <span className="font-medium">{selectedModule.permissions.length}</span>
                        </div>
                        <div className="flex justify-between">
                          <span className="text-muted-foreground">Creado:</span>
                          <span>{new Date(selectedModule.createdAt).toLocaleDateString("es-ES")}</span>
                        </div>
                      </div>
                    </CardContent>
                  </Card>
                </div>
              </TabsContent>
            </Tabs>
          )}
          <DialogFooter>
            <Button variant="outline" onClick={() => setIsDetailDialogOpen(false)}>
              Cerrar
            </Button>
            <Button
              className="bg-gradient-to-r from-primary to-chart-1"
              onClick={() => {
                setIsDetailDialogOpen(false)
                if (selectedModule) {
                  handleEdit(selectedModule)
                }
              }}
            >
              Editar Módulo
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>

      {/* Edit Dialog */}
      <Dialog open={isEditDialogOpen} onOpenChange={setIsEditDialogOpen}>
        <DialogContent className="sm:max-w-[700px] bg-card/80 backdrop-blur-xl border-border max-h-[90vh] overflow-y-auto">
          <DialogHeader>
            <DialogTitle>Editar Módulo</DialogTitle>
            <DialogDescription>Modifica la información del módulo seleccionado.</DialogDescription>
          </DialogHeader>
          {selectedModule && (
            <Tabs defaultValue="basic" className="w-full">
              <TabsList className="grid w-full grid-cols-2">
                <TabsTrigger value="basic">Información Básica</TabsTrigger>
                <TabsTrigger value="permissions">Permisos</TabsTrigger>
              </TabsList>
              <TabsContent value="basic" className="space-y-4">
                <div className="grid grid-cols-2 gap-4">
                  <div className="space-y-2">
                    <Label htmlFor="edit-name">Nombre del Módulo</Label>
                    <Input id="edit-name" defaultValue={selectedModule.name} />
                  </div>
                  <div className="space-y-2">
                    <Label htmlFor="edit-version">Versión</Label>
                    <Input id="edit-version" defaultValue={selectedModule.version} />
                  </div>
                </div>
                <div className="space-y-2">
                  <Label htmlFor="edit-description">Descripción</Label>
                  <Textarea id="edit-description" defaultValue={selectedModule.description} />
                </div>
                <div className="grid grid-cols-2 gap-4">
                  <div className="space-y-2">
                    <Label htmlFor="edit-category">Categoría</Label>
                    <Select defaultValue={selectedModule.category}>
                      <SelectTrigger>
                        <SelectValue />
                      </SelectTrigger>
                      <SelectContent>
                        <SelectItem value="inventory">Inventario</SelectItem>
                        <SelectItem value="hr">Recursos Humanos</SelectItem>
                        <SelectItem value="sales">Ventas</SelectItem>
                        <SelectItem value="finance">Finanzas</SelectItem>
                        <SelectItem value="system">Sistema</SelectItem>
                      </SelectContent>
                    </Select>
                  </div>
                  <div className="space-y-2">
                    <Label htmlFor="edit-status">Estado</Label>
                    <Select defaultValue={selectedModule.status}>
                      <SelectTrigger>
                        <SelectValue />
                      </SelectTrigger>
                      <SelectContent>
                        <SelectItem value="active">Activo</SelectItem>
                        <SelectItem value="maintenance">Mantenimiento</SelectItem>
                        <SelectItem value="suspended">Suspendido</SelectItem>
                        <SelectItem value="inactive">Inactivo</SelectItem>
                      </SelectContent>
                    </Select>
                  </div>
                </div>
              </TabsContent>
              <TabsContent value="permissions" className="space-y-4">
                <div className="space-y-4">
                  <div className="space-y-2">
                    <Label>Gestión de Permisos</Label>
                    <p className="text-sm text-muted-foreground">
                      Habilita o deshabilita los permisos específicos de este módulo.
                    </p>
                  </div>
                  <div className="space-y-3">
                    {selectedModule.permissions.map((permission: (typeof modules)[0]["permissions"][0]) => (
                      <div
                        key={permission.id}
                        className="flex items-center justify-between p-3 bg-accent/20 rounded-lg"
                      >
                        <div>
                          <p className="font-medium text-sm">{permission.name}</p>
                          <p className="text-xs text-muted-foreground font-mono">{permission.id}</p>
                        </div>
                        <Switch defaultChecked={permission.enabled} />
                      </div>
                    ))}
                  </div>
                </div>
              </TabsContent>
            </Tabs>
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
  )
}
