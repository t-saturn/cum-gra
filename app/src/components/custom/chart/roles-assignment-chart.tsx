'use client';

import { Pie, PieChart, Cell, Legend } from 'recharts';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import {
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
  type ChartConfig,
} from '@/components/ui/chart';

const chartConfig = {
  active: {
    label: 'Activas',
    color: 'var(--chart-1)',
  },
  revoked: {
    label: 'Revocadas',
    color: 'var(--chart-2)',
  },
  deleted: {
    label: 'Eliminadas',
    color: 'var(--chart-3)',
  },
} satisfies ChartConfig;

interface RolesAssignmentChartProps {
  assignmentsStats: any;
}

export function RolesAssignmentChart({ assignmentsStats }: RolesAssignmentChartProps) {
  const chartData = [
    {
      name: 'Activas',
      value: assignmentsStats?.active_assignments || 0,
      fill: 'var(--color-active)',
    },
    {
      name: 'Revocadas',
      value: assignmentsStats?.revoked_assignments || 0,
      fill: 'var(--color-revoked)',
    },
    {
      name: 'Eliminadas',
      value: assignmentsStats?.deleted_assignments || 0,
      fill: 'var(--color-deleted)',
    },
  ];

  const total = chartData.reduce((sum, item) => sum + item.value, 0);

  return (
    <Card>
      <CardHeader>
        <CardTitle>Estado de Asignaciones de Roles</CardTitle>
        <CardDescription>
          Distribución de {total.toLocaleString()} asignaciones totales
        </CardDescription>
      </CardHeader>
      <CardContent>
        <ChartContainer config={chartConfig} className="h-[300px] w-full">
          <PieChart>
            <ChartTooltip
              content={
                <ChartTooltipContent
                  hideLabel
                  formatter={(value, name) => (
                    <div className="flex items-center gap-2">
                      <span className="font-medium">{name}:</span>
                      <span className="font-bold">{value}</span>
                      <span className="text-muted-foreground">
                        ({((Number(value) / total) * 100).toFixed(1)}%)
                      </span>
                    </div>
                  )}
                />
              }
            />
            <Pie
              data={chartData}
              dataKey="value"
              nameKey="name"
              cx="50%"
              cy="50%"
              outerRadius={100}
              label={({ name, percent }) =>
                `${name}: ${(percent * 100).toFixed(0)}%`
              }
            >
              {chartData.map((entry, index) => (
                <Cell key={`cell-${index}`} fill={entry.fill} />
              ))}
            </Pie>
            <Legend />
          </PieChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
}