import type { ReactNode } from "react";
import Login from "../pages/Login/main";
import { ROUTES } from "./routeConstants";

export interface RoutesType {
  path: string;
  element: ReactNode;
  isProtected: boolean;
}

export const routesConfig: RoutesType[] = [
  {
    path: ROUTES.LOGIN,
    element: <Login />,
    isProtected: false,
  },
  //   {
  //     path: //path,
  //     element: (
  //       <PrivateRoutes>
  //         {/* protected component */}
  //       </PrivateRoutes>
  //     ),
  //   },
];
