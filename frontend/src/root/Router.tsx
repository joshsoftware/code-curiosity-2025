import { createBrowserRouter, RouterProvider } from "react-router-dom";
import PrivateRoutes from "./PrivateRoutes";
import { routesConfig, type RoutesType } from "./routesConfig";

const generateRoutes = (routes: RoutesType[]) => {
  return routes.map(({ path, element, isProtected }) => {
    let wrappedElement = element;

    if (isProtected) {
      wrappedElement = <PrivateRoutes>{wrappedElement}</PrivateRoutes>;
    }

    return { path, element: wrappedElement };
  });
};

const Router = () => {
  const router = createBrowserRouter(generateRoutes(routesConfig));
  return <RouterProvider router={router} />;
};

export default Router;
