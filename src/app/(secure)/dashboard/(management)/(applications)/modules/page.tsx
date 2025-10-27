'use client';

import { useEffect, useMemo, useState } from 'react';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Badge } from '@/components/ui/badge';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table';
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuLabel, DropdownMenuSeparator, DropdownMenuTrigger } from '@/components/ui/dropdown-menu';
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle } from '@/components/ui/dialog';
import { Search, Plus, Filter, Download, MoreHorizontal, Edit, Trash2, Eye, Grid3X3, Building2, Key, Loader2, Copy, Package, Boxes, Users, Timer } from 'lucide-react';
import { ModulesStatsCards } from '@/components/custom/card/modules-stats-card';
import type { ModuleWithAppDTO } from '@/types/modules';
import { fn_get_modules } from '@/actions/modules/fn_get_modules';

const StatusBadge = ({ status }: { status: string }) => {
  const base = 'border px-2 py-0.5 rounded-full text-xs';
  switch (status) {
    case 'active':
      return <span className={`${base} bg-emerald-500/15 text-emerald-500 border-emerald-500/30`}>Activo</span>;
    case 'maintenance':
      return <span className={`${base} bg-amber-500/15 text-amber-500 border-amber-500/30`}>Mantenimiento</span>;
    case 'suspended':
      return <span className={`${base} bg-red-500/15 text-red-500 border-red-500/30`}>Suspendido</span>;
    case 'inactive':
      return <span className={`${base} bg-zinc-500/15 text-zinc-500 border-zinc-500/30`}>Inactivo</span>;
    case 'development':
      return <span className={`${base} bg-blue-500/15 text-blue-500 border-blue-500/30`}>Desarrollo</span>;
    default:
      return <Badge variant="secondary">{status}</Badge>;
  }
};

