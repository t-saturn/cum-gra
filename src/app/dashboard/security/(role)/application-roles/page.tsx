"use client";

import { useState } from "react";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Textarea } from "@/components/ui/textarea";
import { Checkbox } from "@/components/ui/checkbox";
import {
  Collapsible,
  CollapsibleContent,
  CollapsibleTrigger,
} from "@/components/ui/collapsible";
import {
  Search,
  Filter,
  Plus,
  Edit,
  Trash2,
  Users,
  Shield,
  Settings,
  ChevronDown,
  ChevronRight,
  Eye,
  Copy,
} from "lucide-react";

// Datos simulados de roles de aplicación
const applicationRoles = [
  {
    id: "role_001",
    name: "CRM Admin",
    description: "Administrador completo del sistema CRM",
    applicationId: "app_001",
    applicationName: "CRM System",
    level: "admin",
    permissions: [
      "crm.users.read",
      "crm.users.write",
      "crm.users.delete",
      "crm.leads.read",
      "crm.leads.write",
      "crm.leads.delete",
      "crm.reports.read",
      "crm.reports.export",
      "crm.settings.read",
      "crm.settings.write",
    ],
    userCount: 5,
    isActive: true,
    createdAt: "2024-01-10T10:00:00Z",
    updatedAt: "2024-01-15T14:30:00Z",
    createdBy: "Ana García",
  },
  {
    id: "role_002",
    name: "CRM Sales",
    description: "Vendedor con acceso a leads y oportunidades",
    applicationId: "app_001",
    applicationName: "CRM System",
    level: "user",
    permissions: [
      "crm.leads.read",
      "crm.leads.write",
      "crm.opportunities.read",
      "crm.opportunities.write",
      "crm.contacts.read",
      "crm.contacts.write",
      "crm.reports.read",
    ],
    userCount: 23,
    isActive: true,
    createdAt: "2024-01-10T10:15:00Z",
    updatedAt: "2024-01-12T09:20:00Z",
    createdBy: "Carlos López",
  },
  {
    id: "role_003",
    name: "ERP Finance",
    description: "Acceso completo al módulo financiero",
    applicationId: "app_002",
    applicationName: "ERP System",
    level: "manager",
    permissions: [
      "erp.finance.read",
      "erp.finance.write",
      "erp.accounting.read",
      "erp.accounting.write",
      "erp.budgets.read",
      "erp.budgets.write",
      "erp.reports.financial.read",
      "erp.reports.financial.export",
    ],
    userCount: 8,
    isActive: true,
    createdAt: "2024-01-08T15:30:00Z",
    updatedAt: "2024-01-14T11:45:00Z",
    createdBy: "María Rodríguez",
  },
  {
    id: "role_004",
    name: "ERP Inventory",
    description: "Gestión de inventario y almacenes",
    applicationId: "app_002",
    applicationName: "ERP System",
    level: "user",
    permissions: [
      "erp.inventory.read",
      "erp.inventory.write",
      "erp.warehouse.read",
      "erp.warehouse.write",
      "erp.products.read",
      "erp.products.write",
      "erp.suppliers.read",
      "erp.reports.inventory.read",
    ],
    userCount: 15,
    isActive: true,
    createdAt: "2024-01-09T08:45:00Z",
    updatedAt: "2024-01-13T16:20:00Z",
    createdBy: "David Martín",
  },
  {
    id: "role_005",
    name: "Analytics Viewer",
    description: "Solo lectura de dashboards y reportes",
    applicationId: "app_003",
    applicationName: "Analytics Platform",
    level: "viewer",
    permissions: [
      "analytics.dashboards.read",
      "analytics.reports.read",
      "analytics.charts.read",
      "analytics.data.read",
    ],
    userCount: 45,
    isActive: true,
    createdAt: "2024-01-11T12:00:00Z",
    updatedAt: "2024-01-11T12:00:00Z",
    createdBy: "Laura Fernández",
  },
  {
    id: "role_006",
    name: "Analytics Admin",
    description: "Administrador completo de la plataforma de analytics",
    applicationId: "app_003",
    applicationName: "Analytics Platform",
    level: "admin",
    permissions: [
      "analytics.dashboards.read",
      "analytics.dashboards.write",
      "analytics.dashboards.delete",
      "analytics.reports.read",
      "analytics.reports.write",
      "analytics.reports.delete",
      "analytics.data.read",
      "analytics.data.write",
      "analytics.users.read",
      "analytics.users.write",
      "analytics.settings.read",
      "analytics.settings.write",
    ],
    userCount: 3,
    isActive: true,
    createdAt: "2024-01-11T12:15:00Z",
    updatedAt: "2024-01-15T10:30:00Z",
    createdBy: "Ana García",
  },
];

