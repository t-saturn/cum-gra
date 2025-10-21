'use client';

import { useEffect, useMemo, useState } from 'react';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Badge } from '@/components/ui/badge';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table';
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuLabel, DropdownMenuSeparator, DropdownMenuTrigger } from '@/components/ui/dropdown-menu';
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle, DialogTrigger } from '@/components/ui/dialog';
import { Label } from '@/components/ui/label';
import { Textarea } from '@/components/ui/textarea';
import { Search, Plus, Filter, Download, MoreHorizontal, Edit, Trash2, Eye, Loader2, Copy, Globe, Key, Users, Boxes, ShieldUser, Clock } from 'lucide-react';
import type { ApplicationItem } from '@/types/applications';
import { fn_get_applications } from '@/actions/applications/get_applications';
import { ApplicationsStatsCards } from '@/components/custom/card/application-stats-card';

export default function Page() {
  const [searchTerm, setSearchTerm] = useState('');
  const [isCreateDialogOpen, setIsCreateDialogOpen] = useState(false);
  const [selectedApp, setSelectedApp] = useState<ApplicationItem | null>(null);
  const [applications, setApplications] = useState<ApplicationItem[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const load = async () => {
      try {
        setLoading(true);
        const data = await fn_get_applications(1, 40, false);
        setApplications(data.data);
      } catch (err: any) {
        console.error('Error al cargar aplicaciones:', err);
        setError('No se pudieron cargar las aplicaciones.');
      } finally {
        setLoading(false);
      }
    };
    load();
  }, []);

  const filteredApps = useMemo(() => {
    const q = searchTerm.trim().toLowerCase();
    if (!q) return applications;
    return applications.filter((app) => {
      const inName = app.name.toLowerCase().includes(q);
      const inDesc = (app.description ?? '').toLowerCase().includes(q);
      const inDomain = (app.domain ?? '').toLowerCase().includes(q);
      const inAdmins = (app.admins ?? []).some((a) => a.full_name.toLowerCase().includes(q) || a.email.toLowerCase().includes(q));
      return inName || inDesc || inDomain || inAdmins;
    });
  }, [applications, searchTerm]);

  const getStatusBadge = (status: string) => {
    const base = 'border px-2 py-0.5 rounded-full text-xs';
    switch (status) {
      case 'active':
        return <span className={`${base} bg-emerald-500/15 text-emerald-500 border-emerald-500/30`}>Activa</span>;
      case 'development':
        return <span className={`${base} bg-blue-500/15 text-blue-500 border-blue-500/30`}>Desarrollo</span>;
      case 'suspended':
        return <span className={`${base} bg-red-500/15 text-red-500 border-red-500/30`}>Suspendida</span>;
      case 'inactive':
        return <span className={`${base} bg-zinc-500/15 text-zinc-500 border-zinc-500/30`}>Inactiva</span>;
      default:
        return <Badge variant="secondary">{status}</Badge>;
    }
  };

  const copyToClipboard = (text: string) => navigator.clipboard?.writeText(text);

  if (loading) {
    return (
      <div className="flex justify-center items-center h-64">
        <Loader2 className="w-8 h-8 text-primary animate-spin" />
      </div>
    );
  }

  if (error) {
    return <div className="py-12 text-muted-foreground text-center">{error}</div>;
  }

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <div>
          <h1 className="font-bold text-foreground text-3xl">Aplicaciones</h1>
          <p className="mt-1 text-muted-foreground">Gestiona las aplicaciones OAuth y sus configuraciones</p>
        </div>
        <div className="flex gap-3">
          <Button variant="outline">
            <Download className="mr-2 w-4 h-4" />
            Exportar
          </Button>
          <Dialog open={isCreateDialogOpen} onOpenChange={setIsCreateDialogOpen}>
            <DialogTrigger asChild>
              <Button className="bg-gradient-to-r from-primary hover:from-primary/90 to-chart-1 hover:to-chart-1/90 shadow-lg shadow-primary/25">
                <Plus className="mr-2 w-4 h-4" />
                Nueva Aplicación
              </Button>
            </DialogTrigger>
            <DialogContent className="bg-card/80 backdrop-blur-xl border-border sm:max-w-[700px]">
              <DialogHeader>
                <DialogTitle>Crear Nueva Aplicación</DialogTitle>
                <DialogDescription>Registra una nueva aplicación en el sistema.</DialogDescription>
              </DialogHeader>
              <div className="gap-4 grid py-4">
                <div className="space-y-2">
                  <Label htmlFor="name">Nombre</Label>
                  <Input id="name" placeholder="Ej: Central User Manager" />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="domain">Dominio</Label>
                  <Input id="domain" placeholder="Ej: inventario.empresa.com" />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="description">Descripción</Label>
                  <Textarea id="description" placeholder="Describe la funcionalidad de la aplicación..." />
                </div>
              </div>
              <DialogFooter>
                <Button variant="outline" onClick={() => setIsCreateDialogOpen(false)}>
                  Cancelar
                </Button>
                <Button className="bg-gradient-to-r from-primary to-chart-1">Crear</Button>
              </DialogFooter>
            </DialogContent>
          </Dialog>
        </div>
      </div>

      <ApplicationsStatsCards />

      <Card className="bg-card/50 border-border">
        <CardHeader>
          <div className="flex justify-between items-center">
            <div>
              <CardTitle>Lista de Aplicaciones</CardTitle>
              <CardDescription>
                {filteredApps.length} de {applications.length} aplicaciones
              </CardDescription>
            </div>
            <div className="flex gap-2">
              <div className="relative">
                <Search className="top-1/2 left-3 absolute w-4 h-4 text-muted-foreground -translate-y-1/2 transform" />
                <Input
                  placeholder="Buscar aplicaciones..."
                  value={searchTerm}
                  onChange={(e) => setSearchTerm(e.target.value)}
                  className="bg-background/50 pl-10 focus:border-primary border-border focus:ring-ring w-80"
                />
              </div>
              <Button variant="outline">
                <Filter className="mr-2 w-4 h-4" />
                Filtros
              </Button>
            </div>
          </div>
        </CardHeader>

        <CardContent>
          <div className="border border-border rounded-lg overflow-hidden">
            <Table>
              <TableHeader>
                <TableRow className="bg-accent/50">
                  <TableHead>Aplicación</TableHead>
                  <TableHead>Dominio</TableHead>
                  <TableHead>Administradores</TableHead>
                  <TableHead>Usuarios</TableHead>
                  <TableHead>Creada</TableHead>
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
                          <Boxes className="w-4 h-4 text-primary" />
                          <p className="font-medium text-foreground">{app.name}</p>
                        </div>
                        {app.description && <p className="text-muted-foreground text-sm line-clamp-2">{app.description}</p>}
                        <div className="flex items-center gap-2 text-muted-foreground text-xs">
                          <Key className="w-3 h-3" />
                          <span className="font-mono">{app.client_id}</span>
                          <Button variant="ghost" size="sm" className="px-2 h-6" onClick={() => copyToClipboard(app.client_id)}>
                            <Copy className="w-3 h-3" />
                          </Button>
                        </div>
                      </div>
                    </TableCell>
                    <TableCell>
                      <div className="space-y-1">
                        <div className="flex items-center gap-2">
                          <Globe className="w-4 h-4 text-chart-4" />
                          <span className="font-mono">{app.domain || '—'}</span>
                        </div>
                      </div>
                    </TableCell>
                    <TableCell>
                      <div className="space-y-1">
                        <div className="flex items-center gap-2">
                          {(app.admins?.length ?? 0) > 0 ? (
                            <>
                              <ShieldUser className="w-4 h-4 text-chart-1" />
                              {(app.admins ?? []).slice(0, 2).map((a, idx) => (
                                <p key={idx} className="text-sm">
                                  <span className="font-medium">{a.full_name}</span>
                                  <span className="text-muted-foreground text-xs"> · {a.email}</span>
                                </p>
                              ))}
                              {app.admins!.length > 2 && <p className="text-muted-foreground text-xs">+{app.admins!.length - 2} más</p>}
                            </>
                          ) : (
                            <p className="text-muted-foreground text-sm">—</p>
                          )}
                        </div>
                      </div>
                    </TableCell>
                    <TableCell className="font-medium">
                      <div className="space-y-1">
                        <div className="flex items-center gap-2">
                          <Users className="w-4 h-4 text-chart-2" />
                          <span className="font-medium">{app.users_count}</span>
                        </div>
                      </div>
                    </TableCell>
                    <TableCell className="text-sm">
                      <div className="space-y-1">
                        <div className="flex items-center gap-2">
                          <Clock className="w-4 h-4 text-chart-2" />
                          <span className="font-medium">{new Date(app.created_at).toLocaleDateString('es-ES')}</span>
                        </div>
                      </div>
                    </TableCell>
                    <TableCell>{getStatusBadge(app.status)}</TableCell>
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
                          <DropdownMenuItem onClick={() => setSelectedApp(app)}>
                            <Eye className="mr-2 w-4 h-4" />
                            Ver Detalles
                          </DropdownMenuItem>
                          <DropdownMenuItem disabled>
                            <Edit className="mr-2 w-4 h-4" />
                            Editar (próx.)
                          </DropdownMenuItem>
                          <DropdownMenuSeparator />
                          <DropdownMenuItem className="text-destructive" disabled>
                            <Trash2 className="mr-2 w-4 h-4" />
                            Eliminar (próx.)
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

      <Dialog open={!!selectedApp} onOpenChange={(o) => !o && setSelectedApp(null)}>
        <DialogContent className="bg-card/80 backdrop-blur-xl border-border sm:max-w-[700px] max-h-[90vh] overflow-y-auto">
          <DialogHeader>
            <DialogTitle>{selectedApp?.name}</DialogTitle>
            <DialogDescription>{selectedApp?.description}</DialogDescription>
          </DialogHeader>
          {selectedApp && (
            <div className="space-y-3 text-sm">
              <div className="flex justify-between">
                <span className="text-muted-foreground">Dominio:</span>
                <span className="font-mono">{selectedApp.domain || '—'}</span>
              </div>
              <div className="flex justify-between">
                <span className="text-muted-foreground">Usuarios:</span>
                <span className="font-medium">{selectedApp.users_count}</span>
              </div>
              <div className="flex justify-between">
                <span className="text-muted-foreground">Estado:</span>
                {getStatusBadge(selectedApp.status)}
              </div>
              <div>
                <Label className="font-medium text-sm">Client ID</Label>
                <div className="flex items-center gap-2 mt-1">
                  <Input value={selectedApp.client_id} readOnly className="font-mono text-sm" />
                  <Button variant="outline" size="sm" onClick={() => copyToClipboard(selectedApp.client_id)}>
                    <Copy className="w-4 h-4" />
                  </Button>
                </div>
              </div>
            </div>
          )}
          <DialogFooter>
            <Button variant="outline" onClick={() => setSelectedApp(null)}>
              Cerrar
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </div>
  );
}
