"use client";

import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Progress } from "@/components/ui/progress";
import {
  Users,
  Shield,
  Building2,
  Clock,
  ArrowUpRight,
  AlertTriangle,
  CheckCircle,
  Eye,
  MoreHorizontal,
} from "lucide-react";
import Link from "next/link";
import CardStatsContain from "@/components/custom/card/card-stats-contain";
import { statsDashboard } from "@/mocks/stats-mocks";

const recentActivities = [
  {
    id: 1,
    type: "user_created",
    message: "Nuevo usuario registrado",
    user: "juan.perez@empresa.com",
    time: "Hace 5 minutos",
    status: "success",
    icon: CheckCircle,
  },
  {
    id: 2,
    type: "role_assigned",
    message: "Rol 'Administrador' asignado",
    user: "María García",
    time: "Hace 15 minutos",
    status: "info",
    icon: Shield,
  },
  {
    id: 3,
    type: "login_failed",
    message: "Intento de acceso fallido",
    user: "IP 192.168.1.100",
    time: "Hace 30 minutos",
    status: "warning",
    icon: AlertTriangle,
  },
  {
    id: 4,
    type: "app_registered",
    message: "Nueva aplicación registrada",
    user: "Sistema de Inventario",
    time: "Hace 1 hora",
    status: "success",
    icon: Building2,
  },
];

const pendingTasks = [
  {
    id: 1,
    title: "Revisar solicitudes de acceso pendientes",
    count: 12,
    priority: "high",
    progress: 25,
  },
  {
    id: 2,
    title: "Actualizar permisos de módulos",
    count: 5,
    priority: "medium",
    progress: 60,
  },
  {
    id: 3,
    title: "Verificar usuarios inactivos",
    count: 23,
    priority: "low",
    progress: 80,
  },
];

const quickActions = [
  {
    title: "Crear Usuario",
    icon: Users,
    href: "/dashboard/users/new",
    color: "from-chart-2 to-chart-3",
  },
  {
    title: "Nueva Aplicación",
    icon: Building2,
    href: "/dashboard/applications/new",
    color: "from-chart-4 to-chart-5",
  },
  {
    title: "Configurar Rol",
    icon: Shield,
    href: "/dashboard/roles/application-roles/new",
    color: "from-primary to-chart-1",
  },
];

export default function DashboardOverview() {
  return (
    <div className="space-y-8">
      <div className="flex items-center justify-between flex-col sm:flex-row gap-4">
        <div>
          <h1 className="text-3xl font-bold bg-gradient-to-r from-foreground to-muted-foreground bg-clip-text text-transparent">
            Dashboard
          </h1>
          <p className="text-muted-foreground mt-2 text-lg">
            Resumen general del sistema de gestión de usuarios
          </p>
        </div>
        <div className="flex gap-3">
          <Button variant="outline" className="hover:bg-accent">
            <Eye className="w-4 h-4 mr-2" />
            Ver Reportes
          </Button>
        </div>
      </div>

      {/* Stats Cards */}
      <CardStatsContain stats={statsDashboard} />

      <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
        {/* Recent Activities */}
        <Card className="border-border bg-card/50 backdrop-blur-sm">
          <CardHeader className="flex flex-row items-center justify-between">
            <div>
              <CardTitle className="text-foreground">
                Actividad Reciente
              </CardTitle>
              <CardDescription>
                Últimas acciones realizadas en el sistema
              </CardDescription>
            </div>
            <Button variant="ghost" size="sm">
              <MoreHorizontal className="h-4 w-4" />
            </Button>
          </CardHeader>
          <CardContent className="space-y-4">
            {recentActivities.map((activity) => (
              <div
                key={activity.id}
                className="flex items-start gap-4 p-3 rounded-lg hover:bg-accent/50 transition-colors"
              >
                <div
                  className={`p-2 rounded-full ${
                    activity.status === "success"
                      ? "bg-chart-4/20 text-chart-4"
                      : activity.status === "warning"
                      ? "bg-chart-5/20 text-chart-5"
                      : "bg-chart-2/20 text-chart-2"
                  }`}
                >
                  <activity.icon className="w-4 h-4" />
                </div>
                <div className="flex-1 min-w-0">
                  <p className="text-sm font-medium text-foreground">
                    {activity.message}
                  </p>
                  <p className="text-sm text-muted-foreground">
                    {activity.user}
                  </p>
                  <p className="text-xs text-muted-foreground/80 mt-1">
                    {activity.time}
                  </p>
                </div>
                <Button variant="ghost" size="sm">
                  <ArrowUpRight className="h-4 w-4" />
                </Button>
              </div>
            ))}
          </CardContent>
        </Card>

        {/* Pending Tasks */}
        <Card className="border-border bg-card/50 backdrop-blur-sm">
          <CardHeader className="flex flex-row items-center justify-between">
            <div>
              <CardTitle className="text-foreground">
                Tareas Pendientes
              </CardTitle>
              <CardDescription>
                Acciones que requieren tu atención
              </CardDescription>
            </div>
            <Button variant="ghost" size="sm">
              <MoreHorizontal className="h-4 w-4" />
            </Button>
          </CardHeader>
          <CardContent className="space-y-4">
            {pendingTasks.map((task) => (
              <div key={task.id} className="p-4 bg-accent/30 rounded-lg">
                <div className="flex items-center justify-between mb-3">
                  <div className="flex items-center gap-3">
                    <Clock className="w-4 h-4 text-muted-foreground" />
                    <div>
                      <p className="text-sm font-medium text-foreground">
                        {task.title}
                      </p>
                      <p className="text-xs text-muted-foreground">
                        {task.count} elementos
                      </p>
                    </div>
                  </div>
                  <Badge
                    variant={
                      task.priority === "high"
                        ? "destructive"
                        : task.priority === "medium"
                        ? "default"
                        : "secondary"
                    }
                  >
                    {task.priority === "high"
                      ? "Alta"
                      : task.priority === "medium"
                      ? "Media"
                      : "Baja"}
                  </Badge>
                </div>
                <div className="space-y-2">
                  <div className="flex justify-between text-xs text-muted-foreground">
                    <span>Progreso</span>
                    <span>{task.progress}%</span>
                  </div>
                  <Progress value={task.progress} className="h-2" />
                </div>
              </div>
            ))}
          </CardContent>
        </Card>
      </div>

      {/* Quick Actions */}
      <Card className="border-border bg-card/50 backdrop-blur-sm">
        <CardHeader>
          <CardTitle className="text-foreground">Acciones Rápidas</CardTitle>
          <CardDescription>
            Accesos directos a las funciones más utilizadas
          </CardDescription>
        </CardHeader>
        <CardContent>
          <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
            {quickActions.map((action) => (
              <Link key={action.title} href={action.href}>
                <Card className="border-border hover:shadow-lg hover:shadow-primary/5 transition-all duration-300 cursor-pointer group">
                  <CardContent className="p-6 text-center">
                    <div
                      className={`w-16 h-16 mx-auto mb-4 rounded-2xl bg-gradient-to-r ${action.color} flex items-center justify-center group-hover:scale-110 transition-transform duration-300 shadow-lg`}
                    >
                      <action.icon className="w-8 h-8 text-white" />
                    </div>
                    <h3 className="font-semibold text-foreground group-hover:text-primary transition-colors">
                      {action.title}
                    </h3>
                  </CardContent>
                </Card>
              </Link>
            ))}
          </div>
        </CardContent>
      </Card>
    </div>
  );
}
