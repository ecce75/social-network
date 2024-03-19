import React, { useState } from "react";
import { FriendStatus } from "../buttons/AddFriendsButton";

export interface NotificationProp {
    id: number;
    user_id: number;
    group_id?: number;
    sender_id?: { Int64: number, Valid: boolean };
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

const Notification: React.FC<NotificationProps> = ({ notification, setNotifications, updateNotificationStatus }) => {
    const BE_PORT = process.env.NEXT_PUBLIC_BACKEND_PORT;
    const FE_URL = process.env.NEXT_PUBLIC_FRONTEND_URL;
    const [friendStatuses, setFriendStatuses] = useState<FriendStatus>({});
    // Styles



    const handleButtonClick = (
        senderId: { Int64: number, Valid: boolean } | undefined,
        action: 'accept' | 'decline' | 'accept-group' | 'decline-group',
        notificationId: number
    ) => {
        // Ensure senderId is valid before proceeding
        if (!senderId?.Valid) {
            console.error('Sender ID is invalid or undefined.');
            return;
        }

        const userId = senderId.Int64;

        // Proceed with the action now that userId is confirmed to be a number
        if (action === 'accept') {
            handleAcceptFriendRequest(userId, notificationId);
        } else if (action === 'decline') {
            handleDeclineFriendRequest(userId);
        } else if (action === 'accept-group') {
            handleAcceptGroupJoinRequest(userId);
        } else if (action === 'decline-group') {
            handleDeclineGroupJoinRequest(userId);
        }
    };

    const handleAcceptFriendRequest = (userId: number, notificationId: number) => { // sender id
        // Implement friend request acceptance logic
        fetch(`${FE_URL}:${BE_PORT}/friends/accept/${userId}`, {
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
                updateNotificationStatus(notificationId, 'accepted');
                console.log('Friend request accepted:', response);
                // Update the friend status for this user
                setFriendStatuses(prevStatuses => ({
                    ...prevStatuses,
                    [userId]: 'accepted',
                }));
            })
            .catch(error => console.error('Error:', error));
    };

    const handleDeclineFriendRequest = (userId: number) => {
        // Implement friend request decline logic
        fetch(`${FE_URL}:${BE_PORT}/friends/decline/${userId}`, {
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
                setFriendStatuses(prevStatuses => ({
                    ...prevStatuses,
                    [userId]: 'declined',
                }));
            })
            .catch(error => console.error('Error:', error));
    };

    const handleAcceptGroupJoinRequest = (senderId: number) => {
        // Implement group join request acceptance logic
        console.log('THIS SHOULD ACCEPT GROUP JOIN REQUEST VIA SENDERID')
    };

    const handleDeclineGroupJoinRequest = (senderId: number) => {
        // Implement group join request decline logic
        console.log('THIS SHOULD DECLINE GROUP JOIN REQUEST VIA SENDERID')
    };

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
                    notification.id === id ? { ...notification, is_read: true } : notification
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
                    if (notification.status === 'accepted') {
                        return (
                            <div className="flex">
                                <p className="createdAtWave">You have accepted the friend request!</p>
                            </div>
                        );
                    } else if (notification.status === 'declined') {
                        return (
                            <div className="flex">
                                <p className="createdAtWave">You have declined the friend request!</p>
                            </div>
                        );
                    }
                    return (
                        <div className="flex">
                            <button className="px-2 py-1 m-1 rounded-lg border-none cursor-pointer font-bold bg-green-500 text-white" onClick={() => handleButtonClick(notification.sender_id, 'accept', notification.id)}>
                                Accept
                            </button>
                            <button className="px-2 py-1 m-1 rounded-lg border-none cursor-pointer font-bold bg-red-500 text-white" onClick={() => handleButtonClick(notification.sender_id, 'decline', notification.id)}>
                                Decline
                            </button>
                        </div>
                    );
                case 'group':
                    // Similar check and implementation for group join requests
                    if (notification.status === 'accepted') {
                        return (
                            <div className="flex">
                                <p className="createdAtWave">You have accepted the group request!</p>
                            </div>
                        );
                    } else if (notification.status === 'declined') {
                        return (
                            <div className="flex">
                                <p className="createdAtWave">You have declined the group request!</p>
                            </div>
                        );
                    }
                    return (
                        <div className="flex">
                            <button className="px-2 py-1 m-1 rounded-lg border-none cursor-pointer font-bold bg-green-500 text-white" onClick={() => handleButtonClick(notification.sender_id, 'accept-group', notification.id)}>
                                Accept
                            </button>
                            <button className="px-2 py-1 m-1 rounded-lg border-none cursor-pointer font-bold bg-red-500 text-white" onClick={() => handleButtonClick(notification.sender_id, 'decline-group', notification.id)}>
                                Decline
                            </button>
                        </div>
                    );
                // Add more cases as needed
                default:
                    return null;
            }
        }


        // Render a fallback or null if sender_id is undefined
        return null;
    };


    const notificationTypes = (notification: NotificationProp) => {
        switch (notification.type) {
            case 'friend':
                return (
                    <p>Friend request</p>
                );
            case 'group':
                return (
                    <p>Group request</p>
                );
            case 'post':
                return (
                    <p>New comment</p>
                );
            default:
                return (
                    <p>General</p>
                );
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