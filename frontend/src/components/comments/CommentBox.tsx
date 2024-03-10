import Comment, { CommentProps } from "./Comment";

interface CommentBoxProps {
    comments?: CommentProps[];
}

const CommentBox: React.FC<CommentBoxProps> = ({ comments }: CommentBoxProps) => {
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
                                key={comment.id}
                                id={comment.id}
                                postId={comment.postId}
                                userId={comment.userId}
                                content={comment.content}
                                image={comment.image}
                                created_at={comment.created_at}
                                likes={comment.likes}
                                dislikes={comment.dislikes}
                                username={comment.username}
                                profile_image={comment.profile_image}
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