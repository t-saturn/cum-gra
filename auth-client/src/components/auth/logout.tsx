'use client';

import React, { useState } from 'react';
import { Button } from '../ui/button';
import { toast } from 'sonner';
import { useRouter } from 'next/navigation';

export const Logout: React.FC = () => {
  const router = useRouter();
  const [loading, setLoading] = useState(false);

  async function handle_logout() {
    setLoading(true);
    try {
      const res = await fetch('/api/auth/logout?logout_type=user_logout', {
        method: 'GET',
        credentials: 'include', // send cookies
      });
      const json = await res.json();

      if (!res.ok) {
        toast.error(json.error?.details || json.message || 'Error al cerrar sesión', {
          position: 'top-right',
        });
        return;
      }

      toast.success('Cierre de sesión exitoso', { position: 'top-right' });
      router.push('/auth/login');
    } catch (err: any) {
      toast.error(err.message || 'Error en la petición de logout', { position: 'top-right' });
    } finally {
      setLoading(false);
    }
  }

  return (
    <div>
      <Button variant="ghost" disabled={loading} onClick={handle_logout} className="bg-[#d20f39] text-[#eff1f5] cursor-pointer hover:bg-[#e64553]">
        {loading ? 'Cerrando...' : 'Cerrar sesión'}
      </Button>
    </div>
  );
};
