'use client';

import { useState, useRef } from 'react';
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle } from '@/components/ui/dialog';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { ScrollArea } from '@/components/ui/scroll-area';
import { toast } from 'sonner';
import { Upload, FileSpreadsheet, AlertCircle, CheckCircle2, XCircle, Download } from 'lucide-react';
import { fn_bulk_create_ubigeos, BulkUploadResponse } from '@/actions/ubigeos/fn_bulk_create_ubigeos';
import { fn_download_ubigeos_template } from '@/actions/ubigeos/fn_download_ubigeos_template';

interface BulkUploadModalProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  onSuccess: () => void;
}

export default function BulkUploadUbigeosModal({ open, onOpenChange, onSuccess }: BulkUploadModalProps) {
  const [loading, setLoading] = useState(false);
  const [downloading, setDownloading] = useState(false);
  const [file, setFile] = useState<File | null>(null);
  const [result, setResult] = useState<BulkUploadResponse | null>(null);
  const fileInputRef = useRef<HTMLInputElement>(null);

  const handleDownloadTemplate = async () => {
    try {
      setDownloading(true);
      const blob = await fn_download_ubigeos_template();
      const url = window.URL.createObjectURL(blob);
      const a = document.createElement('a');
      a.href = url;
      a.download = `plantilla_ubigeos_${new Date().toISOString().slice(0, 10)}.xlsx`;
      document.body.appendChild(a);
      a.click();
      window.URL.revokeObjectURL(url);
      document.body.removeChild(a);
      toast.success('Plantilla descargada correctamente');
    } catch (error: any) {
      toast.error(error.message || 'Error al descargar plantilla');
    } finally {
      setDownloading(false);
    }
  };

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const selectedFile = e.target.files?.[0];
    if (!selectedFile) return;

    if (!selectedFile.name.endsWith('.xlsx') && !selectedFile.name.endsWith('.xls')) {
      toast.error('Solo se permiten archivos Excel (.xlsx, .xls)');
      return;
    }

    setFile(selectedFile);
    setResult(null);
  };

  const handleSubmit = async () => {
    if (!file) return;

    try {
      setLoading(true);
      
      const formData = new FormData();
      formData.append('file', file);

      const response = await fn_bulk_create_ubigeos(formData);
      setResult(response);

      if (response.created > 0) {
        toast.success(`${response.created} ubigeos creados correctamente`);
        onSuccess();
      }

      if (response.failed > 0 || response.skipped > 0) {
        toast.warning(`${response.skipped} omitidos, ${response.failed} con errores`);
      }
    } catch (error: any) {
      toast.error(error.message || 'Error en carga masiva');
    } finally {
      setLoading(false);
    }
  };

  const handleClose = () => {
    setFile(null);
    setResult(null);
    if (fileInputRef.current) {
      fileInputRef.current.value = '';
    }
    onOpenChange(false);
  };

  return (
    <Dialog open={open} onOpenChange={handleClose}>
      <DialogContent className="bg-card/80 backdrop-blur-xl border-border sm:max-w-[600px] max-h-[90vh] overflow-y-auto">
        <DialogHeader>
          <DialogTitle>Carga Masiva de Ubigeos</DialogTitle>
          <DialogDescription>
            Sube un archivo Excel con los ubigeos a registrar
          </DialogDescription>
        </DialogHeader>

        <div className="space-y-4">
          {/* Descargar plantilla */}
          <div className="flex items-center justify-between p-4 bg-muted/50 rounded-lg border">
            <div className="flex items-center gap-3">
              <FileSpreadsheet className="w-8 h-8 text-green-600" />
              <div>
                <p className="font-medium">Plantilla Excel</p>
                <p className="text-sm text-muted-foreground">Descarga la plantilla con el formato correcto</p>
              </div>
            </div>
            <Button variant="outline" onClick={handleDownloadTemplate} disabled={downloading}>
              <Download className="w-4 h-4 mr-2" />
              {downloading ? 'Descargando...' : 'Descargar'}
            </Button>
          </div>

          {/* Subir archivo */}
          <div className="space-y-2">
            <Label>Archivo Excel</Label>
            <div
              className="border-2 border-dashed rounded-lg p-8 text-center cursor-pointer hover:border-primary/50 transition-colors"
              onClick={() => fileInputRef.current?.click()}
            >
              <Input
                ref={fileInputRef}
                type="file"
                accept=".xlsx,.xls"
                onChange={handleFileChange}
                className="hidden"
              />
              <Upload className="w-10 h-10 mx-auto text-muted-foreground mb-2" />
              {file ? (
                <div>
                  <p className="text-sm font-medium">{file.name}</p>
                  <p className="text-xs text-muted-foreground mt-1">
                    {(file.size / 1024).toFixed(2)} KB
                  </p>
                </div>
              ) : (
                <p className="text-sm text-muted-foreground">
                  Haz clic o arrastra un archivo Excel aquí
                </p>
              )}
            </div>
          </div>

          {/* Resultado */}
          {result && (
            <div className="space-y-3">
              <Label>Resultado de la carga</Label>
              <div className="grid grid-cols-3 gap-2">
                <div className="p-3 bg-green-500/10 rounded-lg text-center">
                  <CheckCircle2 className="w-5 h-5 mx-auto text-green-500 mb-1" />
                  <p className="text-lg font-bold text-green-600">{result.created}</p>
                  <p className="text-xs text-muted-foreground">Creados</p>
                </div>
                <div className="p-3 bg-yellow-500/10 rounded-lg text-center">
                  <AlertCircle className="w-5 h-5 mx-auto text-yellow-500 mb-1" />
                  <p className="text-lg font-bold text-yellow-600">{result.skipped}</p>
                  <p className="text-xs text-muted-foreground">Omitidos</p>
                </div>
                <div className="p-3 bg-red-500/10 rounded-lg text-center">
                  <XCircle className="w-5 h-5 mx-auto text-red-500 mb-1" />
                  <p className="text-lg font-bold text-red-600">{result.failed}</p>
                  <p className="text-xs text-muted-foreground">Fallidos</p>
                </div>
              </div>

              {result.errors && result.errors.length > 0 && (
                <ScrollArea className="h-[150px] border rounded-lg p-3">
                  <p className="text-sm font-medium mb-2">Errores encontrados:</p>
                  {result.errors.map((err, idx) => (
                    <div key={idx} className="text-sm text-destructive flex items-start gap-2 mb-1">
                      <XCircle className="w-4 h-4 mt-0.5 shrink-0" />
                      <span>Fila {err.row}: {err.message}</span>
                    </div>
                  ))}
                </ScrollArea>
              )}
            </div>
          )}

          {/* Info */}
          {!result && (
            <div className="bg-muted/50 p-4 rounded-lg">
              <p className="text-sm font-medium mb-2">Formato requerido:</p>
              <ul className="text-sm text-muted-foreground space-y-1">
                <li>• Columnas: ubigeo_code, inei_code, department, province, district</li>
                <li>• Máximo 1000 registros por carga</li>
                <li>• Los duplicados serán omitidos automáticamente</li>
              </ul>
            </div>
          )}
        </div>

        <DialogFooter>
          <Button variant="outline" onClick={handleClose}>
            {result ? 'Cerrar' : 'Cancelar'}
          </Button>
          {!result && (
            <Button
              onClick={handleSubmit}
              disabled={!file || loading}
              className="bg-linear-to-r from-primary to-chart-1"
            >
              {loading ? 'Procesando...' : 'Cargar Ubigeos'}
            </Button>
          )}
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}