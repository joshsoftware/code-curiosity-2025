import type { ReactNode } from "react";
import { Layout, type LayoutType } from "@/shared/constants/layout";
import Login from "@/features/Login";
import MyContributions from "@/features/MyContributions";
import UserDashboard from "@/features/UserDashboard";
import {
  LOGIN_PATH,
  MY_CONTRIBUTIONS_PATH,
  REPOSITORY_DETAILS_PATH,
  USER_DASHBOARD_PATH
} from "@/shared/constants/routes";
import RepositoryDetails from "@/features/RepositoryDetails.tsx";
import AdminLogin from "@/features/Admin/AdminLogin";
import { AllUsersList } from "@/features/Admin/Users.tsx";
import ScoreConfigure from "@/features/Admin/ScoreConfigure";
export interface RoutesType {
  path: string;
  element: ReactNode;
  isProtected?: boolean;
  layout: LayoutType;
}

export const routesConfig: RoutesType[] = [
  {
    path: LOGIN_PATH,
    element: <Login />,
    isProtected: false,
    layout: Layout.AuthLayout
  },
  {
    path: USER_DASHBOARD_PATH,
    element: <UserDashboard />,
    isProtected: true,
    layout: Layout.DashboardLayout
  },
  {
    path: MY_CONTRIBUTIONS_PATH,
    element: <MyContributions />,
    isProtected: true,
    layout: Layout.DashboardLayout
  },
  {
    path: REPOSITORY_DETAILS_PATH,
    element: <RepositoryDetails />,
    isProtected: true,
    layout: Layout.DashboardLayout
  },
  {
    path: "/admin/login",
    element: <AdminLogin />,
    isProtected: false,
    layout: Layout.AuthLayout
  },
  {
    path: "/admin/users",
    element: <AllUsersList />,
    isProtected: true,
    layout: Layout.AdminLayout
  },
  {
    path: "/admin/configure/score",
    element: <ScoreConfigure />,
    isProtected: true,
    layout: Layout.AdminLayout
  }
];
