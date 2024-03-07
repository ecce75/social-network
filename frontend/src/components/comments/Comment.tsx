
export interface CommentProps {
    id: number;
    postId: number;
    userId: number;
    content: string;
    createdAt: string;
    likes: number;
    dislikes: number;
    username: string;
    profilePicture?: string;
}

const Comment: React.FC<CommentProps> = ({ 
        id, 
        postId, 
        userId, 
        content, 
        createdAt, 
        likes, 
        dislikes, 
        username, 
        profilePicture 
    }) => {
    //console.log(Id, PostId, UserId, Content, CreatedAt, Likes, Dislikes, Username, ProfilePicture)
    return (
        <div className="chat chat-start">
            {/* Comments inside the CommentsBox.tsx collapsing box*/}
            <div className="chat-image avatar">
                <div className="w-10 rounded-full">
                    {/* TODO: Link Profile picture to comment */}
                    <img alt="Tailwind CSS chat bubble component" src="https://daisyui.com/images/stock/photo-1534528741775-53994a69daeb.jpg" />
                </div>
            </div>
            <div className="chat-header text-black">
                {/* TODO: Link user name to comment */}
                {username}
                {/* TODO: Link time to comment*/}
                <time className="text-xs text-black ">
                    {createdAt}
                </time>
            </div>
            {/* TODO: Link content to comment */}
            <div className="chat-bubble chat-bubble-secondary">
                {content}
            </div>
            {/* TODO: comment uploaded images */}
            <div></div>
        </div>
    );
}

export default Comment;