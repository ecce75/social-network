import Comment, { CommentProps } from "./Comment";

interface CommentBoxProps {
    comments?: CommentProps[];
}

const CommentBox: React.FC<CommentBoxProps> = ({ comments }: CommentBoxProps) => {
    // comments?.map(comment => console.log(comment.Id, comment.PostId, comment.UserId, comment.Content, comment.CreatedAt, comment.Likes, comment.Dislikes, comment.Username, comment.ProfilePicture));
    return (
        <div className="collapse bg-primary">
            <input type="checkbox" />
            {/* Collapsing box that contains comments by other users */}
            <div className="collapse-title text-xl text-white font-medium">
                Comments
            </div>
            <div className="collapse-content">
                {
                    comments != undefined && comments.length > 0 ?
                        comments.map(comment =>
                            <Comment
                                key={(comment.postId * 1000) + comment.id}
                                id={comment.id}
                                postId={comment.postId}
                                userId={comment.userId}
                                content={comment.content}
                                createdAt={comment.createdAt}
                                likes={comment.likes}
                                dislikes={comment.dislikes}
                                username={comment.username}
                                profilePicture={comment.profilePicture}
                            />
                        )
                        :
                        <div>
                            <p>No comments yet.</p>
                        </div>
                }
            </div>
        </div>
    );
}

export default CommentBox;