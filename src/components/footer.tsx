import { Mail, Phone, MapPin, ExternalLink, Shield, FileText } from 'lucide-react';
import { Button } from '@/components/ui/button';
import { Separator } from '@/components/ui/separator';
import Image from 'next/image';

const quickLinks = [
  { label: 'Portal Principal', href: '#', icon: ExternalLink },
  { label: 'Términos de Uso', href: '#', icon: FileText },
  { label: 'Política de Privacidad', href: '#', icon: Shield },
  { label: 'Contacto Técnico', href: '#', icon: Mail },
];

const contactInfo = [
  { label: 'Dirección', value: 'Av. Principal 123, Ciudad', icon: MapPin },
  { label: 'Teléfono', value: '+1 (555) 123-4567', icon: Phone },
  { label: 'Email', value: 'soporte@entidadpublica.gov', icon: Mail },
];

const socialLinks = [
  { name: 'Web', href: 'https://www.gob.pe/regionayacucho' },
  { name: 'Facebook', href: 'https://www.facebook.com/gobiernoregionalayacucho/' },
];

export const Footer: React.FC = () => {
  return (
    <footer className="bg-background mt-auto border-t border-border">
      <div className="mx-auto px-4 sm:px-6 lg:px-8 max-w-7xl">
        <div className="py-8">
          <div className="gap-8 grid grid-cols-1 md:grid-cols-4">
            <div className="md:col-span-2">
              <div className="flex items-center space-x-3 mb-4">
                <div className="flex justify-center items-center w-10 h-10">
                  <Image src="/img/logo.png" alt="Logo Gobierno Regional de Ayacucho" width={24} height={24} className="rounded-full" priority />
                </div>
                <div>
                  <h3 className="font-semibold text-foreground">Gobierno Regional de Ayacucho</h3>
                  <p className="text-muted-foreground text-sm">Sistema de Gestión Interno</p>
                </div>
              </div>
              <p className="mb-4 max-w-md text-muted-foreground text-sm">
                Plataforma integral para la gestión de procesos administrativos y operativos de la Gobierno Regional de Ayacucho, diseñada para optimizar la eficiencia
                institucional.
              </p>

              {/* Enlaces rápidos */}
              <div className="flex flex-wrap gap-2">
                {quickLinks.map((link) => {
                  const IconComponent = link.icon;
                  return (
                    <Button key={link.label} variant="ghost" size="sm" className="px-2 h-8 text-xs" asChild>
                      <a href={link.href} className="flex items-center space-x-1">
                        <IconComponent className="w-3 h-3" />
                        <span>{link.label}</span>
                      </a>
                    </Button>
                  );
                })}
              </div>
            </div>

            <div>
              <h4 className="mb-3 font-medium text-foreground">Información de Contacto</h4>
              <div className="space-y-2">
                {contactInfo.map((contact) => {
                  const IconComponent = contact.icon;
                  return (
                    <div key={contact.label} className="flex items-start space-x-2">
                      <IconComponent className="flex-shrink-0 mt-0.5 w-4 h-4 text-muted-foreground" />
                      <div>
                        <p className="text-muted-foreground text-xs">{contact.label}</p>
                        <p className="text-foreground text-sm">{contact.value}</p>
                      </div>
                    </div>
                  );
                })}
              </div>
            </div>

            <div>
              <h4 className="mb-3 font-medium text-foreground">Estado del Sistema</h4>
              <div className="space-y-3">
                <div className="flex justify-between items-center">
                  <span className="text-muted-foreground text-sm">Servidor</span>
                  <div className="flex items-center space-x-2">
                    <div className="bg-green-500 rounded-full w-2 h-2"></div>
                    <span className="text-foreground text-xs">Online</span>
                  </div>
                </div>
                <div className="flex justify-between items-center">
                  <span className="text-muted-foreground text-sm">Base de Datos</span>
                  <div className="flex items-center space-x-2">
                    <div className="bg-green-500 rounded-full w-2 h-2"></div>
                    <span className="text-foreground text-xs">Conectada</span>
                  </div>
                </div>
                <div className="flex justify-between items-center">
                  <span className="text-muted-foreground text-sm">Última actualización</span>
                  <span className="text-foreground text-xs">Hace 2 min</span>
                </div>
              </div>
            </div>
          </div>
        </div>

        <Separator />

        <div className="py-4">
          <div className="flex sm:flex-row flex-col justify-between items-center space-y-2 sm:space-y-0">
            <div className="flex items-center space-x-4 text-muted-foreground text-sm">
              <span>© 2025 Gobierno Regional de Ayacucho - OTIC</span>
              <span>•</span>
              <span>Todos los derechos reservados</span>
              <span>•</span>
              <span>Versión 1.0.0</span>
            </div>

            <div className="flex items-center space-x-4">
              <span className="text-muted-foreground text-xs">Síguenos:</span>
              {socialLinks.map((social) => (
                <Button key={social.name} variant="ghost" size="sm" className="p-0 w-6 h-6 text-muted-foreground hover:text-foreground" asChild>
                  <a href={social.href} target="_blank" rel="noopener noreferrer">
                    {social.name.charAt(0)}
                  </a>
                </Button>
              ))}
            </div>
          </div>
        </div>
      </div>
    </footer>
  );
};
