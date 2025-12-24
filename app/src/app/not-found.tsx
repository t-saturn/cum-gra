import Link from 'next/link';
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { FileQuestion, ArrowLeft, LayoutDashboard } from 'lucide-react';

export default function NotFound() {
  return (
    <div className="min-h-screen flex items-center justify-center bg-background p-4">
      <Card className="w-full max-w-md shadow-lg border-muted">
        <CardHeader className="flex flex-col items-center space-y-2 pb-2">
          <div className="rounded-full bg-muted p-4 mb-2">
            <FileQuestion className="h-10 w-10 text-muted-foreground" />
          </div>
          <CardTitle className="text-3xl font-bold tracking-tight text-center">
            404
          </CardTitle>
          <p className="text-xl font-medium text-center text-foreground">
            Página no encontrada
          </p>
        </CardHeader>
        
        <CardContent className="text-center">
          <p className="text-muted-foreground">
            Lo sentimos, parece que la página que estás buscando no existe, ha sido movida o no tienes acceso a ella en este momento.
          </p>
        </CardContent>

        <CardFooter className="flex flex-col space-y-2 sm:flex-row sm:space-y-0 sm:space-x-2 justify-center">
          <Button variant="outline" asChild className="w-full sm:w-auto">
            {/* Usamos window.history.back() si fuera un componente cliente, 
                pero en un Server Component 404 es mejor un Link seguro */}
            <Link href="/dashboard">
              <ArrowLeft className="mr-2 h-4 w-4" />
              Volver
            </Link>
          </Button>
          
          <Button asChild className="w-full sm:w-auto">
            <Link href="/dashboard">
              <LayoutDashboard className="mr-2 h-4 w-4" />
              Ir al Dashboard
            </Link>
          </Button>
        </CardFooter>
      </Card>
    </div>
  );
}