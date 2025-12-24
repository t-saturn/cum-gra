'use client';

import { useEffect, useState } from 'react';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { ShieldBan, Users, ListX, Trash2 } from 'lucide-react';
import type { UserRestrictionsStatsResponse } from '@/types/user-restrictions';
import { fn_get_user_restrictions_stats } from '@/actions/users-restrictions/fn_get_user_restrictions_stats';

export function UserRestrictionsStatsCards() {
  const [stats, setStats] = useState<UserRestrictionsStatsResponse>({
    total_restrictions: 0,
    active_restrictions: 0,
    restricted_users: 0,
    deleted_restrictions: 0,
  });
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const loadStats = async () => {
      try {
        const data = await fn_get_user_restrictions_stats();
        setStats(data);
      } catch (error) {
        console.error(error);
      } finally {
        setLoading(false);
      }
    };
    loadStats();
  }, []);

  const cards = [
    { title: 'Total Restricciones', value: stats.total_restrictions, icon: ListX, color: 'text-blue-500', bg: 'bg-blue-500/10' },
    { title: 'Activas', value: stats.active_restrictions, icon: ShieldBan, color: 'text-red-500', bg: 'bg-red-500/10' },
    { title: 'Usuarios Restringidos', value: stats.restricted_users, icon: Users, color: 'text-orange-500', bg: 'bg-orange-500/10' },
    { title: 'Eliminadas', value: stats.deleted_restrictions, icon: Trash2, color: 'text-gray-500', bg: 'bg-gray-500/10' },
  ];

  if (loading) return <div className="grid gap-4 md:grid-cols-4 animate-pulse">{[1,2,3,4].map(i => <div key={i} className="h-24 bg-muted rounded-xl" />)}</div>;

  return (
    <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
      {cards.map((c) => {
        const Icon = c.icon;
        return (
          <Card key={c.title}>
            <CardHeader className="flex flex-row items-center justify-between pb-2">
              <CardTitle className="text-sm font-medium">{c.title}</CardTitle>
              <div className={`${c.bg} p-2 rounded-lg`}><Icon className={`h-4 w-4 ${c.color}`} /></div>
            </CardHeader>
            <CardContent><div className="text-2xl font-bold">{c.value}</div></CardContent>
          </Card>
        );
      })}
    </div>
  );
}