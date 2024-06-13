import { z } from "zod";

export namespace MapDto {
  export const Property = z.object({
    number: z.string(),
    Source: z.string(),
    balanceHolder: z.string(),
    commissioningDate: z.number(),
    consumerAddress: z.object({
      unom: z.number(),
      address: z.string(),
      border: z
        .array(
          z.object({
            Key: z.string(),
            Value: z.string().or(z.number().array().length(2).array().array()).or(z.any())
          })
        )
        .nullable(),
      municipalDistrict: z.string(),
      center: z.object({
        type: z.string(),
        coordinates: z.array(z.number())
      })
    }),
    district: z.string(),
    type: z.string(),
    locationType: z.string(),
    heatingPointAddress: z.object({
      unom: z.number(),
      address: z.string(),
      municipalDistrict: z.string(),
      border: z
        .array(
          z.object({
            Key: z.string(),
            Value: z.string().or(z.number().array().length(2).array().array()).or(z.any())
          })
        )
        .nullable(),
      center: z.object({
        type: z.string(),
        coordinates: z.array(z.number())
      })
    })
  });
  export type Property = z.infer<typeof Property>;
}