const roleStats = {
  total: 156,
  active: 142,
  inactive: 14,
  byLevel: {
    admin: 12,
    manager: 28,
    user: 89,
    viewer: 27,
  },
  byApplication: {
    "CRM System": 45,
    "ERP System": 67,
    "Analytics Platform": 34,
    "HR System": 10,
  },
};

interface Permission {
  id: string;
  name: string;
  category: string;
}

const availablePermissions: Record<string, Permission[]> = {
  "CRM System": [
    { id: "crm.users.read", name: "Ver usuarios", category: "Usuarios" },
    { id: "crm.users.write", name: "Editar usuarios", category: "Usuarios" },
    { id: "crm.users.delete", name: "Eliminar usuarios", category: "Usuarios" },
    { id: "crm.leads.read", name: "Ver leads", category: "Ventas" },
    { id: "crm.leads.write", name: "Editar leads", category: "Ventas" },
    { id: "crm.leads.delete", name: "Eliminar leads", category: "Ventas" },
    {
      id: "crm.opportunities.read",
      name: "Ver oportunidades",
      category: "Ventas",
    },
    {
      id: "crm.opportunities.write",
      name: "Editar oportunidades",
      category: "Ventas",
    },
    { id: "crm.contacts.read", name: "Ver contactos", category: "Contactos" },
    {
      id: "crm.contacts.write",
      name: "Editar contactos",
      category: "Contactos",
    },
    { id: "crm.reports.read", name: "Ver reportes", category: "Reportes" },
    {
      id: "crm.reports.export",
      name: "Exportar reportes",
      category: "Reportes",
    },
    {
      id: "crm.settings.read",
      name: "Ver configuración",
      category: "Configuración",
    },
    {
      id: "crm.settings.write",
      name: "Editar configuración",
      category: "Configuración",
    },
  ],
  "ERP System": [
    { id: "erp.finance.read", name: "Ver finanzas", category: "Finanzas" },
    { id: "erp.finance.write", name: "Editar finanzas", category: "Finanzas" },
    {
      id: "erp.accounting.read",
      name: "Ver contabilidad",
      category: "Contabilidad",
    },
    {
      id: "erp.accounting.write",
      name: "Editar contabilidad",
      category: "Contabilidad",
    },
    {
      id: "erp.inventory.read",
      name: "Ver inventario",
      category: "Inventario",
    },
    {
      id: "erp.inventory.write",
      name: "Editar inventario",
      category: "Inventario",
    },
    { id: "erp.warehouse.read", name: "Ver almacenes", category: "Almacenes" },
    {
      id: "erp.warehouse.write",
      name: "Editar almacenes",
      category: "Almacenes",
    },
    { id: "erp.products.read", name: "Ver productos", category: "Productos" },
    {
      id: "erp.products.write",
      name: "Editar productos",
      category: "Productos",
    },
    {
      id: "erp.suppliers.read",
      name: "Ver proveedores",
      category: "Proveedores",
    },
    {
      id: "erp.reports.financial.read",
      name: "Ver reportes financieros",
      category: "Reportes",
    },
    {
      id: "erp.reports.financial.export",
      name: "Exportar reportes financieros",
      category: "Reportes",
    },
    {
      id: "erp.reports.inventory.read",
      name: "Ver reportes de inventario",
      category: "Reportes",
    },
  ],
  "Analytics Platform": [
    {
      id: "analytics.dashboards.read",
      name: "Ver dashboards",
      category: "Dashboards",
    },
    {
      id: "analytics.dashboards.write",
      name: "Editar dashboards",
      category: "Dashboards",
    },
    {
      id: "analytics.dashboards.delete",
      name: "Eliminar dashboards",
      category: "Dashboards",
    },
    {
      id: "analytics.reports.read",
      name: "Ver reportes",
      category: "Reportes",
    },
    {
      id: "analytics.reports.write",
      name: "Editar reportes",
      category: "Reportes",
    },
    {
      id: "analytics.reports.delete",
      name: "Eliminar reportes",
      category: "Reportes",
    },
    {
      id: "analytics.charts.read",
      name: "Ver gráficos",
      category: "Visualización",
    },
    { id: "analytics.data.read", name: "Ver datos", category: "Datos" },
    { id: "analytics.data.write", name: "Editar datos", category: "Datos" },
    { id: "analytics.users.read", name: "Ver usuarios", category: "Usuarios" },
    {
      id: "analytics.users.write",
      name: "Editar usuarios",
      category: "Usuarios",
    },
    {
      id: "analytics.settings.read",
      name: "Ver configuración",
      category: "Configuración",
    },
    {
      id: "analytics.settings.write",
      name: "Editar configuración",
      category: "Configuración",
    },
  ],
};

