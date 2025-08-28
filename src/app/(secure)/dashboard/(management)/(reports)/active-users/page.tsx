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
import { Search, Download, TrendingUp, Clock, Eye, RefreshCw } from 'lucide-react';
import CardStatsContain from '@/components/custom/card/card-stats-contain';
import { statsActiveUsers } from '@/mocks/stats-mocks';
// Mock data
const activeUsersData = [
  {
    id: '1',
    user: {
      name: 'Juan Carlos Pérez',
      email: 'jperez@empresa.com',
      avatar: '/placeholder.svg?height=40&width=40',
      position: 'Gerente General',
      department: 'Gerencia',
    },
    lastActivity: '2024-01-15 16:45:00',
    sessionsToday: 3,
    totalSessions: 245,
    applicationsUsed: [
      { name: 'Sistema de Inventario', lastAccess: '2024-01-15 16:30:00', sessions: 2 },
      { name: 'Portal RRHH', lastAccess: '2024-01-15 14:20:00', sessions: 1 },
    ],
    deviceInfo: {
      type: 'Desktop',
      browser: 'Chrome 120.0',
      os: 'Windows 11',
      ip: '192.168.1.100',
    },
    activityScore: 95,
    status: 'online',
  },
  {
    id: '2',
    user: {
      name: 'María García López',
      email: 'mgarcia@empresa.com',
      avatar: '/placeholder.svg?height=40&width=40',
      position: 'Analista Senior',
      department: 'Sistemas',
    },
    lastActivity: '2024-01-15 16:20:00',
    sessionsToday: 2,
    totalSessions: 189,
    applicationsUsed: [
      { name: 'Sistema de Inventario', lastAccess: '2024-01-15 16:20:00', sessions: 1 },
      { name: 'App Móvil Ventas', lastAccess: '2024-01-15 15:10:00', sessions: 1 },
    ],
    deviceInfo: {
      type: 'Mobile',
      browser: 'Safari 17.2',
      os: 'iOS 17.2',
      ip: '192.168.1.105',
    },
    activityScore: 88,
    status: 'online',
  },
  {
    id: '3',
    user: {
      name: 'Carlos López Ruiz',
      email: 'clopez@empresa.com',
      avatar: '/placeholder.svg?height=40&width=40',
      position: 'Desarrollador',
      department: 'Sistemas',
    },
    lastActivity: '2024-01-15 15:45:00',
    sessionsToday: 4,
    totalSessions: 312,
    applicationsUsed: [
      { name: 'Sistema de Inventario', lastAccess: '2024-01-15 15:45:00', sessions: 2 },
      { name: 'Portal RRHH', lastAccess: '2024-01-15 13:30:00', sessions: 1 },
      { name: 'Sistema Financiero', lastAccess: '2024-01-15 11:15:00', sessions: 1 },
    ],
    deviceInfo: {
      type: 'Desktop',
      browser: 'Firefox 121.0',
      os: 'Ubuntu 22.04',
      ip: '192.168.1.110',
    },
    activityScore: 92,
    status: 'away',
  },
  {
    id: '4',
    user: {
      name: 'Ana Martínez Silva',
      email: 'amartinez@empresa.com',
      avatar: '/placeholder.svg?height=40&width=40',
      position: 'Especialista TI',
      department: 'Infraestructura',
    },
    lastActivity: '2024-01-15 14:30:00',
    sessionsToday: 1,
    totalSessions: 156,
    applicationsUsed: [{ name: 'Portal RRHH', lastAccess: '2024-01-15 14:30:00', sessions: 1 }],
    deviceInfo: {
      type: 'Desktop',
      browser: 'Edge 120.0',
      os: 'Windows 10',
      ip: '192.168.1.115',
    },
    activityScore: 75,
    status: 'idle',
  },
  {
    id: '5',
    user: {
      name: 'Luis Fernando Torres',
      email: 'ltorres@empresa.com',
      avatar: '/placeholder.svg?height=40&width=40',
      position: 'Gerente RRHH',
      department: 'Recursos Humanos',
    },
    lastActivity: '2024-01-15 12:15:00',
    sessionsToday: 2,
    totalSessions: 203,
    applicationsUsed: [{ name: 'Portal RRHH', lastAccess: '2024-01-15 12:15:00', sessions: 2 }],
    deviceInfo: {
      type: 'Tablet',
      browser: 'Safari 17.1',
      os: 'iPadOS 17.1',
      ip: '192.168.1.120',
    },
    activityScore: 82,
    status: 'offline',
  },
];

