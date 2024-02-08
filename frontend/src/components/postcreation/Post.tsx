import React from 'react';
import CommentText from '../commentstuff/CommentText';
import CommentBox from '../commentstuff/CommentBox';
import PostContent from './PostContent';

interface PostProps {
    title?: string; // New prop for post title
    text?: string;
    pictureUrl?: string;
}

const Post: React.FC<PostProps> = ({ title, text, pictureUrl }) => {
    return (
        <div style={{ border: '1px solid #ccc', borderRadius: '8px', padding: '20px', marginBottom: '20px' }}>
            {/* Post Content */}
            <PostContent
                title={title} // Pass title prop to PostContent
                text={text}
                pictureUrl={pictureUrl}
                placeholderTitle="Beepbaapboop war is bad :c"
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
