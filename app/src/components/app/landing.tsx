'use client';

import Link from 'next/link';
import Image from 'next/image';
import React from 'react';
import { Header } from './header';
import { Footer } from './footer';
import { ArrowRight, ShieldAlert, Users, Settings2, Lock, Activity, Server } from 'lucide-react';

import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Badge } from '@/components/ui/badge';
import { Button } from '@/components/ui/button';

// Contenedor utilitario
const Container: React.FC<React.PropsWithChildren<{ className?: string }>> = ({ children, className }) => (
  <div className={`mx-auto w-full container px-6 ${className ?? ''}`}>{children}</div>
);

// Componente para mostrar la captura del sistema administrativo
const DashboardPreview: React.FC = () => {
  return (
    <div className="mx-auto mt-12 w-full max-w-5xl">
      <div className="rounded-xl border border-border bg-card shadow-2xl overflow-hidden">
        {/* Barra de título estilo navegador */}
        <div className="flex items-center gap-1.5 border-b border-border bg-muted/50 px-4 py-3">
          <div className="w-3 h-3 rounded-full bg-red-400/80" />
          <div className="w-3 h-3 rounded-full bg-yellow-400/80" />
          <div className="w-3 h-3 rounded-full bg-green-400/80" />
          <div className="ml-2 text-xs text-muted-foreground/50 font-medium select-none font-mono">
            cum.regionayacucho.gob.pe
          </div>
        </div>
        
        {/* Contenedor de la imagen */}
        <div className="relative bg-muted/20 aspect-video w-full flex items-center justify-center group">
           
           {/* IMAGEN MODO CLARO: Se muestra por defecto, se oculta en dark mode */}
           <Image 
                src="/img/dashboard-preview-light.webp" 
                alt="Panel de Administración CUM (Claro)" 
                fill 
                className="object-cover object-top block dark:hidden"
                priority
            />

           {/* IMAGEN MODO OSCURO: Se oculta por defecto, se muestra en dark mode */}
           <Image 
                src="/img/dashboard-preview-dark.webp" 
                alt="Panel de Administración CUM (Oscuro)" 
                fill 
                className="object-cover object-top hidden dark:block"
                priority
            />
        </div>
      </div>
    </div>
  );
};

