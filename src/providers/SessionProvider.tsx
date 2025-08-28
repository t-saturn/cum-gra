/* eslint-disable @typescript-eslint/no-explicit-any */
'use client';

import React, { useCallback, useEffect, useMemo, useRef, useState } from 'react';
import { AuthContext } from '@/context/auth-context';
import type { AuthStatus, GatewayMeResponse, GatewayMeData } from '@/types/auth';
import { usePathname, useSearchParams } from 'next/navigation';

// Helpers
function toBase64(str: string) {
  try {
    return btoa(unescape(encodeURIComponent(str)));
  } catch {
    return typeof Buffer !== 'undefined' ? Buffer.from(str, 'utf-8').toString('base64') : '';
  }
}

function normalizePermission(p: string) {
  return (p || '').trim().toLowerCase();
}

function hasPermissionImpl(perms: string[] | undefined, query: string): boolean {
  if (!perms || perms.length === 0) return false;
  const set = new Set(perms.map(normalizePermission));
  const q = normalizePermission(query);
  if (set.has(q)) return true;
  const [mod] = q.split(':');
  if (mod && set.has(mod)) return true;
  return false;
}

// Timings
// Fallback de access token (si /auth/me no trae exp): AJUSTADO A 3 MIN PARA DEV.
// En prod, ponlo a 15 min si ese es tu TTL real.
const ACCESS_TTL_MS = 3 * 60 * 1000;
const REFRESH_EARLY_MS = 60 * 1000; // refrescar ~1 min antes
const DEFAULT_FORCE_MS = Number(process.env.NEXT_PUBLIC_FORCE_REFRESH_MS || 0) || undefined;

// === NUEVO: helper para decidir refresh según nuevo bloque tokens de /auth/me o /introspect ===
function shouldRefreshFromTokens(me: GatewayMeData | null): boolean {
  if (!me) return false;

  // Backends posibles: me.data.tokens.* o me.tokens.* (según cómo lo passee el gateway).
  const anyMe: any = me;
  const tokens = anyMe.tokens || anyMe.data?.tokens || null;
  if (!tokens) return false;

  const at = tokens.access_token;
  const rt = tokens.refresh_token;

  const accessExpired = at && (at.status === 'expired' || /expirado/i.test(at?.token_detail?.message || ''));

  const refreshActive = rt && (rt.status === 'active' || rt?.token_detail?.valid === true);

  // Sólo refrescar cuando:
  // - Access está expirado
  // - Refresh está activo
  // - Y la sesión sigue activa (si viene el campo)
  const sessionStatus = anyMe.status || anyMe.data?.status || anyMe.session?.status || 'active';
  const sessionActive = String(sessionStatus).toLowerCase() === 'active';

  return !!(accessExpired && refreshActive && sessionActive);
}

type SessionProviderProps = {
  children: React.ReactNode;
  client_id: string; // requerido por /api/auth/me
  showWhileChecking?: React.ReactNode;
  defaultRedirect?: string; // adónde volver si no hay sesión (por defecto: ruta actual)
  /** Dev/testing: fuerza rotación cada N ms, ignorando exp (ej: 60_000) */
  forceRefreshEveryMs?: number;
};