export default function ApplicationRolesManagement() {
  const [searchTerm, setSearchTerm] = useState("");
  const [applicationFilter, setApplicationFilter] = useState("all");
  const [levelFilter, setLevelFilter] = useState("all");
  const [statusFilter, setStatusFilter] = useState("all");
  const [isCreateDialogOpen, setIsCreateDialogOpen] = useState(false);
  const [expandedPermissions, setExpandedPermissions] = useState<string[]>([]);

  const [newRole, setNewRole] = useState({
    name: "",
    description: "",
    applicationId: "",
    level: "user",
    permissions: [] as string[],
  });

  const getLevelBadge = (level: string) => {
    const colors = {
      admin: "bg-red-100 text-red-800",
      manager: "bg-orange-100 text-orange-800",
      user: "bg-blue-100 text-blue-800",
      viewer: "bg-gray-100 text-gray-800",
    };
    return colors[level as keyof typeof colors] || "bg-gray-100 text-gray-800";
  };

  const filteredRoles = applicationRoles.filter((role) => {
    const matchesSearch =
      role.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
      role.description.toLowerCase().includes(searchTerm.toLowerCase()) ||
      role.applicationName.toLowerCase().includes(searchTerm.toLowerCase());
    const matchesApplication =
      applicationFilter === "all" || role.applicationName === applicationFilter;
    const matchesLevel = levelFilter === "all" || role.level === levelFilter;
    const matchesStatus =
      statusFilter === "all" ||
      (statusFilter === "active" ? role.isActive : !role.isActive);

    return matchesSearch && matchesApplication && matchesLevel && matchesStatus;
  });

  const togglePermissionExpansion = (roleId: string) => {
    setExpandedPermissions((prev) =>
      prev.includes(roleId)
        ? prev.filter((id) => id !== roleId)
        : [...prev, roleId]
    );
  };

  const handleCreateRole = () => {
    // Aquí iría la lógica para crear el rol
    console.log("Creating role:", newRole);
    setIsCreateDialogOpen(false);
    setNewRole({
      name: "",
      description: "",
      applicationId: "",
      level: "user",
      permissions: [],
    });
  };

  const groupPermissionsByCategory = (
    permissions: string[],
    applicationName: string
  ): Record<string, Permission[]> => {
    const appPermissions: Permission[] =
      availablePermissions[applicationName] || [];
    const grouped: Record<string, Permission[]> = {};

    permissions.forEach((permId) => {
      const perm = appPermissions.find((p) => p.id === permId);
      if (perm) {
        if (!grouped[perm.category]) {
          grouped[perm.category] = [];
        }
        grouped[perm.category].push(perm);
      }
    });

    return grouped;
  };

  return (
    <div className="space-y-6">
      {/* Estadísticas */}
      <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Total Roles</CardTitle>
            <Shield className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">{roleStats.total}</div>
            <p className="text-xs text-muted-foreground">+8 este mes</p>
          </CardContent>
        </Card>
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Roles Activos</CardTitle>
            <div className="h-2 w-2 bg-green-500 rounded-full animate-pulse" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold text-green-600">
              {roleStats.active}
            </div>
            <p className="text-xs text-muted-foreground">
              {Math.round((roleStats.active / roleStats.total) * 100)}% del
              total
            </p>
          </CardContent>
        </Card>
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Roles Admin</CardTitle>
            <Settings className="h-4 w-4 text-red-500" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold text-red-600">
              {roleStats.byLevel.admin}
            </div>
            <p className="text-xs text-muted-foreground">Acceso privilegiado</p>
          </CardContent>
        </Card>
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              Usuarios Asignados
            </CardTitle>
            <Users className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">
              {applicationRoles.reduce((sum, role) => sum + role.userCount, 0)}
            </div>
            <p className="text-xs text-muted-foreground">
              Total de asignaciones
            </p>
          </CardContent>
        </Card>
      </div>

      <Tabs defaultValue="roles" className="space-y-4">
        <TabsList>
          <TabsTrigger value="roles">Roles de Aplicación</TabsTrigger>
          <TabsTrigger value="permissions">Matriz de Permisos</TabsTrigger>
          <TabsTrigger value="analytics">Análisis</TabsTrigger>
        </TabsList>

        <TabsContent value="roles" className="space-y-4">
          <Card>
            <CardHeader>
              <CardTitle>Gestión de Roles por Aplicación</CardTitle>
              <CardDescription>
                Administra los roles específicos para cada aplicación del
                sistema
              </CardDescription>
            </CardHeader>
            <CardContent>
              {/* Filtros y acciones */}
              <div className="flex flex-col gap-4 md:flex-row md:items-center md:justify-between mb-6">
                <div className="flex flex-1 items-center space-x-2">
                  <Search className="h-4 w-4 text-muted-foreground" />
                  <Input
                    placeholder="Buscar roles..."
                    value={searchTerm}
                    onChange={(e) => setSearchTerm(e.target.value)}
                    className="max-w-sm"
                  />
                </div>
                <div className="flex items-center space-x-2">
                  <Filter className="h-4 w-4 text-muted-foreground" />
                  <Select
                    value={applicationFilter}
                    onValueChange={setApplicationFilter}
                  >
                    <SelectTrigger className="w-[180px]">
                      <SelectValue placeholder="Aplicación" />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value="all">
                        Todas las aplicaciones
                      </SelectItem>
                      <SelectItem value="CRM System">CRM System</SelectItem>
                      <SelectItem value="ERP System">ERP System</SelectItem>
                      <SelectItem value="Analytics Platform">
                        Analytics Platform
                      </SelectItem>
                    </SelectContent>
                  </Select>
                  <Select value={levelFilter} onValueChange={setLevelFilter}>
                    <SelectTrigger className="w-[140px]">
                      <SelectValue placeholder="Nivel" />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value="all">Todos</SelectItem>
                      <SelectItem value="admin">Admin</SelectItem>
                      <SelectItem value="manager">Manager</SelectItem>
                      <SelectItem value="user">User</SelectItem>
                      <SelectItem value="viewer">Viewer</SelectItem>
                    </SelectContent>
                  </Select>
                  <Select value={statusFilter} onValueChange={setStatusFilter}>
                    <SelectTrigger className="w-[140px]">
                      <SelectValue placeholder="Estado" />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value="all">Todos</SelectItem>
                      <SelectItem value="active">Activos</SelectItem>
                      <SelectItem value="inactive">Inactivos</SelectItem>
                    </SelectContent>
                  </Select>
                  <Dialog
                    open={isCreateDialogOpen}
                    onOpenChange={setIsCreateDialogOpen}
                  >
                    <DialogTrigger asChild>
                      <Button>
                        <Plus className="h-4 w-4 mr-2" />
                        Nuevo Rol
                      </Button>
                    </DialogTrigger>
                    <DialogContent className="max-w-2xl">
                      <DialogHeader>
                        <DialogTitle>Crear Nuevo Rol</DialogTitle>
                        <DialogDescription>
                          Define un nuevo rol con permisos específicos para una
                          aplicación
                        </DialogDescription>
                      </DialogHeader>
                      <div className="space-y-4">
                        <div className="grid grid-cols-2 gap-4">
                          <div>
                            <Label htmlFor="name">Nombre del Rol</Label>
                            <Input
                              id="name"
                              value={newRole.name}
                              onChange={(e) =>
                                setNewRole({ ...newRole, name: e.target.value })
                              }
                              placeholder="Ej: CRM Manager"
                            />
                          </div>
                          <div>
                            <Label htmlFor="level">Nivel</Label>
                            <Select
                              value={newRole.level}
                              onValueChange={(value) =>
                                setNewRole({ ...newRole, level: value })
                              }
                            >
                              <SelectTrigger>
                                <SelectValue />
                              </SelectTrigger>
                              <SelectContent>
                                <SelectItem value="viewer">Viewer</SelectItem>
                                <SelectItem value="user">User</SelectItem>
                                <SelectItem value="manager">Manager</SelectItem>
                                <SelectItem value="admin">Admin</SelectItem>
                              </SelectContent>
                            </Select>
                          </div>
                        </div>
                        <div>
                          <Label htmlFor="application">Aplicación</Label>
                          <Select
                            value={newRole.applicationId}
                            onValueChange={(value) =>
                              setNewRole({ ...newRole, applicationId: value })
                            }
                          >
                            <SelectTrigger>
                              <SelectValue placeholder="Selecciona una aplicación" />
                            </SelectTrigger>
                            <SelectContent>
                              <SelectItem value="app_001">
                                CRM System
                              </SelectItem>
                              <SelectItem value="app_002">
                                ERP System
                              </SelectItem>
                              <SelectItem value="app_003">
                                Analytics Platform
                              </SelectItem>
                            </SelectContent>
                          </Select>
                        </div>
                        <div>
                          <Label htmlFor="description">Descripción</Label>
                          <Textarea
                            id="description"
                            value={newRole.description}
                            onChange={(e) =>
                              setNewRole({
                                ...newRole,
                                description: e.target.value,
                              })
                            }
                            placeholder="Describe las responsabilidades de este rol..."
                          />
                        </div>
                        {newRole.applicationId && (
                          <div>
                            <Label>Permisos</Label>
                            <div className="border rounded-md p-4 max-h-60 overflow-y-auto">
                              {(() => {
                                const system =
                                  newRole.applicationId === "app_001"
                                    ? "CRM System"
                                    : newRole.applicationId === "app_002"
                                    ? "ERP System"
                                    : "Analytics Platform";
                                const permissions: Permission[] =
                                  availablePermissions[system] || [];

                                return Object.entries(
                                  permissions.reduce(
                                    (
                                      acc: Record<string, Permission[]>,
                                      perm: Permission
                                    ) => {
                                      if (!acc[perm.category])
                                        acc[perm.category] = [];
                                      acc[perm.category].push(perm);
                                      return acc;
                                    },
                                    {}
                                  )
                                ).map(
                                  ([category, perms]: [
                                    string,
                                    Permission[]
                                  ]) => (
                                    <div key={category} className="mb-4">
                                      <h4 className="font-medium mb-2">
                                        {category}
                                      </h4>
                                      <div className="space-y-2">
                                        {perms.map((perm: Permission) => (
                                          <div
                                            key={perm.id}
                                            className="flex items-center space-x-2"
                                          >
                                            <Checkbox
                                              id={perm.id}
                                              checked={newRole.permissions.includes(
                                                perm.id
                                              )}
                                              onCheckedChange={(checked) => {
                                                setNewRole({
                                                  ...newRole,
                                                  permissions: checked
                                                    ? [
                                                        ...newRole.permissions,
                                                        perm.id,
                                                      ]
                                                    : newRole.permissions.filter(
                                                        (p) => p !== perm.id
                                                      ),
                                                });
                                              }}
                                            />
                                            <Label
                                              htmlFor={perm.id}
                                              className="text-sm"
                                            >
                                              {perm.name}
                                            </Label>
                                          </div>
                                        ))}
                                      </div>
                                    </div>
                                  )
                                );
                              })()}
                            </div>
                          </div>
                        )}
                      </div>
                      <DialogFooter>
                        <Button
                          variant="outline"
                          onClick={() => setIsCreateDialogOpen(false)}
                        >
                          Cancelar
                        </Button>
                        <Button onClick={handleCreateRole}>Crear Rol</Button>
                      </DialogFooter>
                    </DialogContent>
                  </Dialog>
                </div>
              </div>

              {/* Tabla de roles */}
              <div className="rounded-md border">
                <Table>
                  <TableHeader>
                    <TableRow>
                      <TableHead>Rol</TableHead>
                      <TableHead>Aplicación</TableHead>
                      <TableHead>Nivel</TableHead>
                      <TableHead>Usuarios</TableHead>
                      <TableHead>Permisos</TableHead>
                      <TableHead>Estado</TableHead>
                      <TableHead className="text-right">Acciones</TableHead>
                    </TableRow>
                  </TableHeader>
                  <TableBody>
                    {filteredRoles.map((role) => (
                      <TableRow key={role.id}>
                        <TableCell>
                          <div>
                            <div className="font-medium">{role.name}</div>
                            <div className="text-sm text-muted-foreground">
                              {role.description}
                            </div>
                          </div>
                        </TableCell>
                        <TableCell>
                          <Badge variant="outline">
                            {role.applicationName}
                          </Badge>
                        </TableCell>
                        <TableCell>
                          <Badge className={getLevelBadge(role.level)}>
                            {role.level.charAt(0).toUpperCase() +
                              role.level.slice(1)}
                          </Badge>
                        </TableCell>
                        <TableCell>
                          <div className="flex items-center space-x-2">
                            <Users className="h-4 w-4 text-muted-foreground" />
                            <span className="font-medium">
                              {role.userCount}
                            </span>
                          </div>
                        </TableCell>
                        <TableCell>
                          <Collapsible>
                            <CollapsibleTrigger asChild>
                              <Button
                                variant="ghost"
                                size="sm"
                                onClick={() =>
                                  togglePermissionExpansion(role.id)
                                }
                              >
                                {expandedPermissions.includes(role.id) ? (
                                  <ChevronDown className="h-4 w-4" />
                                ) : (
                                  <ChevronRight className="h-4 w-4" />
                                )}
                                {role.permissions.length} permisos
                              </Button>
                            </CollapsibleTrigger>
                            <CollapsibleContent className="mt-2">
                              <div className="space-y-2">
                                {Object.entries(
                                  groupPermissionsByCategory(
                                    role.permissions,
                                    role.applicationName
                                  )
                                ).map(([category, perms]) => (
                                  <div key={category}>
                                    <div className="text-xs font-medium text-muted-foreground mb-1">
                                      {category}
                                    </div>
                                    <div className="flex flex-wrap gap-1">
                                      {perms.map((perm) => (
                                        <Badge
                                          key={perm.id}
                                          variant="secondary"
                                          className="text-xs"
                                        >
                                          {perm.name}
                                        </Badge>
                                      ))}
                                    </div>
                                  </div>
                                ))}
                              </div>
                            </CollapsibleContent>
                          </Collapsible>
                        </TableCell>
                        <TableCell>
                          <Badge
                            className={
                              role.isActive
                                ? "bg-green-100 text-green-800"
                                : "bg-gray-100 text-gray-800"
                            }
                          >
                            {role.isActive ? "Activo" : "Inactivo"}
                          </Badge>
                        </TableCell>
                        <TableCell className="text-right">
                          <div className="flex items-center justify-end space-x-2">
                            <Button variant="ghost" size="sm">
                              <Eye className="h-4 w-4" />
                            </Button>
                            <Button variant="ghost" size="sm">
                              <Edit className="h-4 w-4" />
                            </Button>
                            <Button variant="ghost" size="sm">
                              <Copy className="h-4 w-4" />
                            </Button>
                            <Button
                              variant="ghost"
                              size="sm"
                              className="text-red-600 hover:text-red-700"
                            >
                              <Trash2 className="h-4 w-4" />
                            </Button>
                          </div>
                        </TableCell>
                      </TableRow>
                    ))}
                  </TableBody>
                </Table>
              </div>
            </CardContent>
          </Card>
        </TabsContent>

        <TabsContent value="permissions" className="space-y-4">
          <Card>
            <CardHeader>
              <CardTitle>Matriz de Permisos</CardTitle>
              <CardDescription>
                Vista consolidada de todos los permisos por aplicación y rol
              </CardDescription>
            </CardHeader>
            <CardContent>
              <div className="space-y-6">
                {Object.entries(availablePermissions).map(
                  ([appName, permissions]) => (
                    <div key={appName} className="space-y-4">
                      <h3 className="text-lg font-semibold">{appName}</h3>
                      <div className="grid gap-4">
                        {Object.entries(
                          permissions.reduce(
                            (
                              acc: Record<
                                string,
                                { id: string; name: string; category: string }[]
                              >,
                              perm
                            ) => {
                              if (!acc[perm.category]) acc[perm.category] = [];
                              acc[perm.category].push(perm);
                              return acc;
                            },
                            {}
                          )
                        ).map(([category, perms]) => (
                          <Card key={category}>
                            <CardHeader className="pb-3">
                              <CardTitle className="text-base">
                                {category}
                              </CardTitle>
                            </CardHeader>
                            <CardContent>
                              <div className="grid gap-2">
                                {perms.map(
                                  (perm: { id: string; name: string }) => (
                                    <div
                                      key={perm.id}
                                      className="flex items-center justify-between p-2 border rounded"
                                    >
                                      <span className="text-sm">
                                        {perm.name}
                                      </span>
                                      <div className="flex space-x-2">
                                        {applicationRoles
                                          .filter(
                                            (role) =>
                                              role.applicationName === appName
                                          )
                                          .map((role) => (
                                            <Badge
                                              key={role.id}
                                              variant={
                                                role.permissions.includes(
                                                  perm.id
                                                )
                                                  ? "default"
                                                  : "outline"
                                              }
                                              className="text-xs"
                                            >
                                              {role.name}
                                            </Badge>
                                          ))}
                                      </div>
                                    </div>
                                  )
                                )}
                              </div>
                            </CardContent>
                          </Card>
                        ))}
                      </div>
                    </div>
                  )
                )}
              </div>
            </CardContent>
          </Card>
        </TabsContent>

        <TabsContent value="analytics" className="space-y-4">
          <div className="grid gap-4 md:grid-cols-2">
            <Card>
              <CardHeader>
                <CardTitle>Distribución por Nivel</CardTitle>
              </CardHeader>
              <CardContent>
                <div className="space-y-4">
                  {Object.entries(roleStats.byLevel).map(([level, count]) => (
                    <div
                      key={level}
                      className="flex items-center justify-between"
                    >
                      <div className="flex items-center space-x-2">
                        <div
                          className={`w-3 h-3 rounded-full ${
                            level === "admin"
                              ? "bg-red-500"
                              : level === "manager"
                              ? "bg-orange-500"
                              : level === "user"
                              ? "bg-blue-500"
                              : "bg-gray-500"
                          }`}
                        ></div>
                        <span className="capitalize">{level}</span>
                      </div>
                      <div className="flex items-center space-x-2">
                        <div className="w-24 bg-gray-200 rounded-full h-2">
                          <div
                            className={`h-2 rounded-full ${
                              level === "admin"
                                ? "bg-red-500"
                                : level === "manager"
                                ? "bg-orange-500"
                                : level === "user"
                                ? "bg-blue-500"
                                : "bg-gray-500"
                            }`}
                            style={{
                              width: `${(count / roleStats.total) * 100}%`,
                            }}
                          ></div>
                        </div>
                        <span className="text-sm font-medium">{count}</span>
                      </div>
                    </div>
                  ))}
                </div>
              </CardContent>
            </Card>

            <Card>
              <CardHeader>
                <CardTitle>Roles por Aplicación</CardTitle>
              </CardHeader>
              <CardContent>
                <div className="space-y-4">
                  {Object.entries(roleStats.byApplication).map(
                    ([app, count]) => (
                      <div
                        key={app}
                        className="flex items-center justify-between"
                      >
                        <span>{app}</span>
                        <div className="flex items-center space-x-2">
                          <div className="w-24 bg-gray-200 rounded-full h-2">
                            <div
                              className="bg-blue-600 h-2 rounded-full"
                              style={{
                                width: `${(count / roleStats.total) * 100}%`,
                              }}
                            ></div>
                          </div>
                          <span className="text-sm font-medium">{count}</span>
                        </div>
                      </div>
                    )
                  )}
                </div>
              </CardContent>
            </Card>
          </div>
        </TabsContent>
      </Tabs>
    </div>
  );
}
