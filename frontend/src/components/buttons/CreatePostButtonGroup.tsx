
import CreatePostGroup from "../postcreation/CreatePostGroup";



function CreatePostButtonGroup() {
    const openModal = () => {
        const modal = document.getElementById('Modal_Post_Group') as HTMLDialogElement | null;
        if (modal) {
            modal.showModal();
        }
    };

    return (
        <div>
            {/* Open the modal using document.getElementById('ID').showModal() method */}
            <button className="btn btn-primary text-white" onClick={openModal}>Create Post</button>
            <dialog id="Modal_Post_Group" className="modal">
                <div className="modal-box" style={{maxWidth:'none', width: '50%', height: '50%'}}>
                    <h3 className="font-bold text-black text-lg">Group Post Creation</h3>
                    <CreatePostGroup/>
                    
                </div>
                <form method="dialog" className="modal-backdrop">
                    <button>close</button>
                </form>
            </dialog>
        </div>
    )
}

export default CreatePostButtonGroup;
