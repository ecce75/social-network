// hooks/useAuthCheck.js

import { useEffect } from 'react';
import { useRouter } from 'next/navigation';
import { UseAuth } from '@/hooks/utils';

const useAuthCheck = () => {
    const router = useRouter();
    const checkAuth = async () => {
        const auth = await UseAuth();
        if (!auth.is_authenticated) {
            console.log('User not authenticated');
            router.push('/auth');
        }else {
            console.log('User authenticated, updating session.');
            fetch (`/api/users/auth-update`, {
                method: 'PUT',
                credentials: 'include'
            })
            .then(response => {
                console.log("auth-update response: " + response);
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
    }, []);
};

export default useAuthCheck;
