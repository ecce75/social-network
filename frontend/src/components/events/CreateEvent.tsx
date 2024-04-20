import React, { useState } from "react";
import { EventProps } from "@/components/groups/GroupEventFeed";

interface CreateEventProps {
    groupId: string;
    setEvents: React.Dispatch<React.SetStateAction<EventProps[]>>;
    setIsModalOpen: React.Dispatch<React.SetStateAction<boolean>>
}

const CreateEvent: React.FC<CreateEventProps> = ({ groupId, setEvents, setIsModalOpen }) => {
    const [title, setTitle] = useState('');
    const [description, setDescription] = useState('');
    const [location, setLocation] = useState('');
    const [startDateTime, setStartDateTime] = useState('');
    const [endDateTime, setEndDateTime] = useState('');

    const BE_PORT = process.env.NEXT_PUBLIC_BACKEND_PORT;
    const FE_URL = process.env.NEXT_PUBLIC_URL;

    const handleTitleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setTitle(event.target.value);
    };

    const handleLocationChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setLocation(event.target.value);
    };

    const handleMessageChange = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
        setDescription(event.target.value);
    };

    const handleStartDateTimeChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setStartDateTime(event.target.value);
    };

    const handleEndDateTimeChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setEndDateTime(event.target.value);
    };

    const handleSubmit = (event: React.FormEvent) => {
        event.preventDefault(); // Prevent the form from submitting traditionally

        if (!title || !description || !location || !startDateTime) {
            alert('Please fill in all fields.');
            return;
        }

        const eventRequest: any = {
            title: title,
            description: description,
            start_time: new Date(startDateTime).toISOString(),
            location: location,
            group_id: parseInt(groupId),
        };

        if (endDateTime) {
            eventRequest.end_time = new Date(endDateTime).toISOString();
        }

        fetch(`${FE_URL}:${BE_PORT}/events`, {
            method: 'POST',
            credentials: 'include',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(eventRequest),
        })
            .then(response => {
                if (response.status === 200) {
                    return response.json();
                }
            })
            .then(data => {
                if (data) {
                    // Convert backend response to match EventProps interface
                    const frontendEvent = {
                        ...data,
                        id: data.id.toString(),
                        creator_id: data.creator_id.toString(),
                        group_id: data.group_id.toString(),
                        start_time: new Date(data.start_time).toISOString(),
                        end_time: data.end_time ? new Date(data.end_time).toISOString() : '',
                        created_at: new Date(data.created_at).toISOString(),
                    };

                    setEvents((prevEvents) => [frontendEvent, ...prevEvents]);
                    setIsModalOpen(false);
                }
            })
            .catch((error) => {
                console.error('Error fetching group:', error);
            })
    };

    return (
        <div className="">
            <form onSubmit={handleSubmit} className="space-y-4">
                {/* Event Title */}
                <input
                    type="text"
                    placeholder="Event Title"
                    className="input input-bordered w-full max-w-lg my-4"
                    value={title}
                    onChange={handleTitleChange}
                />

                {/* Event Description */}
                <textarea
                    placeholder="Event Description"
                    className="textarea textarea-bordered w-full h-40 text-black"
                    value={description}
                    onChange={handleMessageChange}
                ></textarea>

                <input
                    type="text"
                    placeholder="Event Location"
                    className="input input-bordered w-full max-w-lg my-4"
                    value={location}
                    onChange={handleLocationChange}
                />

                {/* Event Day/Time */}
                <input
                    type="datetime-local"
                    className="input input-bordered w-full max-w-lg"
                    value={startDateTime}
                    onChange={handleStartDateTimeChange}
                />
                <input
                    type="datetime-local"
                    className="input input-bordered w-full max-w-lg"
                    value={endDateTime}
                    onChange={handleEndDateTimeChange}
                />

                {/* Submit Button */}
                <div className="flex justify-end">
                    <button type="submit" className="btn btn-primary">Create Event</button>
                </div>
            </form>
        </div>
    );
};

export default CreateEvent;