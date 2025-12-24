'use server';

import { auth } from '@/lib/auth';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export const fn_download_organic_units_template = async (): Promise<Blob> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    const res = await fetch(`${API_BASE_URL}/api/organic-units/template`, {
      method: 'GET',
      headers: {
        Authorization: `Bearer ${session.accessToken}`,
      },
    });

    if (!res.ok) {
      throw new Error(`Error al descargar plantilla: ${res.statusText}`);
    }

    const arrayBuffer = await res.arrayBuffer();
    return new Blob([arrayBuffer], {
      type: 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet',
    });
  } catch (err) {
    console.error('Error en fn_download_organic_units_template:', err);
    throw err;
  }
};