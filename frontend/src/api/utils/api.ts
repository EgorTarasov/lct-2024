/* eslint-disable @typescript-eslint/no-explicit-any */
import axios, { AxiosRequestConfig, AxiosResponse } from "axios";
import { ZodObject, z } from "zod";
import { authToken } from "./authToken";
import { toast } from "sonner";

axios.defaults.baseURL = "/api/";

const getConfig = (config?: AxiosRequestConfig<unknown>) => ({
  ...config,
  headers: {
    Authorization: `Bearer channel/${authToken.get()}`,
    "Access-Control-Allow-Origin": "*",
    ...config?.headers
  }
});

const handleRequest = async <T extends z.ZodRawShape>(
  req: () => Promise<AxiosResponse>,
  schema: ZodObject<T>
): Promise<z.infer<typeof schema>> => {
  try {
    const res = await req();

    if (schema) {
      const validatedData = schema.parse(res.data);

      return validatedData as any;
    }

    return res as any;
  } catch (error) {
    if (error instanceof z.ZodError) {
      toast.error("Ошибка валидации данных", {
        description: error.errors.map((e) => e.message).join("\n")
      });
      console.error("Validation error:", error.errors);
    } else if (axios.isAxiosError(error)) {
      if (error.response?.status === 401) {
        toast.error("Сессия истекла, пожалуйста, войдите снова");
        authToken.remove();
        window.location.replace("/login");
      } else if (error.response?.status === 500) {
        toast.error("Внутренняя ошибка сервера", {
          description: "Мы все узнаем и починим, попробуйте позже"
        });
      } else {
        toast.error("Неизвестная ошибка при выполнении запроса");
      }
      console.error("Axios error:", error.message, error.status);
    } else {
      toast.error("Неизвестная ошибка при выполнении запроса");
      console.error("Unexpected error:", error);
    }
    throw error;
  }
};

const get = <T extends z.ZodRawShape>(
  path: string,
  schema: ZodObject<T>,
  config?: AxiosRequestConfig<unknown>
): Promise<z.infer<typeof schema>> =>
  handleRequest(() => axios.get(path, getConfig(config)), schema) as any;

const post = <T extends z.ZodRawShape>(
  path: string,
  schema: ZodObject<T>,
  variables?: unknown,
  config?: AxiosRequestConfig<unknown>
): Promise<z.infer<typeof schema>> =>
  handleRequest(() => axios.post(path, variables, getConfig(config)), schema) as any;

const put = <T extends z.ZodRawShape>(
  path: string,
  schema: ZodObject<T>,
  variables?: unknown,
  config?: AxiosRequestConfig<unknown>
): Promise<z.infer<typeof schema>> =>
  handleRequest(() => axios.put(path, variables, getConfig(config)), schema);

const del = <T extends z.ZodRawShape>(
  path: string,
  schema: ZodObject<T>,
  config?: AxiosRequestConfig<unknown>
): Promise<z.infer<typeof schema>> =>
  handleRequest(() => axios.delete(path, getConfig(config)), schema) as any;

export default {
  get,
  post,
  put,
  delete: del
};
