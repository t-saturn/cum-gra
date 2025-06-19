"use client"

import { Popover, PopoverContent, PopoverTrigger } from "@/components/ui/popover"
import { Button } from "@/components/ui/button"
import { Badge } from "@/components/ui/badge"
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar"
import { Settings, LogOut, User, Shield, Clock, MapPin, ChevronRight } from "lucide-react"
import Link from "next/link"
import { useProfile } from "@/context/ProfileContext"

export function UserPopover() {
  const { profile } = useProfile()

  // Función para obtener las iniciales del nombre
  const getInitials = (name: string) => {
    return name
      .split(" ")
      .map((word) => word.charAt(0))
      .join("")
      .toUpperCase()
      .slice(0, 2)
  }

  // Datos simulados adicionales para el perfil
  const userStats = {
    lastLogin: "Hace 2 horas",
    location: "Ayacucho, Perú",
    sessionsActive: 2,
    role: profile.role || "Administrador",
  }

  return (
    <Popover>
      <PopoverTrigger asChild>
        <Button
          variant="ghost"
          className="flex items-center gap-3 hover:bg-accent/50 px-3 py-2 h-auto rounded-xl transition-all duration-200 hover:scale-[1.02]"
        >
          <Avatar className="h-10 w-10 ring-2 ring-primary/20 ring-offset-2 ring-offset-background">
            <AvatarImage src={profile.avatar || "/placeholder.svg?height=40&width=40"} alt={profile.name} />
            <AvatarFallback className="bg-gradient-to-br from-primary to-chart-1 text-primary-foreground font-semibold">
              {getInitials(profile.name)}
            </AvatarFallback>
          </Avatar>
          <div className="hidden text-left md:block">
            <div className="font-medium text-sm text-foreground truncate max-w-32">{profile.name}</div>
            <div className="text-xs text-muted-foreground">{userStats.role}</div>
          </div>
        </Button>
      </PopoverTrigger>
      <PopoverContent
        className="w-80 bg-popover backdrop-blur-xl border-border/50 shadow-2xl shadow-black/10 mr-2"
        align="end"
      >
        <div className="space-y-6">
          {/* User Profile Header */}
          <div className="flex flex-col items-center text-center space-y-4">
            <div className="relative">
              <Avatar className="h-20 w-20 ring-4 ring-primary/20 ring-offset-4 ring-offset-background">
                <AvatarImage src={profile.avatar || "/placeholder.svg?height=80&width=80"} alt={profile.name} />
                <AvatarFallback className="bg-gradient-to-br from-primary to-chart-1 text-primary-foreground font-bold text-xl">
                  {getInitials(profile.name)}
                </AvatarFallback>
              </Avatar>
              <div className="absolute -bottom-1 -right-1 w-6 h-6 bg-green-500 rounded-full border-2 border-background flex items-center justify-center">
                <div className="w-2 h-2 bg-white rounded-full"></div>
              </div>
            </div>

            <div className="space-y-2">
              <h3 className="font-semibold text-lg text-foreground">{profile.name}</h3>
              <Badge variant="secondary" className="bg-primary/10 text-primary border-primary/20 font-medium">
                <Shield className="h-3 w-3 mr-1" />
                {userStats.role}
              </Badge>
            </div>
          </div>

          {/* User Stats */}
          <div className="grid grid-cols-2 gap-3">
            <div className="bg-muted/30 rounded-lg p-3 text-center">
              <div className="flex items-center justify-center gap-1 text-muted-foreground mb-1">
                <Clock className="h-3 w-3" />
                <span className="text-xs">Último acceso</span>
              </div>
              <div className="text-sm font-medium">{userStats.lastLogin}</div>
            </div>
            <div className="bg-muted/30 rounded-lg p-3 text-center">
              <div className="flex items-center justify-center gap-1 text-muted-foreground mb-1">
                <MapPin className="h-3 w-3" />
                <span className="text-xs">Ubicación</span>
              </div>
              <div className="text-sm font-medium">{userStats.location}</div>
            </div>
          </div>

          {/* Quick Actions */}
          <div className="space-y-2">
            <Link href="https://accounts.regionayacucho.gob.pe/user" target="_blank">
              <Button variant="ghost" className="w-full justify-between hover:bg-accent/50 rounded-lg h-12">
                <div className="flex items-center gap-3">
                  <div className="w-8 h-8 rounded-lg bg-primary/10 flex items-center justify-center">
                    <User className="h-4 w-4 text-primary" />
                  </div>
                  <div className="text-left">
                    <div className="font-medium text-sm">Mi Perfil</div>
                    <div className="text-xs text-muted-foreground">Gestionar información personal</div>
                  </div>
                </div>
                <ChevronRight className="h-4 w-4 text-muted-foreground" />
              </Button>
            </Link>

            <Link href="/dashboard/settings">
              <Button variant="ghost" className="w-full justify-between hover:bg-accent/50 rounded-lg h-12">
                <div className="flex items-center gap-3">
                  <div className="w-8 h-8 rounded-lg bg-orange-500/10 flex items-center justify-center">
                    <Settings className="h-4 w-4 text-orange-600" />
                  </div>
                  <div className="text-left">
                    <div className="font-medium text-sm">Configuración</div>
                    <div className="text-xs text-muted-foreground">Preferencias del sistema</div>
                  </div>
                </div>
                <ChevronRight className="h-4 w-4 text-muted-foreground" />
              </Button>
            </Link>
          </div>

          {/* Logout Button */}
          <div className="pt-2 border-t border-border/50">
            <Button
              variant="ghost"
              className="w-full justify-center bg-destructive/5 hover:bg-destructive/10 text-destructive hover:text-destructive rounded-lg h-11 font-medium"
            >
              <LogOut className="h-4 w-4 mr-2" />
              Cerrar Sesión
            </Button>
          </div>

          {/* Footer Info */}
          <div className="pt-2 border-t border-border/50">
            <div className="flex items-center justify-between text-xs text-muted-foreground">
              <span>Sesiones activas: {userStats.sessionsActive}</span>
              <Badge variant="outline" className="text-xs">
                Conectado
              </Badge>
            </div>
          </div>
        </div>
      </PopoverContent>
    </Popover>
  )
}
