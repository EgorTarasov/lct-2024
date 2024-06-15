import { z } from "zod";

export namespace ConsumersDto {
  export const Filters = z
    .object({
      filterName: z.string(),
      values: z.string().array()
    })
    .array();
}
