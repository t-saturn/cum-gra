"use client"

import { useState } from "react"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Badge } from "@/components/ui/badge"
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs"
import {
  Search,
  Download,
  TrendingUp,
  TrendingDown,
  Users,
  Clock,
  Activity,
  Eye,
  Globe,
  Smartphone,
  Monitor,
  RefreshCw,
} from "lucide-react"
import CardStatsContain from "@/components/custom/card/card-stats-contain"
import { statsAppAccess } from "@/mocks/stats-mocks"

// Mock data
const appAccessData = [
  {
    id: "1",
    application: {
      name: "Sistema de Inventario",
      clientId: "inv_client_12345",
      type: "web",
      icon: "📦",
    },
    statistics: {
      totalAccesses: 15420,
      uniqueUsers: 45,
      averageSessionTime: "2h 15m",
      peakHour: "14:00",
      dailyGrowth: 8.5,
      weeklyGrowth: 15.2,
      errorRate: 0.02,
    },
    accessPatterns: {
      desktop: 75,
      mobile: 20,
      tablet: 5,
    },
    topUsers: [
      { name: "Carlos López", accesses: 245, lastAccess: "2024-01-15 16:30:00" },
      { name: "María García", accesses: 189, lastAccess: "2024-01-15 15:45:00" },
      { name: "Juan Pérez", accesses: 156, lastAccess: "2024-01-15 14:20:00" },
    ],
    hourlyDistribution: [
      { hour: "08:00", accesses: 45 },
      { hour: "09:00", accesses: 78 },
      { hour: "10:00", accesses: 92 },
      { hour: "11:00", accesses: 105 },
      { hour: "12:00", accesses: 68 },
      { hour: "13:00", accesses: 85 },
      { hour: "14:00", accesses: 125 },
      { hour: "15:00", accesses: 110 },
      { hour: "16:00", accesses: 95 },
      { hour: "17:00", accesses: 72 },
    ],
    status: "active",
  },
  {
    id: "2",
    application: {
      name: "Portal de Recursos Humanos",
      clientId: "hr_client_67890",
      type: "web",
      icon: "👥",
    },
    statistics: {
      totalAccesses: 8750,
      uniqueUsers: 32,
      averageSessionTime: "1h 45m",
      peakHour: "09:00",
      dailyGrowth: 5.2,
      weeklyGrowth: 12.8,
      errorRate: 0.01,
    },
    accessPatterns: {
      desktop: 85,
      mobile: 10,
      tablet: 5,
    },
    topUsers: [
      { name: "Luis Torres", accesses: 203, lastAccess: "2024-01-15 12:15:00" },
      { name: "Ana Martínez", accesses: 156, lastAccess: "2024-01-15 14:30:00" },
      { name: "Pedro Ruiz", accesses: 134, lastAccess: "2024-01-15 11:45:00" },
    ],
    hourlyDistribution: [
      { hour: "08:00", accesses: 35 },
      { hour: "09:00", accesses: 95 },
      { hour: "10:00", accesses: 78 },
      { hour: "11:00", accesses: 85 },
      { hour: "12:00", accesses: 45 },
      { hour: "13:00", accesses: 52 },
      { hour: "14:00", accesses: 88 },
      { hour: "15:00", accesses: 75 },
      { hour: "16:00", accesses: 65 },
      { hour: "17:00", accesses: 48 },
    ],
    status: "active",
  },
  {
    id: "3",
    application: {
      name: "App Móvil Ventas",
      clientId: "mobile_sales_11111",
      type: "mobile",
      icon: "📱",
    },
    statistics: {
      totalAccesses: 12800,
      uniqueUsers: 25,
      averageSessionTime: "45m",
      peakHour: "11:00",
      dailyGrowth: 12.3,
      weeklyGrowth: 25.6,
      errorRate: 0.015,
    },
    accessPatterns: {
      desktop: 15,
      mobile: 80,
      tablet: 5,
    },
    topUsers: [
      { name: "Roberto Silva", accesses: 312, lastAccess: "2024-01-15 16:45:00" },
      { name: "Carmen Vega", accesses: 278, lastAccess: "2024-01-15 15:30:00" },
      { name: "Diego Morales", accesses: 245, lastAccess: "2024-01-15 14:15:00" },
    ],
    hourlyDistribution: [
      { hour: "08:00", accesses: 25 },
      { hour: "09:00", accesses: 45 },
      { hour: "10:00", accesses: 68 },
      { hour: "11:00", accesses: 95 },
      { hour: "12:00", accesses: 52 },
      { hour: "13:00", accesses: 38 },
      { hour: "14:00", accesses: 75 },
      { hour: "15:00", accesses: 85 },
      { hour: "16:00", accesses: 78 },
      { hour: "17:00", accesses: 65 },
    ],
    status: "active",
  },
  {
    id: "4",
    application: {
      name: "Sistema Financiero",
      clientId: "finance_client_22222",
      type: "web",
      icon: "💰",
    },
    statistics: {
      totalAccesses: 3200,
      uniqueUsers: 12,
      averageSessionTime: "3h 20m",
      peakHour: "10:00",
      dailyGrowth: -2.1,
      weeklyGrowth: -5.8,
      errorRate: 0.008,
    },
    accessPatterns: {
      desktop: 95,
      mobile: 3,
      tablet: 2,
    },
    topUsers: [
      { name: "Fernando Castro", accesses: 145, lastAccess: "2024-01-10 09:30:00" },
      { name: "Isabel Mendoza", accesses: 98, lastAccess: "2024-01-09 15:20:00" },
      { name: "Andrés Herrera", accesses: 87, lastAccess: "2024-01-08 11:10:00" },
    ],
    hourlyDistribution: [
      { hour: "08:00", accesses: 15 },
      { hour: "09:00", accesses: 25 },
      { hour: "10:00", accesses: 35 },
      { hour: "11:00", accesses: 28 },
      { hour: "12:00", accesses: 12 },
      { hour: "13:00", accesses: 8 },
      { hour: "14:00", accesses: 22 },
      { hour: "15:00", accesses: 18 },
      { hour: "16:00", accesses: 15 },
      { hour: "17:00", accesses: 10 },
    ],
    status: "suspended",
  },
]

