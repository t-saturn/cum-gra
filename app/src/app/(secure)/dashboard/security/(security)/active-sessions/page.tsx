import { Suspense } from 'react';
import { Skeleton } from '@/components/ui/skeleton';
import ActiveSessionsContent from './active-sessions-content';

function ActiveSessionsLoadingSkeleton() {
  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <div className="space-y-2">
          <Skeleton className="h-8 w-64" />
          <Skeleton className="h-4 w-96" />
        </div>
        <div className="flex gap-3">
          <Skeleton className="h-10 w-32" />
        </div>
      </div>
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
        {[1, 2, 3, 4].map((i) => (
          <Skeleton key={i} className="h-32" />
        ))}
      </div>
      <Skeleton className="h-96 w-full" />
    </div>
  );
}

export default function ActiveSessionsPage() {
  return (
    <Suspense fallback={<ActiveSessionsLoadingSkeleton />}>
      <ActiveSessionsContent />
    </Suspense>
  );
}