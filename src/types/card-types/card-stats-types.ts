import { LucideIcon } from "lucide-react";

export interface CardStatType {
  title: string;
  value: string;
  change?: string;
  trend?: string;
  trendIcon?: LucideIcon;
  trendText?: string;
  icon: LucideIcon;
  color: string;
  bgColor: string;
  textColor: string;
}
