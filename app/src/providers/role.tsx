'use client';

import { createContext, useContext } from 'react';

export type ModuleInfo = {
  id: string;
  name: string;
  route?: string | null;
  icon?: string | null;
  parent_id?: string | null;
  sort_order: number;
  status: string;
};

export type RoleValue = { 
  id: string; 
  name: string; 
  modules?: ModuleInfo[] 
} | null;

const RoleContext = createContext<RoleValue>(null);

export const RoleProvider = ({ value, children }: { value: RoleValue; children: React.ReactNode }) => (
  <RoleContext.Provider value={value}>{children}</RoleContext.Provider>
);

export const useRole = () => useContext(RoleContext);