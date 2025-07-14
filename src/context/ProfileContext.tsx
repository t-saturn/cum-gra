'use client';

import React from 'react';

import { createContext, useContext, useState, useEffect, ReactNode } from 'react';

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
  // mook simulated

  const { data: session } = {
    data: {
      user: {
        id: '1',
      },
    },
  };

  const [profile, setProfile] = useState<ProfileSchema>({ name: '', role: '', avatar: '' });

  const fetchProfile = async (userId: string) => {
    try {
      const result = {
        userId,
        name: 'Miguel Ramirez',
        role: 'admin',
        avatar: 'https://i.pravatar.cc/150?img=7',
      };
      if (result) setProfile(result);
      else console.error('Error al traer el perfil:');
    } catch (error) {
      console.error('Error al traer el perfil: ' + error);
    }
  };

  useEffect(() => {
    if (!session?.user?.id) return;
    fetchProfile(session.user.id);
  }, [session?.user?.id]);

  return <ProfileContext.Provider value={{ profile, fetchProfile }}>{children}</ProfileContext.Provider>;
};

export const useProfile = () => {
  const context = useContext(ProfileContext);
  if (!context) throw new Error('useProfile debe usarse dentro de un ProfileProvider');

  return context;
};
