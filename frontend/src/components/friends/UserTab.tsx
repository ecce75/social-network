"use client"

import React, {useEffect, useRef, useState} from 'react';
import {useRouter} from 'next/navigation';
import UserInformation from './UserInformation';
import {useChat} from "@/components/chat/ChatContext";


interface UserTabProps {
    userID?: number;
    userName: string;
    avatar: string;
    friendStatus?: 'pending' | 'pending_confirmation' | 'accepted' | 'declined' | 'none'; // Possible friend statuses
    onAddFriend?: () => void; // Optional prop for the add friend functionality
    onAcceptRequest?: () => void; // Optional prop for the Accept request functionality
    onDeclineRequest?: () => void; // Optional prop for the Decline request functionality
}

const UserTab: React.FC<UserTabProps> = ({userID, userName, avatar, friendStatus, onAddFriend, onAcceptRequest, onDeclineRequest}) => {
    const router = useRouter();
    const [showDialog, setShowDialog] = useState(false);
    const [dialogPosition, setDialogPosition] = useState({x: 0, y: 0});
    const dialogRef = useRef<HTMLDivElement>(null);
    const [friendAction, setFriendAction] = useState(0);

    const handleOpenDialog = (e: React.MouseEvent<HTMLDivElement>) => {
        e.stopPropagation();
        const rect = e.currentTarget.getBoundingClientRect();
        setDialogPosition({
            x: e.clientX - (rect?.left ?? 0), // x position within the element
            y: e.clientY - rect?.height ?? 0,  // y position within the element
        });
        setShowDialog(true);
    };

    useEffect(() => {
        const handleClickOutside = (event: MouseEvent) => {
            if (dialogRef.current && !dialogRef.current.contains(event.target as Node)) {
                setShowDialog(false);
            }
        };

        document.addEventListener('mousedown', handleClickOutside);
        return () => {
            document.removeEventListener('mousedown', handleClickOutside);
        };
    }, [dialogRef]);

    const handleViewProfile = (e: React.MouseEvent<HTMLButtonElement>) => {
        e.stopPropagation()
        setShowDialog(false);
        router.push(`/dashboard/profile/${userID}`);
    };



    const {openChat} = useChat();

    const handleChat = (e: React.MouseEvent<HTMLButtonElement>) => {
        e.stopPropagation()
        setShowDialog(false);
        openChat({userID, userName, avatar});
        // Any additional logic for opening a chat
    };


    return (
        <div
            className={`flex justify-between items-center border-2 border-gray-300 bg-primary rounded-md mt-2 ${friendStatus === 'accepted' ? 'cursor-pointer' : ''}`}
            onClick={handleOpenDialog}>
            {/* Group Content */}
            <UserInformation
                userName={userName} // Pass title prop to GroupContent
                pictureUrl={avatar}
                />
            {showDialog && (
                <div

                    ref={dialogRef}
                    className="absolute z-10 p-2 bg-white shadow-lg rounded text-black"
                    style={{
                        top: `${dialogPosition.y}px`,
                        left: `${dialogPosition.x}px`,
                        width: '150px'
                    }} // Adjusted width and text color
                >
                    <ul>
                        <li>
                            <button onClick={(e) => handleViewProfile(e)}
                                    style={{fontSize: '0.875rem', padding: '4px 8px'}} // Smaller font size and padding
                            >View Profile
                            </button>
                        </li>
                        <li>
                            <button onClick={(e) => handleChat(e)}
                                    style={{fontSize: '0.875rem', padding: '4px 8px'}} // Smaller font size and padding
                            >Send Message
                            </button>
                        </li>
                    </ul>
                </div>
            )}
            <div className="flex items-center mr-2">
                {(friendStatus === 'none' || friendStatus === 'declined') && onAddFriend && (
                    <button onClick={(e) => {
                        e.stopPropagation(); // Prevents the parent div click event
                        onAddFriend();
                        setFriendAction(currentValue => currentValue + 1);

                    }} className="btn btn-primary">
                        Add Friend
                    </button>

                )}
                {friendStatus === 'pending' && (
                    <p className="text-xs text-white bg-secondary py-1 px-3 rounded">
                        Friend request sent
                    </p>
                )}
                {friendStatus === 'pending_confirmation' && onAcceptRequest && onDeclineRequest && (
                    <div>
                        <button onClick={(e) => {
                            e.stopPropagation(); // Prevents the parent div click event
                            onAcceptRequest();
                            setFriendAction(currentValue => currentValue + 1);
                        }} className="btn btn-primary">
                            Accept Request
                        </button>
                        <button onClick={(e) => {
                            e.stopPropagation(); // Prevents the parent div click event
                            onDeclineRequest();
                            setFriendAction(currentValue => currentValue + 1);
                        }} className="btn btn-primary">
                            Decline Request
                        </button>
                    </div>
                )}
            </div>
        </div>
    )
        ;
};

export default UserTab;