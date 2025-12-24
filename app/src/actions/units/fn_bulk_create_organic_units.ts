'use server';

import { auth } from '@/lib/auth';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export interface BulkUploadError {
  row: number;
  field?: string;
  message: string;
}

export interface BulkUploadResponse {
  created: number;
  skipped: number;
  failed: number;
  errors: BulkUploadError[];
  details: Array<{ id: number; name: string; acronym: string }>;
}

export const fn_bulk_create_organic_units = async (
  formData: FormData
): Promise<BulkUploadResponse> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    const res = await fetch(`${API_BASE_URL}/api/organic-units/bulk`, {
      method: 'POST',
      headers: {
        Authorization: `Bearer ${session.accessToken}`,
      },
      body: formData,
    });

    if (!res.ok) {
      const errorData = await res.json().catch(() => ({}));
      throw new Error(errorData.error || `Error en carga masiva: ${res.statusText}`);
    }

    return await res.json();
  } catch (err) {
    console.error('Error en fn_bulk_create_organic_units:', err);
    throw err;
  }
};