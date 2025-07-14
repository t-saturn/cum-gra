'use client';

import { useState } from 'react';
import { Badge } from '@/components/ui/badge';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table';
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs';
import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar';
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle, DialogTrigger } from '@/components/ui/dialog';
import { Checkbox } from '@/components/ui/checkbox';
import { Switch } from '@/components/ui/switch';
import { Collapsible, CollapsibleContent, CollapsibleTrigger } from '@/components/ui/collapsible';
import { Search, Filter, Plus, Edit, Trash2, Users, Shield, Crown, Eye, ChevronDown, ChevronRight, UserCheck, Settings } from 'lucide-react';
import CardStatsContain from '@/components/custom/card/card-stats-contain';
import { statsUsersAndRoles } from '@/mocks/stats-mocks';

// Datos simulados de usuarios con roles
const userRoles = [
  {
    id: 'user_001',
    name: 'Ana García',
    email: 'ana.garcia@empresa.com',
    avatar: '/placeholder.svg?height=32&width=32',
    department: 'Administración',
    position: 'Gerente General',
    isActive: true,
    lastLogin: '2024-01-15T14:30:00Z',
    roles: [
      {
        id: 'role_001',
        name: 'Super Admin',
        type: 'system',
        level: 'admin',
        applicationName: 'Sistema',
        permissions: ['*'],
        assignedAt: '2024-01-01T00:00:00Z',
        assignedBy: 'Sistema',
        expiresAt: null,
        isActive: true,
      },
      {
        id: 'role_002',
        name: 'CRM Admin',
        type: 'application',
        level: 'admin',
        applicationName: 'CRM System',
        permissions: ['crm.*'],
        assignedAt: '2024-01-01T00:00:00Z',
        assignedBy: 'Sistema',
        expiresAt: null,
        isActive: true,
      },
    ],
    totalPermissions: 156,
    riskLevel: 'high',
  },
  {
    id: 'user_002',
    name: 'Carlos López',
    email: 'carlos.lopez@empresa.com',
    avatar: '/placeholder.svg?height=32&width=32',
    department: 'Ventas',
    position: 'Director de Ventas',
    isActive: true,
    lastLogin: '2024-01-15T13:45:00Z',
    roles: [
      {
        id: 'role_003',
        name: 'Sales Manager',
        type: 'application',
        level: 'manager',
        applicationName: 'CRM System',
        permissions: ['crm.leads.*', 'crm.opportunities.*', 'crm.reports.read'],
        assignedAt: '2024-01-05T10:00:00Z',
        assignedBy: 'Ana García',
        expiresAt: null,
        isActive: true,
      },
      {
        id: 'role_004',
        name: 'Team Lead',
        type: 'organizational',
        level: 'manager',
        applicationName: 'Sistema',
        permissions: ['team.manage', 'reports.team'],
        assignedAt: '2024-01-05T10:15:00Z',
        assignedBy: 'Ana García',
        expiresAt: null,
        isActive: true,
      },
    ],
    totalPermissions: 23,
    riskLevel: 'medium',
  },
  {
    id: 'user_003',
    name: 'María Rodríguez',
    email: 'maria.rodriguez@empresa.com',
    avatar: '/placeholder.svg?height=32&width=32',
    department: 'Finanzas',
    position: 'Contadora',
    isActive: true,
    lastLogin: '2024-01-15T12:20:00Z',
    roles: [
      {
        id: 'role_005',
        name: 'Finance User',
        type: 'application',
        level: 'user',
        applicationName: 'ERP System',
        permissions: ['erp.finance.read', 'erp.accounting.*', 'erp.reports.financial'],
        assignedAt: '2024-01-08T14:30:00Z',
        assignedBy: 'Ana García',
        expiresAt: null,
        isActive: true,
      },
    ],
    totalPermissions: 12,
    riskLevel: 'low',
  },
  {
    id: 'user_004',
    name: 'David Martín',
    email: 'david.martin@empresa.com',
    avatar: '/placeholder.svg?height=32&width=32',
    department: 'IT',
    position: 'Desarrollador',
    isActive: false,
    lastLogin: '2024-01-10T16:45:00Z',
    roles: [
      {
        id: 'role_006',
        name: 'Developer',
        type: 'application',
        level: 'user',
        applicationName: 'Sistema',
        permissions: ['system.read', 'logs.read', 'debug.access'],
        assignedAt: '2024-01-03T09:00:00Z',
        assignedBy: 'Ana García',
        expiresAt: '2024-02-01T00:00:00Z',
        isActive: false,
      },
    ],
    totalPermissions: 8,
    riskLevel: 'low',
  },
  {
    id: 'user_005',
    name: 'Laura Fernández',
    email: 'laura.fernandez@empresa.com',
    avatar: '/placeholder.svg?height=32&width=32',
    department: 'Recursos Humanos',
    position: 'Especialista en RRHH',
    isActive: true,
    lastLogin: '2024-01-15T11:30:00Z',
    roles: [
      {
        id: 'role_007',
        name: 'HR Specialist',
        type: 'application',
        level: 'user',
        applicationName: 'HR System',
        permissions: ['hr.employees.read', 'hr.employees.write', 'hr.reports.basic'],
        assignedAt: '2024-01-12T13:15:00Z',
        assignedBy: 'Ana García',
        expiresAt: null,
        isActive: true,
      },
      {
        id: 'role_008',
        name: 'User Manager',
        type: 'system',
        level: 'manager',
        applicationName: 'Sistema',
        permissions: ['users.read', 'users.write', 'roles.assign.basic'],
        assignedAt: '2024-01-12T13:20:00Z',
        assignedBy: 'Ana García',
        expiresAt: null,
        isActive: true,
      },
    ],
    totalPermissions: 18,
    riskLevel: 'medium',
  },
];

