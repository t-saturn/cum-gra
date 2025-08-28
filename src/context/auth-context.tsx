'use client';

import { createContext, useContext } from 'react';
import type { AuthStatus, GatewayMeData } from '@/types/auth';

export type AuthContextValue = {
  status: AuthStatus;
  session: GatewayMeData | null;
  refresh: () => Promise<void>;
  logout: (opts?: { redirect?: boolean; to?: string }) => Promise<void>;
  hasPermission: (query: string) => boolean; // flexible: "module" o "module:action"
};

export const AuthContext = createContext<AuthContextValue | undefined>(undefined);

export function useAuthContext(): AuthContextValue {
  const ctx = useContext(AuthContext);
  if (!ctx) throw new Error('useAuthContext debe usarse dentro de <SessionProvider>');
  return ctx;
}
