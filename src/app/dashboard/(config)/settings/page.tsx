'use client';

import { useState } from 'react';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { Switch } from '@/components/ui/switch';
import { Textarea } from '@/components/ui/textarea';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs';
import { Separator } from '@/components/ui/separator';
import { Alert, AlertDescription } from '@/components/ui/alert';
import { Settings, Shield, Bell, Database, Mail, Globe, Key, Server, AlertTriangle, CheckCircle, Save, RotateCcw, Download, Upload, Trash2 } from 'lucide-react';

interface SystemConfig {
  general: {
    systemName: string;
    systemDescription: string;
    defaultLanguage: string;
    timezone: string;
    dateFormat: string;
    sessionTimeout: number;
    maxLoginAttempts: number;
    maintenanceMode: boolean;
  };
  security: {
    passwordMinLength: number;
    passwordRequireSpecial: boolean;
    passwordRequireNumbers: boolean;
    passwordRequireUppercase: boolean;
    passwordExpiration: number;
    twoFactorRequired: boolean;
    ipWhitelist: string[];
    allowedDomains: string[];
    encryptionLevel: string;
  };
  notifications: {
    emailEnabled: boolean;
    smsEnabled: boolean;
    pushEnabled: boolean;
    securityAlerts: boolean;
    systemAlerts: boolean;
    userAlerts: boolean;
    emailServer: string;
    emailPort: number;
    emailUsername: string;
  };
  database: {
    backupEnabled: boolean;
    backupFrequency: string;
    retentionDays: number;
    compressionEnabled: boolean;
    encryptionEnabled: boolean;
    maxConnections: number;
    queryTimeout: number;
  };
  integrations: {
    ldapEnabled: boolean;
    ldapServer: string;
    ldapPort: number;
    samlEnabled: boolean;
    oauthEnabled: boolean;
    apiRateLimit: number;
    webhooksEnabled: boolean;
  };
}

const initialConfig: SystemConfig = {
  general: {
    systemName: 'Central User Manager',
    systemDescription: 'Sistema centralizado de gestión de usuarios y permisos',
    defaultLanguage: 'es',
    timezone: 'America/Lima',
    dateFormat: 'DD/MM/YYYY',
    sessionTimeout: 30,
    maxLoginAttempts: 5,
    maintenanceMode: false,
  },
  security: {
    passwordMinLength: 8,
    passwordRequireSpecial: true,
    passwordRequireNumbers: true,
    passwordRequireUppercase: true,
    passwordExpiration: 90,
    twoFactorRequired: false,
    ipWhitelist: ['192.168.1.0/24', '10.0.0.0/8'],
    allowedDomains: ['empresa.com', 'subsidiaria.com'],
    encryptionLevel: 'AES-256',
  },
  notifications: {
    emailEnabled: true,
    smsEnabled: false,
    pushEnabled: true,
    securityAlerts: true,
    systemAlerts: true,
    userAlerts: true,
    emailServer: 'smtp.empresa.com',
    emailPort: 587,
    emailUsername: 'noreply@empresa.com',
  },
  database: {
    backupEnabled: true,
    backupFrequency: 'daily',
    retentionDays: 30,
    compressionEnabled: true,
    encryptionEnabled: true,
    maxConnections: 100,
    queryTimeout: 30,
  },
  integrations: {
    ldapEnabled: false,
    ldapServer: 'ldap.empresa.com',
    ldapPort: 389,
    samlEnabled: false,
    oauthEnabled: true,
    apiRateLimit: 1000,
    webhooksEnabled: true,
  },
};