const availableRoles = [
  {
    id: 'role_001',
    name: 'Super Admin',
    type: 'system',
    level: 'admin',
    applicationName: 'Sistema',
    description: 'Acceso completo al sistema',
    permissions: ['*'],
    userCount: 1,
    isActive: true,
  },
  {
    id: 'role_002',
    name: 'CRM Admin',
    type: 'application',
    level: 'admin',
    applicationName: 'CRM System',
    description: 'Administrador del sistema CRM',
    permissions: ['crm.*'],
    userCount: 2,
    isActive: true,
  },
  {
    id: 'role_003',
    name: 'Sales Manager',
    type: 'application',
    level: 'manager',
    applicationName: 'CRM System',
    description: 'Gestor de ventas',
    permissions: ['crm.leads.*', 'crm.opportunities.*', 'crm.reports.read'],
    userCount: 5,
    isActive: true,
  },
  {
    id: 'role_004',
    name: 'Sales User',
    type: 'application',
    level: 'user',
    applicationName: 'CRM System',
    description: 'Usuario de ventas',
    permissions: ['crm.leads.read', 'crm.leads.write', 'crm.opportunities.read'],
    userCount: 15,
    isActive: true,
  },
  {
    id: 'role_005',
    name: 'Finance Manager',
    type: 'application',
    level: 'manager',
    applicationName: 'ERP System',
    description: 'Gestor financiero',
    permissions: ['erp.finance.*', 'erp.accounting.*', 'erp.reports.financial.*'],
    userCount: 3,
    isActive: true,
  },
  {
    id: 'role_006',
    name: 'Finance User',
    type: 'application',
    level: 'user',
    applicationName: 'ERP System',
    description: 'Usuario financiero',
    permissions: ['erp.finance.read', 'erp.accounting.read', 'erp.reports.financial.read'],
    userCount: 8,
    isActive: true,
  },
];

const roleStats = {
  totalUsers: 156,
  activeUsers: 142,
  inactiveUsers: 14,
  totalRoles: 45,
  byLevel: {
    admin: 8,
    manager: 23,
    user: 125,
  },
  byType: {
    system: 12,
    application: 89,
    organizational: 55,
  },
  byRisk: {
    low: 98,
    medium: 43,
    high: 15,
  },
};

