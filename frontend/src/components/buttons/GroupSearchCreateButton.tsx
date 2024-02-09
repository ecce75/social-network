import CreateGroup from "../groups/CreateGroup";

function GroupSearchCreateButton (){
    const openModal = () => {
        const modal = document.getElementById('Modal_Create_Group') as HTMLDialogElement | null;
        if (modal) {
            modal.showModal();
        }
    };
    return (
        <div className="flex justify-between">
        <input type="text" placeholder="Search" className="input input-bordered w-full max-w-xs" />


        <div>
            {/* Open the modal using document.getElementById('ID').showModal() method */}
            <button className="btn btn-primary text-white" onClick={openModal}>New Group
            <svg xmlns="http://www.w3.org/2000/svg" className="h-8 w-8" fill="none" viewBox="0 0 20 20" stroke="currentColor"><path strokeLinecap="round" strokeLinejoin="round" strokeWidth="3" d="M10 5a1 1 0 011 1v3h3a1 1 0 110 2h-3v3a1 1 0 11-2 0v-3H6a1 1 0 110-2h3V6a1 1 0 011-1z" /></svg>
            </button>
            <dialog id="Modal_Create_Group" className="modal">
                <div className="modal-box" style={{maxWidth:'none', width: '50%', height: '50%'}}>
                    <h3 className="font-bold text-black text-lg">Create your group</h3>
                    <CreateGroup/>
                </div>
                <form method="dialog" className="modal-backdrop">
                    <button>close</button>
                </form>
            </dialog>
        </div>
        </div>
    )
}

export default GroupSearchCreateButton;