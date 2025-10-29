'use client';

import { NextPage } from 'next';
import Image from 'next/image';
import { RecoverAccount } from '@/components/auth/forgot';

const Page: NextPage = () => {
  return (
    <div className="flex w-full min-h-screen">
      <div className="top-8 left-8 z-10 absolute flex items-center space-x-3">
        <Image src="/img/logo.png" alt="logo" width={40} height={20} />
        <span className="font-bold text-[#d20f39] text-2xl">Gobierno Regional de Ayacucho</span>
      </div>

      {/* Right Side - Form centrado */}
      <div className="flex flex-1 justify-center items-center p-4">
        <div className="w-full max-w-xl">
          <RecoverAccount />
        </div>
      </div>
    </div>
  );
};

export default Page;
