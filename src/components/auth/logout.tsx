'use client';

import React from 'react';
import { Button } from '../ui/button';
import { toast } from 'sonner';
import { useRouter } from 'next/navigation';

export const Logout: React.FC = () => {
  const router = useRouter();

  async function handle_logout() {
    try {
      const res = await fetch('/api/auth/logout?logout_type=user_logout', {
        method: 'GET',
        credentials: 'include', // <-- send cookies
      });
      const json = await res.json();
      console.log('Logout response:', json);

      if (res.ok && json.success) {
        toast.success('Cierre de sesión exitoso');
        router.push('/auth/login');
      } else {
        toast.error(json.error?.details || json.message || 'Error al cerrar sesión');
      }
    } catch (err: any) {
      toast.error(err.message || 'Error en la petición de logout');
    }
  }

  return (
    <div>
      <Button variant="ghost" onClick={handle_logout} className="bg-[#d20f39] text-[#eff1f5] cursor-pointer hover:bg-[#e64553]">
        Cerrar sesión
      </Button>
    </div>
  );
};
