import { type FC, type ReactNode, useEffect } from "react";
import { Navigate, useLocation, useNavigate } from "react-router-dom";

import { LOGIN_PATH } from "@/shared/constants/routes";
import { getAccessToken } from "@/shared/utils/local-storage";

interface WithAuthProps {
  children: ReactNode;
}

const WithAuth: FC<WithAuthProps> = ({ children }) => {
  const navigate = useNavigate();
  const location = useLocation();
  const userAccessToken = getAccessToken();

  useEffect(() => {
    if (!userAccessToken) {
      navigate(LOGIN_PATH, { replace: true });
    }
  }, [userAccessToken, location.pathname, navigate]);

  if (!userAccessToken) {
    return <Navigate to={LOGIN_PATH} replace />;
  }

  return <>{children}</>;
};

export default WithAuth;