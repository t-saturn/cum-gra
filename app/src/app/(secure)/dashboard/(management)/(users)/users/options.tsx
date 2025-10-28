'use client';

import { useState } from 'react';
import { motion } from 'framer-motion';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle, DialogTrigger } from '@/components/ui/dialog';
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from '@/components/ui/alert-dialog';
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuLabel, DropdownMenuSeparator, DropdownMenuTrigger } from '@/components/ui/dropdown-menu';
import { Badge } from '@/components/ui/badge';
import { Sheet, SheetContent, SheetDescription, SheetHeader, SheetTitle } from '@/components/ui/sheet';
import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar';
import { Eye, Edit, Trash2, ArchiveRestore, Ban, Plus, Loader2 } from 'lucide-react';

export type MinimalOrganicUnit = { id: string; acronym: string; name: string };
export type MinimalStructuralPosition = { id: string; name: string };
export type UserListItem = {
  id: string;
  email: string;
  first_name?: string | null;
  last_name?: string | null;
  phone?: string | null;
  dni: string;
  status: 'active' | 'inactive' | 'suspended' | string;
  organic_unit?: MinimalOrganicUnit | null;
  structural_position?: MinimalStructuralPosition | null;
  created_at?: string;
};

// —— Placeholders de Server Actions (reemplaza por tus implementaciones reales) ——
// import { fn_create_user } from "@/actions/users/fn_create_user";
// import { fn_update_user } from "@/actions/users/fn_update_user";
// import { fn_deactivate_user } from "@/actions/users/fn_deactivate_user";
// import { fn_activate_user } from "@/actions/users/fn_activate_user";
// import { fn_delete_user } from "@/actions/users/fn_delete_user";

async function fakeDelay<T>(v: T, ms = 650) {
  return new Promise<T>((res) => setTimeout(() => res(v), ms));
}
const fn_create_user = async (_: any) => fakeDelay({ ok: true });
const fn_update_user = async (_: any) => fakeDelay({ ok: true });
const fn_deactivate_user = async (_: any) => fakeDelay({ ok: true });
const fn_activate_user = async (_: any) => fakeDelay({ ok: true });
const fn_delete_user = async (_: any) => fakeDelay({ ok: true });

function cx(...cn: (string | false | undefined)[]) {
  return cn.filter(Boolean).join(' ');
}

export function CreateUserDialog({ onCreated }: { onCreated?: () => void }) {
  const [open, setOpen] = useState(false);
  const [loading, setLoading] = useState(false);
  const [form, setForm] = useState({
    first_name: '',
    last_name: '',
    email: '',
    dni: '',
    phone: '',
  });

  const submit = async () => {
    try {
      setLoading(true);
      if (!form.email || !form.dni) {
        console.log('Email y DNI son obligatorios');
        return;
      }
      const res = await fn_create_user(form);
      if ((res as any)?.ok) {
        console.log('Usuario creado');
        setOpen(false);
        onCreated?.();
      } else {
        console.log('No se pudo crear el usuario');
      }
    } catch (err) {
      console.error(err);
      console.log('Error creando usuario');
    } finally {
      setLoading(false);
    }
  };

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>
        <Button className="bg-gradient-to-r from-primary hover:from-primary/90 to-chart-1 hover:to-chart-1/90 shadow-lg shadow-primary/25">
          <Plus className="mr-2 w-4 h-4" /> Nuevo Usuario
        </Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-xl">
        <DialogHeader>
          <DialogTitle>Crear usuario</DialogTitle>
          <DialogDescription>Completa los campos para registrar un nuevo usuario.</DialogDescription>
        </DialogHeader>
        <div className="gap-4 grid py-4">
          <div className="gap-3 grid grid-cols-2">
            <div>
              <Label>Nombre</Label>
              <Input value={form.first_name} onChange={(e) => setForm((v) => ({ ...v, first_name: e.target.value }))} placeholder="Ana" />
            </div>
            <div>
              <Label>Apellido</Label>
              <Input value={form.last_name} onChange={(e) => setForm((v) => ({ ...v, last_name: e.target.value }))} placeholder="García" />
            </div>
          </div>
          <div className="gap-3 grid grid-cols-2">
            <div>
              <Label>Email</Label>
              <Input type="email" value={form.email} onChange={(e) => setForm((v) => ({ ...v, email: e.target.value }))} placeholder="ana@empresa.com" />
            </div>
            <div>
              <Label>DNI</Label>
              <Input value={form.dni} onChange={(e) => setForm((v) => ({ ...v, dni: e.target.value }))} placeholder="12345678" />
            </div>
          </div>
          <div>
            <Label>Teléfono</Label>
            <Input value={form.phone} onChange={(e) => setForm((v) => ({ ...v, phone: e.target.value }))} placeholder="999-999-999" />
          </div>
        </div>
        <DialogFooter>
          <Button variant="outline" onClick={() => setOpen(false)}>
            Cancelar
          </Button>
          <Button onClick={submit} disabled={loading}>
            {loading ? <Loader2 className="mr-2 w-4 h-4 animate-spin" /> : null}
            Guardar
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}