export const Content = () => {
  return (
    <section className="w-full pb-20">
      <Container className="flex flex-col items-center text-center">
        
        {/* Hero Section */}
        <div className="space-y-6 pt-20 sm:pt-28 max-w-4xl">
          <Badge variant="outline" className="rounded-full px-4 py-1.5 text-sm font-medium border-primary/20 bg-primary/5 text-primary">
            <ShieldAlert className="w-3.5 h-3.5 mr-2" />
            Acceso Restringido • Solo Personal Autorizado
          </Badge>

          <h1 className="font-extrabold text-foreground text-4xl sm:text-5xl lg:text-6xl tracking-tight leading-[1.1]">
            Sistema Central de <br />
            <span className="text-primary">Gestión de Usuarios</span>
          </h1>

          <p className="text-muted-foreground text-lg sm:text-xl max-w-2xl mx-auto leading-relaxed">
            Plataforma administrativa para la gestión centralizada de identidades, control de acceso basado en roles (RBAC) y auditoría de seguridad para todo el ecosistema institucional.
          </p>

          <div className="flex flex-col sm:flex-row justify-center items-center gap-4 pt-6">
            <Link href="/dashboard" tabIndex={-1}>
              <Button size="lg" className="rounded-full px-8 h-12 text-base font-medium shadow-lg shadow-primary/20 transition-all hover:scale-105">
                Acceder al Panel
                <ArrowRight className="ml-2 w-4 h-4" />
              </Button>
            </Link>
            <Link href="#modules" tabIndex={-1}>
              <Button variant="ghost" size="lg" className="rounded-full px-8 h-12 text-base hover:bg-muted/50">
                Ver Módulos
              </Button>
            </Link>
          </div>
        </div>

        {/* Preview Image */}
        <DashboardPreview />

        {/* Features / Modules Section */}
        <div id="modules" className="mt-24 w-full">
            <div className="text-center mb-16">
                <h2 className="text-3xl font-bold tracking-tight mb-4">Administración Integral</h2>
                <p className="text-muted-foreground text-lg max-w-2xl mx-auto">
                    Herramientas avanzadas diseñadas para administradores de sistemas y oficiales de seguridad.
                </p>
            </div>

            <div className="grid grid-cols-1 md:grid-cols-3 gap-8 max-w-7xl mx-auto text-left">
                {/* Card 1: Gestión de Usuarios */}
                <Card className="group border-border/50 bg-card hover:bg-accent/5 transition-all duration-300 hover:shadow-lg hover:border-primary/20">
                    <CardHeader>
                        <div className="bg-blue-500/10 mb-4 w-12 h-12 rounded-lg flex items-center justify-center group-hover:bg-blue-500/20 transition-colors">
                            <Users className="w-6 h-6 text-blue-600 dark:text-blue-400" />
                        </div>
                        <CardTitle className="text-xl">Directorio de Usuarios</CardTitle>
                    </CardHeader>
                    <CardContent className="text-muted-foreground">
                        Administración del ciclo de vida de usuarios, unidades orgánicas y posiciones estructurales. Sincronización con bases de datos de RR.HH.
                    </CardContent>
                </Card>

                {/* Card 2: Seguridad y Roles */}
                <Card className="group border-border/50 bg-card hover:bg-accent/5 transition-all duration-300 hover:shadow-lg hover:border-primary/20">
                    <CardHeader>
                        <div className="bg-red-500/10 mb-4 w-12 h-12 rounded-lg flex items-center justify-center group-hover:bg-red-500/20 transition-colors">
                            <Lock className="w-6 h-6 text-red-600 dark:text-red-400" />
                        </div>
                        <CardTitle className="text-xl">Control de Accesos (RBAC)</CardTitle>
                    </CardHeader>
                    <CardContent className="text-muted-foreground">
                        Definición granular de roles, permisos por módulo y asignación de privilegios a nivel de aplicación. Gestión de restricciones y bloqueos.
                    </CardContent>
                </Card>

                {/* Card 3: Auditoría */}
                <Card className="group border-border/50 bg-card hover:bg-accent/5 transition-all duration-300 hover:shadow-lg hover:border-primary/20">
                    <CardHeader>
                        <div className="bg-amber-500/10 mb-4 w-12 h-12 rounded-lg flex items-center justify-center group-hover:bg-amber-500/20 transition-colors">
                            <Activity className="w-6 h-6 text-amber-600 dark:text-amber-400" />
                        </div>
                        <CardTitle className="text-xl">Auditoría y Sesiones</CardTitle>
                    </CardHeader>
                    <CardContent className="text-muted-foreground">
                        Monitoreo en tiempo real de sesiones activas (Keycloak). Logs detallados de cambios en permisos y trazabilidad de acciones administrativas.
                    </CardContent>
                </Card>
            </div>
            
            {/* Stats / Tech info strip */}
            <div className="mt-20 border-y bg-muted/30 py-12">
                <div className="grid grid-cols-2 md:grid-cols-4 gap-8 text-center max-w-5xl mx-auto">
                    <div className="space-y-2">
                        <div className="font-bold text-3xl sm:text-4xl text-foreground">SSO</div>
                        <div className="text-sm font-medium text-muted-foreground uppercase tracking-wider">Integración</div>
                    </div>
                    <div className="space-y-2">
                        <div className="font-bold text-3xl sm:text-4xl text-foreground">24/7</div>
                        <div className="text-sm font-medium text-muted-foreground uppercase tracking-wider">Disponibilidad</div>
                    </div>
                    <div className="space-y-2">
                        <div className="font-bold text-3xl sm:text-4xl text-foreground">API</div>
                        <div className="text-sm font-medium text-muted-foreground uppercase tracking-wider">Interoperabilidad</div>
                    </div>
                     <div className="space-y-2">
                        <div className="flex justify-center items-center h-full">
                            <Server className="w-8 h-8 text-muted-foreground/50" />
                        </div>
                    </div>
                </div>
            </div>

        </div>
      </Container>
    </section>
  );
};

const Landing = () => {
  return (
    <main className="flex flex-col bg-background min-h-screen text-foreground">
      <Header />
      <div className="flex grow flex-col items-center w-full">
        <Content />
      </div>
      <Footer />
    </main>
  );
};

export default Landing;