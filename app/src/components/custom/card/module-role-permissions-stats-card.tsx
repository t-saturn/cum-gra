'use client';

import { useEffect, useState } from 'react';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Key, KeyRound, Package, Shield, Trash2 } from 'lucide-react';
import { fn_get_module_role_permissions_stats } from '@/actions/module-role-permissions/fn_get_module_role_permissions_stats';
import type { ModuleRolePermissionsStatsResponse } from '@/types/module-role-permissions';

export function ModuleRolePermissionsStatsCards() {
  const [stats, setStats] = useState<ModuleRolePermissionsStatsResponse>({
    total_permissions: 0,
    active_permissions: 0,
    deleted_permissions: 0,
    unique_modules: 0,
    unique_roles: 0,
    permissions_by_type: {
      read: 0,
      write: 0,
      execute: 0,
      delete: 0,
      admin: 0,
    },
  });
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const loadStats = async () => {
      try {
        const data = await fn_get_module_role_permissions_stats();
        setStats(data);
      } catch (error) {
        console.error('Error loading module role permissions stats:', error);
      } finally {
        setLoading(false);
      }
    };
    loadStats();
  }, []);

  const cards = [
    {
      title: 'Total Permisos',
      value: stats.total_permissions,
      icon: Key,
      color: 'text-blue-500',
      bgColor: 'bg-blue-500/10',
    },
    {
      title: 'Permisos Activos',
      value: stats.active_permissions,
      icon: KeyRound,
      color: 'text-green-500',
      bgColor: 'bg-green-500/10',
    },
    {
      title: 'Módulos Únicos',
      value: stats.unique_modules,
      icon: Package,
      color: 'text-purple-500',
      bgColor: 'bg-purple-500/10',
    },
    {
      title: 'Roles Únicos',
      value: stats.unique_roles,
      icon: Shield,
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