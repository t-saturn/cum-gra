import { ReactNode } from "react";

interface InfoHeaderProps {
  title: string;
  description: string;
  children?: ReactNode;
}

const InfoHeader = ({ title, description, children }: InfoHeaderProps) => {
  return (
    <div className="flex items-center justify-between flex-col sm:flex-row gap-4">
      <div>
        <h1 className="text-3xl font-bold bg-gradient-to-r from-foreground to-muted-foreground bg-clip-text text-transparent">
          {title}
        </h1>
        <p className="text-muted-foreground mt-2 text-lg">
          {description}
        </p>
      </div>
      {children}
    </div>
  );
};

export default InfoHeader;
