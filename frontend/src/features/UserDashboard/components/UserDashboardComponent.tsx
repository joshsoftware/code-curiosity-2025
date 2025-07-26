import Leaderboard from "@/features/UserDashboard/components/Leaderboard";
import RecentActivities from "@/features/UserDashboard/components/RecentActivities";
import Overview from "@/features/UserDashboard/components/Overview";

const UserDashboardComponent = () => {
  return (
    <div className="flex h-full w-full gap-4">
      <Leaderboard className="w-[45%]" />
      <div className="flex w-[55%] flex-col gap-4">
        <Overview className="h-[45%]" />
        <RecentActivities className="h-[55%]" />
      </div>
    </div>
  );
};

export default UserDashboardComponent;
