'use client';

import { useAuthGuard } from '@/hooks/useAuthGuard';
import SessionProvider from '@/providers/SessionProvider';
import { Suspense } from 'react';

const Guard = () => {
  useAuthGuard();
  return null;
};

const Layout = ({ children }: { children: React.ReactNode }) => {
  return (
    <Suspense fallback={null}>
      <SessionProvider client_id="cum" showWhileChecking={<div>Cargando sesión…</div>}>
        <div className="flex flex-col bg-background h-screen overflow-hidden">
          <main className="flex-1 overflow-y-auto">
            <Guard />
            {children}
          </main>
        </div>
        <div className="overflow-hidden"></div>
      </SessionProvider>
    </Suspense>
  );
};

export default Layout;
