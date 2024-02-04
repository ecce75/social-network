"use client"

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
// Your authentication hook or context
import { useAuth } from '../util/utils';
import MainHeader from '@/components/headers/MainHeader';

export default function Home() {
    const router = useRouter();
    useEffect(() => {
        (async () => {
            const auth = await useAuth();
            if (!auth.is_authenticated) {
                router.push('/auth');
            }
        })();
    }, []); // Empty dependency array to run only once on mount

    return (

        <div>
            {/* Header */}
            <header>
                <MainHeader />

                {/* User Profile */}
                <div>
                    <span>John Doe</span>
                </div>

                {/* Navigation */}
                <nav>
                    <ul>
                        <li>Home</li>
                        <li>Friends</li>
                        <li>Messages</li>
                        {/* Add more navigation items */}
                    </ul>
                </nav>
            </header>

            {/* Main Content */}
            <main>
                {/* News Feed */}
                <section>
                    {/* Display posts from friends */}
                </section>

                {/* Sidebar */}
                <aside>
                    {/* Display friend suggestions */}
                </aside>
            </main>

            {/* Footer */}
            <footer>
                {/* Display copyright information */}
            </footer>
        </div>
    );
}
