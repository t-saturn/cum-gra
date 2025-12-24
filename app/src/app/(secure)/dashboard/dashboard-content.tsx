'use client';

import { useEffect, useState } from 'react';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Users, Building2, Shield, FileKey, Activity, TrendingUp, Clock, BarChart3 } from 'lucide-react';
import { fn_get_user_stats } from '@/actions/users/fn_get_user_stats';
import { fn_get_units_stats } from '@/actions/units/fn_get_units_stats';
import { fn_get_positions_stats } from '@/actions/positions/fn_get_positions_stats';
import { fn_get_applications_stats } from '@/actions/applications/fn_get_applications_stats';
import { fn_get_application_roles_stats } from '@/actions/application-roles/fn_get_application_roles_stats';
import { fn_get_modules_stats } from '@/actions/modules/fn_get_modules_stats';
import { fn_get_user_application_roles_stats } from '@/actions/user-application-roles/fn_get_user_application_roles_stats';
import { fn_get_user_restrictions_stats } from '@/actions/users-restrictions/fn_get_user_restrictions_stats';
import { fn_get_sessions_stats } from '@/actions/keycloak/sessions/fn_get_sessions_stats';
import type { UsersStatsResponse } from '@/types/users';
import { UsersGrowthChart } from '@/components/custom/chart/users-growth-chart';
import { ApplicationsDistributionChart } from '@/components/custom/chart/applications-distribution-chart';
import { RolesAssignmentChart } from '@/components/custom/chart/roles-assignment-chart';
import { SessionsActivityChart } from '@/components/custom/chart/sessions-activity-chart';
import { useSession } from 'next-auth/react';

export default function DashboardContent() {
  const { data: session } = useSession();
  const [greeting, setGreeting] = useState('');
  const [loading, setLoading] = useState(true);

  const [userStats, setUserStats] = useState<UsersStatsResponse | null>(null);
  const [unitsStats, setUnitsStats] = useState<any>(null);
  const [positionsStats, setPositionsStats] = useState<any>(null);
  const [appsStats, setAppsStats] = useState<any>(null);
  const [rolesStats, setRolesStats] = useState<any>(null);
  const [modulesStats, setModulesStats] = useState<any>(null);
  const [assignmentsStats, setAssignmentsStats] = useState<any>(null);
  const [restrictionsStats, setRestrictionsStats] = useState<any>(null);
  const [sessionsStats, setSessionsStats] = useState<any>(null);

  useEffect(() => {
    const loadDashboardData = async () => {
      try {
        setLoading(true);

        // Determinar saludo según hora del día
        const hour = new Date().getHours();
        if (hour < 12) setGreeting('Buenos días');
        else if (hour < 19) setGreeting('Buenas tardes');
        else setGreeting('Buenas noches');

        // Cargar todas las estadísticas en paralelo
        const [
          users,
          units,
          positions,
          apps,
          roles,
          modules,
          assignments,
          restrictions,
          sessions,
        ] = await Promise.all([
          fn_get_user_stats(),
          fn_get_units_stats(),
          fn_get_positions_stats(),
          fn_get_applications_stats(),
          fn_get_application_roles_stats(),
          fn_get_modules_stats(),
          fn_get_user_application_roles_stats(),
          fn_get_user_restrictions_stats(),
          fn_get_sessions_stats(),
        ]);

        setUserStats(users);
        setUnitsStats(units);
        setPositionsStats(positions);
        setAppsStats(apps);
        setRolesStats(roles);
        setModulesStats(modules);
        setAssignmentsStats(assignments);
        setRestrictionsStats(restrictions);
        setSessionsStats(sessions);
      } catch (error) {
        console.error('Error loading dashboard data:', error);
      } finally {
        setLoading(false);
      }
    };

    loadDashboardData();
  }, []);

  if (loading) {
    return <DashboardSkeleton />;
  }

  // Obtener nombre del usuario de la sesión
  const userName = session?.user?.name || 'Usuario';

  return (
    <div className="space-y-6">
      {/* Saludo personalizado */}
      <div className="flex flex-col space-y-2">
        <h1 className="text-3xl font-bold tracking-tight">
          {greeting}, {userName}
        </h1>
        <p className="text-muted-foreground">
          Aquí está el resumen de tu sistema de gestión
        </p>
      </div>

      {/* Cards de estadísticas principales */}
      <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
        <StatCard
          title="Usuarios Totales"
          value={userStats?.total_users || 0}
          description={`${userStats?.active_users || 0} activos`}
          icon={Users}
          trend={`+${userStats?.new_users_last_month || 0} este mes`}
          color="text-blue-500"
        />
        <StatCard
          title="Unidades Orgánicas"
          value={unitsStats?.total_organic_units || 0}
          description={`${unitsStats?.active_organic_units || 0} activas`}
          icon={Building2}
          trend={`${unitsStats?.total_employees || 0} empleados`}
          color="text-green-500"
        />
        <StatCard
          title="Aplicaciones"
          value={appsStats?.total_applications || 0}
          description={`${appsStats?.active_applications || 0} activas`}
          icon={FileKey}
          trend={`${appsStats?.total_users || 0} usuarios`}
          color="text-purple-500"
        />
        <StatCard
          title="Sesiones Activas"
          value={sessionsStats?.total_sessions || 0}
          description="Usuarios conectados"
          icon={Activity}
          trend={`${sessionsStats?.unique_users || 0} únicos`}
          color="text-orange-500"
        />
      </div>

      {/* Cards secundarias */}
      <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
        <MiniStatCard
          title="Roles de Aplicación"
          value={rolesStats?.total_roles || 0}
          icon={Shield}
          color="bg-blue-500/10 text-blue-500"
        />
        <MiniStatCard
          title="Módulos"
          value={modulesStats?.total_modules || 0}
          icon={BarChart3}
          color="bg-green-500/10 text-green-500"
        />
        <MiniStatCard
          title="Asignaciones de Roles"
          value={assignmentsStats?.total_assignments || 0}
          icon={TrendingUp}
          color="bg-purple-500/10 text-purple-500"
        />
        <MiniStatCard
          title="Restricciones Activas"
          value={restrictionsStats?.active_restrictions || 0}
          icon={Clock}
          color="bg-orange-500/10 text-orange-500"
        />
      </div>

      {/* Gráficos principales */}
      <div className="grid gap-4 md:grid-cols-2">
        <UsersGrowthChart />
        <ApplicationsDistributionChart appsStats={appsStats} rolesStats={rolesStats} />
      </div>

      <div className="grid gap-4 md:grid-cols-2">
        <RolesAssignmentChart assignmentsStats={assignmentsStats} />
        <SessionsActivityChart sessionsStats={sessionsStats} />
      </div>

      {/* Estadísticas detalladas */}
      <div className="grid gap-4 md:grid-cols-3">
        <DetailedStatCard
          title="Posiciones Estructurales"
          stats={[
            { label: 'Total', value: positionsStats?.total_positions || 0 },
            { label: 'Activas', value: positionsStats?.active_positions || 0 },
            { label: 'Empleados asignados', value: positionsStats?.assigned_employees || 0 },
          ]}
          icon={Shield}
        />
        <DetailedStatCard
          title="Gestión de Usuarios"
          stats={[
            { label: 'Usuarios suspendidos', value: userStats?.suspended_users || 0 },
            { label: 'Nuevos este mes', value: userStats?.new_users_last_month || 0 },
            { label: 'Total en el sistema', value: userStats?.total_users || 0 },
          ]}
          icon={Users}
        />
        <DetailedStatCard
          title="Restricciones"
          stats={[
            { label: 'Total', value: restrictionsStats?.total_restrictions || 0 },
            { label: 'Activas', value: restrictionsStats?.active_restrictions || 0 },
            { label: 'Usuarios restringidos', value: restrictionsStats?.restricted_users || 0 },
          ]}
          icon={Clock}
        />
      </div>
    </div>
  );
}

