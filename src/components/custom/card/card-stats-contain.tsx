import { CardStatType } from '@/types/card-types/card-stats-types';
import CardStat from './card-stats';

interface CardStatsContainProps {
  stats: CardStatType[];
}

const CardStatsContain = ({ stats }: CardStatsContainProps) => {
  return (
    <div className="gap-3 grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4">
      {stats.map((stat, i) => (
        <CardStat key={i} stat={stat} />
      ))}
    </div>
  );
};

export default CardStatsContain;
