import type { FC } from 'react';
import LoginComponent from './loginComponent';
import AuthLayout from '@/shared/components/AuthLayout';

const Login: FC = () => {
  return (
    <AuthLayout>
      <LoginComponent />
    </AuthLayout>
  );
};

export default Login;
