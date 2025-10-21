'use client';

import { useEffect, useState } from 'react';
import { Card, CardHeader, CardTitle, CardContent } from '@/components/ui/card';
import { Loader2, Boxes, UsersRound, CircleCheck } from 'lucide-react';
import { fn_get_applications_stats } from '@/actions/applications/fn_get_applications_stats';
import { ApplicationsStatsResponse } from '@/types/applications';

const cards = [
  { key: 'total_applications', title: 'Total de Aplicaciones', icon: Boxes, color: 'text-primary' },
  { key: 'active_applications', title: 'Activas', icon: CircleCheck, color: 'text-green-600' },
  { key: 'total_users', title: 'Usuarios Totales', icon: UsersRound, color: 'text-blue-600' },
] as const;

export const ApplicationsStatsCards: React.FC = () => {
  const [stats, setStats] = useState<ApplicationsStatsResponse | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchStats = async () => {
      try {
        const data = await fn_get_applications_stats();
        setStats(data);
      } catch (err) {
        console.error('Error al obtener estadísticas de aplicaciones:', err);
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
    return <div className="py-12 text-muted-foreground text-center">No se pudieron cargar las estadísticas de aplicaciones.</div>;
  }

  return (
    <div className="gap-4 grid sm:grid-cols-2 lg:grid-cols-3">
      {cards.map(({ key, title, icon: Icon, color }) => (
        <Card key={key} className="bg-card/60 shadow-sm hover:shadow-md backdrop-blur-xl border-border transition-shadow">
          <CardHeader className="flex flex-row justify-between items-center space-y-0 pb-2">
            <CardTitle className="font-medium text-sm">{title}</CardTitle>
            <Icon className={`w-5 h-5 ${color}`} />
          </CardHeader>
          <CardContent>
            <div className="font-bold text-2xl">{stats[key as keyof ApplicationsStatsResponse]}</div>
          </CardContent>
        </Card>
      ))}
    </div>
  );
};
