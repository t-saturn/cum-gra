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
import { Dialog, DialogContent, DialogDescription, DialogHeader, DialogTitle, DialogTrigger } from '@/components/ui/dialog';
import { Search, Filter, Download, AlertTriangle, XCircle, Clock, User, Settings, Database, Eye, Activity } from 'lucide-react';
import { Label } from '@/components/ui/label';
import CardStatsContain from '@/components/custom/card/card-stats-contain';
import { statsAuditEvents } from '@/mocks/stats-mocks';

// Datos simulados de logs de auditoría
const auditLogs = [
  {
    id: 'log_001',
    timestamp: '2024-01-15T14:30:25Z',
    userId: 'user_001',
    userName: 'Ana García',
    userEmail: 'ana.garcia@empresa.com',
    avatar: '/placeholder.svg?height=32&width=32',
    action: 'LOGIN',
    resource: 'Sistema',
    resourceId: null,
    description: 'Inicio de sesión exitoso',
    ip: '192.168.1.45',
    userAgent: 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36',
    status: 'success',
    severity: 'info',
    category: 'authentication',
    details: {
      method: 'password',
      location: 'Madrid, España',
      device: 'Desktop',
    },
  },
  {
    id: 'log_002',
    timestamp: '2024-01-15T14:25:18Z',
    userId: 'user_002',
    userName: 'Carlos López',
    userEmail: 'carlos.lopez@empresa.com',
    avatar: '/placeholder.svg?height=32&width=32',
    action: 'UPDATE',
    resource: 'Usuario',
    resourceId: 'user_045',
    description: 'Actualización de permisos de usuario',
    ip: '10.0.0.123',
    userAgent: 'Mozilla/5.0 (iPhone; CPU iPhone OS 17_2 like Mac OS X)',
    status: 'success',
    severity: 'medium',
    category: 'user_management',
    details: {
      changes: ['permissions.admin', 'permissions.reports'],
      targetUser: 'Pedro Sánchez',
      previousRole: 'User',
      newRole: 'Admin',
    },
  },
  {
    id: 'log_003',
    timestamp: '2024-01-15T14:20:45Z',
    userId: 'user_003',
    userName: 'María Rodríguez',
    userEmail: 'maria.rodriguez@empresa.com',
    avatar: '/placeholder.svg?height=32&width=32',
    action: 'DELETE',
    resource: 'Aplicación',
    resourceId: 'app_012',
    description: 'Eliminación de aplicación del sistema',
    ip: '172.16.0.89',
    userAgent: 'Mozilla/5.0 (Linux; Android 14; SM-G998B)',
    status: 'success',
    severity: 'high',
    category: 'application_management',
    details: {
      applicationName: 'Legacy CRM',
      affectedUsers: 23,
      dataRetention: '30 days',
    },
  },
  {
    id: 'log_004',
    timestamp: '2024-01-15T14:15:32Z',
    userId: 'user_004',
    userName: 'David Martín',
    userEmail: 'david.martin@empresa.com',
    avatar: '/placeholder.svg?height=32&width=32',
    action: 'LOGIN_FAILED',
    resource: 'Sistema',
    resourceId: null,
    description: 'Intento de inicio de sesión fallido - Credenciales incorrectas',
    ip: '203.0.113.45',
    userAgent: 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7)',
    status: 'failed',
    severity: 'warning',
    category: 'authentication',
    details: {
      attempts: 3,
      lockoutTime: '15 minutes',
      reason: 'Invalid password',
    },
  },
  {
    id: 'log_005',
    timestamp: '2024-01-15T14:10:15Z',
    userId: 'system',
    userName: 'Sistema',
    userEmail: 'system@empresa.com',
    avatar: '/placeholder.svg?height=32&width=32',
    action: 'BACKUP',
    resource: 'Base de Datos',
    resourceId: 'db_main',
    description: 'Backup automático de base de datos completado',
    ip: '127.0.0.1',
    userAgent: 'System/1.0',
    status: 'success',
    severity: 'info',
    category: 'system',
    details: {
      size: '2.3 GB',
      duration: '45 seconds',
      location: 'backup-server-01',
    },
  },
  {
    id: 'log_006',
    timestamp: '2024-01-15T14:05:28Z',
    userId: 'user_005',
    userName: 'Laura Fernández',
    userEmail: 'laura.fernandez@empresa.com',
    avatar: '/placeholder.svg?height=32&width=32',
    action: 'EXPORT',
    resource: 'Reporte',
    resourceId: 'report_financial_q4',
    description: 'Exportación de reporte financiero Q4 2023',
    ip: '198.51.100.67',
    userAgent: 'Mozilla/5.0 (Windows NT 10.0; Win64; x64)',
    status: 'success',
    severity: 'medium',
    category: 'data_access',
    details: {
      format: 'PDF',
      records: 15420,
      fileSize: '8.7 MB',
    },
  },
];

