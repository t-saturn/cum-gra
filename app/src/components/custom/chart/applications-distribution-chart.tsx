'use client';

import { Bar, BarChart, CartesianGrid, XAxis, YAxis } from 'recharts';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import {
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
  type ChartConfig,
} from '@/components/ui/chart';

const chartConfig = {
  applications: {
    label: 'Aplicaciones',
    color: 'var(--chart-1)',
  },
  roles: {
    label: 'Roles',
    color: 'var(--chart-2)',
  },
} satisfies ChartConfig;

interface ApplicationsDistributionChartProps {
  appsStats: any;
  rolesStats: any;
}

export function ApplicationsDistributionChart({
  appsStats,
  rolesStats,
}: ApplicationsDistributionChartProps) {
  const chartData = [
    {
      category: 'Total',
      applications: appsStats?.total_applications || 0,
      roles: rolesStats?.total_roles || 0,
    },
    {
      category: 'Activas',
      applications: appsStats?.active_applications || 0,
      roles: rolesStats?.active_roles || 0,
    },
    {
      category: 'Con Usuarios',
      applications: appsStats?.apps_with_users || 0,
      roles: rolesStats?.roles_with_users || 0,
    },
  ];

  return (
    <Card>
      <CardHeader>
        <CardTitle>Distribución de Aplicaciones y Roles</CardTitle>
        <CardDescription>
          Comparación entre aplicaciones y roles del sistema
        </CardDescription>
      </CardHeader>
      <CardContent>
        <ChartContainer config={chartConfig} className="h-[300px] w-full">
          <BarChart data={chartData}>
            <CartesianGrid vertical={false} strokeDasharray="3 3" />
            <XAxis
              dataKey="category"
              tickLine={false}
              tickMargin={10}
              axisLine={false}
            />
            <YAxis
              tickLine={false}
              axisLine={false}
              tickMargin={10}
            />
            <ChartTooltip
              cursor={false}
              content={<ChartTooltipContent indicator="dashed" />}
            />
            <Bar
              dataKey="applications"
              fill="var(--color-applications)"
              radius={[8, 8, 0, 0]}
            />
            <Bar
              dataKey="roles"
              fill="var(--color-roles)"
              radius={[8, 8, 0, 0]}
            />
          </BarChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
}