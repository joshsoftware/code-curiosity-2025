import {
  Dialog,
  DialogTrigger,
  DialogContent,
  DialogHeader,
  DialogTitle,
} from "@/shared/components/ui/dialog";
import { Button } from "@/shared/components/ui/button";
import { useAllContributionTypes } from "@/api/queries/UserGoals";

const CoinsInfo = () => {
  const { data, isLoading } = useAllContributionTypes();
  const contributions = data?.data ?? [];

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button
          variant="link"
          className="text-cc-app-blue px-0 text-xs font-semibold underline"
        >
          How do points work?
        </Button>
      </DialogTrigger>
      <DialogContent className="max-w-md sm:max-w-lg">
        <DialogHeader>
          <DialogTitle className="text-lg font-bold">
            How does the point system work?
          </DialogTitle>
        </DialogHeader>

        <div className="text-sm leading-relaxed text-gray-600">
          We assign coins for every open-source contribution you make on GitHub.
          Contributions are updated everyday at midnight, so you can see
          contributions you made before 12 am yesterday
        </div>

        <div className="mt-4 text-sm font-semibold text-gray-800">
          Current Point Structure:
        </div>

        <div className="max-h-60 overflow-y-auto rounded-md border border-gray-200 bg-gray-50 p-3">
          {isLoading ? (
            <div className="text-center text-sm text-gray-500">Loading...</div>
          ) : contributions.length === 0 ? (
            <div className="text-center text-sm text-gray-500">
              No data available.
            </div>
          ) : (
            <ul className="space-y-2">
              {contributions.map(c => (
                <li
                  key={c.id}
                  className="flex items-center justify-between rounded-md bg-white px-3 py-2 text-sm shadow-sm"
                >
                  <span>{c.contributionType}</span>
                  <span className="font-semibold text-orange-500">
                    {c.score} pts
                  </span>
                </li>
              ))}
            </ul>
          )}
        </div>

       
      </DialogContent>
    </Dialog>
  );
};

export default CoinsInfo;
