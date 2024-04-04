import React, {useEffect, useState} from "react";
import {FriendStatus} from "../buttons/AddFriendsButton";
import {useRouter} from "next/navigation";

export interface NotificationProp {
    id: number;
    user_id: number;
    group_id?: number;
    sender_id?: number;
    type: string;
    message: string;
    is_read: boolean;
    created_at: string;
    status?: 'pending' | 'accepted' | 'declined';
}

export interface NotificationProps {
    notification: NotificationProp;
    setNotifications: React.Dispatch<React.SetStateAction<NotificationProp[]>>;
    updateNotificationStatus: (notificationId: number, newStatus: any) => void;
}

const Notification: React.FC<NotificationProps> = ({notification, setNotifications, updateNotificationStatus}) => {
    const router = useRouter();
    const BE_PORT = process.env.NEXT_PUBLIC_BACKEND_PORT;
    const FE_URL = process.env.NEXT_PUBLIC_FRONTEND_URL;
    const [friendStatuses, setFriendStatuses] = useState<FriendStatus>({});
    const [groupStatuses, setGroupStatuses] = useState<FriendStatus>({});
    // Styles

    console.log(notification)

    useEffect(() => {
        console.log(groupStatuses)
    }, [groupStatuses])


    // handles requests for friend requests, group invitations
    const handleRequest = (id: number, request: string, requestType: string) => { // sender id
        // Implement friend request acceptance logic
        fetch(`${FE_URL}:${BE_PORT}/${request}/${requestType}/${id}`, {
            method: 'POST',
            credentials: 'include',
            headers: {
                'Content-Type': 'application/json',
            },
        })
            .then(response => {
                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }
                if (request == 'friends') {
                    // Update the friend status for this user
                    setFriendStatuses(prevStatuses => ({
                        ...prevStatuses,
                        [id]: requestType == 'accept' ? 'accepted' : 'declined',
                    }));
                } else if (request == 'invitations') {
                    // Update the friend status for this user
                    console.log('updating group status')
                    setGroupStatuses(prevStatuses => ({
                        ...prevStatuses,
                        [id]: requestType == 'accept' ? 'accepted' : 'declined',
                    }));
                }
            })
            .catch(error => console.error('Error:', error));
    };

    const handleGroupRequest = (id: number, requestType: string) => {
            fetch(`${FE_URL}:${BE_PORT}/invitations/${notification.group_id}/${id}`, {
            method: 'PUT',
            credentials: 'include',
            headers: {
                'Content-Type': 'application/json',
            },
        })
            .then(response => {
                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }
                // Update the group status for this user
                setGroupStatuses(prevStatuses => ({
                    ...prevStatuses,
                    [id]: requestType == 'approve' ? 'accepted' : 'declined',
                }));
            })
            .catch(error => console.error('Error:', error)); }

    const handleEventClick = () => {
        router.push(`/groups/${notification.group_id}`)
    }


    const markNotificationAsRead = (id: number) => {
        // This function should make a request to the backend to mark the notification as read
        // The request should include the notification ID
        fetch(`${FE_URL}:${BE_PORT}/notifications/${id}`, {
            method: 'PUT',
            credentials: 'include',
        }).then(response => {
            if (response.ok) {
                // If the request is successful, update the state of the notifications to mark the notification as read
                setNotifications(prevNotifications => prevNotifications.map(notification =>
                    notification.id === id ? {...notification, is_read: true} : notification
                ));
            } else {
                console.error('Failed to mark notification as read');
            }
        });
    };


    const notificationActions = (notification: NotificationProp) => {
        // Ensure sender_id is defined before rendering buttons
        if (notification.message && !notification.message.includes('approved') && !notification.message.includes('accepted') && !notification.message.includes("You are now")) {
            switch (notification.type) {
                case 'friend':
                    if (notification.sender_id != undefined && friendStatuses[notification.sender_id] == 'accepted') {
                        return (
                            <div className="flex">
                                <p className="createdAtWave">You have accepted the friend request!</p>
                            </div>
                        );
                    } else if (notification.sender_id != undefined && friendStatuses[notification.sender_id] == 'declined') {
                        return (
                            <div className="flex">
                                <p className="createdAtWave">You have declined the friend request!</p>
                            </div>
                        );
                    } else if (notification.message.includes("sent you")) {
                        return (
                            <div className="flex">
                                <button
                                    className="px-2 py-1 m-1 rounded-lg border-none cursor-pointer font-bold bg-green-500 text-white"
                                    onClick={() => handleRequest(notification.sender_id ? notification.sender_id : 0, 'friends', 'accept')}>
                                    Accept
                                </button>
                                <button
                                    className="px-2 py-1 m-1 rounded-lg border-none cursor-pointer font-bold bg-red-500 text-white"
                                    onClick={() => handleRequest(notification.sender_id ? notification.sender_id : 0, 'friends', 'decline')}>
                                    Decline
                                </button>
                            </div>
                        );
                    }
                case 'group':
                    // Similar check and implementation for group join requests
                    if (notification.message.includes("has requested")) {
                        return (
                            <div className="flex">
                                <button
                                    className="px-2 py-1 m-1 rounded-lg border-none cursor-pointer font-bold bg-green-500 text-white"
                                    onClick={() => handleRequest(notification.sender_id ? notification.sender_id : 0, 'groups', 'accept')}>
                                    Approve
                                </button>
                                <button
                                    className="px-2 py-1 m-1 rounded-lg border-none cursor-pointer font-bold bg-red-500 text-white"
                                    onClick={() => handleRequest(notification.sender_id ? notification.sender_id : 0, 'groups', 'decline')}>
                                    Decline
                                </button>
                            </div>
                        );
                    } else if (notification.message.includes("event")) {
                        return (
                            <div className="flex justify-center mt-1">
                                <button
                                    className="px-2 py-1 m-1 rounded-lg border-none cursor-pointer font-bold bg-slate-400 text-white"
                                    onClick={() => handleEventClick()}>
                                    View Event in Group
                                </button>
                            </div>
                        );
                    } else if (notification.message.includes("have been invited")) {
                        if (notification.user_id != undefined && groupStatuses[notification.user_id] == 'accepted') {
                            return (
                                <div className="flex">
                                    <p className="createdAtWave">You have accepted the group request!</p>
                                </div>
                            );
                        } else if (notification.user_id != undefined && groupStatuses[notification.user_id] == 'declined') {
                            return (
                                <div className="flex">
                                    <p className="createdAtWave">You have declined the group request!</p>
                                </div>
                            );
                        }
                        return (
                            <div className="flex">
                                <button
                                    className="px-2 py-1 m-1 rounded-lg border-none cursor-pointer font-bold bg-green-500 text-white"
                                    onClick={() => handleRequest(notification.group_id ? notification.group_id : 0, 'invitations', 'accept')}>
                                    Accept
                                </button>
                                <button
                                    className="px-2 py-1 m-1 rounded-lg border-none cursor-pointer font-bold bg-red-500 text-white"
                                    onClick={() => handleRequest(notification.group_id ? notification.group_id : 0, 'invitations', 'decline')}>
                                    Decline
                                </button>
                            </div>
                        );
                    }
                // Add more cases as needed
                default:
                    return null;
            }
        }


        // Render a fallback or null if sender_id is undefined
        return null;
    };


    const notificationTypes = (notification: NotificationProp) => {
        console.log(notification.type, notification.message)
        switch (notification.type) {
            default:
                return (
                    <p>General</p>
                );
            case 'friend':
                return (
                    <p>Friend request</p>
                );

            case 'post':
                return (
                    <p>New comment</p>
                );
            case 'group':
                console.log("Entered group case", notification.message)
                if (notification.message.includes("event")) {
                    return (
                        <p>New group event</p>
                    )
                } else {
                    return (
                        <p>Group request</p>
                    )
                }


        }
    };


    return (
        <div className="bg-gray-100 p-5 rounded-lg shadow-sm m-2" key={notification.id}>
            <div className="p-2 border-b border-gray-200 mb-2 cursor-pointer" onClick={() => {
                if (!notification.is_read) {
                    markNotificationAsRead(notification.id);
                }
            }}>
                <div className="relative">
                    {!notification.is_read && <span
                        className="absolute top-2.5 right-2.5 h-2.5 w-2.5 bg-red-500 rounded-full"></span>}
                    <h2 className="font-bold text-gray-800">{notificationTypes(notification)}</h2>
                    <p className="text-gray-800">{notification.message}</p>

                    {notificationActions(notification)}
                </div>
            </div>
        </div>
    );


};

export default Notification;