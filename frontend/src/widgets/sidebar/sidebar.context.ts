import { CSSProperties, PropsWithChildren, ReactNode, createContext, useContext } from "react";

export interface SidebarContent {
  content: ReactNode;
  width?: CSSProperties["width"];
}

interface SidebarContextProps {
  isOpen: boolean;
  toggleSidebar: () => void;
  content: SidebarContent | null;
  setContent: (v: SidebarContent | null) => void;
}

export const SidebarContext = createContext<SidebarContextProps | null>(null);

export const useSidebar = () => {
  const ctx = useContext(SidebarContext);

  if (!ctx) throw new Error("SidebarContext is not provided");

  return ctx;
};
