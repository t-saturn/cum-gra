'use server';

import { auth } from '@/lib/auth';
import type { ApplicationsListResponse } from '@/types/applications';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8000';

export const fn_get_applications = async (
  page: number = 1,
  pageSize: number = 10,
  isDeleted: boolean = false
): Promise<ApplicationsListResponse> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    const params = new URLSearchParams({
      page: page.toString(),
      page_size: pageSize.toString(),
      is_deleted: isDeleted.toString(),
    });

    // CORRECCIÓN: Cambiar fetch`...` por fetch(`...`)
    const res = await fetch(`${API_BASE_URL}/api/applications?${params}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
      cache: 'no-store',
    });

    if (!res.ok) {
      throw new Error(`Error al obtener aplicaciones: ${res.statusText}`);
    }

    const data: ApplicationsListResponse = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_get_applications:', err);
    throw err;
  }
};