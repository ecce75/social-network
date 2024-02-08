"use client"

import MainHeader from '@/components/headers/MainHeader';

import PostFeed from '@/components/feeds/PostFeed';

export default function Dashboard() {

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
