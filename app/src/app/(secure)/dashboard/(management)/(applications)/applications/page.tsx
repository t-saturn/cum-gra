'use client';

import { useEffect, useMemo, useState } from 'react';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuLabel, DropdownMenuSeparator, DropdownMenuTrigger } from '@/components/ui/dropdown-menu';
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle, DialogTrigger } from '@/components/ui/dialog';
import { Label } from '@/components/ui/label';
import { Textarea } from '@/components/ui/textarea';
import { Search, Plus, Filter, Download, MoreHorizontal, Edit, Trash2, Eye, Loader2, Copy, Globe, Key, Boxes, ShieldCheck, ChevronLeft, ChevronRight, Check } from 'lucide-react';
import { getKeycloakClients } from '@/actions/keycloak/clients/get-clients';
import { KeycloakClientsStatsCards } from '@/components/custom/card/keycloak-clients-stats-card';
import type { KeycloakApplication } from '@/types/keycloak/clients';
import { toast } from 'sonner';

const PAGE_SIZE_OPTIONS = [5, 10, 20, 30, 50];

export default function Page() {
  const [searchTerm, setSearchTerm] = useState('');
  const [isCreateDialogOpen, setIsCreateDialogOpen] = useState(false);
  const [selectedApp, setSelectedApp] = useState<KeycloakApplication | null>(null);
  const [applications, setApplications] = useState<KeycloakApplication[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [currentPage, setCurrentPage] = useState(1);
  const [itemsPerPage, setItemsPerPage] = useState(10);
  const [copiedId, setCopiedId] = useState<string | null>(null);

  useEffect(() => {
    const load = async () => {
      try {
        setLoading(true);
        const result = await getKeycloakClients();
        if (result.success) {
          setApplications(result.data);
        } else {
          setError(result.error || 'Error desconocido');
        }
      } catch (err: any) {
        console.error('Error al cargar clientes:', err);
        setError('No se pudieron cargar los clientes.');
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
      const inDomain = app.domain.toLowerCase().includes(q);
      const inClientId = app.client_id.toLowerCase().includes(q);
      return inName || inDesc || inDomain || inClientId;
    });
  }, [applications, searchTerm]);

  // Paginación
  const totalPages = Math.ceil(filteredApps.length / itemsPerPage);
  const startIndex = (currentPage - 1) * itemsPerPage;
  const endIndex = startIndex + itemsPerPage;
  const currentApps = filteredApps.slice(startIndex, endIndex);

  // Reset página cuando cambia el filtro o items por página
  useEffect(() => {
    setCurrentPage(1);
  }, [searchTerm, itemsPerPage]);

  const getStatusBadge = (status: string) => {
    const base = 'border px-2 py-0.5 rounded-full text-xs';
    switch (status) {
      case 'active':
        return <span className={`${base} bg-emerald-500/15 text-emerald-500 border-emerald-500/30`}>Activo</span>;
      case 'development':
        return <span className={`${base} bg-blue-500/15 text-blue-500 border-blue-500/30`}>Desarrollo</span>;
      case 'inactive':
        return <span className={`${base} bg-zinc-500/15 text-zinc-500 border-zinc-500/30`}>Inactivo</span>;
      default:
        return <span className={`${base} bg-zinc-500/15 text-zinc-500 border-zinc-500/30`}>{status}</span>;
    }
  };

  const getProtocolBadge = (protocol: string) => {
    const base = 'border px-2 py-0.5 rounded-full text-xs';
    switch (protocol) {
      case 'openid-connect':
        return <span className={`${base} bg-blue-500/15 text-blue-500 border-blue-500/30`}>OAuth 2.0</span>;
      case 'saml':
        return <span className={`${base} bg-purple-500/15 text-purple-500 border-purple-500/30`}>SAML</span>;
      default:
        return <span className={`${base} bg-zinc-500/15 text-zinc-500 border-zinc-500/30`}>{protocol}</span>;
    }
  };

  const copyToClipboard = async (text: string, id: string) => {
    try {
      await navigator.clipboard.writeText(text);
      setCopiedId(id);
      toast.success('Copiado al portapapeles', {
        description: text,
      });
      setTimeout(() => setCopiedId(null), 2000);
    } catch (err) {
      toast.error('Error al copiar', {
        description: 'No se pudo copiar al portapapeles',
      });
    }
  };

  const CopyButton = ({ text, id }: { text: string; id: string }) => {
    const isCopied = copiedId === id;
    return (
      <Button
        variant="ghost"
        size="sm"
        className="px-2 h-6"
        onClick={() => copyToClipboard(text, id)}
      >
        {isCopied ? (
          <Check className="w-3 h-3 text-green-500" />
        ) : (
          <Copy className="w-3 h-3" />
        )}
      </Button>
    );
  };

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
          <h1 className="font-bold text-foreground text-3xl">Aplicaciones OAuth</h1>
          <p className="mt-1 text-muted-foreground">Gestiona los clientes OAuth y SAML del realm</p>
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
                Nuevo Cliente
              </Button>
            </DialogTrigger>
            <DialogContent className="bg-card/80 backdrop-blur-xl border-border sm:max-w-[700px]">
              <DialogHeader>
                <DialogTitle>Crear Nuevo Cliente</DialogTitle>
                <DialogDescription>Registra un nuevo cliente OAuth/SAML en Keycloak.</DialogDescription>
              </DialogHeader>
              <div className="gap-4 grid py-4">
                <div className="space-y-2">
                  <Label htmlFor="clientId">Client ID</Label>
                  <Input id="clientId" placeholder="Ej: mi-aplicacion" />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="name">Nombre</Label>
                  <Input id="name" placeholder="Ej: Mi Aplicación Web" />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="rootUrl">Root URL</Label>
                  <Input id="rootUrl" placeholder="Ej: https://miapp.gore.gob.pe" />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="description">Descripción</Label>
                  <Textarea id="description" placeholder="Describe la funcionalidad del cliente..." />
                </div>
              </div>
              <DialogFooter>
                <Button variant="outline" onClick={() => setIsCreateDialogOpen(false)}>
                  Cancelar
                </Button>
                <Button className="bg-gradient-to-r from-primary to-chart-1" disabled>
                  Crear (próximamente)
                </Button>
              </DialogFooter>
            </DialogContent>
          </Dialog>
        </div>
      </div>

      <KeycloakClientsStatsCards />

      <Card className="bg-card/50 border-border">
        <CardHeader>
          <div className="flex justify-between items-center">
            <div>
              <CardTitle>Lista de Clientes</CardTitle>
              <CardDescription>
                Mostrando {startIndex + 1}-{Math.min(endIndex, filteredApps.length)} de {filteredApps.length} clientes
              </CardDescription>
            </div>
            <div className="flex gap-2">
              <div className="relative">
                <Search className="top-1/2 left-3 absolute w-4 h-4 text-muted-foreground -translate-y-1/2 transform" />
                <Input
                  placeholder="Buscar clientes..."
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
                  <TableHead>Cliente</TableHead>
                  <TableHead>Dominio</TableHead>
                  <TableHead>Protocolo</TableHead>
                  <TableHead>Estado</TableHead>
                  <TableHead className="text-right">Acciones</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {currentApps.length === 0 ? (
                  <TableRow>
                    <TableCell colSpan={5} className="text-center py-8 text-muted-foreground">
                      No se encontraron clientes
                    </TableCell>
                  </TableRow>
                ) : (
                  currentApps.map((app) => (
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
                            <CopyButton text={app.client_id} id={`client-${app.id}`} />
                          </div>
                        </div>
                      </TableCell>
                      <TableCell>
                        <div className="flex items-center gap-2">
                          <Globe className="w-4 h-4 text-chart-4" />
                          <span className="font-mono text-sm">{app.domain}</span>
                        </div>
                      </TableCell>
                      <TableCell>
                        <div className="flex items-center gap-2">
                          <ShieldCheck className="w-4 h-4 text-chart-3" />
                          {getProtocolBadge(app.protocol)}
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
                  ))
                )}
              </TableBody>
            </Table>
          </div>

          {/* Paginación y selector de items por página */}
          <div className="flex items-center justify-between mt-4">
            <div className="flex items-center gap-2">
              <span className="text-sm text-muted-foreground">Mostrar</span>
              <Select
                value={itemsPerPage.toString()}
                onValueChange={(value) => setItemsPerPage(Number(value))}
              >
                <SelectTrigger className="w-[70px] h-9">
                  <SelectValue />
                </SelectTrigger>
                <SelectContent>
                  {PAGE_SIZE_OPTIONS.map((size) => (
                    <SelectItem key={size} value={size.toString()}>
                      {size}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
              <span className="text-sm text-muted-foreground">
                por página
              </span>
            </div>

            {totalPages > 1 && (
              <div className="flex items-center gap-4">
                <div className="text-sm text-muted-foreground">
                  Página {currentPage} de {totalPages}
                </div>
                <div className="flex gap-2">
                  <Button
                    variant="outline"
                    size="sm"
                    onClick={() => setCurrentPage((p) => Math.max(1, p - 1))}
                    disabled={currentPage === 1}
                  >
                    <ChevronLeft className="w-4 h-4 mr-1" />
                    Anterior
                  </Button>
                  <Button
                    variant="outline"
                    size="sm"
                    onClick={() => setCurrentPage((p) => Math.min(totalPages, p + 1))}
                    disabled={currentPage === totalPages}
                  >
                    Siguiente
                    <ChevronRight className="w-4 h-4 ml-1" />
                  </Button>
                </div>
              </div>
            )}
          </div>
        </CardContent>
      </Card>

      <Dialog open={!!selectedApp} onOpenChange={(o) => !o && setSelectedApp(null)}>
        <DialogContent className="bg-card/80 backdrop-blur-xl border-border sm:max-w-[700px] max-h-[90vh] overflow-y-auto">
          <DialogHeader>
            <DialogTitle>{selectedApp?.name}</DialogTitle>
            <DialogDescription>{selectedApp?.description || 'Sin descripción'}</DialogDescription>
          </DialogHeader>
          {selectedApp && (
            <div className="space-y-3 text-sm">
              <div className="flex justify-between">
                <span className="text-muted-foreground">Dominio:</span>
                <span className="font-mono">{selectedApp.domain}</span>
              </div>
              <div className="flex justify-between">
                <span className="text-muted-foreground">Protocolo:</span>
                {getProtocolBadge(selectedApp.protocol)}
              </div>
              <div className="flex justify-between">
                <span className="text-muted-foreground">Estado:</span>
                {getStatusBadge(selectedApp.status)}
              </div>
              <div className="flex justify-between">
                <span className="text-muted-foreground">Habilitado:</span>
                <span className="font-medium">{selectedApp.enabled ? 'Sí' : 'No'}</span>
              </div>
              <div>
                <Label className="font-medium text-sm">Client ID</Label>
                <div className="flex items-center gap-2 mt-1">
                  <Input value={selectedApp.client_id} readOnly className="font-mono text-sm" />
                  <Button
                    variant="outline"
                    size="sm"
                    onClick={() => copyToClipboard(selectedApp.client_id, `modal-${selectedApp.id}`)}
                  >
                    {copiedId === `modal-${selectedApp.id}` ? (
                      <Check className="w-4 h-4 text-green-500" />
                    ) : (
                      <Copy className="w-4 h-4" />
                    )}
                  </Button>
                </div>
              </div>
              {selectedApp.redirect_uris.length > 0 && (
                <div>
                  <Label className="font-medium text-sm">Redirect URIs</Label>
                  <div className="space-y-1 mt-1">
                    {selectedApp.redirect_uris.map((uri, idx) => (
                      <div key={idx} className="flex items-center gap-2">
                        <Input value={uri} readOnly className="font-mono text-xs" />
                        <Button
                          variant="outline"
                          size="sm"
                          onClick={() => copyToClipboard(uri, `uri-${selectedApp.id}-${idx}`)}
                        >
                          {copiedId === `uri-${selectedApp.id}-${idx}` ? (
                            <Check className="w-3 h-3 text-green-500" />
                          ) : (
                            <Copy className="w-3 h-3" />
                          )}
                        </Button>
                      </div>
                    ))}
                  </div>
                </div>
              )}
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