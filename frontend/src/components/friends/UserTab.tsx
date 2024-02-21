import React from 'react';
import { useRouter } from 'next/navigation'; 
import UserInformation from './UserInformation';


interface UserTabProps {
    userName: string;
    avatarUrl: string;
    onAddFriend?: () => void; // Optional prop for the add friend functionality
}

const UserTab: React.FC<UserTabProps> = ({ userName, avatarUrl, onAddFriend  }) => {
    const router = useRouter(); 

    const handleClick = () => {
        router.push('/dashboard/profile/placeholderprofile'); // Redirect to users profile page
    };

    return (
        <div style={{ border: '2px solid #ccc', backgroundColor: '#4F7942', borderRadius: '8px',  marginBottom: '1px', cursor: 'pointer' }} onClick={handleClick}>
            {/* Group Content */}
            <UserInformation
                userName={userName} // Pass title prop to GroupContent
                pictureUrl={avatarUrl}
                placeholderuserName="Mari Tänav"
                placeholderPictureUrl="https://daisyui.com/images/stock/photo-1534528741775-53994a69daeb.jpg"
            />
            {onAddFriend && (
                <button onClick={(e) => {
                    e.stopPropagation(); // Prevents the parent div click event
                    onAddFriend();
                }} className="btn btn-primary">
                    Add Friend
                </button>
            )}
        </div>
    );
};

export default UserTab;