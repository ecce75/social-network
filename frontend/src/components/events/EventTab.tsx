import { EventProps } from "@/components/groups/GroupEventFeed";
import React, { useEffect } from "react";
import { formatDate } from "@/hooks/utils";

interface EventTabProps {
    event: EventProps;
    setEvents: React.Dispatch<React.SetStateAction<EventProps[]>>;
}

const EventTab: React.FC<EventTabProps> = ({ event, setEvents }) => {
    const start = formatDate(event.start_time);
    const end = formatDate(event.end_time);
    const [isAttending, setIsAttending] = React.useState(false);

    console.log('Event:', event);


    useEffect(() => {
        // Check if the current user is attending the event
        if (event.attendance !== null && event.attendance != undefined) {
            const isCurrentUserAttending = event.attendance.some((user) => user.status === 'current_user');
            setIsAttending(isCurrentUserAttending);
            if (isCurrentUserAttending) {
                console.log('User is attending');
            }
        }
    }, []);

    function handleAttend() {
        console.log('Attend button clicked', isAttending);
        try {
            fetch(`${process.env.NEXT_PUBLIC_URL}:${process.env.NEXT_PUBLIC_BACKEND_PORT}/events/${event.id}/${isAttending ? '0' : '1'}`, {
                method: 'PUT',
                credentials: 'include'
            })
                .then(res => res.json())
                .then(data => {
                    console.log('starting fetch', isAttending)
                    setEvents((prevEvents) => {
                        console.log('Prev events', prevEvents)
                        return prevEvents.map((prevEvent) => {
                            console.log("Prevevent", prevEvent)
                            console.log("Event", event)
                            if (prevEvent.id === event.id) {
                                console.log('Event found', isAttending)
                                if (isAttending && prevEvent.attendance != undefined) {
                                    // If the user is currently attending, remove them from the attendance array
                                    console.log('User is attending', prevEvent.attendance.filter((user) => user.id !== data.id))
                                    return {
                                        ...prevEvent,
                                        attendance: prevEvent.attendance.filter((user) => user.id !== data.id)
                                    };
                                } else {
                                    console.log('User is not attending')
                                    // If the user is not currently attending, add them to the attendance array
                                    return {
                                        ...prevEvent,
                                        attendance: prevEvent.attendance ? [...prevEvent.attendance, {
                                            id: data.id,
                                            username: data.username,
                                            avatar_url: data.avatar_url
                                        }] : [{
                                            id: data.id,
                                            username: data.username,
                                            avatar_url: data.avatar_url
                                        }]
                                    };
                                }
                            }
                            console.log('Prev event 2', prevEvent)
                            return prevEvent;
                        })
                    })
                    console.log(setEvents((event) => { console.log(event); return event }))
                    setIsAttending(!isAttending);
                })
        } catch (err) {
            console.error(err)
        }
    }



    const openModal = () => {
        const modal = document.getElementById(event.id) as HTMLDialogElement | null;
        if (modal) {
            modal.showModal();
        }
    };
    return (
        <>
            <div className="card w-full bg-primary text-secondary" onClick={openModal}>
                <div className="card-body p-4 bg-green-50 rounded-md my-3">
                    <h2 className="card-title ml-2">{event.title}</h2>
                    <small className="text-slate-600 text-xs">From: <time
                        className="text-black text-xs opacity-50">{start}</time></small>
                    <small className="text-slate-600 text-xs">Until: <time
                        className="text-black text-xs opacity-50">{end}</time></small>
                    <p className="justify-start">{event.description}</p>
                    <div className="card-actions justify-center">
                        <button className={`btn rounded-full ${!isAttending && 'btn-secondary'}`}
                            onClick={handleAttend}>{isAttending ? 'Not Attend' : 'Attend'}</button>

                    </div>
                </div>
            </div>

            <dialog id={event.id} className="modal">
                <div className="modal-box" style={{ maxWidth: 'none', width: '50%', height: '50%' }}>
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
                    <small className="text-slate-600 text-xl">From: <span
                        className="text-black opacity-50">{start}</span></small>
                    <small className="text-slate-600 text-xl">Until: <span
                        className="text-black opacity-50">{end}</span></small>

                    <div>
                        <h4 className="text-black text-xl underline">Attendees</h4>
                        <div className="flex flex-row">
                            {event.attendance != undefined && event.attendance.map((user) => {
                                return (
                                    <div key={user.id} className="flex flex-col justify-center items-center">
                                        <img src={user.avatar_url} alt={user.username} className="w-12 h-12 rounded-full" />
                                        <p className="text-black text-xs">{user.username}</p>
                                    </div>
                                )
                            })}
                        </div>
                    </div>

                    <div className="card-actions justify-center mt-4">
                        <button className={`btn rounded-full ${!isAttending && 'btn-secondary'}`}
                            onClick={handleAttend}>{isAttending ? 'Not Attend' : 'Attend'}</button>
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