'use client'

import { useState } from 'react'
import { Badge } from '@/components/ui/badge'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { Input } from '@/components/ui/input'
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select'
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table'
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs'
import { Switch } from '@/components/ui/switch'
import { Collapsible, CollapsibleContent, CollapsibleTrigger } from '@/components/ui/collapsible'
import { Search, Filter, Shield, Lock, Eye, Edit, Trash2, ChevronDown, ChevronRight, Settings, Users, Database, BarChart3, Download } from 'lucide-react'

// Datos simulados de permisos de módulos
const modulePermissions = [
  {
    id: 'perm_001',
    moduleId: 'mod_001',
    moduleName: 'Gestión de Usuarios',
    applicationName: 'CRM System',
    permissions: [
      {
        id: 'users.read',
        name: 'Ver usuarios',
        description: 'Permite visualizar la lista de usuarios',
        category: 'Lectura',
        isActive: true,
        riskLevel: 'low',
        assignedRoles: ['Admin', 'Manager', 'HR'],
        assignedUsers: 45
      },
      {
        id: 'users.write',
        name: 'Editar usuarios',
        description: 'Permite modificar información de usuarios',
        category: 'Escritura',
        isActive: true,
        riskLevel: 'medium',
        assignedRoles: ['Admin', 'HR'],
        assignedUsers: 12
      },
      {
        id: 'users.delete',
        name: 'Eliminar usuarios',
        description: 'Permite eliminar usuarios del sistema',
        category: 'Eliminación',
        isActive: true,
        riskLevel: 'high',
        assignedRoles: ['Admin'],
        assignedUsers: 3
      },
      {
        id: 'users.export',
        name: 'Exportar usuarios',
        description: 'Permite exportar datos de usuarios',
        category: 'Exportación',
        isActive: true,
        riskLevel: 'medium',
        assignedRoles: ['Admin', 'Manager'],
        assignedUsers: 8
      }
    ]
  },
  {
    id: 'perm_002',
    moduleId: 'mod_002',
    moduleName: 'Gestión de Leads',
    applicationName: 'CRM System',
    permissions: [
      {
        id: 'leads.read',
        name: 'Ver leads',
        description: 'Permite visualizar leads y oportunidades',
        category: 'Lectura',
        isActive: true,
        riskLevel: 'low',
        assignedRoles: ['Admin', 'Sales', 'Manager'],
        assignedUsers: 67
      },
      {
        id: 'leads.write',
        name: 'Editar leads',
        description: 'Permite crear y modificar leads',
        category: 'Escritura',
        isActive: true,
        riskLevel: 'medium',
        assignedRoles: ['Admin', 'Sales'],
        assignedUsers: 34
      },
      {
        id: 'leads.assign',
        name: 'Asignar leads',
        description: 'Permite asignar leads a vendedores',
        category: 'Asignación',
        isActive: true,
        riskLevel: 'medium',
        assignedRoles: ['Admin', 'Manager'],
        assignedUsers: 15
      },
      {
        id: 'leads.convert',
        name: 'Convertir leads',
        description: 'Permite convertir leads en oportunidades',
        category: 'Conversión',
        isActive: true,
        riskLevel: 'high',
        assignedRoles: ['Admin', 'Sales'],
        assignedUsers: 28
      }
    ]
  },
  {
    id: 'perm_003',
    moduleId: 'mod_003',
    moduleName: 'Módulo Financiero',
    applicationName: 'ERP System',
    permissions: [
      {
        id: 'finance.read',
        name: 'Ver finanzas',
        description: 'Permite visualizar información financiera',
        category: 'Lectura',
        isActive: true,
        riskLevel: 'medium',
        assignedRoles: ['Admin', 'Finance', 'Manager'],
        assignedUsers: 23
      },
      {
        id: 'finance.write',
        name: 'Editar finanzas',
        description: 'Permite modificar registros financieros',
        category: 'Escritura',
        isActive: true,
        riskLevel: 'high',
        assignedRoles: ['Admin', 'Finance'],
        assignedUsers: 8
      },
      {
        id: 'finance.approve',
        name: 'Aprobar transacciones',
        description: 'Permite aprobar transacciones financieras',
        category: 'Aprobación',
        isActive: true,
        riskLevel: 'high',
        assignedRoles: ['Admin', 'Finance Manager'],
        assignedUsers: 5
      },
      {
        id: 'finance.reports',
        name: 'Reportes financieros',
        description: 'Permite generar reportes financieros',
        category: 'Reportes',
        isActive: true,
        riskLevel: 'medium',
        assignedRoles: ['Admin', 'Finance', 'Manager'],
        assignedUsers: 18
      }
    ]
  },
  {
    id: 'perm_004',
    moduleId: 'mod_004',
    moduleName: 'Inventario',
    applicationName: 'ERP System',
    permissions: [
      {
        id: 'inventory.read',
        name: 'Ver inventario',
        description: 'Permite visualizar stock y productos',
        category: 'Lectura',
        isActive: true,
        riskLevel: 'low',
        assignedRoles: ['Admin', 'Warehouse', 'Sales'],
        assignedUsers: 56
      },
      {
        id: 'inventory.write',
        name: 'Editar inventario',
        description: 'Permite modificar stock y productos',
        category: 'Escritura',
        isActive: true,
        riskLevel: 'medium',
        assignedRoles: ['Admin', 'Warehouse'],
        assignedUsers: 23
      },
      {
        id: 'inventory.transfer',
        name: 'Transferir stock',
        description: 'Permite transferir productos entre almacenes',
        category: 'Transferencia',
        isActive: true,
        riskLevel: 'medium',
        assignedRoles: ['Admin', 'Warehouse'],
        assignedUsers: 15
      },
      {
        id: 'inventory.adjust',
        name: 'Ajustar inventario',
        description: 'Permite realizar ajustes de inventario',
        category: 'Ajuste',
        isActive: true,
        riskLevel: 'high',
        assignedRoles: ['Admin', 'Warehouse Manager'],
        assignedUsers: 7
      }
    ]
  }
]

