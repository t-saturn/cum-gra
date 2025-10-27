'use client';

import { useCallback, useEffect, useState } from 'react';
import { UsersListResponse, UserListItem } from '@/types/users';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Badge } from '@/components/ui/badge';
import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table';
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuLabel, DropdownMenuSeparator, DropdownMenuTrigger } from '@/components/ui/dropdown-menu';
import { Search, Plus, Filter, Download, MoreHorizontal, Edit, Trash2, Shield, Eye, Loader2, Phone, Hexagon, Timer } from 'lucide-react';
import { fn_get_users } from '@/actions/users/fn_get_users';
import { UsersStatsCards } from '@/components/custom/card/users-stats-cards';
import { CreateUserDialog, UserActionMenu } from './options';

export default function UsersManagement() {
  const [searchTerm, setSearchTerm] = useState('');
  const [users, setUsers] = useState<UserListItem[]>([]);
  const [loading, setLoading] = useState(true);
  const [total, setTotal] = useState(0);

  const refetch = useCallback(async (page: number = 1, pageSize: number = 20) => {
    try {
      setLoading(true);
      const response: UsersListResponse = await fn_get_users(page, pageSize);
      setUsers(response.data);
      setTotal(response.total);
    } catch (err) {
      console.error('Error recargando usuarios:', err);
    } finally {
      setLoading(false);
    }
  }, []);

  useEffect(() => {
    const fetchUsers = async () => {
      try {
        setLoading(true);
        const response: UsersListResponse = await fn_get_users(1, 20);
        setUsers(response.data);
        setTotal(response.total);
      } catch (err) {
        console.error('Error cargando usuarios:', err);
      } finally {
        setLoading(false);
      }
    };
    fetchUsers();
  }, []);

  const filteredUsers = users.filter(
    (user) =>
      user.email.toLowerCase().includes(searchTerm.toLowerCase()) ||
      user.first_name?.toLowerCase().includes(searchTerm.toLowerCase()) ||
      user.last_name?.toLowerCase().includes(searchTerm.toLowerCase()) ||
      user.dni.includes(searchTerm),
  );

  const getStatusBadge = (status: string) => {
    switch (status) {
      case 'active':
        return <Badge className="bg-chart-4/20 border-chart-4/30 text-chart-2">Activo</Badge>;
      case 'inactive':
        return <Badge className="bg-chart-5/20 border-chart-5/30 text-chart-1">Inactivo</Badge>;
      case 'suspended':
        return <Badge className="bg-yellow-500/20 border-yellow-500/30 text-yellow-600">Suspendido</Badge>;
      default:
        return <Badge variant="secondary">Desconocido</Badge>;
    }
  };

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <div>
          <h1 className="font-bold text-foreground text-3xl">Gestión de Usuarios</h1>
          <p className="mt-1 text-muted-foreground">Administra todos los usuarios del sistema</p>
        </div>
        <div className="flex gap-3">
          <Button variant="outline">
            <Download className="mr-2 w-4 h-4" />
            Exportar
          </Button>

          <CreateUserDialog onCreated={refetch} />
        </div>
      </div>

      <UsersStatsCards />

      <Card className="bg-card/50 border-border">
        <CardHeader>
          <div className="flex justify-between items-center">
            <div>
              <CardTitle>Lista de Usuarios</CardTitle>
              <CardDescription>{loading ? 'Cargando usuarios...' : `${filteredUsers.length} de ${total} usuarios`}</CardDescription>
            </div>
            <div className="flex gap-2">
              <div className="relative">
                <Search className="top-1/2 left-3 absolute w-4 h-4 text-muted-foreground -translate-y-1/2 transform" />
                <Input
                  placeholder="Buscar usuarios..."
                  value={searchTerm}
                  onChange={(e) => setSearchTerm(e.target.value)}
                  className="bg-background/50 pl-10 border-border focus:border-primary focus:ring-ring w-80"
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
          {loading ? (
            <div className="flex justify-center items-center py-16">
              <Loader2 className="w-8 h-8 text-primary animate-spin" />
            </div>
          ) : (
            <div className="border border-border rounded-lg">
              <Table>
                <TableHeader>
                  <TableRow className="bg-accent/50">
                    <TableHead>Usuario</TableHead>
                    <TableHead>DNI</TableHead>
                    <TableHead>Teléfono</TableHead>
                    <TableHead>Estado</TableHead>
                    <TableHead>Unidad Orgánica</TableHead>
                    <TableHead>Cargo</TableHead>
                    <TableHead>Creado</TableHead>
                    <TableHead className="text-right">Acciones</TableHead>
                  </TableRow>
                </TableHeader>
                <TableBody>
                  {filteredUsers.map((user) => (
                    <TableRow key={user.id} className="hover:bg-accent/30">
                      <TableCell>
                        <div className="flex items-center gap-3">
                          <Avatar className="w-10 h-10">
                            <AvatarImage src={`/placeholder.svg?height=40&width=40`} />
                            <AvatarFallback className="bg-gradient-to-r from-primary to-chart-1 font-semibold text-primary-foreground">
                              {user.first_name?.[0] ?? '?'}
                              {user.last_name?.[0] ?? ''}
                            </AvatarFallback>
                          </Avatar>
                          <div>
                            <p className="font-medium text-foreground">
                              {user.first_name} {user.last_name}
                            </p>
                            <p className="text-muted-foreground text-sm">{user.email}</p>
                          </div>
                        </div>
                      </TableCell>
                      <TableCell className="font-mono text-sm">{user.dni}</TableCell>
                      <TableCell className="text-sm">
                        <div className="flex items-center gap-2">
                          {user.phone ? (
                            <>
                              <Phone className="w-4 h-4 text-chart-2" />
                              {user.phone}
                            </>
                          ) : (
                            <span className="text-muted-foreground">—</span>
                          )}
                        </div>
                      </TableCell>
                      <TableCell>{getStatusBadge(user.status)}</TableCell>
                      <TableCell className="text-sm">{user.organic_unit?.acronym ?? '—'}</TableCell>
                      <TableCell className="text-sm">
                        <div className="flex items-center gap-2">
                          {user.structural_position?.name ? (
                            <>
                              <Hexagon className="w-4 h-4 text-chart-1" />
                              {user.structural_position.name}
                            </>
                          ) : (
                            <span className="text-muted-foreground">—</span>
                          )}
                        </div>
                      </TableCell>
                      <TableCell className="text-sm">
                        <div className="flex items-center gap-2">
                          {user.created_at ? (
                            <>
                              <Timer className="w-4 h-4 text-chart-3" />
                              {new Date(user.created_at).toLocaleDateString('es-ES')}
                            </>
                          ) : (
                            <span className="text-muted-foreground">—</span>
                          )}
                        </div>
                      </TableCell>
                      <TableCell className="text-right">
                        <UserActionMenu user={user} onRefresh={refetch} />
                      </TableCell>
                    </TableRow>
                  ))}
                </TableBody>
              </Table>
            </div>
          )}
        </CardContent>
      </Card>
    </div>
  );
}
