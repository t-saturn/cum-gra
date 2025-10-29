import { LogIn, LogOut } from 'lucide-react';
import React from 'react';
import { Button } from '../ui/button';
import { useRouter } from 'next/navigation';

export const Login = () => {
  const router = useRouter();

  const handleClick = () => {
    // Redirección controlada por Next.js (cliente)
    router.push(`/auth/login`);
  };

  return (
    <Button onClick={handleClick} className="flex items-center gap-2 bg-[#d20f39] hover:bg-red-400 text-white cursor-pointer">
      <LogIn size={18} />
      Ingresar
    </Button>
  );
};

export const Logout = () => {
  const handleClick = () => {
    // Delegamos totalmente al mini-backend
    window.location.href = `/api/auth/logout`;
  };

  return (
    <Button onClick={handleClick} className="flex items-center gap-2 bg-[#d20f39] hover:bg-red-400 text-white cursor-pointer">
      <LogOut size={18} />
      Cerrar Sesión
    </Button>
  );
};
