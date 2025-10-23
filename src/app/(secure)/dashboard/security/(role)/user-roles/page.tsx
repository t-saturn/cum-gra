'use client';

import { useEffect, useMemo, useState } from 'react';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { Avatar, AvatarFallback } from '@/components/ui/avatar';
import { Search, Eye, Edit, Trash2, ChevronLeft, ChevronRight, Loader2, Users, KeyRound, Boxes, Shield } from 'lucide-react';

import type { RolesAssignmentsResponse, RoleAssignmentsDTO } from '@/types/roles_assignments';
import { fn_get_roles_assignments } from '@/actions/roles_assignments/fn_get_roles_assignments';

const PAGE_SIZE_OPTIONS = [10, 20, 50, 100];

export default function Page() {
  const [page, setPage] = useState<number>(1);
  const [pageSize, setPageSize] = useState<number>(20);
  const [isDeleted] = useState<boolean>(false);

  const [data, setData] = useState<RolesAssignmentsResponse | null>(null);
  const [rows, setRows] = useState<RoleAssignmentsDTO[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  const [searchTerm, setSearchTerm] = useState('');

  useEffect(() => {
    const load = async () => {
      try {
        setLoading(true);
        setError(null);
        const resp = await fn_get_roles_assignments(page, pageSize, isDeleted);
        setData(resp);
        setRows(resp.data);
      } catch (err: any) {
        console.error('Error al cargar roles-assignments:', err);
        setError('No se pudo cargar la información de asignaciones de roles.');
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

    return rows.filter(({ user, assignments }) => {
      const inUser = includes(user.email) || includes(user.dni) || includes(user.first_name ?? '') || includes(user.last_name ?? '');
      const inAssignments = assignments.some((a) => includes(a.application.name) || includes(a.role.name));
      return inUser || inAssignments;
    });
  }, [rows, searchTerm]);

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
      <div className="flex justify-between items-center">
        <div>
          <h1 className="font-bold text-3xl">Asignaciones de Roles</h1>
          <p className="mt-1 text-muted-foreground">Vista consolidada por usuario con sus roles asignados por aplicación.</p>
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

      <Card className="bg-card/50 border-border">
        <CardHeader>
          <div className="flex justify-between items-center gap-3">
            <div>
              <CardTitle>Usuarios y sus Roles</CardTitle>
              <CardDescription>
                {filteredRows.length} resultados filtrados · {total} totales
              </CardDescription>
            </div>
            <div className="relative w-full max-w-xs">
              <Search className="top-1/2 left-3 absolute w-4 h-4 text-muted-foreground -translate-y-1/2" />
              <Input placeholder="Buscar por nombre, email, DNI, app o rol…" value={searchTerm} onChange={(e) => setSearchTerm(e.target.value)} className="pl-10" />
            </div>
          </div>
        </CardHeader>

        <CardContent>
          <div className="border border-border rounded-lg overflow-hidden">
            <Table>
              <TableHeader>
                <TableRow className="bg-accent/50">
                  <TableHead className="w-[32%]">Usuario</TableHead>
                  <TableHead>Asignaciones (Aplicación → Rol)</TableHead>
                  <TableHead className="w-[12%] text-right">Opciones</TableHead>
                </TableRow>
              </TableHeader>

              <TableBody>
                {filteredRows.map(({ user, assignments }) => {
                  const fullName = `${user.first_name ?? ''} ${user.last_name ?? ''}`.trim() || user.email;
                  return (
                    <TableRow key={user.id} className="hover:bg-accent/30">
                      <TableCell>
                        <div className="flex items-start gap-3">
                          <Avatar className="w-10 h-10">
                            <AvatarFallback>{initials(user.first_name, user.last_name)}</AvatarFallback>
                          </Avatar>
                          <div className="space-y-1">
                            <div className="font-medium">{fullName}</div>
                            <div className="flex items-center gap-2 text-muted-foreground text-sm">
                              <Users className="w-3.5 h-3.5" />
                              <span>{user.email}</span>
                            </div>
                            <div className="flex items-center gap-2 text-muted-foreground text-xs">
                              <KeyRound className="w-3.5 h-3.5" />
                              <span>DNI: {user.dni}</span>
                            </div>
                          </div>
                        </div>
                      </TableCell>

                      <TableCell>
                        <div className="flex flex-wrap gap-3">
                          {assignments.map((a, idx) => (
                            <div key={`${a.application.id}-${a.role.id}-${idx}`} className="flex items-center gap-2 px-3 py-2 border rounded-sm">
                              <div className="flex flex-col leading-tight">
                                <div className="flex items-center gap-2 font-semibold text-base">
                                  <Boxes className="w-4 h-4 text-primary" />
                                  <span>{a.application.name}</span>
                                </div>

                                <div className="flex items-center gap-1 mt-1 text-muted-foreground text-xs">
                                  <Shield className="w-3 h-3" />
                                  <span className="capitalize">{a.role.name}</span>
                                </div>
                              </div>
                            </div>
                          ))}

                          {assignments.length === 0 && <span className="text-muted-foreground text-sm">— Sin roles asignados —</span>}
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
                          <Button variant="ghost" size="sm" className="text-red-600 hover:text-red-700" title="Eliminar asignaciones">
                            <Trash2 className="w-4 h-4" />
                          </Button>
                        </div>
                      </TableCell>
                    </TableRow>
                  );
                })}

                {filteredRows.length === 0 && (
                  <TableRow>
                    <TableCell colSpan={3} className="text-muted-foreground text-center">
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
