// app/page.tsx
'use client';

import { NextPage } from 'next';
import { Building2 } from 'lucide-react';
import { Login } from '@/components/auth/login';

const Page: NextPage = () => {
  return (
    <div className="min-h-screen w-full flex">
      <div className="absolute top-8 left-8 z-10 flex items-center space-x-3">
        <div className="w-10 h-10 bg-gray-900 rounded-lg flex items-center justify-center">
          <Building2 className="w-6 h-6 text-white" />
        </div>
        <span className="text-2xl font-bold text-gray-900">Gobierno Regional de Ayacucho</span>
      </div>

      {/* Right Side - Form centrado */}
      <div className="flex-1 flex items-center justify-center p-4">
        <div className="w-full max-w-xl">
          <Login />
        </div>
      </div>
    </div>
  );
};

export default Page;
