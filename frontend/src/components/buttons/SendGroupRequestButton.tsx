import React, {useEffect} from 'react';

interface SendGroupRequestProps {
    id?: string;
}

const SendGroupRequestButton: React.FC<SendGroupRequestProps> = ({id}) => {

    const sendRequest = () => {
        const BE_PORT = process.env.NEXT_PUBLIC_BACKEND_PORT;
        const FE_URL = process.env.NEXT_PUBLIC_FRONTEND_URL;
        fetch(`${FE_URL}:${BE_PORT}/invitations/request/${id}`, {
            method: 'POST',
            credentials: 'include'
        })
            .then(response => response.json())
            .then(data => {
                console.log(data)
                return (
                    <>
                        <p>Group invitation sent</p>
                    </>
                )
            })

    }
    return (
        <div>
            <button className="btn btn-xs sm:btn-sm md:btn-md lg:btn-md btn-secondary text-white" onClick={sendRequest}>Request to join</button>
        </div>
    )
}

export default SendGroupRequestButton;