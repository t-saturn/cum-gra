import CardCustom from "@/components/custom/card/card-custom";
import { Card, CardContent } from "@/components/ui/card";
import { Building2, Shield, Users } from "lucide-react";
import Link from "next/link";

const quickActions = [
  {
    title: "Crear Usuario",
    icon: Users,
    href: "/dashboard/users",
    color: "from-chart-2 to-chart-3",
  },
  {
    title: "Posición Estructural",
    icon: Building2,
    href: "/dashboard/structural-positions",
    color: "from-chart-4 to-chart-5",
  },
  {
    title: "Configurar Rol",
    icon: Shield,
    href: "/dashboard/security/application-roles",
    color: "from-primary to-chart-1",
  },
];

const QuickActions = () => {
  return (
    <CardCustom
        title="Acciones Rápidas"
        description="Accesos directos a las funciones más utilizadas"
        order="auto"
      >
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
      </CardCustom>
  )
}

export default QuickActions