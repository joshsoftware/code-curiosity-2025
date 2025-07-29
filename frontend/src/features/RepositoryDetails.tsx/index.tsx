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
    <div className="space-y-6 px-30 p-8">
      <div className="flex cursor-pointer items-center space-x-2">
        <ArrowLeft className="h-5 w-5 text-gray-600" />
        <span
          className="text-lg font-medium text-blue-600"
          onClick={() => navigate("/my-contributions")}
        >
          Repository Details
        </span>
      </div>
      <div className="">
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
