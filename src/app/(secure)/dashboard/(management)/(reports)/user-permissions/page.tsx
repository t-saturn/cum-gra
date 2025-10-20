'use client';

import { useState } from 'react';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Badge } from '@/components/ui/badge';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table';
import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs';
import { Collapsible, CollapsibleContent, CollapsibleTrigger } from '@/components/ui/collapsible';
import { Search, Download, Shield, Lock, Unlock, ChevronDown, ChevronRight, AlertTriangle, CheckCircle, XCircle, Crown, Star, Award, RefreshCw } from 'lucide-react';
import CardStatsContain from '@/components/custom/card/card-stats-contain';
import { statsPermissionsForUsers } from '@/mocks/stats-mocks';

// Mock data
const userPermissionsData = [
  {
    id: '1',
    user: {
      name: 'Juan Carlos Pérez',
      email: 'jperez@empresa.com',
      avatar: '/placeholder.svg?height=40&width=40',
      position: 'Gerente General',
      department: 'Gerencia',
      status: 'active',
    },
    roles: [
      {
        id: 'admin',
        name: 'Administrador',
        application: 'Sistema Global',
        level: 'high',
        permissions: [
          { module: 'users', actions: ['read', 'write', 'delete'], granted: true },
          { module: 'applications', actions: ['read', 'write', 'delete'], granted: true },
          { module: 'roles', actions: ['read', 'write', 'delete'], granted: true },
          { module: 'reports', actions: ['read', 'write'], granted: true },
        ],
      },
      {
        id: 'inv_manager',
        name: 'Gerente de Inventario',
        application: 'Sistema de Inventario',
        level: 'medium',
        permissions: [
          { module: 'products', actions: ['read', 'write'], granted: true },
          { module: 'warehouse', actions: ['read', 'write'], granted: true },
          { module: 'reports', actions: ['read'], granted: true },
        ],
      },
    ],
    totalPermissions: 15,
    riskLevel: 'low',
    lastReview: '2024-01-10',
    createdAt: '2024-01-01',
  },
  {
    id: '2',
    user: {
      name: 'María García López',
      email: 'mgarcia@empresa.com',
      avatar: '/placeholder.svg?height=40&width=40',
      position: 'Analista Senior',
      department: 'Sistemas',
      status: 'active',
    },
    roles: [
      {
        id: 'dev_lead',
        name: 'Líder de Desarrollo',
        application: 'Sistema de Inventario',
        level: 'medium',
        permissions: [
          { module: 'products', actions: ['read', 'write'], granted: true },
          { module: 'warehouse', actions: ['read'], granted: true },
          { module: 'users', actions: ['read'], granted: true },
          { module: 'reports', actions: ['read'], granted: true },
        ],
      },
      {
        id: 'hr_analyst',
        name: 'Analista RRHH',
        application: 'Portal RRHH',
        level: 'low',
        permissions: [
          { module: 'employees', actions: ['read'], granted: true },
          { module: 'reports', actions: ['read'], granted: true },
        ],
      },
    ],
    totalPermissions: 8,
    riskLevel: 'low',
    lastReview: '2024-01-12',
    createdAt: '2024-01-02',
  },
  {
    id: '3',
    user: {
      name: 'Carlos López Ruiz',
      email: 'clopez@empresa.com',
      avatar: '/placeholder.svg?height=40&width=40',
      position: 'Desarrollador',
      department: 'Sistemas',
      status: 'active',
    },
    roles: [
      {
        id: 'developer',
        name: 'Desarrollador',
        application: 'Sistema de Inventario',
        level: 'low',
        permissions: [
          { module: 'products', actions: ['read'], granted: true },
          { module: 'warehouse', actions: ['read'], granted: true },
          { module: 'reports', actions: ['read'], granted: false },
        ],
      },
      {
        id: 'finance_user',
        name: 'Usuario Financiero',
        application: 'Sistema Financiero',
        level: 'medium',
        permissions: [
          { module: 'accounting', actions: ['read', 'write'], granted: true },
          { module: 'invoicing', actions: ['read', 'write'], granted: true },
          { module: 'reports', actions: ['read'], granted: true },
        ],
      },
    ],
    totalPermissions: 7,
    riskLevel: 'medium',
    lastReview: '2024-01-08',
    createdAt: '2024-01-03',
  },
  {
    id: '4',
    user: {
      name: 'Ana Martínez Silva',
      email: 'amartinez@empresa.com',
      avatar: '/placeholder.svg?height=40&width=40',
      position: 'Especialista TI',
      department: 'Infraestructura',
      status: 'active',
    },
    roles: [
      {
        id: 'it_specialist',
        name: 'Especialista TI',
        application: 'Sistema Global',
        level: 'high',
        permissions: [
          { module: 'users', actions: ['read', 'write'], granted: true },
          { module: 'applications', actions: ['read'], granted: true },
          { module: 'security', actions: ['read', 'write'], granted: true },
          { module: 'audit', actions: ['read'], granted: true },
        ],
      },
    ],
    totalPermissions: 6,
    riskLevel: 'high',
    lastReview: '2024-01-05',
    createdAt: '2024-01-04',
  },
  {
    id: '5',
    user: {
      name: 'Luis Fernando Torres',
      email: 'ltorres@empresa.com',
      avatar: '/placeholder.svg?height=40&width=40',
      position: 'Gerente RRHH',
      department: 'Recursos Humanos',
      status: 'suspended',
    },
    roles: [
      {
        id: 'hr_manager',
        name: 'Gerente RRHH',
        application: 'Portal RRHH',
        level: 'high',
        permissions: [
          { module: 'employees', actions: ['read', 'write', 'delete'], granted: false },
          { module: 'payroll', actions: ['read', 'write'], granted: false },
          { module: 'performance', actions: ['read', 'write'], granted: false },
          { module: 'reports', actions: ['read', 'write'], granted: false },
        ],
      },
    ],
    totalPermissions: 8,
    riskLevel: 'high',
    lastReview: '2024-01-01',
    createdAt: '2024-01-05',
  },
];

