import { IncidentsPageViewModel } from "@/widgets/incidents/incidents.page.vm";
import { createContext, useContext } from "react";

export const IncidentsContext = createContext<{ vm: IncidentsPageViewModel | null }>({ vm: null });

export const useIncidentsContext = () => {
  const context = useContext(IncidentsContext);
  if (!context.vm) {
    throw new Error("useIncidentsContext must be used within a IncidentsProvider");
  }
  return context.vm;
};
