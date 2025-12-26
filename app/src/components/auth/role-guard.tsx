'use client';

import { useEffect, useState, useRef } from 'react';
import { useRouter, usePathname } from 'next/navigation';
import { RoleProvider, type RoleValue } from '@/providers/role';
import { toast } from 'sonner';
import { fn_get_user_role } from '@/actions/auth/fn_get_user_role';
import { buildSidebarMenu } from '@/lib/build-sidebar-menu';

function extractRoutes(modules: any[]): string[] {
  const routes = new Set<string>();

  const extract = (mods: any[]) => {
    for (const mod of mods) {
      if (mod.route) {
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
  const normalizedPath = pathname.endsWith('/') && pathname !== '/' ? pathname.slice(0, -1) : pathname;

  if (normalizedPath === '/dashboard') {
    return allowedRoutes.includes('/dashboard');
  }

  for (const route of allowedRoutes) {
    if (route === '/dashboard') continue;
    if (normalizedPath === route) return true;
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

        const allowed = isRouteAllowed(pathname, allowedRoutes);

        if (!allowed) {
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

        const message = error?.message || '';

        // Usuario no autenticado -> login
        if (message.includes('No autenticado')) {
          router.replace('/');
          return;
        }

        if (
          message.includes('404') ||
          message.includes('no tiene rol') ||
          message.includes('No se encontraron datos')
        ) {
          router.replace('/unauthorized');
          return;
        }

        if (message.includes('client_id')) {
          toast.error('Falta configuración de la aplicación (client_id)');
          router.replace('/unauthorized');
          return;
        }

        router.replace('/unauthorized');
      }
    };

    setIsAuthorized(false);
    fetchRole();

    return () => {
      isMounted = false;
    };
  }, [pathname, router]);

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