'use client';

import React, { useEffect, useState } from 'react';
import Image from 'next/image';
import { ThemeToggle } from '../theme/theme-toggle';

const Container: React.FC<React.PropsWithChildren<{ className?: string }>> = ({ children, className }) => (
  <div className={`mx-auto w-full container px-6 ${className ?? ''}`}>{children}</div>
);

const Logo: React.FC = () => (
  <div className="flex items-center gap-2 sm:gap-3" data-testid="logo">
    <Image src="/img/logo.png" alt="Logo Gobierno Regional de Ayacucho" width={36} height={36} className="rounded-full" priority />
    <div className="hidden sm:block leading-tight">
      <p className="font-bold tracking-tight">
        <span className="font-black text-[#d20f39] text-xl sm:text-xl">Gobierno Regional</span>
      </p>
      <p className="font-black text-[#d20f39] text-xl sm:text-xl tracking-tight">de Ayacucho</p>
    </div>
  </div>
);

export const Header: React.FC = () => {
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
