import { Link, useLocation } from "react-router-dom";

import { Card } from "@/shared/components/ui/card";
import { USER_DASHBOARD_NAVBAR_OPTIONS } from "@/shared/types/navbar";

const Navbar = () => {
  const location = useLocation();

  const isActive = (path: string) => location.pathname === path;

  return (
    <Card className="bg-cc-app-blue flex h-16 w-full flex-row items-center justify-center gap-2 rounded-full border-none shadow-none">
      {USER_DASHBOARD_NAVBAR_OPTIONS.map(option => (
        <Link
          key={option.path}
          to={option.path}
          className={`rounded-full px-6 py-3 font-medium text-white transition-colors duration-200 ${
            isActive(option.path)
              ? "bg-cc-app-mid-blue"
              : "hover:bg-[#003aa5]"
          } `}
        >
          {option.name}
        </Link>
      ))}
    </Card>
  );
};

export default Navbar;
