'use client';

import { useEffect, useState } from 'react';
import { Card, CardHeader, CardTitle, CardContent } from '@/components/ui/card';
import { Loader2, Boxes, CircleCheck, ShieldCheck, Globe } from 'lucide-react';
import { getKeycloakClientsStats } from '@/actions/keycloak/clients/get-clients-stats';

const cards = [
  { key: 'total_clients', title: 'Total de Clientes', icon: Boxes, color: 'text-primary' },
  { key: 'enabled_clients', title: 'Activos', icon: CircleCheck, color: 'text-green-600' },
  { key: 'oauth_clients', title: 'OAuth 2.0', icon: ShieldCheck, color: 'text-blue-600' },
  { key: 'saml_clients', title: 'SAML', icon: Globe, color: 'text-purple-600' },
] as const;

export const KeycloakClientsStatsCards: React.FC = () => {
  const [stats, setStats] = useState<any>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchStats = async () => {
      try {
        const result = await getKeycloakClientsStats();
        if (result.success) {
          setStats(result.data);
        }
      } catch (err) {
        console.error('Error al obtener estadísticas de clientes:', err);
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
            <Icon className={`w-5 h-5 ${color}`} />
          </CardHeader>
          <CardContent>
            <div className="font-bold text-2xl">{stats[key]}</div>
          </CardContent>
        </Card>
      ))}
    </div>
  );
};