export default function UserPermissionsReport() {
  const [searchTerm, setSearchTerm] = useState('');
  const [departmentFilter, setDepartmentFilter] = useState('all');
  const [riskFilter, setRiskFilter] = useState('all');
  const [statusFilter, setStatusFilter] = useState('all');
  const [expandedUsers, setExpandedUsers] = useState<string[]>([]);

  const filteredUsers = userPermissionsData.filter((user) => {
    const matchesSearch =
      user.user.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
      user.user.email.toLowerCase().includes(searchTerm.toLowerCase()) ||
      user.user.department.toLowerCase().includes(searchTerm.toLowerCase());

    const matchesDepartment = departmentFilter === 'all' || user.user.department === departmentFilter;
    const matchesRisk = riskFilter === 'all' || user.riskLevel === riskFilter;
    const matchesStatus = statusFilter === 'all' || user.user.status === statusFilter;

    return matchesSearch && matchesDepartment && matchesRisk && matchesStatus;
  });

  const getRiskBadge = (risk: string) => {
    switch (risk) {
      case 'high':
        return (
          <Badge className="bg-destructive/20 text-destructive border-destructive/30">
            <AlertTriangle className="w-3 h-3 mr-1" />
            Alto
          </Badge>
        );
      case 'medium':
        return (
          <Badge className="bg-chart-5/20 text-chart-5 border-chart-5/30">
            <AlertTriangle className="w-3 h-3 mr-1" />
            Medio
          </Badge>
        );
      case 'low':
        return (
          <Badge className="bg-chart-4/20 text-chart-4 border-chart-4/30">
            <CheckCircle className="w-3 h-3 mr-1" />
            Bajo
          </Badge>
        );
      default:
        return <Badge variant="secondary">Desconocido</Badge>;
    }
  };

  const getStatusBadge = (status: string) => {
    switch (status) {
      case 'active':
        return (
          <Badge className="bg-chart-4/20 text-chart-4 border-chart-4/30">
            <CheckCircle className="w-3 h-3 mr-1" />
            Activo
          </Badge>
        );
      case 'suspended':
        return (
          <Badge className="bg-destructive/20 text-destructive border-destructive/30">
            <XCircle className="w-3 h-3 mr-1" />
            Suspendido
          </Badge>
        );
      case 'inactive':
        return (
          <Badge className="bg-muted text-muted-foreground border-muted-foreground/30">
            <XCircle className="w-3 h-3 mr-1" />
            Inactivo
          </Badge>
        );
      default:
        return <Badge variant="secondary">Desconocido</Badge>;
    }
  };

  const getRoleLevelBadge = (level: string) => {
    switch (level) {
      case 'high':
        return (
          <Badge className="bg-primary/20 text-primary border-primary/30">
            <Crown className="w-3 h-3 mr-1" />
            Alto
          </Badge>
        );
      case 'medium':
        return (
          <Badge className="bg-chart-2/20 text-chart-2 border-chart-2/30">
            <Star className="w-3 h-3 mr-1" />
            Medio
          </Badge>
        );
      case 'low':
        return (
          <Badge className="bg-chart-3/20 text-chart-3 border-chart-3/30">
            <Award className="w-3 h-3 mr-1" />
            Básico
          </Badge>
        );
      default:
        return <Badge variant="secondary">Desconocido</Badge>;
    }
  };

  const getPermissionIcon = (actions: string[], granted: boolean) => {
    if (!granted) return <XCircle className="w-4 h-4 text-destructive" />;
    if (actions.includes('delete')) return <Shield className="w-4 h-4 text-destructive" />;
    if (actions.includes('write')) return <Lock className="w-4 h-4 text-chart-5" />;
    return <Unlock className="w-4 h-4 text-chart-4" />;
  };

  const toggleUserExpansion = (userId: string) => {
    setExpandedUsers((prev) => (prev.includes(userId) ? prev.filter((id) => id !== userId) : [...prev, userId]));
  };

  const formatDate = (dateString: string) => {
    return new Date(dateString).toLocaleDateString('es-ES');
  };

  const getPermissionSummary = (roles: (typeof userPermissionsData)[0]['roles']) => {
    const allPermissions = roles.flatMap((role) => role.permissions);
    const granted = allPermissions.filter((p) => p.granted).length;
    const total = allPermissions.length;
    return { granted, total };
  };

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold text-foreground">Permisos por Usuario</h1>
          <p className="text-muted-foreground mt-1">Análisis detallado de roles y permisos asignados</p>
        </div>
        <div className="flex gap-3">
          <Button variant="outline">
            <RefreshCw className="w-4 h-4 mr-2" />
            Actualizar
          </Button>
          <Button variant="outline">
            <Download className="w-4 h-4 mr-2" />
            Exportar
          </Button>
        </div>
      </div>

      {/* Stats Cards */}
      <CardStatsContain stats={statsPermissionsForUsers} />

      <Tabs defaultValue="users" className="w-full">
        <TabsList className="grid w-full grid-cols-3">
          <TabsTrigger value="users">Usuarios y Permisos</TabsTrigger>
          <TabsTrigger value="roles">Análisis de Roles</TabsTrigger>
          <TabsTrigger value="matrix">Matriz de Permisos</TabsTrigger>
        </TabsList>

        <TabsContent value="users" className="space-y-4">
          {/* Filters */}
          <Card className="border-border bg-card/50">
            <CardHeader>
              <div className="flex items-center justify-between">
                <div>
                  <CardTitle>Permisos de Usuarios</CardTitle>
                  <CardDescription>
                    {filteredUsers.length} de {userPermissionsData.length} usuarios mostrados
                  </CardDescription>
                </div>
                <div className="flex gap-2">
                  <div className="relative">
                    <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 text-muted-foreground w-4 h-4" />
                    <Input
                      placeholder="Buscar usuarios..."
                      value={searchTerm}
                      onChange={(e) => setSearchTerm(e.target.value)}
                      className="pl-10 w-80 bg-background/50 border-border focus:border-primary focus:ring-ring"
                    />
                  </div>
                  <Select value={departmentFilter} onValueChange={setDepartmentFilter}>
                    <SelectTrigger className="w-40">
                      <SelectValue placeholder="Departamento" />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value="all">Todos</SelectItem>
                      <SelectItem value="Gerencia">Gerencia</SelectItem>
                      <SelectItem value="Sistemas">Sistemas</SelectItem>
                      <SelectItem value="Recursos Humanos">RRHH</SelectItem>
                      <SelectItem value="Infraestructura">Infraestructura</SelectItem>
                    </SelectContent>
                  </Select>
                  <Select value={riskFilter} onValueChange={setRiskFilter}>
                    <SelectTrigger className="w-32">
                      <SelectValue placeholder="Riesgo" />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value="all">Todos</SelectItem>
                      <SelectItem value="high">Alto</SelectItem>
                      <SelectItem value="medium">Medio</SelectItem>
                      <SelectItem value="low">Bajo</SelectItem>
                    </SelectContent>
                  </Select>
                  <Select value={statusFilter} onValueChange={setStatusFilter}>
                    <SelectTrigger className="w-32">
                      <SelectValue placeholder="Estado" />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value="all">Todos</SelectItem>
                      <SelectItem value="active">Activo</SelectItem>
                      <SelectItem value="suspended">Suspendido</SelectItem>
                      <SelectItem value="inactive">Inactivo</SelectItem>
                    </SelectContent>
                  </Select>
                </div>
              </div>
            </CardHeader>
            <CardContent>
              <div className="space-y-4">
                {filteredUsers.map((user) => {
                  const isExpanded = expandedUsers.includes(user.id);
                  const permissionSummary = getPermissionSummary(user.roles);

                  return (
                    <Card key={user.id} className="border-border bg-accent/20">
                      <Collapsible>
                        <CollapsibleTrigger className="w-full" onClick={() => toggleUserExpansion(user.id)}>
                          <CardContent className="p-4">
                            <div className="flex items-center justify-between">
                              <div className="flex items-center gap-4">
                                <Avatar className="w-12 h-12">
                                  <AvatarImage src={user.user.avatar || '/placeholder.svg'} />
                                  <AvatarFallback className="bg-gradient-to-r from-primary to-chart-1 text-primary-foreground font-semibold">
                                    {user.user.name
                                      .split(' ')
                                      .map((n) => n[0])
                                      .join('')}
                                  </AvatarFallback>
                                </Avatar>
                                <div className="text-left">
                                  <p className="font-medium text-foreground">{user.user.name}</p>
                                  <p className="text-sm text-muted-foreground">{user.user.email}</p>
                                  <p className="text-xs text-muted-foreground">
                                    {user.user.position} - {user.user.department}
                                  </p>
                                </div>
                              </div>
                              <div className="flex items-center gap-4">
                                <div className="text-center">
                                  <p className="text-lg font-bold text-primary">{user.roles.length}</p>
                                  <p className="text-xs text-muted-foreground">Roles</p>
                                </div>
                                <div className="text-center">
                                  <p className="text-lg font-bold text-chart-2">
                                    {permissionSummary.granted}/{permissionSummary.total}
                                  </p>
                                  <p className="text-xs text-muted-foreground">Permisos</p>
                                </div>
                                <div className="flex flex-col gap-2">
                                  {getStatusBadge(user.user.status)}
                                  {getRiskBadge(user.riskLevel)}
                                </div>
                                <div className="text-right">
                                  <p className="text-sm text-muted-foreground">Última revisión</p>
                                  <p className="text-sm font-medium">{formatDate(user.lastReview)}</p>
                                </div>
                                {isExpanded ? <ChevronDown className="w-5 h-5 text-muted-foreground" /> : <ChevronRight className="w-5 h-5 text-muted-foreground" />}
                              </div>
                            </div>
                          </CardContent>
                        </CollapsibleTrigger>
                        <CollapsibleContent>
                          <CardContent className="pt-0 px-4 pb-4">
                            <div className="space-y-4">
                              <div className="border-t border-border pt-4">
                                <h4 className="font-semibold mb-3">Roles y Permisos Detallados</h4>
                                <div className="space-y-4">
                                  {user.roles.map((role, roleIndex) => (
                                    <div key={roleIndex} className="bg-background/50 rounded-lg p-4">
                                      <div className="flex items-center justify-between mb-3">
                                        <div>
                                          <div className="flex items-center gap-2">
                                            <h5 className="font-medium">{role.name}</h5>
                                            {getRoleLevelBadge(role.level)}
                                          </div>
                                          <p className="text-sm text-muted-foreground">{role.application}</p>
                                        </div>
                                      </div>
                                      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-3">
                                        {role.permissions.map((permission, permIndex) => (
                                          <div key={permIndex} className="flex items-center justify-between p-3 bg-accent/30 rounded">
                                            <div className="flex items-center gap-2">
                                              {getPermissionIcon(permission.actions, permission.granted)}
                                              <div>
                                                <p className="text-sm font-medium">{permission.module}</p>
                                                <p className="text-xs text-muted-foreground">{permission.actions.join(', ')}</p>
                                              </div>
                                            </div>
                                            <Badge variant={permission.granted ? 'default' : 'destructive'} className="text-xs">
                                              {permission.granted ? 'Concedido' : 'Denegado'}
                                            </Badge>
                                          </div>
                                        ))}
                                      </div>
                                    </div>
                                  ))}
                                </div>
                              </div>
                            </div>
                          </CardContent>
                        </CollapsibleContent>
                      </Collapsible>
                    </Card>
                  );
                })}
              </div>
            </CardContent>
          </Card>
        </TabsContent>

        <TabsContent value="roles" className="space-y-4">
          <Card className="border-border bg-card/50">
            <CardHeader>
              <CardTitle>Análisis de Roles</CardTitle>
              <CardDescription>Distribución y uso de roles en el sistema</CardDescription>
            </CardHeader>
            <CardContent>
              <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
                <div className="space-y-4">
                  <h4 className="font-semibold">Roles por Nivel</h4>
                  <div className="space-y-3">
                    <div className="flex items-center justify-between">
                      <div className="flex items-center gap-2">
                        <Crown className="w-4 h-4 text-primary" />
                        <span className="text-sm">Alto</span>
                      </div>
                      <div className="flex items-center gap-2">
                        <div className="w-20 bg-muted rounded-full h-2">
                          <div className="bg-primary h-2 rounded-full" style={{ width: '35%' }} />
                        </div>
                        <span className="text-sm font-medium">35%</span>
                      </div>
                    </div>
                    <div className="flex items-center justify-between">
                      <div className="flex items-center gap-2">
                        <Star className="w-4 h-4 text-chart-2" />
                        <span className="text-sm">Medio</span>
                      </div>
                      <div className="flex items-center gap-2">
                        <div className="w-20 bg-muted rounded-full h-2">
                          <div className="bg-chart-2 h-2 rounded-full" style={{ width: '45%' }} />
                        </div>
                        <span className="text-sm font-medium">45%</span>
                      </div>
                    </div>
                    <div className="flex items-center justify-between">
                      <div className="flex items-center gap-2">
                        <Award className="w-4 h-4 text-chart-3" />
                        <span className="text-sm">Básico</span>
                      </div>
                      <div className="flex items-center gap-2">
                        <div className="w-20 bg-muted rounded-full h-2">
                          <div className="bg-chart-3 h-2 rounded-full" style={{ width: '20%' }} />
                        </div>
                        <span className="text-sm font-medium">20%</span>
                      </div>
                    </div>
                  </div>
                </div>

                <div className="space-y-4">
                  <h4 className="font-semibold">Usuarios por Riesgo</h4>
                  <div className="space-y-3">
                    <div className="flex items-center justify-between">
                      <div className="flex items-center gap-2">
                        <AlertTriangle className="w-4 h-4 text-destructive" />
                        <span className="text-sm">Alto Riesgo</span>
                      </div>
                      <div className="flex items-center gap-2">
                        <div className="w-20 bg-muted rounded-full h-2">
                          <div className="bg-destructive h-2 rounded-full" style={{ width: '15%' }} />
                        </div>
                        <span className="text-sm font-medium">15%</span>
                      </div>
                    </div>
                    <div className="flex items-center justify-between">
                      <div className="flex items-center gap-2">
                        <AlertTriangle className="w-4 h-4 text-chart-5" />
                        <span className="text-sm">Riesgo Medio</span>
                      </div>
                      <div className="flex items-center gap-2">
                        <div className="w-20 bg-muted rounded-full h-2">
                          <div className="bg-chart-5 h-2 rounded-full" style={{ width: '25%' }} />
                        </div>
                        <span className="text-sm font-medium">25%</span>
                      </div>
                    </div>
                    <div className="flex items-center justify-between">
                      <div className="flex items-center gap-2">
                        <CheckCircle className="w-4 h-4 text-chart-4" />
                        <span className="text-sm">Bajo Riesgo</span>
                      </div>
                      <div className="flex items-center gap-2">
                        <div className="w-20 bg-muted rounded-full h-2">
                          <div className="bg-chart-4 h-2 rounded-full" style={{ width: '60%' }} />
                        </div>
                        <span className="text-sm font-medium">60%</span>
                      </div>
                    </div>
                  </div>
                </div>

                <div className="space-y-4">
                  <h4 className="font-semibold">Aplicaciones Más Usadas</h4>
                  <div className="space-y-3">
                    <div className="flex items-center justify-between">
                      <span className="text-sm">Sistema de Inventario</span>
                      <Badge variant="outline">45 usuarios</Badge>
                    </div>
                    <div className="flex items-center justify-between">
                      <span className="text-sm">Portal RRHH</span>
                      <Badge variant="outline">32 usuarios</Badge>
                    </div>
                    <div className="flex items-center justify-between">
                      <span className="text-sm">Sistema Financiero</span>
                      <Badge variant="outline">12 usuarios</Badge>
                    </div>
                    <div className="flex items-center justify-between">
                      <span className="text-sm">App Móvil Ventas</span>
                      <Badge variant="outline">25 usuarios</Badge>
                    </div>
                  </div>
                </div>
              </div>
            </CardContent>
          </Card>
        </TabsContent>

        <TabsContent value="matrix" className="space-y-4">
          <Card className="border-border bg-card/50">
            <CardHeader>
              <CardTitle>Matriz de Permisos</CardTitle>
              <CardDescription>Vista consolidada de permisos por módulo y acción</CardDescription>
            </CardHeader>
            <CardContent>
              <div className="overflow-x-auto">
                <Table>
                  <TableHeader>
                    <TableRow className="bg-accent/50">
                      <TableHead>Módulo</TableHead>
                      <TableHead className="text-center">Lectura</TableHead>
                      <TableHead className="text-center">Escritura</TableHead>
                      <TableHead className="text-center">Eliminación</TableHead>
                      <TableHead className="text-center">Total Usuarios</TableHead>
                    </TableRow>
                  </TableHeader>
                  <TableBody>
                    <TableRow>
                      <TableCell className="font-medium">Usuarios</TableCell>
                      <TableCell className="text-center">
                        <Badge className="bg-chart-4/20 text-chart-4">89</Badge>
                      </TableCell>
                      <TableCell className="text-center">
                        <Badge className="bg-chart-5/20 text-chart-5">45</Badge>
                      </TableCell>
                      <TableCell className="text-center">
                        <Badge className="bg-destructive/20 text-destructive">12</Badge>
                      </TableCell>
                      <TableCell className="text-center">
                        <Badge variant="outline">89</Badge>
                      </TableCell>
                    </TableRow>
                    <TableRow>
                      <TableCell className="font-medium">Aplicaciones</TableCell>
                      <TableCell className="text-center">
                        <Badge className="bg-chart-4/20 text-chart-4">67</Badge>
                      </TableCell>
                      <TableCell className="text-center">
                        <Badge className="bg-chart-5/20 text-chart-5">23</Badge>
                      </TableCell>
                      <TableCell className="text-center">
                        <Badge className="bg-destructive/20 text-destructive">8</Badge>
                      </TableCell>
                      <TableCell className="text-center">
                        <Badge variant="outline">67</Badge>
                      </TableCell>
                    </TableRow>
                    <TableRow>
                      <TableCell className="font-medium">Productos</TableCell>
                      <TableCell className="text-center">
                        <Badge className="bg-chart-4/20 text-chart-4">45</Badge>
                      </TableCell>
                      <TableCell className="text-center">
                        <Badge className="bg-chart-5/20 text-chart-5">32</Badge>
                      </TableCell>
                      <TableCell className="text-center">
                        <Badge className="bg-destructive/20 text-destructive">5</Badge>
                      </TableCell>
                      <TableCell className="text-center">
                        <Badge variant="outline">45</Badge>
                      </TableCell>
                    </TableRow>
                    <TableRow>
                      <TableCell className="font-medium">Empleados</TableCell>
                      <TableCell className="text-center">
                        <Badge className="bg-chart-4/20 text-chart-4">32</Badge>
                      </TableCell>
                      <TableCell className="text-center">
                        <Badge className="bg-chart-5/20 text-chart-5">18</Badge>
                      </TableCell>
                      <TableCell className="text-center">
                        <Badge className="bg-destructive/20 text-destructive">3</Badge>
                      </TableCell>
                      <TableCell className="text-center">
                        <Badge variant="outline">32</Badge>
                      </TableCell>
                    </TableRow>
                    <TableRow>
                      <TableCell className="font-medium">Reportes</TableCell>
                      <TableCell className="text-center">
                        <Badge className="bg-chart-4/20 text-chart-4">78</Badge>
                      </TableCell>
                      <TableCell className="text-center">
                        <Badge className="bg-chart-5/20 text-chart-5">34</Badge>
                      </TableCell>
                      <TableCell className="text-center">
                        <Badge className="bg-destructive/20 text-destructive">2</Badge>
                      </TableCell>
                      <TableCell className="text-center">
                        <Badge variant="outline">78</Badge>
                      </TableCell>
                    </TableRow>
                  </TableBody>
                </Table>
              </div>
            </CardContent>
          </Card>
        </TabsContent>
      </Tabs>
    </div>
  );
}