export default function ModulesManagement() {
  const [searchTerm, setSearchTerm] = useState('');
  const [modules, setModules] = useState<ModuleWithAppDTO[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const [page, setPage] = useState(1);
  const [pageSize] = useState(20);
  const [total, setTotal] = useState(0);

  const [selected, setSelected] = useState<ModuleWithAppDTO | null>(null);
  const [isDetailOpen, setIsDetailOpen] = useState(false);

  useEffect(() => {
    const load = async () => {
      try {
        setLoading(true);
        setError(null);
        const list = await fn_get_modules(page, pageSize, false);
        setModules(list.data);
        setTotal(list.total);
      } catch (err: any) {
        console.error('Error al cargar módulos:', err);
        setError('No se pudieron cargar los módulos.');
      } finally {
        setLoading(false);
      }
    };
    load();
  }, [page, pageSize]);

  const filtered = useMemo(() => {
    const q = searchTerm.trim().toLowerCase();
    if (!q) return modules;
    return modules.filter((m) => {
      const inName = m.name.toLowerCase().includes(q);
      const inItem = (m.item ?? '').toLowerCase().includes(q);
      const inRoute = (m.route ?? '').toLowerCase().includes(q);
      const inApp = (m.application_name ?? '').toLowerCase().includes(q);
      const inClient = (m.application_client_id ?? '').toLowerCase().includes(q);
      return inName || inItem || inRoute || inApp || inClient;
    });
  }, [modules, searchTerm]);

  const totalPages = Math.max(1, Math.ceil(total / pageSize));
  const copy = (t: string) => navigator.clipboard?.writeText(t);

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
          <h1 className="font-bold text-foreground text-3xl">Módulos</h1>
          <p className="mt-1 text-muted-foreground">Gestiona los módulos y funciones de las aplicaciones</p>
        </div>
        <div className="flex gap-3">
          <Button variant="outline">
            <Download className="mr-2 w-4 h-4" />
            Exportar
          </Button>
          <Button className="bg-gradient-to-r from-primary hover:from-primary/90 to-chart-1 hover:to-chart-1/90 shadow-lg shadow-primary/25">
            <Plus className="mr-2 w-4 h-4" />
            Nuevo Módulo
          </Button>
        </div>
      </div>

      <ModulesStatsCards />

      <Card className="bg-card/50 border-border">
        <CardHeader>
          <div className="flex justify-between items-center">
            <div>
              <CardTitle>Lista de Módulos</CardTitle>
              <CardDescription>
                {filtered.length} de {total} módulos
              </CardDescription>
            </div>
            <div className="flex gap-2">
              <div className="relative">
                <Search className="top-1/2 left-3 absolute w-4 h-4 text-muted-foreground -translate-y-1/2 transform" />
                <Input
                  placeholder="Buscar módulo, ruta, app, client_id..."
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
                  <TableHead>Módulo</TableHead>
                  <TableHead>Ruta</TableHead>
                  <TableHead>Aplicación</TableHead>
                  <TableHead>Usuarios</TableHead>
                  <TableHead>Creado</TableHead>
                  <TableHead>Estado</TableHead>
                  <TableHead className="text-right">Acciones</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {filtered.map((m) => (
                  <TableRow key={m.id} className="hover:bg-accent/30">
                    <TableCell>
                      <div className="space-y-1">
                        <div className="flex items-center gap-2">
                          <Package className="w-4 h-4 text-primary" />
                          <p className="font-medium text-foreground">{m.name}</p>
                        </div>
                        {(m.item || m.icon) && (
                          <p className="flex items-center gap-2 text-muted-foreground text-xs">
                            {m.item && <span className="font-mono">item: {m.item}</span>}
                            {m.icon && (
                              <>
                                <span>·</span>
                                <span className="font-mono">icon: {m.icon}</span>
                              </>
                            )}
                          </p>
                        )}
                        <div className="flex items-center gap-2 text-muted-foreground text-xs">
                          <Key className="w-3 h-3" />
                          <span className="font-mono">{m.id}</span>
                          <Button variant="ghost" size="sm" className="px-2 h-6" onClick={() => copy(m.id)}>
                            <Copy className="w-3 h-3" />
                          </Button>
                        </div>
                      </div>
                    </TableCell>

                    <TableCell>
                      <span className="font-mono text-sm">{m.route ?? '—'}</span>
                    </TableCell>

                    <TableCell>
                      <div className="space-y-1">
                        <div className="flex items-center gap-2">
                          <Boxes className="w-4 h-4 text-chart-1" />
                          <p className="font-medium text-sm">{m.application_name ?? '—'}</p>
                        </div>
                        {m.application_client_id && <p className="font-mono text-muted-foreground text-xs">{m.application_client_id}</p>}
                      </div>
                    </TableCell>

                    <TableCell className="font-medium">
                      <div className="space-y-1">
                        <div className="flex items-center gap-2">
                          <Users className="w-4 h-4 text-chart-2" />
                          <span className="font-medium">{m.users_count}</span>
                        </div>
                      </div>
                    </TableCell>
                    <TableCell className="text-sm">
                      <div className="flex items-center gap-2">
                        {m.created_at ? (
                          <>
                            <Timer className="w-4 h-4 text-chart-3" />
                            {new Date(m.created_at).toLocaleDateString('es-ES')}
                          </>
                        ) : (
                          <span className="text-muted-foreground">—</span>
                        )}
                      </div>
                    </TableCell>
                    <TableCell>
                      <StatusBadge status={m.status} />
                    </TableCell>

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
                          <DropdownMenuItem
                            onClick={() => {
                              setSelected(m);
                              setIsDetailOpen(true);
                            }}
                          >
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

          {/* Paginación */}
          <div className="flex justify-between items-center mt-4">
            <p className="text-muted-foreground text-sm">
              Página {page} de {totalPages} · {total} resultados
            </p>
            <div className="flex gap-2">
              <Button variant="outline" size="sm" disabled={page <= 1} onClick={() => setPage((p) => Math.max(1, p - 1))}>
                Anterior
              </Button>
              <Button variant="outline" size="sm" disabled={page >= totalPages} onClick={() => setPage((p) => Math.min(totalPages, p + 1))}>
                Siguiente
              </Button>
            </div>
          </div>
        </CardContent>
      </Card>

      {/* Detalle */}
      <Dialog open={isDetailOpen} onOpenChange={(o) => !o && setIsDetailOpen(false)}>
        <DialogContent className="bg-card/80 backdrop-blur-xl border-border sm:max-w-[800px] max-h-[90vh] overflow-y-auto">
          <DialogHeader>
            <DialogTitle className="flex items-center gap-2">
              <Grid3X3 className="w-5 h-5" />
              {selected?.name}
            </DialogTitle>
            <DialogDescription>{selected?.route}</DialogDescription>
          </DialogHeader>

          {selected && (
            <div className="space-y-4 text-sm">
              <div className="gap-4 grid grid-cols-1 md:grid-cols-2">
                <Card className="bg-accent/20 border-border">
                  <CardContent className="p-4">
                    <h4 className="mb-2 font-semibold">Información General</h4>
                    <div className="space-y-2">
                      <div className="flex justify-between">
                        <span className="text-muted-foreground">ID:</span>
                        <span className="font-mono">{selected.id}</span>
                      </div>
                      {selected.item && (
                        <div className="flex justify-between">
                          <span className="text-muted-foreground">Item:</span>
                          <span className="font-mono">{selected.item}</span>
                        </div>
                      )}
                      {selected.parent_id && (
                        <div className="flex justify-between">
                          <span className="text-muted-foreground">Padre:</span>
                          <span className="font-mono">{selected.parent_id}</span>
                        </div>
                      )}
                      <div className="flex justify-between">
                        <span className="text-muted-foreground">Orden:</span>
                        <span>{selected.sort_order}</span>
                      </div>
                      <div className="flex justify-between">
                        <span className="text-muted-foreground">Estado:</span>
                        <StatusBadge status={selected.status} />
                      </div>
                      <div className="flex justify-between">
                        <span className="text-muted-foreground">Creado:</span>
                        <span>{new Date(selected.created_at).toLocaleString('es-ES')}</span>
                      </div>
                      <div className="flex justify-between">
                        <span className="text-muted-foreground">Actualizado:</span>
                        <span>{new Date(selected.updated_at).toLocaleString('es-ES')}</span>
                      </div>
                    </div>
                  </CardContent>
                </Card>

                <Card className="bg-accent/20 border-border">
                  <CardContent className="p-4">
                    <h4 className="mb-2 font-semibold">Aplicación Asociada</h4>
                    <div className="space-y-2">
                      <div className="flex items-center gap-2">
                        <Building2 className="w-4 h-4 text-chart-2" />
                        <span className="font-medium">{selected.application_name ?? '—'}</span>
                      </div>
                      {selected.application_client_id && (
                        <div className="flex items-center gap-2">
                          <Key className="w-4 h-4 text-muted-foreground" />
                          <span className="font-mono text-xs">{selected.application_client_id}</span>
                        </div>
                      )}
                      <div className="flex justify-between items-center">
                        <span className="text-muted-foreground">Usuarios:</span>
                        <span className="font-medium">{selected.users_count}</span>
                      </div>
                    </div>
                  </CardContent>
                </Card>
              </div>
            </div>
          )}

          <DialogFooter>
            <Button variant="outline" onClick={() => setIsDetailOpen(false)}>
              Cerrar
            </Button>
            <Button disabled className="bg-gradient-to-r from-primary to-chart-1">
              Editar (próx.)
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </div>
  );
}
