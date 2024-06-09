import { z } from "zod";

export namespace UserDto {
  export const Item = z.object({
    firstName: z.string(),
    lastName: z.string()
  });
  export type Item = z.infer<typeof Item>;
}
