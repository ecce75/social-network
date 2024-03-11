import { useState } from "react";
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
};

export interface NotificationComponentProps {
    notifications?: NotificationProp[];
    setNotifications: React.Dispatch<React.SetStateAction<NotificationProp[]>>;
    updateNotificationStatus: (notificationId: number, newStatus: any) => void;
}

const NotificationComponent: React.FC<NotificationComponentProps> = ({ notifications = [], setNotifications, updateNotificationStatus }) => {
    const BE_PORT = process.env.NEXT_PUBLIC_BACKEND_PORT;
    const FE_URL = process.env.NEXT_PUBLIC_FRONTEND_URL;
    const [friendStatuses, setFriendStatuses] = useState<FriendStatus>({});
    // Styles
    const containerStyle = {
        backgroundColor: '#f8f9fa', // This should match the light background color of the page
        padding: '20px',
        borderRadius: '8px', // If the design uses rounded corners
        boxShadow: '0 2px 4px rgba(0, 0, 0, 0.1)', // Optional: if there is a shadow in the design
        margin: '10px',
    };

    const notificationStyle = {
        padding: '10px',
        borderBottom: '1px solid #eaeaea', // Use a color that fits with the design
        marginBottom: '10px',
    };

    const messageTypeStyle = {
        fontWeight: 'bold',
        color: '#2a2a2a', // Color for the message type
    };

    const messageStyle = {
        color: '#333', // Regular text color
    };

    const dotStyle: React.CSSProperties = {
        position: 'absolute',
        top: '10px',
        right: '10px',
        height: '10px',
        width: '10px',
        backgroundColor: 'red',
        borderRadius: '50%',
    };


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
            handleDeclineGroupJoinequest(userId);
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
        console.log('THIS SHOULD ACCEPT FRIENT REQUEST VIA SENDERID')
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

    const handleDeclineGroupJoinequest = (senderId: number) => {
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
                setNotifications(notifications?.map(notification => {
                    if (notification.id === id) {
                        return { ...notification, is_read: true };
                    }
                    return notification;
                }));
            } else {
                console.error('Failed to mark notification as read');
            }
        });
    };

    const parentContainerStyle = {
        display: 'flex',
    }

    // Additional styles for buttons
    const buttonStyle = {
        padding: '5px 10px',
        margin: '5px',
        borderRadius: '15px',
        border: 'none',
        cursor: 'pointer',
        fontWeight: 'bold',
    };

    const acceptButtonStyle = {
        ...buttonStyle,
        backgroundColor: '#4CAF50', // Green color for accept
        color: 'white',
    };

    const declineButtonStyle = {
        ...buttonStyle,
        backgroundColor: '#f44336', // Red color for decline
        color: 'white',
    };

    const viewButtonStyle = {
        ...buttonStyle,
        backgroundColor: '#2196F3', // Blue color for view
        color: 'white',
    };

    const notificationActions = (notification: NotificationProp) => {
        // Ensure sender_id is defined before rendering buttons
        if (notification.message && !notification.message.includes('approved') && !notification.message.includes('accepted') && !notification.message.includes("You are now")) {
            switch (notification.type) {
                case 'friend':
                    if (notification.status === 'accepted') {
                        return (
                            <div style={parentContainerStyle}>
                                <p className="createdAtWave">You have accepted the friend request!</p>
                            </div>
                        );
                    } else if (notification.status === 'declined') {
                        return (
                            <div style={parentContainerStyle}>
                                <p className="createdAtWave">You have declined the friend request!</p>
                            </div>
                        );
                    }
                    return (
                        <div style={parentContainerStyle}>
                            <button style={acceptButtonStyle} onClick={() => handleButtonClick(notification.sender_id, 'accept', notification.id)}>
                                Accept
                            </button>
                            <button style={declineButtonStyle} onClick={() => handleButtonClick(notification.sender_id, 'decline', notification.id)}>
                                Decline
                            </button>
                        </div>
                    );
                case 'group':
                    // Similar check and implementation for group join requests
                    if (notification.status === 'accepted') {
                        return (
                            <div style={parentContainerStyle}>
                                <p className="createdAtWave">You have accepted the group request!</p>
                            </div>
                        );
                    } else if (notification.status === 'declined') {
                        return (
                            <div style={parentContainerStyle}>
                                <p className="createdAtWave">You have declined the group request!</p>
                            </div>
                        );
                    }
                    return (
                        <div style={parentContainerStyle}>
                            <button style={acceptButtonStyle} onClick={() => handleButtonClick(notification.sender_id, 'accept-group', notification.id)}>
                                Accept
                            </button>
                            <button style={declineButtonStyle} onClick={() => handleButtonClick(notification.sender_id, 'decline-group', notification.id)}>
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
        <div>
            {notifications.map((notification) => {
                return (
                    <div style={containerStyle} key={notification.id}>
                        <div style={notificationStyle} onClick={() => {
                            if (!notification.is_read) {
                                markNotificationAsRead(notification.id);
                            }
                        }}>
                            <div style={{ position: 'relative' }}>
                                {!notification.is_read && <span style={dotStyle}></span>}
                                <h2 style={messageTypeStyle}>{notificationTypes(notification)}</h2>
                                <p style={messageStyle}>{notification.message}</p>
                                {notificationActions(notification)}
                            </div>
                        </div>
                    </div>
                );
            })}
        </div>
    );
};

export default NotificationComponent;