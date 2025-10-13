import { ReactNode } from 'react';

interface InfoHeaderProps {
  title: string;
  description: string;
  children?: ReactNode;
}

const InfoHeader = ({ title, description, children }: InfoHeaderProps) => {
  return (
    <div className="flex sm:flex-row flex-col justify-between items-center gap-4">
      <div>
        <h1 className="bg-clip-text bg-gradient-to-r from-foreground to-muted-foreground font-bold text-transparent text-3xl">{title}</h1>
        <p className="mt-2 text-muted-foreground text-lg">{description}</p>
      </div>
      {children}
    </div>
  );
};

export default InfoHeader;
