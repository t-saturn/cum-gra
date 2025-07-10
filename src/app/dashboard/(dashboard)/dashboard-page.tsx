"use client";

import CardStatsContain from "@/components/custom/card/card-stats-contain";
import { statsDashboard } from "@/mocks/stats-mocks";
import PendingTasks from "./pending-tasks";
import RecentActivities from "./recent-activities";
import QuickActions from "./quick-actions";
import InfoHeader from "@/components/custom/info-header";
import { Button } from "@/components/ui/button";
import { Eye } from "lucide-react";
import Link from "next/link";

export default function DashboardOverview() {
  return (
    <div className="space-y-8">
      <InfoHeader
        title="Dashboard"
        description="Resumen general del sistema de gestión de usuarios"
      >
        <Link href={'/dashboard/active-users'}>
          <Button variant="outline" className="hover:bg-accent">
            <Eye className="w-4 h-4 mr-2" />
            Ver Reportes
          </Button>
        </Link>
      </InfoHeader>

      {/* Stats Cards */}
      <CardStatsContain stats={statsDashboard} />

      <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
        {/* Recent Activities */}
        <RecentActivities />

        {/* Pending Tasks */}
        <PendingTasks />
      </div>

      {/* Quick Actions */}
      <QuickActions />
    </div>
  );
}
