"use client";

import { useState } from "react";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Switch } from "@/components/ui/switch";
import { Textarea } from "@/components/ui/textarea";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { Badge } from "@/components/ui/badge";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { Separator } from "@/components/ui/separator";
import { Alert, AlertDescription } from "@/components/ui/alert";
import { MapPin, Clock, Save, Camera, Download, Trash2, Eye, EyeOff, AlertTriangle, CheckCircle } from "lucide-react";
import { User, Shield, Bell, Activity, Key, Smartphone, Calendar, QrCode } from "lucide-react";

interface UserProfile {
  personal: {
    firstName: string;
    lastName: string;
    email: string;
    phone: string;
    avatar: string;
    department: string;
    position: string;
    location: string;
    timezone: string;
    language: string;
    bio: string;
  };
  security: {
    twoFactorEnabled: boolean;
    passwordLastChanged: string;
    trustedDevices: number;
    activeSessions: number;
    loginNotifications: boolean;
    securityAlerts: boolean;
  };
  preferences: {
    theme: string;
    notifications: {
      email: boolean;
      push: boolean;
      sms: boolean;
    };
    privacy: {
      profileVisible: boolean;
      activityVisible: boolean;
      contactVisible: boolean;
    };
  };
}

const initialProfile: UserProfile = {
  personal: {
    firstName: "Admin",
    lastName: "Usuario",
    email: "admin@empresa.com",
    phone: "+51 999 123 456",
    avatar: "/placeholder.svg?height=100&width=100",
    department: "Tecnología",
    position: "Administrador del Sistema",
    location: "Lima, Perú",
    timezone: "America/Lima",
    language: "es",
    bio: "Administrador principal del sistema CUM con más de 5 años de experiencia en gestión de usuarios y seguridad.",
  },
  security: {
    twoFactorEnabled: true,
    passwordLastChanged: "2024-01-15",
    trustedDevices: 3,
    activeSessions: 2,
    loginNotifications: true,
    securityAlerts: true,
  },
  preferences: {
    theme: "system",
    notifications: {
      email: true,
      push: true,
      sms: false,
    },
    privacy: {
      profileVisible: true,
      activityVisible: false,
      contactVisible: true,
    },
  },
};

const recentActivity = [
  {
    id: 1,
    action: "Inicio de sesión",
    device: "Chrome en Windows",
    location: "Lima, Perú",
    timestamp: "2024-01-20 09:15:00",
    status: "success",
  },
  {
    id: 2,
    action: "Cambio de configuración",
    device: "Sistema Web",
    location: "Lima, Perú",
    timestamp: "2024-01-19 16:30:00",
    status: "success",
  },
  {
    id: 3,
    action: "Intento de acceso fallido",
    device: "Mobile App",
    location: "IP desconocida",
    timestamp: "2024-01-18 22:45:00",
    status: "warning",
  },
  {
    id: 4,
    action: "Actualización de perfil",
    device: "Chrome en Windows",
    location: "Lima, Perú",
    timestamp: "2024-01-17 14:20:00",
    status: "success",
  },
];

