'use client';

import { useState } from 'react';
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle } from '@/components/ui/dialog';
import { Button } from '@/components/ui/button';
import { Badge } from '@/components/ui/badge';
import { ScrollArea } from '@/components/ui/scroll-area';
import { toast } from 'sonner';
import { fn_create_user } from '@/actions/users/fn_create_user';
import type { KeycloakUserSimple } from '@/actions/keycloak/users/fn_get_keycloak_users';
import { RefreshCw, Check, User } from 'lucide-react';

interface SyncKeycloakUsersModalProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  unsyncedUsers: KeycloakUserSimple[];
  onSuccess: () => void;
}

export default function SyncKeycloakUsersModal({ open, onOpenChange, unsyncedUsers, onSuccess }: SyncKeycloakUsersModalProps) {
  const [syncing, setSyncing] = useState<string | null>(null);
  const [syncedUsers, setSyncedUsers] = useState<Set<string>>(new Set());

  const handleSyncUser = async (kcUser: KeycloakUserSimple) => {
    setSyncing(kcUser.username);

    try {
      // Crear usuario en el backend usando datos de Keycloak
      await fn_create_user({
        email: kcUser.email || `${kcUser.username}@regionayacucho.gob.pe`,
        dni: kcUser.username,
        first_name: kcUser.firstName || 'Usuario',
        last_name: kcUser.lastName || 'Keycloak',
        status: kcUser.enabled ? 'active' : 'inactive',
        sync_to_keycloak: false, // No volver a crear en Keycloak
      });

      setSyncedUsers((prev) => new Set(prev).add(kcUser.username));
      toast.success(`Usuario ${kcUser.firstName} ${kcUser.lastName} sincronizado correctamente`);
    } catch (error: any) {
      toast.error(error.message || `Error al sincronizar ${kcUser.firstName} ${kcUser.lastName}`);
    } finally {
      setSyncing(null);
    }
  };

  const handleSyncAll = async () => {
    for (const kcUser of unsyncedUsers) {
      if (!syncedUsers.has(kcUser.username)) {
        await handleSyncUser(kcUser);
      }
    }
    toast.success('Sincronización completada');
    onSuccess();
    onOpenChange(false);
  };

  const getStatusBadge = (enabled: boolean) => {
    return enabled ? (
      <Badge variant="outline" className="text-xs bg-green-500/10 text-green-500">
        Habilitado
      </Badge>
    ) : (
      <Badge variant="outline" className="text-xs bg-red-500/10 text-red-500">
        Deshabilitado
      </Badge>
    );
  };

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="sm:max-w-[700px] max-h-[80vh]">
        <DialogHeader>
          <DialogTitle>Sincronizar Usuarios de Keycloak</DialogTitle>
          <DialogDescription>
            Se encontraron {unsyncedUsers.length} usuario(s) en Keycloak que no están registrados en el sistema. Selecciona los usuarios que deseas
            sincronizar.
          </DialogDescription>
        </DialogHeader>

        <ScrollArea className="max-h-[400px] pr-4">
          <div className="space-y-3">
            {unsyncedUsers.map((kcUser) => {
              const isSynced = syncedUsers.has(kcUser.username);
              const isSyncing = syncing === kcUser.username;

              return (
                <div key={kcUser.id} className="flex items-start justify-between p-4 border rounded-lg bg-card hover:bg-accent/50 transition">
                  <div className="flex-1 space-y-1">
                    <div className="flex items-center gap-2">
                      <User className="w-4 h-4 text-primary" />
                      <p className="font-medium">
                        {kcUser.firstName || 'Sin nombre'} {kcUser.lastName || ''}
                      </p>
                      {getStatusBadge(kcUser.enabled)}
                    </div>
                    <p className="text-sm text-muted-foreground">Username: {kcUser.username}</p>
                    {kcUser.email && <p className="text-sm text-muted-foreground">Email: {kcUser.email}</p>}
                    {kcUser.createdTimestamp && (
                      <p className="text-xs text-muted-foreground">
                        Creado: {new Date(kcUser.createdTimestamp).toLocaleDateString('es-PE')}
                      </p>
                    )}
                  </div>

                  <Button
                    size="sm"
                    variant={isSynced ? 'outline' : 'default'}
                    onClick={() => handleSyncUser(kcUser)}
                    disabled={isSyncing || isSynced}
                    className="ml-4"
                  >
                    {isSyncing ? (
                      <>
                        <RefreshCw className="w-4 h-4 mr-2 animate-spin" />
                        Sincronizando...
                      </>
                    ) : isSynced ? (
                      <>
                        <Check className="w-4 h-4 mr-2 text-green-500" />
                        Sincronizado
                      </>
                    ) : (
                      <>
                        <RefreshCw className="w-4 h-4 mr-2" />
                        Sincronizar
                      </>
                    )}
                  </Button>
                </div>
              );
            })}

            {unsyncedUsers.length === 0 && (
              <div className="flex flex-col items-center justify-center py-8 text-center">
                <Check className="w-12 h-12 text-green-500 mb-2" />
                <p className="text-muted-foreground">Todos los usuarios están sincronizados</p>
              </div>
            )}
          </div>
        </ScrollArea>

        <DialogFooter>
          <Button variant="outline" onClick={() => onOpenChange(false)} disabled={!!syncing}>
            Cerrar
          </Button>
          {unsyncedUsers.length > 0 && syncedUsers.size < unsyncedUsers.length && (
            <Button onClick={handleSyncAll} disabled={!!syncing} className="bg-linear-to-r from-primary to-chart-1">
              <RefreshCw className="w-4 h-4 mr-2" />
              Sincronizar Todos
            </Button>
          )}
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}