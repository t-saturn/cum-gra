'use client';

import { useEffect, useState } from 'react';
import { Card, CardHeader, CardTitle, CardContent } from '@/components/ui/card';
import { Building2, Building, Trash2, Users, Loader2 } from 'lucide-react';
import { fn_get_units_stats } from '@/actions/units/fn_get_units_stats';
import { OrganicUnitsStatsResponse } from '@/types/units';

const cards = [
  { key: 'total_organic_units', title: 'Total de Unidades', icon: Building2, color: 'text-primary' },
  { key: 'active_organic_units', title: 'Activas', icon: Building, color: 'text-green-600' },
  { key: 'deleted_organic_units', title: 'Eliminadas', icon: Trash2, color: 'text-red-600' },
  { key: 'total_employees', title: 'Empleados Totales', icon: Users, color: 'text-blue-600' },
] as const;

export const OrganicUnitsStatsCards: React.FC = () => {
  const [stats, setStats] = useState<OrganicUnitsStatsResponse | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchStats = async () => {
      try {
        const data = await fn_get_units_stats();
        setStats(data);
      } catch (err) {
        console.error('Error al obtener estadísticas de unidades orgánicas:', err);
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
    return <div className="py-12 text-muted-foreground text-center">No se pudieron cargar las estadísticas de unidades orgánicas.</div>;
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
            <div className="font-bold text-2xl">{stats[key as keyof OrganicUnitsStatsResponse]}</div>
          </CardContent>
        </Card>
      ))}
    </div>
  );
};
