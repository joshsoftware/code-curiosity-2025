import { useState } from "react";
import { useRepositories } from "@/api/queries/Repositories";
import { Separator } from "@/shared/components/ui/separator";
import RepositoriesCard from "./RepositoriesCard";

const Repositories = () => {
  const { data, isLoading } = useRepositories();
  const repositoriesData = data?.data || [];

  const [search, setSearch] = useState("");
  const [language, setLanguage] = useState("All");
  const [sort, setSort] = useState("latest");

  if (isLoading)
    return (
      <div className="flex h-full w-full items-center justify-center">
        <div className="border-t-cc-app-blue h-12 w-12 animate-spin rounded-full border-4 border-gray-200"></div>
      </div>
    );

  const filteredRepos = repositoriesData
    .filter(
      repo =>
        repo.repoName.toLowerCase().includes(search.toLowerCase()) ||
        repo.description?.toLowerCase().includes(search.toLowerCase()) ||
        repo.languages?.some((lang: string) =>
          lang.toLowerCase().includes(search.toLowerCase())
        )
    )
    .filter(repo =>
      language === "All" ? true : repo.languages?.includes(language)
    )
    .sort((a, b) => {
      if (sort === "latest")
        return (
          new Date(b.updateDate).getTime() - new Date(a.updateDate).getTime()
        );
      if (sort === "oldest")
        return (
          new Date(a.updateDate).getTime() - new Date(b.updateDate).getTime()
        );
      return 0;
    });

  return (
    <div className="mx-auto max-w-4xl space-y-4 p-6">
      <div className="flex flex-col gap-2 md:flex-row md:items-center md:justify-between">
        <input
          type="text"
          placeholder="Find a repository..."
          value={search}
          onChange={e => setSearch(e.target.value)}
          className="border-cc-app-blue text-cc-app-blue placeholder-cc-app-blue focus:border-cc-app-blue focus:ring-cc-app-blue/70 h-10 flex-1 rounded-lg border bg-blue-50 px-4 text-sm shadow-sm transition focus:ring-2 focus:outline-none"
        />

        <div className="flex gap-2 md:ml-4">
          <select
            value={language}
            onChange={e => setLanguage(e.target.value)}
            className="border-cc-app-blue text-cc-app-blue focus:border-cc-app-blue focus:ring-cc-app-blue/70 h-10 rounded-lg border bg-blue-50 px-4 text-sm shadow-sm transition focus:ring-2 focus:outline-none"
          >
            <option value="All">All Languages</option>
            {[...new Set(repositoriesData.flatMap(r => r.languages))].map(
              lang => (
                <option key={lang} value={lang}>
                  {lang}
                </option>
              )
            )}
          </select>

          <select
            value={sort}
            onChange={e => setSort(e.target.value)}
            className="border-cc-app-blue text-cc-app-blue focus:border-cc-app-blue focus:ring-cc-app-blue/70 h-10 rounded-lg border bg-blue-50 px-4 text-sm shadow-sm transition focus:ring-2 focus:outline-none"
          >
            <option value="latest">Latest Updated</option>
            <option value="oldest">Oldest Updated</option>
          </select>
        </div>
      </div>

      {filteredRepos.length === 0 ? (
        <div className="flex flex-col items-center justify-center py-20 text-center">
          <p className="text-lg font-semibold text-gray-700">
            No repositories found
          </p>
          <p className="mt-2 text-sm text-gray-500">
            Try adjusting your search or filters
          </p>
        </div>
      ) : (
        filteredRepos.map(repo => (
          <div key={repo.id}>
            <RepositoriesCard
              id={repo.id}
              name={repo.repoName}
              languages={repo.languages}
              description={repo.description}
              updatedOn={repo.updateDate}
              coins={repo.totalCoinsEarned}
            />
            <Separator />
          </div>
        ))
      )}
    </div>
  );
};

export default Repositories;
