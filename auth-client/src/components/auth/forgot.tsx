'use client';

import React, { useState } from 'react';
import { Eye, EyeOff, Mail, CheckCircle2, Hash } from 'lucide-react';
import { useRouter } from 'next/navigation';
import { toast } from 'sonner';

import { Card, CardHeader, CardTitle, CardDescription, CardContent, CardFooter } from '../ui/card';
import { Button } from '../ui/button';
import { Label } from '@radix-ui/react-label';
import { Input } from '../ui/input';

const SEND_CODE_ENDPOINT = '/api/auth/recover/send-code';
const CONFIRM_ENDPOINT = '/api/auth/recover/confirm';

export const RecoverAccount: React.FC = () => {
  const router = useRouter();

  const [loadingSend, setLoadingSend] = useState(false);
  const [loadingConfirm, setLoadingConfirm] = useState(false);

  const [email, setEmail] = useState('');
  const [codeSent, setCodeSent] = useState(false);
  const [code, setCode] = useState('');

  const [passwordVisible1, setPasswordVisible1] = useState(false);
  const [passwordVisible2, setPasswordVisible2] = useState(false);
  const [newPass, setNewPass] = useState('');
  const [confirmPass, setConfirmPass] = useState('');

  async function handleSendCode(e: React.FormEvent) {
    e.preventDefault();
    if (!email) {
      toast.error('Ingresa tu correo electrónico', { position: 'top-right' });
      return;
    }

    try {
      setLoadingSend(true);
      const res = await fetch(SEND_CODE_ENDPOINT, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email }),
      });
      const json = await res.json();

      if (!res.ok) {
        toast.error(json.error?.details || json.message || 'No se pudo enviar el código', { position: 'top-right' });
        return;
      }

      toast.success('Código enviado a tu correo', { position: 'top-right' });
      setCodeSent(true);
    } catch {
      toast.error('Error enviando el código', { position: 'top-right' });
    } finally {
      setLoadingSend(false);
    }
  }

  async function handleConfirm(e: React.FormEvent) {
    e.preventDefault();

    if (!email || !code) {
      toast.error('Completa correo y código', { position: 'top-right' });
      return;
    }
    if (!newPass || !confirmPass) {
      toast.error('Ingresa la nueva contraseña y su confirmación', { position: 'top-right' });
      return;
    }
    if (newPass !== confirmPass) {
      toast.error('Las contraseñas no coinciden', { position: 'top-right' });
      return;
    }

    try {
      setLoadingConfirm(true);
      const res = await fetch(CONFIRM_ENDPOINT, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, code, new_password: newPass }),
      });
      const json = await res.json();

      if (!res.ok) {
        toast.error(json.error?.details || json.message || 'No se pudo cambiar la contraseña', { position: 'top-right' });
        return;
      }

      toast.success('Contraseña actualizada. Inicia sesión.', { position: 'top-right' });
      router.push('/auth/login');
    } catch {
      toast.error('Error al confirmar el cambio', { position: 'top-right' });
    } finally {
      setLoadingConfirm(false);
    }
  }

  const disabledConfirm = !codeSent || !email || !code || !newPass || !confirmPass || newPass !== confirmPass || loadingConfirm;

  return (
    <Card className="w-full max-w-md mx-auto">
      <CardHeader>
        <CardTitle>Recuperar cuenta</CardTitle>
        <CardDescription>Ingresa tu correo, recibe un código y establece una nueva contraseña.</CardDescription>
      </CardHeader>

      <CardContent>
        <form className="flex flex-col gap-6" onSubmit={codeSent ? handleConfirm : handleSendCode}>
          {/* Email */}
          <div className="grid gap-2">
            <Label htmlFor="email" className="capitalize">
              correo electrónico
            </Label>
            <div className="relative">
              <Input
                id="email"
                name="email"
                type="email"
                placeholder="nicola.tesla@gmail.com"
                required
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                disabled={loadingSend || loadingConfirm || codeSent}
              />
              <Mail className="w-4 h-4 absolute right-3 top-1/2 -translate-y-1/2 opacity-60" />
            </div>
          </div>

          {/* Send code */}
          <Button
            type="button"
            variant="ghost"
            className="w-full hover:bg-[#e64553] font-bold hover:text-[#eff1f5]  cursor-pointer"
            onClick={handleSendCode}
            disabled={loadingSend || !email || codeSent}
          >
            {loadingSend ? 'Enviando…' : codeSent ? 'Código enviado' : 'Enviar código'}
          </Button>

          {/* Code */}
          <div className="grid gap-2">
            <Label htmlFor="code" className="capitalize">
              código de verificación
            </Label>
            <div className="relative">
              <Input
                id="code"
                name="code"
                type="text"
                placeholder="Ingresa el código"
                required={codeSent}
                value={code}
                onChange={(e) => setCode(e.target.value)}
                disabled={!codeSent || loadingConfirm}
              />
              <Hash className="w-4 h-4 absolute right-3 top-1/2 -translate-y-1/2 opacity-60" />
            </div>
          </div>

          {/* New password */}
          <div className="grid gap-2">
            <Label htmlFor="newPass" className="capitalize">
              nueva contraseña
            </Label>
            <div className="relative">
              <Input
                id="newPass"
                name="newPass"
                placeholder="********"
                type={passwordVisible1 ? 'text' : 'password'}
                value={newPass}
                onChange={(e) => setNewPass(e.target.value)}
                disabled={!codeSent || loadingConfirm}
                required={codeSent}
              />
              <button
                type="button"
                onClick={() => setPasswordVisible1((v) => !v)}
                aria-label={passwordVisible1 ? 'Ocultar contraseña' : 'Mostrar contraseña'}
                title={passwordVisible1 ? 'Ocultar contraseña' : 'Mostrar contraseña'}
                className="absolute right-3 top-1/2 -translate-y-1/2 text-gray-400 focus:outline-none"
              >
                {passwordVisible1 ? <EyeOff className="w-5 h-5" /> : <Eye className="w-5 h-5" />}
              </button>
            </div>
          </div>

          {/* Confirm password */}
          <div className="grid gap-2">
            <Label htmlFor="confirmPass" className="capitalize">
              confirmar nueva contraseña
            </Label>
            <div className="relative">
              <Input
                id="confirmPass"
                name="confirmPass"
                placeholder="********"
                type={passwordVisible2 ? 'text' : 'password'}
                value={confirmPass}
                onChange={(e) => setConfirmPass(e.target.value)}
                disabled={!codeSent || loadingConfirm}
                required={codeSent}
              />
              <button
                type="button"
                onClick={() => setPasswordVisible2((v) => !v)}
                aria-label={passwordVisible2 ? 'Ocultar contraseña' : 'Mostrar contraseña'}
                title={passwordVisible2 ? 'Ocultar contraseña' : 'Mostrar contraseña'}
                className="absolute right-3 top-1/2 -translate-y-1/2 text-gray-400 focus:outline-none"
              >
                {passwordVisible2 ? <EyeOff className="w-5 h-5" /> : <Eye className="w-5 h-5" />}
              </button>
            </div>
            {confirmPass && newPass !== confirmPass && <p className="text-sm text-red-500">Las contraseñas no coinciden</p>}
          </div>

          {/* Actions */}
          <div className="flex flex-col gap-3">
            <Button
              type={codeSent ? 'submit' : 'button'}
              className="w-full font-bold bg-[#d20f39] hover:bg-[#e64553]"
              onClick={codeSent ? undefined : handleSendCode}
              disabled={disabledConfirm && codeSent}
            >
              {codeSent ? (
                <>
                  <CheckCircle2 className="w-5 h-5 mr-2" /> Confirmar
                </>
              ) : (
                <span>Enviar código</span>
              )}
            </Button>

            <Button type="button" variant="outline" className="w-full" onClick={() => router.push('/auth/login')}>
              Cancelar
            </Button>
          </div>
        </form>
      </CardContent>

      <CardFooter className="justify-center">
        <Button variant="link" onClick={() => router.push('/auth/login')} className="cursor-pointer">
          Volver a iniciar sesión
        </Button>
      </CardFooter>
    </Card>
  );
};
