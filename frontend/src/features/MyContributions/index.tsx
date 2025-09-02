import Repositories from "./components/Repositories";

const MyContributions = () => {
  return (
    <div className="space-y-6 p-8 px-30">
      <div className="flex cursor-pointer items-center space-x-2">
        <span className="text-cc-app-blue text-lg font-medium">
          My Contributed Repositories
        </span>
      </div>
      <div>
        <Repositories />
      </div>
    </div>
  );
};

export default MyContributions;
