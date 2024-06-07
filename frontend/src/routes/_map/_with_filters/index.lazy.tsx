import { createLazyFileRoute } from "@tanstack/react-router";

export const Route = createLazyFileRoute("/_map/_with_filters/")({
  component: () => <div>Hello /(map)/!</div>
});
