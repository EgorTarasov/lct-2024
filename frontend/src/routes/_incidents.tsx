import { ScrollArea } from "@/components/ui/scroll-area";
import { IncidentsOverlay } from "@/widgets/incidents/incidents-overlay";
import { IncidentsSidebar } from "@/widgets/incidents/incidents-sidebar";
import { IncidentsPageViewModel } from "@/widgets/incidents/incidents.page.vm";
import { Outlet, createFileRoute } from "@tanstack/react-router";
import { observer } from "mobx-react-lite";
import { useState } from "react";

const Page = observer(() => {
  const [vm] = useState(() => new IncidentsPageViewModel());

  return (
    <main className="size-full flex">
      <IncidentsSidebar vm={vm} />
      <IncidentsOverlay hasSidebar />
      <ScrollArea className="flex flex-col pt-16 flex-1">
        <Outlet />
      </ScrollArea>
    </main>
  );
});

export const Route = createFileRoute("/_incidents")({
  component: Page
});
