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
import { Textarea } from '@/components/ui/textarea';
import { Switch } from '@/components/ui/switch';
import { Search, Filter, Plus, Edit, Trash2, Ban, Clock, MapPin, Shield, AlertTriangle, Eye, Calendar, Users, Lock, Unlock } from 'lucide-react';
import CardStatsContain from '@/components/custom/card/card-stats-contain';
import { statsRestrictions } from '@/mocks/stats-mocks';

// Datos simulados de restricciones de usuario
const userRestrictions = [
  {
    id: 'rest_001',
    userId: 'user_001',
    userName: 'Ana García',
    userEmail: 'ana.garcia@empresa.com',
    avatar: '/placeholder.svg?height=32&width=32',
    restrictions: [
      {
        id: 'time_001',
        type: 'time',
        name: 'Horario Laboral',
        description: 'Acceso permitido solo en horario laboral',
        isActive: true,
        severity: 'medium',
        startTime: '08:00',
        endTime: '18:00',
        timezone: 'Europe/Madrid',
        weekdays: ['monday', 'tuesday', 'wednesday', 'thursday', 'friday'],
        createdAt: '2024-01-10T10:00:00Z',
        createdBy: 'Sistema',
      },
      {
        id: 'ip_001',
        type: 'ip',
        name: 'Red Corporativa',
        description: 'Acceso solo desde IPs corporativas',
        isActive: true,
        severity: 'high',
        allowedIPs: ['192.168.1.0/24', '10.0.0.0/16'],
        blockedIPs: [],
        createdAt: '2024-01-10T10:15:00Z',
        createdBy: 'Admin',
      },
    ],
    riskLevel: 'medium',
    lastViolation: '2024-01-14T15:30:00Z',
    violationCount: 2,
    isActive: true,
  },
  {
    id: 'rest_002',
    userId: 'user_002',
    userName: 'Carlos López',
    userEmail: 'carlos.lopez@empresa.com',
    avatar: '/placeholder.svg?height=32&width=32',
    restrictions: [
      {
        id: 'geo_001',
        type: 'location',
        name: 'Restricción Geográfica',
        description: 'Acceso bloqueado desde ciertos países',
        isActive: true,
        severity: 'high',
        allowedCountries: ['ES', 'FR', 'DE'],
        blockedCountries: ['CN', 'RU'],
        allowedCities: [],
        blockedCities: [],
        createdAt: '2024-01-12T14:20:00Z',
        createdBy: 'Security Team',
      },
      {
        id: 'device_001',
        type: 'device',
        name: 'Dispositivos Autorizados',
        description: 'Solo dispositivos registrados',
        isActive: true,
        severity: 'medium',
        allowedDevices: ['desktop-001', 'mobile-002'],
        maxDevices: 3,
        requireDeviceRegistration: true,
        createdAt: '2024-01-12T14:25:00Z',
        createdBy: 'IT Admin',
      },
    ],
    riskLevel: 'high',
    lastViolation: '2024-01-15T09:45:00Z',
    violationCount: 5,
    isActive: true,
  },
  {
    id: 'rest_003',
    userId: 'user_003',
    userName: 'María Rodríguez',
    userEmail: 'maria.rodriguez@empresa.com',
    avatar: '/placeholder.svg?height=32&width=32',
    restrictions: [
      {
        id: 'session_001',
        type: 'session',
        name: 'Límite de Sesiones',
        description: 'Máximo 2 sesiones concurrentes',
        isActive: true,
        severity: 'medium',
        maxConcurrentSessions: 2,
        sessionTimeout: 3600, // 1 hora en segundos
        idleTimeout: 1800, // 30 minutos en segundos
        createdAt: '2024-01-08T11:30:00Z',
        createdBy: 'Security Admin',
      },
      {
        id: 'app_001',
        type: 'application',
        name: 'Aplicaciones Restringidas',
        description: 'Acceso limitado a ciertas aplicaciones',
        isActive: true,
        severity: 'low',
        allowedApplications: ['CRM System', 'Email'],
        blockedApplications: ['Admin Panel', 'Finance Module'],
        timeRestrictions: {
          'Finance Module': {
            startTime: '09:00',
            endTime: '17:00',
            weekdays: ['monday', 'tuesday', 'wednesday', 'thursday', 'friday'],
          },
        },
        createdAt: '2024-01-08T11:35:00Z',
        createdBy: 'Department Manager',
      },
    ],
    riskLevel: 'low',
    lastViolation: null,
    violationCount: 0,
    isActive: true,
  },
  {
    id: 'rest_004',
    userId: 'user_004',
    userName: 'David Martín',
    userEmail: 'david.martin@empresa.com',
    avatar: '/placeholder.svg?height=32&width=32',
    restrictions: [
      {
        id: 'temp_001',
        type: 'temporary',
        name: 'Suspensión Temporal',
        description: 'Acceso suspendido por violación de políticas',
        isActive: true,
        severity: 'high',
        startDate: '2024-01-15T00:00:00Z',
        endDate: '2024-01-22T23:59:59Z',
        reason: 'Múltiples intentos de acceso no autorizado',
        canAppeal: true,
        appealDeadline: '2024-01-20T23:59:59Z',
        createdAt: '2024-01-15T08:00:00Z',
        createdBy: 'Security Officer',
      },
    ],
    riskLevel: 'high',
    lastViolation: '2024-01-14T22:15:00Z',
    violationCount: 8,
    isActive: false,
  },
];

