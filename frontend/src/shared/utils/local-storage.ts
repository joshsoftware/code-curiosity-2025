import { ACCESS_TOKEN_KEY, USER_DATA_KEY } from "@/shared/constants/local-storage";
import type { User } from "../types/types";

export const getAccessToken = (): string | null => {
  return localStorage.getItem(ACCESS_TOKEN_KEY) || null;
};

export const setAccessToken = (token: string) => {
  localStorage.setItem(ACCESS_TOKEN_KEY, token);
};

export const clearAccessToken = () => {
  localStorage.removeItem(ACCESS_TOKEN_KEY);
};

export const getUserData = () => {
  try {
    const userData = localStorage.getItem(USER_DATA_KEY);
    return userData ? (JSON.parse(userData) as User) : null;
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
  } catch (error) {
    return null;
  }
};

export const setUserData = (data: any) => {
  localStorage.setItem(USER_DATA_KEY, JSON.stringify(data));
};

export const clearUserCredentials = () => {
  localStorage.removeItem(USER_DATA_KEY);
  localStorage.removeItem(ACCESS_TOKEN_KEY);
};