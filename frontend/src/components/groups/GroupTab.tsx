import React from 'react';
import { useRouter } from 'next/navigation'; // Import useRouter from next/router
import GroupInformation from './GroupInformation';

export interface GroupTabProps {
    id: number;
    creatorId: number;
    title: string;
    description: string;
    image: string;
    createdAt?: Date;
    updatedAt?: Date;
}

const GroupTab: React.FC<GroupTabProps> = ({ id, creatorId, title, description, image, createdAt, updatedAt }) => {
    const router = useRouter(); // Initialize useRouter

    const handleClick = () => {
        router.push(`/dashboard/groups/${id}`); // Redirect to groups page
    };

    return (
        <div style={{ border: '2px solid #ccc', backgroundColor: '#4F7942', borderRadius: '8px', padding: '20px', marginBottom: '20px', cursor: 'pointer' }} onClick={handleClick}>
            {/* Group Content */}
            {
                <GroupInformation
                            key={id}
                            title={title} // Pass title prop to GroupContent
                            text={description}
                            pictureUrl={image}
                        />
            }
        </div>
    );
};

export default GroupTab;