const restrictionStats = {
  total: 89,
  active: 76,
  inactive: 13,
  byType: {
    time: 23,
    ip: 18,
    location: 15,
    device: 12,
    session: 11,
    application: 8,
    temporary: 2,
  },
  bySeverity: {
    low: 34,
    medium: 38,
    high: 17,
  },
  violations: {
    today: 12,
    thisWeek: 45,
    thisMonth: 156,
  },
};

export default function UserRestrictionsManagement() {
  const [searchTerm, setSearchTerm] = useState('');
  const [typeFilter, setTypeFilter] = useState('all');
  const [severityFilter, setSeverityFilter] = useState('all');
  const [statusFilter, setStatusFilter] = useState('all');
  const [isCreateDialogOpen, setIsCreateDialogOpen] = useState(false);

  const getSeverityBadge = (severity: string) => {
    const colors = {
      low: 'bg-green-100 text-green-800',
      medium: 'bg-yellow-100 text-yellow-800',
      high: 'bg-red-100 text-red-800',
    };
    return colors[severity as keyof typeof colors] || 'bg-gray-100 text-gray-800';
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
      case 'time':
        return <Clock className="h-4 w-4" />;
      case 'ip':
        return <Shield className="h-4 w-4" />;
      case 'location':
        return <MapPin className="h-4 w-4" />;
      case 'device':
        return <Users className="h-4 w-4" />;
      case 'session':
        return <Lock className="h-4 w-4" />;
      case 'application':
        return <Ban className="h-4 w-4" />;
      case 'temporary':
        return <Calendar className="h-4 w-4" />;
      default:
        return <AlertTriangle className="h-4 w-4" />;
    }
  };

  const filteredUsers = userRestrictions.filter((user) => {
    const matchesSearch = user.userName.toLowerCase().includes(searchTerm.toLowerCase()) || user.userEmail.toLowerCase().includes(searchTerm.toLowerCase());
    const matchesStatus = statusFilter === 'all' || (statusFilter === 'active' ? user.isActive : !user.isActive);

    return matchesSearch && matchesStatus;
  });

  const getAllRestrictions = () => {
    return userRestrictions.flatMap((user) =>
      user.restrictions.map((restriction) => ({
        ...restriction,
        userName: user.userName,
        userEmail: user.userEmail,
        userAvatar: user.avatar,
        userRisk: user.riskLevel,
      })),
    );
  };

  const filteredRestrictions = getAllRestrictions().filter((restriction) => {
    const matchesSearch =
      restriction.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
      restriction.description.toLowerCase().includes(searchTerm.toLowerCase()) ||
      restriction.userName.toLowerCase().includes(searchTerm.toLowerCase());
    const matchesType = typeFilter === 'all' || restriction.type === typeFilter;
    const matchesSeverity = severityFilter === 'all' || restriction.severity === severityFilter;

    return matchesSearch && matchesType && matchesSeverity;
  });

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
      <CardStatsContain stats={statsRestrictions} />

      <Tabs defaultValue="users" className="space-y-4">
        <TabsList>
          <TabsTrigger value="users">Por Usuario</TabsTrigger>
          <TabsTrigger value="restrictions">Todas las Restricciones</TabsTrigger>
          <TabsTrigger value="violations">Violaciones</TabsTrigger>
          <TabsTrigger value="analytics">Análisis</TabsTrigger>
        </TabsList>

        <TabsContent value="users" className="space-y-4">
          <Card>
            <CardHeader>
              <CardTitle>Restricciones por Usuario</CardTitle>
              <CardDescription>Vista de usuarios con sus restricciones aplicadas</CardDescription>
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
                  <Dialog open={isCreateDialogOpen} onOpenChange={setIsCreateDialogOpen}>
                    <DialogTrigger asChild>
                      <Button>
                        <Plus className="h-4 w-4 mr-2" />
                        Nueva Restricción
                      </Button>
                    </DialogTrigger>
                    <DialogContent className="max-w-2xl">
                      <DialogHeader>
                        <DialogTitle>Crear Nueva Restricción</DialogTitle>
                        <DialogDescription>Define una nueva restricción para un usuario específico</DialogDescription>
                      </DialogHeader>
                      <div className="space-y-4">
                        <div className="grid grid-cols-2 gap-4">
                          <div>
                            <Label>Usuario</Label>
                            <Select>
                              <SelectTrigger>
                                <SelectValue placeholder="Seleccionar usuario" />
                              </SelectTrigger>
                              <SelectContent>
                                <SelectItem value="user_001">Ana García</SelectItem>
                                <SelectItem value="user_002">Carlos López</SelectItem>
                                <SelectItem value="user_003">María Rodríguez</SelectItem>
                              </SelectContent>
                            </Select>
                          </div>
                          <div>
                            <Label>Tipo de Restricción</Label>
                            <Select>
                              <SelectTrigger>
                                <SelectValue placeholder="Seleccionar tipo" />
                              </SelectTrigger>
                              <SelectContent>
                                <SelectItem value="time">Horario</SelectItem>
                                <SelectItem value="ip">IP/Red</SelectItem>
                                <SelectItem value="location">Ubicación</SelectItem>
                                <SelectItem value="device">Dispositivo</SelectItem>
                                <SelectItem value="session">Sesión</SelectItem>
                                <SelectItem value="application">Aplicación</SelectItem>
                                <SelectItem value="temporary">Temporal</SelectItem>
                              </SelectContent>
                            </Select>
                          </div>
                        </div>
                        <div>
                          <Label>Nombre</Label>
                          <Input placeholder="Nombre de la restricción" />
                        </div>
                        <div>
                          <Label>Descripción</Label>
                          <Textarea placeholder="Describe el propósito de esta restricción..." />
                        </div>
                        <div className="grid grid-cols-2 gap-4">
                          <div>
                            <Label>Severidad</Label>
                            <Select>
                              <SelectTrigger>
                                <SelectValue placeholder="Seleccionar severidad" />
                              </SelectTrigger>
                              <SelectContent>
                                <SelectItem value="low">Baja</SelectItem>
                                <SelectItem value="medium">Media</SelectItem>
                                <SelectItem value="high">Alta</SelectItem>
                              </SelectContent>
                            </Select>
                          </div>
                          <div className="flex items-center space-x-2">
                            <Switch id="active" />
                            <Label htmlFor="active">Activar inmediatamente</Label>
                          </div>
                        </div>
                      </div>
                      <DialogFooter>
                        <Button variant="outline" onClick={() => setIsCreateDialogOpen(false)}>
                          Cancelar
                        </Button>
                        <Button onClick={() => setIsCreateDialogOpen(false)}>Crear Restricción</Button>
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
                              {user.userName
                                .split(' ')
                                .map((n) => n[0])
                                .join('')}
                            </AvatarFallback>
                          </Avatar>
                          <div>
                            <CardTitle className="text-lg">{user.userName}</CardTitle>
                            <CardDescription>{user.userEmail}</CardDescription>
                          </div>
                        </div>
                        <div className="flex items-center space-x-2">
                          <Badge className={getRiskBadge(user.riskLevel)}>Riesgo {user.riskLevel === 'low' ? 'Bajo' : user.riskLevel === 'medium' ? 'Medio' : 'Alto'}</Badge>
                          <Badge variant="outline">{user.restrictions.length} restricciones</Badge>
                          {user.violationCount > 0 && <Badge variant="destructive">{user.violationCount} violaciones</Badge>}
                          <Badge className={user.isActive ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800'}>{user.isActive ? 'Activo' : 'Suspendido'}</Badge>
                        </div>
                      </div>
                    </CardHeader>
                    <CardContent>
                      <div className="grid gap-3">
                        {user.restrictions.map((restriction) => (
                          <div key={restriction.id} className="flex items-center justify-between p-3 border rounded-lg">
                            <div className="flex items-center space-x-3">
                              {getTypeIcon(restriction.type)}
                              <div>
                                <div className="font-medium">{restriction.name}</div>
                                <div className="text-sm text-muted-foreground">{restriction.description}</div>
                              </div>
                            </div>
                            <div className="flex items-center space-x-2">
                              <Badge className={getSeverityBadge(restriction.severity)}>
                                {restriction.severity === 'low' ? 'Baja' : restriction.severity === 'medium' ? 'Media' : 'Alta'}
                              </Badge>
                              <Switch checked={restriction.isActive} />
                              <Button variant="ghost" size="sm">
                                <Edit className="h-4 w-4" />
                              </Button>
                            </div>
                          </div>
                        ))}
                      </div>
                    </CardContent>
                  </Card>
                ))}
              </div>
            </CardContent>
          </Card>
        </TabsContent>

        <TabsContent value="restrictions" className="space-y-4">
          <Card>
            <CardHeader>
              <CardTitle>Todas las Restricciones</CardTitle>
              <CardDescription>Lista completa de todas las restricciones del sistema</CardDescription>
            </CardHeader>
            <CardContent>
              {/* Filtros */}
              <div className="flex flex-col gap-4 md:flex-row md:items-center md:justify-between mb-6">
                <div className="flex flex-1 items-center space-x-2">
                  <Search className="h-4 w-4 text-muted-foreground" />
                  <Input placeholder="Buscar restricciones..." value={searchTerm} onChange={(e) => setSearchTerm(e.target.value)} className="max-w-sm" />
                </div>
                <div className="flex items-center space-x-2">
                  <Filter className="h-4 w-4 text-muted-foreground" />
                  <Select value={typeFilter} onValueChange={setTypeFilter}>
                    <SelectTrigger className="w-[140px]">
                      <SelectValue placeholder="Tipo" />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value="all">Todos</SelectItem>
                      <SelectItem value="time">Horario</SelectItem>
                      <SelectItem value="ip">IP/Red</SelectItem>
                      <SelectItem value="location">Ubicación</SelectItem>
                      <SelectItem value="device">Dispositivo</SelectItem>
                      <SelectItem value="session">Sesión</SelectItem>
                      <SelectItem value="application">Aplicación</SelectItem>
                      <SelectItem value="temporary">Temporal</SelectItem>
                    </SelectContent>
                  </Select>
                  <Select value={severityFilter} onValueChange={setSeverityFilter}>
                    <SelectTrigger className="w-[140px]">
                      <SelectValue placeholder="Severidad" />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value="all">Todas</SelectItem>
                      <SelectItem value="low">Baja</SelectItem>
                      <SelectItem value="medium">Media</SelectItem>
                      <SelectItem value="high">Alta</SelectItem>
                    </SelectContent>
                  </Select>
                </div>
              </div>

              {/* Tabla de restricciones */}
              <div className="rounded-md border">
                <Table>
                  <TableHeader>
                    <TableRow>
                      <TableHead>Restricción</TableHead>
                      <TableHead>Usuario</TableHead>
                      <TableHead>Tipo</TableHead>
                      <TableHead>Severidad</TableHead>
                      <TableHead>Creado</TableHead>
                      <TableHead>Estado</TableHead>
                      <TableHead className="text-right">Acciones</TableHead>
                    </TableRow>
                  </TableHeader>
                  <TableBody>
                    {filteredRestrictions.map((restriction) => (
                      <TableRow key={restriction.id}>
                        <TableCell>
                          <div className="flex items-center space-x-2">
                            {getTypeIcon(restriction.type)}
                            <div>
                              <div className="font-medium">{restriction.name}</div>
                              <div className="text-sm text-muted-foreground">{restriction.description}</div>
                            </div>
                          </div>
                        </TableCell>
                        <TableCell>
                          <div className="flex items-center space-x-2">
                            <Avatar className="h-6 w-6">
                              <AvatarImage src={restriction.userAvatar || '/placeholder.svg'} />
                              <AvatarFallback className="text-xs">
                                {restriction.userName
                                  .split(' ')
                                  .map((n) => n[0])
                                  .join('')}
                              </AvatarFallback>
                            </Avatar>
                            <div>
                              <div className="font-medium text-sm">{restriction.userName}</div>
                              <div className="text-xs text-muted-foreground">{restriction.userEmail}</div>
                            </div>
                          </div>
                        </TableCell>
                        <TableCell>
                          <Badge variant="outline" className="capitalize">
                            {restriction.type === 'time'
                              ? 'Horario'
                              : restriction.type === 'ip'
                              ? 'IP/Red'
                              : restriction.type === 'location'
                              ? 'Ubicación'
                              : restriction.type === 'device'
                              ? 'Dispositivo'
                              : restriction.type === 'session'
                              ? 'Sesión'
                              : restriction.type === 'application'
                              ? 'Aplicación'
                              : restriction.type === 'temporary'
                              ? 'Temporal'
                              : restriction.type}
                          </Badge>
                        </TableCell>
                        <TableCell>
                          <Badge className={getSeverityBadge(restriction.severity)}>
                            {restriction.severity === 'low' ? 'Baja' : restriction.severity === 'medium' ? 'Media' : 'Alta'}
                          </Badge>
                        </TableCell>
                        <TableCell>
                          <div className="text-sm">{formatDate(restriction.createdAt)}</div>
                          <div className="text-xs text-muted-foreground">por {restriction.createdBy}</div>
                        </TableCell>
                        <TableCell>
                          <Switch checked={restriction.isActive} />
                        </TableCell>
                        <TableCell className="text-right">
                          <div className="flex items-center justify-end space-x-2">
                            <Button variant="ghost" size="sm">
                              <Eye className="h-4 w-4" />
                            </Button>
                            <Button variant="ghost" size="sm">
                              <Edit className="h-4 w-4" />
                            </Button>
                            <Button variant="ghost" size="sm" className="text-red-600 hover:text-red-700">
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

        <TabsContent value="violations" className="space-y-4">
          <Card>
            <CardHeader>
              <CardTitle>Violaciones de Restricciones</CardTitle>
              <CardDescription>Registro de violaciones y intentos de acceso no autorizados</CardDescription>
            </CardHeader>
            <CardContent>
              <div className="space-y-4">
                {userRestrictions
                  .filter((user) => user.violationCount > 0)
                  .sort((a, b) => b.violationCount - a.violationCount)
                  .map((user) => (
                    <div key={user.id} className="flex items-center justify-between p-4 border rounded-lg bg-red-50">
                      <div className="flex items-center space-x-3">
                        <AlertTriangle className="h-5 w-5 text-red-500" />
                        <Avatar className="h-8 w-8">
                          <AvatarImage src={user.avatar || '/placeholder.svg'} />
                          <AvatarFallback>
                            {user.userName
                              .split(' ')
                              .map((n) => n[0])
                              .join('')}
                          </AvatarFallback>
                        </Avatar>
                        <div>
                          <div className="font-medium">{user.userName}</div>
                          <div className="text-sm text-muted-foreground">
                            {user.violationCount} violaciones - Última: {user.lastViolation ? formatDate(user.lastViolation) : 'N/A'}
                          </div>
                        </div>
                      </div>
                      <div className="flex items-center space-x-2">
                        <Badge className={getRiskBadge(user.riskLevel)}>{user.riskLevel === 'low' ? 'Bajo' : user.riskLevel === 'medium' ? 'Medio' : 'Alto'}</Badge>
                        <Button variant="outline" size="sm">
                          <Eye className="h-4 w-4 mr-2" />
                          Ver Detalles
                        </Button>
                        {user.isActive ? (
                          <Button variant="destructive" size="sm">
                            <Ban className="h-4 w-4 mr-2" />
                            Suspender
                          </Button>
                        ) : (
                          <Button variant="outline" size="sm">
                            <Unlock className="h-4 w-4 mr-2" />
                            Reactivar
                          </Button>
                        )}
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
                <CardTitle>Restricciones por Tipo</CardTitle>
              </CardHeader>
              <CardContent>
                <div className="space-y-4">
                  {Object.entries(restrictionStats.byType).map(([type, count]) => (
                    <div key={type} className="flex items-center justify-between">
                      <div className="flex items-center space-x-2">
                        {getTypeIcon(type)}
                        <span className="capitalize">
                          {type === 'time'
                            ? 'Horario'
                            : type === 'ip'
                            ? 'IP/Red'
                            : type === 'location'
                            ? 'Ubicación'
                            : type === 'device'
                            ? 'Dispositivo'
                            : type === 'session'
                            ? 'Sesión'
                            : type === 'application'
                            ? 'Aplicación'
                            : type === 'temporary'
                            ? 'Temporal'
                            : type}
                        </span>
                      </div>
                      <div className="flex items-center space-x-2">
                        <div className="w-24 bg-gray-200 rounded-full h-2">
                          <div className="bg-blue-600 h-2 rounded-full" style={{ width: `${(count / restrictionStats.total) * 100}%` }}></div>
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
                <CardTitle>Distribución por Severidad</CardTitle>
              </CardHeader>
              <CardContent>
                <div className="space-y-4">
                  {Object.entries(restrictionStats.bySeverity).map(([severity, count]) => (
                    <div key={severity} className="flex items-center justify-between">
                      <div className="flex items-center space-x-2">
                        <div className={`w-3 h-3 rounded-full ${severity === 'low' ? 'bg-green-500' : severity === 'medium' ? 'bg-yellow-500' : 'bg-red-500'}`}></div>
                        <span className="capitalize">{severity === 'low' ? 'Baja' : severity === 'medium' ? 'Media' : 'Alta'}</span>
                      </div>
                      <div className="flex items-center space-x-2">
                        <div className="w-24 bg-gray-200 rounded-full h-2">
                          <div
                            className={`h-2 rounded-full ${severity === 'low' ? 'bg-green-500' : severity === 'medium' ? 'bg-yellow-500' : 'bg-red-500'}`}
                            style={{ width: `${(count / restrictionStats.total) * 100}%` }}
                          ></div>
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
