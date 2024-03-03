"use client"

import {useAuth} from '../../util/utils';
import PostFeed from '@/components/feeds/PostFeed';
import { useRouter } from 'next/navigation';
import { useEffect } from 'react';

export default function Dashboard() {
    const router = useRouter();
    const checkAuth = async () => {
        const auth = await useAuth();
        if (!auth.is_authenticated) {
            router.push('/auth');
        }
    };
    useEffect(() => {
        // Call useAuth immediately on mount
        checkAuth();
        // Set up the interval to call useAuth every 15 minutes
        const intervalId = setInterval(checkAuth, 1800000); // 900000 ms is 15 minutes
        // Cleanup function to clear the interval when the component unmounts
        return () => clearInterval(intervalId);
    }, []); // Empty dependency array to run only once on mount

    return (

        <div>
            {/* Main Content */}
            <main>
            <section>
                <PostFeed />
            </section>
            </main>
            
        </div>
    );
}
