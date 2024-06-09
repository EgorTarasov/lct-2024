import { z } from "zod";

export namespace MapDto {
  export const Property = z
    .object({
      globalID: z.number(),
      polygon: z
        .object({
          Key: z.string(),
          Value: z.any()
        })
        .array()
    })
    .array();
}
