import { createLazyFileRoute } from "@tanstack/react-router";

export const Route = createLazyFileRoute(
  "/_map/_heat_sources/heat_source/$heatSourceId/_consumers/consumers/"
)({});
