

function JoinRequestsButton() {
    const openModal = () => {
        const modal = document.getElementById('Modal_Join_Request') as HTMLDialogElement | null;
        if (modal) {
            modal.showModal();
        }
    };

    return (
        <div>
            {/* Open the modal using document.getElementById('ID').showModal() method */}
            <button className="btn btn-secondary text-white" onClick={openModal}>Requests</button>
            <dialog id="Modal_Join_Request" className="modal">
                <div className="modal-box" style={{maxWidth:'none', width: '50%', height: '50%'}}>
                    <h3 className="font-bold text-black text-lg">Incoming join requests</h3>
                    
                </div>
                <form method="dialog" className="modal-backdrop">
                    <button>close</button>
                </form>
            </dialog>
        </div>
    )
}

export default JoinRequestsButton;