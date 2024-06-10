import { NavigateFn } from "@tanstack/react-router";

export const checkGrant = (x: { navigate: NavigateFn }, allowed: boolean) => {
  if (allowed) {
    return;
  }

  x.navigate({
    to: "/login"
  });
};
