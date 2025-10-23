'use client';

import { useEffect, useState } from 'react';
import { Card, CardHeader, CardTitle, CardContent } from '@/components/ui/card';
import { Loader2, ShieldAlert, UserRoundX, CircleAlert, CircleX } from 'lucide-react';
import { fn_get_users_restrictions_stats } from '@/actions/users_restrictions/fn_get_users_restrictions_stats';
import type { UsersRestrictionsStatsResponse } from '@/types/users_restrictions';

const cards = [
  { key: 'total_restrictions', title: 'Total de Restricciones', icon: CircleAlert, color: 'text-primary' },
  { key: 'active_restrictions', title: 'Activas', icon: ShieldAlert, color: 'text-green-600' },
  { key: 'restricted_users', title: 'Usuarios Restringidos', icon: UserRoundX, color: 'text-amber-600' },
  { key: 'deleted_restrictions', title: 'Eliminadas', icon: CircleX, color: 'text-red-600' },
] as const;

export const UsersRestrictionsStatsCards: React.FC = () => {
  const [stats, setStats] = useState<UsersRestrictionsStatsResponse | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchStats = async () => {
      try {
        const data = await fn_get_users_restrictions_stats();
        setStats(data);
      } catch (err) {
        console.error('Error al obtener estadísticas de restricciones de usuarios:', err);
      } finally {
        setLoading(false);
      }
    };
    fetchStats();
  }, []);

  if (loading) {
    return (
      <div className="flex justify-center items-center py-12">
        <Loader2 className="w-8 h-8 text-primary animate-spin" />
      </div>
    );
  }

  if (!stats) {
    return <div className="py-12 text-muted-foreground text-center">No se pudieron cargar las estadísticas.</div>;
  }

  return (
    <div className="gap-4 grid sm:grid-cols-2 lg:grid-cols-4">
      {cards.map(({ key, title, icon: Icon, color }) => (
        <Card key={key} className="bg-card/60 shadow-sm hover:shadow-md backdrop-blur-xl border-border transition-shadow">
          <CardHeader className="flex flex-row justify-between items-center space-y-0 pb-2">
            <CardTitle className="font-medium text-sm">{title}</CardTitle>
            <Icon className={`h-5 w-5 ${color}`} />
          </CardHeader>
          <CardContent>
            <div className="font-bold text-2xl">{stats[key as keyof UsersRestrictionsStatsResponse]}</div>
          </CardContent>
        </Card>
      ))}
    </div>
  );
};
