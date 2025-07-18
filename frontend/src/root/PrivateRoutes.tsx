import { useEffect, type ReactNode } from "react";
import { Navigate } from "react-router-dom";
import { toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

interface PrivateRoutesProps {
  children: ReactNode;
}

const PrivateRoutes: React.FC<PrivateRoutesProps> = ({ children }) => {
  const localStorageToken = localStorage.getItem("token");

  useEffect(() => {
    if (!localStorageToken) {
      toast.warning("You need to login to proceed!");
    }
  }, [localStorageToken]);

  return localStorageToken ? children : <Navigate to="/user/login" replace />;
};

export default PrivateRoutes;
