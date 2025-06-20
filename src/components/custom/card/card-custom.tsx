import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { ReactNode } from "react";

type Order = "auto" | "col" | "row";

interface CardCustomProps {
  title: string;
  description: string;
  order?: Order;
  children: ReactNode;
}

const CardCustom = ({ title, description, order = "auto", children }: CardCustomProps) => {

  return (
    <Card className="border-border bg-card/50 backdrop-blur-sm">
        <CardHeader>
          <CardTitle className="text-foreground">{title}</CardTitle>
          <CardDescription>
            {description}
          </CardDescription>
        </CardHeader>
        <CardContent>
          <div className={`grid gap-6 ${order === "col" ? 'grid-cols-1' : 'grid-cols-1 md:grid-cols-3'}`}>
            {children}
          </div>
        </CardContent>
      </Card>
  )
}

export default CardCustom