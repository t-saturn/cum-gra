'use client';

import React from 'react';
import Image from 'next/image';
import Link from 'next/link';
import { motion } from 'framer-motion';
import { ThemeToggle } from '@/components/theme/theme-toggle';
import { ShieldCheck, Users, Lock, ArrowRight } from 'lucide-react';

const Container: React.FC<React.PropsWithChildren<{ className?: string }>> = ({ children, className }) => (
  <div className={`mx-auto w-full max-w-6xl px-6 ${className ?? ''}`}>{children}</div>
);

const Logo: React.FC = () => (
  <div className="flex items-center gap-2 sm:gap-3" data-testid="logo">
    <Image src="/img/logo.png" alt="Logo Gobierno Regional de Ayacucho" width={36} height={36} className="rounded-full" priority />
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

const LoginButton: React.FC<{ continueTo?: string }> = ({ continueTo }) => {
  const href = continueTo ? `/auth/login?continue=${encodeURIComponent(continueTo)}` : '/auth/login';
  return (
    <Link
      href={href}
      data-testid="login-button"
      className="group inline-flex items-center gap-2 bg-[#d20f39] shadow hover:shadow-md px-5 py-2.5 rounded-2xl focus:outline-none focus-visible:ring-[#d20f39]/60 focus-visible:ring-2 font-semibold text-white text-sm transition"
    >
      Iniciar sesión
      <ArrowRight className="w-4 h-4 transition group-hover:translate-x-0.5" />
    </Link>
  );
};

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
  <section className="relative py-10 sm:py-14 lg:min-h-[70vh]">
    <AnimatedBackdrop />
    <Container className="items-center gap-8 sm:gap-12 grid grid-cols-1 lg:grid-cols-2">
      <div className="space-y-6 sm:space-y-7">
        <motion.h1
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.6 }}
          className="font-black text-3xl sm:text-4xl md:text-5xl leading-tight tracking-tight"
        >
          Plataforma Única de Autenticación
          <span className="block text-[#d20f39]">Gobierno Regional de Ayacucho</span>
        </motion.h1>

        <motion.p
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.8, delay: 0.1 }}
          className="max-w-xl text-muted-foreground text-sm sm:text-base"
        >
          Accede de forma segura a los servicios institucionales con inicio de sesión único (SSO), cumplimiento, auditoría y protección de cuentas.
        </motion.p>

        <motion.div initial={{ opacity: 0, y: 10 }} animate={{ opacity: 1, y: 0 }} transition={{ duration: 0.6, delay: 0.2 }} className="flex flex-wrap items-center gap-3">
          <LoginButton />

          <a href="#beneficios" className="group inline-flex items-center gap-2 hover:shadow px-4 py-2 border rounded-2xl font-medium text-sm transition">
            Conoce más
            <ArrowRight className="w-4 h-4 transition group-hover:translate-x-0.5" />
          </a>
        </motion.div>

        <motion.ul
          initial="hidden"
          animate="show"
          variants={{ hidden: {}, show: { transition: { staggerChildren: 0.08 } } }}
          className="flex flex-wrap gap-3 sm:gap-4 mt-4 text-muted-foreground text-xs sm:text-sm"
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
        <div className="relative bg-background/60 shadow-sm backdrop-blur p-5 sm:p-6 border rounded-3xl overflow-hidden">
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
  <div className="p-4 border rounded-2xl" data-testid={`stat-${label}`}>
    <div className="flex justify-between items-center">
      <span className="text-muted-foreground">{label}</span>
      <div className="opacity-70">{icon}</div>
    </div>
    <div className="mt-1 font-bold text-2xl tracking-tight">{value}</div>
  </div>
);

