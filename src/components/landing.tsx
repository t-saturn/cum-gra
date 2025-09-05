'use client';

import { ThemeToggle } from '@/components/theme/theme-toggle';
import Image from 'next/image';
import { motion } from 'framer-motion';
import React from 'react';
import { Login } from './auth/auth';
import { ArrowRight, Lock, ShieldCheck, Users } from 'lucide-react';

const Container: React.FC<React.PropsWithChildren<{ className?: string }>> = ({ children, className }) => (
  <div className={`mx-auto w-full max-w-6xl px-6 ${className ?? ''}`}>{children}</div>
);

const Logo: React.FC = () => (
  <div className="flex items-center gap-3">
    <Image src="/img/logo.png" alt="Logo Gobierno Regional de Ayacucho" width={40} height={40} className="rounded-full" priority />
    <div className="leading-tight">
      <p className="font-bold tracking-tight">
        <span className="font-black text-[#d20f39] text-xl">Gobierno Regional</span>
      </p>
      <p className="font-black text-[#d20f39] text-xl tracking-tight">de Ayacucho</p>
    </div>
  </div>
);

const Header: React.FC = () => (
  <header className="z-20 relative bg-background">
    <Container className="flex justify-between items-center py-4">
      <Logo />
      <ThemeToggle />
    </Container>
  </header>
);

const AnimatedBackdrop: React.FC = () => (
  <div className="-z-10 absolute inset-0 pointer-events-none">
    <div className="absolute inset-0 bg-[radial-gradient(circle_at_center,rgba(210,15,57,0.06),transparent_60%)]" />
    <div className="absolute inset-0 bg-[linear-gradient(to_bottom_right,rgba(210,15,57,0.06),transparent)]" />

    <motion.div
      className="-top-24 -left-24 absolute bg-[#d20f39]/20 blur-3xl rounded-full w-72 h-72"
      animate={{ x: [0, 20, -10, 0], y: [0, -10, 10, 0], rotate: [0, 15, -10, 0] }}
      transition={{ duration: 16, repeat: Infinity, ease: 'easeInOut' }}
    />
    <motion.div
      className="right-0 bottom-0 absolute bg-rose-500/10 blur-3xl rounded-full w-80 h-80"
      animate={{ x: [0, -15, 10, 0], y: [0, 10, -10, 0] }}
      transition={{ duration: 18, repeat: Infinity, ease: 'easeInOut' }}
    />

    <div className="absolute inset-0 opacity-40 [mask-image:radial-gradient(circle_at_center,black,transparent_70%)]">
      <div className="bg-[radial-gradient(#000_1px,transparent_1px)] w-full h-full [background-size:12px_12px]" />
    </div>
  </div>
);

const Hero: React.FC = () => (
  <section className="relative min-h-[70vh]">
    <AnimatedBackdrop />
    <Container className="items-center gap-10 grid grid-cols-1 lg:grid-cols-2 py-16 min-h-[70vh]">
      <div className="space-y-7">
        <motion.h1
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.6 }}
          className="font-black text-4xl md:text-5xl leading-tight tracking-tight"
        >
          Plataforma Única de Autenticación
          <span className="block text-[#d20f39]">Gobierno Regional de Ayacucho</span>
        </motion.h1>

        <motion.p initial={{ opacity: 0, y: 20 }} animate={{ opacity: 1, y: 0 }} transition={{ duration: 0.8, delay: 0.1 }} className="max-w-xl text-muted-foreground text-base">
          Accede de forma segura a los servicios institucionales con inicio de sesión único (SSO), cumplimiento, auditoría y protección de cuentas.
        </motion.p>

        <motion.div initial={{ opacity: 0, y: 10 }} animate={{ opacity: 1, y: 0 }} transition={{ duration: 0.6, delay: 0.2 }} className="flex flex-wrap items-center gap-3">
          <Login />

          <a href="#beneficios" className="group inline-flex items-center gap-2 hover:shadow px-4 py-2 border rounded-2xl font-medium text-sm transition">
            Conoce más
            <ArrowRight className="w-4 h-4 transition group-hover:translate-x-0.5" />
          </a>
        </motion.div>

        <motion.ul
          initial="hidden"
          animate="show"
          variants={{ hidden: {}, show: { transition: { staggerChildren: 0.08 } } }}
          className="flex flex-wrap gap-4 mt-4 text-muted-foreground text-sm"
        >
          {['Cifrado moderno', 'Auditoría de accesos', 'Integración con sistemas internos'].map((item) => (
            <motion.li key={item} variants={{ hidden: { opacity: 0, y: 6 }, show: { opacity: 1, y: 0 } }} className="inline-flex items-center gap-2 px-3 py-1 border rounded-full">
              <span className="bg-[#d20f39] rounded-full w-1.5 h-1.5" />
              {item}
            </motion.li>
          ))}
        </motion.ul>
      </div>

      <motion.div
        initial={{ opacity: 0, scale: 0.98 }}
        animate={{ opacity: 1, scale: 1 }}
        transition={{ duration: 0.6, delay: 0.15 }}
        className="relative lg:justify-self-end mt-8 lg:mt-0 lg:max-w-none max-w-md sm:max-w-lg"
      >
        <div className="relative bg-background/60 backdrop-blur p-5 sm:p-6 border rounded-3xl overflow-hidden">
          <div className="absolute inset-0 bg-gradient-to-br from-rose-500/10 via-transparent to-transparent" aria-hidden />

          <div className="relative gap-4 sm:gap-5 grid">
            <div className="gap-3 md:gap-4 grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3">
              <Stat icon={<Lock className="w-5 h-5" />} label="Intentos bloqueados" value="1,248" />
              <Stat icon={<ShieldCheck className="w-5 h-5" />} label="Sesiones activas" value="527" />
              <Stat icon={<Users className="w-5 h-5" />} label="Usuarios" value="15,932" />
            </div>

            <div className="p-4 border rounded-2xl">
              <p className="font-medium text-sm">Estado del servicio</p>
              <div className="flex items-center gap-2 mt-2 text-sm">
                <span className="bg-emerald-500 rounded-full w-2 h-2 animate-pulse" />
                Operativo
              </div>
              <div className="bg-gradient-to-r from-transparent via-[#d20f39]/10 to-transparent mt-4 rounded-xl w-full h-20 sm:h-24 overflow-hidden">
                <AnimatedBars />
              </div>
            </div>
          </div>
        </div>
      </motion.div>
    </Container>
  </section>
);

