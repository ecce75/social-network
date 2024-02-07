import Comments from "./Comments";
function CommentBox (){

return (
        <div className="collapse bg-primary">
        <input type="checkbox" /> 
        {/* Collapsing box that contains comments by other users */}
        <div className="collapse-title text-xl text-black font-medium">
            Comments
        </div>
        <div className="collapse-content"> 
            <Comments/>
        </div>
        </div>
);
}

export default CommentBox;