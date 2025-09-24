import { ACCESS_TOKEN_KEY } from "@/shared/constants/local-storage";

export const getAccessToken = (): string | null => {
  return localStorage.getItem(ACCESS_TOKEN_KEY) || null;
};

export const setAccessToken = (token: string) => {
  localStorage.setItem(ACCESS_TOKEN_KEY, token);
};

export const clearAccessToken = () => {
  localStorage.removeItem(ACCESS_TOKEN_KEY);
};
