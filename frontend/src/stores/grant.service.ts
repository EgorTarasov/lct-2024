import { makeAutoObservable } from "mobx";
import { AuthService } from "./auth.service";
import { Auth } from "@/types/auth.type";

const isLoggedIn = (state: Auth.State): state is Auth.Authenticated =>
  state.state === "authenticated";

class grantService {
  constructor() {
    makeAutoObservable(this);
  }

  get canReadDirectory() {
    if (!isLoggedIn(AuthService.auth)) return false;

    return true;
  }
}

export const GrantService = new grantService();
