import { createContext, useMemo, useState, type ReactNode } from 'react';

export type User = {
  githubId: string;
  githubUsername: string;
  avatarUrl: string;
};

export interface UserContextInterface {
  user: User;
  login: (user: User, token: string) => void;
  logout: () => void;
}

const defaultUser: User = {
  githubId: '',
  githubUsername: '',
  avatarUrl: ''
};

const defaultState: UserContextInterface = {
  user: defaultUser,
  login: () => {
    throw new Error('login must be used within UserProvider');
  },
  logout: () => {
    throw new Error('logout must be used within UserProvider');
  }
};

export const UserContext = createContext(defaultState);

type UserProviderProps = {
  children: ReactNode;
};

export const UserProvider = ({ children }: UserProviderProps) => {
  const [user, setUser] = useState<User>(defaultUser);

  const login = (newUser: User, token: string) => {
    setUser(newUser);
    localStorage.setItem('token', token);
  };

  const logout = () => {
    setUser(defaultUser);
    localStorage.removeItem('token');
  };

  const value = useMemo(() => ({ user, login, logout }), [user]);

  return <UserContext.Provider value={value}>{children}</UserContext.Provider>;
};
