'use server';

import { auth } from '@/lib/auth';
import type { DepartmentItem } from '@/types/ubigeos';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export const fn_get_departments = async (): Promise<DepartmentItem[]> => {
  const session = await auth();
  if (!session?.accessToken) throw new Error('No hay sesión activa');

  const res = await fetch(`${API_BASE_URL}/api/ubigeos/departments`, {
    headers: {
      Authorization: `Bearer ${session.accessToken}`,
    },
    cache: 'no-store',
  });

  if (!res.ok) {
    throw new Error('Error al obtener departamentos');
  }

  return res.json();
};
