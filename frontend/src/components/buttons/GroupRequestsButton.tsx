import React, {useEffect} from 'react';
import UserTab from "@/components/friends/UserTab";

interface GroupRequestsButtonProps {
    groupId?: string;
}
interface User {
    id: string;
    username: string;
    image : string;
}

const GroupRequestsButton: React.FC<GroupRequestsButtonProps> = ({groupId}) => {
    const BE_PORT = process.env.NEXT_PUBLIC_BACKEND_PORT;
    const FE_URL = process.env.NEXT_PUBLIC_FRONTEND_URL;
    const [users, setUsers] = React.useState<User[]>([]);

    useEffect(() => {
        try {
            fetch(`${FE_URL}:${BE_PORT}/groups/${groupId}/requests`, {
                method: 'GET',
                credentials: 'include'
            })
                .then(response => response.json())
                .then(data => {
                    data.map((request: any) => {
                        const newUser: User = {
                            id: request.id,
                            username: request.username,
                            image: request.image
                        }
                        setUsers([...users, newUser]);
                    })
                    console.log('Group join requests:', data);

                })
        } catch (error) {
            console.error('Error fetching groups join requests:', error);
        }
    },[]);
    useEffect(() => {console.log(users)}, [users]);
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
                <div className="modal-box" style={{maxWidth:'none', width: '50%', height: '50%'}}>
                    <h3 className="font-bold text-black text-lg">Incoming join requests</h3>
                    {users.map((user: User) => {
                        return (
                            <UserTab
                                key={user.id}
                            userName={user.username}
                            avatar={user.image}
                            friendStatus="declined"
                            />
                        )
                    }) }
                    
                </div>
                <form method="dialog" className="modal-backdrop">
                    <button>close</button>
                </form>
            </dialog>
        </div>
    )
}

export default GroupRequestsButton;