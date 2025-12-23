'use client';

import { useEffect, useState } from 'react';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Activity, Users, Clock, TrendingUp } from 'lucide-react';
import { fn_get_sessions_stats } from '@/actions/keycloak/sessions/fn_get_sessions_stats';
import type { SessionsStatsResponse } from '@/types/sessions';

export function SessionsStatsCards() {
  const [stats, setStats] = useState<SessionsStatsResponse>({
    total_sessions: 0,
    unique_users: 0,
    active_last_hour: 0,
    sessions_by_client: {},
  });
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const loadStats = async () => {
      try {
        const data = await fn_get_sessions_stats();
        setStats(data);
      } catch (error) {
        console.error('Error loading sessions stats:', error);
      } finally {
        setLoading(false);
      }
    };
    loadStats();
  }, []);

  // Obtener el cliente con más sesiones
  const topClient = Object.keys(stats.sessions_by_client).length > 0
    ? Object.entries(stats.sessions_by_client)
        .sort((a, b) => b[1] - a[1])[0]
    : null;

  const cards = [
    {
      title: 'Sesiones Activas',
      value: stats.total_sessions,
      icon: Activity,
      color: 'text-blue-500',
      bgColor: 'bg-blue-500/10',
    },
    {
      title: 'Usuarios Únicos',
      value: stats.unique_users,
      icon: Users,
      color: 'text-green-500',
      bgColor: 'bg-green-500/10',
    },
    {
      title: 'Activos (Última Hora)',
      value: stats.active_last_hour,
      icon: Clock,
      color: 'text-orange-500',
      bgColor: 'bg-orange-500/10',
    },
    {
      title: 'Cliente Principal',
      value: topClient ? `${topClient[0]} (${topClient[1]})` : 'N/A',
      icon: TrendingUp,
      color: 'text-purple-500',
      bgColor: 'bg-purple-500/10',
      isText: true,
    },
  ];

  if (loading) {
    return (
      <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
        {[1, 2, 3, 4].map((i) => (
          <Card key={i} className="animate-pulse">
            <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
              <div className="h-4 w-24 bg-muted rounded" />
              <div className="h-8 w-8 bg-muted rounded" />
            </CardHeader>
            <CardContent>
              <div className="h-8 w-16 bg-muted rounded" />
            </CardContent>
          </Card>
        ))}
      </div>
    );
  }

  return (
    <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
      {cards.map((card) => {
        const Icon = card.icon;
        return (
          <Card key={card.title}>
            <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
              <CardTitle className="text-sm font-medium">{card.title}</CardTitle>
              <div className={`${card.bgColor} p-2 rounded-lg`}>
                <Icon className={`h-4 w-4 ${card.color}`} />
              </div>
            </CardHeader>
            <CardContent>
              <div className={`${card.isText ? 'text-lg' : 'text-2xl'} font-bold truncate`}>
                {card.isText ? card.value : (card.value as number).toLocaleString()}
              </div>
            </CardContent>
          </Card>
        );
      })}
    </div>
  );
}