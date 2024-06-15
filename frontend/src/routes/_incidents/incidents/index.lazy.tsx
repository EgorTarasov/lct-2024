import { createLazyFileRoute } from "@tanstack/react-router";

export const Route = createLazyFileRoute("/_incidents/incidents/")({
  component: () => <div className="flex h-[2000px]">Hello /_incidents/incidents/!</div>
});
