'use server';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:9191';

export async function fn_get_positions_stats() {
  try {
    const res = await fetch(`${API_BASE_URL}/positions/stats`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
      // Opcional: puedes agregar cache: 'no-store' si quieres datos siempre frescos
      cache: 'no-store',
    });

    if (!res.ok) throw new Error(`Error HTTP ${res.status}: ${res.statusText}`);

    const data = await res.json();
    return data;
  } catch (error) {
    console.error('Error obteniendo estadísticas de posiciones:', error);
    throw error;
  }
}
