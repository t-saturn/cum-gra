'use client';

import { useSession, signIn } from 'next-auth/react';
import { useEffect } from 'react';

const SessionGuard = ({ children }: { children: React.ReactNode }) => {
  const { data: session, status } = useSession();

  useEffect(() => {
    if (session?.error === 'RefreshAccessTokenError') {
      signIn('keycloak', { callbackUrl: '/dashboard' });
    }
  }, [session]);

  useEffect(() => {
    if (status === 'unauthenticated') {
      signIn('keycloak', { callbackUrl: '/dashboard' });
    }
  }, [status]);

  if (status === 'loading') return null;
  if (status === 'unauthenticated') return null;

  return <>{children}</>;
};

export default SessionGuard;