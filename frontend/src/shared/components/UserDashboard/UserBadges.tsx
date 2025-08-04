import { useUserBadges } from "@/api/queries/UserBadges";
import type { Badge } from "@/shared/types/types";
import bronzeBadge from "@/assets/bronzeBadge.svg";
import silverBadge from "@/assets/silverBadge.svg";
import goldBadge from "@/assets/goldBadge.svg";

const badgeColorMap: Record<string, string> = {
  BEGINNER: bronzeBadge,
  INTERMEDIATE: silverBadge,
  ADVANCED: goldBadge
};

const UserBadges = () => {
  const { data } = useUserBadges();
  const badges = data?.data ?? [];

  const grouped = badges.reduce<Record<string, Badge[]>>((acc, badge) => {
    const type = badge.badgeType.toUpperCase();
    if (!acc[type]) acc[type] = [];
    acc[type].push(badge);
    return acc;
  }, {});

  return (
    <div>
      <p className="text-cc-app-light-blue mb-4 text-left font-semibold">
        BADGES
      </p>
      <div className="flex flex-wrap gap-4">
        {Object.entries(grouped).map(([type, badgeList]) => {
          const badge = badgeColorMap[type] ?? "";
          return (
            <div
              key={type}
              className="group relative flex flex-col items-center"
            >
              <img src={badge} alt="Badge" className="h-10 w-10" />
              {badgeList.length > 1 && (
                <span className="mt-1 text-xs text-white">×{badgeList.length}</span>
              )}
              <div className="absolute bottom-8 z-10 hidden w-max rounded bg- px-2 py-1 text-xs text-white group-hover:block">
                {type} <br />
              </div>
            </div>
          );
        })}
      </div>
    </div>
  );
};

export default UserBadges;
