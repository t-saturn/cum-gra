'use client';

import { NextPage } from 'next';
import Image from 'next/image';
import { RecoverAccount } from '@/components/auth/forgot';

const Page: NextPage = () => {
  return (
    <div className="min-h-screen w-full flex">
      <div className="absolute top-8 left-8 z-10 flex items-center space-x-3">
        <Image src="/img/logo.png" alt="logo" width={40} height={20} />
        <span className="text-2xl font-bold text-[#d20f39]">Gobierno Regional de Ayacucho</span>
      </div>

      {/* Right Side - Form centrado */}
      <div className="flex-1 flex items-center justify-center p-4">
        <div className="w-full max-w-xl">
          <RecoverAccount />
        </div>
      </div>
    </div>
  );
};

export default Page;
