'use client';

import { useEffect, useState, useRef } from 'react';
import { useRouter, usePathname } from 'next/navigation';
import { RoleProvider, type RoleValue } from '@/providers/role';
import { toast } from 'sonner';
import { fn_get_user_role } from '@/actions/auth/fn_get_user_role';
import { buildSidebarMenu } from '@/lib/build-sidebar-menu';

// Extraer todas las rutas de los módulos (incluyendo children)
function extractRoutes(modules: any[]): string[] {
  const routes = new Set<string>();

  const extract = (mods: any[]) => {
    for (const mod of mods) {
      if (mod.route) {
        // Normalizar ruta
        const normalizedRoute =
          mod.route.endsWith('/') && mod.route !== '/' ? mod.route.slice(0, -1) : mod.route;
        routes.add(normalizedRoute);
      }
      if (mod.children && mod.children.length > 0) {
        extract(mod.children);
      }
    }
  };

  extract(modules);
  return Array.from(routes);
}

function isRouteAllowed(pathname: string, allowedRoutes: string[]): boolean {
  // Normalizar pathname
  const normalizedPath = pathname.endsWith('/') && pathname !== '/' ? pathname.slice(0, -1) : pathname;

  // /dashboard exacto siempre permitido si está en la lista
  if (normalizedPath === '/dashboard') {
    return allowedRoutes.includes('/dashboard');
  }

  // Para otras rutas, buscar coincidencia exacta o subruta de una ruta permitida
  // PERO la ruta permitida debe ser más específica que /dashboard
  for (const route of allowedRoutes) {
    // Ignorar /dashboard para la lógica de subrutas
    if (route === '/dashboard') continue;

    // Coincidencia exacta
    if (normalizedPath === route) return true;

    // Subruta permitida (ej: /dashboard/users/123 si /dashboard/users está permitido)
    if (normalizedPath.startsWith(route + '/')) return true;
  }

  return false;
}

const RoleGuard = ({ children }: { children: React.ReactNode }) => {
  const router = useRouter();
  const pathname = usePathname();
  const [role, setRole] = useState<RoleValue | null | 'loading'>('loading');
  const [isAuthorized, setIsAuthorized] = useState<boolean>(false);
  const modulesRef = useRef<any[] | null>(null);

  useEffect(() => {
    let isMounted = true;

    const fetchRole = async () => {
      try {
        const data = await fn_get_user_role();

        if (!isMounted) return;

        if (!data.modules || data.modules.length === 0) {
          router.replace('/unauthorized');
          return;
        }

        modulesRef.current = data.modules;

        const allowedRoutes = extractRoutes(data.modules);
        console.log('Rutas permitidas:', allowedRoutes);
        console.log('Ruta actual:', pathname);

        const allowed = isRouteAllowed(pathname, allowedRoutes);
        console.log('¿Ruta permitida?:', allowed);

        if (!allowed) {
          console.log('Redirigiendo a /unauthorized');
          router.replace('/unauthorized');
          setIsAuthorized(false);
          return;
        }

        setIsAuthorized(true);

        const sidebarMenu = buildSidebarMenu(data.modules);

        setRole({
          id: data.id,
          name: data.role,
          modules: data.modules,
          sidebarMenu,
        });
      } catch (error: any) {
        if (!isMounted) return;

        console.error('Error fetching role:', error);
        const message = error?.message || '';

        if (message.includes('No autenticado')) {
          setRole(null);
          return;
        }

        if (message.includes('404') || message.includes('no tiene rol')) {
          router.replace('/unauthorized');
          return;
        }

        if (message.includes('client_id')) {
          toast.error('Falta configuración de la aplicación (client_id)');
          setRole(null);
          return;
        }

        setRole(null);
      }
    };

    setIsAuthorized(false);
    fetchRole();

    return () => {
      isMounted = false;
    };
  }, [pathname, router]);

  useEffect(() => {
    if (role === null) {
      router.replace('/');
    }
  }, [role, router]);

  if (role === 'loading' || !isAuthorized) {
    return (
      <div className="flex h-screen items-center justify-center">
        <div className="h-8 w-8 animate-spin rounded-full border-4 border-primary border-t-transparent" />
      </div>
    );
  }

  if (role === null) {
    return null;
  }

  return <RoleProvider value={role}>{children}</RoleProvider>;
};

export default RoleGuard;