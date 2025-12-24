'use server';
import { auth } from '@/lib/auth';
const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export const fn_delete_user_restriction = async (id: string) => {
  const session = await auth();
  if (!session?.accessToken) throw new Error('No hay sesión activa');

  const res = await fetch(`${API_BASE_URL}/api/user-restrictions/${id}`, {
    method: 'DELETE',
    headers: { Authorization: `Bearer ${session.accessToken}` },
  });

  if (!res.ok) throw new Error('Error al eliminar restricción');
  return await res.json();
};

export const fn_restore_user_restriction = async (id: string) => {
  const session = await auth();
  if (!session?.accessToken) throw new Error('No hay sesión activa');

  const res = await fetch(`${API_BASE_URL}/api/user-restrictions/${id}/restore`, {
    method: 'PATCH',
    headers: { Authorization: `Bearer ${session.accessToken}` },
  });

  if (!res.ok) throw new Error('Error al restaurar restricción');
  return await res.json();
};