// Componente StatCard
interface StatCardProps {
  title: string;
  value: number;
  description: string;
  icon: React.ElementType;
  trend?: string;
  color: string;
}

function StatCard({ title, value, description, icon: Icon, trend, color }: StatCardProps) {
  return (
    <Card>
      <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
        <CardTitle className="text-sm font-medium">{title}</CardTitle>
        <Icon className={`h-4 w-4 ${color}`} />
      </CardHeader>
      <CardContent>
        <div className="text-2xl font-bold">{value.toLocaleString()}</div>
        <p className="text-xs text-muted-foreground mt-1">{description}</p>
        {trend && (
          <p className="text-xs text-green-500 mt-2 font-medium">{trend}</p>
        )}
      </CardContent>
    </Card>
  );
}

// Componente MiniStatCard
interface MiniStatCardProps {
  title: string;
  value: number;
  icon: React.ElementType;
  color: string;
}

function MiniStatCard({ title, value, icon: Icon, color }: MiniStatCardProps) {
  return (
    <Card>
      <CardContent className="flex items-center p-6">
        <div className={`rounded-full p-3 ${color} mr-4`}>
          <Icon className="h-5 w-5" />
        </div>
        <div>
          <p className="text-sm font-medium text-muted-foreground">{title}</p>
          <p className="text-2xl font-bold">{value.toLocaleString()}</p>
        </div>
      </CardContent>
    </Card>
  );
}

// Componente DetailedStatCard
interface DetailedStatCardProps {
  title: string;
  stats: Array<{ label: string; value: number }>;
  icon: React.ElementType;
}

function DetailedStatCard({ title, stats, icon: Icon }: DetailedStatCardProps) {
  return (
    <Card>
      <CardHeader>
        <div className="flex items-center space-x-2">
          <Icon className="h-5 w-5 text-primary" />
          <CardTitle className="text-base">{title}</CardTitle>
        </div>
      </CardHeader>
      <CardContent>
        <div className="space-y-3">
          {stats.map((stat, index) => (
            <div key={index} className="flex items-center justify-between">
              <span className="text-sm text-muted-foreground">{stat.label}</span>
              <span className="text-sm font-semibold">{stat.value.toLocaleString()}</span>
            </div>
          ))}
        </div>
      </CardContent>
    </Card>
  );
}

function DashboardSkeleton() {
  return (
    <div className="space-y-6">
      <div className="space-y-2">
        <div className="h-8 w-64 bg-muted animate-pulse rounded" />
        <div className="h-4 w-96 bg-muted animate-pulse rounded" />
      </div>
      <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
        {[...Array(4)].map((_, i) => (
          <div key={i} className="h-32 bg-muted animate-pulse rounded-lg" />
        ))}
      </div>
      <div className="grid gap-4 md:grid-cols-2">
        {[...Array(2)].map((_, i) => (
          <div key={i} className="h-96 bg-muted animate-pulse rounded-lg" />
        ))}
      </div>
    </div>
  );
}