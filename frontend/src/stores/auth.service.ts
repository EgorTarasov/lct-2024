import { AuthEndpoint } from "@/api/endpoints/auth.endpoint";
import { UserEndpoint } from "@/api/endpoints/user.endpoint";
import { authToken } from "@/api/utils/authToken";
import { Auth } from "@/types/auth.type";
import { makeAutoObservable } from "mobx";

class AuthServiceViewModel {
  public auth: Auth.State = { state: "loading" };

  constructor() {
    makeAutoObservable(this);
    void this.init();
  }

  private async init() {
    if (!authToken.get()) {
      this.auth = { state: "anonymous" };
      return;
    }

    try {
      const user = await UserEndpoint.current();
      this.auth = { state: "authenticated", user };
    } catch {
      this.auth = { state: "anonymous" };
    }
  }

  login = async (v: AuthEndpoint.LoginTemplate): Promise<boolean> => {
    try {
      await AuthEndpoint.login(v);

      const user = await UserEndpoint.current();
      this.auth = { state: "authenticated", user };
      return true;
    } catch {
      return false;
    }
  };

  // loginViaVk = async (code: unknown): Promise<boolean> => {
  //   try {
  //     await AuthEndpoint.loginViaVk(code);

  //     const user = await UserEndpoint.current();
  //     this.auth = { state: "authenticated", user };
  //     return true;
  //   } catch {
  //     toast({
  //       variant: "destructive",
  //       description: "Не удалось войти через ВКонтакте",
  //       title: "Ошибка входа"
  //     });
  //     return false;
  //   }
  // };

  register = async (v: AuthEndpoint.RegisterTemplate): Promise<boolean> => {
    try {
      await AuthEndpoint.register(v);

      const user = await UserEndpoint.current();
      this.auth = { state: "authenticated", user };
      return true;
    } catch {
      return false;
    }
  };

  logout() {
    this.auth = { state: "anonymous" };
  }

  // async updatePassword(token: string, newPassword: string) {
  //   try {
  //     await AuthEndpoint.resetPassword(token, newPassword);

  //     const user = await UserEndpoint.current();
  //     this.auth = { state: "authenticated", user };
  //     return true;
  //   } catch {
  //     return false;
  //   }
  // }
}

export const AuthService = new AuthServiceViewModel();