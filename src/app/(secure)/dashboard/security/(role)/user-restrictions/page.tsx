'use client';

import { useEffect, useMemo, useState } from 'react';
import { Badge } from '@/components/ui/badge';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table';
import { Avatar, AvatarFallback } from '@/components/ui/avatar';
import { Search, Edit, Trash2, Eye, ChevronLeft, ChevronRight, Loader2, Users, KeyRound, Shield, Package } from 'lucide-react';
import { UsersRestrictionsStatsCards } from '@/components/custom/card/users-restrictions-stats-cards';

import { fn_get_users_restrictions } from '@/actions/users_restrictions/fn_get_users_restrictions';
import type { RolesRestrictResponse, RoleRestrictDTO, UserAppAssignmentDTO } from '@/types/users_restrictions';
import type { ModuleMinimalDTO } from '@/types/roles_app';

const PAGE_SIZE_OPTIONS = [10, 20, 50, 100];

type FlatRow = {
  userId: string;
  firstName: string | null;
  lastName: string | null;
  email: string;
  dni: string;
  appId: string;
  appName: string;
  clientId: string;
  roleId: string | null;
  roleName: string | null;
  modules: ModuleMinimalDTO[] | null;
  modulesRestrict: ModuleMinimalDTO[] | null;
};

export default function Page() {
  const [page, setPage] = useState<number>(1);
  const [pageSize, setPageSize] = useState<number>(20);
  const [isDeleted] = useState<boolean>(false);

  const [data, setData] = useState<RolesRestrictResponse | null>(null);
  const [rows, setRows] = useState<RoleRestrictDTO[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  const [searchTerm, setSearchTerm] = useState('');

  useEffect(() => {
    const load = async () => {
      try {
        setLoading(true);
        setError(null);
        const resp = await fn_get_users_restrictions(page, pageSize, isDeleted);
        setData(resp);
        setRows(resp.data);
      } catch (err: any) {
        console.error('Error al cargar users-restrictions:', err);
        setError('No se pudo cargar la información de restricciones por usuario.');
      } finally {
        setLoading(false);
      }
    };
    load();
  }, [page, pageSize, isDeleted]);

  const flatRows: FlatRow[] = useMemo(() => {
    return rows.flatMap((r) => {
      const { user, apps } = r;
      const full: FlatRow[] = (apps ?? []).map((a: UserAppAssignmentDTO) => ({
        userId: user.id,
        firstName: user.first_name ?? null,
        lastName: user.last_name ?? null,
        email: user.email,
        dni: user.dni,
        appId: a.app.id,
        appName: a.app.name,
        clientId: a.app.client_id,
        roleId: a.role?.id ?? null,
        roleName: a.role?.name ?? null,
        modules: a.modules ?? null,
        modulesRestrict: a.modules_restrict ?? null,
      }));
      return full.length > 0
        ? full
        : [
            {
              userId: user.id,
              firstName: user.first_name ?? null,
              lastName: user.last_name ?? null,
              email: user.email,
              dni: user.dni,
              appId: '',
              appName: '',
              clientId: '',
              roleId: null,
              roleName: null,
              modules: null,
              modulesRestrict: null,
            },
          ];
    });
  }, [rows]);

  // Filtro por texto
  const filtered = useMemo(() => {
    const q = searchTerm.trim().toLowerCase();
    if (!q) return flatRows;

    const inc = (s?: string | null) => (s ?? '').toLowerCase().includes(q);
    const modMatch = (mods: ModuleMinimalDTO[] | null) => (mods ?? []).some((m) => inc(m.name) || inc(m.icon ?? '') || inc(m.id));

    return flatRows.filter((r) => {
      const fullName = `${r.firstName ?? ''} ${r.lastName ?? ''}`.trim();
      return inc(fullName) || inc(r.email) || inc(r.dni) || inc(r.appName) || inc(r.clientId) || inc(r.roleName ?? '') || modMatch(r.modules) || modMatch(r.modulesRestrict);
    });
  }, [flatRows, searchTerm]);

  const total = data?.total ?? 0;
  const canPrev = page > 1;
  const canNext = page * pageSize < total;

  const initials = (first?: string | null, last?: string | null) => `${(first ?? '').slice(0, 1)}${(last ?? '').slice(0, 1)}`.toUpperCase() || 'U';

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
      <UsersRestrictionsStatsCards />

      <Card className="bg-card/50 border-border">
        <CardHeader>
          <div className="flex justify-between items-center gap-3">
            <div>
              <CardTitle>Restricciones por Usuario</CardTitle>
              <CardDescription>
                {filtered.length} resultados filtrados · {total} totales
              </CardDescription>
            </div>
            <div className="relative w-full max-w-xs">
              <Search className="top-1/2 left-3 absolute w-4 h-4 text-muted-foreground -translate-y-1/2" />
              <Input placeholder="Buscar por usuario, email, DNI, app, rol o módulo…" value={searchTerm} onChange={(e) => setSearchTerm(e.target.value)} className="pl-10" />
            </div>
          </div>
        </CardHeader>

        <CardContent>
          <div className="flex justify-between items-center mb-3">
            <div className="text-muted-foreground text-sm">Mostrando filas a partir de asignaciones App/Role por usuario.</div>
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

          <div className="border border-border rounded-lg overflow-hidden">
            <Table>
              <TableHeader>
                <TableRow className="bg-accent/50">
                  <TableHead className="w-[26%]">Usuario</TableHead>
                  <TableHead className="w-[20%]">Aplicación</TableHead>
                  <TableHead className="w-[14%]">Rol</TableHead>
                  <TableHead>Módulos Permitidos</TableHead>
                  <TableHead className="w-[18%]">Módulos Restringidos</TableHead>
                  <TableHead className="w-[10%] text-right">Acciones</TableHead>
                </TableRow>
              </TableHeader>

              <TableBody>
                {filtered.map((r, idx) => {
                  const fullName = `${r.firstName ?? ''} ${r.lastName ?? ''}`.trim() || r.email;
                  return (
                    <TableRow key={`${r.userId}-${r.appId}-${r.roleId ?? 'norole'}-${idx}`} className="hover:bg-accent/30">
                      <TableCell>
                        <div className="flex items-start gap-3">
                          <Avatar className="w-10 h-10">
                            <AvatarFallback>{initials(r.firstName, r.lastName)}</AvatarFallback>
                          </Avatar>
                          <div className="space-y-1">
                            <div className="font-medium">{fullName}</div>
                            <div className="flex items-center gap-2 text-muted-foreground text-sm">
                              <Users className="w-3.5 h-3.5" />
                              <span>{r.email}</span>
                            </div>
                            <div className="flex items-center gap-2 text-muted-foreground text-xs">
                              <KeyRound className="w-3.5 h-3.5" />
                              <span>DNI: {r.dni}</span>
                            </div>
                          </div>
                        </div>
                      </TableCell>

                      <TableCell>
                        {r.appId ? (
                          <div className="space-y-1">
                            <div className="flex items-center gap-2">
                              <Package className="w-4 h-4 text-primary" />
                              <span className="font-medium">{r.appName}</span>
                            </div>
                            <div className="font-mono text-muted-foreground text-xs">{r.clientId}</div>
                            <div className="font-mono text-muted-foreground text-xs">{r.appId}</div>
                          </div>
                        ) : (
                          <span className="text-muted-foreground text-sm">— Sin aplicación —</span>
                        )}
                      </TableCell>

                      <TableCell>
                        {r.roleName ? (
                          <div className="space-y-1">
                            <div className="flex items-center gap-2">
                              <Shield className="w-4 h-4 text-chart-1" />
                              <span className="font-medium capitalize">{r.roleName}</span>
                            </div>
                            {r.roleId && <div className="font-mono text-muted-foreground text-xs">{r.roleId}</div>}
                          </div>
                        ) : (
                          <span className="text-muted-foreground text-sm">— Sin rol —</span>
                        )}
                      </TableCell>

                      <TableCell>
                        <div className="flex flex-wrap gap-1">
                          {(r.modules ?? []).length > 0 ? (
                            r.modules!.map((m) => (
                              <Badge key={m.id} variant="secondary" className="text-xs">
                                <Package className="inline-block mr-1 w-3 h-3 align-middle" />
                                {m.name}
                                {m.icon && <span className="opacity-70 ml-1 text-[10px]">[{m.icon}]</span>}
                              </Badge>
                            ))
                          ) : (
                            <span className="text-muted-foreground text-sm">—</span>
                          )}
                        </div>
                      </TableCell>

                      <TableCell>
                        <div className="flex flex-wrap gap-1">
                          {(r.modulesRestrict ?? []).length > 0 ? (
                            r.modulesRestrict!.map((m) => (
                              <Badge key={m.id} variant="destructive" className="text-xs">
                                <Package className="inline-block mr-1 w-3 h-3 align-middle" />
                                {m.name}
                                {m.icon && <span className="opacity-70 ml-1 text-[10px]">[{m.icon}]</span>}
                              </Badge>
                            ))
                          ) : (
                            <span className="text-muted-foreground text-sm">—</span>
                          )}
                        </div>
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
                  );
                })}

                {filtered.length === 0 && (
                  <TableRow>
                    <TableCell colSpan={6} className="text-muted-foreground text-center">
                      No hay resultados que coincidan con la búsqueda.
                    </TableCell>
                  </TableRow>
                )}
              </TableBody>
            </Table>
          </div>

          <div className="flex justify-between items-center mt-4">
            <div className="text-muted-foreground text-sm">
              Página <span className="font-medium">{page}</span> · Mostrando <span className="font-medium">{filtered.length}</span> de <span className="font-medium">{total}</span>{' '}
              registros
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
