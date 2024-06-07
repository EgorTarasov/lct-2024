import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/_map/_with_filters")({
  component: () => <div>Hello /_default!</div>
});