export function EditUserDialog({
  user,
  onUpdated,
  trigger,
  open,
  onOpenChange,
}: {
  user: UserListItem;
  onUpdated?: () => void;
  trigger?: React.ReactNode;
  open?: boolean;
  onOpenChange?: (v: boolean) => void;
}) {
  const [internalOpen, setInternalOpen] = useState(false);
  const isControlled = typeof open === 'boolean';
  const openState = isControlled ? open! : internalOpen;
  const setOpen = isControlled ? onOpenChange! : setInternalOpen;

  const [loading, setLoading] = useState(false);
  const [form, setForm] = useState({
    first_name: user.first_name ?? '',
    last_name: user.last_name ?? '',
    email: user.email,
    dni: user.dni,
    phone: user.phone ?? '',
  });

  const submit = async () => {
    try {
      setLoading(true);
      const res = await fn_update_user({ id: user.id, ...form });
      if ((res as any)?.ok) {
        console.log('Usuario actualizado');
        setOpen(false);
        onUpdated?.();
      } else {
        console.log('No se pudo actualizar');
      }
    } catch (e) {
      console.error(e);
      console.log('Error actualizando');
    } finally {
      setLoading(false);
    }
  };

  return (
    <Dialog open={openState} onOpenChange={setOpen}>
      {trigger ? <DialogTrigger asChild>{trigger}</DialogTrigger> : null}
      <DialogContent className="sm:max-w-xl">
        <DialogHeader>
          <DialogTitle>Editar usuario</DialogTitle>
          <DialogDescription>Actualiza la información del usuario.</DialogDescription>
        </DialogHeader>
        <div className="gap-4 grid py-4">
          <div className="gap-3 grid grid-cols-2">
            <div>
              <Label>Nombre</Label>
              <Input value={form.first_name} onChange={(e) => setForm((v) => ({ ...v, first_name: e.target.value }))} />
            </div>
            <div>
              <Label>Apellido</Label>
              <Input value={form.last_name} onChange={(e) => setForm((v) => ({ ...v, last_name: e.target.value }))} />
            </div>
          </div>
          <div className="gap-3 grid grid-cols-2">
            <div>
              <Label>Email</Label>
              <Input type="email" value={form.email} onChange={(e) => setForm((v) => ({ ...v, email: e.target.value }))} />
            </div>
            <div>
              <Label>DNI</Label>
              <Input value={form.dni} onChange={(e) => setForm((v) => ({ ...v, dni: e.target.value }))} />
            </div>
          </div>
          <div>
            <Label>Teléfono</Label>
            <Input value={form.phone} onChange={(e) => setForm((v) => ({ ...v, phone: e.target.value }))} />
          </div>
        </div>
        <DialogFooter>
          <Button variant="outline" onClick={() => setOpen(false)}>
            Cancelar
          </Button>
          <Button onClick={submit} disabled={loading}>
            {loading ? <Loader2 className="mr-2 w-4 h-4 animate-spin" /> : null}
            Guardar cambios
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}

export function StatusBadge({ value }: { value: string }) {
  switch (value) {
    case 'active':
      return <Badge className="bg-chart-4/20 border-chart-4/30 text-chart-2">Activo</Badge>;
    case 'inactive':
      return <Badge className="bg-chart-5/20 border-chart-5/30 text-chart-1">Inactivo</Badge>;
    case 'suspended':
      return <Badge className="bg-yellow-500/20 border-yellow-500/30 text-yellow-600">Suspendido</Badge>;
    default:
      return <Badge variant="secondary">Desconocido</Badge>;
  }
}

function Info({ label, value, className }: { label: string; value: React.ReactNode; className?: string }) {
  return (
    <div className={cx('rounded-xl border border-border p-3', className)}>
      <div className="text-muted-foreground text-xs">{label}</div>
      <div className="mt-1 text-sm">{value}</div>
    </div>
  );
}

export function UserActionMenu({ user, onRefresh }: { user: UserListItem; onRefresh?: () => void }) {
  const [viewOpen, setViewOpen] = useState(false);
  const [editOpen, setEditOpen] = useState(false);
  const [deactivateOpen, setDeactivateOpen] = useState(false);
  const [activateOpen, setActivateOpen] = useState(false);
  const [deleteOpen, setDeleteOpen] = useState(false);
  const [loadingAction, setLoadingAction] = useState(false);

  const runDeactivate = async () => {
    try {
      setLoadingAction(true);
      const res = await fn_deactivate_user({ id: user.id });
      if ((res as any)?.ok) {
        console.log('Usuario dado de baja');
        onRefresh?.();
      } else {
        console.log('No se pudo dar de baja');
      }
    } finally {
      setLoadingAction(false);
      setDeactivateOpen(false);
    }
  };

  const runActivate = async () => {
    try {
      setLoadingAction(true);
      const res = await fn_activate_user({ id: user.id });
      if ((res as any)?.ok) {
        console.log('Usuario reactivado');
        onRefresh?.();
      } else {
        console.log('No se pudo reactivar');
      }
    } finally {
      setLoadingAction(false);
      setActivateOpen(false);
    }
  };

  const runDelete = async () => {
    try {
      setLoadingAction(true);
      const res = await fn_delete_user({ id: user.id });
      if ((res as any)?.ok) {
        console.log('Usuario eliminado');
        onRefresh?.();
      } else {
        console.log('No se pudo eliminar');
      }
    } finally {
      setLoadingAction(false);
      setDeleteOpen(false);
    }
  };

  return (
    <>
      <DropdownMenu>
        <DropdownMenuTrigger asChild>
          <Button variant="ghost" size="sm">
            <MoreIcon />
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent align="end" className="bg-card/80 backdrop-blur-xl border-border">
          <DropdownMenuLabel>Acciones</DropdownMenuLabel>
          <DropdownMenuSeparator />

          {/* Ver Detalles */}
          <DropdownMenuItem
            className="cursor-pointer"
            onSelect={(e) => {
              e.preventDefault();
              setViewOpen(true);
            }}
          >
            <Eye className="mr-2 w-4 h-4" /> Ver Detalles
          </DropdownMenuItem>

          {/* Editar */}
          <DropdownMenuItem
            className="cursor-pointer"
            onSelect={(e) => {
              e.preventDefault();
              setEditOpen(true);
            }}
          >
            <Edit className="mr-2 w-4 h-4" /> Editar
          </DropdownMenuItem>

          <DropdownMenuSeparator />

          {user.status === 'active' ? (
            <DropdownMenuItem
              className="text-yellow-600 cursor-pointer"
              onSelect={(e) => {
                e.preventDefault();
                setDeactivateOpen(true);
              }}
            >
              <Ban className="mr-2 w-4 h-4" /> Dar de baja
            </DropdownMenuItem>
          ) : (
            <DropdownMenuItem
              className="text-emerald-600 cursor-pointer"
              onSelect={(e) => {
                e.preventDefault();
                setActivateOpen(true);
              }}
            >
              <ArchiveRestore className="mr-2 w-4 h-4" /> Dar de alta
            </DropdownMenuItem>
          )}

          <DropdownMenuItem
            className="text-destructive cursor-pointer"
            onSelect={(e) => {
              e.preventDefault();
              setDeleteOpen(true);
            }}
          >
            <Trash2 className="mr-2 w-4 h-4" /> Eliminar
          </DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>

      <Sheet open={viewOpen} onOpenChange={setViewOpen}>
        <SheetContent className="sm:max-w-xl">
          <SheetHeader>
            <SheetTitle>Información del usuario</SheetTitle>
            <SheetDescription>Revisa los datos principales antes de realizar cambios.</SheetDescription>
          </SheetHeader>
          <div className="space-y-4 mt-4">
            <div className="flex items-center gap-3">
              <Avatar className="w-12 h-12">
                <AvatarImage src={`/placeholder.svg?height=48&width=48`} />
                <AvatarFallback className="bg-gradient-to-r from-primary to-chart-1 font-semibold text-primary-foreground">
                  {(user.first_name?.[0] ?? '?') + (user.last_name?.[0] ?? '')}
                </AvatarFallback>
              </Avatar>
              <div>
                <div className="font-semibold">
                  {user.first_name} {user.last_name}
                </div>
                <div className="text-muted-foreground text-sm">{user.email}</div>
              </div>
            </div>
            <div className="gap-3 grid grid-cols-2">
              <Info label="DNI" value={user.dni} />
              <Info label="Teléfono" value={user.phone ?? '—'} />
              <Info label="Estado" value={<StatusBadge value={user.status} />} />
              <Info label="Unidad orgánica" value={user.organic_unit?.acronym ?? '—'} />
              <Info label="Cargo" value={user.structural_position?.name ?? '—'} className="col-span-2" />
            </div>
          </div>
        </SheetContent>
      </Sheet>

      <EditUserDialog user={user} onUpdated={onRefresh} open={editOpen} onOpenChange={setEditOpen} />

      <AlertDialog open={deactivateOpen} onOpenChange={setDeactivateOpen}>
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>Dar de baja al usuario</AlertDialogTitle>
            <AlertDialogDescription>El usuario no podrá iniciar sesión ni operar en el sistema. Puedes revertir esta acción reactivándolo.</AlertDialogDescription>
          </AlertDialogHeader>
          <AlertDialogFooter>
            <AlertDialogCancel>Cancelar</AlertDialogCancel>
            <AlertDialogAction onClick={runDeactivate} disabled={loadingAction}>
              {loadingAction ? <Loader2 className="mr-2 w-4 h-4 animate-spin" /> : <Ban className="mr-2 w-4 h-4" />}
              Confirmar baja
            </AlertDialogAction>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>

      <AlertDialog open={activateOpen} onOpenChange={setActivateOpen}>
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>Reactivar usuario</AlertDialogTitle>
            <AlertDialogDescription>Esto restaurará el acceso del usuario al sistema.</AlertDialogDescription>
          </AlertDialogHeader>
          <AlertDialogFooter>
            <AlertDialogCancel>Cancelar</AlertDialogCancel>
            <AlertDialogAction onClick={runActivate} disabled={loadingAction}>
              {loadingAction ? <Loader2 className="mr-2 w-4 h-4 animate-spin" /> : <ArchiveRestore className="mr-2 w-4 h-4" />}
              Confirmar alta
            </AlertDialogAction>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>

      <AlertDialog open={deleteOpen} onOpenChange={setDeleteOpen}>
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>¿Eliminar definitivamente?</AlertDialogTitle>
            <AlertDialogDescription>Esta acción no se puede deshacer. Se eliminará el usuario y sus relaciones no críticas.</AlertDialogDescription>
          </AlertDialogHeader>
          <AlertDialogFooter>
            <AlertDialogCancel>Cancelar</AlertDialogCancel>
            <AlertDialogAction onClick={runDelete} disabled={loadingAction} className="bg-destructive hover:bg-destructive/90">
              {loadingAction ? <Loader2 className="mr-2 w-4 h-4 animate-spin" /> : <Trash2 className="mr-2 w-4 h-4" />}
              Eliminar
            </AlertDialogAction>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>
    </>
  );
}

function MoreIcon() {
  return (
    <motion.svg
      xmlns="http://www.w3.org/2000/svg"
      viewBox="0 0 24 24"
      width="20"
      height="20"
      className="text-foreground"
      initial={{ rotate: 0 }}
      whileHover={{ rotate: 90 }}
      transition={{ type: 'spring', stiffness: 200, damping: 12 }}
    >
      <path fill="currentColor" d="M3 12a2 2 0 1 0 4 0a2 2 0 1 0 -4 0m7 0a2 2 0 1 0 4 0a2 2 0 1 0 -4 0m7 0a2 2 0 1 0 4 0a2 2 0 1 0 -4 0" />
    </motion.svg>
  );
}
