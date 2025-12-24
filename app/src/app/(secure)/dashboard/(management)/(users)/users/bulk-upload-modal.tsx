'use client';

import { useState, useRef } from 'react';
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle } from '@/components/ui/dialog';
import { Button } from '@/components/ui/button';
import { Progress } from '@/components/ui/progress';
import { ScrollArea } from '@/components/ui/scroll-area';
import { Badge } from '@/components/ui/badge';
import { toast } from 'sonner';
import { Upload, FileSpreadsheet, CheckCircle2, XCircle, AlertCircle } from 'lucide-react';
import { fn_bulk_create_users } from '@/actions/users/fn_bulk_create_users';
import type { BulkCreateUsersResponse } from '@/types/users';

interface BulkUploadModalProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  onSuccess: () => void;
}

export default function BulkUploadModal({ open, onOpenChange, onSuccess }: BulkUploadModalProps) {
  const [file, setFile] = useState<File | null>(null);
  const [uploading, setUploading] = useState(false);
  const [result, setResult] = useState<BulkCreateUsersResponse | null>(null);
  const fileInputRef = useRef<HTMLInputElement>(null);

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const selectedFile = e.target.files?.[0];
    if (selectedFile) {
      if (selectedFile.name.endsWith('.xlsx') || selectedFile.name.endsWith('.xls')) {
        setFile(selectedFile);
        setResult(null);
      } else {
        toast.error('Por favor selecciona un archivo Excel (.xlsx o .xls)');
      }
    }
  };

  const handleUpload = async () => {
    if (!file) {
      toast.error('Por favor selecciona un archivo');
      return;
    }

    setUploading(true);
    try {
      const response = await fn_bulk_create_users(file);
      setResult(response);
      
      if (response.failure_count === 0) {
        toast.success(`${response.success_count} usuarios creados exitosamente`);
      } else {
        toast.warning(`${response.success_count} exitosos, ${response.failure_count} fallidos`);
      }
      
      onSuccess();
    } catch (error: any) {
      toast.error(error.message || 'Error en la carga masiva');
    } finally {
      setUploading(false);
    }
  };

  const handleClose = () => {
    setFile(null);
    setResult(null);
    onOpenChange(false);
  };

  const successRate = result 
    ? Math.round((result.success_count / (result.success_count + result.failure_count)) * 100)
    : 0;

  return (
    <Dialog open={open} onOpenChange={handleClose}>
      <DialogContent className="sm:max-w-[600px] max-h-[80vh]">
        <DialogHeader>
          <DialogTitle>Carga Masiva de Usuarios</DialogTitle>
          <DialogDescription>
            Sube un archivo Excel con los datos de múltiples usuarios
          </DialogDescription>
        </DialogHeader>

        <div className="space-y-4">
          {!result ? (
            <>
              <div
                className="border-2 border-dashed border-border rounded-lg p-8 text-center cursor-pointer hover:border-primary transition"
                onClick={() => fileInputRef.current?.click()}
              >
                <input
                  ref={fileInputRef}
                  type="file"
                  accept=".xlsx,.xls"
                  onChange={handleFileChange}
                  className="hidden"
                />
                
                <FileSpreadsheet className="mx-auto h-12 w-12 text-muted-foreground mb-4" />
                
                {file ? (
                  <div className="space-y-2">
                    <p className="font-medium text-foreground">{file.name}</p>
                    <p className="text-sm text-muted-foreground">
                      {(file.size / 1024).toFixed(2)} KB
                    </p>
                    <Button variant="outline" size="sm" onClick={(e) => {
                      e.stopPropagation();
                      setFile(null);
                    }}>
                      Cambiar archivo
                    </Button>
                  </div>
                ) : (
                  <div className="space-y-2">
                    <p className="font-medium text-foreground">
                      Haz clic o arrastra un archivo Excel
                    </p>
                    <p className="text-sm text-muted-foreground">
                      Formatos soportados: .xlsx, .xls
                    </p>
                  </div>
                )}
              </div>

              <div className="bg-muted/50 p-4 rounded-lg space-y-2">
                <p className="text-sm font-medium flex items-center gap-2">
                  <AlertCircle className="h-4 w-4" />
                  Instrucciones
                </p>
                <ul className="text-sm text-muted-foreground space-y-1 list-disc list-inside">
                  <li>Descarga la plantilla Excel primero</li>
                  <li>Llena los datos en la hoja "Usuarios"</li>
                  <li>Consulta las otras hojas para IDs válidos</li>
                  <li>Si no ingresas contraseña, se usará el DNI</li>
                  <li>Los usuarios se crearán en Keycloak automáticamente</li>
                </ul>
              </div>
            </>
          ) : (
            <div className="space-y-4">
              <div className="flex items-center justify-between">
                <div className="space-y-1">
                  <p className="text-sm font-medium">Resultados de la carga</p>
                  <p className="text-xs text-muted-foreground">
                    {result.success_count + result.failure_count} usuarios procesados
                  </p>
                </div>
                <Badge variant={result.failure_count === 0 ? 'default' : 'secondary'}>
                  {successRate}% exitoso
                </Badge>
              </div>

              <Progress value={successRate} className="h-2" />

              <div className="grid grid-cols-2 gap-4">
                <div className="bg-green-500/10 border border-green-500/20 rounded-lg p-3">
                  <div className="flex items-center gap-2">
                    <CheckCircle2 className="h-4 w-4 text-green-500" />
                    <span className="text-sm font-medium">Exitosos</span>
                  </div>
                  <p className="text-2xl font-bold text-green-500 mt-1">
                    {result.success_count}
                  </p>
                </div>

                <div className="bg-red-500/10 border border-red-500/20 rounded-lg p-3">
                  <div className="flex items-center gap-2">
                    <XCircle className="h-4 w-4 text-red-500" />
                    <span className="text-sm font-medium">Fallidos</span>
                  </div>
                  <p className="text-2xl font-bold text-red-500 mt-1">
                    {result.failure_count}
                  </p>
                </div>
              </div>

              {result.results.length > 0 && (
                <ScrollArea className="h-[300px] border rounded-lg p-4">
                  <div className="space-y-2">
                    {result.results.map((item, index) => (
                      <div
                        key={index}
                        className={`flex items-start justify-between p-3 rounded-lg border ${
                          item.success
                            ? 'bg-green-500/5 border-green-500/20'
                            : 'bg-red-500/5 border-red-500/20'
                        }`}
                      >
                        <div className="flex-1">
                          <div className="flex items-center gap-2">
                            {item.success ? (
                              <CheckCircle2 className="h-4 w-4 text-green-500 shrink-0" />
                            ) : (
                              <XCircle className="h-4 w-4 text-red-500 shrink-0" />
                            )}
                            <div>
                              <p className="font-medium text-sm">{item.email}</p>
                              <p className="text-xs text-muted-foreground">DNI: {item.dni}</p>
                              {item.error && (
                                <p className="text-xs text-red-500 mt-1">{item.error}</p>
                              )}
                            </div>
                          </div>
                        </div>
                      </div>
                    ))}
                  </div>
                </ScrollArea>
              )}
            </div>
          )}
        </div>

        <DialogFooter>
          <Button variant="outline" onClick={handleClose} disabled={uploading}>
            {result ? 'Cerrar' : 'Cancelar'}
          </Button>
          {!result && file && (
            <Button onClick={handleUpload} disabled={uploading}>
              {uploading ? (
                <>
                  <Upload className="mr-2 h-4 w-4 animate-pulse" />
                  Procesando...
                </>
              ) : (
                <>
                  <Upload className="mr-2 h-4 w-4" />
                  Subir Archivo
                </>
              )}
            </Button>
          )}
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}