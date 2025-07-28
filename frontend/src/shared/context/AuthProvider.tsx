import { type ReactNode, createContext, useMemo, useState } from "react";

import { clearAccessToken, setAccessToken } from "@/shared/utils/local-storage";

export type UserCredentials = {
  githubId: string;
  githubUsername: string;
  avatarUrl: string;
};

export interface AuthContextInterface {
  userCredentials: UserCredentials | null;
  login: (userCredentials: UserCredentials, token: string) => void;
  logout: () => void;
}

const AuthContext = createContext<AuthContextInterface>({
  userCredentials: null,
  login: () => {
    throw new Error("AuthContext: login called outside AuthProvider");
  },
  logout: () => {
    throw new Error("AuthContext: logout called outside AuthProvider");
  }
});

type AuthProviderProps = {
  children: ReactNode;
};

export const AuthProvider = ({ children }: AuthProviderProps) => {
  const [userCredentials, setUserCredentials] =
    useState<UserCredentials | null>(null);

  const login = (userCredentials: UserCredentials, token: string) => {
    setUserCredentials(userCredentials);
    setAccessToken(token);
  };

  const logout = () => {
    setUserCredentials(null);
    clearAccessToken();
  };

  const value = useMemo(
    () => ({ userCredentials, login, logout }),
    [userCredentials]
  );

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
};

export { AuthContext };
