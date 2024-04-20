import React, {useEffect, useState} from "react";
import EventTab from "@/components/events/EventTab";

interface EventWithGroupName {
    id: string;
    creator_id: string;
    group_id: string;
    title: string;
    description: string;
    location: string;
    start_time: string;
    end_time: string;
    created_at: string;
    group_name?: string;
}

function EventsButton() {
    const BE_PORT = process.env.NEXT_PUBLIC_BACKEND_PORT;
    const FE_URL = process.env.NEXT_PUBLIC_URL;

    const [events, setEvents] = useState<EventWithGroupName[]>([]);

    useEffect(() => {
        fetch(`${FE_URL}:${BE_PORT}/events/me`, {
            method: 'GET',
            credentials: 'include',
        }).then(async (response) => {
            if (response.ok) {
                console.log('Events fetched successfully');
                if (response.bodyUsed === true) {
                    const data = await response.json();
                    console.log(data)
                    setEvents(data);
                }
            } else {
                console.error('Failed to fetch events');
            }

        }).catch((error) => {
            console.error('Failed to fetch events', error);
        });
    }, [])

    const groupedEvents = events.reduce((acc: {[key: string]: EventWithGroupName[];}, event) => {
        if (event.group_name) {
            if (!acc[event.group_name]) {
                acc[event.group_name] = [];
            }
            acc[event.group_name].push(event);

        }
        return acc;
    }, {});

    return (
        <div className="flex flex-col">
            <ul className="border-2 border-secondary rounded-lg p-1 flex items-center bg-secondary text-white mt-5" >
                <svg xmlns="http://www.w3.org/2000/svg" className="h-8 w-8" fill="none" viewBox="0 0 20 20"
                     stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="1.5"
                          d="M18 3a1 1 0 00-1.447-.894L8.763 6H5a3 3 0 000 6h.28l1.771 5.316A1 1 0 008 18h1a1 1 0 001-1v-4.382l6.553 3.276A1 1 0 0018 15V3z"/>
                </svg>
                <span>  Your Events</span>
            </ul>
            {/* Iterate over groupedEvents to render each group and its events */}
            {Object.entries(groupedEvents).map(([groupName, events]) => (
                <div key={groupName}>
                    <h3 className="flex text-white font-semibold text-lg mt-2 justify-center">{groupName}</h3> {/* Display the group name */}
                    <div>
                    {events.map((event: any) => (
                            <EventTab key={event.id} event={event} setEvents={setEvents}/> // Render each event using EventTab
                        ))}
                    </div>
                </div>
            ))}


        </div>
    );
}

export default EventsButton;