import React, { type FC, type ReactNode } from "react";
import { User, BarChart3, Users, ChevronRight, LogOut } from "lucide-react";
import { useNavigate } from "react-router-dom";
import { ADMIN_LOGIN_PATH } from "../constants/routes";
import { Button } from "../components/ui/button";
import { clearAccessToken } from "../utils/local-storage";

interface AdminLayoutProps {
  children: ReactNode;
}

const AdminLayout: FC<AdminLayoutProps> = ({ children }) => {
  const [activeItem, setActiveItem] = React.useState("Users");
  const navigate = useNavigate();

  const menuItems = [
    {
      name: "Configure Score",
      icon: BarChart3,
      path: "/admin/configure/score"
    },
    { name: "Users", icon: Users, path: "/admin/users" }
  ];

  const handleMenuClick = (itemName: string, path: string) => {
    setActiveItem(itemName);
    navigate(path);
  };

  return (
    <div className="from-bg-cc-app-blue to-cc-app-mid-blue h-screen bg-gradient-to-br">
      <div className="bg-cc-app-blue fixed top-0 left-0 z-40 flex h-full w-72 flex-col justify-between border-r border-gray-200 shadow-xl">
        <div>
          <div className="relative border-b border-gray-100 p-6">
            <div className="absolute top-0 right-0 opacity-5">
              <div className="grid grid-cols-6 gap-1 p-2">
                {Array.from({ length: 24 }).map((_, i) => (
                  <div key={i} className="h-1 w-1 rounded-full bg-blue-600" />
                ))}
              </div>
            </div>

            <div className="relative flex items-center space-x-3">
              <div className="from-cc-app-blue to-cc-app-mid-blue flex h-10 w-10 items-center justify-center rounded-xl bg-gradient-to-br shadow-lg">
                <span className="text-lg font-bold text-white">&lt;/&gt;</span>
              </div>
              <div className="flex flex-col">
                <span className="text-lg font-bold tracking-tight text-white">
                  CODE CURIOSITY
                </span>
                <span className="text-xs font-medium tracking-wider text-white uppercase">
                  Admin Portal
                </span>
              </div>
            </div>
          </div>

          <nav className="space-y-2 p-4">
            {menuItems.map(item => {
              const IconComponent = item.icon;
              const isActive = activeItem === item.name;

              return (
                <button
                  key={item.name}
                  onClick={() => handleMenuClick(item.name, item.path)}
                  className={`group relative flex w-full items-center justify-between rounded-xl px-4 py-3 text-left font-medium transition-all duration-200 ${
                    isActive
                      ? "from-cc-app-blue to-cc-app-mid-blue scale-[1.02] transform bg-gradient-to-r text-white shadow-lg"
                      : "text-white hover:bg-gray-50 hover:text-gray-900 hover:shadow-md"
                  }`}
                >
                  <div className="flex items-center space-x-3">
                    <div
                      className={`flex h-8 w-8 items-center justify-center rounded-lg transition-colors duration-200 ${
                        isActive
                          ? "bg-white/20"
                          : "bg-gray-100 group-hover:bg-gray-200"
                      }`}
                    >
                      <IconComponent
                        className={`h-4 w-4 ${
                          isActive
                            ? "text-white"
                            : "text-gray-600 group-hover:text-gray-800"
                        }`}
                      />
                    </div>
                    <span className="text-sm">{item.name}</span>
                  </div>

                  <ChevronRight
                    className={`h-4 w-4 transition-all duration-200 ${
                      isActive
                        ? "rotate-90 transform text-white/80"
                        : "text-gray-400 group-hover:translate-x-1 group-hover:transform group-hover:text-gray-600"
                    }`}
                  />

                  {isActive && (
                    <div className="absolute top-1/2 left-0 h-8 w-1 -translate-y-1/2 rounded-r-full bg-white" />
                  )}
                </button>
              );
            })}
          </nav>
        </div>

        <div className="p-4">
          <Button
            onClick={() => {
              clearAccessToken();
              location.href = ADMIN_LOGIN_PATH;
            }}
            className="bg-cc-app-mid-blue hover:bg-cc-app-blue w-full rounded-md px-4 py-2 text-sm text-white transition-colors"
          >
            Logout <LogOut />
          </Button>
        </div>
      </div>

      <div className="ml-72 flex h-full flex-col">
        <header className="bg-cc-app-blue relative border-b border-gray-200 shadow-sm">
          <div className="flex items-center justify-between px-8 py-4">
            <div className="flex items-center space-x-4">
              <div className="flex flex-col">
                <h1 className="text-2xl font-bold tracking-tight text-white">
                  {activeItem}
                </h1>
                <p className="mt-1 text-sm text-white">
                  Manage your {activeItem.toLowerCase()} settings and
                  configurations
                </p>
              </div>
            </div>

            <div className="flex items-center space-x-4">
              <div className="flex items-center space-x-3 rounded-full border border-gray-200 bg-gray-50 px-4 py-2 transition-colors duration-200 hover:bg-gray-100">
                <div className="from-cc-app-blue to-cc-app-mid-blue flex h-8 w-8 items-center justify-center rounded-full bg-gradient-to-br">
                  <User className="h-4 w-4 text-white" />
                </div>
                <div className="flex flex-col">
                  <span className="text-sm font-medium text-gray-900">
                    Admin
                  </span>
                  <span className="text-xs text-gray-500">Administrator</span>
                </div>
              </div>
            </div>
          </div>
        </header>

        <main className="bg-cc-app-gray-background flex-1 overflow-auto bg-gradient-to-br">
          <div className="p-8">{children}</div>
        </main>
      </div>
    </div>
  );
};

export default AdminLayout;
