import React, { type FC, type ReactNode } from "react";
import { User, BarChart3, Users, ChevronRight } from "lucide-react";
import { useNavigate } from "react-router-dom";

interface AdminLayoutProps {
  children: ReactNode;
}

const AdminLayout: FC<AdminLayoutProps> = ({ children }) => {
  const [activeItem, setActiveItem] = React.useState("Users");
  const navigate = useNavigate();

  const menuItems = [
    { name: "Configure Score", icon: BarChart3, path: "/admin/configure/score" },
    { name: "Users", icon: Users, path: "/admin/users" }
  ];

  const handleMenuClick = (itemName: string, path: string) => {
    setActiveItem(itemName);
    navigate(path);
  };

  return (
    <div className="h-screen bg-gradient-to-br from-bg-cc-app-blue to-cc-app-mid-blue">
      <div className="fixed left-0 top-0 z-40 h-full w-72  bg-cc-app-blue shadow-xl border-r border-gray-200">
        <div className="relative p-6 border-b border-gray-100">
          <div className="absolute top-0 right-0 opacity-5">
            <div className="grid grid-cols-6 gap-1 p-2">
              {Array.from({ length: 24 }).map((_, i) => (
                <div key={i} className="h-1 w-1 bg-blue-600 rounded-full" />
              ))}
            </div>
          </div>
          
          <div className="relative flex items-center space-x-3">
            <div className="flex h-10 w-10 items-center justify-center rounded-xl bg-gradient-to-br from-cc-app-blue to-cc-app-mid-blue shadow-lg">
              <span className="text-lg font-bold text-white">&lt;/&gt;</span>
            </div>
            <div className="flex flex-col">
              <span className="text-lg font-bold text-white tracking-tight">
                CODE CURIOSITY
              </span>
              <span className="text-xs font-medium text-white uppercase tracking-wider">
                Admin Portal
              </span>
            </div>
          </div>
        </div>

        <nav className="flex-1 p-4 space-y-2">
          {menuItems.map((item) => {
            const IconComponent = item.icon;
            const isActive = activeItem === item.name;
            
            return (
              <button
                key={item.name}
                onClick={() => handleMenuClick(item.name, item.path)}
                className={`group relative flex w-full items-center justify-between rounded-xl px-4 py-3 text-left font-medium transition-all duration-200 ${
                  isActive
                    ? "bg-gradient-to-r from-cc-app-blue to-cc-app-mid-blue text-white shadow-lg transform scale-[1.02]"
                    : "text-white hover:bg-gray-50 hover:text-gray-900 hover:shadow-md"
                }`}
              >
                <div className="flex items-center space-x-3">
                  <div className={`flex h-8 w-8 items-center justify-center rounded-lg transition-colors duration-200 ${
                    isActive 
                      ? "bg-white/20" 
                      : "bg-gray-100 group-hover:bg-gray-200"
                  }`}>
                    <IconComponent className={`h-4 w-4 ${
                      isActive ? "text-white" : "text-gray-600 group-hover:text-gray-800"
                    }`} />
                  </div>
                  <span className="text-sm">{item.name}</span>
                </div>
                
                <ChevronRight className={`h-4 w-4 transition-all duration-200 ${
                  isActive 
                    ? "text-white/80 transform rotate-90" 
                    : "text-gray-400 group-hover:text-gray-600 group-hover:transform group-hover:translate-x-1"
                }`} />
                
                {isActive && (
                  <div className="absolute left-0 top-1/2 h-8 w-1 -translate-y-1/2 bg-white rounded-r-full" />
                )}
              </button>
            );
          })}
        </nav>

        <div className="absolute bottom-4 left-4 opacity-5">
          <div className="grid grid-cols-6 gap-1">
            {Array.from({ length: 18 }).map((_, i) => (
              <div key={i} className="h-1 w-1 bg-blue-600 rounded-full" />
            ))}
          </div>
        </div>
      </div>

      <div className="ml-72 flex flex-col h-full">
        <header className="relative bg-cc-app-blue shadow-sm border-b border-gray-200">
          <div className="flex items-center justify-between px-8 py-4">
            <div className="flex items-center space-x-4">
              <div className="flex flex-col">
                <h1 className="text-2xl font-bold text-white tracking-tight">
                  {activeItem}
                </h1>
                <p className="text-sm text-white mt-1">
                  Manage your {activeItem.toLowerCase()} settings and configurations
                </p>
              </div>
            </div>
            
            <div className="flex items-center space-x-4">
              <div className="flex items-center space-x-3 rounded-full bg-gray-50 px-4 py-2 border border-gray-200 hover:bg-gray-100 transition-colors duration-200">
                <div className="flex h-8 w-8 items-center justify-center rounded-full bg-gradient-to-br from-cc-app-blue to-cc-app-mid-blue">
                  <User className="h-4 w-4 text-white" />
                </div>
                <div className="flex flex-col">
                  <span className="text-sm font-medium text-gray-900">Admin</span>
                  <span className="text-xs text-gray-500">Administrator</span>
                </div>
              </div>
            </div>
          </div>
        </header>

        <main className="flex-1 overflow-auto bg-gradient-to-br bg-cc-app-gray-background">
          <div className="p-8">
            {children}
          </div>
        </main>
      </div>
    </div>
  );
};

export default AdminLayout;
