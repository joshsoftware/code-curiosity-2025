import { Button } from '@/components/ui/button';
import { Card, CardContent, CardFooter, CardHeader } from '@/components/ui/card';
import { type FC } from 'react';
import Coder from '@/assets/coder.svg';
import { env } from '@/utils/endpoints';

const LoginComponent: FC = () => {
  const handleGithubLogin = () => {
    window.location.href = env.GithubAuthUrl || '';
  };

  return (
    <Card className="flex justify-center w-full max-w-md rounded-2xl shadow-none bg-white text-center h-[60vh]">
      <CardHeader className="flex justify-center">
        <img src={Coder} alt="Developer Illustration" className="w-40 h-auto" />
      </CardHeader>

      <CardContent className="space-y-6">
        <Button
          onClick={handleGithubLogin}
          className="w-3/4 h-10 bg-cc-app-orange text-white font-semibold rounded-md"
        >
          Login via GitHub
        </Button>
        <hr className="border-gray-200" />
      </CardContent>

      <CardFooter className="space-y-6 justify-end flex-col">
        <p className="text-sm text-gray-600">
          No idea where to start? Try this{' '}
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
