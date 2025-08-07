import React from 'react';
import { Button } from '../ui/button';
import { toast } from 'sonner';
import { useRouter } from 'next/navigation';

export const Logout = () => {
  const router = useRouter();
  async function handle_logout() {
    const res = await fetch('/api/auth/logout', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        session_id: '82ef5a24-ede9-4a9c-8441-048200f8a50a', // replace with actual session ID logic
        logout_type: 'user_logout',
      }),
    });
    const json = await res.json();
    if (res.ok && json.success) {
      // redirect to login
      router.push('/login');
    } else {
      toast.error(json.error?.details || json.message);
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
