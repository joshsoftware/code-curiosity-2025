import { type FC } from "react";
import clsx from "clsx";
import type { Language } from "@/shared/types/types";
import { LangColor } from "@/shared/constants/constants";

interface LanguageCardProps {
  title?: string;
  languages: Language[];
  className?: string;
}

const LanguageCard: FC<LanguageCardProps> = ({
  title = "Languages",
  languages,
  className
}) => {
  return (
    <div
      className={clsx(
        "rounded-xl border border-gray-300 bg-white p-4 h-max",
        className
      )}
    >
      <h3 className="mb-3 text-lg font-semibold text-gray-900">{title}</h3>

      <div className="mb-4 flex h-2 w-full overflow-hidden rounded-full bg-gray-100">
        {languages.map((language, index) => {
          console.log(language.name);
          const bgColor = LangColor[language.name] || "bg-gray-400";
          return (
            <div
              key={`${language.name}-${index}`}
              className={`h-full transition-all duration-300 ${bgColor}`}
              style={{ width: `${language.percentage}%` }}
            />
          );
        })}
      </div>

      <div className="grid grid-cols-2 gap-2">
        {languages.map((language, index) => {
          const color = LangColor[language.name] || "bg-gray-400";
          return (
            <div
              key={`${language.name}-legend-${index}`}
              className="flex items-center gap-2"
            >
              <div className={`h-3 w-3 rounded-full ${color}`} />
              <span className="text-sm text-gray-700">
                {language.name} {language.percentage}%
              </span>
            </div>
          );
        })}
      </div>
    </div>
  );
};

export default LanguageCard;
export type { LanguageCardProps };
