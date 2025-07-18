/* eslint-disable react-refresh/only-export-components */
import {
  createContext,
  useState,
  type Dispatch,
  type ReactNode,
  type SetStateAction,
} from "react";

export type User = {
  githubId: string;
  githubUsername: string;
  avatarUrl: string;
};

export interface UserContextInterface {
  user: User;
  setUser: Dispatch<SetStateAction<User>>;
}

const defaultState = {
  user: {
    githubId: "",
    githubUsername: "",
    avatarUrl: "",
  },
  setUser: (user: User) => {},
} as UserContextInterface;

export const UserContext = createContext(defaultState);

type UserProviderProps = {
  children: ReactNode;
};

export const UserProvider = ({ children }: UserProviderProps) => {
  const [user, setUser] = useState<User>(defaultState.user);

  return (
    <UserContext.Provider value={{ user, setUser }}>
      {children}
    </UserContext.Provider>
  );
};
