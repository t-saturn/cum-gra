import { Fullscreen } from 'lucide-react';

import { Button } from '@/components/ui/button';
import { SidebarTrigger } from '@/components/ui/sidebar';
import { ThemeToggle } from '../theme/theme-toggle';
import { UserPopover } from './user-popover';
import { AppsPopover } from './apps-popover';

function Navbar() {

  const handleToggleFullscreen = () => {
    if (!document.fullscreenElement)
      document.documentElement.requestFullscreen().catch((err) => {
        console.error('Error al intentar poner en fullscreen:', err);
      });
    else
      document.exitFullscreen().catch((err) => {
        console.error('Error al intentar salir de fullscreen:', err);
      });
  };

  return (
    <header className="sticky top-0 z-10 flex h-16 items-center gap-4 border-b bg-card/70 p-6 rounded-t-lg backdrop-blur-sm">
      <SidebarTrigger className="hover:cursor-pointer" />

      <div className="ml-auto flex items-center gap-4 text-muted-foreground">
        <ThemeToggle />
        <Button variant="ghost" size="icon" className="hover:cursor-pointer" onClick={handleToggleFullscreen}>
          <Fullscreen className="h-5 w-5" />
          <span className="sr-only">Fullscreen</span>
        </Button>
        <AppsPopover />
        <UserPopover />
      </div>
    </header>
  );
}

export default Navbar;