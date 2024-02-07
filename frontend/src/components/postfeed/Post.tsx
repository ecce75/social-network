import React from 'react';
import CommentText from '../textareas/CommentText';
import CommentBox from '../textareas/CommentBox';
import PostContent from './PostContent';

interface PostProps {
    text?: string;
    pictureUrl?: string;
}

const Post: React.FC<PostProps> = ({ text, pictureUrl }) => {
    return (
        <div style={{ border: '1px solid #ccc', borderRadius: '8px', padding: '20px', marginBottom: '20px' }}>
            {/* Post Content */}
            <PostContent
                text={text}
                pictureUrl={pictureUrl}
                textColor="black"
                placeholderText="Poor children living in poverty after the war"
                placeholderPictureUrl="https://cdn.pixabay.com/photo/2014/11/13/06/12/boy-529067_960_720.jpg"
            />
            
            {/* Chatbox for commenting and like button */}
            <CommentText />
            
            {/* Comments */}
            <div style={{ marginTop: '20px' }}>
                <CommentBox/>
            </div>
        </div>
    );
};

export default Post;
