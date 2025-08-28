'use client';

import { ThemeToggle } from '@/components/theme/theme-toggle';
import Image from 'next/image';
import React from 'react';
import { Login } from './auth/auth';

const Logo = () => (
  <div className="flex items-center gap-2">
    <Image src="/img/logo.png" alt="Logo Gobierno Regional de Ayacucho" width={32} height={32} className="rounded-full" priority />
    <span className="font-bold text-xl tracking-tight">
      <span className="font-black text-[#d20f39] hover:text-red-400 text-2xl">Gobierno Regional de Ayacucho</span>
    </span>
  </div>
);

const LandingHeader = () => {
  return (
    <header className="flex justify-between items-center mx-auto px-6 pt-6 w-full max-w-6xl">
      <Logo />
      <ThemeToggle />
    </header>
  );
};

const LandingContent = () => {
  return (
    <main className="flex-1 items-center gap-10 grid grid-cols-1 lg:grid-cols-2 mx-auto px-6 py-14 w-full max-w-6xl">
      <div className="space-y-6">
        <h1 className="font-black text-3xl leading-tight">
          Auth service client
          <br />
          single sign on app main view
        </h1>
        <p className="max-w-xl">A community-driven color scheme meant for coding, designing, and much more!</p>

        <div className="flex flex-wrap gap-3 pt-2">
          <Login />
        </div>
      </div>

      <div className="hidden lg:block min-h-[380px]" />
    </main>
  );
};

export const Landing = () => {
  return (
    <div className="h-screen overflow-y-scroll snap-mandatory snap-y no-scrollbar">
      <section className="grid grid-rows-[auto_1fr_auto] min-h-screen snap-start">
        <LandingHeader />
        <LandingContent />
      </section>
    </div>
  );
};
