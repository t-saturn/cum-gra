'use client';

import { useAuthContext } from '@/context/auth-context';

export function usePermission() {
  const { hasPermission } = useAuthContext();
  return {
    can: (q: string) => hasPermission(q),
  };
}
