'use client';

import React, { useState } from 'react';
import { Eye, EyeOff, GlobeLock, ScanFace } from 'lucide-react';
import { toast } from 'sonner';
import { useRouter } from 'next/navigation';

import { Card, CardHeader, CardTitle, CardDescription, CardContent, CardFooter } from '../ui/card';
import { Button } from '../ui/button';
import { Label } from '@radix-ui/react-label';
import { Input } from '../ui/input';

export const Login: React.FC = () => {
  const router = useRouter();
  const [passwordVisible, setPasswordVisible] = useState(false);

  const handle_login = () => {
    // Aquí iría la lógica de autenticación
    console.log('Iniciando sesión...');
    toast.success('Inicio de sesión exitoso', {
      position: 'top-right',
      duration: 3000,
    });
    router.push('/');
  };

  return (
    <Card className="w-full max-w-md mx-auto">
      <CardHeader>
        <CardTitle>Login to your account</CardTitle>
        <CardDescription>Enter your email and password below to login</CardDescription>
      </CardHeader>

      <CardContent>
        <form
          onSubmit={(e) => {
            e.preventDefault();
            handle_login();
          }}
          className="flex flex-col gap-6"
        >
          {/* Email */}
          <div className="grid gap-2">
            <Label htmlFor="email">Email</Label>
            <Input id="email" type="email" placeholder="m@example.com" required />
          </div>

          {/* Password */}
          <div className="grid gap-2">
            <div className="flex items-center gap-4">
              <Label htmlFor="password">Password</Label>
              <a href="/forgot" className="ml-auto text-sm underline-offset-4 hover:underline text-red-500">
                Forgot your password?
              </a>
            </div>
            <div className="relative">
              <Input id="password" type={passwordVisible ? 'text' : 'password'} required />
              <button
                type="button"
                onClick={() => setPasswordVisible((v) => !v)}
                aria-label={passwordVisible ? 'Ocultar contraseña' : 'Mostrar contraseña'}
                title={passwordVisible ? 'Ocultar contraseña' : 'Mostrar contraseña'}
                className="absolute right-3 top-1/2 -translate-y-1/2 text-gray-400 focus:outline-none"
              >
                {passwordVisible ? <EyeOff className="w-5 h-5" /> : <Eye className="w-5 h-5" />}
              </button>
            </div>
          </div>

          <Button variant={'ghost'} type="submit" className="w-full bg-red-600 font-bold text-white">
            Login
          </Button>
        </form>
      </CardContent>

      <CardFooter className="flex-col gap-8">
        <Button variant="outline" className="w-full" onClick={handle_login}>
          <GlobeLock className="mr-2 w-5 h-5" />
          Continuar con Google
        </Button>
        <Button variant="outline" className="w-full" onClick={handle_login}>
          <ScanFace className="mr-2 w-5 h-5" />
          Autenticación Facial
        </Button>
      </CardFooter>
    </Card>
  );
};
