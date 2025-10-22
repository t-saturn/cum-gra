// app/(ruta)/roles-app/page.tsx
'use client';

import { useEffect, useMemo, useState } from 'react';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Badge } from '@/components/ui/badge';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { Search, Boxes, Shield, LayoutDashboard, Eye, Edit, Trash2, Key, ChevronLeft, ChevronRight, Loader2, Copy, Package } from 'lucide-react';

import type { RolesAppsResponse, RoleAppModulesItemDTO, ModuleMinimalDTO } from '@/types/roles_app';
import { fn_get_roles_apps } from '@/actions/roles_app/fn_get_roles_app';
import { ApplicationsStatsCards } from '@/components/custom/card/application-stats-card';

const PAGE_SIZE_OPTIONS = [10, 20, 50, 100];

export default function Page() {
  const [page, setPage] = useState<number>(1);
  const [pageSize, setPageSize] = useState<number>(20);
  const [isDeleted] = useState<boolean>(false);

  const [data, setData] = useState<RolesAppsResponse | null>(null);
  const [rows, setRows] = useState<RoleAppModulesItemDTO[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  const [searchTerm, setSearchTerm] = useState<string>('');

  useEffect(() => {
    const load = async () => {
      try {
        setLoading(true);
        setError(null);
        const resp = await fn_get_roles_apps(page, pageSize, isDeleted);
        setData(resp);
        setRows(resp.data);
      } catch (err: any) {
        console.error('Error al cargar roles-app:', err);
        setError('No se pudo cargar la información de roles por aplicación.');
      } finally {
        setLoading(false);
      }
    };
    load();
  }, [page, pageSize, isDeleted]);

  const filteredRows = useMemo(() => {
    const q = searchTerm.trim().toLowerCase();
    if (!q) return rows;
    const includes = (s?: string | null) => (s ?? '').toLowerCase().includes(q);

    return rows.filter((item) => {
      const inApp = includes(item.app.name) || includes(item.app.client_id) || includes(item.app.id);
      const inRole = includes(item.role.name) || includes(item.role.id);
      const inRoleModules = (item.role_modules ?? []).some((m: ModuleMinimalDTO) => includes(m.name) || includes(m.icon ?? '') || includes(m.id));
      const inAppModules = (item.app_modules ?? []).some((m: ModuleMinimalDTO) => includes(m.name) || includes(m.icon ?? '') || includes(m.id));
      return inApp || inRole || inRoleModules || inAppModules;
    });
  }, [rows, searchTerm]);

  const total = data?.total ?? 0;
  const canPrev = page > 1;
  const canNext = page * pageSize < total;

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
      {/* Header */}
      <div className="flex justify-between items-center">
        <div>
          <h1 className="font-bold text-3xl">Roles por Aplicación</h1>
          <p className="mt-1 text-muted-foreground">
            Vista consolidada en una sola tabla: <span className="font-medium">Aplicación</span>, <span className="font-medium">Rol</span> y{' '}
            <span className="font-medium">Módulos del rol</span>.
          </p>
        </div>
        <div className="flex items-center gap-2">
          <Select
            value={String(pageSize)}
            onValueChange={(v) => {
              setPageSize(Number(v));
              setPage(1);
            }}
          >
            <SelectTrigger className="w-[130px]">
              <SelectValue placeholder="Tamaño pág." />
            </SelectTrigger>
            <SelectContent>
              {PAGE_SIZE_OPTIONS.map((opt) => (
                <SelectItem key={opt} value={String(opt)}>
                  {opt} / pág
                </SelectItem>
              ))}
            </SelectContent>
          </Select>
        </div>
      </div>

      <ApplicationsStatsCards />

      <Card className="bg-card/50 border-border">
        <CardHeader>
          <div className="flex justify-between items-center gap-3">
            <div>
              <CardTitle>Relación App · Rol · Módulos</CardTitle>
              <CardDescription>
                {filteredRows.length} resultados filtrados · {total} totales
              </CardDescription>
            </div>
            <div className="relative w-full max-w-xs">
              <Search className="top-1/2 left-3 absolute w-4 h-4 text-muted-foreground -translate-y-1/2" />
              <Input placeholder="Buscar por app, client_id, rol o módulo…" value={searchTerm} onChange={(e) => setSearchTerm(e.target.value)} className="pl-10" />
            </div>
          </div>
        </CardHeader>

        <CardContent>
          <div className="border border-border rounded-lg overflow-hidden">
            <Table>
              <TableHeader>
                <TableRow className="bg-accent/50">
                  <TableHead className="w-[30%]">Aplicación / Client ID</TableHead>
                  <TableHead className="w-[20%]">Rol</TableHead>
                  <TableHead>Módulos del Rol</TableHead>
                  <TableHead className="w-[12%] text-right">Acciones</TableHead>
                </TableRow>
              </TableHeader>

              <TableBody>
                {filteredRows.map((item) => (
                  <TableRow key={`${item.app.id}-${item.role.id}`} className="hover:bg-accent/30">
                    <TableCell>
                      <div className="space-y-1">
                        <div className="flex items-center gap-2">
                          <Boxes className="w-4 h-4 text-primary" />
                          <span className="font-medium">{item.app.name}</span>
                        </div>
                        <div className="flex items-center gap-2 text-muted-foreground text-xs">
                          <Key className="w-3 h-3" />
                          <code className="font-mono">{item.app.client_id}</code>
                          <Button variant="ghost" size="sm" className="px-2 h-6" onClick={() => copyToClipboard(item.app.client_id)} title="Copiar client_id">
                            <Copy className="w-3 h-3" />
                          </Button>
                        </div>
                        <div className="font-mono text-muted-foreground text-xs">{item.app.id}</div>
                      </div>
                    </TableCell>

                    <TableCell>
                      <div className="flex items-center gap-2">
                        <Shield className="w-4 h-4 text-chart-1" />
                        <span className="font-medium">{item.role.name}</span>
                      </div>
                      <div className="font-mono text-muted-foreground text-xs">{item.role.id}</div>
                    </TableCell>

                    <TableCell>
                      <div className="flex flex-wrap gap-1">
                        {(item.role_modules ?? []).length > 0 ? (
                          item.role_modules.map((m) => (
                            <Badge key={m.id} variant="secondary" className="text-xs">
                              <span className="mr-1 align-middle">
                                <Package className="inline-block w-3 h-3" />
                              </span>
                              {m.name}
                              {m.icon && <span className="opacity-70 ml-1 text-[10px]">[{m.icon}]</span>}
                            </Badge>
                          ))
                        ) : (
                          <span className="text-muted-foreground text-sm">— Sin módulos asignados al rol —</span>
                        )}
                      </div>
                      {/* Si quieres mostrar también cuántos módulos totales tiene la app: */}
                      {item.app_modules && item.app_modules.length > 0 && (
                        <div className="mt-2 text-muted-foreground text-xs">
                          Módulos totales de la app: <span className="font-medium">{item.app_modules.length}</span>
                        </div>
                      )}
                    </TableCell>

                    <TableCell className="text-right">
                      <div className="flex justify-end items-center gap-2">
                        <Button variant="ghost" size="sm" title="Ver">
                          <Eye className="w-4 h-4" />
                        </Button>
                        <Button variant="ghost" size="sm" title="Editar">
                          <Edit className="w-4 h-4" />
                        </Button>
                        <Button variant="ghost" size="sm" className="text-red-600 hover:text-red-700" title="Eliminar">
                          <Trash2 className="w-4 h-4" />
                        </Button>
                      </div>
                    </TableCell>
                  </TableRow>
                ))}

                {filteredRows.length === 0 && (
                  <TableRow>
                    <TableCell colSpan={4} className="text-muted-foreground text-center">
                      No hay resultados que coincidan con la búsqueda.
                    </TableCell>
                  </TableRow>
                )}
              </TableBody>
            </Table>
          </div>

          <div className="flex justify-between items-center mt-4">
            <div className="text-muted-foreground text-sm">
              Página <span className="font-medium">{page}</span> · Mostrando <span className="font-medium">{filteredRows.length}</span> de{' '}
              <span className="font-medium">{total}</span> registros
            </div>

            <div className="flex items-center gap-2">
              <Button variant="outline" size="icon" disabled={!canPrev} onClick={() => setPage((p) => Math.max(1, p - 1))} title="Anterior">
                <ChevronLeft className="w-4 h-4" />
              </Button>
              <div className="px-2 text-muted-foreground text-sm">
                pág. <span className="font-medium">{page}</span>
              </div>
              <Button variant="outline" size="icon" disabled={!canNext} onClick={() => setPage((p) => p + 1)} title="Siguiente">
                <ChevronRight className="w-4 h-4" />
              </Button>
            </div>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}
