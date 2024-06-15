import { createLazyFileRoute } from "@tanstack/react-router";
import { observer } from "mobx-react-lite";

const Page = observer(() => {
  const unom = Route.useParams().unom;

  return <div>Hello /_incidents/incidents/$unom! {unom}</div>;
});

export const Route = createLazyFileRoute("/_incidents/incidents/$unom")({
  component: Page
});
