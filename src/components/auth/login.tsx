import { component$, useStyles$, useSignal } from "@builder.io/qwik";
import { Building2, Chrome, Eye, EyeOff, Scan } from "lucide-react";

// Helper component to wrap lucide-react icons for JSX compatibility
const IconWrapper = component$<{ Icon: any; class?: string }>(
  ({ Icon, class: className }) => {
    return <Icon class={className} />;
  },
);

export default component$(() => {
  useStyles$(`
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
  `);

  const showPassword = useSignal(false);

  return (
    <div class="flex flex-col lg:flex-row h-screen">
      {/* Left Side - Illustration */}
      <div class="hidden lg:flex w-full h-full bg-gradient-to-br from-red-50 via-pink-50 to-red-100 relative overflow-hidden">
        {/* Background Elements */}
        <div class="absolute inset-0 w-full h-full">
          {/* Decorative circles */}
          <div class="absolute top-20 left-20 w-4 h-4 bg-[#D90D1E] rounded-full animate-pulse"></div>
          <div class="absolute top-40 right-32 w-3 h-3 bg-[#B50A17] rounded-full animate-bounce"></div>
          <div class="absolute bottom-32 left-16 w-5 h-5 bg-[#D90D1E] rounded-full animate-pulse"></div>
          <div class="absolute bottom-20 right-20 w-4 h-4 bg-[#B50A17] rounded-full animate-bounce"></div>
          <div class="absolute top-60 left-1/3 w-3 h-3 bg-[#D90D1E] rounded-full animate-pulse"></div>

          {/* Abstract shapes */}
          <div class="absolute top-0 left-0 w-full h-full">
            <svg
              class="absolute top-10 left-10 w-32 h-32 text-red-200 opacity-60"
              viewBox="0 0 100 100"
            >
              <path
                d="M20,20 Q50,5 80,20 Q95,50 80,80 Q50,95 20,80 Q5,50 20,20"
                fill="currentColor"
              />
            </svg>
            <svg
              class="absolute bottom-20 right-10 w-40 h-40 text-red-200 opacity-50"
              viewBox="0 0 100 100"
            >
              <circle cx="50" cy="50" r="40" fill="currentColor" />
            </svg>
            <svg
              class="absolute top-1/2 left-1/4 w-24 h-24 text-red-300 opacity-40"
              viewBox="0 0 100 100"
            >
              <polygon points="50,10 90,90 10,90" fill="currentColor" />
            </svg>
          </div>

          {/* World map illustration */}
          <div class="absolute inset-0 flex items-center justify-center">
            <div class="relative w-96 h-64 opacity-30">
              <svg viewBox="0 0 400 250" class="w-full h-full text-red-300">
                <path
                  d="M50,80 Q80,70 120,85 L140,90 Q160,85 180,90 L200,95 Q220,90 250,95 L280,100 Q300,95 330,100"
                  stroke="currentColor"
                  stroke-width="2"
                  fill="none"
                />
                <path
                  d="M60,120 Q90,110 130,125 L150,130 Q170,125 190,130 L210,135 Q230,130 260,135"
                  stroke="currentColor"
                  stroke-width="2"
                  fill="none"
                />
                <circle cx="100" cy="90" r="3" fill="currentColor" />
                <circle cx="200" cy="100" r="3" fill="currentColor" />
                <circle cx="280" cy="110" r="3" fill="currentColor" />
              </svg>
            </div>
          </div>

          {/* Ship illustration */}
          <div class="absolute bottom-1/4 left-1/4">
            <svg
              width="80"
              height="40"
              viewBox="0 0 80 40"
              class="text-gray-700"
            >
              <path
                d="M10,30 L70,30 L65,25 L60,20 L50,15 L30,15 L20,20 L15,25 Z"
                fill="currentColor"
              />
              <rect x="25" y="10" width="3" height="15" fill="currentColor" />
              <rect x="35" y="8" width="3" height="17" fill="currentColor" />
              <rect x="45" y="12" width="3" height="13" fill="currentColor" />
              <path d="M25,10 L35,10 L30,5 Z" fill="#D90D1E" />
              <path d="M35,8 L45,8 L40,3 Z" fill="#D90D1E" />
            </svg>
          </div>
        </div>

        {/* Logo */}
        <div class="absolute top-8 left-8 z-10">
          <div class="flex items-center space-x-3">
            <div class="w-10 h-10 bg-gray-900 rounded-lg flex items-center justify-center">
              <IconWrapper Icon={Building2} class="w-6 h-6 text-white" />
            </div>
            <span class="text-2xl font-bold text-gray-900">
              Gobierno Regional de Ayacucho
            </span>
          </div>
        </div>
      </div>

      {/* Right Side - Login Form */}
      <div class="lg:absolute lg:right-20 flex-1 h-full flex items-center justify-center p-8 lg:p-12 bg-transparent">
        <div class="w-full max-w-md bg-transparent">
          {/* Mobile Header */}
          <div class="text-center mb-8 lg:hidden">
            <div class="flex items-center justify-center space-x-3 mb-4">
              <div class="w-10 h-10 bg-gray-900 rounded-lg flex items-center justify-center">
                <IconWrapper Icon={Building2} class="w-6 h-6 text-white" />
              </div>
              <span class="text-xl font-bold text-gray-900">
                Gobierno Regional de Ayacucho
              </span>
            </div>
          </div>

          {/* Floating Login Card */}
          <div class="bg-white shadow-2xl border-0 rounded-2xl overflow-hidden animate-fade-in">
            <div class="p-8 lg:p-10">
              <div class="space-y-8">
                <div class="space-y-2">
                  <h1 class="text-3xl font-bold text-gray-900">
                    Iniciar sesión
                  </h1>
                </div>

                <form class="space-y-6">
                  {/* Email Field */}
                  <div class="space-y-2">
                    <label
                      for="email"
                      class="text-sm font-medium text-gray-700"
                    >
                      Email o Usuario
                    </label>
                    <input
                      id="email"
                      type="email"
                      placeholder="Dirección de email"
                      class="h-12 w-full px-3 border border-gray-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-[#D90D1E] focus:border-[#D90D1E]"
                      required
                    />
                  </div>

                  {/* Password Field */}
                  <div class="space-y-2">
                    <div class="flex items-center justify-between">
                      <label
                        for="password"
                        class="text-sm font-medium text-gray-700"
                      >
                        Contraseña
                      </label>
                      <a
                        href="#"
                        class="text-sm text-[#D90D1E] hover:text-[#B50A17]"
                      >
                        ¿Olvidaste tu contraseña?
                      </a>
                    </div>
                    <div class="relative">
                      <input
                        id="password"
                        type={showPassword.value ? "text" : "password"}
                        placeholder="Contraseña"
                        class="h-12 w-full px-3 pr-10 border border-gray-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-[#D90D1E] focus:border-[#D90D1E]"
                        required
                      />
                      <button
                        type="button"
                        class="absolute right-3 top-1/2 transform -translate-y-1/2 w-5 h-5 text-gray-400"
                        onClick$={() =>
                          (showPassword.value = !showPassword.value)
                        }
                      >
                        <IconWrapper
                          Icon={showPassword.value ? Eye : EyeOff}
                          class="w-5 h-5 text-gray-400"
                        />
                      </button>
                    </div>
                  </div>

                  {/* Sign In Button */}
                  <button
                    type="submit"
                    class="w-full h-12 bg-[#D90D1E] hover:bg-[#B50A17] text-white font-semibold rounded-lg transition-colors duration-200"
                  >
                    Iniciar sesión
                  </button>

                  {/* Divider */}
                  <div class="relative">
                    <div class="absolute inset-0 flex items-center">
                      <span class="w-full border-t border-gray-200"></span>
                    </div>
                    <div class="relative flex justify-center text-sm">
                      <span class="bg-white px-4 text-gray-500 font-medium">
                        O
                      </span>
                    </div>
                  </div>

                  {/* Alternative Login Methods */}
                  <div class="space-y-3">
                    {/* Google Login */}
                    <button
                      type="button"
                      class="w-full h-12 border border-gray-200 hover:border-gray-300 hover:bg-gray-50 rounded-lg transition-colors duration-200 bg-transparent flex items-center justify-center"
                    >
                      <div class="flex items-center space-x-3">
                        <IconWrapper
                          Icon={Chrome}
                          class="w-5 h-5 text-[#4285F4]"
                        />
                        <span class="text-gray-700 font-medium">
                          Iniciar sesión con Google
                        </span>
                      </div>
                    </button>

                    {/* Face Authentication */}
                    <button
                      type="button"
                      class="w-full h-12 border border-gray-200 hover:border-gray-300 hover:bg-gray-50 rounded-lg transition-colors duration-200 bg-transparent flex items-center justify-center"
                    >
                      <div class="flex items-center space-x-3">
                        <IconWrapper
                          Icon={Scan}
                          class="w-5 h-5 text-[#D90D1E]"
                        />
                        <span class="text-gray-700 font-medium">
                          Autenticación Facial
                        </span>
                      </div>
                    </button>
                  </div>
                </form>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
});
