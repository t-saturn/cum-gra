import CardCustom from "@/components/custom/card/card-custom";
import { AlertTriangle, Building2, CheckCircle, Shield } from "lucide-react";

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

const RecentActivities = () => {
  return (
    <CardCustom
          title="Actividad Reciente"
          description="Últimas acciones realizadas en el sistema"
          order="col"
        >
          {recentActivities.map((activity) => (
            <div
              key={activity.id}
              className="flex items-start gap-4 p-3 rounded-lg bg-accent/50 transition-colors"
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
                <p className="text-sm text-muted-foreground">{activity.user}</p>
                <p className="text-xs text-muted-foreground/80 mt-1">
                  {activity.time}
                </p>
              </div>
            </div>
          ))}
        </CardCustom>
  )
}

export default RecentActivities