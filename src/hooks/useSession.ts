'use client';

import { useAuthContext } from '@/context/auth-context';

export function useSession() {
  const ctx = useAuthContext();
  return {
    status: ctx.status,
    session: ctx.session,
    userId: ctx.session?.user_id ?? null,
    role: ctx.session?.role ?? null,
    name: ctx.session?.name ?? null,
    email: ctx.session?.email ?? null,
    permissions: ctx.session?.module_permissions ?? [],
    refresh: ctx.refresh,
    logout: ctx.logout,
    hasPermission: ctx.hasPermission,
  };
}