export default function SessionProvider({ children, client_id, showWhileChecking = null, defaultRedirect, forceRefreshEveryMs = DEFAULT_FORCE_MS }: SessionProviderProps) {
  const [status, setStatus] = useState<AuthStatus>('checking');
  const [session, setSession] = useState<GatewayMeData | null>(null);

  const isRedirectingRef = useRef(false);
  const lastCheckRef = useRef(0);
  const refreshTimerRef = useRef<ReturnType<typeof setTimeout> | null>(null);
  const nextRefreshAtRef = useRef<number | null>(null);
  const isRefreshingRef = useRef(false);

  const pathname = usePathname();
  const searchParams = useSearchParams();
  const currentRel = useMemo(() => {
    const qs = searchParams?.toString();
    return qs ? `${pathname}?${qs}` : pathname;
  }, [pathname, searchParams]);

  const clearRefreshTimer = () => {
    if (refreshTimerRef.current) {
      clearTimeout(refreshTimerRef.current as any);
      refreshTimerRef.current = null;
    }
    nextRefreshAtRef.current = null;
  };

  const redirectToIntrospect = useCallback(
    (to?: string) => {
      if (isRedirectingRef.current) return;
      isRedirectingRef.current = true;
      const nextTo = to ?? defaultRedirect ?? currentRel ?? '/';
      window.location.href = `/api/auth/introspect?to=${encodeURIComponent(nextTo)}`;
    },
    [currentRel, defaultRedirect],
  );

  const doRefresh = useCallback(async () => {
    if (isRefreshingRef.current) return;
    isRefreshingRef.current = true;
    try {
      const res = await fetch('/api/auth/refresh', {
        method: 'GET',
        credentials: 'include',
        cache: 'no-store',
        headers: { Accept: 'application/json' },
      });

      if (res.status === 200) {
        // Re-consultar /me para reprogramar el siguiente timer con la nueva expiración
        await fetchMe();
        return;
      }

      // Refresh falló → salir a login
      setSession(null);
      setStatus('unauthenticated');
      clearRefreshTimer();
      redirectToIntrospect();
    } catch {
      setSession(null);
      setStatus('unauthenticated');
      clearRefreshTimer();
      redirectToIntrospect();
    } finally {
      isRefreshingRef.current = false;
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []); // deps vacías: se referenciará fetchMe vía una asignación más abajo

  // Para romper la dependencia circular con fetchMe/scheduleRefresh,
  // guardamos una referencia estable a doRefresh y fetchMe:
  const doRefreshRef = useRef(doRefresh);
  useEffect(() => {
    doRefreshRef.current = doRefresh;
  }, [doRefresh]);

  const scheduleRefresh = useCallback(
    (me: GatewayMeData | null) => {
      if (!me) return;

      // Caso FORZADO (dev/testing) — igual que ya tienes
      if (typeof forceRefreshEveryMs === 'number' && forceRefreshEveryMs > 0) {
        // ... (igual que tu código actual)
        return;
      }

      // === NUEVO: si nos dan remaining_seconds tras un refresh, lo usamos como prioridad ===
      const anyMe: any = me;
      let remainingMsFromMe: number | null = null;
      if (typeof anyMe.remaining_seconds === 'number' && anyMe.remaining_seconds > 0) {
        remainingMsFromMe = anyMe.remaining_seconds * 1000;
      }

      // Caso normal: usa expiración del access si existe; si no, fallback
      let expMs: number | null = null;
      if (anyMe.access_expires_at) {
        const t = Date.parse(anyMe.access_expires_at as string);
        if (!Number.isNaN(t)) expMs = t;
      } else if (anyMe.exp) {
        const t = Number(anyMe.exp) * 1000;
        if (!Number.isNaN(t)) expMs = t;
      }

      const now = Date.now();

      // Si tenemos remaining_seconds, lo preferimos para calcular el próximo refresh
      let target = now + (remainingMsFromMe ?? 0);
      if (!remainingMsFromMe) {
        target = (expMs ?? now + ACCESS_TTL_MS) - REFRESH_EARLY_MS;
      }

      const jitter = Math.floor(Math.random() * 5000); // 0–5s
      const fireAt = Math.max(target - jitter, now + 3000);

      if (nextRefreshAtRef.current && nextRefreshAtRef.current <= fireAt + 500) return;

      if (refreshTimerRef.current) clearTimeout(refreshTimerRef.current as any);
      nextRefreshAtRef.current = fireAt;

      const delay = Math.max(fireAt - now, 3000);
      refreshTimerRef.current = setTimeout(() => {
        nextRefreshAtRef.current = null;
        void doRefreshRef.current();
      }, delay);
    },
    [forceRefreshEveryMs],
  );

  const fetchMe = useCallback(async () => {
    try {
      const res = await fetch(`/api/auth/me?client_id=${encodeURIComponent(client_id)}`, {
        method: 'GET',
        credentials: 'include',
        cache: 'no-store',
        headers: { Accept: 'application/json' },
      });

      if (res.status === 200) {
        const data = (await res.json()) as GatewayMeResponse;

        // Caso éxito con sesión válida
        if (data?.success && data?.data) {
          // === NUEVO: si el backend (introspect/me) dice que access expiró pero refresh activo → refrescar sin redirigir
          if (shouldRefreshFromTokens(data.data)) {
            // No marquemos aún unauthenticated; intentamos refresh
            scheduleRefresh(data.data); // por si trae remaining_seconds / exp
            await doRefreshRef.current(); // fuerza refresh inmediato
            return; // el propio refresh hará el fetchMe de vuelta
          }

          // Caso normal: sesión OK
          setSession(data.data);
          setStatus('authenticated');
          scheduleRefresh(data.data);
          return;
        }
      }

      // No autenticado real → redirigir
      setSession(null);
      setStatus('unauthenticated');
      clearRefreshTimer();
      redirectToIntrospect();
    } catch {
      setSession(null);
      setStatus('unauthenticated');
      clearRefreshTimer();
      redirectToIntrospect();
    }
  }, [client_id, redirectToIntrospect, scheduleRefresh]);

  // Ahora que fetchMe está memorizado, lo exponemos a doRefreshRef para romper el ciclo
  const fetchMeRef = useRef(fetchMe);
  useEffect(() => {
    fetchMeRef.current = fetchMe;
    // re-vincula doRefresh para usar el fetchMe más reciente
    doRefreshRef.current = async () => {
      if (isRefreshingRef.current) return;
      isRefreshingRef.current = true;
      try {
        const res = await fetch('/api/auth/refresh', {
          method: 'GET',
          credentials: 'include',
          cache: 'no-store',
          headers: { Accept: 'application/json' },
        });

        if (res.status === 200) {
          await fetchMeRef.current();
          return;
        }
        setSession(null);
        setStatus('unauthenticated');
        clearRefreshTimer();
        redirectToIntrospect();
      } catch {
        setSession(null);
        setStatus('unauthenticated');
        clearRefreshTimer();
        redirectToIntrospect();
      } finally {
        isRefreshingRef.current = false;
      }
    };
  }, [fetchMe, redirectToIntrospect]);

  const refresh = useCallback(async () => {
    await doRefreshRef.current();
  }, []);

  const logout = useCallback(
    async (opts?: { redirect?: boolean; to?: string }) => {
      try {
        const url = new URL('/api/auth/logout', window.location.origin);
        if (opts?.redirect != null) url.searchParams.set('redirect', String(!!opts.redirect));
        if (opts?.to) {
          const abs = new URL(opts.to, window.location.origin).toString();
          url.searchParams.set('callback_url', toBase64(abs));
        }

        const res = await fetch(url.toString(), {
          method: 'GET',
          credentials: 'include',
          cache: 'no-store',
        });

        if (res.status >= 300 && res.status < 400) {
          const location = res.headers.get('location') ?? '/';
          window.location.href = location;
          return;
        }

        let json: any = null;
        try {
          json = await res.json();
        } catch {}

        setSession(null);
        setStatus('unauthenticated');
        clearRefreshTimer();

        if (json?.success && opts?.redirect === false) return;
        redirectToIntrospect(opts?.to ?? '/');
      } catch {
        setSession(null);
        setStatus('unauthenticated');
        clearRefreshTimer();
        redirectToIntrospect(opts?.to ?? '/');
      }
    },
    [redirectToIntrospect],
  );

  // ===== Efectos =====

  // Chequeo inicial + cleanup
  useEffect(() => {
    fetchMe();
    return () => {
      clearRefreshTimer();
    };
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  // Re-check por eventos de ventana (sin polling constante)
  useEffect(() => {
    const recheckIfIdle = () => {
      const now = Date.now();
      if (now - lastCheckRef.current < 2000) return; // rate-limit 2s
      lastCheckRef.current = now;
      if (!isRedirectingRef.current) fetchMeRef.current();
    };

    const onFocus = () => recheckIfIdle();
    const onOnline = () => recheckIfIdle();
    const onVisible = () => {
      if (document.visibilityState === 'visible') recheckIfIdle();
    };

    window.addEventListener('focus', onFocus);
    window.addEventListener('online', onOnline);
    document.addEventListener('visibilitychange', onVisible);

    return () => {
      window.removeEventListener('focus', onFocus);
      window.removeEventListener('online', onOnline);
      document.removeEventListener('visibilitychange', onVisible);
    };
  }, []);

  // Si quedas unauthenticated en cualquier momento, dispara redirección
  useEffect(() => {
    if (status === 'unauthenticated' && !isRedirectingRef.current) {
      clearRefreshTimer();
      redirectToIntrospect();
    }
  }, [status, redirectToIntrospect]);

  const value = useMemo(
    () => ({
      status,
      session,
      refresh, // puedes llamarlo para forzar rotación manual
      logout,
      hasPermission: (q: string) => hasPermissionImpl(session?.module_permissions, q),
    }),
    [status, session, refresh, logout],
  );

  if (status === 'checking' && showWhileChecking) return <>{showWhileChecking}</>;

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
}