export default function UserRolesManagement() {
  const [searchTerm, setSearchTerm] = useState('');
  const [departmentFilter, setDepartmentFilter] = useState('all');
  const [roleFilter, setRoleFilter] = useState('all');
  const [statusFilter, setStatusFilter] = useState('all');
  const [isAssignDialogOpen, setIsAssignDialogOpen] = useState(false);
  const [expandedUsers, setExpandedUsers] = useState<string[]>([]);

  const [assignmentForm, setAssignmentForm] = useState({
    userId: '',
    roleIds: [] as string[],
    expiresAt: '',
    reason: '',
  });

  const getLevelBadge = (level: string) => {
    const colors = {
      admin: 'bg-red-100 text-red-800',
      manager: 'bg-orange-100 text-orange-800',
      user: 'bg-blue-100 text-blue-800',
      viewer: 'bg-gray-100 text-gray-800',
    };
    return colors[level as keyof typeof colors] || 'bg-gray-100 text-gray-800';
  };

  const getRiskBadge = (risk: string) => {
    const colors = {
      low: 'bg-green-100 text-green-800',
      medium: 'bg-yellow-100 text-yellow-800',
      high: 'bg-red-100 text-red-800',
    };
    return colors[risk as keyof typeof colors] || 'bg-gray-100 text-gray-800';
  };

  const getTypeIcon = (type: string) => {
    switch (type) {
      case 'system':
        return <Crown className="h-4 w-4" />;
      case 'application':
        return <Settings className="h-4 w-4" />;
      case 'organizational':
        return <Users className="h-4 w-4" />;
      default:
        return <Shield className="h-4 w-4" />;
    }
  };

  const toggleUserExpansion = (userId: string) => {
    setExpandedUsers((prev) => (prev.includes(userId) ? prev.filter((id) => id !== userId) : [...prev, userId]));
  };

  const filteredUsers = userRoles.filter((user) => {
    const matchesSearch =
      user.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
      user.email.toLowerCase().includes(searchTerm.toLowerCase()) ||
      user.department.toLowerCase().includes(searchTerm.toLowerCase());
    const matchesDepartment = departmentFilter === 'all' || user.department === departmentFilter;
    const matchesRole = roleFilter === 'all' || user.roles.some((role) => role.name === roleFilter);
    const matchesStatus = statusFilter === 'all' || (statusFilter === 'active' ? user.isActive : !user.isActive);

    return matchesSearch && matchesDepartment && matchesRole && matchesStatus;
  });

  const handleAssignRoles = () => {
    console.log('Assigning roles:', assignmentForm);
    setIsAssignDialogOpen(false);
    setAssignmentForm({
      userId: '',
      roleIds: [],
      expiresAt: '',
      reason: '',
    });
  };

  const formatDate = (dateString: string) => {
    return new Date(dateString).toLocaleString('es-ES', {
      year: 'numeric',
      month: '2-digit',
      day: '2-digit',
      hour: '2-digit',
      minute: '2-digit',
    });
  };

  return (
    <div className="space-y-6">
      {/* Estadísticas */}
      <CardStatsContain stats={statsUsersAndRoles} />

      <Tabs defaultValue="users" className="space-y-4">
        <TabsList>
          <TabsTrigger value="users">Usuarios con Roles</TabsTrigger>
          <TabsTrigger value="roles">Roles Disponibles</TabsTrigger>
          <TabsTrigger value="assignments">Asignaciones</TabsTrigger>
          <TabsTrigger value="analytics">Análisis</TabsTrigger>
        </TabsList>

        <TabsContent value="users" className="space-y-4">
          <Card>
            <CardHeader>
              <CardTitle>Gestión de Roles por Usuario</CardTitle>
              <CardDescription>Vista de usuarios con sus roles y permisos asignados</CardDescription>
            </CardHeader>
            <CardContent>
              {/* Filtros */}
              <div className="flex flex-col gap-4 md:flex-row md:items-center md:justify-between mb-6">
                <div className="flex flex-1 items-center space-x-2">
                  <Search className="h-4 w-4 text-muted-foreground" />
                  <Input placeholder="Buscar usuarios..." value={searchTerm} onChange={(e) => setSearchTerm(e.target.value)} className="max-w-sm" />
                </div>
                <div className="flex items-center space-x-2">
                  <Filter className="h-4 w-4 text-muted-foreground" />
                  <Select value={departmentFilter} onValueChange={setDepartmentFilter}>
                    <SelectTrigger className="w-[140px]">
                      <SelectValue placeholder="Departamento" />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value="all">Todos</SelectItem>
                      <SelectItem value="Administración">Administración</SelectItem>
                      <SelectItem value="Ventas">Ventas</SelectItem>
                      <SelectItem value="Finanzas">Finanzas</SelectItem>
                      <SelectItem value="IT">IT</SelectItem>
                      <SelectItem value="Recursos Humanos">RRHH</SelectItem>
                    </SelectContent>
                  </Select>
                  <Select value={roleFilter} onValueChange={setRoleFilter}>
                    <SelectTrigger className="w-[140px]">
                      <SelectValue placeholder="Rol" />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value="all">Todos</SelectItem>
                      <SelectItem value="Super Admin">Super Admin</SelectItem>
                      <SelectItem value="CRM Admin">CRM Admin</SelectItem>
                      <SelectItem value="Sales Manager">Sales Manager</SelectItem>
                      <SelectItem value="Finance User">Finance User</SelectItem>
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
                  <Dialog open={isAssignDialogOpen} onOpenChange={setIsAssignDialogOpen}>
                    <DialogTrigger asChild>
                      <Button>
                        <Plus className="h-4 w-4 mr-2" />
                        Asignar Roles
                      </Button>
                    </DialogTrigger>
                    <DialogContent className="max-w-2xl">
                      <DialogHeader>
                        <DialogTitle>Asignar Roles a Usuario</DialogTitle>
                        <DialogDescription>Selecciona un usuario y los roles que deseas asignar</DialogDescription>
                      </DialogHeader>
                      <div className="space-y-4">
                        <div>
                          <Label>Usuario</Label>
                          <Select value={assignmentForm.userId} onValueChange={(value) => setAssignmentForm({ ...assignmentForm, userId: value })}>
                            <SelectTrigger>
                              <SelectValue placeholder="Seleccionar usuario" />
                            </SelectTrigger>
                            <SelectContent>
                              {userRoles.map((user) => (
                                <SelectItem key={user.id} value={user.id}>
                                  {user.name} - {user.email}
                                </SelectItem>
                              ))}
                            </SelectContent>
                          </Select>
                        </div>
                        <div>
                          <Label>Roles Disponibles</Label>
                          <div className="border rounded-md p-4 max-h-60 overflow-y-auto">
                            <div className="space-y-3">
                              {availableRoles.map((role) => (
                                <div key={role.id} className="flex items-center space-x-3">
                                  <Checkbox
                                    id={role.id}
                                    checked={assignmentForm.roleIds.includes(role.id)}
                                    onCheckedChange={(checked) => {
                                      if (checked) {
                                        setAssignmentForm({
                                          ...assignmentForm,
                                          roleIds: [...assignmentForm.roleIds, role.id],
                                        });
                                      } else {
                                        setAssignmentForm({
                                          ...assignmentForm,
                                          roleIds: assignmentForm.roleIds.filter((id) => id !== role.id),
                                        });
                                      }
                                    }}
                                  />
                                  <div className="flex-1">
                                    <div className="flex items-center space-x-2">
                                      {getTypeIcon(role.type)}
                                      <Label htmlFor={role.id} className="font-medium">
                                        {role.name}
                                      </Label>
                                      <Badge className={getLevelBadge(role.level)}>{role.level}</Badge>
                                    </div>
                                    <p className="text-sm text-muted-foreground">{role.description}</p>
                                  </div>
                                </div>
                              ))}
                            </div>
                          </div>
                        </div>
                        <div className="grid grid-cols-2 gap-4">
                          <div>
                            <Label>Fecha de Expiración (Opcional)</Label>
                            <Input type="datetime-local" value={assignmentForm.expiresAt} onChange={(e) => setAssignmentForm({ ...assignmentForm, expiresAt: e.target.value })} />
                          </div>
                          <div>
                            <Label>Razón de Asignación</Label>
                            <Input
                              placeholder="Motivo de la asignación..."
                              value={assignmentForm.reason}
                              onChange={(e) => setAssignmentForm({ ...assignmentForm, reason: e.target.value })}
                            />
                          </div>
                        </div>
                      </div>
                      <DialogFooter>
                        <Button variant="outline" onClick={() => setIsAssignDialogOpen(false)}>
                          Cancelar
                        </Button>
                        <Button onClick={handleAssignRoles}>Asignar Roles</Button>
                      </DialogFooter>
                    </DialogContent>
                  </Dialog>
                </div>
              </div>

              {/* Lista de usuarios */}
              <div className="space-y-4">
                {filteredUsers.map((user) => (
                  <Card key={user.id}>
                    <CardHeader>
                      <div className="flex items-center justify-between">
                        <div className="flex items-center space-x-3">
                          <Avatar className="h-10 w-10">
                            <AvatarImage src={user.avatar || '/placeholder.svg'} />
                            <AvatarFallback>
                              {user.name
                                .split(' ')
                                .map((n) => n[0])
                                .join('')}
                            </AvatarFallback>
                          </Avatar>
                          <div>
                            <CardTitle className="text-lg">{user.name}</CardTitle>
                            <CardDescription>
                              {user.email} • {user.department} • {user.position}
                            </CardDescription>
                          </div>
                        </div>
                        <div className="flex items-center space-x-2">
                          <Badge className={getRiskBadge(user.riskLevel)}>Riesgo {user.riskLevel === 'low' ? 'Bajo' : user.riskLevel === 'medium' ? 'Medio' : 'Alto'}</Badge>
                          <Badge variant="outline">{user.roles.length} roles</Badge>
                          <Badge variant="outline">{user.totalPermissions} permisos</Badge>
                          <Badge className={user.isActive ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800'}>{user.isActive ? 'Activo' : 'Inactivo'}</Badge>
                          <Collapsible>
                            <CollapsibleTrigger asChild>
                              <Button variant="ghost" size="sm" onClick={() => toggleUserExpansion(user.id)}>
                                {expandedUsers.includes(user.id) ? <ChevronDown className="h-4 w-4" /> : <ChevronRight className="h-4 w-4" />}
                              </Button>
                            </CollapsibleTrigger>
                          </Collapsible>
                        </div>
                      </div>
                    </CardHeader>
                    <Collapsible open={expandedUsers.includes(user.id)}>
                      <CollapsibleContent>
                        <CardContent>
                          <div className="space-y-3">
                            <div className="text-sm text-muted-foreground mb-3">Último acceso: {formatDate(user.lastLogin)}</div>
                            {user.roles.map((role) => (
                              <div key={role.id} className="flex items-center justify-between p-3 border rounded-lg">
                                <div className="flex items-center space-x-3">
                                  {getTypeIcon(role.type)}
                                  <div>
                                    <div className="font-medium">{role.name}</div>
                                    <div className="text-sm text-muted-foreground">
                                      {role.applicationName} • Asignado por {role.assignedBy}
                                    </div>
                                  </div>
                                </div>
                                <div className="flex items-center space-x-2">
                                  <Badge className={getLevelBadge(role.level)}>{role.level}</Badge>
                                  <Badge variant="outline">{role.type}</Badge>
                                  {role.expiresAt && <Badge variant="destructive">Expira: {formatDate(role.expiresAt)}</Badge>}
                                  <Switch checked={role.isActive} />
                                  <Button variant="ghost" size="sm">
                                    <Edit className="h-4 w-4" />
                                  </Button>
                                  <Button variant="ghost" size="sm" className="text-red-600 hover:text-red-700">
                                    <Trash2 className="h-4 w-4" />
                                  </Button>
                                </div>
                              </div>
                            ))}
                          </div>
                        </CardContent>
                      </CollapsibleContent>
                    </Collapsible>
                  </Card>
                ))}
              </div>
            </CardContent>
          </Card>
        </TabsContent>

        <TabsContent value="roles" className="space-y-4">
          <Card>
            <CardHeader>
              <CardTitle>Roles Disponibles</CardTitle>
              <CardDescription>Lista de todos los roles disponibles para asignación</CardDescription>
            </CardHeader>
            <CardContent>
              <div className="rounded-md border">
                <Table>
                  <TableHeader>
                    <TableRow>
                      <TableHead>Rol</TableHead>
                      <TableHead>Tipo</TableHead>
                      <TableHead>Nivel</TableHead>
                      <TableHead>Aplicación</TableHead>
                      <TableHead>Usuarios</TableHead>
                      <TableHead>Permisos</TableHead>
                      <TableHead>Estado</TableHead>
                      <TableHead className="text-right">Acciones</TableHead>
                    </TableRow>
                  </TableHeader>
                  <TableBody>
                    {availableRoles.map((role) => (
                      <TableRow key={role.id}>
                        <TableCell>
                          <div className="flex items-center space-x-2">
                            {getTypeIcon(role.type)}
                            <div>
                              <div className="font-medium">{role.name}</div>
                              <div className="text-sm text-muted-foreground">{role.description}</div>
                            </div>
                          </div>
                        </TableCell>
                        <TableCell>
                          <Badge variant="outline" className="capitalize">
                            {role.type === 'system' ? 'Sistema' : role.type === 'application' ? 'Aplicación' : role.type === 'organizational' ? 'Organizacional' : role.type}
                          </Badge>
                        </TableCell>
                        <TableCell>
                          <Badge className={getLevelBadge(role.level)}>{role.level.charAt(0).toUpperCase() + role.level.slice(1)}</Badge>
                        </TableCell>
                        <TableCell>
                          <Badge variant="outline">{role.applicationName}</Badge>
                        </TableCell>
                        <TableCell>
                          <div className="flex items-center space-x-2">
                            <Users className="h-4 w-4 text-muted-foreground" />
                            <span className="font-medium">{role.userCount}</span>
                          </div>
                        </TableCell>
                        <TableCell>
                          <Badge variant="secondary">{role.permissions.length} permisos</Badge>
                        </TableCell>
                        <TableCell>
                          <Badge className={role.isActive ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800'}>{role.isActive ? 'Activo' : 'Inactivo'}</Badge>
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
                              <UserCheck className="h-4 w-4" />
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

        <TabsContent value="assignments" className="space-y-4">
          <Card>
            <CardHeader>
              <CardTitle>Historial de Asignaciones</CardTitle>
              <CardDescription>Registro de todas las asignaciones y cambios de roles</CardDescription>
            </CardHeader>
            <CardContent>
              <div className="space-y-4">
                {userRoles
                  .flatMap((user) =>
                    user.roles.map((role) => ({
                      ...role,
                      userName: user.name,
                      userEmail: user.email,
                      userAvatar: user.avatar,
                      userDepartment: user.department,
                    })),
                  )
                  .sort((a, b) => new Date(b.assignedAt).getTime() - new Date(a.assignedAt).getTime())
                  .map((assignment) => (
                    <div key={`${assignment.userName}-${assignment.id}`} className="flex items-center justify-between p-4 border rounded-lg">
                      <div className="flex items-center space-x-3">
                        <Avatar className="h-8 w-8">
                          <AvatarImage src={assignment.userAvatar || '/placeholder.svg'} />
                          <AvatarFallback>
                            {assignment.userName
                              .split(' ')
                              .map((n) => n[0])
                              .join('')}
                          </AvatarFallback>
                        </Avatar>
                        <div>
                          <div className="font-medium">
                            {assignment.name} asignado a {assignment.userName}
                          </div>
                          <div className="text-sm text-muted-foreground">
                            {assignment.userEmail} • {assignment.userDepartment} • {formatDate(assignment.assignedAt)}
                          </div>
                        </div>
                      </div>
                      <div className="flex items-center space-x-2">
                        <Badge className={getLevelBadge(assignment.level)}>{assignment.level}</Badge>
                        <Badge variant="outline">{assignment.applicationName}</Badge>
                        {assignment.expiresAt && <Badge variant="destructive">Expira: {formatDate(assignment.expiresAt)}</Badge>}
                        <div className="text-sm text-muted-foreground">por {assignment.assignedBy}</div>
                      </div>
                    </div>
                  ))}
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
                    <div key={level} className="flex items-center justify-between">
                      <div className="flex items-center space-x-2">
                        <div className={`w-3 h-3 rounded-full ${level === 'admin' ? 'bg-red-500' : level === 'manager' ? 'bg-orange-500' : 'bg-blue-500'}`}></div>
                        <span className="capitalize">{level}</span>
                      </div>
                      <div className="flex items-center space-x-2">
                        <div className="w-24 bg-gray-200 rounded-full h-2">
                          <div
                            className={`h-2 rounded-full ${level === 'admin' ? 'bg-red-500' : level === 'manager' ? 'bg-orange-500' : 'bg-blue-500'}`}
                            style={{ width: `${(count / roleStats.totalUsers) * 100}%` }}
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
                <CardTitle>Distribución por Tipo</CardTitle>
              </CardHeader>
              <CardContent>
                <div className="space-y-4">
                  {Object.entries(roleStats.byType).map(([type, count]) => (
                    <div key={type} className="flex items-center justify-between">
                      <div className="flex items-center space-x-2">
                        {getTypeIcon(type)}
                        <span className="capitalize">
                          {type === 'system' ? 'Sistema' : type === 'application' ? 'Aplicación' : type === 'organizational' ? 'Organizacional' : type}
                        </span>
                      </div>
                      <div className="flex items-center space-x-2">
                        <div className="w-24 bg-gray-200 rounded-full h-2">
                          <div className="bg-blue-600 h-2 rounded-full" style={{ width: `${(count / roleStats.totalUsers) * 100}%` }}></div>
                        </div>
                        <span className="text-sm font-medium">{count}</span>
                      </div>
                    </div>
                  ))}
                </div>
              </CardContent>
            </Card>
          </div>
        </TabsContent>
      </Tabs>
    </div>
  );
}