const permissionStats = {
  total: 234,
  active: 198,
  inactive: 36,
  byRisk: {
    low: 89,
    medium: 102,
    high: 43
  },
  byCategory: {
    'Lectura': 78,
    'Escritura': 65,
    'Eliminación': 23,
    'Exportación': 34,
    'Aprobación': 18,
    'Reportes': 16
  }
}

export default function ModulePermissionsManagement() {
  const [searchTerm, setSearchTerm] = useState('')
  const [moduleFilter, setModuleFilter] = useState('all')
  const [categoryFilter, setCategoryFilter] = useState('all')
  const [riskFilter, setRiskFilter] = useState('all')
  const [expandedModules, setExpandedModules] = useState<string[]>([])

  const getRiskBadge = (risk: string) => {
    const colors = {
      low: 'bg-green-100 text-green-800',
      medium: 'bg-yellow-100 text-yellow-800',
      high: 'bg-red-100 text-red-800'
    }
    return colors[risk as keyof typeof colors] || 'bg-gray-100 text-gray-800'
  }

  const getCategoryIcon = (category: string) => {
    switch (category.toLowerCase()) {
      case 'lectura': return <Eye className="h-4 w-4" />
      case 'escritura': return <Edit className="h-4 w-4" />
      case 'eliminación': return <Trash2 className="h-4 w-4" />
      case 'exportación': return <Download className="h-4 w-4" />
      case 'aprobación': return <Shield className="h-4 w-4" />
      case 'reportes': return <BarChart3 className="h-4 w-4" />
      default: return <Settings className="h-4 w-4" />
    }
  }

  const toggleModuleExpansion = (moduleId: string) => {
    setExpandedModules(prev => 
      prev.includes(moduleId) 
        ? prev.filter(id => id !== moduleId)
        : [...prev, moduleId]
    )
  }

  const filteredModules = modulePermissions.filter(module => {
    const matchesSearch = module.moduleName.toLowerCase().includes(searchTerm.toLowerCase()) ||
                         module.applicationName.toLowerCase().includes(searchTerm.toLowerCase())
    const matchesModule = moduleFilter === 'all' || module.moduleName === moduleFilter

    return matchesSearch && matchesModule
  })

  const getAllPermissions = () => {
    return modulePermissions.flatMap(module => 
      module.permissions.map(perm => ({
        ...perm,
        moduleName: module.moduleName,
        applicationName: module.applicationName
      }))
    )
  }

  const filteredPermissions = getAllPermissions().filter(permission => {
    const matchesSearch = permission.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
                         permission.description.toLowerCase().includes(searchTerm.toLowerCase())
    const matchesCategory = categoryFilter === 'all' || permission.category === categoryFilter
    const matchesRisk = riskFilter === 'all' || permission.riskLevel === riskFilter

    return matchesSearch && matchesCategory && matchesRisk
  })

  return (
    <div className="space-y-6">
      {/* Estadísticas */}
      <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Total Permisos</CardTitle>
            <Shield className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">{permissionStats.total}</div>
            <p className="text-xs text-muted-foreground">
              Across all modules
            </p>
          </CardContent>
        </Card>
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Permisos Activos</CardTitle>
            <div className="h-2 w-2 bg-green-500 rounded-full animate-pulse" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold text-green-600">{permissionStats.active}</div>
            <p className="text-xs text-muted-foreground">
              {Math.round((permissionStats.active / permissionStats.total) * 100)}% del total
            </p>
          </CardContent>
        </Card>
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Alto Riesgo</CardTitle>
            <Lock className="h-4 w-4 text-red-500" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold text-red-600">{permissionStats.byRisk.high}</div>
            <p className="text-xs text-muted-foreground">
              Requieren supervisión
            </p>
          </CardContent>
        </Card>
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Usuarios Asignados</CardTitle>
            <Users className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">
              {getAllPermissions().reduce((sum, perm) => sum + perm.assignedUsers, 0)}
            </div>
            <p className="text-xs text-muted-foreground">
              Total de asignaciones
            </p>
          </CardContent>
        </Card>
      </div>

      <Tabs defaultValue="modules" className="space-y-4">
        <TabsList>
          <TabsTrigger value="modules">Por Módulos</TabsTrigger>
          <TabsTrigger value="permissions">Todos los Permisos</TabsTrigger>
          <TabsTrigger value="matrix">Matriz de Permisos</TabsTrigger>
          <TabsTrigger value="analytics">Análisis</TabsTrigger>
        </TabsList>

        <TabsContent value="modules" className="space-y-4">
          <Card>
            <CardHeader>
              <CardTitle>Permisos por Módulo</CardTitle>
              <CardDescription>
                Vista organizada de permisos agrupados por módulo del sistema
              </CardDescription>
            </CardHeader>
            <CardContent>
              {/* Filtros */}
              <div className="flex flex-col gap-4 md:flex-row md:items-center md:justify-between mb-6">
                <div className="flex flex-1 items-center space-x-2">
                  <Search className="h-4 w-4 text-muted-foreground" />
                  <Input
                    placeholder="Buscar módulos..."
                    value={searchTerm}
                    onChange={(e) => setSearchTerm(e.target.value)}
                    className="max-w-sm"
                  />
                </div>
                <div className="flex items-center space-x-2">
                  <Filter className="h-4 w-4 text-muted-foreground" />
                  <Select value={moduleFilter} onValueChange={setModuleFilter}>
                    <SelectTrigger className="w-[180px]">
                      <SelectValue placeholder="Módulo" />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value="all">Todos los módulos</SelectItem>
                      <SelectItem value="Gestión de Usuarios">Gestión de Usuarios</SelectItem>
                      <SelectItem value="Gestión de Leads">Gestión de Leads</SelectItem>
                      <SelectItem value="Módulo Financiero">Módulo Financiero</SelectItem>
                      <SelectItem value="Inventario">Inventario</SelectItem>
                    </SelectContent>
                  </Select>
                  <Button variant="outline" size="sm">
                    <Download className="h-4 w-4 mr-2" />
                    Exportar
                  </Button>
                </div>
              </div>

              {/* Lista de módulos */}
              <div className="space-y-4">
                {filteredModules.map((module) => (
                  <Card key={module.id}>
                    <CardHeader>
                      <div className="flex items-center justify-between">
                        <div className="flex items-center space-x-3">
                          <Database className="h-5 w-5 text-muted-foreground" />
                          <div>
                            <CardTitle className="text-lg">{module.moduleName}</CardTitle>
                            <CardDescription>{module.applicationName}</CardDescription>
                          </div>
                        </div>
                        <div className="flex items-center space-x-2">
                          <Badge variant="outline">
                            {module.permissions.length} permisos
                          </Badge>
                          <Collapsible>
                            <CollapsibleTrigger asChild>
                              <Button 
                                variant="ghost" 
                                size="sm"
                                onClick={() => toggleModuleExpansion(module.id)}
                              >
                                {expandedModules.includes(module.id) ? (
                                  <ChevronDown className="h-4 w-4" />
                                ) : (
                                  <ChevronRight className="h-4 w-4" />
                                )}
                              </Button>
                            </CollapsibleTrigger>
                          </Collapsible>
                        </div>
                      </div>
                    </CardHeader>
                    <Collapsible open={expandedModules.includes(module.id)}>
                      <CollapsibleContent>
                        <CardContent>
                          <div className="grid gap-4">
                            {module.permissions.map((permission) => (
                              <div key={permission.id} className="flex items-center justify-between p-4 border rounded-lg">
                                <div className="flex items-center space-x-3">
                                  {getCategoryIcon(permission.category)}
                                  <div>
                                    <div className="font-medium">{permission.name}</div>
                                    <div className="text-sm text-muted-foreground">{permission.description}</div>
                                  </div>
                                </div>
                                <div className="flex items-center space-x-3">
                                  <Badge className={getRiskBadge(permission.riskLevel)}>
                                    {permission.riskLevel === 'low' ? 'Bajo' : 
                                     permission.riskLevel === 'medium' ? 'Medio' : 'Alto'}
                                  </Badge>
                                  <Badge variant="outline">
                                    {permission.category}
                                  </Badge>
                                  <div className="flex items-center space-x-2">
                                    <Users className="h-4 w-4 text-muted-foreground" />
                                    <span className="text-sm">{permission.assignedUsers}</span>
                                  </div>
                                  <Switch checked={permission.isActive} />
                                </div>
                              </div>
                            ))}
                          </div>
                        </CardContent>
                      </CollapsibleContent>
                    </Collapsible>
                  </Card>
                ))}
              </div>
            </CardContent>
          </Card>
        </TabsContent>

        <TabsContent value="permissions" className="space-y-4">
          <Card>
            <CardHeader>
              <CardTitle>Todos los Permisos</CardTitle>
              <CardDescription>
                Lista completa de todos los permisos del sistema
              </CardDescription>
            </CardHeader>
            <CardContent>
              {/* Filtros */}
              <div className="flex flex-col gap-4 md:flex-row md:items-center md:justify-between mb-6">
                <div className="flex flex-1 items-center space-x-2">
                  <Search className="h-4 w-4 text-muted-foreground" />
                  <Input
                    placeholder="Buscar permisos..."
                    value={searchTerm}
                    onChange={(e) => setSearchTerm(e.target.value)}
                    className="max-w-sm"
                  />
                </div>
                <div className="flex items-center space-x-2">
                  <Filter className="h-4 w-4 text-muted-foreground" />
                  <Select value={categoryFilter} onValueChange={setCategoryFilter}>
                    <SelectTrigger className="w-[140px]">
                      <SelectValue placeholder="Categoría" />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value="all">Todas</SelectItem>
                      <SelectItem value="Lectura">Lectura</SelectItem>
                      <SelectItem value="Escritura">Escritura</SelectItem>
                      <SelectItem value="Eliminación">Eliminación</SelectItem>
                      <SelectItem value="Exportación">Exportación</SelectItem>
                      <SelectItem value="Aprobación">Aprobación</SelectItem>
                      <SelectItem value="Reportes">Reportes</SelectItem>
                    </SelectContent>
                  </Select>
                  <Select value={riskFilter} onValueChange={setRiskFilter}>
                    <SelectTrigger className="w-[140px]">
                      <SelectValue placeholder="Riesgo" />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value="all">Todos</SelectItem>
                      <SelectItem value="low">Bajo</SelectItem>
                      <SelectItem value="medium">Medio</SelectItem>
                      <SelectItem value="high">Alto</SelectItem>
                    </SelectContent>
                  </Select>
                </div>
              </div>

              {/* Tabla de permisos */}
              <div className="rounded-md border">
                <Table>
                  <TableHeader>
                    <TableRow>
                      <TableHead>Permiso</TableHead>
                      <TableHead>Módulo</TableHead>
                      <TableHead>Categoría</TableHead>
                      <TableHead>Riesgo</TableHead>
                      <TableHead>Roles Asignados</TableHead>
                      <TableHead>Usuarios</TableHead>
                      <TableHead>Estado</TableHead>
                      <TableHead className="text-right">Acciones</TableHead>
                    </TableRow>
                  </TableHeader>
                  <TableBody>
                    {filteredPermissions.map((permission) => (
                      <TableRow key={permission.id}>
                        <TableCell>
                          <div className="flex items-center space-x-2">
                            {getCategoryIcon(permission.category)}
                            <div>
                              <div className="font-medium">{permission.name}</div>
                              <div className="text-sm text-muted-foreground">{permission.description}</div>
                            </div>
                          </div>
                        </TableCell>
                        <TableCell>
                          <div>
                            <div className="font-medium">{permission.moduleName}</div>
                            <div className="text-sm text-muted-foreground">{permission.applicationName}</div>
                          </div>
                        </TableCell>
                        <TableCell>
                          <Badge variant="outline">{permission.category}</Badge>
                        </TableCell>
                        <TableCell>
                          <Badge className={getRiskBadge(permission.riskLevel)}>
                            {permission.riskLevel === 'low' ? 'Bajo' : 
                             permission.riskLevel === 'medium' ? 'Medio' : 'Alto'}
                          </Badge>
                        </TableCell>
                        <TableCell>
                          <div className="flex flex-wrap gap-1">
                            {permission.assignedRoles.slice(0, 2).map((role, index) => (
                              <Badge key={index} variant="secondary" className="text-xs">
                                {role}
                              </Badge>
                            ))}
                            {permission.assignedRoles.length > 2 && (
                              <Badge variant="secondary" className="text-xs">
                                +{permission.assignedRoles.length - 2}
                              </Badge>
                            )}
                          </div>
                        </TableCell>
                        <TableCell>
                          <div className="flex items-center space-x-2">
                            <Users className="h-4 w-4 text-muted-foreground" />
                            <span className="font-medium">{permission.assignedUsers}</span>
                          </div>
                        </TableCell>
                        <TableCell>
                          <Switch checked={permission.isActive} />
                        </TableCell>
                        <TableCell className="text-right">
                          <div className="flex items-center justify-end space-x-2">
                            <Button variant="ghost" size="sm">
                              <Eye className="h-4 w-4" />
                            </Button>
                            <Button variant="ghost" size="sm">
                              <Edit className="h-4 w-4" />
                            </Button>
                          </div>
                        </TableCell>
                      </TableRow>
                    ))}
                  </TableBody>
                </Table>
              </div>
            </CardContent>
          </Card>
        </TabsContent>

        <TabsContent value="matrix" className="space-y-4">
          <Card>
            <CardHeader>
              <CardTitle>Matriz de Permisos</CardTitle>
              <CardDescription>
                Vista consolidada de permisos por categoría y nivel de riesgo
              </CardDescription>
            </CardHeader>
            <CardContent>
              <div className="space-y-6">
                {Object.entries(permissionStats.byCategory).map(([category, count]) => (
                  <div key={category} className="space-y-4">
                    <div className="flex items-center justify-between">
                      <h3 className="text-lg font-semibold flex items-center space-x-2">
                        {getCategoryIcon(category)}
                        <span>{category}</span>
                      </h3>
                      <Badge variant="outline">{count} permisos</Badge>
                    </div>
                    <div className="grid gap-2">
                      {filteredPermissions
                        .filter(perm => perm.category === category)
                        .map((permission) => (
                          <div key={permission.id} className="flex items-center justify-between p-3 border rounded">
                            <div className="flex items-center space-x-3">
                              <div>
                                <div className="font-medium">{permission.name}</div>
                                <div className="text-sm text-muted-foreground">
                                  {permission.moduleName} - {permission.applicationName}
                                </div>
                              </div>
                            </div>
                            <div className="flex items-center space-x-2">
                              <Badge className={getRiskBadge(permission.riskLevel)}>
                                {permission.riskLevel === 'low' ? 'Bajo' : 
                                 permission.riskLevel === 'medium' ? 'Medio' : 'Alto'}
                              </Badge>
                              <div className="flex items-center space-x-1">
                                <Users className="h-4 w-4 text-muted-foreground" />
                                <span className="text-sm">{permission.assignedUsers}</span>
                              </div>
                              <Switch checked={permission.isActive} />
                            </div>
                          </div>
                        ))}
                    </div>
                  </div>
                ))}
              </div>
            </CardContent>
          </Card>
        </TabsContent>

        <TabsContent value="analytics" className="space-y-4">
          <div className="grid gap-4 md:grid-cols-2">
            <Card>
              <CardHeader>
                <CardTitle>Distribución por Riesgo</CardTitle>
              </CardHeader>
              <CardContent>
                <div className="space-y-4">
                  {Object.entries(permissionStats.byRisk).map(([risk, count]) => (
                    <div key={risk} className="flex items-center justify-between">
                      <div className="flex items-center space-x-2">
                        <div className={`w-3 h-3 rounded-full ${
                          risk === 'low' ? 'bg-green-500' :
                          risk === 'medium' ? 'bg-yellow-500' : 'bg-red-500'
                        }`}></div>
                        <span className="capitalize">Riesgo {risk === 'low' ? 'Bajo' : risk === 'medium' ? 'Medio' : 'Alto'}</span>
                      </div>
                      <div className="flex items-center space-x-2">
                        <div className="w-24 bg-gray-200 rounded-full h-2">
                          <div className={`h-2 rounded-full ${
                            risk === 'low' ? 'bg-green-500' :
                            risk === 'medium' ? 'bg-yellow-500' : 'bg-red-500'
                          }`} style={{width: `${(count / permissionStats.total) * 100}%`}}></div>
                        </div>
                        <span className="text-sm font-medium">{count}</span>
                      </div>
                    </div>
                  ))}
                </div>
              </CardContent>
            </Card>

            <Card>
              <CardHeader>
                <CardTitle>Permisos por Categoría</CardTitle>
              </CardHeader>
              <CardContent>
                <div className="space-y-4">
                  {Object.entries(permissionStats.byCategory).map(([category, count]) => (
                    <div key={category} className="flex items-center justify-between">
                      <div className="flex items-center space-x-2">
                        {getCategoryIcon(category)}
                        <span>{category}</span>
                      </div>
                      <div className="flex items-center space-x-2">
                        <div className="w-24 bg-gray-200 rounded-full h-2">
                          <div className="bg-blue-600 h-2 rounded-full" style={{width: `${(count / permissionStats.total) * 100}%`}}></div>
                        </div>
                        <span className="text-sm font-medium">{count}</span>
                      </div>
                    </div>
                  ))}
                </div>
              </CardContent>
            </Card>
          </div>
        </TabsContent>
      </Tabs>
    </div>
  )
}
