import { CSSProperties, ReactNode, createContext, useContext } from "react";

export interface MainSidebarContent {
  content: ReactNode;
}

interface MainSidebarContextProps {
  isOpen: boolean;
  toggleSidebar: () => void;
  content: MainSidebarContent | null;
  setContent: (v: MainSidebarContent | null) => void;
}

export const MainSidebarContext = createContext<MainSidebarContextProps | null>(null);

export const useSidebar = () => {
  const ctx = useContext(MainSidebarContext);

  if (!ctx) throw new Error("SidebarContext is not provided");

  return ctx;
};
