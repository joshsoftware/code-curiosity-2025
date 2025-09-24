import { type FC, type ReactNode, useEffect } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import { CheckCircle } from "lucide-react";

import { Card } from "@/shared/components/ui/card";
import { LOGIN_PATH, USER_DASHBOARD_PATH } from "@/shared/constants/routes";
import { getAccessToken } from "@/shared/utils/local-storage";

interface AuthLayoutProps {
  children: ReactNode;
}

const AuthLayout: FC<AuthLayoutProps> = ({ children }) => {
  const location = useLocation();
  const navigate = useNavigate();

  useEffect(() => {
    const userAccessToken = getAccessToken();
    const shouldRedirect = [LOGIN_PATH].includes(location.pathname);

    if (userAccessToken && shouldRedirect) {
      navigate(USER_DASHBOARD_PATH);
    }
  }, [navigate, location]);

  return (
    <div className="h-screen bg-white p-4">
      <div className="from-cc-app-light-blue via-cc-app-mid-blue to-cc-app-blue flex h-full w-full rounded-2xl bg-gradient-to-bl">
        <Card className="relative hidden w-1/2 overflow-hidden border-none bg-transparent shadow-none lg:flex">
          <div className="absolute top-8 left-8 opacity-20">
            <div className="grid grid-cols-8 gap-2">
              {Array.from({ length: 64 }).map((_, i) => (
                <div key={i} className="h-2 w-2 bg-white" />
              ))}
            </div>
          </div>

          <div className="flex h-full w-full items-center justify-center p-20">
            <div className="justify-star flex flex-1 items-center">
              <div className="max-w-md space-y-10">
                <div>
                  <h1 className="mb-1 text-4xl font-bold tracking-tight text-white">
                    Code Curiosity
                  </h1>
                </div>

                <div className="space-y-6">
                  {[
                    "Earn and Upskill",
                    "Set Your Goals",
                    "Leader Board",
                    "Open Source Contribution"
                  ].map((text, i) => (
                    <div
                      key={i}
                      className="group flex cursor-pointer items-center space-x-4 transition"
                    >
                      <CheckCircle className="h-6 w-6 text-blue-200 transition-colors duration-200 group-hover:text-white" />
                      <span className="text-lg font-medium text-blue-100 transition-colors duration-200 group-hover:text-white">
                        {text}
                      </span>
                    </div>
                  ))}
                </div>
              </div>
            </div>
          </div>

          <div className="absolute right-6 bottom-6 opacity-20">
            <div className="grid grid-cols-8 gap-2">
              {Array.from({ length: 64 }).map((_, i) => (
                <div key={i} className="h-2 w-2 bg-white" />
              ))}
            </div>
          </div>
        </Card>

        <Card className="bg-cc-app-gray-background flex w-1/2 flex-1 items-center justify-center p-6">
          {children}
        </Card>
      </div>
    </div>
  );
};

export default AuthLayout;
