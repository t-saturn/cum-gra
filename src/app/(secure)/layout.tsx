import { ReactNode } from "react";
import { auth } from "@/lib/auth";
import { redirect } from "next/navigation";
import { getUserRoleForCurrentApp } from "@/lib/roles";
import { RoleProvider } from "@/providers/role";
import { toBase64Url } from "@/helpers";
import { ProfileProvider } from "@/context/ProfileContext";
import { SidebarProvider } from "@/components/ui/sidebar";
import LayoutClient from '@/components/layout/layout';

const AUTH_ORIGIN = process.env.NEXT_PUBLIC_AUTH_ORIGIN!;
const APP_ORIGIN = process.env.NEXT_PUBLIC_APP_ORIGIN!;

export default async function ProtectedLayout({ children }: { children: ReactNode }) {
  const session = await auth();

  if (!session?.user?.id) {
    const callback = new URL("/dashboard", APP_ORIGIN).toString();
    const cb64 = toBase64Url(callback);
    console.log(session)
    const signin = new URL("/auth/signin", AUTH_ORIGIN);
    signin.searchParams.set("cb64", cb64);
    redirect(signin.toString());
  }

  const role = await getUserRoleForCurrentApp(String(session.user.id));
  if (!role) redirect("/unauthorized");

  return (<RoleProvider value={role}> <ProfileProvider>
    <SidebarProvider>
      <LayoutClient>{children}</LayoutClient>
    </SidebarProvider>
  </ProfileProvider></RoleProvider>)
}
