import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation'; // Import useRouter from next/router
import GroupContent from './GroupInformation';
import GroupInformation from './GroupInformation';

interface GroupTabProps {
    id: number;
    creatorId: number;
    title: string;
    description: string;
    image: string;
    createdAt?: Date;
    updatedAt?: Date;
}

const GroupTab: React.FC<GroupTabProps> = ({ title, description, image }) => {
    const BE_PORT = process.env.NEXT_PUBLIC_BACKEND_PORT;
    const FE_URL = process.env.NEXT_PUBLIC_FRONTEND_URL;
    const [groups, setGroups] = useState<GroupTabProps[]>([]); // Initialize groups state as an empty array
    const router = useRouter(); // Initialize useRouter

    useEffect(() => {
        // Fetch groups
        fetch(`${FE_URL}:${BE_PORT}/groups`, {
            method: 'GET',
            credentials: 'include' // Send cookies with the request
        })
            .then(response => response.json())
            .then(data => {
                if (data === null) {
                    return;
                }
                setGroups(data);
                console.log('Groups:', data);
            })
            .catch(error => console.error('Error fetching groups:', error));
    }, [BE_PORT, FE_URL]);

    const handleClick = () => {
        router.push('/dashboard/groups'); // Redirect to groups page
    };

    return (
        <div style={{ border: '2px solid #ccc', backgroundColor: '#4F7942', borderRadius: '8px', padding: '20px', marginBottom: '20px', cursor: 'pointer' }} onClick={handleClick}>
            {/* Group Content */}
            {
                groups.length > 0 ?
                    groups.map(group =>
                        <GroupInformation
                            key={group.id}
                            title={group.title} // Pass title prop to GroupContent
                            text={group.description}
                            pictureUrl={group.image}
                        />
                    )
                    :
                    <div>
                        <p>No groups found</p>
                    </div>
            }
        </div>
    );
};

export default GroupTab;
