'use server';

const API_BASE_URL = process.env.API_BASE_URL ?? 'http://localhost:9191';

export async function fn_get_units_stats() {
  try {
    const res = await fetch(`${API_BASE_URL}/units/stats`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
      cache: 'no-store',
    });

    if (!res.ok) throw new Error(`Error HTTP ${res.status}: ${res.statusText}`);

    const data = await res.json();
    return data;
  } catch (error) {
    console.error('Error obteniendo estadísticas de unidades orgánicas:', error);
    throw error;
  }
}
