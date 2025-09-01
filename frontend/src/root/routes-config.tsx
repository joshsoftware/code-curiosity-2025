import type { ReactNode } from "react";

import Login from "@/features/Login";
import MyContributions from "@/features/MyContributions";
import UserDashboard from "@/features/UserDashboard";
import {
  LOGIN_PATH,
  MY_CONTRIBUTIONS_PATH,
  USER_DASHBOARD_PATH
} from "@/shared/constants/routes";

export interface RoutesType {
  path: string;
  element: ReactNode;
  isProtected?: boolean;
}

export const routesConfig: RoutesType[] = [
  {
    path: LOGIN_PATH,
    element: <Login />,
    isProtected: false
  },
  {
    path: USER_DASHBOARD_PATH,
    element: <UserDashboard />,
    isProtected: false
  },
  {
    path: MY_CONTRIBUTIONS_PATH,
    element: <MyContributions />,
    isProtected: false
  }
];
