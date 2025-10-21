'use client';

import { useEffect, useState } from 'react';
import { Card, CardHeader, CardTitle, CardContent } from '@/components/ui/card';
import { Loader2, CircleQuestionMark, CircleCheck, CircleX, UsersRound } from 'lucide-react';
import { fn_get_positions_stats } from '@/actions/positions/fn_get_positions_stats';

type PositionsStatsResponse = {
  total_positions: number;
  active_positions: number;
  deleted_positions: number;
  assigned_employees: number;
};

const cards = [
  { key: 'total_positions', title: 'Total de Cargos', icon: CircleQuestionMark, color: 'text-primary' },
  { key: 'active_positions', title: 'Activos', icon: CircleCheck, color: 'text-green-600' },
  { key: 'deleted_positions', title: 'Eliminados', icon: CircleX, color: 'text-red-600' },
  { key: 'assigned_employees', title: 'Asignados', icon: UsersRound, color: 'text-blue-600' },
] as const;

export const PositionsStatsCards: React.FC = () => {
  const [stats, setStats] = useState<PositionsStatsResponse | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchStats = async () => {
      try {
        const data = await fn_get_positions_stats();
        setStats(data);
      } catch (err) {
        console.error('Error al obtener estadísticas de posiciones:', err);
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

  if (!stats) return <div className="py-12 text-muted-foreground text-center">No se pudieron cargar las estadísticas de cargos.</div>;

  return (
    <div className="gap-4 grid sm:grid-cols-2 lg:grid-cols-4">
      {cards.map(({ key, title, icon: Icon, color }) => (
        <Card key={key} className="bg-card/60 shadow-sm hover:shadow-md backdrop-blur-xl border-border transition-shadow">
          <CardHeader className="flex flex-row justify-between items-center space-y-0 pb-2">
            <CardTitle className="font-medium text-sm">{title}</CardTitle>
            <Icon className={`w-5 h-5 ${color}`} />
          </CardHeader>
          <CardContent>
            <div className="font-bold text-2xl">{stats[key as keyof PositionsStatsResponse]}</div>
          </CardContent>
        </Card>
      ))}
    </div>
  );
};
