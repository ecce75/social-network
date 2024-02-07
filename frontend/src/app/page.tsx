"use client"

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
// Your authentication hook or context
import { useAuth } from '../util/utils';
import MainHeader from '@/components/headers/MainHeader';
import LeftNavBar from '@/components/leftnavbar/LeftNavBar'
import background from '../../public/assets/background.png';
import PostFeed from '@/components/feeds/PostFeed';

export default function Home() {
    const router = useRouter();
    useEffect(() => {
        (async () => {
            const auth = await useAuth();
            if (!auth.is_authenticated) {
                router.push('/auth');
            }
            else{
                router.push('/dashboard');
            }
        })();
    }, []); // Empty dependency array to run only once on mount

    return (

        <div>
                <div role="alert" className="alert bg-primary bottom-0 w-full fixed">
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" className="stroke-info shrink-0 w-6 h-6"><path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>
                <span>Redirecting</span>
                <span className="loading loading-spinner loading-lg"></span>
                </div>
            
        </div>
    );
}
