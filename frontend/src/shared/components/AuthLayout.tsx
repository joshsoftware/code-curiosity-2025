import { useEffect, type FC, type ReactNode } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';
import { getAccessToken, getUserData } from '../utils/local-storage';
import { ROUTES } from '@/root/routeConstants';
import { Card } from '@/components/ui/card';
import { CheckCircle } from 'lucide-react';

interface AuthLayoutProps {
  children: ReactNode;
}

const AuthLayout: FC<AuthLayoutProps> = ({ children }) => {
  const location = useLocation();
  const navigate = useNavigate();

  useEffect(() => {
    const userAccessToken = getAccessToken();
    const userData = getUserData();
    const shouldRedirect = [ROUTES.LOGIN].includes(location.pathname);

    if (userAccessToken && userData && shouldRedirect) {
      navigate(ROUTES.LANDING);
    }
  }, [navigate, location.pathname]);

  return (
    <div className="h-screen bg-white p-4">
      <div className="flex h-full w-full bg-gradient-to-br from-cc-mid-blue via-cc-app-blue to-cc-app-blue rounded-2xl">
        <Card className="hidden lg:flex w-1/2 bg-transparent relative overflow-hidden shadow-none border-none">
          <div className="absolute top-8 left-8 opacity-20">
            <div className="grid grid-cols-8 gap-2">
              {Array.from({ length: 64 }).map((_, i) => (
                <div key={i} className="w-2 h-2 bg-white " />
              ))}
            </div>
          </div>

          <div className="p-20 flex items-center justify-center w-full h-full">
            <div className=" flex flex-1 items-center justify-star">
              <div className="space-y-10 max-w-md">
                <div>
                  <h1 className="text-4xl font-bold text-white mb-1 tracking-tight">
                    Code Curiosity
                  </h1>
                </div>

                <div className="space-y-6">
                  {[
                    'Earn and Upskill',
                    'Set Your Goals',
                    'Leader Board',
                    'Open Source Contribution'
                  ].map((text, i) => (
                    <div
                      key={i}
                      className="flex items-center space-x-4 group cursor-pointer transition"
                    >
                      <CheckCircle className="w-6 h-6 text-blue-200 group-hover:text-white transition-colors duration-200" />
                      <span className="text-lg text-blue-100 group-hover:text-white font-medium transition-colors duration-200">
                        {text}
                      </span>
                    </div>
                  ))}
                </div>
              </div>
            </div>
          </div>

          {/* Bottom Grid Decoration */}
          <div className="absolute bottom-6 right-6 opacity-20">
            <div className="grid grid-cols-8 gap-2">
              {Array.from({ length: 64 }).map((_, i) => (
                <div key={i} className="w-2 h-2 bg-white" />
              ))}
            </div>
          </div>
        </Card>

        <Card className="flex-1 w-1/2  bg-cc-app-gray-background flex items-center justify-center p-6">
          {children}
        </Card>
      </div>
    </div>
  );
};

export default AuthLayout;
