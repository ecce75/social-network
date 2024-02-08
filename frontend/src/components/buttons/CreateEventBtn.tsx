import CreatePost from "../postcreation/CreatePost";

function CreateEventButton() {
    const openModal = () => {
        const modal = document.getElementById('Modal_Create_Event') as HTMLDialogElement | null;
        if (modal) {
            modal.showModal();
        }
    };

    return (
        <div>
            {/* Open the modal using document.getElementById('ID').showModal() method */}
            <button className="btn btn-primary text-white" onClick={openModal}>Create Event</button>
            <dialog id="Modal_Create_Event" className="modal">
                <div className="modal-box" style={{maxWidth:'none', width: '50%', height: '50%'}}>
                    <h3 className="font-bold text-black text-lg">Create an Event</h3>
                    
                </div>
                <form method="dialog" className="modal-backdrop">
                    <button>close</button>
                </form>
            </dialog>
        </div>
    )
}

export default CreateEventButton;