export default function SettingsManagement() {
  const [config, setConfig] = useState<SystemConfig>(initialConfig);
  const [hasChanges, setHasChanges] = useState(false);
  const [saving, setSaving] = useState(false);

  const handleSave = async () => {
    setSaving(true);
    // Simular guardado
    await new Promise((resolve) => setTimeout(resolve, 1000));
    setSaving(false);
    setHasChanges(false);
  };

  const handleReset = () => {
    setConfig(initialConfig);
    setHasChanges(false);
  };

  const updateConfig = <T extends keyof SystemConfig, K extends keyof SystemConfig[T]>(section: T, field: K & string, value: SystemConfig[T][K]) => {
    setConfig((prev) => ({
      ...prev,
      [section]: {
        ...prev[section],
        [field]: value,
      },
    }));
    setHasChanges(true);
  };

  return (
    <div className="space-y-6">
      {hasChanges && (
        <Alert>
          <AlertTriangle className="h-4 w-4" />
          <AlertDescription>Tienes cambios sin guardar. No olvides guardar tus configuraciones.</AlertDescription>
        </Alert>
      )}

      <div className="flex gap-4">
        <Button onClick={handleSave} disabled={!hasChanges || saving}>
          <Save className="w-4 h-4 mr-2" />
          {saving ? 'Guardando...' : 'Guardar Cambios'}
        </Button>
        <Button variant="outline" onClick={handleReset} disabled={!hasChanges}>
          <RotateCcw className="w-4 h-4 mr-2" />
          Restablecer
        </Button>
        <Button variant="outline">
          <Download className="w-4 h-4 mr-2" />
          Exportar Config
        </Button>
        <Button variant="outline">
          <Upload className="w-4 h-4 mr-2" />
          Importar Config
        </Button>
      </div>

      <Tabs defaultValue="general" className="space-y-6">
        <TabsList className="grid w-full grid-cols-5">
          <TabsTrigger value="general" className="flex items-center gap-2">
            <Settings className="w-4 h-4" />
            General
          </TabsTrigger>
          <TabsTrigger value="security" className="flex items-center gap-2">
            <Shield className="w-4 h-4" />
            Seguridad
          </TabsTrigger>
          <TabsTrigger value="notifications" className="flex items-center gap-2">
            <Bell className="w-4 h-4" />
            Notificaciones
          </TabsTrigger>
          <TabsTrigger value="database" className="flex items-center gap-2">
            <Database className="w-4 h-4" />
            Base de Datos
          </TabsTrigger>
          <TabsTrigger value="integrations" className="flex items-center gap-2">
            <Globe className="w-4 h-4" />
            Integraciones
          </TabsTrigger>
        </TabsList>

        <TabsContent value="general" className="space-y-6">
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center gap-2">
                <Settings className="w-5 h-5" />
                Configuraciones Generales
              </CardTitle>
              <CardDescription>Configuraciones básicas del sistema y comportamiento general</CardDescription>
            </CardHeader>
            <CardContent className="space-y-6">
              <div className="grid grid-cols-2 gap-6">
                <div className="space-y-2">
                  <Label htmlFor="systemName">Nombre del Sistema</Label>
                  <Input id="systemName" value={config.general.systemName} onChange={(e) => updateConfig('general', 'systemName', e.target.value)} />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="defaultLanguage">Idioma por Defecto</Label>
                  <Select value={config.general.defaultLanguage} onValueChange={(value) => updateConfig('general', 'defaultLanguage', value)}>
                    <SelectTrigger>
                      <SelectValue />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value="es">Español</SelectItem>
                      <SelectItem value="en">English</SelectItem>
                      <SelectItem value="pt">Português</SelectItem>
                    </SelectContent>
                  </Select>
                </div>
              </div>

              <div className="space-y-2">
                <Label htmlFor="systemDescription">Descripción del Sistema</Label>
                <Textarea id="systemDescription" value={config.general.systemDescription} onChange={(e) => updateConfig('general', 'systemDescription', e.target.value)} rows={3} />
              </div>

              <div className="grid grid-cols-3 gap-6">
                <div className="space-y-2">
                  <Label htmlFor="timezone">Zona Horaria</Label>
                  <Select value={config.general.timezone} onValueChange={(value) => updateConfig('general', 'timezone', value)}>
                    <SelectTrigger>
                      <SelectValue />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value="America/Lima">America/Lima (UTC-5)</SelectItem>
                      <SelectItem value="America/New_York">America/New_York (UTC-5)</SelectItem>
                      <SelectItem value="Europe/Madrid">Europe/Madrid (UTC+1)</SelectItem>
                    </SelectContent>
                  </Select>
                </div>
                <div className="space-y-2">
                  <Label htmlFor="sessionTimeout">Timeout de Sesión (min)</Label>
                  <Input
                    id="sessionTimeout"
                    type="number"
                    value={config.general.sessionTimeout}
                    onChange={(e) => updateConfig('general', 'sessionTimeout', Number.parseInt(e.target.value))}
                  />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="maxLoginAttempts">Máx. Intentos de Login</Label>
                  <Input
                    id="maxLoginAttempts"
                    type="number"
                    value={config.general.maxLoginAttempts}
                    onChange={(e) => updateConfig('general', 'maxLoginAttempts', Number.parseInt(e.target.value))}
                  />
                </div>
              </div>

              <Separator />

              <div className="flex items-center justify-between">
                <div className="space-y-1">
                  <Label>Modo Mantenimiento</Label>
                  <p className="text-sm text-muted-foreground">Activa el modo mantenimiento para bloquear el acceso al sistema</p>
                </div>
                <Switch checked={config.general.maintenanceMode} onCheckedChange={(checked) => updateConfig('general', 'maintenanceMode', checked)} />
              </div>
            </CardContent>
          </Card>
        </TabsContent>

        <TabsContent value="security" className="space-y-6">
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center gap-2">
                <Shield className="w-5 h-5" />
                Configuraciones de Seguridad
              </CardTitle>
              <CardDescription>Políticas de contraseñas, autenticación y control de acceso</CardDescription>
            </CardHeader>
            <CardContent className="space-y-6">
              <div className="space-y-4">
                <h4 className="font-semibold flex items-center gap-2">
                  <Key className="w-4 h-4" />
                  Políticas de Contraseñas
                </h4>

                <div className="grid grid-cols-2 gap-6">
                  <div className="space-y-2">
                    <Label htmlFor="passwordMinLength">Longitud Mínima</Label>
                    <Input
                      id="passwordMinLength"
                      type="number"
                      value={config.security.passwordMinLength}
                      onChange={(e) => updateConfig('security', 'passwordMinLength', Number.parseInt(e.target.value))}
                    />
                  </div>
                  <div className="space-y-2">
                    <Label htmlFor="passwordExpiration">Expiración (días)</Label>
                    <Input
                      id="passwordExpiration"
                      type="number"
                      value={config.security.passwordExpiration}
                      onChange={(e) => updateConfig('security', 'passwordExpiration', Number.parseInt(e.target.value))}
                    />
                  </div>
                </div>

                <div className="space-y-4">
                  <div className="flex items-center justify-between">
                    <Label>Requerir Caracteres Especiales</Label>
                    <Switch checked={config.security.passwordRequireSpecial} onCheckedChange={(checked) => updateConfig('security', 'passwordRequireSpecial', checked)} />
                  </div>
                  <div className="flex items-center justify-between">
                    <Label>Requerir Números</Label>
                    <Switch checked={config.security.passwordRequireNumbers} onCheckedChange={(checked) => updateConfig('security', 'passwordRequireNumbers', checked)} />
                  </div>
                  <div className="flex items-center justify-between">
                    <Label>Requerir Mayúsculas</Label>
                    <Switch checked={config.security.passwordRequireUppercase} onCheckedChange={(checked) => updateConfig('security', 'passwordRequireUppercase', checked)} />
                  </div>
                  <div className="flex items-center justify-between">
                    <Label>Autenticación de Dos Factores Obligatoria</Label>
                    <Switch checked={config.security.twoFactorRequired} onCheckedChange={(checked) => updateConfig('security', 'twoFactorRequired', checked)} />
                  </div>
                </div>
              </div>

              <Separator />

              <div className="space-y-4">
                <h4 className="font-semibold">Control de Acceso</h4>

                <div className="space-y-2">
                  <Label>IPs Permitidas (una por línea)</Label>
                  <Textarea
                    value={config.security.ipWhitelist.join('\n')}
                    onChange={(e) =>
                      updateConfig(
                        'security',
                        'ipWhitelist',
                        e.target.value.split('\n').filter((ip) => ip.trim()),
                      )
                    }
                    rows={4}
                    placeholder="192.168.1.0/24&#10;10.0.0.0/8"
                  />
                </div>

                <div className="space-y-2">
                  <Label>Dominios Permitidos</Label>
                  <Textarea
                    value={config.security.allowedDomains.join('\n')}
                    onChange={(e) =>
                      updateConfig(
                        'security',
                        'allowedDomains',
                        e.target.value.split('\n').filter((domain) => domain.trim()),
                      )
                    }
                    rows={3}
                    placeholder="empresa.com&#10;subsidiaria.com"
                  />
                </div>

                <div className="space-y-2">
                  <Label htmlFor="encryptionLevel">Nivel de Encriptación</Label>
                  <Select value={config.security.encryptionLevel} onValueChange={(value) => updateConfig('security', 'encryptionLevel', value)}>
                    <SelectTrigger>
                      <SelectValue />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value="AES-128">AES-128</SelectItem>
                      <SelectItem value="AES-256">AES-256</SelectItem>
                      <SelectItem value="RSA-2048">RSA-2048</SelectItem>
                    </SelectContent>
                  </Select>
                </div>
              </div>
            </CardContent>
          </Card>
        </TabsContent>

        <TabsContent value="notifications" className="space-y-6">
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center gap-2">
                <Bell className="w-5 h-5" />
                Configuraciones de Notificaciones
              </CardTitle>
              <CardDescription>Configuración de canales de notificación y alertas del sistema</CardDescription>
            </CardHeader>
            <CardContent className="space-y-6">
              <div className="space-y-4">
                <h4 className="font-semibold">Canales de Notificación</h4>

                <div className="space-y-4">
                  <div className="flex items-center justify-between">
                    <div className="space-y-1">
                      <Label>Notificaciones por Email</Label>
                      <p className="text-sm text-muted-foreground">Enviar notificaciones vía correo electrónico</p>
                    </div>
                    <Switch checked={config.notifications.emailEnabled} onCheckedChange={(checked) => updateConfig('notifications', 'emailEnabled', checked)} />
                  </div>

                  <div className="flex items-center justify-between">
                    <div className="space-y-1">
                      <Label>Notificaciones SMS</Label>
                      <p className="text-sm text-muted-foreground">Enviar notificaciones vía SMS</p>
                    </div>
                    <Switch checked={config.notifications.smsEnabled} onCheckedChange={(checked) => updateConfig('notifications', 'smsEnabled', checked)} />
                  </div>

                  <div className="flex items-center justify-between">
                    <div className="space-y-1">
                      <Label>Notificaciones Push</Label>
                      <p className="text-sm text-muted-foreground">Enviar notificaciones push en navegador</p>
                    </div>
                    <Switch checked={config.notifications.pushEnabled} onCheckedChange={(checked) => updateConfig('notifications', 'pushEnabled', checked)} />
                  </div>
                </div>
              </div>

              <Separator />

              <div className="space-y-4">
                <h4 className="font-semibold">Tipos de Alertas</h4>

                <div className="space-y-4">
                  <div className="flex items-center justify-between">
                    <Label>Alertas de Seguridad</Label>
                    <Switch checked={config.notifications.securityAlerts} onCheckedChange={(checked) => updateConfig('notifications', 'securityAlerts', checked)} />
                  </div>
                  <div className="flex items-center justify-between">
                    <Label>Alertas del Sistema</Label>
                    <Switch checked={config.notifications.systemAlerts} onCheckedChange={(checked) => updateConfig('notifications', 'systemAlerts', checked)} />
                  </div>
                  <div className="flex items-center justify-between">
                    <Label>Alertas de Usuario</Label>
                    <Switch checked={config.notifications.userAlerts} onCheckedChange={(checked) => updateConfig('notifications', 'userAlerts', checked)} />
                  </div>
                </div>
              </div>

              {config.notifications.emailEnabled && (
                <>
                  <Separator />
                  <div className="space-y-4">
                    <h4 className="font-semibold flex items-center gap-2">
                      <Mail className="w-4 h-4" />
                      Configuración de Email
                    </h4>

                    <div className="grid grid-cols-2 gap-6">
                      <div className="space-y-2">
                        <Label htmlFor="emailServer">Servidor SMTP</Label>
                        <Input id="emailServer" value={config.notifications.emailServer} onChange={(e) => updateConfig('notifications', 'emailServer', e.target.value)} />
                      </div>
                      <div className="space-y-2">
                        <Label htmlFor="emailPort">Puerto</Label>
                        <Input
                          id="emailPort"
                          type="number"
                          value={config.notifications.emailPort}
                          onChange={(e) => updateConfig('notifications', 'emailPort', Number.parseInt(e.target.value))}
                        />
                      </div>
                    </div>

                    <div className="space-y-2">
                      <Label htmlFor="emailUsername">Usuario/Email</Label>
                      <Input id="emailUsername" value={config.notifications.emailUsername} onChange={(e) => updateConfig('notifications', 'emailUsername', e.target.value)} />
                    </div>
                  </div>
                </>
              )}
            </CardContent>
          </Card>
        </TabsContent>

        <TabsContent value="database" className="space-y-6">
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center gap-2">
                <Database className="w-5 h-5" />
                Configuraciones de Base de Datos
              </CardTitle>
              <CardDescription>Configuración de respaldos, rendimiento y mantenimiento de la base de datos</CardDescription>
            </CardHeader>
            <CardContent className="space-y-6">
              <div className="space-y-4">
                <h4 className="font-semibold">Respaldos Automáticos</h4>

                <div className="flex items-center justify-between">
                  <div className="space-y-1">
                    <Label>Habilitar Respaldos</Label>
                    <p className="text-sm text-muted-foreground">Crear respaldos automáticos de la base de datos</p>
                  </div>
                  <Switch checked={config.database.backupEnabled} onCheckedChange={(checked) => updateConfig('database', 'backupEnabled', checked)} />
                </div>

                {config.database.backupEnabled && (
                  <div className="grid grid-cols-2 gap-6">
                    <div className="space-y-2">
                      <Label htmlFor="backupFrequency">Frecuencia de Respaldo</Label>
                      <Select value={config.database.backupFrequency} onValueChange={(value) => updateConfig('database', 'backupFrequency', value)}>
                        <SelectTrigger>
                          <SelectValue />
                        </SelectTrigger>
                        <SelectContent>
                          <SelectItem value="hourly">Cada Hora</SelectItem>
                          <SelectItem value="daily">Diario</SelectItem>
                          <SelectItem value="weekly">Semanal</SelectItem>
                          <SelectItem value="monthly">Mensual</SelectItem>
                        </SelectContent>
                      </Select>
                    </div>
                    <div className="space-y-2">
                      <Label htmlFor="retentionDays">Retención (días)</Label>
                      <Input
                        id="retentionDays"
                        type="number"
                        value={config.database.retentionDays}
                        onChange={(e) => updateConfig('database', 'retentionDays', Number.parseInt(e.target.value))}
                      />
                    </div>
                  </div>
                )}

                <div className="space-y-4">
                  <div className="flex items-center justify-between">
                    <Label>Compresión de Respaldos</Label>
                    <Switch checked={config.database.compressionEnabled} onCheckedChange={(checked) => updateConfig('database', 'compressionEnabled', checked)} />
                  </div>
                  <div className="flex items-center justify-between">
                    <Label>Encriptación de Respaldos</Label>
                    <Switch checked={config.database.encryptionEnabled} onCheckedChange={(checked) => updateConfig('database', 'encryptionEnabled', checked)} />
                  </div>
                </div>
              </div>

              <Separator />

              <div className="space-y-4">
                <h4 className="font-semibold flex items-center gap-2">
                  <Server className="w-4 h-4" />
                  Rendimiento
                </h4>

                <div className="grid grid-cols-2 gap-6">
                  <div className="space-y-2">
                    <Label htmlFor="maxConnections">Máx. Conexiones</Label>
                    <Input
                      id="maxConnections"
                      type="number"
                      value={config.database.maxConnections}
                      onChange={(e) => updateConfig('database', 'maxConnections', Number.parseInt(e.target.value))}
                    />
                  </div>
                  <div className="space-y-2">
                    <Label htmlFor="queryTimeout">Timeout de Query (seg)</Label>
                    <Input
                      id="queryTimeout"
                      type="number"
                      value={config.database.queryTimeout}
                      onChange={(e) => updateConfig('database', 'queryTimeout', Number.parseInt(e.target.value))}
                    />
                  </div>
                </div>
              </div>

              <div className="flex gap-4">
                <Button variant="outline">
                  <Database className="w-4 h-4 mr-2" />
                  Probar Conexión
                </Button>
                <Button variant="outline">
                  <Download className="w-4 h-4 mr-2" />
                  Crear Respaldo Manual
                </Button>
                <Button variant="outline">
                  <Trash2 className="w-4 h-4 mr-2" />
                  Limpiar Logs Antiguos
                </Button>
              </div>
            </CardContent>
          </Card>
        </TabsContent>

        <TabsContent value="integrations" className="space-y-6">
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center gap-2">
                <Globe className="w-5 h-5" />
                Integraciones Externas
              </CardTitle>
              <CardDescription>Configuración de integraciones con sistemas externos y APIs</CardDescription>
            </CardHeader>
            <CardContent className="space-y-6">
              <div className="space-y-6">
                <div className="space-y-4">
                  <div className="flex items-center justify-between">
                    <div className="space-y-1">
                      <Label>Integración LDAP</Label>
                      <p className="text-sm text-muted-foreground">Autenticación contra directorio LDAP</p>
                    </div>
                    <Switch checked={config.integrations.ldapEnabled} onCheckedChange={(checked) => updateConfig('integrations', 'ldapEnabled', checked)} />
                  </div>

                  {config.integrations.ldapEnabled && (
                    <div className="grid grid-cols-2 gap-6 ml-6">
                      <div className="space-y-2">
                        <Label htmlFor="ldapServer">Servidor LDAP</Label>
                        <Input id="ldapServer" value={config.integrations.ldapServer} onChange={(e) => updateConfig('integrations', 'ldapServer', e.target.value)} />
                      </div>
                      <div className="space-y-2">
                        <Label htmlFor="ldapPort">Puerto</Label>
                        <Input
                          id="ldapPort"
                          type="number"
                          value={config.integrations.ldapPort}
                          onChange={(e) => updateConfig('integrations', 'ldapPort', Number.parseInt(e.target.value))}
                        />
                      </div>
                    </div>
                  )}
                </div>

                <Separator />

                <div className="space-y-4">
                  <div className="flex items-center justify-between">
                    <div className="space-y-1">
                      <Label>Autenticación SAML</Label>
                      <p className="text-sm text-muted-foreground">Single Sign-On con SAML 2.0</p>
                    </div>
                    <Switch checked={config.integrations.samlEnabled} onCheckedChange={(checked) => updateConfig('integrations', 'samlEnabled', checked)} />
                  </div>

                  <div className="flex items-center justify-between">
                    <div className="space-y-1">
                      <Label>OAuth 2.0</Label>
                      <p className="text-sm text-muted-foreground">Autenticación con proveedores OAuth</p>
                    </div>
                    <Switch checked={config.integrations.oauthEnabled} onCheckedChange={(checked) => updateConfig('integrations', 'oauthEnabled', checked)} />
                  </div>

                  <div className="flex items-center justify-between">
                    <div className="space-y-1">
                      <Label>Webhooks</Label>
                      <p className="text-sm text-muted-foreground">Notificaciones HTTP a sistemas externos</p>
                    </div>
                    <Switch checked={config.integrations.webhooksEnabled} onCheckedChange={(checked) => updateConfig('integrations', 'webhooksEnabled', checked)} />
                  </div>
                </div>

                <Separator />

                <div className="space-y-4">
                  <h4 className="font-semibold">API y Límites</h4>

                  <div className="space-y-2">
                    <Label htmlFor="apiRateLimit">Límite de Rate (req/min)</Label>
                    <Input
                      id="apiRateLimit"
                      type="number"
                      value={config.integrations.apiRateLimit}
                      onChange={(e) => updateConfig('integrations', 'apiRateLimit', Number.parseInt(e.target.value))}
                    />
                  </div>
                </div>

                <div className="flex gap-4">
                  <Button variant="outline">
                    <CheckCircle className="w-4 h-4 mr-2" />
                    Probar LDAP
                  </Button>
                  <Button variant="outline">
                    <Key className="w-4 h-4 mr-2" />
                    Generar API Key
                  </Button>
                  <Button variant="outline">
                    <Globe className="w-4 h-4 mr-2" />
                    Ver Webhooks
                  </Button>
                </div>
              </div>
            </CardContent>
          </Card>
        </TabsContent>
      </Tabs>
    </div>
  );
}
