'use client';

import { useEffect, useState } from 'react';
import { Card, CardHeader, CardTitle, CardContent } from '@/components/ui/card';
import { Loader2, UsersRound, UserRoundCheck, UserRoundX, UserRoundPlus } from 'lucide-react';
import { UsersStatsResponse } from '@/types/users';
import { getUsersStats } from '@/actions/users/fn_get_user_stats';

const cards = [
  { key: 'total_users', title: 'Total de Usuarios', icon: UsersRound, color: 'text-primary' },
  { key: 'active_users', title: 'Activos', icon: UserRoundCheck, color: 'text-green-600' },
  { key: 'suspended_users', title: 'Suspendidos', icon: UserRoundX, color: 'text-yellow-600' },
  { key: 'new_users_last_month', title: 'Nuevos (último mes)', icon: UserRoundPlus, color: 'text-blue-600' },
] as const;

export const UsersStatsCards: React.FC = () => {
  const [stats, setStats] = useState<UsersStatsResponse | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchStats = async () => {
      try {
        const data = await getUsersStats();
        setStats(data);
      } catch (err) {
        console.error('Error al obtener estadísticas:', err);
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

  if (!stats) return <div className="py-12 text-muted-foreground text-center">No se pudieron cargar las estadísticas de usuarios.</div>;

  return (
    <div className="gap-4 grid sm:grid-cols-2 lg:grid-cols-4">
      {cards.map(({ key, title, icon: Icon, color }) => (
        <Card key={key} className="bg-card/60 shadow-sm hover:shadow-md backdrop-blur-xl border-border transition-shadow">
          <CardHeader className="flex flex-row justify-between items-center space-y-0 pb-2">
            <CardTitle className="font-medium text-sm">{title}</CardTitle>
            <Icon className={`w-5 h-5 ${color}`} />
          </CardHeader>
          <CardContent>
            <div className="font-bold text-2xl">{stats[key as keyof UsersStatsResponse]}</div>
          </CardContent>
        </Card>
      ))}
    </div>
  );
};
