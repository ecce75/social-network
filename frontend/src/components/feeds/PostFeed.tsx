import React from 'react';
import Post from '../postcreation/Post';
import CreatePost from '../postcreation/CreatePost';
import CreatePostButton from '../buttons/CreatePostButton';

const PostFeed: React.FC = () => {
    return (
        /* Change % for post feed width*/
        <section style={{ width: '45%', margin: 'auto', backgroundColor: '#e5e7eb', padding: '20px', maxHeight: '110vh', overflowY: 'auto' }}>
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
