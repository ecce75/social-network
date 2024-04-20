// hooks/useAuthCheck.js

import { useEffect } from 'react';
import { useRouter } from 'next/navigation';
import { useAuth } from '@/hooks/utils';

const useAuthCheck = () => {
    const router = useRouter();
    const checkAuth = async () => {
        // eslint-disable-next-line react-hooks/rules-of-hooks
        const auth = await useAuth();
        if (!auth.is_authenticated) {
            router.push('/auth');
        }else {
            fetch (`${process.env.NEXT_PUBLIC_URL}:${process.env.NEXT_PUBLIC_BACKEND_PORT}/api/users/auth-update`, {
                method: 'PUT',
                credentials: 'include'
            })
            .then(response => {
                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }
            })
        }
    };

    useEffect(() => {
        checkAuth();
        // Optionally set up a recurring check
        const intervalId = setInterval(checkAuth, 1800000); // 30 minutes
        return () => clearInterval(intervalId);
    }, [router]);
};

export default useAuthCheck;
