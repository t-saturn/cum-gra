'use client';

import { ChevronDown, LogOut } from 'lucide-react';
import Image from 'next/image';
import Link from 'next/link';
import { usePathname } from 'next/navigation';
import { useState, useRef } from 'react';

import { Collapsible, CollapsibleContent, CollapsibleTrigger } from '@/components/ui/collapsible';
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarGroup,
  SidebarGroupLabel,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarMenuSub,
  SidebarRail,
  useSidebar,
} from '@/components/ui/sidebar';
import { fn_get_sidebar_menu } from '@/helpers/sidebar-helper';
import { useProfile } from '@/context/ProfileContext';
import { SidebarItem, SidebarSubItem } from '@/types/sidebar-types';

export default function AppSidebar({ hoveredItem, setHoveredItem }: { hoveredItem: string | null; setHoveredItem: (item: string | null) => void }) {
  const { profile } = useProfile();
  const { state, isMobile, setOpenMobile } = useSidebar();
  const isCollapsed = state === 'collapsed';
  const pathname = usePathname();
  const [hoverPosition, setHoverPosition] = useState<number>(0);
  const menuRefs = useRef<Map<string, HTMLDivElement>>(new Map());

  const sidebarMenus = fn_get_sidebar_menu(profile.role);

  const handleCloseSession = () => {
    // simulated logout - redirect login page
    window.location.href = '/login';
  };

  const handleMouseEnter = (item: string, event: React.MouseEvent) => {
    if (isCollapsed && sidebarMenus.some((group) => group.menu.some((menuItem: SidebarItem) => menuItem.label === item && menuItem.items))) {
      setHoveredItem(item);
      const target = event.currentTarget as HTMLDivElement;
      const rect = target.getBoundingClientRect();
      setHoverPosition(rect.top);
    }
  };

  const handleMouseLeave = () => setHoveredItem(null);

  return (
    <>
      <Sidebar className="rounded-lg border border-border shadow-sm relative h-full overflow-hidden" collapsible="icon">
        <SidebarHeader className="bg-card flex items-center border-b rounded-t-lg">
          {!isCollapsed ? (
            <div className="flex items-center gap-4 font-semibold p-6">
              <Image src="/img/logo.png" alt="logo" width={40} height={20} />
              <div className="flex flex-col">
                <span className="text-xl font-bold">CUM</span>
                <span className="text-sm font-light text-muted-foreground">Central User Manager</span>
              </div>
            </div>
          ) : (
            <div className="p-2">
              <Image src="/img/logo.png" alt="logo" width={20} height={20} />
            </div>
          )}
        </SidebarHeader>
        <SidebarContent className="bg-card rounded-b-lg">
          {sidebarMenus.map((menubar, index) => (
            <SidebarGroup key={index}>
              {!isCollapsed && <SidebarGroupLabel className="px-6 text-xs font-semibold uppercase text-muted-foreground">{menubar.title}</SidebarGroupLabel>}
              <SidebarMenu>
                {menubar.menu.map((item: SidebarItem, i: number) => {
                  const isActive = item.url === pathname || (item.items && item.items.some((subitem: SidebarSubItem) => subitem.url === pathname));

                  const isSubItemActive = item.items && item.items.some((subitem: SidebarSubItem) => subitem.url === pathname);

                  return (
                    <SidebarMenuItem key={i}>
                      {item.items ? (
                        <Collapsible asChild className="group/collapsible hover:cursor-pointer">
                          <div
                            ref={(el) => {
                              if (el) menuRefs.current.set(item.label, el);
                            }}
                          >
                            <CollapsibleTrigger asChild>
                              <SidebarMenuButton
                                tooltip={item.items ? undefined : item.label}
                                className={`hover:bg-primary hover:text-[#eff1f5] ${
                                  isActive
                                    ? isSubItemActive && !isCollapsed
                                      ? 'bg-secondary group-data-[state=closed]/collapsible:bg-primary'
                                      : 'bg-primary text-[#eff1f5]'
                                    : 'data-[active=true]:bg-primary data-[active=true]:text-[#eff1f5]'
                                } hover:cursor-pointer`}
                                onMouseEnter={(e) => handleMouseEnter(item.label, e)}
                                onMouseLeave={handleMouseLeave}
                              >
                                {item.icon && <item.icon width={16} />}
                                {!isCollapsed && <span>{item.label}</span>}
                                <ChevronDown className="ml-auto transition-transform duration-200 group-data-[state=open]/collapsible:rotate-180" />
                              </SidebarMenuButton>
                            </CollapsibleTrigger>
                            <CollapsibleContent>
                              <SidebarMenuSub>
                                {item.items.map((subitem: SidebarSubItem, j: number) => {
                                  const isSubItemActive = subitem.url === pathname;
                                  return (
                                    <Link
                                      href={subitem.url}
                                      key={j}
                                      className={`px-2 py-1 text-xs flex items-center gap-2 rounded-lg ${
                                        isSubItemActive ? 'bg-primary text-[#eff1f5]' : 'hover:bg-primary hover:text-[#eff1f5]'
                                      }`}
                                      onClick={() => {
                                        if (isMobile) {
                                          setOpenMobile(false);
                                        }
                                      }}
                                    >
                                      {subitem.icon && <subitem.icon width={16} />}
                                      <span>{subitem.label}</span>
                                    </Link>
                                  );
                                })}
                              </SidebarMenuSub>
                            </CollapsibleContent>
                          </div>
                        </Collapsible>
                      ) : (
                        <Link href={item.url}>
                          <SidebarMenuButton
                            className={`hover:bg-primary hover:text-[#eff1f5] ${
                              isActive ? 'bg-primary text-[#eff1f5]' : 'data-[active=true]:bg-primary data-[active=true]:text-[#eff1f5]'
                            } hover:cursor-pointer`}
                            onMouseEnter={(e) => handleMouseEnter(item.label, e)}
                            onMouseLeave={handleMouseLeave}
                            onClick={() => {
                              if (isMobile) {
                                setOpenMobile(false);
                              }
                            }}
                            tooltip={item.items ? undefined : item.label}
                          >
                            {item.icon && <item.icon width={16} />}
                            {!isCollapsed && <span>{item.label}</span>}
                          </SidebarMenuButton>
                        </Link>
                      )}
                    </SidebarMenuItem>
                  );
                })}
              </SidebarMenu>
            </SidebarGroup>
          ))}
        </SidebarContent>
        <SidebarFooter className="bg-card">
          <SidebarMenuButton
            tooltip="Cerrar sesión"
            className="w-full justify-center bg-destructive/5 hover:bg-destructive/10 text-destructive hover:text-destructive rounded-lg h-11 font-medium hover:cursor-pointer"
            onClick={handleCloseSession}
          >
            <LogOut width={16} className={`transition-all duration-150 ${isCollapsed ? 'rotate-180' : ''}`} />
            {!isCollapsed && <span>Cerrar sesión</span>}
          </SidebarMenuButton>
        </SidebarFooter>
        <SidebarRail />
      </Sidebar>

      {hoveredItem && isCollapsed && sidebarMenus.some((group) => group.menu.some((item: SidebarItem) => item.label === hoveredItem && item.items)) && (
        <div
          className="absolute left-11 z-50 w-48 rounded-md border bg-card py-1 shadow-lg"
          style={{ top: `${hoverPosition}px` }}
          onMouseEnter={() => {
            const item = sidebarMenus.flatMap((group) => group.menu).find((i) => i.label === hoveredItem);
            if (item?.items) setHoveredItem(hoveredItem);
          }}
          onMouseLeave={() => setHoveredItem(null)}
        >
          <div className="px-3 py-2 text-sm font-medium text-muted-foreground">{hoveredItem}</div>
          <div className="border-t"></div>
          {sidebarMenus
            .flatMap((group) => group.menu)
            .find((item) => item.label === hoveredItem)
            ?.items?.map((subitem: SidebarSubItem, index: number) => {
              const isSubItemActive = subitem.url === pathname;
              return (
                <Link
                  href={subitem.url}
                  key={index}
                  className={`flex items-center gap-2 m-2 px-3 py-2 text-sm rounded-lg ${isSubItemActive ? 'bg-primary text-white' : 'hover:bg-primary hover:text-white'}`}
                >
                  {subitem.icon && <subitem.icon className="h-4 w-4" />}
                  <span>{subitem.label}</span>
                </Link>
              );
            })}
        </div>
      )}
    </>
  );
}
