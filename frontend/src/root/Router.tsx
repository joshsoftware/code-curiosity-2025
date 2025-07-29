import { RouterProvider, createBrowserRouter } from "react-router-dom";

import WithAuth from "@/shared/HOC/WithAuth";
import { type RoutesType, routesConfig } from "@/root/routes-config";
import { Layout } from "@/shared/constants/layout";
import AuthLayout from "@/shared/layout/AuthLayout";
import UserDashboardLayout from "@/shared/layout/UserDashboardLayout";

const generateRoutes = (routes: RoutesType[]) => {
  return routes.map(({ path, element, isProtected, layout }) => {
    let wrappedElement = element;

    if (isProtected) {
      wrappedElement = <WithAuth>{wrappedElement}</WithAuth>;
    }

    if (layout == Layout.AuthLayout) {
      wrappedElement = <AuthLayout>{wrappedElement}</AuthLayout>;
    }

    if (layout == Layout.DashboardLayout) {
      wrappedElement = (
        <UserDashboardLayout>{wrappedElement}</UserDashboardLayout>
      );
    }

    return { path, element: wrappedElement };
  });
};

const Router = () => {
  const router = createBrowserRouter(generateRoutes(routesConfig));
  return <RouterProvider router={router} />;
};

export default Router;
