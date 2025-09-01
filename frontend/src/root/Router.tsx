import { RouterProvider, createBrowserRouter } from "react-router-dom";

import WithAuth from "@/shared/HOC/WithAuth";
import { type RoutesType, routesConfig } from "@/root/routes-config";

const generateRoutes = (routes: RoutesType[]) => {
  return routes.map(({ path, element, isProtected }) => {
    let wrappedElement = element;

    if (isProtected) {
      wrappedElement = <WithAuth>{wrappedElement}</WithAuth>;
    }

    return { path, element: wrappedElement };
  });
};

const Router = () => {
  const router = createBrowserRouter(generateRoutes(routesConfig));
  return <RouterProvider router={router} />;
};

export default Router;
