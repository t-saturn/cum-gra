'use client';

import React, { useEffect, useMemo, useState } from 'react';
import { Eye, EyeOff, GlobeLock, ScanFace } from 'lucide-react';
import { useSearchParams } from 'next/navigation';
import { Card, CardHeader, CardTitle, CardDescription, CardContent, CardFooter } from '../ui/card';
import { Button } from '../ui/button';
import { Label } from '@radix-ui/react-label';
import { Input } from '../ui/input';
import { getDeviceInfo } from '@/helpers/device-info';
import { DeviceInfo } from '@/types';

const genFlowIdHex = (len = 32) => {
  const bytes = new Uint8Array(len / 2);
  crypto.getRandomValues(bytes);
  return Array.from(bytes, (b) => b.toString(16).padStart(2, '0')).join('');
};

const API_BASE = process.env.NEXT_PUBLIC_API_BASE ?? 'http://localhost:5555';
const FRONT_BASE = process.env.NEXT_PUBLIC_FRONT_BASE ?? 'http://localhost:3000';

const ERROR_MAP: Record<string, string> = {
  INVALID_CREDENTIALS: 'Credenciales inválidas',
  INVALID_REQUEST: 'Solicitud inválida',
  INVALID_CALLBACK_URL: 'URL de retorno no permitida',
  REDIS_UPSERT_FAILED: 'Error temporal, intente nuevamente',
  UNAUTHORIZED: 'No autorizado',
};

export const Login: React.FC = () => {
  const sp = useSearchParams();
  const [passwordVisible, setPasswordVisible] = useState(false);
  const [loading, setLoading] = useState(false);
  const [loginFlowId, setLoginFlowId] = useState<string>('');

  // error_code pasado por redirect del gateway
  const errorCode = sp.get('error_code') || '';

  async function handle_login(e: React.FormEvent<HTMLFormElement>): Promise<void> {
    e.preventDefault();
    setLoading(true);

    const form = e.currentTarget;

    try {
      // Obtener la info real del dispositivo
      const device_info: DeviceInfo = await getDeviceInfo();

      const payload = {
        email: (form.email as HTMLInputElement).value,
        password: (form.password as HTMLInputElement).value,
        login_flow_id: loginFlowId,
        application_id: 'app_001',
        callback_url: callbackAbs,
        device_info, // usamos la data real
      };

      const res = await fetch('/api/auth/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
        body: JSON.stringify(payload),
      });

      const json = await res.json().catch(() => ({}));

      if (res.status === 401) {
        const msg = json?.error?.details || ERROR_MAP[json?.error?.code] || 'Credenciales inválidas';
        const url = new URL(window.location.href);
        url.searchParams.set('error_code', json?.error?.code || 'UNAUTHORIZED');
        window.history.replaceState(null, '', url.toString());
        alert(msg);
        setLoading(false);
        return;
      }

      if (json?.redirect) {
        const finalURL = new URL(json.redirect, API_BASE).toString();
        localStorage.removeItem('login_flow_id');
        window.location.assign(finalURL);
        return;
      }

      alert('Respuesta inesperada del servidor.');
    } catch (err) {
      console.error(err);
      alert('No se pudo contactar al servidor.');
    } finally {
      setLoading(false);
    }
  }

  // callback absoluto (ej: http://localhost:3000/home o lo que venga en ?callback)
  const callbackAbs = useMemo(() => {
    // const cb = sp.get('callback') || `${FRONT_BASE}/home`;
    const cb = sp.get('callback') || `${FRONT_BASE}`;
    try {
      return new URL(cb, FRONT_BASE).toString();
    } catch {
      // return `${FRONT_BASE}/home`;
      return `${FRONT_BASE}`;
    }
  }, [sp]);

  useEffect(() => {
    let id = localStorage.getItem('login_flow_id');
    if (!id) {
      id = genFlowIdHex();
      localStorage.setItem('login_flow_id', id);
    }
    setLoginFlowId(id);
  }, []);

  const friendlyError = errorCode ? ERROR_MAP[errorCode] || 'Ocurrió un error' : '';

  return (
    <Card className="w-full max-w-md mx-auto">
      <CardHeader>
        <CardTitle>Iniciar sesión</CardTitle>
        <CardDescription>Ingrese su correo electrónico y contraseña</CardDescription>
      </CardHeader>

      <CardContent>
        <form onSubmit={handle_login} className="flex flex-col gap-6">
          {friendlyError && <p className="text-red-600 text-sm">{friendlyError}</p>}

          <div className="grid gap-2">
            <Label htmlFor="email">correo electrónico</Label>
            <Input id="email" name="email" type="email" placeholder="nicola.tesla@gmail.com" required />
          </div>

          <div className="grid gap-2">
            <div className="flex items-center gap-4">
              <Label htmlFor="password">contraseña</Label>
              <a href="/auth/forgot" className="ml-auto text-sm underline-offset-4 hover:underline text-red-500">
                ¿olvidaste tu contraseña?
              </a>
            </div>
            <div className="relative">
              <Input id="password" name="password" placeholder="********" type={passwordVisible ? 'text' : 'password'} required />
              <button
                type="button"
                onClick={() => setPasswordVisible((v) => !v)}
                className="absolute right-3 top-1/2 -translate-y-1/2 text-gray-400 focus:outline-none"
                aria-label={passwordVisible ? 'Ocultar contraseña' : 'Mostrar contraseña'}
                title={passwordVisible ? 'Ocultar contraseña' : 'Mostrar contraseña'}
              >
                {passwordVisible ? <EyeOff className="w-5 h-5" /> : <Eye className="w-5 h-5" />}
              </button>
            </div>
          </div>

          <Button variant="ghost" type="submit" className="w-full bg-red-600 font-bold text-white" disabled={loading}>
            {loading ? 'Logging in…' : 'Login'}
          </Button>
        </form>
      </CardContent>

      <CardFooter className="flex-col gap-8">
        <Button variant="outline" className="w-full " onClick={() => window.location.assign(`${API_BASE}/api/auth/login/google`)}>
          <GlobeLock className="mr-2 w-5 h-5" />
          Continuar con Google
        </Button>
        <Button variant="outline" className="w-full " onClick={() => window.location.assign(`${API_BASE}/api/auth/login/facial`)}>
          <ScanFace className="mr-2 w-5 h-5" />
          Autenticación Facial
        </Button>
      </CardFooter>
    </Card>
  );
};
