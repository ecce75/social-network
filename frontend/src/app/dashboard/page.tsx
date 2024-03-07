"use client"

import PostFeed from '@/components/feeds/PostFeed';
import useAuthCheck from "@/hooks/authCheck";

export default function Dashboard() {
    useAuthCheck();
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
