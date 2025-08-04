import React, { type FC, type ReactNode } from "react";
import { User, BarChart3, Users } from "lucide-react";

interface AdminLayoutProps {
  children: ReactNode;
}

const AdminLayout: FC<AdminLayoutProps> = ({ children }) => {
  const [activeItem, setActiveItem] = React.useState("Users");

  const menuItems = [
    { name: "Configure Score", icon: BarChart3 },
    { name: "Users", icon: Users }
  ];

  return (
    <div className="flex h-screen bg-gray-100">
      <div className="flex w-64 flex-col bg-gradient-to-b from-blue-600 to-blue-800 text-white">
        <div className="border-b border-blue-500 p-4">
          <div className="flex items-center space-x-2">
            <div className="text-xl font-bold">
              <span className="text-white">&lt;/&gt;</span>
            </div>
            <span className="text-sm font-medium tracking-wide">
              CODE CURIOSITY
            </span>
          </div>
        </div>

        <nav className="flex-1 pt-4">
          {menuItems.map(item => {
            const IconComponent = item.icon;
            return (
              <button
                key={item.name}
                onClick={() => setActiveItem(item.name)}
                className={`flex w-full items-center px-4 py-3 text-left text-sm font-medium transition-colors duration-200 hover:bg-blue-700 ${
                  activeItem === item.name
                    ? "border-r-2 border-white bg-blue-700"
                    : ""
                }`}
              >
                <IconComponent className="mr-3 h-4 w-4" />
                {item.name}
              </button>
            );
          })}
        </nav>
      </div>

      <div className="flex flex-1 flex-col">
        <div className="flex items-center justify-between bg-blue-600 px-6 py-3 text-white shadow-sm">
          <div className="text-lg font-medium">{activeItem}</div>
          <div className="flex items-center space-x-2">
            <User className="h-5 w-5" />
            <span className="text-sm">Hi Admin</span>
          </div>
        </div>

        <div className="flex-1 overflow-auto">{children}</div>
      </div>
    </div>
  );
};

export default AdminLayout;
