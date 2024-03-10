import { formatDate } from '@/hooks/utils'
import { useState } from 'react';

export interface CommentProps {
    id: number;
    postId: number;
    userId: number;
    content: string;
    image?: string;
    created_at: string;
    likes: number;
    dislikes: number;
    username: string;
    profile_image?: string;
}

const Comment: React.FC<CommentProps> = ({
    id,
    postId,
    userId,
    content,
    image,
    created_at,
    likes,
    dislikes,
    username,
    profile_image
}) => {
    const formattedCreatedAt = formatDate(created_at);
    const [isModalOpen, setIsModalOpen] = useState(false);
    return (
        <div className="chat chat-start">
            {/* Comments inside the CommentsBox.tsx collapsing box*/}
            <div className="chat-image avatar">
                <div className="w-10 rounded-full">
                    {/* TODO: Link Profile picture to comment */}
                    <img alt="Tailwind CSS chat bubble component" src={profile_image} />
                </div>
            </div>
            <div className="chat-header text-black">
                <span className="mr-2">{username}</span>
                <time className="text-xs text-black">{formattedCreatedAt}</time>
            </div>
            <div className="chat-bubble chat-bubble-secondary">
                {content}
                <div style={{marginLeft: "5px"}}>
                    {image && <img src={image} alt="" style={{ width: "100%", height: 'auto', cursor: 'pointer' }} onClick={() => setIsModalOpen(true)} />}
                </div>
            </div>
            {/* TODO: comment uploaded images */}
            {isModalOpen && (
                <div style={{ position: 'fixed', top: 0, left: 0, width: '100%', height: '100%', backgroundColor: 'rgba(0, 0, 0, 0.5)', display: 'flex', justifyContent: 'center', alignItems: 'center' }} onClick={() => setIsModalOpen(false)}>
                    <img src={image} alt="" style={{ maxHeight: '80%', maxWidth: '80%' }} onClick={e => e.stopPropagation()} />
                </div>
            )}
        </div>
    );
}

export default Comment;