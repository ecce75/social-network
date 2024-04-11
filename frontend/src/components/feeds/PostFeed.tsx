import React, { useEffect, useState } from 'react';
import Post from '../postcreation/Post';
import CreatePostButton from '../buttons/CreatePostButton';
import { PostProps } from '../postcreation/Post';
import { CommentProps } from '../comments/Comment';

const PostFeed: React.FC = () => {
    const BE_PORT = process.env.NEXT_PUBLIC_BACKEND_PORT;
    const FE_URL = process.env.NEXT_PUBLIC_URL;
    const [posts, setPosts] = useState<PostProps[]>([]);
    const [comments, setComments] = useState<{ [postId: number]: CommentProps[] }>([]);

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
                setPosts(data);
            })
            .catch(error => console.error('Error fetching posts:', error));
    }, []);

    useEffect(() => {
        posts.forEach(post => {
            fetch(`${FE_URL}:${BE_PORT}/post/${post.id}/comments`, {
                method: 'GET',
                credentials: 'include' // Send cookies with the request
            })
                .then(response => response.json())
                .then(data => {
                    if (data !== null) {
                        setComments(prevComments => ({ ...prevComments, [post.id]: data }));
                    }
                })
                .catch(error => console.error('Error fetching comments: ', error));
        })
    }, [posts]);

    return (
        /* Change % for post feed width*/
        <section style={{ width: '45%', margin: 'auto', backgroundColor: '#e5e7eb', padding: '20px', height: '100vh', overflowY: 'auto' }}>
            <div style={{ display: 'flex', flexDirection: 'column' }}>
                {/* Post Creation Form */}
                <div style={{ marginBottom: '20px' }}>
                    <CreatePostButton />
                </div>
                {/* Posts */}
                <div style={{ marginBottom: '20px' }}>
                    {/* Display posts, need to implement*/}
                    {
                        posts.length > 0 ?
                            posts.map(post =>
                                <Post
                                    key={post.id}
                                    {...post}
                                    comments={comments[post.id]}
                                    setComments={setComments}
                                />
                            )
                            :
                            <div>
                                <p>No posts found</p>
                            </div>
                    }
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
