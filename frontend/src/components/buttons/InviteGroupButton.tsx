

function InviteGroupButton() {
    const openModal = () => {
        const modal = document.getElementById('Modal_Invite_Group') as HTMLDialogElement | null;
        if (modal) {
            modal.showModal();
        }
    };

    return (
        <div>
            {/* Open the modal using document.getElementById('ID').showModal() method */}
            <button className="btn btn-secondary text-white sm:btn-sm md:btn-md lg:btn-md" onClick={openModal}>Invite People</button>
            <dialog id="Modal_Invite_Group" className="modal">
                <div className="modal-box" style={{maxWidth:'none', width: '50%', height: '50%'}}>
                    <h3 className="font-bold text-black text-lg">Invite people to your group</h3>
                    
                </div>
                <form method="dialog" className="modal-backdrop">
                    <button>close</button>
                </form>
            </dialog>
        </div>
    )
}

export default InviteGroupButton;