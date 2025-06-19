'use client';

import React from 'react';

import AppSidebar from './sidebar';

import { SidebarInset } from '@/components/ui/sidebar';
import Navbar from './navbar';

interface LayoutProps {
  children: React.ReactNode;
}

export default function Layout({ children }: LayoutProps) {
  const [hoveredItem, setHoveredItem] = React.useState<string | null>(null);

  return (
    <div className="flex h-screen bg-background p-2 w-full gap-2">
      <AppSidebar hoveredItem={hoveredItem} setHoveredItem={setHoveredItem} />
      <SidebarInset className="flex-1 rounded-lg border border-border bg-card shadow-sm overflow-y-auto">
        <Navbar />
        <div className="p-4">{children}</div>
      </SidebarInset>
    </div>
  );
}