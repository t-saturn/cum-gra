'use client';

import { Popover, PopoverContent, PopoverTrigger } from '@/components/ui/popover';
import { Button } from '@/components/ui/button';
import { Badge } from '@/components/ui/badge';
import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar';
import { Settings, LogOut, User, Shield, Clock, MapPin, ChevronRight } from 'lucide-react';
import Link from 'next/link';
import { useProfile } from '@/context/ProfileContext';
import { to_cb64 } from '@/helpers';

const AUTH_ORIGIN = process.env.NEXT_PUBLIC_AUTH_ORIGIN!; // p.ej. http://sso.localtest.me:30000
const APP_ORIGIN = process.env.NEXT_PUBLIC_APP_ORIGIN!; // p.ej. http://cum.localtest.me:30001

export function UserPopover() {
  const { profile } = useProfile();

  const signoutUrl = new URL('/api/auth/signout', AUTH_ORIGIN);
  signoutUrl.searchParams.set('cb64', to_cb64(`${APP_ORIGIN}/`));

  const getInitials = (name: string) => {
    return name
      .split(' ')
      .map((word) => word.charAt(0))
      .join('')
      .toUpperCase()
      .slice(0, 2);
  };

  const userStats = {
    lastLogin: 'Hace 2 horas',
    location: 'Ayacucho, Perú',
    sessionsActive: 2,
    role: profile.role || 'Administrador',
  };

  return (
    <Popover>
      <PopoverTrigger asChild>
        <Button variant="ghost" className="flex items-center gap-3 hover:bg-accent/50 px-3 py-2 rounded-xl h-auto hover:scale-[1.02] transition-all duration-200">
          <Avatar className="ring-2 ring-primary/20 ring-offset-2 ring-offset-background w-10 h-10">
            <AvatarImage src={profile.avatar || '/placeholder.svg?height=40&width=40'} alt={profile.name} />
            <AvatarFallback className="bg-gradient-to-br from-primary to-chart-1 font-semibold text-primary-foreground">{getInitials(profile.name)}</AvatarFallback>
          </Avatar>
          <div className="hidden md:block text-left">
            <div className="max-w-32 font-medium text-foreground text-sm truncate">{profile.name}</div>
            <div className="text-muted-foreground text-xs">{userStats.role}</div>
          </div>
        </Button>
      </PopoverTrigger>
      <PopoverContent className="bg-popover shadow-2xl shadow-black/10 backdrop-blur-xl mr-2 border-border/50 w-80" align="end">
        <div className="space-y-6">
          <div className="flex flex-col items-center space-y-4 text-center">
            <div className="relative">
              <Avatar className="ring-4 ring-primary/20 ring-offset-4 ring-offset-background w-20 h-20">
                <AvatarImage src={profile.avatar || '/placeholder.svg?height=80&width=80'} alt={profile.name} />
                <AvatarFallback className="bg-gradient-to-br from-primary to-chart-1 font-bold text-primary-foreground text-xl">{getInitials(profile.name)}</AvatarFallback>
              </Avatar>
              <div className="-right-1 -bottom-1 absolute flex justify-center items-center bg-green-500 border-2 border-background rounded-full w-6 h-6">
                <div className="bg-white rounded-full w-2 h-2"></div>
              </div>
            </div>

            <div className="space-y-2">
              <h3 className="font-semibold text-foreground text-lg">{profile.name}</h3>
              <Badge variant="secondary" className="bg-primary/10 border-primary/20 font-medium text-primary">
                <Shield className="mr-1 w-3 h-3" />
                {userStats.role}
              </Badge>
            </div>
          </div>

          <div className="gap-3 grid grid-cols-2">
            <div className="bg-muted/30 p-3 rounded-lg text-center">
              <div className="flex justify-center items-center gap-1 mb-1 text-muted-foreground">
                <Clock className="w-3 h-3" />
                <span className="text-xs">Último acceso</span>
              </div>
              <div className="font-medium text-sm">{userStats.lastLogin}</div>
            </div>
            <div className="bg-muted/30 p-3 rounded-lg text-center">
              <div className="flex justify-center items-center gap-1 mb-1 text-muted-foreground">
                <MapPin className="w-3 h-3" />
                <span className="text-xs">Ubicación</span>
              </div>
              <div className="font-medium text-sm">{userStats.location}</div>
            </div>
          </div>

          <div className="space-y-2">
            <Link href="https://accounts.regionayacucho.gob.pe/user" target="_blank">
              <Button variant="ghost" className="justify-between hover:bg-accent/50 rounded-lg w-full h-12">
                <div className="flex items-center gap-3">
                  <div className="flex justify-center items-center bg-primary/10 rounded-lg w-8 h-8">
                    <User className="w-4 h-4 text-primary" />
                  </div>
                  <div className="text-left">
                    <div className="font-medium text-sm">Mi Perfil</div>
                    <div className="text-muted-foreground text-xs">Gestionar información personal</div>
                  </div>
                </div>
                <ChevronRight className="w-4 h-4 text-muted-foreground" />
              </Button>
            </Link>

            <Link href="/dashboard/settings">
              <Button variant="ghost" className="justify-between hover:bg-accent/50 rounded-lg w-full h-12">
                <div className="flex items-center gap-3">
                  <div className="flex justify-center items-center bg-orange-500/10 rounded-lg w-8 h-8">
                    <Settings className="w-4 h-4 text-orange-600" />
                  </div>
                  <div className="text-left">
                    <div className="font-medium text-sm">Configuración</div>
                    <div className="text-muted-foreground text-xs">Preferencias del sistema</div>
                  </div>
                </div>
                <ChevronRight className="w-4 h-4 text-muted-foreground" />
              </Button>
            </Link>
          </div>

          <div className="mt-2 pt-2 border-t border-border/50">
            <a
              href={signoutUrl.toString()}
              className="flex items-center gap-2 hover:bg-red-600 px-3 py-2 rounded-lg w-full font-medium text-red-600 hover:text-white text-sm transition-colors duration-200"
            >
              <LogOut className="w-4 h-4" />
              Cerrar sesión
            </a>
          </div>

          <div className="pt-2 border-t border-border/50">
            <div className="flex justify-between items-center text-muted-foreground text-xs">
              <span>Sesiones activas: {userStats.sessionsActive}</span>
              <Badge variant="outline" className="text-xs">
                Conectado
              </Badge>
            </div>
          </div>
        </div>
      </PopoverContent>
    </Popover>
  );
}
