import { Card } from "@/shared/components/ui/card";
import { Separator } from "@/shared/components/ui/separator";
import UserProfileDetails from "@/shared/components/UserDashboard/UserProfileDetails";
import UserBadges from "@/shared/components/UserDashboard/UserBadges";
import UserGoals from "@/shared/components/UserDashboard/UserGoals";

const UserProfileCard = () => {
  return (
    <Card className="bg-cc-app-blue h-full w-full max-w-md overflow-hidden border-none shadow-none">
      <div className="no-scrollbar relative h-full overflow-auto p-8 text-center">
        <div className="absolute top-0 left-6 opacity-20">
          <div className="grid grid-cols-6 gap-2">
            {Array.from({ length: 30 }).map((_, i) => (
              <div key={i} className="h-2 w-2 bg-white" />
            ))}
          </div>
        </div>
        <UserProfileDetails />
        <Separator className="my-6 bg-gradient-to-r from-transparent via-white to-transparent opacity-70" />
        <UserBadges />
        <Separator className="my-6 bg-gradient-to-r from-transparent via-white to-transparent opacity-70" />
        <UserGoals />
      </div>
    </Card>
  );
};

export default UserProfileCard;
