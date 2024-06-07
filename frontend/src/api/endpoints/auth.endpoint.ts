import api from "api/utils/api";
import { AuthDto } from "../models/auth.model";

export namespace AuthEndpoint {
  export const login = async (username: string, password: string) =>
    api.post("/auth/login", AuthDto.Token, {
      username,
      password
    });

  export type RegisterTemplate = {
    username: string;
    password: string;
    firstName: string;
    lastName: string;
  };
  export const register = async (v: RegisterTemplate) =>
    api.post("/auth/register", AuthDto.Token, v);
}
