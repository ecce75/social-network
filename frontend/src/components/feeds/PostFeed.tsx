import React, { useEffect } from 'react';
import Post from '../postcreation/Post';
import CreatePost from '../postcreation/CreatePost';
import CreatePostButton from '../buttons/CreatePostButton';

const PostFeed: React.FC = () => {
    const BE_PORT = process.env.NEXT_PUBLIC_BACKEND_PORT;
    const FE_URL = process.env.NEXT_PUBLIC_FRONTEND_URL;
    useEffect(() => {
        // Fetch posts
        fetch(`${FE_URL}:${BE_PORT}/posts`, {
            method: 'GET',
            credentials: 'include' // Send cookies with the request
        })
            .then(response => response.json())
            .then(data => {
                if (data === null) {
                    return;
                }
            })
            .catch(error => console.error('Error fetching posts:', error));
    }, []);
    return (
        /* Change % for post feed width*/
        <section style={{ width: '45%', margin: 'auto', backgroundColor: '#e5e7eb', padding: '20px', height: '100vh', overflowY: 'auto' }}>
            <div style={{ display: 'flex', flexDirection: 'column' }}>
                {/* Post Creation Form */}
                <div style={{ marginBottom: '20px' }}>
                    <CreatePostButton/>
                </div>
                {/* Posts */}
                <div style={{ marginBottom: '20px' }}>
                    {/* Display posts, need to implement*/}
                    <Post/>
                    <Post/>
                </div>
                {/* News */}
                <div>
                    {/* Display news */}
                </div>
            </div>
        </section>
    );
};

export default PostFeed;
