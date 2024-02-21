import React, { useEffect, useState } from 'react';
import UserTab from "@/components/friends/UserTab";

interface User {
    id: string;
    username: string;
    avatar_url: string;
}

interface FriendStatus {
    [key: string]: 'pending' | 'pending_confirmation' | 'accepted' | 'declined' | 'none'; // Possible friend statuses
}

const AddFriendsButton: React.FC = () => {
    const [users, setUsers] = useState<User[]>([]);
    const [friendStatuses, setFriendStatuses] = useState<FriendStatus>({});

    useEffect(() => {
        fetch('http://localhost:8080/api/users/list', {
            method: 'GET',
            credentials: 'include', // Send cookies with the request
        })
            .then(response => response.json())
            .then(data => {
            if (data != null) {
                setUsers(data as User[]);
                // Fetch friend status for each user
                data.forEach((user: User) => {
                    checkFriendStatus(user.id);
                });
            }
            })
            .catch(error => console.error('Error fetching users:', error));
    }, []);

    const checkFriendStatus = (userId: string) => {
        fetch(`http://localhost:8080/friends/check/${userId}`, {
            method: 'GET',
            credentials: 'include',
        })
            .then(response => response.json())
            .then(data => {
                setFriendStatuses(prevStatuses => ({
                    ...prevStatuses,
                    [userId]: data || 'none', // Set to 'none' if no status is returned
                }));
            })
            .catch(error => console.error('Error checking friend status:', error));
    };

    const handleAddFriend = (userId: string) => {
        fetch(`http://localhost:8080/friends/request/${userId}`, {
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
                // Assuming the server responds with the updated friend status
                return response.json();
            })
            .then(data => {
                console.log('Friend request sent:', data);
                // Update the friend status for this user
                setFriendStatuses(prevStatuses => ({
                    ...prevStatuses,
                    [userId]: data.status,
                }));
            })
            .catch(error => console.error('Error:', error));
    };
    const handleAcceptRequest = (userId: string) => {
         fetch(`http://localhost:8080/friends/accept/${userId}`, {
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
                // Assuming the server responds with the updated friend status
                return response.json();
            })
            .then(data => {
                console.log('Friend request accepted:', data);
                // Update the friend status for this user
                setFriendStatuses(prevStatuses => ({
                    ...prevStatuses,
                    [userId]: 'accepted',
                }));
            })
            .catch(error => console.error('Error:', error));
    };

    const openModal = () => {
        const modal = document.getElementById('addFriendsModal');
        if (modal instanceof HTMLDialogElement) {
            modal.showModal();
        }
    };

    return (
        <>
            <button className="btn btn-secondary text-white mt-2 rounded-xl" onClick={openModal}>
                Add Friends
            </button>
            {users != null ? (
            <dialog id="addFriendsModal" className="modal">
                <div className="modal-box" style={{ maxWidth: 'none', width: '50%', height: '50%', overflowY: 'auto' }}>
                    <h3 className="font-bold text-lg">Choose friends to add</h3>
                    {users.map((user) => (
                        <UserTab
                            key={user.id}
                            userName={user.username}
                            avatarUrl={user.avatar_url}
                            friendStatus={friendStatuses[user.id]}
                            onAddFriend={() => handleAddFriend(user.id)}
                            onAcceptRequest={() => handleAcceptRequest(user.id)}
                        />
                    ))}
                </div>
                <form method="dialog" className="modal-backdrop">
                    <button>close</button>
                </form>
            </dialog>
            ) : (
                <p>No new friends to add</p>
            )}
        </>
    );
};

export default AddFriendsButton;
