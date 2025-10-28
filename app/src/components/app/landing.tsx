'use client';

import Link from 'next/link';
import React, { useEffect, useRef, useState } from 'react';
import { Header } from './header';
import { Footer } from './footer';
import { animate } from 'animejs';
import { ArrowRight, ShieldCheck, Users, KeyRound, ServerCog, Activity, Link2, CheckCircle2, Lock, Cloud, Box, Boxes, Package } from 'lucide-react';

import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Badge } from '@/components/ui/badge';

const Container: React.FC<React.PropsWithChildren<{ className?: string }>> = ({ children, className }) => (
  <div className={`mx-auto w-full container px-6 ${className ?? ''}`}>{children}</div>
);

const PreviewSSOAnimation: React.FC = () => {
  const phase1Ref = useRef<HTMLDivElement>(null);
  const phase2Ref = useRef<HTMLDivElement>(null);
  const boxRefs = useRef<HTMLDivElement[]>([]);
  const [phase, setPhase] = useState<'loading' | 'final'>('loading');

  const setBoxRef = (el: HTMLDivElement | null, idx: number) => {
    if (!el) return;
    boxRefs.current[idx] = el;
  };

  useEffect(() => {
    if (phase1Ref.current) {
      phase1Ref.current.style.opacity = '1';
      phase1Ref.current.style.pointerEvents = 'auto';
    }
    if (phase2Ref.current) {
      phase2Ref.current.style.opacity = '0';
      phase2Ref.current.style.pointerEvents = 'none';
      phase2Ref.current.style.transform = 'scale(0.9)';
    }

    const highlight = (activeIndex: number) => {
      boxRefs.current.forEach((el, i) => {
        if (!el) return;
        animate(el, { opacity: i === activeIndex ? 1 : 0.25, scale: i === activeIndex ? 1.1 : 0.95, duration: 350, ease: 'outQuad' });
      });
    };

    let step = 0;
    const cycles = 6; // 0,1,2,0,1,2...
    const stepMs = 420;

    highlight(0);
    const interval = setInterval(() => {
      step += 1;
      if (step < cycles) {
        highlight(step % 3);
        return;
      }
      clearInterval(interval);

      if (phase1Ref.current) {
        animate(phase1Ref.current, { opacity: 0, duration: 400, ease: 'inOutQuad' });
      }

      setTimeout(() => {
        if (!phase2Ref.current) return;
        phase2Ref.current.style.pointerEvents = 'auto';
        animate(phase2Ref.current, { opacity: 1, scale: 1, duration: 650, ease: 'outBack' });

        animate('.pkg-rot', { rotate: [0, 360], duration: 5000, ease: 'linear', loop: true });

        setPhase('final');
      }, 420);
    }, stepMs);

    return () => {
      clearInterval(interval);
    };
  }, []);

  return (
    <div className="mx-auto mt-10 p-4 sm:p-6 border rounded-2xl w-full max-w-4xl">
      <p className="mb-2 text-muted-foreground text-sm">Vista previa plataforma</p>

      <div className="relative flex justify-center items-center bg-muted/30 border border-border/50 rounded-xl h-[280px] sm:h-[360px] overflow-hidden">
        <div ref={phase1Ref} className="absolute inset-0 flex justify-center items-center gap-8" aria-hidden={phase !== 'loading'}>
          {[0, 1, 2].map((i) => (
            <div key={i} ref={(el) => setBoxRef(el, i)} className="transition will-change-transform" style={{ opacity: i === 0 ? 1 : 0.25 }}>
              <Box className="w-14 sm:w-16 h-14 sm:h-16 text-[#d20f39]" />
            </div>
          ))}
        </div>

        <div ref={phase2Ref} className="absolute inset-0" aria-hidden={phase !== 'final'}>
          <div className="absolute inset-0 flex justify-center items-center">
            <Boxes className="w-20 sm:w-24 h-20 sm:h-24 text-[#d20f39]" />
          </div>

          <div className="top-6 sm:top-8 left-6 sm:left-8 absolute pkg-rot">
            <Package className="w-10 sm:w-12 h-10 sm:h-12 text-[#d20f39]" />
          </div>
          <div className="top-6 sm:top-8 right-6 sm:right-8 absolute pkg-rot">
            <Package className="w-10 sm:w-12 h-10 sm:h-12 text-[#d20f39]" />
          </div>
          <div className="bottom-6 sm:bottom-8 left-6 sm:left-8 absolute pkg-rot">
            <Package className="w-10 sm:w-12 h-10 sm:h-12 text-[#d20f39]" />
          </div>
          <div className="right-6 sm:right-8 bottom-6 sm:bottom-8 absolute pkg-rot">
            <Package className="w-10 sm:w-12 h-10 sm:h-12 text-[#d20f39]" />
          </div>
        </div>
      </div>
    </div>
  );
};

