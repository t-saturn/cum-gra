'use client';

import { NextPage } from 'next';
import { useState } from 'react';
import { Building2, Eye, EyeOff, ScanFace, GlobeLock } from 'lucide-react';
import { useRouter } from 'next/navigation';
import { loginAction } from '@/actions/auth-login';

const Login: NextPage = () => {
  const router = useRouter();
  const [passwordVisible, setPasswordVisible] = useState(false);

  return (
    <>
      <div className="min-h-screen bg-white flex flex-col lg:flex-row">
        {/* Left Side - Illustration */}
        <div className="hidden lg:flex w-full h-screen bg-gradient-to-br from-red-50 via-pink-50 to-red-100 relative overflow-hidden">
          {/* Decorative circles */}
          <div className="absolute top-20 left-20 w-4 h-4 bg-[#D90D1E] rounded-full animate-pulse" />
          <div className="absolute top-40 right-32 w-3 h-3 bg-[#B50A17] rounded-full animate-bounce" />
          <div className="absolute bottom-32 left-16 w-5 h-5 bg-[#D90D1E] rounded-full animate-pulse" />
          <div className="absolute bottom-20 right-20 w-4 h-4 bg-[#B50A17] rounded-full animate-bounce" />
          <div className="absolute top-60 left-1/3 w-3 h-3 bg-[#D90D1E] rounded-full animate-pulse" />

          {/* Abstract shapes */}
          <div className="absolute inset-0">
            <svg className="absolute top-10 left-10 w-32 h-32 text-red-200 opacity-60" viewBox="0 0 100 100">
              <path d="M20,20 Q50,5 80,20 Q95,50 80,80 Q50,95 20,80 Q5,50 20,20" fill="currentColor" />
            </svg>
            <svg className="absolute bottom-20 right-10 w-40 h-40 text-red-200 opacity-50" viewBox="0 0 100 100">
              <circle cx="50" cy="50" r="40" fill="currentColor" />
            </svg>
            <svg className="absolute top-1/2 left-1/4 w-24 h-24 text-red-300 opacity-40" viewBox="0 0 100 100">
              <polygon points="50,10 90,90 10,90" fill="currentColor" />
            </svg>
          </div>

          {/* World map lines */}
          <div className="absolute inset-0 flex items-center justify-center">
            <div className="relative w-96 h-64 opacity-30">
              <svg viewBox="0 0 400 250" className="w-full h-full text-red-300">
                <path d="M50,80 Q80,70 120,85 L140,90 Q160,85 180,90 L200,95 Q220,90 250,95 L280,100 Q300,95 330,100" stroke="currentColor" strokeWidth={2} fill="none" />
                <path d="M60,120 Q90,110 130,125 L150,130 Q170,125 190,130 L210,135 Q230,130 260,135" stroke="currentColor" strokeWidth={2} fill="none" />
                <circle cx="100" cy="90" r="3" fill="currentColor" />
                <circle cx="200" cy="100" r="3" fill="currentColor" />
                <circle cx="280" cy="110" r="3" fill="currentColor" />
              </svg>
            </div>
          </div>

          {/* Ship */}
          <div className="absolute bottom-1/4 left-1/4">
            <svg width="80" height="40" viewBox="0 0 80 40" className="text-gray-700">
              <path d="M10,30 L70,30 L65,25 L60,20 L50,15 L30,15 L20,20 L15,25 Z" fill="currentColor" />
              <rect x="25" y="10" width="3" height="15" fill="currentColor" />
              <rect x="35" y="8" width="3" height="17" fill="currentColor" />
              <rect x="45" y="12" width="3" height="13" fill="currentColor" />
              <path d="M25,10 L35,10 L30,5 Z" fill="#D90D1E" />
              <path d="M35,8 L45,8 L40,3 Z" fill="#D90D1E" />
            </svg>
          </div>

          {/* Logo */}
          <div className="absolute top-8 left-8 z-10 flex items-center space-x-3">
            <div className="w-10 h-10 bg-gray-900 rounded-lg flex items-center justify-center">
              <Building2 className="w-6 h-6 text-white" />
            </div>
            <span className="text-2xl font-bold text-gray-900">Gobierno Regional de Ayacucho</span>
          </div>
        </div>

        {/* Right Side - Form */}
        <div className="flex-1 h-screen flex items-center justify-center p-8 lg:p-12 lg:absolute lg:right-20 bg-transparent">
          <div className="w-full max-w-md">
            {/* Mobile Logo */}
            <div className="mb-8 lg:hidden text-center">
              <div className="flex items-center justify-center space-x-3 mb-4">
                <div className="w-10 h-10 bg-gray-900 rounded-lg flex items-center justify-center">
                  <Building2 className="w-6 h-6 text-white" />
                </div>
                <span className="text-xl font-bold text-gray-900">Gobierno Regional de Ayacucho</span>
              </div>
            </div>

            {/* Card */}
            <div className="bg-white shadow-2xl rounded-2xl overflow-hidden animate-fade-in">
              <div className="p-8 lg:p-10 space-y-8">
                <h1 className="text-3xl font-bold text-gray-900">Iniciar sesión</h1>

                <form
                  className="space-y-6"
                  method="null"
                  action={async (formData: FormData) => {
                    try {
                      await loginAction(formData);
                      router.push('/');
                    } catch (err: any) {
                      console.error(err);
                    }
                  }}
                >
                  {/* Email */}
                  <div className="space-y-2">
                    <label htmlFor="email" className="text-sm font-medium text-gray-700">
                      Email o Usuario
                    </label>
                    <input
                      id="email"
                      name="email"
                      type="email"
                      placeholder="Dirección de email"
                      required
                      className="h-12 w-full px-3 text-gray-600 border border-gray-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-[#D90D1E] focus:border-[#D90D1E]"
                    />
                  </div>

                  {/* Password */}
                  <div className="space-y-2">
                    <div className="flex items-center justify-between">
                      <label htmlFor="password" className="text-sm font-medium text-gray-700">
                        Contraseña
                      </label>
                      <a href="#" className="text-sm text-[#D90D1E] hover:text-[#B50A17]">
                        ¿Olvidaste tu contraseña?
                      </a>
                    </div>
                    <div className="relative">
                      <input
                        id="password"
                        name="password"
                        type={passwordVisible ? 'text' : 'password'}
                        placeholder="Contraseña"
                        required
                        className="h-12 w-full text-gray-600 px-3 pr-10 border border-gray-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-[#D90D1E] focus:border-[#D90D1E]"
                      />
                      <button
                        type="button"
                        onClick={() => setPasswordVisible(!passwordVisible)}
                        aria-label={passwordVisible ? 'Ocultar contraseña' : 'Mostrar contraseña'}
                        title={passwordVisible ? 'Ocultar contraseña' : 'Mostrar contraseña'}
                        className="absolute right-3 top-1/2 -translate-y-1/2 text-gray-400 focus:outline-none"
                      >
                        {passwordVisible ? <EyeOff className="w-5 h-5" /> : <Eye className="w-5 h-5" />}
                      </button>
                    </div>
                  </div>

                  {/* Submit */}
                  <button type="submit" className="w-full h-12 bg-[#D90D1E] hover:bg-[#B50A17] text-white font-semibold rounded-lg transition-colors duration-200">
                    Iniciar sesión
                  </button>

                  {/* Divider */}
                  <div className="relative">
                    <div className="absolute inset-0 flex items-center">
                      <span className="w-full border-t border-gray-200" />
                    </div>
                    <div className="relative flex justify-center text-sm">
                      <span className="bg-white px-4 text-gray-500 font-medium">O</span>
                    </div>
                  </div>

                  {/* Alternative Methods */}
                  <div className="space-y-3">
                    <button
                      type="button"
                      className="w-full h-12 border border-gray-200 hover:border-gray-300 hover:bg-gray-50 rounded-lg transition-colors duration-200 flex items-center justify-center"
                    >
                      <GlobeLock className="w-5 h-5 text-[#4285F4]" />
                      <span className="ml-3 text-gray-700 font-medium">Iniciar sesión con Google</span>
                    </button>
                    <button
                      type="button"
                      className="w-full h-12 border border-gray-200 hover:border-gray-300 hover:bg-gray-50 rounded-lg transition-colors duration-200 flex items-center justify-center"
                    >
                      <ScanFace className="w-5 h-5 text-[#D90D1E]" />
                      <span className="ml-3 text-gray-700 font-medium">Autenticación Facial</span>
                    </button>
                  </div>
                </form>
              </div>
            </div>
          </div>
        </div>
      </div>

      <style jsx>{`
        @keyframes fade-in {
          from {
            opacity: 0;
            transform: translateY(20px);
          }
          to {
            opacity: 1;
            transform: translateY(0);
          }
        }
        .animate-fade-in {
          animation: fade-in 0.6s ease-out;
        }
      `}</style>
    </>
  );
};

export default Login;
