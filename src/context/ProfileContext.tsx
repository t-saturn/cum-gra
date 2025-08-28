'use client';

import { useSession } from '@/hooks/useSession';
import React, { createContext, useContext, useState, useEffect, ReactNode } from 'react';

interface ProfileSchema {
  name: string;
  role: string;
  avatar: string;
}

interface ProfileContextType {
  profile: ProfileSchema;
  fetchProfile: (userId: string) => Promise<void>;
}

const ProfileContext = createContext<ProfileContextType | undefined>(undefined);

export const ProfileProvider = ({ children }: { children: ReactNode }) => {
  const { role, userId, name } = useSession();

  const [profile, setProfile] = useState<ProfileSchema>({
    name: '',
    role: '',
    avatar: '',
  });

  const fetchProfile = async (userId: string) => {
    try {
      // aquí luego podrías hacer un fetch real a tu API
      const result = {
        userId,
        name: name ?? 'Usuario',
        role: role ?? 'invitado',
        avatar: 'https://i.pravatar.cc/150?img=7',
      };

      if (result) setProfile(result);
      else console.error('Error al traer el perfil:');
    } catch (error) {
      console.error('Error al traer el perfil: ' + error);
    }
  };

  useEffect(() => {
    if (!userId) return; // espera a tener un userId válido
    fetchProfile(userId);
  }, [userId]);

  return <ProfileContext.Provider value={{ profile, fetchProfile }}>{children}</ProfileContext.Provider>;
};

export const useProfile = () => {
  const context = useContext(ProfileContext);
  if (!context) throw new Error('useProfile debe usarse dentro de un ProfileProvider');
  return context;
};
