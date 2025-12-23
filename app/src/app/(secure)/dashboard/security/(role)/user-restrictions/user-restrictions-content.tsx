'use client';

import { useState, useEffect, useTransition } from 'react';
import { useRouter, useSearchParams } from 'next/navigation';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table';
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuLabel, DropdownMenuSeparator, DropdownMenuTrigger } from '@/components/ui/dropdown-menu';
import { Badge } from '@/components/ui/badge';
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle } from '@/components/ui/dialog';
import { Plus, MoreHorizontal, Edit, Trash2, RotateCcw, Filter, Layers } from 'lucide-react';
import { toast } from 'sonner';

import { fn_get_applications } from '@/actions/applications/fn_get_applications';

import RestrictionModal from './restriction-modal';
import BulkRestrictionModal from './bulk-restriction-modal';
import type { UserRestrictionItem } from '@/types/user-restrictions';
import { fn_get_user_restrictions } from '@/actions/users-restrictions/fn_get_users_restrictions';
import { fn_delete_user_restriction, fn_restore_user_restriction } from '@/actions/users-restrictions/fn_delete_user_restriction';
import { UserRestrictionsStatsCards } from '@/components/custom/card/users-restrictions-stats-cards';

export default function UserRestrictionsContent() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const [isPending, startTransition] = useTransition();

  const page = Number(searchParams.get('page')) || 1;
  const showDeleted = searchParams.get('deleted') === 'true';

  const [restrictions, setRestrictions] = useState<UserRestrictionItem[]>([]);
  const [applications, setApplications] = useState<any[]>([]);
  const [selectedApp, setSelectedApp] = useState('');
  const [loading, setLoading] = useState(false);
  const [total, setTotal] = useState(0);

  // Modals
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [isBulkModalOpen, setIsBulkModalOpen] = useState(false);
  const [isDeleteOpen, setIsDeleteOpen] = useState(false);
  const [selectedItem, setSelectedItem] = useState<UserRestrictionItem | null>(null);

  // Cargar Apps
  useEffect(() => {
    fn_get_applications(1, 100, false).then(res => {
        setApplications(res.data);
        if (res.data.length > 0 && !selectedApp) setSelectedApp(res.data[0].id);
    });
  }, []);

  const loadData = async () => {
    if (!selectedApp) return;
    setLoading(true);
    try {
      const res = await fn_get_user_restrictions(page, 10, { 
        application_id: selectedApp, 
        is_deleted: showDeleted 
      });
      setRestrictions(res.data);
      setTotal(res.total);
    } catch {
      toast.error('Error cargando restricciones');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    if (selectedApp) loadData();
  }, [page, showDeleted, selectedApp]);

  const handleAppChange = (val: string) => {
    setSelectedApp(val);
    router.replace('?page=1'); // Reset page
  };

  const toggleDeleted = () => {
    const params = new URLSearchParams(searchParams);
    if (showDeleted) params.delete('deleted');
    else params.set('deleted', 'true');
    router.replace(`?${params.toString()}`);
  };

  const handleDelete = async () => {
    if (!selectedItem) return;
    try {
      await fn_delete_user_restriction(selectedItem.id);
      toast.success('Eliminado correctamente');
      setIsDeleteOpen(false);
      loadData();
    } catch (e: any) {
      toast.error(e.message);
    }
  };

  const handleRestore = async (id: string) => {
    try {
      await fn_restore_user_restriction(id);
      toast.success('Restaurado correctamente');
      loadData();
    } catch (e: any) {
      toast.error(e.message);
    }
  };

  const getBadgeColor = (type: string) => {
    switch (type) {
        case 'block': return 'destructive';
        case 'read_only': return 'secondary';
        case 'limit': return 'outline';
        default: return 'default';
    }
  };

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <div>
          <h1 className="text-3xl font-bold text-foreground">Restricciones de Usuario</h1>
          <p className="text-muted-foreground mt-1">Controla excepciones y bloqueos específicos por módulo.</p>
        </div>
        <div className="flex gap-2">
            <Button variant="outline" onClick={() => setIsBulkModalOpen(true)} disabled={!selectedApp}>
                <Layers className="mr-2 h-4 w-4" /> Masivo
            </Button>
            <Button onClick={() => { setSelectedItem(null); setIsModalOpen(true); }} className="bg-linear-to-r from-primary to-chart-1" disabled={!selectedApp}>
                <Plus className="mr-2 h-4 w-4" /> Nueva
            </Button>
        </div>
      </div>

      <UserRestrictionsStatsCards />

      <Card className="bg-card/50 border-border">
        <CardHeader>
            <div className="flex justify-between items-center">
                <div className="flex items-center gap-4">
                    <CardTitle>Listado</CardTitle>
                    <div className="flex items-center gap-2 bg-muted/50 p-1 rounded-lg border">
                        <Layers className="w-4 h-4 ml-2 text-primary" />
                        <Select value={selectedApp} onValueChange={handleAppChange}>
                            <SelectTrigger className="w-[200px] border-none shadow-none h-8 bg-transparent focus:ring-0"><SelectValue placeholder="App..." /></SelectTrigger>
                            <SelectContent>{applications.map(a => <SelectItem key={a.id} value={a.id}>{a.name}</SelectItem>)}</SelectContent>
                        </Select>
                    </div>
                </div>
                <Button variant="ghost" size="sm" onClick={toggleDeleted} className={showDeleted ? 'text-destructive bg-destructive/10' : ''}>
                    <Filter className="mr-2 h-4 w-4" /> {showDeleted ? 'Ver Activos' : 'Ver Eliminados'}
                </Button>
            </div>
        </CardHeader>
        <CardContent>
            {!selectedApp ? (
                <div className="text-center py-10 text-muted-foreground">Selecciona una aplicación para ver sus restricciones</div>
            ) : (
                <Table>
                    <TableHeader>
                        <TableRow className="bg-accent/50">
                            <TableHead>Usuario</TableHead>
                            <TableHead>Módulo</TableHead>
                            <TableHead>Tipo</TableHead>
                            <TableHead>Razón</TableHead>
                            <TableHead>Estado</TableHead>
                            <TableHead className="text-right">Acciones</TableHead>
                        </TableRow>
                    </TableHeader>
                    <TableBody>
                        {loading ? <TableRow><TableCell colSpan={6} className="text-center py-8">Cargando...</TableCell></TableRow> :
                         restrictions.length === 0 ? <TableRow><TableCell colSpan={6} className="text-center py-8">No hay registros</TableCell></TableRow> :
                         restrictions.map(r => (
                            <TableRow key={r.id} className={r.is_deleted ? 'opacity-50' : ''}>
                                <TableCell>
                                    <div className="font-medium">{r.user_full_name}</div>
                                    <div className="text-xs text-muted-foreground">{r.user_email}</div>
                                </TableCell>
                                <TableCell>{r.module_name}</TableCell>
                                <TableCell>
                                    <Badge variant={getBadgeColor(r.restriction_type) as any}>
                                        {r.restriction_type === 'limit' ? `Limit: ${r.max_permission_level}` : r.restriction_type}
                                    </Badge>
                                </TableCell>
                                <TableCell className="max-w-[200px] truncate" title={r.reason || ''}>{r.reason || '-'}</TableCell>
                                <TableCell>
                                    {r.expires_at && new Date(r.expires_at) < new Date() ? 
                                        <Badge variant="outline" className="text-orange-500 border-orange-200">Expirado</Badge> : 
                                        r.is_deleted ? <Badge variant="destructive">Eliminado</Badge> : 
                                        <Badge className="bg-green-500/10 text-green-500 border-green-500/20">Activo</Badge>
                                    }
                                </TableCell>
                                <TableCell className="text-right">
                                    <DropdownMenu>
                                        <DropdownMenuTrigger asChild><Button variant="ghost" size="sm"><MoreHorizontal className="h-4 w-4" /></Button></DropdownMenuTrigger>
                                        <DropdownMenuContent align="end">
                                            <DropdownMenuLabel>Acciones</DropdownMenuLabel>
                                            <DropdownMenuSeparator />
                                            {!r.is_deleted ? (
                                                <>
                                                    <DropdownMenuItem onClick={() => { setSelectedItem(r); setIsModalOpen(true); }}>
                                                        <Edit className="mr-2 h-4 w-4" /> Editar
                                                    </DropdownMenuItem>
                                                    <DropdownMenuItem className="text-destructive focus:text-destructive" onClick={() => { setSelectedItem(r); setIsDeleteOpen(true); }}>
                                                        <Trash2 className="mr-2 h-4 w-4" /> Eliminar
                                                    </DropdownMenuItem>
                                                </>
                                            ) : (
                                                <DropdownMenuItem onClick={() => handleRestore(r.id)}>
                                                    <RotateCcw className="mr-2 h-4 w-4" /> Restaurar
                                                </DropdownMenuItem>
                                            )}
                                        </DropdownMenuContent>
                                    </DropdownMenu>
                                </TableCell>
                            </TableRow>
                         ))
                        }
                    </TableBody>
                </Table>
            )}
        </CardContent>
      </Card>

      {/* Aquí pasamos defaultApplicationId */}
      <RestrictionModal 
        open={isModalOpen} 
        onOpenChange={setIsModalOpen} 
        restriction={selectedItem} 
        onSuccess={loadData} 
        defaultApplicationId={selectedApp} 
      />

      <BulkRestrictionModal
        open={isBulkModalOpen}
        onOpenChange={setIsBulkModalOpen}
        onSuccess={loadData}
        defaultApplicationId={selectedApp}
      />

      <Dialog open={isDeleteOpen} onOpenChange={setIsDeleteOpen}>
        <DialogContent>
            <DialogHeader>
                <DialogTitle>¿Eliminar Restricción?</DialogTitle>
                <DialogDescription>El usuario recuperará el acceso normal al módulo.</DialogDescription>
            </DialogHeader>
            <DialogFooter>
                <Button variant="outline" onClick={() => setIsDeleteOpen(false)}>Cancelar</Button>
                <Button variant="destructive" onClick={handleDelete}>Eliminar</Button>
            </DialogFooter>
        </DialogContent>
      </Dialog>
    </div>
  );
}