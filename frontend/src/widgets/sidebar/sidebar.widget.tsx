import React, { FC, ReactElement, useContext, useEffect } from "react";
import { SidebarContent, SidebarContext } from "./sidebar.context";

interface Props extends Omit<SidebarContent, "content"> {
  children: ReactElement;
}

export const Sidebar: FC<Props> = (x) => {
  const ctx = useContext(SidebarContext);

  useEffect(() => {
    if (ctx) {
      ctx.setContent({
        content: React.cloneElement(x.children)
      });
    }
  }, [ctx, x.children]);

  return null;
};
