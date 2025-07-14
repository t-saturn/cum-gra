import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { LucideIcon } from 'lucide-react';
import { ReactNode } from 'react';

type Order = 'auto' | 'col' | 'row';

interface CardCustomProps {
  Icon?: LucideIcon;
  title: string;
  description: string;
  order?: Order;
  children: ReactNode;
}

const CardCustom = ({ title, description, order = 'auto', children, Icon }: CardCustomProps) => {
  return (
    <Card className="border-border bg-card/50 backdrop-blur-sm">
      <CardHeader>
        <CardTitle className="flex items-center gap-2">
          {Icon && <Icon className="w-5 h-5" />}
          {title}
        </CardTitle>
        <CardDescription>{description}</CardDescription>
      </CardHeader>
      <CardContent>
        <div className={`grid gap-6 ${order === 'col' ? 'grid-cols-1' : 'grid-cols-1 md:grid-cols-3'}`}>{children}</div>
      </CardContent>
    </Card>
  );
};

export default CardCustom;
