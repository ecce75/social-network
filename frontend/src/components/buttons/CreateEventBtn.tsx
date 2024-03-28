import CreateEvent from "@/components/events/CreateEvent";
import {EventProps} from "@/components/groups/GroupEventFeed";
import React from "react";

function CreateEventButton({groupId, setEvents}: {
    groupId: string,
    setEvents: (value: (((prevState: EventProps[]) => EventProps[]) | EventProps[])) => void
}) {
    const [isModalOpen, setIsModalOpen] = React.useState(false);



    return (
        <div>
            {/* Open the modal using document.getElementById('ID').showModal() method */}
            <button className="btn btn-xs sm:btn-sm md:btn-md lg:btn-md btn-primary text-white"
                    onClick={() =>{setIsModalOpen(true)}}>Create Event
            </button>
            <dialog open={isModalOpen} onClose={() =>{setIsModalOpen(false)}} id="Modal_Create_Event" className="modal">
                <div className="modal-box" style={{maxWidth: 'none', width: '50%', height: '65%'}}>
                    <h3 className="font-bold text-black text-lg">Create an Event</h3>
                    <CreateEvent groupId={groupId}
                                 setEvents={setEvents}
                                 setIsModalOpen={setIsModalOpen}
                    />
                </div>

                <form method="dialog" className="modal-backdrop">
                    <button>close</button>
                </form>
            </dialog>
        </div>
    )
}

export default CreateEventButton;