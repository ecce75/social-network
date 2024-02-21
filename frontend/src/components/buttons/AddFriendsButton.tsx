import React, {useEffect, useState} from 'react';
import UserTab from "@/components/friends/UserTab";

interface User {
    id: string;
    username: string;
    avatar_url: string;
}

const AddFriendsButton: React.FC = () => {
    const [users, setUsers] = useState<User[]>([]);

    useEffect(() => {
        fetch('http://localhost:8080/api/users/list', {
            method: 'GET',
            credentials: 'include', // Send cookies with the request
        })
            .then(response => response.json())
            .then(data => setUsers(data as User[]))
            .catch(error => console.error('Error fetching users:', error));
    }, []);

    const handleAddFriend = (userId: string) => {
        fetch(`http://localhost:8080/friends/request/${userId}`, { // Use backticks here
            method: 'POST',
            credentials: 'include', // Send cookies with the request
            headers: {
                'Content-Type': 'application/json',
            },
        })
            .then(response => {
                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }
                // return response.json();
            })
            .then(data => console.log('Friend request sent:', data))
            .catch(error => console.error('Error:', error));
    };


    const openModal = () => {
        const modal = document.getElementById('addFriendsModal');
        if (modal instanceof HTMLDialogElement) {
            modal.showModal();
        }
    };
    console.log(users)

    return (
        <>
            <button className="btn btn-primary text-white" onClick={openModal}>
                Add Friends
            </button>
            <dialog id="addFriendsModal" className="modal">
                <div className="modal-box" style={{maxWidth: 'none', width: '50%', height: '50%', overflowY: 'auto'}}>
                    <h3 className="font-bold text-lg">Choose friends to add</h3>
                    {users.map((user) => (
                        <UserTab
                            key={user.id}
                            userName={user.username}
                            avatarUrl={user.avatar_url}
                            onAddFriend={() => handleAddFriend(user.id)}
                        />
                    ))}
                </div>
                <form method="dialog" className="modal-backdrop">
                    <button>close</button>
                </form>
            </dialog>
        </>
    );
};

export default AddFriendsButton;
