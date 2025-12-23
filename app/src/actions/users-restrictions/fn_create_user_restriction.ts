'use server';

import { auth } from '@/lib/auth';
import type { CreateRestrictionInput, UserRestrictionItem } from '@/types/user-restrictions';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export const fn_create_user_restriction = async (input: CreateRestrictionInput): Promise<UserRestrictionItem> => {
  const session = await auth();
  if (!session?.accessToken) throw new Error('No hay sesión activa');

  const res = await fetch(`${API_BASE_URL}/api/user-restrictions`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${session.accessToken}`,
    },
    body: JSON.stringify(input),
  });

  if (!res.ok) {
    const error = await res.json();
    throw new Error(error.error || error.message || 'Error al crear restricción');
  }
  return await res.json();
};