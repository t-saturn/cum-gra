'use client';

import { useEffect, useState } from 'react';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Shield, ShieldCheck, Trash2, Package, Users } from 'lucide-react';
import { fn_get_application_roles_stats } from '@/actions/application-roles/fn_get_application_roles_stats';
import type { ApplicationRolesStatsResponse } from '@/types/application-roles';

export function ApplicationRolesStatsCards() {
  const [stats, setStats] = useState<ApplicationRolesStatsResponse>({
    total_roles: 0,
    active_roles: 0,
    deleted_roles: 0,
    roles_with_modules: 0,
    roles_with_users: 0,
  });
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const loadStats = async () => {
      try {
        const data = await fn_get_application_roles_stats();
        setStats(data);
      } catch (error) {
        console.error('Error loading application roles stats:', error);
      } finally {
        setLoading(false);
      }
    };
    loadStats();
  }, []);

  const cards = [
    {
      title: 'Total Roles',
      value: stats.total_roles,
      icon: Shield,
      color: 'text-blue-500',
      bgColor: 'bg-blue-500/10',
    },
    {
      title: 'Roles Activos',
      value: stats.active_roles,
      icon: ShieldCheck,
      color: 'text-green-500',
      bgColor: 'bg-green-500/10',
    },
    {
      title: 'Con Módulos',
      value: stats.roles_with_modules,
      icon: Package,
      color: 'text-purple-500',
      bgColor: 'bg-purple-500/10',
    },
    {
      title: 'Con Usuarios',
      value: stats.roles_with_users,
      icon: Users,
      color: 'text-orange-500',
      bgColor: 'bg-orange-500/10',
    },
  ];

  if (loading) {
    return (
      <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
        {[1, 2, 3, 4].map((i) => (
          <Card key={i} className="animate-pulse">
            <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
              <div className="h-4 w-24 bg-muted rounded" />
              <div className="h-8 w-8 bg-muted rounded" />
            </CardHeader>
            <CardContent>
              <div className="h-8 w-16 bg-muted rounded" />
            </CardContent>
          </Card>
        ))}
      </div>
    );
  }

  return (
    <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
      {cards.map((card) => {
        const Icon = card.icon;
        return (
          <Card key={card.title}>
            <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
              <CardTitle className="text-sm font-medium">{card.title}</CardTitle>
              <div className={`${card.bgColor} p-2 rounded-lg`}>
                <Icon className={`h-4 w-4 ${card.color}`} />
              </div>
            </CardHeader>
            <CardContent>
              <div className="text-2xl font-bold">{card.value.toLocaleString()}</div>
            </CardContent>
          </Card>
        );
      })}
    </div>
  );
}