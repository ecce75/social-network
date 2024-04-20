import React from 'react';

interface AcceptInvitationProps {
    id?: string;
}

const AcceptInvitationButton: React.FC<AcceptInvitationProps> = ({ id }) => {
    const [status, setStatus] = React.useState<'accepted' | 'declined' | 'none'>("none");

    const acceptInvitation = () => {
        console.log('Accepting invitation to join group:', id);
        const BE_PORT = process.env.NEXT_PUBLIC_BACKEND_PORT;
        const FE_URL = process.env.NEXT_PUBLIC_URL;
        fetch(`${FE_URL}:${BE_PORT}/invitations/accept/${id}`, {
            method: 'POST',
            credentials: 'include'
        })
            // .then(response => response.json())
            .then(response => {
                if (response.status === 200) {
                    window.location.reload();
                }
            })

    }

    const declineInvitation = () => {
        console.log('Declining invitation to join group:', id);
        const BE_PORT = process.env.NEXT_PUBLIC_BACKEND_PORT;
        const FE_URL = process.env.NEXT_PUBLIC_URL;
        fetch(`${FE_URL}:${BE_PORT}/invitations/decline/${id}`, {
            method: 'POST',
            credentials: 'include'
        })
            // .then(response => response.json())
            .then(response => {
                if (response.status === 200) {
                    setStatus('declined');
                }
            })

    }
    return (
        <div>
            {status == 'none' ? (
                <div className="flex flex-col items-center space-y-4">
                    <h2 className="text-base text-center font-bold text-white">You have been invited to this group</h2>
                    <div className="flex space-x-4">
                        <button
                            className="px-2 py-2 font-semibold text-white text-sm bg-secondary rounded-md hover:bg-green-900"
                            onClick={acceptInvitation}
                        >
                            Accept Invitation
                        </button>
                        <button
                            className="px-2 py-2 font-semibold text-white text-sm bg-secondary rounded-md hover:bg-green-900"
                            onClick={declineInvitation}
                        >
                            Decline Invitation
                        </button>
                    </div>
                </div>

            ) : status == 'accepted' ? (
                <button
                    className="btn btn-xs sm:btn-sm md:btn-md lg:btn-md btn-secondary text-white"
                    disabled
                >
                    Join Request sent
                </button>
            ) : (
                <button
                    className="btn btn-xs sm:btn-sm md:btn-md lg:btn-md btn-secondary text-white"
                    disabled
                >
                    Invitation Declined
                </button>

            )}
        </div>
    )
}

export default AcceptInvitationButton;