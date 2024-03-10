
export interface NotificationProp {
    id: number;
    user_id: number;
    group_id?: number;
    sender_id?: number;
    type: string;
    message: string;
    is_read: boolean;
    created_at: string;
};

export interface NotificationComponentProps {
    notifications?: NotificationProp[];
    setNotifications: React.Dispatch<React.SetStateAction<NotificationProp[]>>;
}

const NotificationComponent: React.FC<NotificationComponentProps> = ({ notifications = [], setNotifications }) => {
    const BE_PORT = process.env.NEXT_PUBLIC_BACKEND_PORT;
    const FE_URL = process.env.NEXT_PUBLIC_FRONTEND_URL;
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
                console.log('Notification with id: '+ id +' marked as read')
            } else {
                console.error('Failed to mark notification as read');
            }
        });
    };

    return (
        <div>
            {notifications?.map((notification: NotificationProp) => {
                return (
                    <div style={containerStyle} key={notification.id}>
                        <div style={notificationStyle} onClick={() => markNotificationAsRead(notification.id)}>
                            <div style={{ position: 'relative' }}>
                                {!notification.is_read && <span style={dotStyle}></span>}
                                <h2 style={messageTypeStyle}>{notification.type}</h2>
                                <p style={messageStyle}>{notification.message}</p>
                            </div>
                        </div>
                    </div>
                );
            })}
        </div>
    );
};

export default NotificationComponent;