const auditStats = {
  total: 15420,
  today: 342,
  success: 13890,
  failed: 1530,
  byCategory: {
    authentication: 4520,
    user_management: 3210,
    application_management: 2890,
    data_access: 2450,
    system: 2350,
  },
  bySeverity: {
    info: 8920,
    warning: 4200,
    medium: 1800,
    high: 500,
  },
};

export default function AuditLogsManagement() {
  const [searchTerm, setSearchTerm] = useState('');
  const [actionFilter, setActionFilter] = useState('all');
  const [statusFilter, setStatusFilter] = useState('all');
  const [severityFilter, setSeverityFilter] = useState('all');
  const [categoryFilter] = useState('all');
  const [selectedLog, setSelectedLog] = useState<(typeof auditLogs)[0]>();

  const getStatusBadge = (status: string) => {
    const colors = {
      success: 'bg-green-100 text-green-800',
      failed: 'bg-red-100 text-red-800',
      warning: 'bg-yellow-100 text-yellow-800',
    };
    return colors[status as keyof typeof colors] || 'bg-gray-100 text-gray-800';
  };

  const getSeverityBadge = (severity: string) => {
    const colors = {
      info: 'bg-blue-100 text-blue-800',
      warning: 'bg-yellow-100 text-yellow-800',
      medium: 'bg-orange-100 text-orange-800',
      high: 'bg-red-100 text-red-800',
    };
    return colors[severity as keyof typeof colors] || 'bg-gray-100 text-gray-800';
  };

  const getActionIcon = (action: string) => {
    switch (action) {
      case 'LOGIN':
      case 'LOGOUT':
        return <User className="h-4 w-4" />;
      case 'UPDATE':
      case 'CREATE':
        return <Settings className="h-4 w-4" />;
      case 'DELETE':
        return <XCircle className="h-4 w-4" />;
      case 'BACKUP':
        return <Database className="h-4 w-4" />;
      case 'LOGIN_FAILED':
        return <AlertTriangle className="h-4 w-4" />;
      default:
        return <Activity className="h-4 w-4" />;
    }
  };

  const filteredLogs = auditLogs.filter((log) => {
    const matchesSearch =
      log.userName.toLowerCase().includes(searchTerm.toLowerCase()) ||
      log.description.toLowerCase().includes(searchTerm.toLowerCase()) ||
      log.resource.toLowerCase().includes(searchTerm.toLowerCase());
    const matchesAction = actionFilter === 'all' || log.action === actionFilter;
    const matchesStatus = statusFilter === 'all' || log.status === statusFilter;
    const matchesSeverity = severityFilter === 'all' || log.severity === severityFilter;
    const matchesCategory = categoryFilter === 'all' || log.category === categoryFilter;

    return matchesSearch && matchesAction && matchesStatus && matchesSeverity && matchesCategory;
  });

  const formatTimestamp = (timestamp: string) => {
    return new Date(timestamp).toLocaleString('es-ES', {
      year: 'numeric',
      month: '2-digit',
      day: '2-digit',
      hour: '2-digit',
      minute: '2-digit',
      second: '2-digit',
    });
  };

  return (
    <div className="space-y-6">
      {/* Estadísticas */}

      <CardStatsContain stats={statsAuditEvents} />

      <Tabs defaultValue="logs" className="space-y-4">
        <TabsList>
          <TabsTrigger value="logs">Logs de Auditoría</TabsTrigger>
          <TabsTrigger value="analytics">Análisis</TabsTrigger>
          <TabsTrigger value="alerts">Alertas</TabsTrigger>
        </TabsList>

        <TabsContent value="logs" className="space-y-4">
          <Card>
            <CardHeader>
              <CardTitle>Registro de Auditoría</CardTitle>
              <CardDescription>Historial completo de todas las actividades del sistema</CardDescription>
            </CardHeader>
            <CardContent>
              {/* Filtros */}
              <div className="flex flex-col gap-4 md:flex-row md:items-center md:justify-between mb-6">
                <div className="flex flex-1 items-center space-x-2">
                  <Search className="h-4 w-4 text-muted-foreground" />
                  <Input placeholder="Buscar por usuario, acción o recurso..." value={searchTerm} onChange={(e) => setSearchTerm(e.target.value)} className="max-w-sm" />
                </div>
                <div className="flex items-center space-x-2">
                  <Filter className="h-4 w-4 text-muted-foreground" />
                  <Select value={actionFilter} onValueChange={setActionFilter}>
                    <SelectTrigger className="w-[140px]">
                      <SelectValue placeholder="Acción" />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value="all">Todas</SelectItem>
                      <SelectItem value="LOGIN">Login</SelectItem>
                      <SelectItem value="LOGOUT">Logout</SelectItem>
                      <SelectItem value="CREATE">Crear</SelectItem>
                      <SelectItem value="UPDATE">Actualizar</SelectItem>
                      <SelectItem value="DELETE">Eliminar</SelectItem>
                      <SelectItem value="EXPORT">Exportar</SelectItem>
                    </SelectContent>
                  </Select>
                  <Select value={statusFilter} onValueChange={setStatusFilter}>
                    <SelectTrigger className="w-[140px]">
                      <SelectValue placeholder="Estado" />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value="all">Todos</SelectItem>
                      <SelectItem value="success">Exitoso</SelectItem>
                      <SelectItem value="failed">Fallido</SelectItem>
                      <SelectItem value="warning">Advertencia</SelectItem>
                    </SelectContent>
                  </Select>
                  <Select value={severityFilter} onValueChange={setSeverityFilter}>
                    <SelectTrigger className="w-[140px]">
                      <SelectValue placeholder="Severidad" />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value="all">Todas</SelectItem>
                      <SelectItem value="info">Info</SelectItem>
                      <SelectItem value="warning">Advertencia</SelectItem>
                      <SelectItem value="medium">Media</SelectItem>
                      <SelectItem value="high">Alta</SelectItem>
                    </SelectContent>
                  </Select>
                  <Button variant="outline" size="sm">
                    <Download className="h-4 w-4 mr-2" />
                    Exportar
                  </Button>
                </div>
              </div>

              {/* Tabla de logs */}
              <div className="rounded-md border">
                <Table>
                  <TableHeader>
                    <TableRow>
                      <TableHead>Fecha/Hora</TableHead>
                      <TableHead>Usuario</TableHead>
                      <TableHead>Acción</TableHead>
                      <TableHead>Recurso</TableHead>
                      <TableHead>Estado</TableHead>
                      <TableHead>Severidad</TableHead>
                      <TableHead>IP</TableHead>
                      <TableHead className="text-right">Acciones</TableHead>
                    </TableRow>
                  </TableHeader>
                  <TableBody>
                    {filteredLogs.map((log) => (
                      <TableRow key={log.id}>
                        <TableCell>
                          <div className="flex items-center space-x-2">
                            <Clock className="h-4 w-4 text-muted-foreground" />
                            <span className="font-mono text-sm">{formatTimestamp(log.timestamp)}</span>
                          </div>
                        </TableCell>
                        <TableCell>
                          <div className="flex items-center space-x-3">
                            <Avatar className="h-8 w-8">
                              <AvatarImage src={log.avatar || '/placeholder.svg'} />
                              <AvatarFallback>
                                {log.userName
                                  .split(' ')
                                  .map((n) => n[0])
                                  .join('')}
                              </AvatarFallback>
                            </Avatar>
                            <div>
                              <div className="font-medium">{log.userName}</div>
                              <div className="text-sm text-muted-foreground">{log.userEmail}</div>
                            </div>
                          </div>
                        </TableCell>
                        <TableCell>
                          <div className="flex items-center space-x-2">
                            {getActionIcon(log.action)}
                            <span className="font-medium">{log.action}</span>
                          </div>
                        </TableCell>
                        <TableCell>
                          <div>
                            <div className="font-medium">{log.resource}</div>
                            {log.resourceId && <div className="text-sm text-muted-foreground">{log.resourceId}</div>}
                          </div>
                        </TableCell>
                        <TableCell>
                          <Badge className={getStatusBadge(log.status)}>{log.status === 'success' ? 'Exitoso' : log.status === 'failed' ? 'Fallido' : 'Advertencia'}</Badge>
                        </TableCell>
                        <TableCell>
                          <Badge className={getSeverityBadge(log.severity)}>
                            {log.severity === 'info' ? 'Info' : log.severity === 'warning' ? 'Advertencia' : log.severity === 'medium' ? 'Media' : 'Alta'}
                          </Badge>
                        </TableCell>
                        <TableCell>
                          <span className="font-mono text-sm">{log.ip}</span>
                        </TableCell>
                        <TableCell className="text-right">
                          <Dialog>
                            <DialogTrigger asChild>
                              <Button variant="ghost" size="sm" onClick={() => setSelectedLog(log)}>
                                <Eye className="h-4 w-4" />
                              </Button>
                            </DialogTrigger>
                            <DialogContent className="max-w-2xl">
                              <DialogHeader>
                                <DialogTitle>Detalles del Evento de Auditoría</DialogTitle>
                                <DialogDescription>Información completa del evento seleccionado</DialogDescription>
                              </DialogHeader>
                              {selectedLog && (
                                <div className="space-y-4">
                                  <div className="grid grid-cols-2 gap-4">
                                    <div>
                                      <Label className="text-sm font-medium">Fecha/Hora</Label>
                                      <p className="text-sm">{formatTimestamp(selectedLog.timestamp)}</p>
                                    </div>
                                    <div>
                                      <Label className="text-sm font-medium">Usuario</Label>
                                      <p className="text-sm">{selectedLog.userName}</p>
                                    </div>
                                    <div>
                                      <Label className="text-sm font-medium">Acción</Label>
                                      <p className="text-sm">{selectedLog.action}</p>
                                    </div>
                                    <div>
                                      <Label className="text-sm font-medium">Recurso</Label>
                                      <p className="text-sm">{selectedLog.resource}</p>
                                    </div>
                                    <div>
                                      <Label className="text-sm font-medium">IP</Label>
                                      <p className="text-sm font-mono">{selectedLog.ip}</p>
                                    </div>
                                    <div>
                                      <Label className="text-sm font-medium">Estado</Label>
                                      <Badge className={getStatusBadge(selectedLog.status)}>{selectedLog.status}</Badge>
                                    </div>
                                  </div>
                                  <div>
                                    <Label className="text-sm font-medium">Descripción</Label>
                                    <p className="text-sm">{selectedLog.description}</p>
                                  </div>
                                  <div>
                                    <Label className="text-sm font-medium">User Agent</Label>
                                    <p className="text-sm font-mono break-all">{selectedLog.userAgent}</p>
                                  </div>
                                  {selectedLog.details && (
                                    <div>
                                      <Label className="text-sm font-medium">Detalles Adicionales</Label>
                                      <pre className="text-sm bg-gray-100 p-3 rounded-md overflow-auto">{JSON.stringify(selectedLog.details, null, 2)}</pre>
                                    </div>
                                  )}
                                </div>
                              )}
                            </DialogContent>
                          </Dialog>
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
                <CardTitle>Eventos por Categoría</CardTitle>
              </CardHeader>
              <CardContent>
                <div className="space-y-4">
                  {Object.entries(auditStats.byCategory).map(([category, count]) => (
                    <div key={category} className="flex items-center justify-between">
                      <span className="capitalize">{category.replace('_', ' ')}</span>
                      <div className="flex items-center space-x-2">
                        <div className="w-24 bg-gray-200 rounded-full h-2">
                          <div
                            className="bg-blue-600 h-2 rounded-full"
                            style={{
                              width: `${(count / auditStats.total) * 100}%`,
                            }}
                          ></div>
                        </div>
                        <span className="text-sm font-medium">{count.toLocaleString()}</span>
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
                  {Object.entries(auditStats.bySeverity).map(([severity, count]) => (
                    <div key={severity} className="flex items-center justify-between">
                      <div className="flex items-center space-x-2">
                        <div
                          className={`w-3 h-3 rounded-full ${
                            severity === 'info' ? 'bg-blue-500' : severity === 'warning' ? 'bg-yellow-500' : severity === 'medium' ? 'bg-orange-500' : 'bg-red-500'
                          }`}
                        ></div>
                        <span className="capitalize">{severity}</span>
                      </div>
                      <div className="flex items-center space-x-2">
                        <div className="w-24 bg-gray-200 rounded-full h-2">
                          <div
                            className={`h-2 rounded-full ${
                              severity === 'info' ? 'bg-blue-500' : severity === 'warning' ? 'bg-yellow-500' : severity === 'medium' ? 'bg-orange-500' : 'bg-red-500'
                            }`}
                            style={{
                              width: `${(count / auditStats.total) * 100}%`,
                            }}
                          ></div>
                        </div>
                        <span className="text-sm font-medium">{count.toLocaleString()}</span>
                      </div>
                    </div>
                  ))}
                </div>
              </CardContent>
            </Card>
          </div>
        </TabsContent>

        <TabsContent value="alerts" className="space-y-4">
          <Card>
            <CardHeader>
              <CardTitle>Alertas Críticas</CardTitle>
              <CardDescription>Eventos que requieren atención inmediata</CardDescription>
            </CardHeader>
            <CardContent>
              <div className="space-y-4">
                {auditLogs
                  .filter((log) => log.severity === 'high' || log.status === 'failed')
                  .map((log) => (
                    <div key={log.id} className="flex items-center justify-between p-4 border rounded-lg bg-red-50">
                      <div className="flex items-center space-x-3">
                        <AlertTriangle className="h-5 w-5 text-red-500" />
                        <div>
                          <div className="font-medium">{log.description}</div>
                          <div className="text-sm text-muted-foreground">
                            {log.userName} - {formatTimestamp(log.timestamp)}
                          </div>
                        </div>
                      </div>
                      <div className="flex items-center space-x-2">
                        <Badge className={getSeverityBadge(log.severity)}>{log.severity}</Badge>
                        <Button variant="outline" size="sm">
                          <Eye className="h-4 w-4 mr-2" />
                          Ver Detalles
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
