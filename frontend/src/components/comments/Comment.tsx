import { formatDate } from '@/hooks/utils'

export interface CommentProps {
    id: number;
    postId: number;
    userId: number;
    content: string;
    created_at: string;
    likes: number;
    dislikes: number;
    username: string;
    image?: string;
}

const Comment: React.FC<CommentProps> = ({
    id,
    postId,
    userId,
    content,
    created_at,
    likes,
    dislikes,
    username,
    image
}) => {
    const formattedCreatedAt = formatDate(created_at)

    return (
        <div className="chat chat-start">
            {/* Comments inside the CommentsBox.tsx collapsing box*/}
            <div className="chat-image avatar">
                <div className="w-10 rounded-full">
                    {/* TODO: Link Profile picture to comment */}
                    <img alt="Tailwind CSS chat bubble component" src={image} />
                </div>
            </div>
            <div className="chat-header text-black">
                <span className="mr-2">{username}</span>
                <time className="text-xs text-black">{formattedCreatedAt}</time>
            </div>
            <div className="chat-bubble chat-bubble-secondary">
                {content}
            </div>
            {/* TODO: comment uploaded images */}
            <div></div>
        </div>
    );
}

export default Comment;