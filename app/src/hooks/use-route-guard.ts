// src/hooks/use-route-guard.ts
'use client';

import { useEffect, useState } from 'react';
import { usePathname, useRouter } from 'next/navigation';

interface Module {
  id: string;
  route: string | null;
  children?: Module[];
}

// Extraer todas las rutas permitidas de los módulos
function extractAllowedRoutes(modules: Module[]): string[] {
  const routes: string[] = ['/dashboard']; // Siempre permitir dashboard base

  const extract = (mods: Module[]) => {
    for (const mod of mods) {
      if (mod.route) {
        routes.push(mod.route);
      }
      if (mod.children && mod.children.length > 0) {
        extract(mod.children);
      }
    }
  };

  extract(modules);
  return routes;
}

// Verificar si una ruta está permitida
function isRouteAllowed(pathname: string, allowedRoutes: string[]): boolean {
  return allowedRoutes.some((route) => {
    // Coincidencia exacta
    if (pathname === route) return true;
    // Permitir subrutas (ej: /dashboard/users/123 si /dashboard/users está permitido)
    if (pathname.startsWith(route + '/')) return true;
    return false;
  });
}

export function useRouteGuard(modules: Module[] | null) {
  const pathname = usePathname();
  const router = useRouter();
  const [isAuthorized, setIsAuthorized] = useState<boolean | 'checking'>('checking');

  useEffect(() => {
    if (!modules) {
      setIsAuthorized(false);
      return;
    }

    const allowedRoutes = extractAllowedRoutes(modules);
    const allowed = isRouteAllowed(pathname, allowedRoutes);

    if (!allowed) {
      router.replace('/unauthorized');
      setIsAuthorized(false);
    } else {
      setIsAuthorized(true);
    }
  }, [pathname, modules, router]);

  return isAuthorized;
}