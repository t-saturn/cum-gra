'use client';

import Link from 'next/link';
import React from 'react';
import Image from 'next/image';
import { ArrowRight } from 'lucide-react';
import { ThemeToggle } from './theme/theme-toggle';
import { Footer } from './footer';

const Container: React.FC<React.PropsWithChildren<{ className?: string }>> = ({ children, className }) => (
  <div className={`mx-auto w-full max-w-6xl px-6 ${className ?? ''}`}>{children}</div>
);

const Logo: React.FC = () => (
  <div className="flex items-center gap-2 sm:gap-3" data-testid="logo">
    <Image src="/img/logo.png" alt="Logo Gobierno Regional de Ayacucho" width={40} height={40} className="rounded-full" priority />
    <div className="hidden sm:block leading-tight">
      <p className="font-bold tracking-tight">
        <span className="font-black text-[#d20f39] text-xl sm:text-2xl">Gobierno Regional</span>
      </p>
      <p className="font-black text-[#d20f39] text-xl sm:text-2xl tracking-tight">de Ayacucho</p>
    </div>
  </div>
);

const Header: React.FC = () => (
  <header className="top-0 z-50 fixed inset-x-0 bg-background/80 backdrop-blur border-b">
    <Container className="flex justify-between items-center py-3 sm:py-4">
      <Logo />
      <ThemeToggle />
    </Container>
  </header>
);

const Landing = () => {
  return (
    <div className="flex flex-col min-h-screen">
      <Header />

      <main className="flex flex-col flex-1 justify-center items-center px-6 text-center">
        <div className="space-y-6 max-w-3xl">
          <h1 className="font-extrabold text-foreground text-3xl sm:text-5xl tracking-tight">Plataforma Central de Gestión de Usuarios</h1>
          <p className="text-muted-foreground text-lg sm:text-xl">
            Administra, autentica y supervisa usuarios en un sistema integral y seguro para optimizar los procesos institucionales.
          </p>

          <div className="pt-4">
            <Link
              href="/dashboard"
              className="group inline-flex items-center gap-2 bg-[#d20f39] shadow hover:shadow-md px-6 py-3 rounded-2xl focus:outline-none focus-visible:ring-[#d20f39]/60 focus-visible:ring-2 font-semibold text-white text-base transition"
            >
              Continuar
              <ArrowRight className="w-5 h-5 transition group-hover:translate-x-1" />
            </Link>
          </div>
        </div>

        <div className="mt-10">
          <Image src="/img/landing-preview.png" alt="Vista previa plataforma" width={800} height={400} className="shadow-lg border rounded-2xl" priority />
        </div>
      </main>

      <Footer />
    </div>
  );
};

export default Landing;
