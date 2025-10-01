import Repository from "./components/Repository";
import Languages from "./components/Languages";
import RecentActivities from "./components/RepositoryActivities";
import ContributorsList from "./components/Contributors";
import { ArrowLeft } from "lucide-react";
import { useNavigate } from "react-router-dom";
import { Separator } from "@/shared/components/ui/separator";

const RepositoryDetails = () => {
  const navigate = useNavigate();
  return (
    <div className="space-y-6 p-8 px-30">
      <div
        className="flex cursor-pointer items-center space-x-2"
        onClick={() => navigate("/my-contributions")}
      >
        <ArrowLeft className="h-5 w-5 text-gray-600" />
        <span className="text-cc-app-blue text-lg font-medium">
          Repository Details
        </span>
      </div>
      <div>
        <div className="rounded-xl border border-gray-300">
          <Repository />
          <Separator className="w-max border-t border-dashed border-gray-300" />
          <ContributorsList />
        </div>
        <div className="mt-6 grid grid-cols-1 gap-4 md:grid-cols-2">
          <RecentActivities />
          <Languages />
        </div>
      </div>
    </div>
  );
};

export default RepositoryDetails;