export default function AppAccessReport() {
  const [searchTerm, setSearchTerm] = useState("")
  const [typeFilter, setTypeFilter] = useState("all")
  const [statusFilter, setStatusFilter] = useState("all")

  const filteredApps = appAccessData.filter((app) => {
    const matchesSearch = app.application.name.toLowerCase().includes(searchTerm.toLowerCase())
    const matchesType = typeFilter === "all" || app.application.type === typeFilter
    const matchesStatus = statusFilter === "all" || app.status === statusFilter

    return matchesSearch && matchesType && matchesStatus
  })

  const getStatusBadge = (status: string) => {
    switch (status) {
      case "active":
        return (
          <Badge className="bg-chart-4/20 text-chart-4 border-chart-4/30">
            <Activity className="w-3 h-3 mr-1" />
            Activa
          </Badge>
        )
      case "suspended":
        return (
          <Badge className="bg-destructive/20 text-destructive border-destructive/30">
            <Clock className="w-3 h-3 mr-1" />
            Suspendida
          </Badge>
        )
      default:
        return <Badge variant="secondary">Desconocido</Badge>
    }
  }

  const getTypeBadge = (type: string) => {
    switch (type) {
      case "web":
        return (
          <Badge className="bg-primary/20 text-primary border-primary/30">
            <Globe className="w-3 h-3 mr-1" />
            Web
          </Badge>
        )
      case "mobile":
        return (
          <Badge className="bg-chart-2/20 text-chart-2 border-chart-2/30">
            <Smartphone className="w-3 h-3 mr-1" />
            Móvil
          </Badge>
        )
      case "desktop":
        return (
          <Badge className="bg-chart-3/20 text-chart-3 border-chart-3/30">
            <Monitor className="w-3 h-3 mr-1" />
            Escritorio
          </Badge>
        )
      default:
        return <Badge variant="secondary">Otro</Badge>
    }
  }

  const formatDateTime = (dateString: string) => {
    return new Date(dateString).toLocaleString("es-ES")
  }

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold text-foreground">Accesos por Aplicación</h1>
          <p className="text-muted-foreground mt-1">Análisis detallado del uso de aplicaciones del sistema</p>
        </div>
        <div className="flex gap-3">
          <Button variant="outline">
            <RefreshCw className="w-4 h-4 mr-2" />
            Actualizar
          </Button>
          <Button variant="outline">
            <Download className="w-4 h-4 mr-2" />
            Exportar
          </Button>
        </div>
      </div>

      {/* Stats Cards */}
      <CardStatsContain stats={statsAppAccess} />

      <Tabs defaultValue="overview" className="w-full">
        <TabsList className="grid w-full grid-cols-3">
          <TabsTrigger value="overview">Vista General</TabsTrigger>
          <TabsTrigger value="detailed">Análisis Detallado</TabsTrigger>
          <TabsTrigger value="patterns">Patrones de Uso</TabsTrigger>
        </TabsList>

        <TabsContent value="overview" className="space-y-4">
          {/* Filters */}
          <Card className="border-border bg-card/50">
            <CardHeader>
              <div className="flex items-center justify-between">
                <div>
                  <CardTitle>Resumen de Aplicaciones</CardTitle>
                  <CardDescription>
                    {filteredApps.length} de {appAccessData.length} aplicaciones mostradas
                  </CardDescription>
                </div>
                <div className="flex gap-2">
                  <div className="relative">
                    <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 text-muted-foreground w-4 h-4" />
                    <Input
                      placeholder="Buscar aplicaciones..."
                      value={searchTerm}
                      onChange={(e) => setSearchTerm(e.target.value)}
                      className="pl-10 w-80 bg-background/50 border-border focus:border-primary focus:ring-ring"
                    />
                  </div>
                  <Select value={typeFilter} onValueChange={setTypeFilter}>
                    <SelectTrigger className="w-32">
                      <SelectValue placeholder="Tipo" />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value="all">Todos</SelectItem>
                      <SelectItem value="web">Web</SelectItem>
                      <SelectItem value="mobile">Móvil</SelectItem>
                      <SelectItem value="desktop">Escritorio</SelectItem>
                    </SelectContent>
                  </Select>
                  <Select value={statusFilter} onValueChange={setStatusFilter}>
                    <SelectTrigger className="w-32">
                      <SelectValue placeholder="Estado" />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value="all">Todos</SelectItem>
                      <SelectItem value="active">Activa</SelectItem>
                      <SelectItem value="suspended">Suspendida</SelectItem>
                    </SelectContent>
                  </Select>
                </div>
              </div>
            </CardHeader>
            <CardContent>
              <div className="rounded-lg border border-border">
                <Table>
                  <TableHeader>
                    <TableRow className="bg-accent/50">
                      <TableHead>Aplicación</TableHead>
                      <TableHead>Tipo</TableHead>
                      <TableHead>Total Accesos</TableHead>
                      <TableHead>Usuarios Únicos</TableHead>
                      <TableHead>Sesión Promedio</TableHead>
                      <TableHead>Crecimiento</TableHead>
                      <TableHead>Estado</TableHead>
                      <TableHead className="text-right">Acciones</TableHead>
                    </TableRow>
                  </TableHeader>
                  <TableBody>
                    {filteredApps.map((app) => (
                      <TableRow key={app.id} className="hover:bg-accent/30">
                        <TableCell>
                          <div className="flex items-center gap-3">
                            <div className="text-2xl">{app.application.icon}</div>
                            <div>
                              <p className="font-medium text-foreground">{app.application.name}</p>
                              <p className="text-sm text-muted-foreground font-mono">{app.application.clientId}</p>
                            </div>
                          </div>
                        </TableCell>
                        <TableCell>{getTypeBadge(app.application.type)}</TableCell>
                        <TableCell>
                          <div className="text-center">
                            <p className="text-lg font-bold text-primary">
                              {app.statistics.totalAccesses.toLocaleString()}
                            </p>
                            <p className="text-xs text-muted-foreground">Pico: {app.statistics.peakHour}</p>
                          </div>
                        </TableCell>
                        <TableCell>
                          <div className="flex items-center gap-2">
                            <Users className="w-4 h-4 text-chart-2" />
                            <span className="font-medium">{app.statistics.uniqueUsers}</span>
                          </div>
                        </TableCell>
                        <TableCell>
                          <div className="flex items-center gap-2">
                            <Clock className="w-4 h-4 text-chart-3" />
                            <span className="font-medium">{app.statistics.averageSessionTime}</span>
                          </div>
                        </TableCell>
                        <TableCell>
                          <div className="flex items-center gap-1">
                            {app.statistics.dailyGrowth >= 0 ? (
                              <TrendingUp className="w-4 h-4 text-chart-4" />
                            ) : (
                              <TrendingDown className="w-4 h-4 text-destructive" />
                            )}
                            <span className={app.statistics.dailyGrowth >= 0 ? "text-chart-4" : "text-destructive"}>
                              {app.statistics.dailyGrowth >= 0 ? "+" : ""}
                              {app.statistics.dailyGrowth.toFixed(1)}%
                            </span>
                          </div>
                        </TableCell>
                        <TableCell>{getStatusBadge(app.status)}</TableCell>
                        <TableCell className="text-right">
                          <Button variant="ghost" size="sm">
                            <Eye className="w-4 h-4" />
                          </Button>
                        </TableCell>
                      </TableRow>
                    ))}
                  </TableBody>
                </Table>
              </div>
            </CardContent>
          </Card>
        </TabsContent>

        <TabsContent value="detailed" className="space-y-4">
          <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
            {filteredApps.map((app) => (
              <Card key={app.id} className="border-border bg-card/50">
                <CardHeader>
                  <div className="flex items-center justify-between">
                    <div className="flex items-center gap-3">
                      <div className="text-2xl">{app.application.icon}</div>
                      <div>
                        <CardTitle className="text-lg">{app.application.name}</CardTitle>
                        <CardDescription>{app.application.clientId}</CardDescription>
                      </div>
                    </div>
                    {getStatusBadge(app.status)}
                  </div>
                </CardHeader>
                <CardContent className="space-y-4">
                  <div className="grid grid-cols-2 gap-4">
                    <div className="text-center p-3 bg-accent/20 rounded-lg">
                      <p className="text-2xl font-bold text-primary">{app.statistics.totalAccesses.toLocaleString()}</p>
                      <p className="text-sm text-muted-foreground">Total Accesos</p>
                    </div>
                    <div className="text-center p-3 bg-accent/20 rounded-lg">
                      <p className="text-2xl font-bold text-chart-2">{app.statistics.uniqueUsers}</p>
                      <p className="text-sm text-muted-foreground">Usuarios Únicos</p>
                    </div>
                  </div>

                  <div>
                    <h4 className="font-semibold mb-2">Usuarios Más Activos</h4>
                    <div className="space-y-2">
                      {app.topUsers.map((user, index) => (
                        <div key={index} className="flex items-center justify-between p-2 bg-accent/20 rounded">
                          <div>
                            <p className="font-medium text-sm">{user.name}</p>
                            <p className="text-xs text-muted-foreground">Último: {formatDateTime(user.lastAccess)}</p>
                          </div>
                          <Badge variant="outline">{user.accesses} accesos</Badge>
                        </div>
                      ))}
                    </div>
                  </div>

                  <div>
                    <h4 className="font-semibold mb-2">Distribución por Dispositivo</h4>
                    <div className="space-y-2">
                      <div className="flex items-center justify-between">
                        <span className="text-sm">Desktop</span>
                        <div className="flex items-center gap-2">
                          <div className="w-20 bg-muted rounded-full h-2">
                            <div
                              className="bg-primary h-2 rounded-full"
                              style={{ width: `${app.accessPatterns.desktop}%` }}
                            />
                          </div>
                          <span className="text-sm font-medium">{app.accessPatterns.desktop}%</span>
                        </div>
                      </div>
                      <div className="flex items-center justify-between">
                        <span className="text-sm">Mobile</span>
                        <div className="flex items-center gap-2">
                          <div className="w-20 bg-muted rounded-full h-2">
                            <div
                              className="bg-chart-2 h-2 rounded-full"
                              style={{ width: `${app.accessPatterns.mobile}%` }}
                            />
                          </div>
                          <span className="text-sm font-medium">{app.accessPatterns.mobile}%</span>
                        </div>
                      </div>
                      <div className="flex items-center justify-between">
                        <span className="text-sm">Tablet</span>
                        <div className="flex items-center gap-2">
                          <div className="w-20 bg-muted rounded-full h-2">
                            <div
                              className="bg-chart-3 h-2 rounded-full"
                              style={{ width: `${app.accessPatterns.tablet}%` }}
                            />
                          </div>
                          <span className="text-sm font-medium">{app.accessPatterns.tablet}%</span>
                        </div>
                      </div>
                    </div>
                  </div>
                </CardContent>
              </Card>
            ))}
          </div>
        </TabsContent>

        <TabsContent value="patterns" className="space-y-4">
          <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
            {filteredApps.map((app) => (
              <Card key={app.id} className="border-border bg-card/50">
                <CardHeader>
                  <div className="flex items-center gap-3">
                    <div className="text-2xl">{app.application.icon}</div>
                    <div>
                      <CardTitle className="text-lg">{app.application.name}</CardTitle>
                      <CardDescription>Patrón de accesos por hora</CardDescription>
                    </div>
                  </div>
                </CardHeader>
                <CardContent>
                  <div className="space-y-4">
                    <div className="grid grid-cols-10 gap-1">
                      {app.hourlyDistribution.map((data, index) => (
                        <div key={index} className="text-center">
                          <div className="relative h-20 bg-accent/20 rounded flex items-end justify-center p-1">
                            <div
                              className="bg-gradient-to-t from-primary to-chart-1 rounded-sm w-full transition-all duration-300"
                              style={{
                                height: `${
                                  (data.accesses / Math.max(...app.hourlyDistribution.map((h) => h.accesses))) * 100
                                }%`,
                                minHeight: "4px",
                              }}
                            />
                            <div className="absolute -top-5 left-1/2 transform -translate-x-1/2 text-xs font-medium text-foreground">
                              {data.accesses}
                            </div>
                          </div>
                          <p className="text-xs text-muted-foreground mt-1">{data.hour}</p>
                        </div>
                      ))}
                    </div>
                    <div className="flex items-center justify-between text-sm text-muted-foreground">
                      <span>Pico: {app.statistics.peakHour}</span>
                      <span>Máx: {Math.max(...app.hourlyDistribution.map((h) => h.accesses))} accesos</span>
                    </div>
                  </div>
                </CardContent>
              </Card>
            ))}
          </div>
        </TabsContent>
      </Tabs>
    </div>
  )
}
