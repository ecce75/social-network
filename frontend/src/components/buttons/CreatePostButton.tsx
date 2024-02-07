import CreatePost from "../postfeed/CreatePost";

function CreatePostButton() {
    const openModal = () => {
        const modal = document.getElementById('my_modal_2') as HTMLDialogElement | null;
        if (modal) {
            modal.showModal();
        }
    };

    return (
        <div>
            {/* Open the modal using document.getElementById('ID').showModal() method */}
            <button className="btn btn-primary text-black" onClick={openModal}>Create Post</button>
            <dialog id="my_modal_2" className="modal">
                <div className="modal-box" style={{maxWidth:'none', width: '50%', height: '50%'}}>
                    <h3 className="font-bold text-black text-lg">Write about your day!</h3>
                    <CreatePost/>
                </div>
                <form method="dialog" className="modal-backdrop">
                    <button>close</button>
                </form>
            </dialog>
        </div>
    )
}

export default CreatePostButton;
