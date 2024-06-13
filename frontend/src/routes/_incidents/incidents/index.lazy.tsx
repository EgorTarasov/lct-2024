import { createLazyFileRoute } from "@tanstack/react-router";

export const Route = createLazyFileRoute("/_incidents/incidents/")({
  component: () => <div className="flex bg-red-200 h-[2000px]">Hello /_incidents/incidents/!</div>
});
