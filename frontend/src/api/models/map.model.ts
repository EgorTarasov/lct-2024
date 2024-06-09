import { z } from "zod";

export namespace MapDto {
  export const Moek = z
    .object({
      loadActual: z.number(),
      loadAvg: z.number(),
      loadHeating: z.number(),
      loadVent: z.number(),
      number: z.string(),
      point: z.string(),
      polygon: z.string(),
      src: z.string(),
      type: z.string()
    })
    .array();
}
