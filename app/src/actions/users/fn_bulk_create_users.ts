'use server';

import { auth } from '@/lib/auth';
import type { BulkCreateUsersResponse } from '@/types/users';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export const fn_bulk_create_users = async (file: File): Promise<BulkCreateUsersResponse> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    const formData = new FormData();
    formData.append('file', file);

    const res = await fetch(`${API_BASE_URL}/api/users/bulk`, {
      method: 'POST',
      headers: {
        Authorization: `Bearer ${session.accessToken}`,
      },
      body: formData,
    });

    if (!res.ok) {
      const error = await res.json();
      throw new Error(error.error || `Error en carga masiva: ${res.statusText}`);
    }

    const data: BulkCreateUsersResponse = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_bulk_create_users:', err);
    throw err;
  }
};