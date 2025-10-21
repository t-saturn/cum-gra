'use client';

import { useEffect, useState } from 'react';
import { Card, CardHeader, CardTitle, CardContent } from '@/components/ui/card';
import { Boxes, Loader2, UsersRound, CircleX, CircleCheck } from 'lucide-react';
import { fn_get_modules_stats } from '@/actions/modules/fn_get_modules_stats';
import { ModulesStatsResponse } from '@/types/modules';

const cards = [
  { key: 'total_modules', title: 'Total de Módulos', icon: Boxes, color: 'text-primary' },
  { key: 'active_modules', title: 'Activos', icon: CircleCheck, color: 'text-green-600' },
  { key: 'deleted_modules', title: 'Eliminados', icon: CircleX, color: 'text-red-600' },
  { key: 'total_users', title: 'Usuarios Totales', icon: UsersRound, color: 'text-blue-600' },
] as const;

export const ModulesStatsCards: React.FC = () => {
  const [stats, setStats] = useState<ModulesStatsResponse | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchStats = async () => {
      try {
        const data = await fn_get_modules_stats();
        setStats(data);
      } catch (err) {
        console.error('Error al obtener estadísticas de módulos:', err);
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
    return <div className="py-12 text-muted-foreground text-center">No se pudieron cargar las estadísticas de módulos.</div>;
  }

  return (
    <div className="gap-4 grid sm:grid-cols-2 lg:grid-cols-4">
      {cards.map(({ key, title, icon: Icon, color }) => (
        <Card key={key} className="bg-card/60 shadow-sm hover:shadow-md backdrop-blur-xl border-border transition-shadow">
          <CardHeader className="flex flex-row justify-between items-center space-y-0 pb-2">
            <CardTitle className="font-medium text-sm">{title}</CardTitle>
            <Icon className={`w-5 h-5 ${color}`} />
          </CardHeader>
          <CardContent>
            <div className="font-bold text-2xl">{stats[key as keyof ModulesStatsResponse]}</div>
          </CardContent>
        </Card>
      ))}
    </div>
  );
};
