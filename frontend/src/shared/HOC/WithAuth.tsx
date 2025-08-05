import { type FC, type ReactNode, useEffect } from "react";
import { Navigate, useLocation, useNavigate } from "react-router-dom";

import { LOGIN_PATH } from "@/shared/constants/routes";
import { getAccessToken } from "@/shared/utils/local-storage";

interface WithAuthProps {
  children: ReactNode;
}

const WithAuth: FC<WithAuthProps> = ({ children }) => {
  const navigate = useNavigate();
  const location = useLocation();
  const userAccessToken = getAccessToken();

  useEffect(() => {
    if (!userAccessToken) {
      navigate(LOGIN_PATH, { replace: true });
    }
  }, [userAccessToken, location.pathname, navigate]);

  if (!userAccessToken) {
    return <Navigate to={LOGIN_PATH} replace />;
  }

  return <>{children}</>;
};

export default WithAuth;




// import { useEffect } from 'react'
// import { useNavigate, useSearchParams } from 'react-router-dom'

// export default function AuthCallback() {
//   const [searchParams] = useSearchParams()
//   const navigate = useNavigate()

//   useEffect(() => {
//     const code = searchParams.get('code')
//     if (!code) return

//     // Call backend to exchange code for token
//     fetch(`http://localhost:8080/github/callback?code=${code}`)
//       .then(res => res.json())
//       .then(data => {
//         const token = data.accessToken
//         if (token) {
//           // Store token in localStorage (or sessionStorage)
//           localStorage.setItem('accessToken', token)

//           // Navigate to home or dashboard
//           navigate('/dashboard')
//         } else {
//           console.error('No token received')
//         }
//       })
//       .catch(err => {
//         console.error('GitHub login failed:', err)
//       })
//   }, [])

//   return <div>Logging in...</div>
// }