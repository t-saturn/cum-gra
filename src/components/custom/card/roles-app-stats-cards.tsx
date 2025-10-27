'use client';

import { useEffect, useState } from 'react';
import { Card, CardHeader, CardTitle, CardContent } from '@/components/ui/card';
import { Loader2, Shield, CircleCheck, Crown, UsersRound } from 'lucide-react';
import { fn_get_roles_apps_stats } from '@/actions/roles_app/fn_get_roles_app_stats';
import type { RolesAppsStatsResponse } from '@/types/roles_app';

const cards = [
  { key: 'total_roles', title: 'Total de Roles', icon: Shield, color: 'text-primary' },
  { key: 'active_roles', title: 'Activos', icon: CircleCheck, color: 'text-green-600' },
  { key: 'admin_roles', title: 'Roles Admin', icon: Crown, color: 'text-amber-600' },
  { key: 'assigned_users', title: 'Usuarios Asignados', icon: UsersRound, color: 'text-blue-600' },
] as const;

export const RolesAppsStatsCards: React.FC = () => {
  const [stats, setStats] = useState<RolesAppsStatsResponse | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchStats = async () => {
      try {
        const data = await fn_get_roles_apps_stats();
        setStats(data);
      } catch (err) {
        console.error('Error al obtener estadísticas de roles-app:', err);
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
    return <div className="py-12 text-muted-foreground text-center">No se pudieron cargar las estadísticas de roles.</div>;
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
            <div className="font-bold text-2xl">{stats[key as keyof RolesAppsStatsResponse]}</div>
          </CardContent>
        </Card>
      ))}
    </div>
  );
};
