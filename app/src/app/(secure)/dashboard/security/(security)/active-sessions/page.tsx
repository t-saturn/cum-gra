'use client';

import { useState } from 'react';
import { Badge } from '@/components/ui/badge';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table';
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs';
import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar';
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from '@/components/ui/alert-dialog';
import { Search, Filter, Download, Monitor, Smartphone, Tablet, MapPin, Clock, AlertTriangle, X, Eye, Ban } from 'lucide-react';
import CardStatsContain from '@/components/custom/card/card-stats-contain';
import { statsSessions } from '@/mocks/stats-mocks';

// Datos simulados de sesiones activas
const activeSessions = [
  {
    id: 'sess_001',
    userId: 'user_001',
    userName: 'Ana García',
    userEmail: 'ana.garcia@empresa.com',
    avatar: '/placeholder.svg?height=32&width=32',
    device: 'Desktop',
    browser: 'Chrome 120.0',
    os: 'Windows 11',
    ip: '192.168.1.45',
    location: 'Madrid, España',
    startTime: '2024-01-15T08:30:00Z',
    lastActivity: '2024-01-15T14:25:00Z',
    duration: '5h 55m',
    applications: ['CRM', 'ERP', 'Analytics'],
    status: 'active',
    riskLevel: 'low',
    sessionType: 'web',
  },
  {
    id: 'sess_002',
    userId: 'user_002',
    userName: 'Carlos López',
    userEmail: 'carlos.lopez@empresa.com',
    avatar: '/placeholder.svg?height=32&width=32',
    device: 'Mobile',
    browser: 'Safari Mobile',
    os: 'iOS 17.2',
    ip: '10.0.0.123',
    location: 'Barcelona, España',
    startTime: '2024-01-15T09:15:00Z',
    lastActivity: '2024-01-15T14:20:00Z',
    duration: '5h 5m',
    applications: ['Mobile App', 'Notifications'],
    status: 'active',
    riskLevel: 'medium',
    sessionType: 'mobile',
  },
  {
    id: 'sess_003',
    userId: 'user_003',
    userName: 'María Rodríguez',
    userEmail: 'maria.rodriguez@empresa.com',
    avatar: '/placeholder.svg?height=32&width=32',
    device: 'Tablet',
    browser: 'Chrome Mobile',
    os: 'Android 14',
    ip: '172.16.0.89',
    location: 'Valencia, España',
    startTime: '2024-01-15T07:45:00Z',
    lastActivity: '2024-01-15T14:18:00Z',
    duration: '6h 33m',
    applications: ['Dashboard', 'Reports'],
    status: 'idle',
    riskLevel: 'low',
    sessionType: 'tablet',
  },
  {
    id: 'sess_004',
    userId: 'user_004',
    userName: 'David Martín',
    userEmail: 'david.martin@empresa.com',
    avatar: '/placeholder.svg?height=32&width=32',
    device: 'Desktop',
    browser: 'Firefox 121.0',
    os: 'macOS Sonoma',
    ip: '203.0.113.45',
    location: 'Sevilla, España',
    startTime: '2024-01-15T10:00:00Z',
    lastActivity: '2024-01-15T14:15:00Z',
    duration: '4h 15m',
    applications: ['Admin Panel', 'User Management'],
    status: 'active',
    riskLevel: 'high',
    sessionType: 'web',
  },
  {
    id: 'sess_005',
    userId: 'user_005',
    userName: 'Laura Fernández',
    userEmail: 'laura.fernandez@empresa.com',
    avatar: '/placeholder.svg?height=32&width=32',
    device: 'Desktop',
    browser: 'Edge 120.0',
    os: 'Windows 10',
    ip: '198.51.100.67',
    location: 'Bilbao, España',
    startTime: '2024-01-15T08:00:00Z',
    lastActivity: '2024-01-15T13:45:00Z',
    duration: '5h 45m',
    applications: ['Finance', 'Accounting'],
    status: 'warning',
    riskLevel: 'medium',
    sessionType: 'web',
  },
];

