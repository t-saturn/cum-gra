'use client';

import { useEffect, useRef } from 'react';
import { useAuthContext } from '@/context/auth-context';
import { usePathname, useSearchParams, useRouter } from 'next/navigation';

export function useAuthGuard(opts?: { require?: string; onForbidden?: string }) {
  const { status, hasPermission } = useAuthContext();
  const pathname = usePathname();
  const search = useSearchParams();
  const router = useRouter();
  const redirectingRef = useRef(false);

  const currentRel = (() => {
    const qs = search?.toString();
    return qs ? `${pathname}?${qs}` : pathname;
  })();

  useEffect(() => {
    if (redirectingRef.current) return;

    if (status === 'unauthenticated') {
      redirectingRef.current = true;
      // Redirección centralizada al bridge local
      window.location.href = `/api/auth/introspect?to=${encodeURIComponent(currentRel)}`;
      return;
    }

    if (status === 'authenticated' && opts?.require) {
      if (!hasPermission(opts.require)) {
        router.replace(opts.onForbidden ?? '/403');
      }
    }
  }, [status, opts, currentRel, router, hasPermission]);
}
