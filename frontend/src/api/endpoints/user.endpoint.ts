import { UserDto } from "../models/user.dto";
import api from "../utils/api";

export namespace UserEndpoint {
  export const current = () => api.get("/auth/me", UserDto.Item);
}
