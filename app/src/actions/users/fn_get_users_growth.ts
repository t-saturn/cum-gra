'use server';

import { auth } from '@/lib/auth';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:8080';

export interface UserGrowthData {
  date: string;
  total: number;
  active: number;
  new: number;
}

export const fn_get_users_growth = async (days: number = 90): Promise<UserGrowthData[]> => {
  try {
    const session = await auth();
    if (!session?.accessToken) {
      throw new Error('No hay sesión activa');
    }

    // Obtener TODOS los usuarios (sin paginación para análisis completo)
    const res = await fetch(`${API_BASE_URL}/api/users?page=1&page_size=10000`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${session.accessToken}`,
      },
      cache: 'no-store',
    });

    if (!res.ok) {
      throw new Error(`Error al obtener usuarios: ${res.statusText}`);
    }

    const data = await res.json();
    const users = data.data || [];

    // Generar datos por día
    const growthData: UserGrowthData[] = [];
    const today = new Date();
    today.setHours(0, 0, 0, 0);

    for (let i = days - 1; i >= 0; i--) {
      const currentDate = new Date(today);
      currentDate.setDate(currentDate.getDate() - i);
      const nextDate = new Date(currentDate);
      nextDate.setDate(nextDate.getDate() + 1);

      // Usuarios creados hasta esta fecha (total acumulado)
      const totalUsers = users.filter((user: any) => {
        const createdAt = new Date(user.created_at);
        return createdAt < nextDate;
      }).length;

      // Usuarios activos hasta esta fecha
      const activeUsers = users.filter((user: any) => {
        const createdAt = new Date(user.created_at);
        return createdAt < nextDate && user.status === 'active' && !user.is_deleted;
      }).length;

      // Usuarios nuevos en este día específico
      const newUsers = users.filter((user: any) => {
        const createdAt = new Date(user.created_at);
        createdAt.setHours(0, 0, 0, 0);
        return createdAt.getTime() === currentDate.getTime();
      }).length;

      growthData.push({
        date: currentDate.toISOString().split('T')[0],
        total: totalUsers,
        active: activeUsers,
        new: newUsers,
      });
    }

    return growthData;
  } catch (err) {
    console.error('Error en fn_get_users_growth:', err);
    throw err;
  }
};