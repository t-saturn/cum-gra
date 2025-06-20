import CardCustom from "@/components/custom/card/card-custom";
import { Badge } from "@/components/ui/badge";
import { Progress } from "@/components/ui/progress";
import { Clock } from "lucide-react";

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

const PendingTasks = () => {
  return (
    <CardCustom
      title="Tareas Pendientes"
      description="Acciones que requieren tu atención"
      order="col"
    >
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
    </CardCustom>
  );
};

export default PendingTasks;
