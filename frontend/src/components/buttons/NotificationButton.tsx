"use client";

import { useEffect, useState } from 'react';
import Notification, { NotificationProp } from '../notification/Notification';

function NotificationButton() {

    const BE_PORT = process.env.NEXT_PUBLIC_BACKEND_PORT;
    const FE_URL = process.env.NEXT_PUBLIC_URL;
    const [notifications, setNotifications] = useState<NotificationProp[]>([]);
    const [showDropdown, setShowDropdown] = useState(false);
    const hasUnread = notifications?.some(notification => !notification.is_read);

    useEffect(() => {
        const fetchNotifications = async () => {
            let url = `${FE_URL}:${BE_PORT}/notifications`;
            const response = await fetch(url, {
                method: 'GET',
                credentials: 'include',
            });
            if (response.ok) {
                const data = await response.json();
                setNotifications(data);
                if (data === null) {
                    setNotifications([]);
                }
            } else {
                console.error('Failed to fetch notifications');
            }
        };
        fetchNotifications();
    }, []);

    const updateNotificationStatus = (notificationId: number, newStatus: any) => {
        setNotifications(prevNotifications =>
            prevNotifications.map(notification =>
                notification.id === notificationId ? { ...notification, status: newStatus } : notification
            )
        );
    };

    const titleStyle = {
        color: '#4a4a4a', // Adjust color to match the design of the page
        marginBottom: '15px',
        marginLeft: '30%',
        fontWeight: 'bold',
    };

    return (
        <div className="relative">
            <button
                className={`btn-2 bg-primary btn-ghost btn-circle ${hasUnread ? 'border-red-500' : ''}`}
                style={{ display: 'flex', justifyContent: 'center', alignItems: 'center' }}
                onClick={() => setShowDropdown(!showDropdown)}
            >
                <svg xmlns="http://www.w3.org/2000/svg" className="h-8 w-8" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9" /></svg>
            </button>
            {showDropdown && (
                <div className="absolute right-0 w-64 mt-2 py-2 border rounded-2xl shadow-xl" style={{ maxHeight: 'calc(3 * 150px)', overflowY: 'auto', backgroundColor: 'rgba(255, 255, 255, 0.6)', zIndex: 1 }}>
                    <h1 style={(titleStyle)}>Notifications</h1>
                    {notifications.slice().reverse().map(notification => (
                        <Notification
                            key={notification.id}
                            notification={notification}
                            setNotifications={setNotifications}
                            updateNotificationStatus={updateNotificationStatus}
                        />
                    ))}
                </div>
            )}
        </div>
    )
}

export default NotificationButton;