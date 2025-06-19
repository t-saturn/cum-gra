"use client"

import { useState } from "react"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Badge } from "@/components/ui/badge"
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar"
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import {
  Search,
  Plus,
  Filter,
  Download,
  MoreHorizontal,
  Edit,
  Trash2,
  Shield,
  Eye,
  UserCheck,
  UserX,
} from "lucide-react"

// Mock data
const users = [
  {
    id: "1",
    email: "juan.perez@empresa.com",
    firstName: "Juan",
    lastName: "Pérez",
    dni: "12345678",
    phone: "+51 999 888 777",
    status: "active",
    emailVerified: true,
    phoneVerified: false,
    twoFactorEnabled: true,
    structuralPosition: "Gerente General",
    organicUnit: "Gerencia",
    lastLogin: "2024-01-15 10:30:00",
    createdAt: "2024-01-01 09:00:00",
  },
  {
    id: "2",
    email: "maria.garcia@empresa.com",
    firstName: "María",
    lastName: "García",
    dni: "87654321",
    phone: "+51 888 777 666",
    status: "active",
    emailVerified: true,
    phoneVerified: true,
    twoFactorEnabled: false,
    structuralPosition: "Analista Senior",
    organicUnit: "Sistemas",
    lastLogin: "2024-01-15 08:15:00",
    createdAt: "2024-01-02 14:30:00",
  },
  {
    id: "3",
    email: "carlos.lopez@empresa.com",
    firstName: "Carlos",
    lastName: "López",
    dni: "11223344",
    phone: "+51 777 666 555",
    status: "suspended",
    emailVerified: false,
    phoneVerified: false,
    twoFactorEnabled: false,
    structuralPosition: "Desarrollador",
    organicUnit: "Sistemas",
    lastLogin: "2024-01-10 16:45:00",
    createdAt: "2024-01-03 11:15:00",
  },
]

