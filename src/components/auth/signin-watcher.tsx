'use client';

import { useSession, getSession } from 'next-auth/react';
import { useEffect, useRef } from 'react';
import { useSearchParams, useRouter } from 'next/navigation';
import { from_cb64 } from '@/helpers';

export const SignInWatcher = () => {
  const { status } = useSession();
  const sp = useSearchParams();
  const router = useRouter();
  const redirected = useRef(false);

  const targetFromParams = () => {
    const cb64 = sp.get('cb64');
    try {
      return cb64 ? from_cb64(cb64) : '/home';
    } catch {
      return '/home';
    }
  };

  const tryRedirect = async () => {
    if (redirected.current) return;
    const s = await getSession();
    if (s) {
      redirected.current = true;
      router.replace(targetFromParams());
    }
  };

  useEffect(() => {
    if (!redirected.current && status === 'authenticated') {
      redirected.current = true;
      router.replace(targetFromParams());
    }
  }, [status, router, sp]);

  useEffect(() => {
    let stopped = false;

    const kick = async () => {
      if (!stopped && !redirected.current) await tryRedirect();
    };

    kick();

    const id = setInterval(kick, 2000);
    return () => {
      stopped = true;
      clearInterval(id);
    };
  }, [sp, router]);

  useEffect(() => {
    const onVis = () => {
      if (document.visibilityState === 'visible') {
        void tryRedirect();
      }
    };
    document.addEventListener('visibilitychange', onVis);
    return () => document.removeEventListener('visibilitychange', onVis);
  }, []);

  return null;
};