const AnimatedBars: React.FC = () => {
  // const bars = new Array(24).fill(0);
  // const bars = Array.from({ length: 24 }, (_, i) => i);
  const bars = Array.from({ length: 24 }, () => 0);

  return (
    <div className="flex items-end gap-1 p-2 h-full" data-testid="animated-bars">
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

// const Footer: React.FC = () => (
//   <footer className="border-t">
//     <Container className="flex md:flex-row flex-col justify-between items-center gap-3 py-6 text-muted-foreground text-sm">
//       <p>© {new Date().getFullYear()} Gobierno Regional de Ayacucho [Oficina de Tecnologías de la Información y Comunicación] — SSO</p>
//       <nav className="flex items-center gap-4">
//         <a href="#" className="hover:text-foreground">
//           Privacidad
//         </a>
//         <a href="#" className="hover:text-foreground">
//           Términos
//         </a>
//         <a href="#" className="hover:text-foreground">
//           Estado del servicio
//         </a>
//       </nav>
//     </Container>
//   </footer>
// );
const Footer: React.FC = () => {
  const year = new Date().getFullYear();

  // Frases que se alternan con efecto de tipeo
  const phrases = ['Gobierno Regional de Ayacucho', 'Oficina de Tecnologías de la Información y Comunicación'];

  const [phraseIndex, setPhraseIndex] = React.useState(0);
  const [subIndex, setSubIndex] = React.useState(0);
  const [deleting, setDeleting] = React.useState(false);
  const [blink, setBlink] = React.useState(true);

  // Parpadeo del cursor
  React.useEffect(() => {
    const blinkInterval = setInterval(() => setBlink((v) => !v), 500);
    return () => clearInterval(blinkInterval);
  }, []);

  // Efecto de escribir/borrar
  React.useEffect(() => {
    const current = phrases[phraseIndex];
    const isComplete = subIndex === current.length;

    const typeSpeed = 55;
    const deleteSpeed = 38;
    const pauseEnd = 1200; // pausa al terminar de escribir
    const pauseStart = 400; // pausa antes de empezar a escribir la siguiente

    let timeout: ReturnType<typeof setTimeout>;

    if (!deleting && isComplete) {
      // Pausa y empieza a borrar
      timeout = setTimeout(() => setDeleting(true), pauseEnd);
    } else if (deleting && subIndex === 0) {
      // Cambia a la siguiente frase y empieza a escribir
      timeout = setTimeout(() => {
        setDeleting(false);
        setPhraseIndex((i) => (i + 1) % phrases.length);
      }, pauseStart);
    } else {
      // Avanza un carácter (escritura o borrado)
      timeout = setTimeout(() => setSubIndex((i) => i + (deleting ? -1 : 1)), deleting ? deleteSpeed : typeSpeed);
    }

    return () => clearTimeout(timeout);
  }, [subIndex, deleting, phraseIndex]);

  // Reinicia el subIndex cuando cambiamos de frase
  React.useEffect(() => {
    if (!deleting) setSubIndex(0);
  }, [phraseIndex, deleting]);

  const typed = phrases[phraseIndex].slice(0, subIndex);

  return (
    <footer className="border-t">
      <Container className="flex md:flex-row flex-col justify-between items-center gap-3 py-6 text-muted-foreground text-sm">
        <p aria-live="polite" className="md:text-left text-center">
          © {year} <span className="font-bold whitespace-nowrap">{typed}</span>
          <span className={`ml-1 inline-block w-[1ch] select-none ${blink ? 'opacity-100' : 'opacity-0'}`}>|</span>
        </p>

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
};

const Landing = () => {
  return (
    <div className="relative min-h-screen text-foreground">
      <Header />

      <main className="h-screen overflow-y-scroll snap-mandatory snap-wrapper snap-y no-scrollbar">
        <section className="pt-20 sm:pt-24 min-h-screen snap-start">
          <Hero />
          <Footer />
        </section>
      </main>

      <style jsx global>{`
        .no-scrollbar::-webkit-scrollbar {
          display: none;
        }
        .no-scrollbar {
          -ms-overflow-style: none;
          scrollbar-width: none;
        }
      `}</style>
    </div>
  );
};

export default Landing;
