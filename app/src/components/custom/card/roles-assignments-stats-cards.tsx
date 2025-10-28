'use client';

import { useEffect, useState } from 'react';
import { Card, CardHeader, CardTitle, CardContent } from '@/components/ui/card';
import { Loader2, UsersRound, Crown, CircleCheck, CircleX } from 'lucide-react';
import { fn_get_roles_assignments_stats } from '@/actions/roles_assignments/fn_get_roles_assignments_stats';
import type { RolesAssignmentsStatsResponse } from '@/types/roles_assignments';

const cards = [
  { key: 'total_users', title: 'Usuarios Totales', icon: UsersRound, color: 'text-primary' },
  { key: 'admin_users', title: 'Usuarios Admin', icon: Crown, color: 'text-amber-600' },
  { key: 'users_with_roles', title: 'Con Roles', icon: CircleCheck, color: 'text-green-600' },
  { key: 'users_without_roles', title: 'Sin Roles', icon: CircleX, color: 'text-red-600' },
] as const;

export const RolesAssignmentsStatsCards: React.FC = () => {
  const [stats, setStats] = useState<RolesAssignmentsStatsResponse | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchStats = async () => {
      try {
        const data = await fn_get_roles_assignments_stats();
        setStats(data);
      } catch (err) {
        console.error('Error al obtener estadísticas de roles-assignments:', err);
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
    return <div className="py-12 text-muted-foreground text-center">No se pudieron cargar las estadísticas de asignaciones.</div>;
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
            <div className="font-bold text-2xl">{stats[key as keyof RolesAssignmentsStatsResponse]}</div>
          </CardContent>
        </Card>
      ))}
    </div>
  );
};
