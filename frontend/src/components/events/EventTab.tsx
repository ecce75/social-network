import {EventProps} from "@/components/groups/GroupEventFeed";
import React from "react";
import {formatDate} from "@/hooks/utils";

const EventTab: React.FC<EventProps> = (event) => {
    const start = formatDate(event.start_time);
    const end = formatDate(event.end_time);

    const openModal = () => {
        const modal = document.getElementById('event_modal') as HTMLDialogElement | null;
        if (modal) {
            modal.showModal();
        }
    };
    return (
        <>
        <div className="card w-full bg-primary text-secondary" onClick={openModal}>
            <div className="card-body p-4 bg-green-50 rounded-md my-3">
                <h2 className="card-title ml-2">{event.title}</h2>
                <small className="text-slate-600 text-xs">From: <time className="text-black text-xs opacity-50">{start}</time></small>
                <small className="text-slate-600 text-xs">Until: <time className="text-black text-xs opacity-50">{end}</time></small>
                <p className="justify-start">{event.description}</p>
                <div className="card-actions justify-center">
                    <button className="btn rounded-full">Attend</button>
                </div>
            </div>
        </div>

           <dialog id="event_modal" className="modal">
    <div className="modal-box" style={{maxWidth:'none', width: '50%', height: '50%'}}>
        <h3 className="text-black text-2xl">Event</h3>
        <h3 className="font-semibold text-black text-4xl py-4">{event.title}</h3>
        <div className="description-box border-2 border-gray-300 p-4 rounded-md">
            <h4 className="text-black text-xl underline">Description</h4>
            <p className="justify-start">{event.description}</p>
        </div>
        <div className="location-box border-2 border-gray-300 p-4 rounded-md mt-4">
            <h4 className="text-black text-xl underline">Location</h4>
            <p className="justify-start">{event.location}</p>
        </div>
        <small className="text-slate-600 text-xl">From: <span className="text-black opacity-50">{start}</span></small>
        <small className="text-slate-600 text-xl">Until: <span className="text-black opacity-50">{end}</span></small>
        <div className="card-actions justify-center mt-4">
            <button className="btn btn-primary rounded-full">Attend</button>
        </div>
    </div>
    <form method="dialog" className="modal-backdrop">
        <button>close</button>
    </form>
</dialog>

    </>
    )
}

export default EventTab;