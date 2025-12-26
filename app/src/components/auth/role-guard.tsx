'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import { RoleProvider, type RoleValue } from '@/providers/role';
import { toast } from 'sonner';
import { fn_get_user_role } from '@/actions/auth/fn_get_user_role';

const RoleGuard = ({ children }: { children: React.ReactNode }) => {
  const router = useRouter();
  const [role, setRole] = useState<RoleValue | null | 'loading'>('loading');

  useEffect(() => {
    let stop = false;

    const fetchRole = async () => {
      try {
        const data = await fn_get_user_role();

        console.log('Role data:', data);

        // CASO: Respuesta OK, pero array de módulos vacío
        if (!data.modules || data.modules.length === 0) {
          router.replace('/unauthorized');
          return;
        }

        if (!stop) {
          setRole({
            id: data.id,
            name: data.role,
            modules: data.modules,
          });
        }
      } catch (error: any) {
        console.error('Error fetching role:', error);
        const message = error?.message || '';

        // CASO: No autenticado
        if (message.includes('No autenticado')) {
          if (!stop) setRole(null);
          return;
        }

        // CASO: Sin rol asignado (404 del backend)
        if (message.includes('404') || message.includes('no tiene rol')) {
          router.replace('/unauthorized');
          return;
        }

        // CASO: Falta client_id
        if (message.includes('client_id')) {
          toast.error('Falta configuración de la aplicación (client_id)');
          if (!stop) setRole(null);
          return;
        }

        // Otros errores
        if (!stop) setRole(null);
      }
    };

    fetchRole();
    return () => {
      stop = true;
    };
  }, [router]);

  useEffect(() => {
    if (role === null) router.replace('/dashboard');
  }, [role, router]);

  if (role === 'loading') return null;
  if (role === null) return null;

  return <RoleProvider value={role}>{children}</RoleProvider>;
};

export default RoleGuard;