const activityStats = {
  totalActiveUsers: 1234,
  onlineNow: 89,
  peakHour: '14:00 - 15:00',
  averageSessionTime: '2h 45m',
  dailyGrowth: 5.2,
  weeklyGrowth: 12.8,
  monthlyGrowth: 23.4,
};

const hourlyActivity = [
  { hour: '00:00', users: 12 },
  { hour: '01:00', users: 8 },
  { hour: '02:00', users: 5 },
  { hour: '03:00', users: 3 },
  { hour: '04:00', users: 2 },
  { hour: '05:00', users: 4 },
  { hour: '06:00', users: 15 },
  { hour: '07:00', users: 35 },
  { hour: '08:00', users: 68 },
  { hour: '09:00', users: 89 },
  { hour: '10:00', users: 95 },
  { hour: '11:00', users: 102 },
  { hour: '12:00', users: 78 },
  { hour: '13:00', users: 85 },
  { hour: '14:00', users: 112 },
  { hour: '15:00', users: 98 },
  { hour: '16:00', users: 89 },
  { hour: '17:00', users: 65 },
  { hour: '18:00', users: 45 },
  { hour: '19:00', users: 32 },
  { hour: '20:00', users: 25 },
  { hour: '21:00', users: 18 },
  { hour: '22:00', users: 15 },
  { hour: '23:00', users: 12 },
];

