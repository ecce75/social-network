import React, { useEffect } from 'react';
import UserTab from "@/components/friends/UserTab";

interface GroupRequestsButtonProps {
    groupId?: string;
}
interface User {
    id: number;
    username: string;
    image: string;
}

const GroupRequestsButton: React.FC<GroupRequestsButtonProps> = ({ groupId }) => {
    const BE_PORT = process.env.NEXT_PUBLIC_BACKEND_PORT;
    const FE_URL = process.env.NEXT_PUBLIC_URL;
    const [users, setUsers] = React.useState<User[]>([]);
    const [groupStatuses, setGroupStatuses] = React.useState<{ [key: string]: 'approved' | 'declined' | 'pending' }>({});

    useEffect(() => {
        try {
            fetch(`${FE_URL}:${BE_PORT}/groups/${groupId}/requests`, {
                method: 'GET',
                credentials: 'include'
            })
                .then(response => response.json())
                .then(data => {
                    if (data != null) {
                        data.map((request: any) => {
                            const newUser: User = {
                                id: request.join_user_id,
                                username: request.username,
                                image: request.image
                            }
                            setGroupStatuses(prevStatuses => ({
                                ...prevStatuses,
                                [request.join_user_id]: request.status
                            }));
                            if (!users.some(user => user.id === newUser.id)) {
                                setUsers([...users, newUser]);
                            }
                        })
                    }
                })
        } catch (error) {
            console.error('Error fetching groups join requests:', error);
        }
    }, []);



    const onAcceptRequest = (userId: number) => {
        try {
            fetch(`${FE_URL}:${BE_PORT}/invitations/approve/${groupId}/${userId}`, {
                method: 'PUT',
                credentials: 'include'
            })
                .then(response => {
                    if (response.status === 200) {
                        setGroupStatuses(prevStatuses => ({
                            ...prevStatuses,
                            [userId]: "approved"
                        }));
                    }
                })

        } catch (error) {
            console.error('Error accepting request:', error);
        }
    }

    const onDeclineRequest = (userId: number) => {
        try {
            fetch(`${FE_URL}:${BE_PORT}/invitations/decline/${groupId}/${userId}`, {
                method: 'PUT',
                credentials: 'include'
            })
                .then(response => {
                    if (response.status === 200) {
                        setGroupStatuses(prevStatuses => ({
                            ...prevStatuses,
                            [userId]: "declined"
                        }));
                    }
                })

        } catch (error) {
            console.error('Error accepting request:', error);
        }
    }


    const openModal = () => {
        const modal = document.getElementById('Modal_Join_Request') as HTMLDialogElement | null;
        if (modal) {
            modal.showModal();
        }
    };

    return (
        <div>
            {/* Open the modal using document.getElementById('ID').showModal() method */}
            <button className={`btn btn-xs sm:btn-sm md:btn-md lg:btn-md btn-secondary text-white `} onClick={openModal}>Requests</button>
            <dialog id="Modal_Join_Request" className="modal">
                <div className="modal-box" style={{ maxWidth: 'none', width: '50%', height: '50%' }}>
                    <h3 className="font-bold text-black text-lg">Incoming join requests</h3>
                    {users.length > 0 ? users.map((user: User) => {
                        return (
                            <UserTab
                                key={user.id}
                                userName={user.username}
                                avatar={user.image}
                                groupStatus={groupStatuses[user.id]}
                                onAcceptRequest={() => { onAcceptRequest(user.id) }}
                                onDeclineRequest={() => { onDeclineRequest(user.id) }}
                            />
                        )
                    }) : (
                        <h2 className="text-xl mt-3 font-semibold text-green-800">No requests yet.</h2>
                    )
                    }

                </div>
                <form method="dialog" className="modal-backdrop">
                    <button>close</button>
                </form>
            </dialog>
        </div>
    )
}

export default GroupRequestsButton;