export const Content = () => {
  return (
    <section className="w-full">
      <Container className="flex flex-col items-center text-center">
        <div className="space-y-6 pt-28 sm:pt-32 max-w-3xl">
          <Badge variant="secondary" className="rounded-full">
            Plataforma SSO • Gestión de Usuarios
          </Badge>

          <h1 className="font-extrabold text-foreground text-xl sm:text-3xl tracking-tight">Plataforma Central de Gestión de Usuarios</h1>

          <p className="text-muted-foreground text-lg sm:text-xl">
            Administra, autentica y supervisa usuarios en un sistema integral y seguro para optimizar los procesos institucionales.
          </p>

          <div className="flex justify-center items-center gap-3 pt-2">
            <Link
              href="/dashboard"
              className="group inline-flex items-center gap-2 bg-[#d20f39] shadow hover:shadow-md px-6 py-3 rounded-2xl focus:outline-none focus-visible:ring-[#d20f39]/60 focus-visible:ring-2 font-semibold text-white transition"
            >
              Continuar
              <ArrowRight className="w-5 h-5 transition group-hover:translate-x-1" />
            </Link>
            <Link href="#features" className="inline-flex items-center hover:bg-accent px-6 py-3 border rounded-2xl transition">
              Ver características
            </Link>
          </div>

          <PreviewSSOAnimation />
        </div>

        <div id="features" className="gap-6 grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 mx-auto mt-14 w-full max-w-6xl">
          <Card className="hover:shadow-md transition">
            <CardHeader>
              <div className="bg-primary/10 p-2 rounded-lg w-fit">
                <Users className="w-5 h-5 text-primary" />
              </div>
              <CardTitle>Gestión Centralizada</CardTitle>
              <CardDescription>Usuarios, roles y permisos desde un solo lugar.</CardDescription>
            </CardHeader>
            <CardContent className="text-muted-foreground text-sm">Administra altas, bajas y cambios con auditoría y trazabilidad.</CardContent>
          </Card>

          <Card className="hover:shadow-md transition">
            <CardHeader>
              <div className="bg-primary/10 p-2 rounded-lg w-fit">
                <ShieldCheck className="w-5 h-5 text-primary" />
              </div>
              <CardTitle>Seguridad de Nivel Institucional</CardTitle>
              <CardDescription>Autenticación robusta y políticas de acceso.</CardDescription>
            </CardHeader>
            <CardContent className="text-muted-foreground text-sm">Políticas de contraseña, bloqueo, y controles por unidad orgánica.</CardContent>
          </Card>

          <Card className="hover:shadow-md transition">
            <CardHeader>
              <div className="bg-primary/10 p-2 rounded-lg w-fit">
                <KeyRound className="w-5 h-5 text-primary" />
              </div>
              <CardTitle>SSO y Federaciones</CardTitle>
              <CardDescription>Inicio de sesión único entre aplicaciones.</CardDescription>
            </CardHeader>
            <CardContent className="text-muted-foreground text-sm">Callback seguro por dominio y sincronización de sesión en tiempo real.</CardContent>
          </Card>

          <Card className="hover:shadow-md transition">
            <CardHeader>
              <div className="bg-primary/10 p-2 rounded-lg w-fit">
                <ServerCog className="w-5 h-5 text-primary" />
              </div>
              <CardTitle>Escalable y Modular</CardTitle>
              <CardDescription>Arquitectura lista para múltiples apps.</CardDescription>
            </CardHeader>
            <CardContent className="text-muted-foreground text-sm">Módulos por aplicaciones, unidades, posiciones y restricciones.</CardContent>
          </Card>

          <Card className="hover:shadow-md transition">
            <CardHeader>
              <div className="bg-primary/10 p-2 rounded-lg w-fit">
                <Activity className="w-5 h-5 text-primary" />
              </div>
              <CardTitle>Monitoreo y Auditoría</CardTitle>
              <CardDescription>Historial de cambios y eventos clave.</CardDescription>
            </CardHeader>
            <CardContent className="text-muted-foreground text-sm">Registro de contraseñas, revocaciones y bitácora de acceso.</CardContent>
          </Card>

          <Card className="hover:shadow-md transition">
            <CardHeader>
              <div className="bg-primary/10 p-2 rounded-lg w-fit">
                <Link2 className="w-5 h-5 text-primary" />
              </div>
              <CardTitle>Integraciones</CardTitle>
              <CardDescription>Conexiones con servicios internos.</CardDescription>
            </CardHeader>
            <CardContent className="text-muted-foreground text-sm">Enlace con intranet, módulos documentales y sistemas de RR.HH.</CardContent>
          </Card>
        </div>

        <div className="mx-auto mt-16 w-full max-w-5xl text-left">
          <h3 className="mb-4 font-bold text-xl">¿Cómo funciona?</h3>
          <div className="gap-6 grid md:grid-cols-3">
            <Card className="relative">
              <CardHeader>
                <div className="flex justify-center items-center bg-primary/10 rounded-full w-8 h-8 font-bold text-primary">1</div>
                <CardTitle>Autenticación</CardTitle>
                <CardDescription>SSO centralizado (App SSO).</CardDescription>
              </CardHeader>
              <CardContent className="text-muted-foreground text-sm">El usuario inicia sesión en la app SSO; se emite la cookie de sesión por dominio.</CardContent>
            </Card>

            <Card className="relative">
              <CardHeader>
                <div className="flex justify-center items-center bg-primary/10 rounded-full w-8 h-8 font-bold text-primary">2</div>
                <CardTitle>Redirección Segura</CardTitle>
                <CardDescription>cb64 y validación de origen.</CardDescription>
              </CardHeader>
              <CardContent className="text-muted-foreground text-sm">La app destino recibe el callback y sincroniza el contexto del usuario (ID/rol).</CardContent>
            </Card>

            <Card className="relative">
              <CardHeader>
                <div className="flex justify-center items-center bg-primary/10 rounded-full w-8 h-8 font-bold text-primary">3</div>
                <CardTitle>Acceso y Auditoría</CardTitle>
                <CardDescription>Permisos y bitácora.</CardDescription>
              </CardHeader>
              <CardContent className="text-muted-foreground text-sm">Se controla el acceso a módulos y se registran eventos clave para auditoría.</CardContent>
            </Card>
          </div>
        </div>

        <div className="mx-auto mt-16 w-full max-w-5xl">
          <Card>
            <CardHeader className="text-left">
              <CardTitle className="flex items-center gap-2">
                <Lock className="w-5 h-5 text-primary" />
                Seguridad y Cumplimiento
              </CardTitle>
              <CardDescription>Prácticas y controles aplicados.</CardDescription>
            </CardHeader>
            <CardContent className="text-muted-foreground text-sm text-left">
              <ul className="gap-3 grid sm:grid-cols-2">
                <li className="flex items-center gap-2">
                  <CheckCircle2 className="w-4 h-4 text-green-600" /> Hash Argon2 + políticas de contraseña
                </li>
                <li className="flex items-center gap-2">
                  <CheckCircle2 className="w-4 h-4 text-green-600" /> CSRF token para rutas sensibles
                </li>
                <li className="flex items-center gap-2">
                  <CheckCircle2 className="w-4 h-4 text-green-600" /> Cookies por dominio y SameSite Lax
                </li>
                <li className="flex items-center gap-2">
                  <CheckCircle2 className="w-4 h-4 text-green-600" /> Auditoría de cambios y revocaciones
                </li>
              </ul>
              <div className="flex flex-wrap items-center gap-2 mt-4">
                <Badge variant="outline" className="gap-1">
                  <Cloud className="w-3.5 h-3.5" /> Alta disponibilidad
                </Badge>
                <Badge variant="outline" className="gap-1">
                  <ShieldCheck className="w-3.5 h-3.5" /> Buenas prácticas OWASP
                </Badge>
              </div>
            </CardContent>
          </Card>
        </div>
      </Container>
    </section>
  );
};

const Landing = () => {
  return (
    <main className="flex flex-col bg-background min-h-screen text-foreground">
      <Header />
      <div className="flex flex-grow justify-center items-start px-4 min-h-[calc(100vh-4rem)]">
        <Content />
      </div>
      <div className="mt-20">
        <Footer />
      </div>
    </main>
  );
};

export default Landing;
