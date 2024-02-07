function CommentText() {
    return (
        <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
            {/* Text field for commenting and "Comment" Button */}
            <div style={{ display: 'flex', alignItems: 'center' }}>
                <input type="text" placeholder="Nice dog! ..." className="input input-bordered input-secondary w-full max-w-xs" />
                <button className="btn bg-primary">Comment</button>
            </div>
            {/* Like Button */}
            <button className="btn bg-primary">
                <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" /></svg>
            </button>
        </div>
    );
}

export default CommentText;