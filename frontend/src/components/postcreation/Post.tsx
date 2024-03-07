import React from 'react';
import PostInformation from './PostInfromation';
import CommentBox from '../comments/CommentBox';
import { CommentProps } from '../comments/Comment';
import CreateComment from '../comments/CreateComment';


export interface PostProps {
    id: number;
    userId: number;
    groupId?: number;
    title: string;
    content?: string;
    imageUrl?: string;
    privacySetting: string;
    createdAt: Date;
    likes: number;
    dislikes: number;
    comments?: CommentProps[];
}

const Post: React.FC<PostProps> = ({ id, userId, groupId, title, content, imageUrl, privacySetting, createdAt, likes, dislikes, comments }) => {
    return (
        <div style={{ border: '1px solid #ccc', borderRadius: '8px', padding: '20px', marginBottom: '20px' }}>
            {/* Post Content */}
            <PostInformation
                title={title}
                text={content}
                pictureUrl={imageUrl}
            />
            
            {/* Chatbox for commenting and like button */}
            <CreateComment 
                postId={id}
            />
            
            {/* Comments */}
            <div style={{ marginTop: '20px' }}>
                <CommentBox
                    comments={comments}
                />
            </div>
        </div>
    );
};

export default Post;
