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
  const [loading, setLoading] = useState(false);

  async function handle_login(e: React.FormEvent) {
    e.preventDefault();
    setLoading(true);

    const form = e.currentTarget as HTMLFormElement;
    const data = {
      email: (form.email as HTMLInputElement).value,
      password: (form.password as HTMLInputElement).value,
    };

    try {
      const res = await fetch('/api/auth/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(data),
      });

      const json = await res.json();
      if (!res.ok) {
        // backendJson.error.details holds the human message
        toast.error(json.error?.details || json.message, { position: 'top-right' });
        setLoading(false);
        return;
      }

      toast.success('Inicio de sesión exitoso', { position: 'top-right' });
      router.push('/');
    } catch (json: any) {
      toast.error('Error en el login', { position: 'top-right' });
    } finally {
      setLoading(false);
    }
  }

  return (
    <Card className="w-full max-w-md mx-auto">
      <CardHeader>
        <CardTitle>Iniciar sesión en su cuenta</CardTitle>
        <CardDescription>Ingrese su correo electrónico y contraseña a continuación para iniciar sesión</CardDescription>
      </CardHeader>

      <CardContent>
        <form onSubmit={handle_login} className="flex flex-col gap-6">
          {/* Email */}
          <div className="grid gap-2">
            <Label htmlFor="email" className="capitalize">
              correo electrónico
            </Label>
            <Input id="email" name="email" type="email" placeholder="nicola.tesla@gmail.com" required />
          </div>

          {/* Password */}
          <div className="grid gap-2">
            <div className="flex items-center gap-4">
              <Label htmlFor="password" className="capitalize">
                contraseña
              </Label>
              <a href="/auth/forgot" className=" ml-auto text-sm underline-offset-4 hover:underline text-red-500">
                olvidaste tu contraseña?
              </a>
            </div>
            <div className="relative">
              <Input id="password" name="password" placeholder="********" type={passwordVisible ? 'text' : 'password'} required />
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

          {/* Remember me */}
          <div className="flex items-center justify-end">
            <input type="checkbox" id="remember" name="remember" className="mr-2" />
            <Label htmlFor="remember" className="m-0">
              Guardar
            </Label>
          </div>

          <Button variant="ghost" type="submit" className="w-full bg-red-600 font-bold text-white " disabled={loading}>
            {loading ? 'Logging in…' : 'Login'}
          </Button>
        </form>
      </CardContent>

      <CardFooter className="flex-col gap-8">
        <Button variant="outline" className="w-full " onClick={() => router.push('/api/auth/login/google')}>
          <GlobeLock className="mr-2 w-5 h-5" />
          Continuar con Google
        </Button>
        <Button variant="outline" className="w-full " onClick={() => router.push('/api/auth/login/facial')}>
          <ScanFace className="mr-2 w-5 h-5" />
          Autenticación Facial
        </Button>
      </CardFooter>
    </Card>
  );
};
