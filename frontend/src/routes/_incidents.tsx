import { IncidentsContext } from "@/components/hoc/incidents-context";
import { ScrollArea } from "@/components/ui/scroll-area";
import { IncidentsOverlay } from "@/widgets/incidents/incidents-overlay";
import { IncidentsSidebar, IncidentsSidebarMobile } from "@/widgets/incidents/incidents-sidebar";
import { IncidentsPageViewModel } from "@/widgets/incidents/incidents.page.vm";
import { Outlet, createFileRoute } from "@tanstack/react-router";
import { observer } from "mobx-react-lite";
import { useState } from "react";

const Page = observer(() => {
  const [vm] = useState(() => new IncidentsPageViewModel());

  return (
    <IncidentsContext.Provider value={{ vm }}>
      <main className="size-full flex">
        <IncidentsSidebarMobile vm={vm} />
        <IncidentsSidebar vm={vm} />
        <IncidentsOverlay vm={vm} hasSidebar />
        <ScrollArea className="flex flex-col pt-[72px] flex-1">
          <Outlet />
        </ScrollArea>
      </main>
    </IncidentsContext.Provider>
  );
});

export const Route = createFileRoute("/_incidents")({
  component: Page
});
