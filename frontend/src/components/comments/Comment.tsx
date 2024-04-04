import { formatDate } from '@/hooks/utils';
import { useState } from 'react';
import { BiSolidLike, BiSolidDislike } from 'react-icons/bi'; // Import like and dislike icons

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
    const [likeCount, setLikeCount] = useState(likes);
    const [dislikeCount, setDislikeCount] = useState(dislikes);
    const [liked, setLiked] = useState(false);
    const [disliked, setDisliked] = useState(false);

    const handleLike = () => {
        if (!liked) {
            setLikeCount(likeCount + 1);
            setLiked(true);
            if (disliked) {
                setDisliked(false);
                setDislikeCount(dislikeCount - 1);
            }
        } else {
            setLikeCount(likeCount - 1);
            setLiked(false);
        }
        // TODO: Implement logic for sending like to the server
    };

    const handleDislike = () => {
        if (!disliked) {
            setDislikeCount(dislikeCount + 1);
            setDisliked(true);
            if (liked) {
                setLiked(false);
                setLikeCount(likeCount - 1);
            }
        } else {
            setDislikeCount(dislikeCount - 1);
            setDisliked(false);
        }
        // TODO: Implement logic for sending dislike to the server
    };

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
                {image && <img src={image} alt="" style={{ width: "100%", height: 'auto', cursor: 'pointer' }} />}
                <div style={{ marginLeft: "5px", display: "flex" }}>
                    <button onClick={handleLike} style={{ display: "inline-flex", alignItems: "center" }}>
                        <BiSolidLike style={{ color: liked ? "blue" : "black" }} />
                        {likeCount > 0 && <span>{likeCount}</span>}
                    </button>
                    <button onClick={handleDislike} style={{ display: "inline-flex", alignItems: "center", marginLeft: "10px" }}>
                        <BiSolidDislike style={{ color: disliked ? "red" : "black" }} />
                        {dislikeCount > 0 && <span>{dislikeCount}</span>}
                    </button>
                </div>
            </div>
        </div>
    );
}

export default Comment;
