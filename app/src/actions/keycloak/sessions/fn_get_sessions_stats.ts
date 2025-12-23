'use server';

import { fn_get_all_sessions } from './fn_get_all_sessions';
import type { SessionsStatsResponse } from '@/types/sessions';

export async function fn_get_sessions_stats(): Promise<SessionsStatsResponse> {
  try {
    const sessions = await fn_get_all_sessions();
    
    const uniqueUsers = new Set(sessions.map(s => s.userId)).size;
    
    const oneHourAgo = Date.now() - (60 * 60 * 1000);
    const activeLastHour = sessions.filter(s => s.lastAccess > oneHourAgo).length;
    
    // Contar sesiones por cliente
    const sessionsByClient: Record<string, number> = {};
    sessions.forEach(session => {
      const key = session.clientName || session.clientId;
      sessionsByClient[key] = (sessionsByClient[key] || 0) + 1;
    });

    return {
      total_sessions: sessions.length,
      unique_users: uniqueUsers,
      active_last_hour: activeLastHour,
      sessions_by_client: sessionsByClient,
    };
  } catch (error) {
    console.error('Error calculating sessions stats:', error);
    return {
      total_sessions: 0,
      unique_users: 0,
      active_last_hour: 0,
      sessions_by_client: {},
    };
  }
}