export default function UsersManagement() {
  const [searchTerm, setSearchTerm] = useState("")

  const filteredUsers = users.filter(
    (user) =>
      user.email.toLowerCase().includes(searchTerm.toLowerCase()) ||
      user.firstName.toLowerCase().includes(searchTerm.toLowerCase()) ||
      user.lastName.toLowerCase().includes(searchTerm.toLowerCase()) ||
      user.dni.includes(searchTerm),
  )

  const getStatusBadge = (status: string) => {
    switch (status) {
      case "active":
        return <Badge className="bg-chart-4/20 text-chart-4 border-chart-4/30">Activo</Badge>
      case "suspended":
        return <Badge className="bg-chart-5/20 text-chart-5 border-chart-5/30">Suspendido</Badge>
      case "deleted":
        return <Badge className="bg-destructive/20 text-destructive border-destructive/30">Eliminado</Badge>
      default:
        return <Badge variant="secondary">Desconocido</Badge>
    }
  }

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold text-foreground">Gestión de Usuarios</h1>
          <p className="text-muted-foreground mt-1">Administra todos los usuarios del sistema</p>
        </div>
        <div className="flex gap-3">
          <Button variant="outline">
            <Download className="w-4 h-4 mr-2" />
            Exportar
          </Button>
          <Button className="bg-gradient-to-r from-primary to-chart-1 hover:from-primary/90 hover:to-chart-1/90 shadow-lg shadow-primary/25">
            <Plus className="w-4 h-4 mr-2" />
            Nuevo Usuario
          </Button>
        </div>
      </div>

      {/* Stats Cards */}
      <div className="grid grid-cols-1 md:grid-cols-4 gap-4">
        <Card className="border-border bg-card/50">
          <CardContent className="p-4">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm text-muted-foreground">Total Usuarios</p>
                <p className="text-2xl font-bold text-foreground">2,847</p>
              </div>
              <UserCheck className="w-8 h-8 text-chart-2" />
            </div>
          </CardContent>
        </Card>
        <Card className="border-border bg-card/50">
          <CardContent className="p-4">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm text-muted-foreground">Usuarios Activos</p>
                <p className="text-2xl font-bold text-chart-4">2,654</p>
              </div>
              <UserCheck className="w-8 h-8 text-chart-4" />
            </div>
          </CardContent>
        </Card>
        <Card className="border-border bg-card/50">
          <CardContent className="p-4">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm text-muted-foreground">Suspendidos</p>
                <p className="text-2xl font-bold text-chart-5">156</p>
              </div>
              <UserX className="w-8 h-8 text-chart-5" />
            </div>
          </CardContent>
        </Card>
        <Card className="border-border bg-card/50">
          <CardContent className="p-4">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm text-muted-foreground">Nuevos (30d)</p>
                <p className="text-2xl font-bold text-primary">37</p>
              </div>
              <Plus className="w-8 h-8 text-primary" />
            </div>
          </CardContent>
        </Card>
      </div>

      {/* Filters and Search */}
      <Card className="border-border bg-card/50">
        <CardHeader>
          <div className="flex items-center justify-between">
            <div>
              <CardTitle>Lista de Usuarios</CardTitle>
              <CardDescription>
                {filteredUsers.length} de {users.length} usuarios
              </CardDescription>
            </div>
            <div className="flex gap-2">
              <div className="relative">
                <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 text-muted-foreground w-4 h-4" />
                <Input
                  placeholder="Buscar usuarios..."
                  value={searchTerm}
                  onChange={(e) => setSearchTerm(e.target.value)}
                  className="pl-10 w-80 bg-background/50 border-border focus:border-primary focus:ring-ring"
                />
              </div>
              <Button variant="outline">
                <Filter className="w-4 h-4 mr-2" />
                Filtros
              </Button>
            </div>
          </div>
        </CardHeader>
        <CardContent>
          <div className="rounded-lg border border-border">
            <Table>
              <TableHeader>
                <TableRow className="bg-accent/50">
                  <TableHead>Usuario</TableHead>
                  <TableHead>DNI</TableHead>
                  <TableHead>Teléfono</TableHead>
                  <TableHead>Estado</TableHead>
                  <TableHead>Verificación</TableHead>
                  <TableHead>Posición</TableHead>
                  <TableHead>Último Acceso</TableHead>
                  <TableHead className="text-right">Acciones</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {filteredUsers.map((user) => (
                  <TableRow key={user.id} className="hover:bg-accent/30">
                    <TableCell>
                      <div className="flex items-center gap-3">
                        <Avatar className="w-10 h-10">
                          <AvatarImage src={`/placeholder.svg?height=40&width=40`} />
                          <AvatarFallback className="bg-gradient-to-r from-primary to-chart-1 text-primary-foreground font-semibold">
                            {user.firstName[0]}
                            {user.lastName[0]}
                          </AvatarFallback>
                        </Avatar>
                        <div>
                          <p className="font-medium text-foreground">
                            {user.firstName} {user.lastName}
                          </p>
                          <p className="text-sm text-muted-foreground">{user.email}</p>
                        </div>
                      </div>
                    </TableCell>
                    <TableCell className="font-mono text-sm">{user.dni}</TableCell>
                    <TableCell className="text-sm">{user.phone}</TableCell>
                    <TableCell>{getStatusBadge(user.status)}</TableCell>
                    <TableCell>
                      <div className="flex gap-1">
                        <Badge variant={user.emailVerified ? "default" : "secondary"} className="text-xs">
                          {user.emailVerified ? "✓" : "✗"} Email
                        </Badge>
                        <Badge variant={user.phoneVerified ? "default" : "secondary"} className="text-xs">
                          {user.phoneVerified ? "✓" : "✗"} Tel
                        </Badge>
                        {user.twoFactorEnabled && (
                          <Badge className="text-xs bg-chart-4/20 text-chart-4 border-chart-4/30">2FA</Badge>
                        )}
                      </div>
                    </TableCell>
                    <TableCell>
                      <div>
                        <p className="text-sm font-medium">{user.structuralPosition}</p>
                        <p className="text-xs text-muted-foreground">{user.organicUnit}</p>
                      </div>
                    </TableCell>
                    <TableCell className="text-sm">{new Date(user.lastLogin).toLocaleDateString("es-ES")}</TableCell>
                    <TableCell className="text-right">
                      <DropdownMenu>
                        <DropdownMenuTrigger asChild>
                          <Button variant="ghost" size="sm">
                            <MoreHorizontal className="w-4 h-4" />
                          </Button>
                        </DropdownMenuTrigger>
                        <DropdownMenuContent align="end" className="bg-card/80 backdrop-blur-xl border-border">
                          <DropdownMenuLabel>Acciones</DropdownMenuLabel>
                          <DropdownMenuSeparator />
                          <DropdownMenuItem>
                            <Eye className="w-4 h-4 mr-2" />
                            Ver Detalles
                          </DropdownMenuItem>
                          <DropdownMenuItem>
                            <Edit className="w-4 h-4 mr-2" />
                            Editar
                          </DropdownMenuItem>
                          <DropdownMenuItem>
                            <Shield className="w-4 h-4 mr-2" />
                            Gestionar Roles
                          </DropdownMenuItem>
                          <DropdownMenuSeparator />
                          <DropdownMenuItem className="text-destructive">
                            <Trash2 className="w-4 h-4 mr-2" />
                            Eliminar
                          </DropdownMenuItem>
                        </DropdownMenuContent>
                      </DropdownMenu>
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </div>
        </CardContent>
      </Card>
    </div>
  )
}
