function GroupSearchCreateButton (){

    return (
        <div className="flex justify-between">
        <input type="text" placeholder="Search" className="input input-bordered w-full max-w-xs" />
        <button className="btn">New Group
            <svg xmlns="http://www.w3.org/2000/svg" className="h-8 w-8" fill="none" viewBox="0 0 20 20" stroke="currentColor"><path strokeLinecap="round" strokeLinejoin="round" strokeWidth="3" d="M10 5a1 1 0 011 1v3h3a1 1 0 110 2h-3v3a1 1 0 11-2 0v-3H6a1 1 0 110-2h3V6a1 1 0 011-1z" /></svg>
        </button>
        </div>
    )
}

export default GroupSearchCreateButton;