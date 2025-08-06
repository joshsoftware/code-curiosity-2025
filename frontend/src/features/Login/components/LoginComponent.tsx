import { Button } from "@/shared/components/ui/button";
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader
} from "@/shared/components/ui/card";
import { GITHUB_AUTH_URL } from "@/shared/constants/endpoints";
import Coder from "@/assets/coder.svg";
import { useNavigate, useSearchParams } from "react-router-dom";
import { useEffect } from "react";
import { setAccessToken } from "@/shared/utils/local-storage";
import { useGithubOauthLogin } from "@/api/queries/Auth";
import { toast } from "sonner";
import { useLoggedInUser } from "@/api/queries/UserProfileDetails";
import { ACCOUNT_INFO_PATH } from "@/shared/constants/routes";
import githubIcon from "@/assets/github-white-icon.svg";
const LoginComponent = () => {
  const [searchParams] = useSearchParams();
  const navigate = useNavigate();

  const code = searchParams.get("code");
  const { data, isSuccess, isError } = useGithubOauthLogin(code);

  const { refetch } = useLoggedInUser(false);

  useEffect(() => {
    if (!code) return;

    if (isSuccess && data?.data) {
      const token = data.data;
      setAccessToken(token);

      refetch().then(({ data: userData }) => {
        if (userData?.data?.isBlocked) {
          navigate(ACCOUNT_INFO_PATH);
        }
      });

      navigate("/");
    }

    if (isError) {
      toast.error("OAuth login failed:");
    }
  }, [isSuccess, isError, data, navigate]);

  const handleGithubLogin = () => {
    window.location.href = GITHUB_AUTH_URL || "";
  };

  return (
    <Card className="flex h-[60vh] w-full max-w-md justify-center rounded-2xl bg-white text-center shadow-none">
      <CardHeader className="flex justify-center">
        <img src={Coder} alt="Developer Illustration" className="h-auto w-40" />
      </CardHeader>

      <CardContent className="space-y-6">
        <Button
          onClick={handleGithubLogin}
          className="bg-cc-app-orange hover:bg-cc-app-blue h-10 w-3/4 rounded-md font-semibold text-white"
        >
          <img src={githubIcon} className="h-4 w-4" /> Sign in with GitHub
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
