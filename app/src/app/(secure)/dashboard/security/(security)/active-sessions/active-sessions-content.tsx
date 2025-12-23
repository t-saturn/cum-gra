'use client';

import { useState, useEffect } from 'react';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Badge } from '@/components/ui/badge';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table';
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuLabel, DropdownMenuSeparator, DropdownMenuTrigger } from '@/components/ui/dropdown-menu';
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle } from '@/components/ui/dialog';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { Search, RefreshCw, MoreHorizontal, Eye, LogOut, Activity, User, Globe, Clock, AlertCircle, ChevronLeft, ChevronRight } from 'lucide-react';
import { toast } from 'sonner';
import { fn_get_all_sessions } from '@/actions/keycloak/sessions/fn_get_all_sessions';
import { fn_get_session_counts_by_client } from '@/actions/keycloak/sessions/fn_get_session_counts_by_client';
import { fn_logout_user } from '@/actions/keycloak/sessions/fn_logout_user';
import type { SessionItem, ClientSessionCount } from '@/types/sessions';
import { SessionsStatsCards } from '@/components/custom/card/sessions-stats-card';

export default function ActiveSessionsContent() {
  const [sessions, setSessions] = useState<SessionItem[]>([]);
  const [clientCounts, setClientCounts] = useState<ClientSessionCount[]>([]);
  const [searchTerm, setSearchTerm] = useState('');
  const [selectedClient, setSelectedClient] = useState<string>('all');
  const [loading, setLoading] = useState(true);
  const [refreshing, setRefreshing] = useState(false);
  const [error, setError] = useState<string | null>(null);

  // Paginación
  const [currentPage, setCurrentPage] = useState(1);
  const [pageSize, setPageSize] = useState(10);

  // Modals
  const [isDetailModalOpen, setIsDetailModalOpen] = useState(false);
  const [isLogoutDialogOpen, setIsLogoutDialogOpen] = useState(false);
  const [selectedSession, setSelectedSession] = useState<SessionItem | null>(null);

  const loadSessions = async () => {
    try {
      setLoading(true);
      setError(null);
      const [sessionsData, countsData] = await Promise.all([
        fn_get_all_sessions(),
        fn_get_session_counts_by_client(),
      ]);
      setSessions(sessionsData);
      setClientCounts(countsData);
    } catch (err: any) {
      setError(err.message ?? 'Error desconocido');
      toast.error('Error al cargar sesiones activas');
    } finally {
      setLoading(false);
    }
  };

  const handleRefresh = async () => {
    setRefreshing(true);
    await loadSessions();
    setRefreshing(false);
    toast.success('Sesiones actualizadas');
  };

  useEffect(() => {
    loadSessions();
  }, []);

  // Filtrar sesiones
  const filteredSessions = sessions.filter((session) => {
    const matchesSearch =
      session.username.toLowerCase().includes(searchTerm.toLowerCase()) ||
      session.userId.toLowerCase().includes(searchTerm.toLowerCase()) ||
      session.ipAddress.toLowerCase().includes(searchTerm.toLowerCase()) ||
      session.clientName.toLowerCase().includes(searchTerm.toLowerCase());

    const matchesClient = selectedClient === 'all' || session.clientId === selectedClient;

    return matchesSearch && matchesClient;
  });

  // Calcular paginación
  const totalItems = filteredSessions.length;
  const totalPages = Math.ceil(totalItems / pageSize);
  const startIndex = (currentPage - 1) * pageSize;
  const endIndex = startIndex + pageSize;
  const paginatedSessions = filteredSessions.slice(startIndex, endIndex);

  // Reset a página 1 cuando cambian los filtros
  useEffect(() => {
    setCurrentPage(1);
  }, [searchTerm, selectedClient, pageSize]);

  const handlePageChange = (newPage: number) => {
    setCurrentPage(newPage);
  };

  const handlePageSizeChange = (newSize: string) => {
    setPageSize(Number(newSize));
    setCurrentPage(1);
  };

  const handleViewDetails = (session: SessionItem) => {
    setSelectedSession(session);
    setIsDetailModalOpen(true);
  };

  const handleLogoutClick = (session: SessionItem) => {
    setSelectedSession(session);
    setIsLogoutDialogOpen(true);
  };

  const handleLogout = async () => {
    if (!selectedSession) return;

    try {
      await fn_logout_user(selectedSession.userId);
      toast.success(`Sesión de ${selectedSession.username} cerrada correctamente`);
      setIsLogoutDialogOpen(false);
      setSelectedSession(null);
      loadSessions();
    } catch (err: any) {
      toast.error(err.message || 'Error al cerrar sesión');
    }
  };

  const formatDate = (timestamp: number) => {
    return new Date(timestamp).toLocaleString('es-PE', {
      year: 'numeric',
      month: 'short',
      day: 'numeric',
      hour: '2-digit',
      minute: '2-digit',
    });
  };

  const getSessionDuration = (start: number) => {
    const duration = Date.now() - start;
    const hours = Math.floor(duration / (1000 * 60 * 60));
    const minutes = Math.floor((duration % (1000 * 60 * 60)) / (1000 * 60));

    if (hours > 0) {
      return `${hours}h ${minutes}m`;
    }
    return `${minutes}m`;
  };

  const getActivityBadge = (lastAccess: number) => {
    const timeSinceActivity = Date.now() - lastAccess;
    const minutes = Math.floor(timeSinceActivity / (1000 * 60));

    if (minutes < 5) {
      return <Badge className="bg-green-500/20 border-green-500/30 text-green-500">Activo</Badge>;
    } else if (minutes < 30) {
      return <Badge className="bg-yellow-500/20 border-yellow-500/30 text-yellow-500">Inactivo ({minutes}m)</Badge>;
    } else {
      return <Badge className="bg-muted border-muted-foreground/30 text-muted-foreground">Inactivo ({Math.floor(minutes / 60)}h)</Badge>;
    }
  };

  if (error) return <p className="py-10 text-destructive text-center">{error}</p>;

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <div>
          <h1 className="font-bold text-foreground text-3xl">Sesiones Activas</h1>
          <p className="mt-1 text-muted-foreground">Monitorea y gestiona las sesiones activas de usuarios en Keycloak</p>
        </div>
        <div className="flex gap-3">
          <Button variant="outline" onClick={handleRefresh} disabled={refreshing || loading}>
            <RefreshCw className={`mr-2 w-4 h-4 ${refreshing ? 'animate-spin' : ''}`} />
            Actualizar
          </Button>
        </div>
      </div>

      <SessionsStatsCards />

      <Card className="bg-card/50 border-border">
        <CardHeader>
          <div className="flex justify-between items-center">
            <div>
              <CardTitle>Lista de Sesiones</CardTitle>
              <CardDescription>
                Mostrando {startIndex + 1} - {Math.min(endIndex, totalItems)} de {totalItems} sesiones
              </CardDescription>
            </div>
            <div className="flex gap-2">
              <div className="relative">
                <Search className="top-1/2 left-3 absolute w-4 h-4 text-muted-foreground -translate-y-1/2 transform" />
                <Input placeholder="Buscar por usuario, IP..." value={searchTerm} onChange={(e) => setSearchTerm(e.target.value)} className="bg-background/50 pl-10 w-80" />
              </div>
              <Select value={selectedClient} onValueChange={setSelectedClient}>
                <SelectTrigger className="w-[250px]">
                  <SelectValue placeholder="Filtrar por cliente" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="all">Todos los clientes</SelectItem>
                  {clientCounts.map((client) => (
                    <SelectItem key={client.clientId} value={client.clientId}>
                      {client.clientName} ({client.sessionCount})
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
            </div>
          </div>
        </CardHeader>

        <CardContent>
          <div className="border border-border rounded-lg">
            <Table>
              <TableHeader>
                <TableRow className="bg-accent/50">
                  <TableHead>Usuario</TableHead>
                  <TableHead>Cliente</TableHead>
                  <TableHead>IP Address</TableHead>
                  <TableHead>Inicio</TableHead>
                  <TableHead>Duración</TableHead>
                  <TableHead>Estado</TableHead>
                  <TableHead className="text-right">Acciones</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {loading ? (
                  <TableRow>
                    <TableCell colSpan={7} className="text-center text-muted-foreground py-8">
                      Cargando sesiones...
                    </TableCell>
                  </TableRow>
                ) : paginatedSessions.length === 0 ? (
                  <TableRow>
                    <TableCell colSpan={7} className="text-center text-muted-foreground py-8">
                      No se encontraron sesiones activas
                    </TableCell>
                  </TableRow>
                ) : (
                  paginatedSessions.map((session) => (
                    <TableRow key={session.id}>
                      <TableCell>
                        <div className="flex items-center gap-2">
                          <User className="w-4 h-4 text-primary" />
                          <div>
                            <p className="font-medium text-foreground">{session.username}</p>
                            <p className="text-xs text-muted-foreground">{session.userId.slice(0, 8)}...</p>
                          </div>
                        </div>
                      </TableCell>
                      <TableCell>
                        <div className="flex items-center gap-2">
                          <Activity className="w-4 h-4 text-chart-1" />
                          <span className="text-sm">{session.clientName}</span>
                        </div>
                      </TableCell>
                      <TableCell>
                        <div className="flex items-center gap-2">
                          <Globe className="w-4 h-4 text-chart-4" />
                          <code className="text-xs bg-muted px-2 py-1 rounded">{session.ipAddress}</code>
                        </div>
                      </TableCell>
                      <TableCell>
                        <div className="flex items-center gap-2">
                          <Clock className="w-4 h-4 text-muted-foreground" />
                          <span className="text-sm">{formatDate(session.start)}</span>
                        </div>
                      </TableCell>
                      <TableCell>
                        <span className="text-sm font-medium">{getSessionDuration(session.start)}</span>
                      </TableCell>
                      <TableCell>{getActivityBadge(session.lastAccess)}</TableCell>
                      <TableCell className="text-right">
                        <DropdownMenu>
                          <DropdownMenuTrigger asChild>
                            <Button variant="ghost" size="sm">
                              <MoreHorizontal className="w-4 h-4" />
                            </Button>
                          </DropdownMenuTrigger>
                          <DropdownMenuContent align="end">
                            <DropdownMenuLabel>Acciones</DropdownMenuLabel>
                            <DropdownMenuSeparator />
                            <DropdownMenuItem onClick={() => handleViewDetails(session)}>
                              <Eye className="mr-2 w-4 h-4" />
                              Ver Detalles
                            </DropdownMenuItem>
                            <DropdownMenuItem className="text-destructive" onClick={() => handleLogoutClick(session)}>
                              <LogOut className="mr-2 w-4 h-4" />
                              Cerrar Sesión
                            </DropdownMenuItem>
                          </DropdownMenuContent>
                        </DropdownMenu>
                      </TableCell>
                    </TableRow>
                  ))
                )}
              </TableBody>
            </Table>
          </div>

          {/* Paginación */}
          <div className="flex items-center justify-between mt-4">
            <div className="flex items-center gap-2">
              <span className="text-sm text-muted-foreground">Mostrar</span>
              <Select value={String(pageSize)} onValueChange={handlePageSizeChange}>
                <SelectTrigger className="w-[100px]">
                  <SelectValue />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="5">5</SelectItem>
                  <SelectItem value="10">10</SelectItem>
                  <SelectItem value="20">20</SelectItem>
                  <SelectItem value="30">30</SelectItem>
                  <SelectItem value="50">50</SelectItem>
                  <SelectItem value="100">100</SelectItem>
                </SelectContent>
              </Select>
              <span className="text-sm text-muted-foreground">por página</span>
            </div>

            <div className="flex items-center gap-4">
              <div className="text-sm text-muted-foreground">
                Página {currentPage} de {totalPages || 1}
              </div>
              <div className="flex gap-2">
                <Button variant="outline" size="sm" onClick={() => handlePageChange(currentPage - 1)} disabled={currentPage === 1 || loading}>
                  <ChevronLeft className="w-4 h-4 mr-1" />
                  Anterior
                </Button>
                <Button variant="outline" size="sm" onClick={() => handlePageChange(currentPage + 1)} disabled={currentPage === totalPages || totalPages === 0 || loading}>
                  Siguiente
                  <ChevronRight className="w-4 h-4 ml-1" />
                </Button>
              </div>
            </div>
          </div>
        </CardContent>
      </Card>

      {/* Modal Detalles */}
      <Dialog open={isDetailModalOpen} onOpenChange={setIsDetailModalOpen}>
        <DialogContent className="sm:max-w-[600px]">
          <DialogHeader>
            <DialogTitle>Detalles de la Sesión</DialogTitle>
            <DialogDescription>Información completa de la sesión activa</DialogDescription>
          </DialogHeader>
          {selectedSession && (
            <div className="space-y-4">
              <div className="grid grid-cols-2 gap-4">
                <div>
                  <p className="text-sm text-muted-foreground">Usuario</p>
                  <p className="font-medium">{selectedSession.username}</p>
                </div>
                <div>
                  <p className="text-sm text-muted-foreground">Estado</p>
                  <div className="mt-1">{getActivityBadge(selectedSession.lastAccess)}</div>
                </div>
              </div>

              <div>
                <p className="text-sm text-muted-foreground">ID de Usuario</p>
                <code className="text-sm bg-muted px-2 py-1 rounded block mt-1">{selectedSession.userId}</code>
              </div>

              <div>
                <p className="text-sm text-muted-foreground">ID de Sesión</p>
                <code className="text-sm bg-muted px-2 py-1 rounded block mt-1 break-all">{selectedSession.id}</code>
              </div>

              <div className="grid grid-cols-2 gap-4">
                <div>
                  <p className="text-sm text-muted-foreground">Cliente</p>
                  <p className="font-medium">{selectedSession.clientName}</p>
                </div>
                <div>
                  <p className="text-sm text-muted-foreground">IP Address</p>
                  <code className="text-sm bg-muted px-2 py-1 rounded">{selectedSession.ipAddress}</code>
                </div>
              </div>

              <div className="grid grid-cols-2 gap-4">
                <div>
                  <p className="text-sm text-muted-foreground">Inicio de Sesión</p>
                  <p className="text-sm">{formatDate(selectedSession.start)}</p>
                </div>
                <div>
                  <p className="text-sm text-muted-foreground">Última Actividad</p>
                  <p className="text-sm">{formatDate(selectedSession.lastAccess)}</p>
                </div>
              </div>

              <div>
                <p className="text-sm text-muted-foreground">Duración Total</p>
                <p className="font-medium text-lg">{getSessionDuration(selectedSession.start)}</p>
              </div>
            </div>
          )}
          <DialogFooter>
            <Button variant="outline" onClick={() => setIsDetailModalOpen(false)}>
              Cerrar
            </Button>
            <Button
              variant="destructive"
              onClick={() => {
                setIsDetailModalOpen(false);
                if (selectedSession) handleLogoutClick(selectedSession);
              }}
            >
              <LogOut className="mr-2 w-4 h-4" />
              Cerrar Sesión
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>

      {/* Dialog Confirmar Cierre de Sesión */}
      <Dialog open={isLogoutDialogOpen} onOpenChange={setIsLogoutDialogOpen}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>¿Cerrar sesión del usuario?</DialogTitle>
            <DialogDescription>
              Esta acción cerrará todas las sesiones activas del usuario <strong>{selectedSession?.username}</strong>. El usuario deberá iniciar sesión nuevamente.
            </DialogDescription>
          </DialogHeader>
          <div className="bg-destructive/10 border border-destructive/20 rounded-lg p-4 flex gap-3">
            <AlertCircle className="w-5 h-5 text-destructive shrink-0 mt-0.5" />
            <div className="text-sm">
              <p className="font-medium text-destructive mb-1">Advertencia</p>
              <p className="text-muted-foreground">Esta acción no se puede deshacer. El usuario perderá el acceso inmediatamente a todas las aplicaciones.</p>
            </div>
          </div>
          <DialogFooter>
            <Button variant="outline" onClick={() => setIsLogoutDialogOpen(false)}>
              Cancelar
            </Button>
            <Button variant="destructive" onClick={handleLogout}>
              <LogOut className="mr-2 w-4 h-4" />
              Cerrar Sesión
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </div>
  );
}