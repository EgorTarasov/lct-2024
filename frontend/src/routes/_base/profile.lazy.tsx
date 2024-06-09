import { createLazyFileRoute } from "@tanstack/react-router";

export const Route = createLazyFileRoute("/_base/profile")({
  component: () => <div>Hello /_base/profile!</div>
});
