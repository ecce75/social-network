import React, {useEffect} from 'react';
import CreateEventButton from '../buttons/CreateEventBtn';
import EventTab from "@/components/events/EventTab";

interface User {
    id: string;
    username: string;
    avatar_url: string;
    status?: string;
}

export interface EventProps {
    id: string;
    creator_id: string;
    group_id: string;
    title: string;
    description: string;
    location: string;
    start_time: string;
    end_time: string;
    created_at: string;
    attendance?: User[]
}

interface GroupEventFeedProps {

    groupId: string;
}


const GroupEventFeed: React.FC<GroupEventFeedProps> = ({groupId}) => {
    const BE_PORT = process.env.NEXT_PUBLIC_BACKEND_PORT;
    const FE_URL = process.env.NEXT_PUBLIC_URL;
    const [events, setEvents] = React.useState<EventProps[]>([]);

    useEffect(() => {
        try {
            fetch(`${FE_URL}:${BE_PORT}/events/group/${groupId}`, {
                method: 'GET',
                credentials: 'include'
            })
                .then(response => response.json())
                .then(async (data) => {
                    if (data != null) {
                        // Fetch attendance for each event
                        const eventsWithAttendance = await Promise.all(data.map(async (event: EventProps) => {
                            const attendanceResponse = await fetch(`${FE_URL}:${BE_PORT}/events/attendance/${event.id}`, {
                                method: 'GET',
                                credentials: 'include'
                            });
                            const attendanceData = await attendanceResponse.json();
                            if (attendanceData != null) {
                                return {...event, attendance: attendanceData};
                            }
                            return {...event, attendance: null};
                        }));
                        console.log('Events with attendance:', eventsWithAttendance)
                        setEvents(eventsWithAttendance);
                    }
                })
        } catch (error) {
            console.error('Error:', error);
        }
    }, [])



    return (
        <div>

            <CreateEventButton
                groupId={groupId}
                setEvents={setEvents}
            />

            <div style={{
                border: '2px solid #ccc',
                backgroundColor: '#4F7942',
                borderRadius: '8px',
                padding: '10px',
                marginTop: '10px'
            }}>
                <h3 style={{color: 'white', fontWeight: 'bold', fontSize: '20px'}}>Events</h3>
            </div>

            {/* Events list */}
            <div style={{
                border: '2px solid #ccc',
                backgroundColor: '#4F7942',
                borderRadius: '8px',
                height: '50vh',
                padding: '20px',
                marginBottom: '20px',
                overflowY: 'auto'
            }}>
                {/* List */}
                <ul style={{display: 'flex', flexDirection: 'column', marginBottom: '20px'}}>
                    {/* Map through the list of events and render each item */}
                    {events.length > 0 ? events.map((event) => {
                        return (
                            <EventTab
                                key={event.id}
                                event={event}
                                setEvents={setEvents}
                            />
                        )
                    }) : (
                        <p>Groups has no events</p>
                    )}


                </ul>
            </div>
        </div>
    );
};

export default GroupEventFeed;
