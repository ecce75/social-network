import React from 'react';
import { useRouter } from 'next/navigation'; 
import UserInformation from './UserInformation';


interface UserTabProps {
    userName: string;
    avatarUrl: string;
    friendStatus?: 'pending' | 'pending_confirmation' | 'accepted' | 'declined' | 'none'; // Possible friend statuses
    onAddFriend?: () => void; // Optional prop for the add friend functionality
    onAcceptRequest?: () => void; // Optional prop for the Accept request functionality
}

const UserTab: React.FC<UserTabProps> = ({ userName, avatarUrl, friendStatus, onAddFriend, onAcceptRequest  }) => {
    const router = useRouter(); 

    const handleClick = () => {
        router.push('/dashboard/profile/placeholderprofile'); // Redirect to users profile page
    };

    return (
        <div className="flex justify-between items-center border-2 border-gray-300 bg-primary rounded-md mb-1 cursor-pointer" onClick={handleClick}>
            {/* Group Content */}
            <UserInformation
                userName={userName} // Pass title prop to GroupContent
                pictureUrl={avatarUrl}
                // placeholderuserName="Mari Tänav"
                // placeholderPictureUrl="https://daisyui.com/images/stock/photo-1534528741775-53994a69daeb.jpg"
            />
            <div className="flex items-center mr-2">
            {(friendStatus === 'none' || friendStatus === 'declined') && onAddFriend && (
                <button onClick={(e) => {
                    e.stopPropagation(); // Prevents the parent div click event
                    onAddFriend();
                }} className="btn btn-primary">
                    Add Friend
                </button>
            )}
            {friendStatus === 'pending' && (
                <p className="text-xs text-white bg-secondary py-1 px-3 rounded">
                    Friend request sent
                </p>
            )}
            {friendStatus === 'pending_confirmation' && onAcceptRequest && (
                <button onClick={(e) => {
                    e.stopPropagation(); // Prevents the parent div click event
                    onAcceptRequest();
                }} className="btn btn-primary">
                    Accept Request
                </button>
            )}
            {friendStatus === 'accepted' && (
                <p className="text-xs text-white bg-secondary py-1 px-3 rounded">
                    Friend request accepted
                </p>
            )}
            </div>
        </div>
    );
};

export default UserTab;