const sessionStats = {
  total: 156,
  active: 89,
  idle: 45,
  warning: 22,
  byDevice: {
    desktop: 78,
    mobile: 45,
    tablet: 33,
  },
  byRisk: {
    low: 98,
    medium: 43,
    high: 15,
  },
};

export default function ActiveSessionsManagement() {
  const [searchTerm, setSearchTerm] = useState('');
  const [statusFilter, setStatusFilter] = useState('all');
  const [deviceFilter, setDeviceFilter] = useState('all');
  const [riskFilter, setRiskFilter] = useState('all');

  const getStatusBadge = (status: string) => {
    const variants = {
      active: 'default',
      idle: 'secondary',
      warning: 'destructive',
    };
    const colors = {
      active: 'bg-green-100 text-green-800',
      idle: 'bg-yellow-100 text-yellow-800',
      warning: 'bg-red-100 text-red-800',
    };
    return {
      variant: variants[status as keyof typeof variants],
      className: colors[status as keyof typeof colors],
    };
  };

  const getRiskBadge = (risk: string) => {
    const colors = {
      low: 'bg-green-100 text-green-800',
      medium: 'bg-yellow-100 text-yellow-800',
      high: 'bg-red-100 text-red-800',
    };
    return colors[risk as keyof typeof colors];
  };

  const getDeviceIcon = (device: string) => {
    switch (device.toLowerCase()) {
      case 'desktop':
        return <Monitor className="h-4 w-4" />;
      case 'mobile':
        return <Smartphone className="h-4 w-4" />;
      case 'tablet':
        return <Tablet className="h-4 w-4" />;
      default:
        return <Monitor className="h-4 w-4" />;
    }
  };

  const filteredSessions = activeSessions.filter((session) => {
    const matchesSearch =
      session.userName.toLowerCase().includes(searchTerm.toLowerCase()) || session.userEmail.toLowerCase().includes(searchTerm.toLowerCase()) || session.ip.includes(searchTerm);
    const matchesStatus = statusFilter === 'all' || session.status === statusFilter;
    const matchesDevice = deviceFilter === 'all' || session.device.toLowerCase() === deviceFilter;
    const matchesRisk = riskFilter === 'all' || session.riskLevel === riskFilter;

    return matchesSearch && matchesStatus && matchesDevice && matchesRisk;
  });

  return (
    <div className="space-y-6">
      {/* Estadísticas */}
      <CardStatsContain stats={statsSessions} />

      <Tabs defaultValue="sessions" className="space-y-4">
        <TabsList>
          <TabsTrigger value="sessions">Sesiones Activas</TabsTrigger>
          <TabsTrigger value="analytics">Análisis</TabsTrigger>
          <TabsTrigger value="security">Seguridad</TabsTrigger>
        </TabsList>

        <TabsContent value="sessions" className="space-y-4">
          <Card>
            <CardHeader>
              <CardTitle>Gestión de Sesiones Activas</CardTitle>
              <CardDescription>Monitorea y controla todas las sesiones activas en tiempo real</CardDescription>
            </CardHeader>
            <CardContent>
              {/* Filtros */}
              <div className="flex flex-col gap-4 md:flex-row md:items-center md:justify-between mb-6">
                <div className="flex flex-1 items-center space-x-2">
                  <Search className="h-4 w-4 text-muted-foreground" />
                  <Input placeholder="Buscar por usuario, email o IP..." value={searchTerm} onChange={(e) => setSearchTerm(e.target.value)} className="max-w-sm" />
                </div>
                <div className="flex items-center space-x-2">
                  <Filter className="h-4 w-4 text-muted-foreground" />
                  <Select value={statusFilter} onValueChange={setStatusFilter}>
                    <SelectTrigger className="w-[140px]">
                      <SelectValue placeholder="Estado" />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value="all">Todos</SelectItem>
                      <SelectItem value="active">Activo</SelectItem>
                      <SelectItem value="idle">Inactivo</SelectItem>
                      <SelectItem value="warning">Advertencia</SelectItem>
                    </SelectContent>
                  </Select>
                  <Select value={deviceFilter} onValueChange={setDeviceFilter}>
                    <SelectTrigger className="w-[140px]">
                      <SelectValue placeholder="Dispositivo" />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value="all">Todos</SelectItem>
                      <SelectItem value="desktop">Desktop</SelectItem>
                      <SelectItem value="mobile">Mobile</SelectItem>
                      <SelectItem value="tablet">Tablet</SelectItem>
                    </SelectContent>
                  </Select>
                  <Select value={riskFilter} onValueChange={setRiskFilter}>
                    <SelectTrigger className="w-[140px]">
                      <SelectValue placeholder="Riesgo" />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value="all">Todos</SelectItem>
                      <SelectItem value="low">Bajo</SelectItem>
                      <SelectItem value="medium">Medio</SelectItem>
                      <SelectItem value="high">Alto</SelectItem>
                    </SelectContent>
                  </Select>
                  <Button variant="outline" size="sm">
                    <Download className="h-4 w-4 mr-2" />
                    Exportar
                  </Button>
                </div>
              </div>

              {/* Tabla de sesiones */}
              <div className="rounded-md border">
                <Table>
                  <TableHeader>
                    <TableRow>
                      <TableHead>Usuario</TableHead>
                      <TableHead>Dispositivo</TableHead>
                      <TableHead>Ubicación</TableHead>
                      <TableHead>Duración</TableHead>
                      <TableHead>Estado</TableHead>
                      <TableHead>Riesgo</TableHead>
                      <TableHead>Aplicaciones</TableHead>
                      <TableHead className="text-right">Acciones</TableHead>
                    </TableRow>
                  </TableHeader>
                  <TableBody>
                    {filteredSessions.map((session) => (
                      <TableRow key={session.id}>
                        <TableCell>
                          <div className="flex items-center space-x-3">
                            <Avatar className="h-8 w-8">
                              <AvatarImage src={session.avatar || '/placeholder.svg'} />
                              <AvatarFallback>
                                {session.userName
                                  .split(' ')
                                  .map((n) => n[0])
                                  .join('')}
                              </AvatarFallback>
                            </Avatar>
                            <div>
                              <div className="font-medium">{session.userName}</div>
                              <div className="text-sm text-muted-foreground">{session.userEmail}</div>
                            </div>
                          </div>
                        </TableCell>
                        <TableCell>
                          <div className="flex items-center space-x-2">
                            {getDeviceIcon(session.device)}
                            <div>
                              <div className="font-medium">{session.device}</div>
                              <div className="text-sm text-muted-foreground">{session.browser}</div>
                            </div>
                          </div>
                        </TableCell>
                        <TableCell>
                          <div className="flex items-center space-x-2">
                            <MapPin className="h-4 w-4 text-muted-foreground" />
                            <div>
                              <div className="font-medium">{session.location}</div>
                              <div className="text-sm text-muted-foreground">{session.ip}</div>
                            </div>
                          </div>
                        </TableCell>
                        <TableCell>
                          <div className="flex items-center space-x-2">
                            <Clock className="h-4 w-4 text-muted-foreground" />
                            <span className="font-medium">{session.duration}</span>
                          </div>
                        </TableCell>
                        <TableCell>
                          <Badge className={getStatusBadge(session.status).className}>
                            {session.status === 'active' ? 'Activo' : session.status === 'idle' ? 'Inactivo' : 'Advertencia'}
                          </Badge>
                        </TableCell>
                        <TableCell>
                          <Badge className={getRiskBadge(session.riskLevel)}>{session.riskLevel === 'low' ? 'Bajo' : session.riskLevel === 'medium' ? 'Medio' : 'Alto'}</Badge>
                        </TableCell>
                        <TableCell>
                          <div className="flex flex-wrap gap-1">
                            {session.applications.slice(0, 2).map((app, index) => (
                              <Badge key={index} variant="outline" className="text-xs">
                                {app}
                              </Badge>
                            ))}
                            {session.applications.length > 2 && (
                              <Badge variant="outline" className="text-xs">
                                +{session.applications.length - 2}
                              </Badge>
                            )}
                          </div>
                        </TableCell>
                        <TableCell className="text-right">
                          <div className="flex items-center justify-end space-x-2">
                            <Button variant="ghost" size="sm">
                              <Eye className="h-4 w-4" />
                            </Button>
                            <AlertDialog>
                              <AlertDialogTrigger asChild>
                                <Button variant="ghost" size="sm" className="text-red-600 hover:text-red-700">
                                  <X className="h-4 w-4" />
                                </Button>
                              </AlertDialogTrigger>
                              <AlertDialogContent>
                                <AlertDialogHeader>
                                  <AlertDialogTitle>Terminar Sesión</AlertDialogTitle>
                                  <AlertDialogDescription>
                                    ¿Estás seguro de que quieres terminar la sesión de {session.userName}? Esta acción no se puede deshacer.
                                  </AlertDialogDescription>
                                </AlertDialogHeader>
                                <AlertDialogFooter>
                                  <AlertDialogCancel>Cancelar</AlertDialogCancel>
                                  <AlertDialogAction className="bg-red-600 hover:bg-red-700">Terminar Sesión</AlertDialogAction>
                                </AlertDialogFooter>
                              </AlertDialogContent>
                            </AlertDialog>
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

        <TabsContent value="analytics" className="space-y-4">
          <div className="grid gap-4 md:grid-cols-2">
            <Card>
              <CardHeader>
                <CardTitle>Distribución por Dispositivo</CardTitle>
              </CardHeader>
              <CardContent>
                <div className="space-y-4">
                  <div className="flex items-center justify-between">
                    <div className="flex items-center space-x-2">
                      <Monitor className="h-4 w-4" />
                      <span>Desktop</span>
                    </div>
                    <div className="flex items-center space-x-2">
                      <div className="w-24 bg-gray-200 rounded-full h-2">
                        <div
                          className="bg-blue-600 h-2 rounded-full"
                          style={{
                            width: `${(sessionStats.byDevice.desktop / sessionStats.total) * 100}%`,
                          }}
                        ></div>
                      </div>
                      <span className="text-sm font-medium">{sessionStats.byDevice.desktop}</span>
                    </div>
                  </div>
                  <div className="flex items-center justify-between">
                    <div className="flex items-center space-x-2">
                      <Smartphone className="h-4 w-4" />
                      <span>Mobile</span>
                    </div>
                    <div className="flex items-center space-x-2">
                      <div className="w-24 bg-gray-200 rounded-full h-2">
                        <div
                          className="bg-green-600 h-2 rounded-full"
                          style={{
                            width: `${(sessionStats.byDevice.mobile / sessionStats.total) * 100}%`,
                          }}
                        ></div>
                      </div>
                      <span className="text-sm font-medium">{sessionStats.byDevice.mobile}</span>
                    </div>
                  </div>
                  <div className="flex items-center justify-between">
                    <div className="flex items-center space-x-2">
                      <Tablet className="h-4 w-4" />
                      <span>Tablet</span>
                    </div>
                    <div className="flex items-center space-x-2">
                      <div className="w-24 bg-gray-200 rounded-full h-2">
                        <div
                          className="bg-purple-600 h-2 rounded-full"
                          style={{
                            width: `${(sessionStats.byDevice.tablet / sessionStats.total) * 100}%`,
                          }}
                        ></div>
                      </div>
                      <span className="text-sm font-medium">{sessionStats.byDevice.tablet}</span>
                    </div>
                  </div>
                </div>
              </CardContent>
            </Card>

            <Card>
              <CardHeader>
                <CardTitle>Análisis de Riesgo</CardTitle>
              </CardHeader>
              <CardContent>
                <div className="space-y-4">
                  <div className="flex items-center justify-between">
                    <div className="flex items-center space-x-2">
                      <div className="w-3 h-3 bg-green-500 rounded-full"></div>
                      <span>Riesgo Bajo</span>
                    </div>
                    <div className="flex items-center space-x-2">
                      <div className="w-24 bg-gray-200 rounded-full h-2">
                        <div
                          className="bg-green-500 h-2 rounded-full"
                          style={{
                            width: `${(sessionStats.byRisk.low / sessionStats.total) * 100}%`,
                          }}
                        ></div>
                      </div>
                      <span className="text-sm font-medium">{sessionStats.byRisk.low}</span>
                    </div>
                  </div>
                  <div className="flex items-center justify-between">
                    <div className="flex items-center space-x-2">
                      <div className="w-3 h-3 bg-yellow-500 rounded-full"></div>
                      <span>Riesgo Medio</span>
                    </div>
                    <div className="flex items-center space-x-2">
                      <div className="w-24 bg-gray-200 rounded-full h-2">
                        <div
                          className="bg-yellow-500 h-2 rounded-full"
                          style={{
                            width: `${(sessionStats.byRisk.medium / sessionStats.total) * 100}%`,
                          }}
                        ></div>
                      </div>
                      <span className="text-sm font-medium">{sessionStats.byRisk.medium}</span>
                    </div>
                  </div>
                  <div className="flex items-center justify-between">
                    <div className="flex items-center space-x-2">
                      <div className="w-3 h-3 bg-red-500 rounded-full"></div>
                      <span>Riesgo Alto</span>
                    </div>
                    <div className="flex items-center space-x-2">
                      <div className="w-24 bg-gray-200 rounded-full h-2">
                        <div
                          className="bg-red-500 h-2 rounded-full"
                          style={{
                            width: `${(sessionStats.byRisk.high / sessionStats.total) * 100}%`,
                          }}
                        ></div>
                      </div>
                      <span className="text-sm font-medium">{sessionStats.byRisk.high}</span>
                    </div>
                  </div>
                </div>
              </CardContent>
            </Card>
          </div>
        </TabsContent>

        <TabsContent value="security" className="space-y-4">
          <Card>
            <CardHeader>
              <CardTitle>Alertas de Seguridad</CardTitle>
              <CardDescription>Sesiones que requieren atención inmediata</CardDescription>
            </CardHeader>
            <CardContent>
              <div className="space-y-4">
                {activeSessions
                  .filter((s) => s.riskLevel === 'high' || s.status === 'warning')
                  .map((session) => (
                    <div key={session.id} className="flex items-center justify-between p-4 border rounded-lg bg-red-50">
                      <div className="flex items-center space-x-3">
                        <AlertTriangle className="h-5 w-5 text-red-500" />
                        <div>
                          <div className="font-medium">{session.userName}</div>
                          <div className="text-sm text-muted-foreground">{session.riskLevel === 'high' ? 'Sesión de alto riesgo' : 'Actividad sospechosa detectada'}</div>
                        </div>
                      </div>
                      <div className="flex items-center space-x-2">
                        <Button variant="outline" size="sm">
                          <Eye className="h-4 w-4 mr-2" />
                          Revisar
                        </Button>
                        <Button variant="destructive" size="sm">
                          <Ban className="h-4 w-4 mr-2" />
                          Bloquear
                        </Button>
                      </div>
                    </div>
                  ))}
              </div>
            </CardContent>
          </Card>
        </TabsContent>
      </Tabs>
    </div>
  );
}
