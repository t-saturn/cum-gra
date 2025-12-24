'use server';

import { auth } from '@/lib/auth';
import type { BulkCreateRestrictionsInput, BulkCreateResponse } from '@/types/user-restrictions';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export const fn_bulk_create_user_restrictions = async (input: BulkCreateRestrictionsInput): Promise<BulkCreateResponse> => {
  const session = await auth();
  if (!session?.accessToken) throw new Error('No hay sesión activa');

  const res = await fetch(`${API_BASE_URL}/api/user-restrictions/bulk-create`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${session.accessToken}`,
    },
    body: JSON.stringify(input),
  });

  if (!res.ok) {
    const error = await res.json();
    throw new Error(error.error || 'Error en creación masiva');
  }
  return await res.json();
};