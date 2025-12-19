'use client';

import { useState } from 'react';
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle } from '@/components/ui/dialog';
import { Button } from '@/components/ui/button';
import { Badge } from '@/components/ui/badge';
import { ScrollArea } from '@/components/ui/scroll-area';
import { toast } from 'sonner';
import { fn_create_application } from '@/actions/applications/fn_create_application';
import type { KeycloakClientSimple } from '@/actions/applications/fn_get_keycloak_clients';
import { RefreshCw, Check, Boxes } from 'lucide-react';

interface SyncKeycloakModalProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  unsyncedClients: KeycloakClientSimple[];
  onSuccess: () => void;
}

export default function SyncKeycloakModal({ open, onOpenChange, unsyncedClients, onSuccess }: SyncKeycloakModalProps) {
  const [syncing, setSyncing] = useState<string | null>(null);
  const [syncedClients, setSyncedClients] = useState<Set<string>>(new Set());

  const handleSyncClient = async (client: KeycloakClientSimple) => {
    setSyncing(client.client_id);

    try {
      // Crear aplicación en el backend usando datos de Keycloak
      await fn_create_application({
        name: client.name,
        client_id: client.client_id,
        domain: client.base_url || 'https://app.regionayacucho.gob.pe',
        description: client.description || `Aplicación sincronizada desde Keycloak (${client.protocol})`,
        status: 'development',
        client_secret: 'sync-from-keycloak', // El backend debería generar uno o dejarlo vacío
      });

      setSyncedClients(prev => new Set(prev).add(client.client_id));
      toast.success(`Cliente ${client.name} sincronizado correctamente`);
    } catch (error: any) {
      toast.error(error.message || `Error al sincronizar ${client.name}`);
    } finally {
      setSyncing(null);
    }
  };

  const handleSyncAll = async () => {
    for (const client of unsyncedClients) {
      if (!syncedClients.has(client.client_id)) {
        await handleSyncClient(client);
      }
    }
    toast.success('Sincronización completada');
    onSuccess();
    onOpenChange(false);
  };

  const getProtocolBadge = (protocol: string) => {
    switch (protocol) {
      case 'openid-connect':
        return <Badge variant="outline" className="text-xs">OAuth 2.0</Badge>;
      case 'saml':
        return <Badge variant="outline" className="text-xs">SAML</Badge>;
      default:
        return <Badge variant="outline" className="text-xs">{protocol}</Badge>;
    }
  };

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="sm:max-w-[700px] max-h-[80vh]">
        <DialogHeader>
          <DialogTitle>Sincronizar Clientes de Keycloak</DialogTitle>
          <DialogDescription>
            Se encontraron {unsyncedClients.length} cliente(s) en Keycloak que no están registrados en el sistema.
            Selecciona los clientes que deseas sincronizar.
          </DialogDescription>
        </DialogHeader>

        <ScrollArea className="max-h-[400px] pr-4">
          <div className="space-y-3">
            {unsyncedClients.map((client) => {
              const isSynced = syncedClients.has(client.client_id);
              const isSyncing = syncing === client.client_id;

              return (
                <div
                  key={client.client_id}
                  className="flex items-start justify-between p-4 border rounded-lg bg-card hover:bg-accent/50 transition"
                >
                  <div className="flex-1 space-y-1">
                    <div className="flex items-center gap-2">
                      <Boxes className="w-4 h-4 text-primary" />
                      <p className="font-medium">{client.name}</p>
                      {getProtocolBadge(client.protocol)}
                    </div>
                    <p className="text-sm text-muted-foreground">
                      Client ID: <code className="text-xs bg-muted px-1 py-0.5 rounded">{client.client_id}</code>
                    </p>
                    {client.description && (
                      <p className="text-sm text-muted-foreground line-clamp-2">{client.description}</p>
                    )}
                    {client.base_url && (
                      <p className="text-xs text-muted-foreground">{client.base_url}</p>
                    )}
                  </div>

                  <Button
                    size="sm"
                    variant={isSynced ? 'outline' : 'default'}
                    onClick={() => handleSyncClient(client)}
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

            {unsyncedClients.length === 0 && (
              <div className="flex flex-col items-center justify-center py-8 text-center">
                <Check className="w-12 h-12 text-green-500 mb-2" />
                <p className="text-muted-foreground">Todos los clientes están sincronizados</p>
              </div>
            )}
          </div>
        </ScrollArea>

        <DialogFooter>
          <Button variant="outline" onClick={() => onOpenChange(false)} disabled={!!syncing}>
            Cerrar
          </Button>
          {unsyncedClients.length > 0 && syncedClients.size < unsyncedClients.length && (
            <Button 
              onClick={handleSyncAll} 
              disabled={!!syncing}
              className="bg-linear-to-r from-primary to-chart-1"
            >
              <RefreshCw className="w-4 h-4 mr-2" />
              Sincronizar Todos
            </Button>
          )}
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}