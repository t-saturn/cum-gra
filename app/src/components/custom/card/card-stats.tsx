import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { CardStatType } from '@/types/card-types/card-stats-types';

const CardStat = ({ stat }: { stat: CardStatType }) => {
  return (
    <Card key={stat.title} className="bg-card/50 hover:shadow-lg hover:shadow-primary/5 backdrop-blur-sm border-border transition-all duration-300">
      <CardHeader className="flex flex-row justify-between items-center space-y-0 pb-2">
        <CardTitle className="font-medium text-muted-foreground text-sm">{stat.title}</CardTitle>
        <div className={`p-2 rounded-lg ${stat.bgColor}`}>
          <stat.icon className={`h-4 w-4 ${stat.textColor}`} />
        </div>
      </CardHeader>
      <CardContent>
        <div className="mb-2 font-bold text-foreground text-3xl">{stat.value}</div>
        {stat.trend && stat.trendIcon && (
          <div className="flex items-center text-sm">
            <stat.trendIcon className={`h-4 w-4 text-${stat.color} mr-1`} />
            <span className={`text-${stat.color}`}>{stat.change}</span>
            <span className="ml-1 text-muted-foreground">{stat.trendText}</span>
          </div>
        )}
      </CardContent>
    </Card>
  );
};

export default CardStat;
