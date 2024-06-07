import { observer } from "mobx-react-lite";
import { useSidebar } from "./sidebar.context";
import { twMerge } from "tailwind-merge";

export const SidebarContainer = observer(() => {
  const ctx = useSidebar();

  return (
    <aside
      className={twMerge("absolute left-0 bottom-0 top-0 bg-card")}
      style={{ width: ctx.content?.width }}>
      {ctx.content?.content}
    </aside>
  );
});
