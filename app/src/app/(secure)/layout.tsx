import { ReactNode } from 'react';
import { ProfileProvider } from '@/context/profile';
import { SidebarProvider } from '@/components/ui/sidebar';
import LayoutClient from '@/components/layout/layout';
import SessionGuard from '@/components/auth/session-guard';
import RoleGuard from '@/components/auth/role-guard';

export const dynamic = 'force-dynamic';

export default async function ProtectedLayout({ children }: { children: ReactNode }) {
  
  return (
    <SessionGuard>
      <RoleGuard>
        <ProfileProvider>
          <SidebarProvider>
            <LayoutClient>{children}</LayoutClient>
          </SidebarProvider>
        </ProfileProvider>
      </RoleGuard>
    </SessionGuard>
  );
}
