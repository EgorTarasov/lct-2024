import { z } from "zod";

export namespace UserDto {
  export const Item = z.object({});
  export type Item = z.infer<typeof Item>;
}
