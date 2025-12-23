'use server';

import { auth } from '@/lib/auth';
import type { ApplicationRoleItem } from '@/types/application-roles';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export interface CreateApplicationRoleInput {
  name: string;
  description?: string;
  application_id: string;
}

export const fn_create_application_role = async (input: CreateApplicationRoleInput): Promise<ApplicationRoleItem> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    const res = await fetch(`${API_BASE_URL}/api/application-roles`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
      body: JSON.stringify(input),
    });

    if (!res.ok) {
      const error = await res.json();
      throw new Error(error.error || `Error al crear rol: ${res.statusText}`);
    }

    const data: ApplicationRoleItem = await res.json();
    return data;
  } catch (err) {
    console.error('Error en fn_create_application_role:', err);
    throw err;
  }
};