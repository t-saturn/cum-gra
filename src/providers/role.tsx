"use client";

import { createContext, useContext } from "react";

export type RoleValue = { id: string; name: string } | null;

const RoleContext = createContext<RoleValue>(null);

export function RoleProvider({ value, children }: { value: RoleValue; children: React.ReactNode }) {
  return <RoleContext.Provider value={value}>{children}</RoleContext.Provider>;
}

export function useRole() {
  return useContext(RoleContext);
}
