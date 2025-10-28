import { SidebarMenuGroup } from '@/types/sidebar-types';
import {
  Boxes,
  CircleQuestionMark,
  Hexagon,
  Layers,
  LayoutDashboard,
  Package,
  UsersRound,
  Settings2,
  UserCircle,
  Ban,
  Shield,
  Activity,
  ScrollText,
  ShieldCheck,
  UserCog,
} from 'lucide-react';

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
        icon: UsersRound,
        url: '/dashboard/users',
        roles: ['admin'],
        items: [
          { label: 'Usuarios', icon: UsersRound, url: '/dashboard/users' },
          {
            label: 'Posiciones Estructurales',
            icon: CircleQuestionMark,
            url: '/dashboard/structural-positions',
          },
          {
            label: 'Unidades Orgánicas',
            icon: Hexagon,
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