export default function AccountManagement() {
  const [profile, setProfile] = useState<UserProfile>(initialProfile);
  const [hasChanges, setHasChanges] = useState(false);
  const [saving, setSaving] = useState(false);
  const [showPassword, setShowPassword] = useState(false);

  const handleSave = async () => {
    setSaving(true);
    // Simular guardado
    await new Promise((resolve) => setTimeout(resolve, 1000));
    setSaving(false);
    setHasChanges(false);
  };

  const updateProfile = (section: keyof UserProfile, field: string, value: string | boolean | number) => {
    setProfile((prev) => ({
      ...prev,
      [section]: {
        ...prev[section],
        [field]: value,
      },
    }));
    setHasChanges(true);
  };

  const updateNestedProfile = <T extends keyof UserProfile, K extends keyof UserProfile[T], V extends UserProfile[T][K]>(
    section: T,
    subsection: K & string,
    field: string,
    value: V
  ) => {
    setProfile((prev) => ({
      ...prev,
      [section]: {
        ...prev[section],
        [subsection]: {
          ...(prev[section] as UserProfile[T])[subsection],
          [field]: value,
        },
      },
    }));
    setHasChanges(true);
  };

  return (
    <div className="space-y-6">
      {hasChanges && (
        <Alert>
          <AlertTriangle className="h-4 w-4" />
          <AlertDescription>Tienes cambios sin guardar. No olvides guardar tu perfil.</AlertDescription>
        </Alert>
      )}

      <div className="flex gap-4">
        <Button onClick={handleSave} disabled={!hasChanges || saving}>
          <Save className="w-4 h-4 mr-2" />
          {saving ? "Guardando..." : "Guardar Cambios"}
        </Button>
        <Button variant="outline">
          <Download className="w-4 h-4 mr-2" />
          Exportar Datos
        </Button>
      </div>

      <Tabs defaultValue="profile" className="space-y-6">
        <TabsList className="grid w-full grid-cols-4">
          <TabsTrigger value="profile" className="flex items-center gap-2">
            <User className="w-4 h-4" />
            Perfil
          </TabsTrigger>
          <TabsTrigger value="security" className="flex items-center gap-2">
            <Shield className="w-4 h-4" />
            Seguridad
          </TabsTrigger>
          <TabsTrigger value="preferences" className="flex items-center gap-2">
            <Bell className="w-4 h-4" />
            Preferencias
          </TabsTrigger>
          <TabsTrigger value="activity" className="flex items-center gap-2">
            <Activity className="w-4 h-4" />
            Actividad
          </TabsTrigger>
        </TabsList>

        <TabsContent value="profile" className="space-y-6">
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center gap-2">
                <User className="w-5 h-5" />
                Información Personal
              </CardTitle>
              <CardDescription>Actualiza tu información personal y datos de contacto</CardDescription>
            </CardHeader>
            <CardContent className="space-y-6">
              <div className="flex items-center gap-6">
                <div className="relative">
                  <Avatar className="w-24 h-24">
                    <AvatarImage src={profile.personal.avatar || "/placeholder.svg"} />
                    <AvatarFallback className="text-2xl bg-gradient-to-r from-blue-500 to-purple-600 text-white">
                      {profile.personal.firstName[0]}
                      {profile.personal.lastName[0]}
                    </AvatarFallback>
                  </Avatar>
                  <Button size="sm" variant="outline" className="absolute -bottom-2 -right-2 rounded-full w-8 h-8 p-0">
                    <Camera className="w-4 h-4" />
                  </Button>
                </div>
                <div className="space-y-2">
                  <h3 className="text-lg font-semibold">
                    {profile.personal.firstName} {profile.personal.lastName}
                  </h3>
                  <p className="text-muted-foreground">{profile.personal.position}</p>
                  <Badge variant="secondary">{profile.personal.department}</Badge>
                </div>
              </div>

              <Separator />

              <div className="grid grid-cols-2 gap-6">
                <div className="space-y-2">
                  <Label htmlFor="firstName">Nombre</Label>
                  <Input id="firstName" value={profile.personal.firstName} onChange={(e) => updateProfile("personal", "firstName", e.target.value)} />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="lastName">Apellido</Label>
                  <Input id="lastName" value={profile.personal.lastName} onChange={(e) => updateProfile("personal", "lastName", e.target.value)} />
                </div>
              </div>

              <div className="grid grid-cols-2 gap-6">
                <div className="space-y-2">
                  <Label htmlFor="email">Email</Label>
                  <Input id="email" type="email" value={profile.personal.email} onChange={(e) => updateProfile("personal", "email", e.target.value)} />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="phone">Teléfono</Label>
                  <Input id="phone" value={profile.personal.phone} onChange={(e) => updateProfile("personal", "phone", e.target.value)} />
                </div>
              </div>

              <div className="grid grid-cols-2 gap-6">
                <div className="space-y-2">
                  <Label htmlFor="department">Departamento</Label>
                  <Select value={profile.personal.department} onValueChange={(value) => updateProfile("personal", "department", value)}>
                    <SelectTrigger>
                      <SelectValue />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value="Tecnología">Tecnología</SelectItem>
                      <SelectItem value="Recursos Humanos">Recursos Humanos</SelectItem>
                      <SelectItem value="Finanzas">Finanzas</SelectItem>
                      <SelectItem value="Operaciones">Operaciones</SelectItem>
                    </SelectContent>
                  </Select>
                </div>
                <div className="space-y-2">
                  <Label htmlFor="position">Posición</Label>
                  <Input id="position" value={profile.personal.position} onChange={(e) => updateProfile("personal", "position", e.target.value)} />
                </div>
              </div>

              <div className="grid grid-cols-2 gap-6">
                <div className="space-y-2">
                  <Label htmlFor="location">Ubicación</Label>
                  <Input id="location" value={profile.personal.location} onChange={(e) => updateProfile("personal", "location", e.target.value)} />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="timezone">Zona Horaria</Label>
                  <Select value={profile.personal.timezone} onValueChange={(value) => updateProfile("personal", "timezone", value)}>
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
              </div>

              <div className="space-y-2">
                <Label htmlFor="bio">Biografía</Label>
                <Textarea id="bio" value={profile.personal.bio} onChange={(e) => updateProfile("personal", "bio", e.target.value)} rows={4} placeholder="Cuéntanos sobre ti..." />
              </div>
            </CardContent>
          </Card>
        </TabsContent>

        <TabsContent value="security" className="space-y-6">
          <div className="grid gap-6">
            <Card>
              <CardHeader>
                <CardTitle className="flex items-center gap-2">
                  <Shield className="w-5 h-5" />
                  Configuración de Seguridad
                </CardTitle>
                <CardDescription>Gestiona tu contraseña y configuraciones de autenticación</CardDescription>
              </CardHeader>
              <CardContent className="space-y-6">
                <div className="space-y-4">
                  <div className="flex items-center justify-between">
                    <div className="space-y-1">
                      <Label>Autenticación de Dos Factores</Label>
                      <p className="text-sm text-muted-foreground">Protege tu cuenta con un segundo factor de autenticación</p>
                    </div>
                    <div className="flex items-center gap-2">
                      <Switch checked={profile.security.twoFactorEnabled} onCheckedChange={(checked) => updateProfile("security", "twoFactorEnabled", checked)} />
                      {profile.security.twoFactorEnabled && (
                        <Badge variant="outline" className="text-green-600 border-green-600">
                          <CheckCircle className="w-3 h-3 mr-1" />
                          Activo
                        </Badge>
                      )}
                    </div>
                  </div>

                  {profile.security.twoFactorEnabled && (
                    <div className="ml-6 space-y-4">
                      <Button variant="outline" size="sm">
                        <QrCode className="w-4 h-4 mr-2" />
                        Ver Código QR
                      </Button>
                      <Button variant="outline" size="sm">
                        <Key className="w-4 h-4 mr-2" />
                        Códigos de Respaldo
                      </Button>
                    </div>
                  )}
                </div>

                <Separator />

                <div className="space-y-4">
                  <h4 className="font-semibold">Cambiar Contraseña</h4>

                  <div className="space-y-4">
                    <div className="space-y-2">
                      <Label htmlFor="currentPassword">Contraseña Actual</Label>
                      <div className="relative">
                        <Input id="currentPassword" type={showPassword ? "text" : "password"} placeholder="Ingresa tu contraseña actual" />
                        <Button
                          type="button"
                          variant="ghost"
                          size="sm"
                          className="absolute right-0 top-0 h-full px-3 py-2 hover:bg-transparent"
                          onClick={() => setShowPassword(!showPassword)}
                        >
                          {showPassword ? <EyeOff className="h-4 w-4" /> : <Eye className="h-4 w-4" />}
                        </Button>
                      </div>
                    </div>

                    <div className="grid grid-cols-2 gap-4">
                      <div className="space-y-2">
                        <Label htmlFor="newPassword">Nueva Contraseña</Label>
                        <Input id="newPassword" type="password" placeholder="Nueva contraseña" />
                      </div>
                      <div className="space-y-2">
                        <Label htmlFor="confirmPassword">Confirmar Contraseña</Label>
                        <Input id="confirmPassword" type="password" placeholder="Confirma la nueva contraseña" />
                      </div>
                    </div>

                    <Button>
                      <Key className="w-4 h-4 mr-2" />
                      Cambiar Contraseña
                    </Button>
                  </div>
                </div>

                <Separator />

                <div className="space-y-4">
                  <h4 className="font-semibold">Notificaciones de Seguridad</h4>

                  <div className="space-y-4">
                    <div className="flex items-center justify-between">
                      <Label>Notificar inicios de sesión</Label>
                      <Switch checked={profile.security.loginNotifications} onCheckedChange={(checked) => updateProfile("security", "loginNotifications", checked)} />
                    </div>
                    <div className="flex items-center justify-between">
                      <Label>Alertas de seguridad</Label>
                      <Switch checked={profile.security.securityAlerts} onCheckedChange={(checked) => updateProfile("security", "securityAlerts", checked)} />
                    </div>
                  </div>
                </div>
              </CardContent>
            </Card>

            <Card>
              <CardHeader>
                <CardTitle>Estado de Seguridad</CardTitle>
                <CardDescription>Información sobre tu actividad de seguridad actual</CardDescription>
              </CardHeader>
              <CardContent>
                <div className="grid grid-cols-2 gap-6">
                  <div className="space-y-2">
                    <div className="flex items-center gap-2">
                      <Calendar className="w-4 h-4 text-muted-foreground" />
                      <span className="text-sm text-muted-foreground">Última actualización de contraseña</span>
                    </div>
                    <p className="font-medium">{profile.security.passwordLastChanged}</p>
                  </div>

                  <div className="space-y-2">
                    <div className="flex items-center gap-2">
                      <Smartphone className="w-4 h-4 text-muted-foreground" />
                      <span className="text-sm text-muted-foreground">Dispositivos de confianza</span>
                    </div>
                    <p className="font-medium">{profile.security.trustedDevices} dispositivos</p>
                  </div>

                  <div className="space-y-2">
                    <div className="flex items-center gap-2">
                      <Activity className="w-4 h-4 text-muted-foreground" />
                      <span className="text-sm text-muted-foreground">Sesiones activas</span>
                    </div>
                    <p className="font-medium">{profile.security.activeSessions} sesiones</p>
                  </div>

                  <div className="space-y-2">
                    <Button variant="outline" size="sm">
                      <Eye className="w-4 h-4 mr-2" />
                      Ver Dispositivos
                    </Button>
                  </div>
                </div>
              </CardContent>
            </Card>
          </div>
        </TabsContent>

        <TabsContent value="preferences" className="space-y-6">
          <div className="grid gap-6">
            <Card>
              <CardHeader>
                <CardTitle className="flex items-center gap-2">
                  <Bell className="w-5 h-5" />
                  Preferencias de Notificaciones
                </CardTitle>
                <CardDescription>Configura cómo y cuándo recibir notificaciones</CardDescription>
              </CardHeader>
              <CardContent className="space-y-6">
                <div className="space-y-4">
                  <h4 className="font-semibold">Canales de Notificación</h4>

                  <div className="space-y-4">
                    <div className="flex items-center justify-between">
                      <div className="space-y-1">
                        <Label>Notificaciones por Email</Label>
                        <p className="text-sm text-muted-foreground">Recibir notificaciones en tu correo</p>
                      </div>
                      <Switch
                        checked={profile.preferences.notifications.email}
                        onCheckedChange={(checked) =>
                          updateNestedProfile("preferences", "notifications", "email", {
                            email: checked,
                            push: profile.preferences.notifications.push,
                            sms: profile.preferences.notifications.sms,
                          })
                        }
                      />
                    </div>

                    <div className="flex items-center justify-between">
                      <div className="space-y-1">
                        <Label>Notificaciones Push</Label>
                        <p className="text-sm text-muted-foreground">Notificaciones en el navegador</p>
                      </div>
                      <Switch
                        checked={profile.preferences.notifications.push}
                        onCheckedChange={(checked) =>
                          updateNestedProfile("preferences", "notifications", "push", {
                            email: profile.preferences.notifications.email,
                            push: checked,
                            sms: profile.preferences.notifications.sms,
                          })
                        }
                      />
                    </div>

                    <div className="flex items-center justify-between">
                      <div className="space-y-1">
                        <Label>Notificaciones SMS</Label>
                        <p className="text-sm text-muted-foreground">Mensajes de texto a tu teléfono</p>
                      </div>
                      <Switch
                        checked={profile.preferences.notifications.sms}
                        onCheckedChange={(checked) =>
                          updateNestedProfile("preferences", "notifications", "sms", {
                            email: profile.preferences.notifications.email,
                            push: profile.preferences.notifications.push,
                            sms: checked,
                          })
                        }
                      />
                    </div>
                  </div>
                </div>

                <Separator />

                <div className="space-y-4">
                  <h4 className="font-semibold">Configuración de Privacidad</h4>

                  <div className="space-y-4">
                    <div className="flex items-center justify-between">
                      <Label>Perfil visible para otros usuarios</Label>
                      <Switch
                        checked={profile.preferences.privacy.profileVisible}
                        onCheckedChange={(checked) =>
                          updateNestedProfile("preferences", "privacy", "profileVisible", {
                            profileVisible: checked,
                            activityVisible: profile.preferences.privacy.activityVisible,
                            contactVisible: profile.preferences.privacy.contactVisible,
                          })
                        }
                      />
                    </div>
                    <div className="flex items-center justify-between">
                      <Label>Actividad visible</Label>
                      <Switch
                        checked={profile.preferences.privacy.activityVisible}
                        onCheckedChange={(checked) =>
                          updateNestedProfile("preferences", "privacy", "activityVisible", {
                            profileVisible: profile.preferences.privacy.profileVisible,
                            activityVisible: checked,
                            contactVisible: profile.preferences.privacy.contactVisible,
                          })
                        }
                      />
                    </div>
                    <div className="flex items-center justify-between">
                      <Label>Información de contacto visible</Label>
                      <Switch
                        checked={profile.preferences.privacy.contactVisible}
                        onCheckedChange={(checked) =>
                          updateNestedProfile("preferences", "privacy", "contactVisible", {
                            profileVisible: profile.preferences.privacy.profileVisible,
                            activityVisible: profile.preferences.privacy.activityVisible,
                            contactVisible: checked,
                          })
                        }
                      />
                    </div>
                  </div>
                </div>
              </CardContent>
            </Card>

            <Card>
              <CardHeader>
                <CardTitle>Preferencias de Interfaz</CardTitle>
                <CardDescription>Personaliza la apariencia y comportamiento de la interfaz</CardDescription>
              </CardHeader>
              <CardContent className="space-y-6">
                <div className="space-y-2">
                  <Label htmlFor="theme">Tema</Label>
                  <Select value={profile.preferences.theme} onValueChange={(value) => updateProfile("preferences", "theme", value)}>
                    <SelectTrigger>
                      <SelectValue />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value="light">Claro</SelectItem>
                      <SelectItem value="dark">Oscuro</SelectItem>
                      <SelectItem value="system">Sistema</SelectItem>
                    </SelectContent>
                  </Select>
                </div>

                <div className="space-y-2">
                  <Label htmlFor="language">Idioma</Label>
                  <Select value={profile.personal.language} onValueChange={(value) => updateProfile("personal", "language", value)}>
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
              </CardContent>
            </Card>
          </div>
        </TabsContent>

        <TabsContent value="activity" className="space-y-6">
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center gap-2">
                <Activity className="w-5 h-5" />
                Actividad Reciente
              </CardTitle>
              <CardDescription>Historial de tu actividad y accesos al sistema</CardDescription>
            </CardHeader>
            <CardContent>
              <div className="space-y-4">
                {recentActivity.map((activity) => (
                  <div key={activity.id} className="flex items-center justify-between p-4 border rounded-lg">
                    <div className="flex items-center gap-4">
                      <div className={`w-2 h-2 rounded-full ${activity.status === "success" ? "bg-green-500" : activity.status === "warning" ? "bg-yellow-500" : "bg-red-500"}`} />
                      <div>
                        <p className="font-medium">{activity.action}</p>
                        <div className="flex items-center gap-4 text-sm text-muted-foreground">
                          <span className="flex items-center gap-1">
                            <Smartphone className="w-3 h-3" />
                            {activity.device}
                          </span>
                          <span className="flex items-center gap-1">
                            <MapPin className="w-3 h-3" />
                            {activity.location}
                          </span>
                          <span className="flex items-center gap-1">
                            <Clock className="w-3 h-3" />
                            {activity.timestamp}
                          </span>
                        </div>
                      </div>
                    </div>
                    <Badge variant={activity.status === "success" ? "default" : "destructive"}>{activity.status === "success" ? "Exitoso" : "Alerta"}</Badge>
                  </div>
                ))}
              </div>

              <div className="flex justify-center mt-6">
                <Button variant="outline">Ver Historial Completo</Button>
              </div>
            </CardContent>
          </Card>

          <Card>
            <CardHeader>
              <CardTitle>Gestión de Datos</CardTitle>
              <CardDescription>Opciones para gestionar tus datos personales</CardDescription>
            </CardHeader>
            <CardContent className="space-y-4">
              <div className="flex gap-4">
                <Button variant="outline">
                  <Download className="w-4 h-4 mr-2" />
                  Descargar Mis Datos
                </Button>
                <Button variant="outline" className="text-destructive hover:text-destructive">
                  <Trash2 className="w-4 h-4 mr-2" />
                  Eliminar Cuenta
                </Button>
              </div>
              <p className="text-sm text-muted-foreground">
                Puedes descargar una copia de todos tus datos o solicitar la eliminación de tu cuenta. La eliminación de cuenta es irreversible.
              </p>
            </CardContent>
          </Card>
        </TabsContent>
      </Tabs>
    </div>
  );
}
