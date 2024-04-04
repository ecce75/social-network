import React from 'react';

interface SendGroupRequestProps {
    id?: string;
}

const SendGroupRequestButton: React.FC<SendGroupRequestProps> = ({ id }) => {
    const [invitationSent, setInvitationSent] = React.useState<boolean>(false);

    const sendRequest = () => {
        const BE_PORT = process.env.NEXT_PUBLIC_BACKEND_PORT;
        const FE_URL = process.env.NEXT_PUBLIC_URL;
        fetch(`${FE_URL}:${BE_PORT}/invitations/request/${id}`, {
            method: 'POST',
            credentials: 'include'
        })
            // .then(response => response.json())
            .then(response => {
                if (response.status === 201) {
                    setInvitationSent(true);
                }
            })

    }
    return (
        <div>
            {!invitationSent ? (
                <button
                    className="btn btn-xs sm:btn-sm md:btn-md lg:btn-md btn-secondary text-white"
                    onClick={sendRequest}
                >
                    Request to join
                </button>
            ) : (
                <button
                    className="btn btn-xs sm:btn-sm md:btn-md lg:btn-md btn-secondary text-white"
                    disabled
                >
                    Join Request sent
                </button>
            )}
        </div>
    )
}

export default SendGroupRequestButton;