export default function ActiveUsersReport() {
  const [searchTerm, setSearchTerm] = useState('');
  const [statusFilter, setStatusFilter] = useState('all');
  const [departmentFilter, setDepartmentFilter] = useState('all');

  const filteredUsers = activeUsersData.filter((user) => {
    const matchesSearch =
      user.user.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
      user.user.email.toLowerCase().includes(searchTerm.toLowerCase()) ||
      user.user.department.toLowerCase().includes(searchTerm.toLowerCase());

    const matchesStatus = statusFilter === 'all' || user.status === statusFilter;
    const matchesDepartment = departmentFilter === 'all' || user.user.department === departmentFilter;

    return matchesSearch && matchesStatus && matchesDepartment;
  });

  const getStatusBadge = (status: string) => {
    switch (status) {
      case 'online':
        return (
          <Badge className="bg-chart-4/20 text-chart-4 border-chart-4/30">
            <div className="w-2 h-2 bg-chart-4 rounded-full mr-1 animate-pulse" />
            En línea
          </Badge>
        );
      case 'away':
        return (
          <Badge className="bg-chart-5/20 text-chart-5 border-chart-5/30">
            <div className="w-2 h-2 bg-chart-5 rounded-full mr-1" />
            Ausente
          </Badge>
        );
      case 'idle':
        return (
          <Badge className="bg-chart-3/20 text-chart-3 border-chart-3/30">
            <div className="w-2 h-2 bg-chart-3 rounded-full mr-1" />
            Inactivo
          </Badge>
        );
      case 'offline':
        return (
          <Badge className="bg-muted text-muted-foreground border-muted-foreground/30">
            <div className="w-2 h-2 bg-muted-foreground rounded-full mr-1" />
            Desconectado
          </Badge>
        );
      default:
        return <Badge variant="secondary">Desconocido</Badge>;
    }
  };

  const getActivityScoreColor = (score: number) => {
    if (score >= 90) return 'text-chart-4';
    if (score >= 75) return 'text-chart-5';
    if (score >= 60) return 'text-chart-3';
    return 'text-muted-foreground';
  };

  const formatDateTime = (dateString: string) => {
    return new Date(dateString).toLocaleString('es-ES');
  };

  const getDeviceIcon = (type: string) => {
    switch (type.toLowerCase()) {
      case 'desktop':
        return '🖥️';
      case 'mobile':
        return '📱';
      case 'tablet':
        return '📱';
      default:
        return '💻';
    }
  };

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold text-foreground">Usuarios Activos</h1>
          <p className="text-muted-foreground mt-1">Monitoreo en tiempo real de la actividad de usuarios</p>
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
      <CardStatsContain stats={statsActiveUsers} />

      <Tabs defaultValue="users" className="w-full">
        <TabsList className="grid w-full grid-cols-3">
          <TabsTrigger value="users">Lista de Usuarios</TabsTrigger>
          <TabsTrigger value="activity">Actividad por Hora</TabsTrigger>
          <TabsTrigger value="analytics">Analíticas</TabsTrigger>
        </TabsList>

        <TabsContent value="users" className="space-y-4">
          {/* Filters */}
          <Card className="border-border bg-card/50">
            <CardHeader>
              <div className="flex items-center justify-between">
                <div>
                  <CardTitle>Usuarios Activos en Tiempo Real</CardTitle>
                  <CardDescription>
                    {filteredUsers.length} de {activeUsersData.length} usuarios mostrados
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
                  <Select value={statusFilter} onValueChange={setStatusFilter}>
                    <SelectTrigger className="w-40">
                      <SelectValue placeholder="Estado" />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value="all">Todos los estados</SelectItem>
                      <SelectItem value="online">En línea</SelectItem>
                      <SelectItem value="away">Ausente</SelectItem>
                      <SelectItem value="idle">Inactivo</SelectItem>
                      <SelectItem value="offline">Desconectado</SelectItem>
                    </SelectContent>
                  </Select>
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
                </div>
              </div>
            </CardHeader>
            <CardContent>
              <div className="rounded-lg border border-border">
                <Table>
                  <TableHeader>
                    <TableRow className="bg-accent/50">
                      <TableHead>Usuario</TableHead>
                      <TableHead>Estado</TableHead>
                      <TableHead>Última Actividad</TableHead>
                      <TableHead>Sesiones Hoy</TableHead>
                      <TableHead>Aplicaciones</TableHead>
                      <TableHead>Dispositivo</TableHead>
                      <TableHead>Puntuación</TableHead>
                      <TableHead className="text-right">Acciones</TableHead>
                    </TableRow>
                  </TableHeader>
                  <TableBody>
                    {filteredUsers.map((user) => (
                      <TableRow key={user.id} className="hover:bg-accent/30">
                        <TableCell>
                          <div className="flex items-center gap-3">
                            <Avatar className="w-10 h-10">
                              <AvatarImage src={user.user.avatar || '/placeholder.svg'} />
                              <AvatarFallback className="bg-gradient-to-r from-primary to-chart-1 text-primary-foreground font-semibold">
                                {user.user.name
                                  .split(' ')
                                  .map((n) => n[0])
                                  .join('')}
                              </AvatarFallback>
                            </Avatar>
                            <div>
                              <p className="font-medium text-foreground">{user.user.name}</p>
                              <p className="text-sm text-muted-foreground">{user.user.email}</p>
                              <p className="text-xs text-muted-foreground">
                                {user.user.position} - {user.user.department}
                              </p>
                            </div>
                          </div>
                        </TableCell>
                        <TableCell>{getStatusBadge(user.status)}</TableCell>
                        <TableCell>
                          <div className="space-y-1">
                            <p className="text-sm font-medium">{formatDateTime(user.lastActivity)}</p>
                            <div className="flex items-center gap-1 text-xs text-muted-foreground">
                              <Clock className="w-3 h-3" />
                              <span>IP: {user.deviceInfo.ip}</span>
                            </div>
                          </div>
                        </TableCell>
                        <TableCell>
                          <div className="text-center">
                            <p className="text-lg font-bold text-primary">{user.sessionsToday}</p>
                            <p className="text-xs text-muted-foreground">de {user.totalSessions} total</p>
                          </div>
                        </TableCell>
                        <TableCell>
                          <div className="space-y-1">
                            {user.applicationsUsed.slice(0, 2).map((app, index) => (
                              <div key={index} className="flex items-center gap-2">
                                <div className="w-2 h-2 bg-chart-4 rounded-full" />
                                <span className="text-sm">{app.name}</span>
                                <Badge variant="outline" className="text-xs">
                                  {app.sessions}
                                </Badge>
                              </div>
                            ))}
                            {user.applicationsUsed.length > 2 && <p className="text-xs text-muted-foreground">+{user.applicationsUsed.length - 2} más</p>}
                          </div>
                        </TableCell>
                        <TableCell>
                          <div className="space-y-1">
                            <div className="flex items-center gap-2">
                              <span className="text-lg">{getDeviceIcon(user.deviceInfo.type)}</span>
                              <span className="text-sm font-medium">{user.deviceInfo.type}</span>
                            </div>
                            <p className="text-xs text-muted-foreground">{user.deviceInfo.browser}</p>
                            <p className="text-xs text-muted-foreground">{user.deviceInfo.os}</p>
                          </div>
                        </TableCell>
                        <TableCell>
                          <div className="text-center">
                            <p className={`text-lg font-bold ${getActivityScoreColor(user.activityScore)}`}>{user.activityScore}</p>
                            <div className="w-full bg-muted rounded-full h-2 mt-1">
                              <div className="bg-gradient-to-r from-primary to-chart-1 h-2 rounded-full transition-all duration-300" style={{ width: `${user.activityScore}%` }} />
                            </div>
                          </div>
                        </TableCell>
                        <TableCell className="text-right">
                          <Button variant="ghost" size="sm">
                            <Eye className="w-4 h-4" />
                          </Button>
                        </TableCell>
                      </TableRow>
                    ))}
                  </TableBody>
                </Table>
              </div>
            </CardContent>
          </Card>
        </TabsContent>

        <TabsContent value="activity" className="space-y-4">
          <Card className="border-border bg-card/50">
            <CardHeader>
              <CardTitle>Actividad por Hora</CardTitle>
              <CardDescription>Distribución de usuarios activos durante las últimas 24 horas</CardDescription>
            </CardHeader>
            <CardContent>
              <div className="space-y-4">
                <div className="grid grid-cols-12 gap-2">
                  {hourlyActivity.map((data, index) => (
                    <div key={index} className="text-center">
                      <div className="relative h-32 bg-accent/20 rounded-lg flex items-end justify-center p-1">
                        <div
                          className="bg-gradient-to-t from-primary to-chart-1 rounded-sm w-full transition-all duration-300 hover:from-primary/80 hover:to-chart-1/80"
                          style={{
                            height: `${(data.users / Math.max(...hourlyActivity.map((h) => h.users))) * 100}%`,
                            minHeight: '4px',
                          }}
                        />
                        <div className="absolute -top-6 left-1/2 transform -translate-x-1/2 text-xs font-medium text-foreground">{data.users}</div>
                      </div>
                      <p className="text-xs text-muted-foreground mt-2">{data.hour}</p>
                    </div>
                  ))}
                </div>
                <div className="flex items-center justify-center gap-4 text-sm text-muted-foreground">
                  <div className="flex items-center gap-2">
                    <div className="w-3 h-3 bg-gradient-to-r from-primary to-chart-1 rounded-sm" />
                    <span>Usuarios Activos</span>
                  </div>
                  <div className="flex items-center gap-2">
                    <TrendingUp className="w-4 h-4 text-chart-4" />
                    <span>Pico: {Math.max(...hourlyActivity.map((h) => h.users))} usuarios</span>
                  </div>
                </div>
              </div>
            </CardContent>
          </Card>
        </TabsContent>

        <TabsContent value="analytics" className="space-y-4">
          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            <Card className="border-border bg-card/50">
              <CardHeader>
                <CardTitle>Crecimiento de Usuarios</CardTitle>
                <CardDescription>Tendencia de usuarios activos en el tiempo</CardDescription>
              </CardHeader>
              <CardContent>
                <div className="space-y-4">
                  <div className="grid grid-cols-3 gap-4">
                    <div className="text-center p-4 bg-accent/20 rounded-lg">
                      <div className="flex items-center justify-center gap-2">
                        <TrendingUp className="w-4 h-4 text-chart-4" />
                        <span className="text-lg font-bold text-chart-4">+{activityStats.dailyGrowth}%</span>
                      </div>
                      <p className="text-sm text-muted-foreground">Diario</p>
                    </div>
                    <div className="text-center p-4 bg-accent/20 rounded-lg">
                      <div className="flex items-center justify-center gap-2">
                        <TrendingUp className="w-4 h-4 text-chart-4" />
                        <span className="text-lg font-bold text-chart-4">+{activityStats.weeklyGrowth}%</span>
                      </div>
                      <p className="text-sm text-muted-foreground">Semanal</p>
                    </div>
                    <div className="text-center p-4 bg-accent/20 rounded-lg">
                      <div className="flex items-center justify-center gap-2">
                        <TrendingUp className="w-4 h-4 text-chart-4" />
                        <span className="text-lg font-bold text-chart-4">+{activityStats.monthlyGrowth}%</span>
                      </div>
                      <p className="text-sm text-muted-foreground">Mensual</p>
                    </div>
                  </div>
                </div>
              </CardContent>
            </Card>

            <Card className="border-border bg-card/50">
              <CardHeader>
                <CardTitle>Distribución por Dispositivo</CardTitle>
                <CardDescription>Tipos de dispositivos más utilizados</CardDescription>
              </CardHeader>
              <CardContent>
                <div className="space-y-4">
                  <div className="space-y-3">
                    <div className="flex items-center justify-between">
                      <div className="flex items-center gap-2">
                        <span className="text-lg">🖥️</span>
                        <span className="text-sm font-medium">Desktop</span>
                      </div>
                      <div className="flex items-center gap-2">
                        <div className="w-24 bg-muted rounded-full h-2">
                          <div className="bg-primary h-2 rounded-full" style={{ width: '65%' }} />
                        </div>
                        <span className="text-sm font-medium">65%</span>
                      </div>
                    </div>
                    <div className="flex items-center justify-between">
                      <div className="flex items-center gap-2">
                        <span className="text-lg">📱</span>
                        <span className="text-sm font-medium">Mobile</span>
                      </div>
                      <div className="flex items-center gap-2">
                        <div className="w-24 bg-muted rounded-full h-2">
                          <div className="bg-chart-2 h-2 rounded-full" style={{ width: '25%' }} />
                        </div>
                        <span className="text-sm font-medium">25%</span>
                      </div>
                    </div>
                    <div className="flex items-center justify-between">
                      <div className="flex items-center gap-2">
                        <span className="text-lg">📱</span>
                        <span className="text-sm font-medium">Tablet</span>
                      </div>
                      <div className="flex items-center gap-2">
                        <div className="w-24 bg-muted rounded-full h-2">
                          <div className="bg-chart-3 h-2 rounded-full" style={{ width: '10%' }} />
                        </div>
                        <span className="text-sm font-medium">10%</span>
                      </div>
                    </div>
                  </div>
                </div>
              </CardContent>
            </Card>
          </div>
        </TabsContent>
      </Tabs>
    </div>
  );
}
