import type { FC, ReactNode } from "react";

import Navbar from "@/shared/components/UserDashboard/Navbar";
import UserProfileCard from "@/shared/components/UserDashboard/UserProfileCard";

interface UserDashboardLayoutProps {
  children?: ReactNode;
}

const UserDashboardLayout: FC<UserDashboardLayoutProps> = ({ children }) => {
  return (
    <div className="bg-cc-app-gray-background flex h-screen w-full gap-6 p-8">
      <UserProfileCard />
      <div className="flex h-full w-full flex-col gap-4 rounded-xl bg-white p-4">
        <Navbar />
        <div className="no-scrollbar h-full w-full overflow-auto rounded-xl bg-white">
          {children}
        </div>
      </div>
    </div>
  );
};

export default UserDashboardLayout;
