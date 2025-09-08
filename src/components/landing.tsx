'use client';

import Link from 'next/link';
import React, { useEffect, useState } from 'react';
import Image from 'next/image';
import { ArrowRight } from 'lucide-react';
import { Footer } from './footer';
import { motion } from 'framer-motion';
import { ThemeToggle } from './theme/theme-toggle';

const Container: React.FC<React.PropsWithChildren<{ className?: string }>> = ({ children, className }) => (
  <div className={`mx-auto w-full container px-6 ${className ?? ''}`}>{children}</div>
);

const Logo: React.FC = () => (
  <div className="flex items-center gap-2 sm:gap-3" data-testid="logo">
    <Image
      src="https://raw.githubusercontent.com/t-saturn/resources/gra/img/logo.png"
      alt="Logo Gobierno Regional de Ayacucho"
      width={36}
      height={36}
      className="rounded-full"
      priority
    />
    <div className="hidden sm:block leading-tight">
      <p className="font-bold tracking-tight">
        <span className="font-black text-[#d20f39] text-xl sm:text-xl">Gobierno Regional</span>
      </p>
      <p className="font-black text-[#d20f39] text-xl sm:text-xl tracking-tight">de Ayacucho</p>
    </div>
  </div>
);

const Header: React.FC = () => {
  const [scrolled, setScrolled] = useState(false);

  useEffect(() => {
    const onScroll = () => setScrolled(window.scrollY > 8);
    onScroll();
    window.addEventListener('scroll', onScroll, { passive: true });
    return () => window.removeEventListener('scroll', onScroll);
  }, []);

  return (
    <header
      className={[
        'fixed inset-x-0 top-0 z-50 border-b transition-colors duration-300 backdrop-blur',
        scrolled ? 'bg-background/80 border-border' : 'bg-transparent border-transparent',
      ].join(' ')}
    >
      <Container className="flex justify-between items-center py-3 sm:py-4">
        <Logo />
        <ThemeToggle />
      </Container>
    </header>
  );
};

const FooterReveal: React.FC = () => {
  const [visible, setVisible] = useState(false);

  useEffect(() => {
    const onScroll = () => setVisible(window.scrollY > 8);
    onScroll();
    window.addEventListener('scroll', onScroll, { passive: true });
    return () => window.removeEventListener('scroll', onScroll);
  }, []);

  return (
    <div
      className={['transition-all duration-300', visible ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-6 pointer-events-none select-none'].join(' ')}
      aria-hidden={!visible}
    >
      <Footer />
    </div>
  );
};

const Landing = () => {
  return (
    <div className="flex flex-col min-h-screen">
      <Header />

      <main className="flex flex-1 justify-center items-center">
        <section className="w-full">
          <Container className="flex flex-col items-center text-center">
            <div className="space-y-6 pt-28 sm:pt-32 max-w-3xl">
              <h1 className="font-extrabold text-foreground text-xl sm:text-3xl tracking-tight">Plataforma Central de Gestión de Usuarios</h1>
              <p className="text-muted-foreground text-lg sm:text-xl">
                Administra, autentica y supervisa usuarios en un sistema integral y seguro para optimizar los procesos institucionales.
              </p>

              <div className="pt-2">
                <Link
                  href="/dashboard"
                  className="group inline-flex items-center gap-2 bg-[#d20f39] shadow hover:shadow-md px-6 py-3 rounded-2xl focus:outline-none focus-visible:ring-[#d20f39]/60 focus-visible:ring-2 font-semibold text-white transition"
                >
                  Continuar
                  <ArrowRight className="w-5 h-5 transition group-hover:translate-x-1" />
                </Link>
              </div>

              <div className="mx-auto mt-10 p-4 sm:p-6 border rounded-2xl w-full max-w-4xl">
                <p className="mb-2 text-muted-foreground text-sm">Vista previa plataforma</p>
                <div className="flex justify-center items-center border border-border/50 rounded-xl h-[280px] sm:h-[360px] [perspective:1000px]">
                  <motion.div
                    className="relative w-30 h-30 [transform-style:preserve-3d]"
                    initial={{ rotateX: 0, rotateY: 0 }}
                    animate={{ rotateX: 360, rotateY: 360 }}
                    transition={{ repeat: Infinity, duration: 8, ease: 'linear' }}
                  >
                    {/* Frente */}
                    <div className="absolute inset-0 bg-[#d20f39]/60 border border-border [transform:translateZ(60px)]" />
                    {/* Atrás */}
                    <div className="absolute inset-0 bg-[#d20f39]/30 border border-border [transform:rotateY(180deg)_translateZ(60px)]" />
                    {/* Derecha */}
                    <div className="absolute inset-0 bg-[#d20f39]/50 border border-border [transform:rotateY(90deg)_translateZ(60px)]" />
                    {/* Izquierda */}
                    <div className="absolute inset-0 bg-[#d20f39]/40 border border-border [transform:rotateY(-90deg)_translateZ(60px)]" />
                    {/* Arriba */}
                    <div className="absolute inset-0 bg-[#d20f39]/45 border border-border [transform:rotateX(90deg)_translateZ(60px)]" />
                    {/* Abajo */}
                    <div className="absolute inset-0 bg-[#d20f39]/25 border border-border [transform:rotateX(-90deg)_translateZ(60px)]" />
                  </motion.div>
                </div>
              </div>
            </div>
          </Container>
        </section>
      </main>

      <FooterReveal />
    </div>
  );
};

export default Landing;
