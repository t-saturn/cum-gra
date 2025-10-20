import { SidebarMenuGroup } from '@/types/sidebar-types';
import { Boxes, Layers, LayoutDashboard, Package } from 'lucide-react';
import { Eye, UserX, ScrollText, ShieldCheck, UserCog } from 'lucide-react';
import { Briefcase, Building } from 'lucide-react';
import { Settings2, UserCircle, UserCheck, Lock, Ban } from 'lucide-react';
import { Users, BarChart3, Shield, Activity } from 'lucide-react';

export const baseMenus: SidebarMenuGroup[] = [
  {
    title: 'Menú',
    menu: [
      {
        label: 'Dashboard',
        icon: LayoutDashboard,
        url: '/dashboard',
        roles: ['user', 'admin'],
      },
    ],
  },
  {
    title: 'Gestión',
    menu: [
      {
        label: 'Usuarios',
        icon: Users,
        url: '/dashboard/users',
        roles: ['admin'],
        items: [
          { label: 'Usuarios', icon: UserCheck, url: '/dashboard/users' },
          {
            label: 'Posiciones Estructurales',
            icon: Briefcase,
            url: '/dashboard/structural-positions',
          },
          {
            label: 'Unidades Orgánicas',
            icon: Building,
            url: '/dashboard/organic-units',
          },
        ],
      },
      {
        label: 'Aplicaciones',
        icon: Layers,
        url: '/dashboard/applications',
        roles: ['admin'],
        items: [
          {
            label: 'Aplicaciones',
            icon: Boxes,
            url: '/dashboard/applications',
          },
          { label: 'Módulos', icon: Package, url: '/dashboard/modules' },
        ],
      },
    ],
  },
  {
    title: 'Seguridad',
    menu: [
      {
        label: 'Seguridad',
        icon: Shield,
        url: '/dashboard/security',
        roles: ['admin'],
        items: [
          {
            label: 'Sesiones Activas',
            icon: Activity,
            url: '/dashboard/security/active-sessions',
          },
          {
            label: 'Logs de Auditoria',
            icon: ScrollText,
            url: '/dashboard/security/audit-logs',
          },
        ],
      },
      {
        label: 'Roles y Permisos',
        icon: ShieldCheck,
        url: '/dashboard/security',
        roles: ['admin'],
        items: [
          {
            label: 'Roles de Aplicación',
            icon: Shield,
            url: '/dashboard/security/application-roles',
          },
          {
            label: 'Asignación de Roles',
            icon: UserCog,
            url: '/dashboard/security/user-roles',
          },
          {
            label: 'Permisos de Módulos',
            icon: Lock,
            url: '/dashboard/security/module-permissions',
          },
          {
            label: 'Restricción de Usuario',
            icon: Ban,
            url: '/dashboard/security/user-restrictions',
          },
        ],
      },
    ],
  },
  {
    title: 'Configuración',
    menu: [
      {
        label: 'Ajustes',
        icon: Settings2,
        url: '/dashboard/settings',
        roles: ['admin'],
      },
      {
        label: 'Cuenta',
        icon: UserCircle,
        url: '/dashboard/account',
        roles: ['admin'],
      },
    ],
  },
];
