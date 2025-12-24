'use client';

import { useMemo } from 'react';
import { Bar, BarChart, CartesianGrid, XAxis, YAxis } from 'recharts';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import {
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
  type ChartConfig,
} from '@/components/ui/chart';
import { Activity, AlertCircle } from 'lucide-react';

const chartConfig = {
  count: {
    label: 'Sesiones',
    color: 'var(--chart-1)',
  },
} satisfies ChartConfig;

interface SessionsActivityChartProps {
  sessionsStats: any;
}

export function SessionsActivityChart({ sessionsStats }: SessionsActivityChartProps) {
  const chartData = useMemo(() => {
    // Si no hay datos de sesiones por cliente, mostrar mensaje
    if (!sessionsStats?.sessions_by_client || Object.keys(sessionsStats.sessions_by_client).length === 0) {
      return [];
    }

    // Convertir el objeto sessions_by_client a array y tomar los top 6
    const sortedClients = Object.entries(sessionsStats.sessions_by_client)
      .sort(([, a], [, b]) => (b as number) - (a as number))
      .slice(0, 6)
      .map(([clientName, count]) => ({
        client: clientName.length > 20 ? clientName.substring(0, 17) + '...' : clientName,
        fullName: clientName,
        count: count as number,
      }));

    return sortedClients;
  }, [sessionsStats]);

  return (
    <Card>
      <CardHeader>
        <div className="flex items-center justify-between">
          <div>
            <CardTitle>Sesiones Activas por Aplicación</CardTitle>
            <CardDescription>
              Distribución de sesiones actuales
            </CardDescription>
          </div>
          <div className="flex items-center gap-2 text-sm">
            <Activity className="h-4 w-4 text-green-500" />
            <span className="font-medium">
              {sessionsStats?.total_sessions || 0} activas
            </span>
          </div>
        </div>
      </CardHeader>
      <CardContent>
        {chartData.length === 0 ? (
          <div className="flex flex-col items-center justify-center h-[300px] text-center">
            <AlertCircle className="h-12 w-12 text-muted-foreground mb-4" />
            <p className="text-sm text-muted-foreground">
              No hay sesiones activas en este momento
            </p>
            <p className="text-xs text-muted-foreground mt-2">
              Las sesiones aparecerán cuando los usuarios inicien sesión
            </p>
          </div>
        ) : (
          <ChartContainer config={chartConfig} className="h-[300px] w-full">
            <BarChart data={chartData}>
              <CartesianGrid vertical={false} strokeDasharray="3 3" />
              <XAxis
                dataKey="client"
                tickLine={false}
                axisLine={false}
                tickMargin={8}
                angle={-45}
                textAnchor="end"
                height={80}
              />
              <YAxis
                tickLine={false}
                axisLine={false}
                tickMargin={8}
              />
              <ChartTooltip
                cursor={false}
                content={
                  <ChartTooltipContent
                    labelFormatter={(_, payload) => {
                      return payload?.[0]?.payload?.fullName || '';
                    }}
                    formatter={(value) => [`${value} sesiones`, 'Total']}
                  />
                }
              />
              <Bar
                dataKey="count"
                fill="var(--color-count)"
                radius={[8, 8, 0, 0]}
              />
            </BarChart>
          </ChartContainer>
        )}
      </CardContent>
    </Card>
  );
}