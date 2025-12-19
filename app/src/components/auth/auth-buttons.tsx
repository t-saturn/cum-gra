'use client';

import { useSession, signIn } from 'next-auth/react';
import { Button } from '@/components/ui/button';
import { keycloakSignOut } from '@/lib/keycloak-logout';

export const AuthButtons: React.FC = () => {
  const { data: session } = useSession();

  if (!session)
    return (
      <Button onClick={() => signIn('keycloak', { callbackUrl: '/dashboard' })} variant="default">
        Iniciar sesión
      </Button>
    );

  return (
    <Button onClick={() => keycloakSignOut()} variant="destructive">
      Cerrar sesión
    </Button>
  );
};