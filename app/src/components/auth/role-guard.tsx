'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import { RoleProvider, type RoleValue } from '@/providers/role';
import { toast } from 'sonner';

const APP_CLIENT_ID = process.env.NEXT_PUBLIC_APP_CLIENT_ID!;

type RoleApiResponse = {
  id: string;
  role?: string;
  modules?: any[];
};

const RoleGuard = ({ children }: { children: React.ReactNode }) => {
  const router = useRouter();
  const [role, setRole] = useState<RoleValue | null | 'loading'>('loading');

  useEffect(() => {
    let stop = false;

    const fetchRole = async () => {
      try {
        if (!APP_CLIENT_ID) {
          toast.error('Falta env NEXT_PUBLIC_APP_CLIENT_ID');
          if (!stop) setRole(null);
          return;
        }

        const res = await fetch(`/api/me/role?client_id=${encodeURIComponent(APP_CLIENT_ID)}`, {
          method: 'GET',
          cache: 'no-store',
        });

        // CASO 1: No autenticado (Sin sesión Keycloak) -> Dejar que SessionGuard maneje o poner null
        if (res.status === 401) {
          if (!stop) setRole(null);
          return;
        }

        // CASO 2: Autenticado, pero SIN ACCESO a esta App (Backend devuelve 404)
        if (res.status === 404) {
          console.warn('Usuario autenticado pero sin rol asignado para este cliente.');
          router.replace('/unauthorized');
          return; // IMPORTANTE: Detener ejecución para no hacer setRole(null) y causar bucle
        }

        if (!res.ok) throw new Error('role fetch failed');

        const data = (await res.json()) as RoleApiResponse;

        // CASO 3: Respuesta 200, pero array de módulos vacío (Defensa en profundidad)
        if (!data.modules || data.modules.length === 0) {
          router.replace('/unauthorized');
          return;
        }

        if (!stop) {
          setRole(
            data?.role
              ? { id: data.id, name: data.role, modules: data.modules || [] }
              : null
          );
        }
      } catch (error) {
        console.error('Error fetching role:', error);
        // Si hay un error de red o desconocido, podrías mandar a unauthorized o null
        // Para evitar bucles en dashboard, mejor mandamos a unauthorized si falla drásticamente
        if (!stop) setRole(null); 
      }
    };

    fetchRole();
    return () => {
      stop = true;
    };
  }, [router]);

  useEffect(() => {
    // Si role es null (ej. 401), intentamos ir a dashboard (que probablemente redirija a login por el SessionGuard)
    // Pero si fue 404, el return temprano de arriba evitó que lleguemos aquí con null, ya que redirigió a unauthorized.
    if (role === null) router.replace('/dashboard');
  }, [role, router]);

  if (role === 'loading') return null;
  if (role === null) return null;

  return <RoleProvider value={role}>{children}</RoleProvider>;
};

export default RoleGuard;