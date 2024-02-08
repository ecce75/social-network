import React from 'react';
import { useRouter } from 'next/navigation'; // Import useRouter from next/router
import GroupContent from './GroupContent';

interface PostProps {
    title?: string; // New prop for post title
    text?: string;
    pictureUrl?: string;
}

const Group: React.FC<PostProps> = ({ title, text, pictureUrl }) => {
    const router = useRouter(); // Initialize useRouter

    const handleClick = () => {
        router.push('/dashboard/groups/placeholdergroup'); // Redirect to groups page
    };

    return (
        <div style={{ border: '2px solid #ccc', backgroundColor: '#4F7942', borderRadius: '8px', padding: '20px', marginBottom: '20px', cursor: 'pointer' }} onClick={handleClick}>
            {/* Post Content */}
            <GroupContent
                title={title} // Pass title prop to GroupContent
                text={text}
                pictureUrl={pictureUrl}
                placeholderTitle="Shoe Emporium"
                placeholderText="Join us for footwear everything!"
                placeholderPictureUrl="https://iili.io/J1ucEoF.jpg"
            />
        </div>
    );
};

export default Group;
