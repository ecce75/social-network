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
    image_url?: string;
    privacySetting: string;
    created_at: Date;
    likes: number;
    dislikes: number;
    creator: string;
    creator_avatar?: string;
    comments?: CommentProps[];
    setComments: React.Dispatch<React.SetStateAction<{[postId: number]: CommentProps[]}>>;
}

const Post: React.FC<PostProps> = ({ id, userId, groupId, title, content, image_url, privacySetting, created_at, likes, dislikes, creator, creator_avatar, comments, setComments }) => {
    return (
        <div style={{ border: '1px solid #ccc', borderRadius: '8px', padding: '20px', marginBottom: '20px' }}>
            {/* Post Content */}
            <PostInformation
                title={title}
                text={content}
                pictureUrl={image_url}
                createdAt={created_at}
                creator={creator}
                creator_avatar={creator_avatar}
            />
            
            {/* Chatbox for commenting and like button */}
            <CreateComment 
                postId={id}
                setComments={setComments}
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
