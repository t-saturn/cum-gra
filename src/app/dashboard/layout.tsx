import { redirect } from "next/navigation";
import LayoutClient from "@/components/layout/layout";
import { ProfileProvider } from "@/context/ProfileContext";
import { SidebarProvider } from "@/components/ui/sidebar";
export default async function DashboardLayout({
  children,
}: Readonly<{ children: React.ReactNode }>) {
  const session = {
    user: {
      name: "John Doe",
      email: "johndoe@example.com",
      image: "https://i.pravatar.cc/300",
    },
  };

  if (!session?.user) redirect("/");

  return (
      <ProfileProvider>
        <SidebarProvider>
          <LayoutClient>{children}</LayoutClient>
        </SidebarProvider>
      </ProfileProvider>
  );
}
