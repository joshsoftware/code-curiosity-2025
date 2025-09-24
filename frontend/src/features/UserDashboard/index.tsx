import UserDashboardLayout from "@/shared/layout/UserDashboardLayout";
import UserDashboardComponent from "@/features/UserDashboard/components/UserDashboardComponent";

const UserDashboard = () => {
  return (
    <UserDashboardLayout>
      <UserDashboardComponent />
    </UserDashboardLayout>
  );
};

export default UserDashboard;
