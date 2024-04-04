import React, { useEffect } from 'react';
import UserTab from "@/components/friends/UserTab";

interface InviteGroupButtonProps {
    groupID?: string;
}

const InviteGroupButton: React.FC<InviteGroupButtonProps> = ({ groupID }) => {
    const BE_PORT = process.env.NEXT_PUBLIC_BACKEND_PORT;
    const FE_URL = process.env.NEXT_PUBLIC_URL;
    const [nonMembers, setNonMembers] = React.useState<{ id: number, username: string, image: string }[]>([]);
    const [invited, setInvited] = React.useState<{ [key: string]: boolean }>({});

    useEffect(() => {
        try {
            fetch(`${FE_URL}:${BE_PORT}/groups/${groupID}/non-members`, {
                method: 'GET',
                credentials: 'include'
            })
                .then(response => response.json())
                .then(data => data.map((user: any) => {
                    const newUser = {
                        id: user.id,
                        username: user.username,
                        image: user.avatar_url
                    }
                    setNonMembers(prevNonMembers => [...prevNonMembers, newUser])
                }


                ))
        }
        catch (error) {
            console.error('Error fetching non-members:', error);
        }

    }, []);


    const openModal = () => {
        const modal = document.getElementById('Modal_Invite_Group') as HTMLDialogElement | null;
        if (modal) {
            modal.showModal();
        }
    };

    function handleInviteToGroup(id: number) {
        try {
            fetch(`${FE_URL}:${BE_PORT}/invitations/invite/${groupID}/${id}`, {
                method: 'POST',
                credentials: 'include'
            })
                .then(response => {
                    if (response.ok) {
                        setInvited(prevInvited => ({ ...prevInvited, [id]: true }));
                    }
                })

        }
        catch (error) {
            console.error('Error inviting user to group:', error);
        }

    }


    return (
        <div>
            {/* Open the modal using document.getElementById('ID').showModal() method */}
            <button className={`btn btn-xs sm:btn-sm md:btn-md lg:btn-md btn-secondary text-white`} onClick={openModal}>Invite People</button>
            <dialog id="Modal_Invite_Group" className="modal">
                <div className="modal-box" style={{ maxWidth: 'none', width: '50%', height: '50%' }}>
                    <h3 className="font-bold text-black text-lg">Invite people to your group</h3>
                    <div>
                        {nonMembers.length > 0 && nonMembers.map((user: any) => {
                            return (
                                <UserTab key={user.id} userID={user.id} userName={user.username} avatar={user.image} onInviteToGroup={() => { handleInviteToGroup(user.id) }} invitedToGroup={invited[user.id] ? invited[user.id] : false} />

                            )
                        })}
                    </div>
                </div>
                <form method="dialog" className="modal-backdrop">
                    <button>close</button>
                </form>


            </dialog>
        </div>
    )
}

export default InviteGroupButton;