import { useEffect, useState } from 'react';
import { BiSolidLike, BiSolidDislike } from 'react-icons/bi';

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


const Comment: React.FC<CommentProps> = ({ id, postId, userId, content, image, created_at, likes, dislikes, username, profile_image }) => {
    const BE_PORT = process.env.NEXT_PUBLIC_BACKEND_PORT;
    const FE_URL = process.env.NEXT_PUBLIC_FRONTEND_URL;
    const formattedCreatedAt = created_at;
    const [likeCount, setLikeCount] = useState(likes);
    const [dislikeCount, setDislikeCount] = useState(dislikes);
    const [liked, setLiked] = useState(false);
    const [disliked, setDisliked] = useState(false);

    useEffect(() => {
        const fetchData = async () => {
            try {
                // Fetch initial counts
                const response = await fetch(`${FE_URL}:${BE_PORT}/vote`, {
                    method: 'GET',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    credentials: 'include'
                });
                if (!response.ok) {
                    throw new Error('Failed to fetch initial counts');
                }
                const { likes: initialLikes, dislikes: initialDislikes } = await response.json();
                setLikeCount(initialLikes);
                setDislikeCount(initialDislikes);
                
                // Fetch user's vote action
                const userVoteResponse = await fetch(`${FE_URL}:${BE_PORT}/vote?id=${id}`, {
                    method: 'GET',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    credentials: 'include'
                });
                if (!userVoteResponse.ok) {
                    throw new Error('Failed to fetch user vote');
                }
                const { action } = await userVoteResponse.json();
                if (action === 'like') {
                    setLiked(true);
                } else if (action === 'dislike') {
                    setDisliked(true);
                }
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        };
    
        fetchData();
    }, [id]);
    

    const handleLike = async () => {
        if (!liked) {
            try {
                const response = await fetch(`${FE_URL}:${BE_PORT}/vote`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({
                        item: 'comment',
                        item_id: id,
                        action: 'like',
                    }),
                    credentials: 'include'
                });
                if (!response.ok) {
                    throw new Error('Failed to vote');
                }
                setLikeCount(likeCount + 1); // Update like count after successful vote
                setLiked(true);
                if (disliked) {
                    setDisliked(false);
                    setDislikeCount(dislikeCount - 1);
                }
            } catch (error) {
                console.error('Error voting:', error);
            }
        } else {
            setLikeCount(likeCount - 1);
            setLiked(false);
        }
    };
    
    const handleDislike = async () => {
        if (!disliked) {
            try {
                const response = await fetch(`${FE_URL}:${BE_PORT}/vote`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({
                        item: 'comment',
                        item_id: id,
                        action: 'dislike',
                    }),
                    credentials: 'include'
                });
                if (!response.ok) {
                    throw new Error('Failed to vote');
                }
                setDislikeCount(dislikeCount + 1); // Update dislike count after successful vote
                setDisliked(true);
                if (liked) {
                    setLiked(false);
                    setLikeCount(likeCount - 1);
                }
            } catch (error) {
                console.error('Error voting:', error);
            }
        } else {
            setDislikeCount(dislikeCount - 1);
            setDisliked(false);
        }
    };
    

    return (
        <div className="chat chat-start">
            <div className="chat-image avatar">
                <div className="w-10 rounded-full">
                    <img alt="Profile" src={profile_image} />
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
