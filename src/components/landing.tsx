'use client';

import Link from 'next/link';
import React from 'react';
import { ArrowRight } from 'lucide-react';
import { motion } from 'framer-motion';
import { Header } from './app/header';
import { Footer } from './app/footer';

const Container: React.FC<React.PropsWithChildren<{ className?: string }>> = ({ children, className }) => (
  <div className={`mx-auto w-full container px-6 ${className ?? ''}`}>{children}</div>
);


export const Content = () => {
  return (
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
                <div className="absolute inset-0 bg-[#d20f39]/60 border border-border [transform:translateZ(60px)]" />
                <div className="absolute inset-0 bg-[#d20f39]/30 border border-border [transform:rotateY(180deg)_translateZ(60px)]" />
                <div className="absolute inset-0 bg-[#d20f39]/50 border border-border [transform:rotateY(90deg)_translateZ(60px)]" />
                <div className="absolute inset-0 bg-[#d20f39]/40 border border-border [transform:rotateY(-90deg)_translateZ(60px)]" />
                <div className="absolute inset-0 bg-[#d20f39]/45 border border-border [transform:rotateX(90deg)_translateZ(60px)]" />
                <div className="absolute inset-0 bg-[#d20f39]/25 border border-border [transform:rotateX(-90deg)_translateZ(60px)]" />
              </motion.div>
            </div>
          </div>
        </div>
      </Container>
    </section>
  )
}

const Landing = () => {
  return (
    <main className="flex flex-col bg-background min-h-screen text-foreground">
      <Header />

      <div className="flex flex-grow justify-center items-center px-4 min-h-[calc(100vh-4rem)]">
        <Content />
      </div>

      <div className='mt-20'>
        <Footer />
      </div>
    </main>
  );
};

export default Landing;
