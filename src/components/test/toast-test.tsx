import React from 'react';
import { toast } from 'sonner';
import { Button } from '../ui/button';

const ToastTest = () => {
  return (
    <div>
      <button
        className="bg-red-400 p-2 px-4 rounded-md cursor-pointer hover:bg-red-500"
        onClick={() =>
          toast.error('Event has been created', {
            description: 'Sunday, December 03, 2023 at 9:00 AM',
            position: 'top-right',
            duration: 500,
            action: {
              label: 'Undo',
              onClick: () => console.log('Undo'),
            },
          })
        }
      >
        Cerrar sesion
      </button>
      <div className="flex gap-4">
        <Button onClick={() => toast('Normal toast')}>Normal</Button>
        <Button
          onClick={() =>
            toast('Action toast', {
              action: { label: 'Undo', onClick: () => console.log('Undo clicked') },
            })
          }
        >
          Action
        </Button>
        <Button variant={'ghost'} onClick={() => toast.success('Success toast')}>
          Success
        </Button>
        <Button onClick={() => toast.info('Info toast')}>Info</Button>
        <Button onClick={() => toast.warning('Warning toast')}>Warning</Button>
        <Button onClick={() => toast.error('Error toast')}>Error</Button>
        <Button
          onClick={() =>
            toast.promise(new Promise((resolve) => setTimeout(resolve, 2000)), {
              loading: 'Loading...',
              success: 'Loaded!',
              error: 'Failed to load.',
            })
          }
        >
          Loading
        </Button>
        <Button onClick={() => toast('Default toast')}>Default</Button>
      </div>
    </div>
  );
};

export default ToastTest;
