import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { CardStatType } from "@/types/card-types/card-stats-types";

const CardStat = ({ stat }: { stat: CardStatType }) => {
  return (
    <Card
      key={stat.title}
      className="border-border bg-card/50 backdrop-blur-sm hover:shadow-lg hover:shadow-primary/5 transition-all duration-300"
    >
      <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
        <CardTitle className="text-sm font-medium text-muted-foreground">
          {stat.title}
        </CardTitle>
        <div className={`p-2 rounded-lg ${stat.bgColor}`}>
          <stat.icon className={`h-4 w-4 ${stat.textColor}`} />
        </div>
      </CardHeader>
      <CardContent>
        <div className="text-3xl font-bold text-foreground mb-2">
          {stat.value}
        </div>
        {stat.trend && stat.trendIcon && (
          <div className="flex items-center text-sm">
            <stat.trendIcon className={`h-4 w-4 text-${stat.color} mr-1`} />
            <span
              className={
                `text-${stat.color}`
              }
            >
              {stat.change}
            </span>
            <span className="text-muted-foreground ml-1">
              {stat.trendText}
            </span>
          </div>
        )}
      </CardContent>
    </Card>
  );
};

export default CardStat;
