import { ThemeToggle } from '@/components/theme/theme-toggle';
import Image from 'next/image';

const Home = () => {
  return (
    <div className="justify-items-center items-center gap-16 grid grid-rows-[20px_1fr_20px] p-8 sm:p-20 pb-20 min-h-screen font-[family-name:var(--font-geist-sans)]">
      <header className="flex flex-row justify-between items-center gap-2 w-full">
        <div className="flex flex-row items-center gap-2">
          <Image src="/img/logo.png" alt="Logo Gobierno Regional de Ayacucho" width={32} height={32} className="rounded-full" priority />
          <h1 className="font-black text-red-500 dark:text-red-400 text-3xl">Gobierno Regional de Ayacucho</h1>
        </div>
        <ThemeToggle />
      </header>

      <main className="flex flex-col items-center sm:items-start gap-[32px] row-start-2">conten-page</main>

      <footer className="flex flex-wrap justify-center items-center gap-[24px] row-start-3"></footer>
    </div>
  );
};

export default Home;
