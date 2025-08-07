import React from 'react';
import { Button } from '../ui/button';
import { toast } from 'sonner';

export const Logout = () => {
  const handle_logout = () => {
    // Logic for handling logout
    console.log('Logging out...');
    toast.success('You have been logged out successfully', {
      position: 'top-right',
      duration: 3000,
    });
  };

  return (
    <div>
      <Button variant="ghost" onClick={handle_logout} className="bg-[#d20f39] text-[#eff1f5] cursor-pointer hover:bg-[#e64553]">
        Cerrar sesión
      </Button>
    </div>
  );
};
