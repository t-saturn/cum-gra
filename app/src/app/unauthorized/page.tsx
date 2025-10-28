const UnauthorizedPage: React.FC = () => {
  return (
    <main className="flex justify-center items-center p-6 min-h-[60vh]">
      <div className="space-y-3 max-w-md text-center">
        <h1 className="font-semibold text-2xl">No autorizado</h1>
        <p className="text-zinc-600">Tu cuenta no tiene un rol asignado para esta aplicación.</p>
      </div>
    </main>
  );
};

export default UnauthorizedPage;
