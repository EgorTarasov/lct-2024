import React, { FC, ReactElement, useContext, useEffect } from "react";
import { MainSidebarContent, MainSidebarContext } from "./main-sidebar.context";
import { ReactNode } from "@tanstack/react-router";

interface Props extends Omit<MainSidebarContent, "content"> {
  children: ReactElement;
}

export const MainSidebar: FC<Props> = (x) => {
  const ctx = useContext(MainSidebarContext);

  useEffect(() => {
    if (ctx) {
      ctx.setContent({
        content: React.cloneElement(x.children)
      });
    }
  }, [x.children]);

  return null;
};
