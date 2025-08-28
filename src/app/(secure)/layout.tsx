'use client';

import { useAuthGuard } from '@/hooks/useAuthGuard';
import SessionProvider from '@/providers/SessionProvider';

const Guard = () => {
  useAuthGuard();
  return null;
};

const Layout = ({ children }: { children: React.ReactNode }) => {
  return (
    <SessionProvider client_id="cum" showWhileChecking={<div>Cargando sesión…</div>}>
      <div className="flex flex-col bg-background h-screen overflow-hidden">
        <main className="flex-1 overflow-y-auto">
          <Guard />
          {children}
        </main>
      </div>
      <div className="overflow-hidden"></div>
    </SessionProvider>
  );
};

export default Layout;
