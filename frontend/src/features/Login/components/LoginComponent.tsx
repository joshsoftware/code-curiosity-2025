import { Button } from "@/shared/components/ui/button";
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader
} from "@/shared/components/ui/card";
import { GITHUB_AUTH_URL } from "@/shared/constants/endpoints";
import Coder from "@/assets/coder.svg";
import { ACCESS_TOKEN_KEY } from "@/shared/constants/local-storage";

const LoginComponent = () => {
  const handleGithubLogin = () => {
    window.location.href = GITHUB_AUTH_URL || "";
    localStorage.setItem(
      ACCESS_TOKEN_KEY,
      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjIsIklzQWRtaW4iOmZhbHNlLCJleHAiOjE3NTQzNzUwMDF9.FL-pl22Idc5Ge2iEiXTxYx5c1WBH06GZg5AYoonHiuI"
    );
  };

  return (
    <Card className="flex h-[60vh] w-full max-w-md justify-center rounded-2xl bg-white text-center shadow-none">
      <CardHeader className="flex justify-center">
        <img src={Coder} alt="Developer Illustration" className="h-auto w-40" />
      </CardHeader>

      <CardContent className="space-y-6">
        <Button
          onClick={handleGithubLogin}
          className="bg-cc-app-orange h-10 w-3/4 rounded-md font-semibold text-white"
        >
          Login via GitHub
        </Button>
        <hr className="border-gray-200" />
      </CardContent>

      <CardFooter className="flex-col justify-end space-y-6">
        <p className="text-sm text-gray-600">
          No idea where to start? Try this{" "}
          <a
            href="https://www.codetriage.com"
            target="_blank"
            rel="noopener noreferrer"
            className="text-app underline hover:text-blue-800"
          >
            Code Triage
          </a>
        </p>
      </CardFooter>
    </Card>
  );
};

export default LoginComponent;