const Stat: React.FC<{ icon: React.ReactNode; label: string; value: string }> = ({ icon, label, value }) => (
  <div className="p-4 border rounded-2xl">
    <div className="flex justify-between items-center">
      <span className="text-muted-foreground">{label}</span>
      <div className="opacity-70">{icon}</div>
    </div>
    <div className="mt-1 font-bold text-2xl tracking-tight">{value}</div>
  </div>
);

const AnimatedBars: React.FC = () => {
  const bars = Array.from({ length: 24 }, () => 0);
  return (
    <div className="flex items-end gap-1 p-2 h-full">
      {bars.map((_, i) => (
        <motion.div
          key={i}
          className="flex-1 bg-[#d20f39] rounded-full w-2"
          initial={{ height: 6 + (i % 5) * 4 }}
          animate={{ height: [10, 40, 16, 28, 12, 48, 10] }}
          transition={{ duration: 2.2 + (i % 5) * 0.2, repeat: Infinity, repeatType: 'mirror', ease: 'easeInOut', delay: i * 0.03 }}
        />
      ))}
    </div>
  );
};

// Benefits section
const Benefits: React.FC = () => (
  <section id="beneficios" className="relative">
    <Container className="py-16">
      <motion.h2
        initial={{ opacity: 0, y: 10 }}
        whileInView={{ opacity: 1, y: 0 }}
        viewport={{ once: true, margin: '-100px' }}
        transition={{ duration: 0.5 }}
        className="font-black text-3xl text-center tracking-tight"
      >
        Beneficios Clave
      </motion.h2>

      <div className="gap-5 grid grid-cols-1 md:grid-cols-3 mt-10">
        <BenefitCard title="Seguridad y Cumplimiento" desc="Políticas de acceso, auditoría y controles para proteger la información institucional." />
        <BenefitCard title="Experiencia Unificada" desc="Un solo inicio de sesión para múltiples sistemas, con redirecciones confiables." />
        <BenefitCard title="Escalable y Modular" desc="Arquitectura preparada para integrarse con nuevas aplicaciones y servicios." />
      </div>
    </Container>
  </section>
);

const BenefitCard: React.FC<{ title: string; desc: string }> = ({ title, desc }) => (
  <motion.div
    initial={{ opacity: 0, y: 8 }}
    whileInView={{ opacity: 1, y: 0 }}
    viewport={{ once: true, margin: '-80px' }}
    transition={{ duration: 0.4 }}
    className="relative shadow-sm p-6 border rounded-3xl overflow-hidden"
  >
    <div className="-top-10 -right-10 absolute bg-rose-500/10 blur-2xl rounded-full w-28 h-28" />
    <h3 className="font-bold text-lg">{title}</h3>
    <p className="mt-2 text-muted-foreground text-sm">{desc}</p>
  </motion.div>
);

const Footer: React.FC = () => (
  <footer className="border-t">
    <Container className="flex md:flex-row flex-col justify-between items-center gap-3 py-6 text-muted-foreground text-sm">
      <p>© {new Date().getFullYear()} Gobierno Regional de Ayacucho — SSO</p>
      <nav className="flex items-center gap-4">
        <a href="#" className="hover:text-foreground">
          Privacidad
        </a>
        <a href="#" className="hover:text-foreground">
          Términos
        </a>
        <a href="#" className="hover:text-foreground">
          Estado del servicio
        </a>
      </nav>
    </Container>
  </footer>
);

export default function Landing() {
  return (
    <div className="bg-background min-h-screen text-foreground">
      <Header />

      <Hero />

      <Benefits />

      <Footer />
    </div>
  );
}
