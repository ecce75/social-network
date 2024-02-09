import React from 'react';

import ProfileInformation from './ProfileInformation';

interface ProfilePageInfoProps {
    title?: string; // New prop for post title
    text?: string;
    pictureUrl?: string;
}

const ProfilePageInfo: React.FC<ProfilePageInfoProps> = ({ title, text, pictureUrl }) => {
    return (
        <div>
            <div style={{ border: '2px solid #ccc', backgroundColor: '#4F7942', borderRadius: '8px', padding: '20px', marginBottom: '20px' }}>
                {/* Group Info*/}
                <ProfileInformation
                    title={title} // Pass title prop to GroupContent
                    text={text}
                    pictureUrl={pictureUrl}
                    placeholderTitle="John Doe"
                    placeholderText="Mis sa nuhid mu profiilil"
                    placeholderPictureUrl="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTHH7mMDHH0S0oWu5HT4FiCTd900_jix22KWhOj6VDlww&s"
                />
            </div>

            
            <div style={{ border: '2px solid #ccc', backgroundColor: '#4F7942', borderRadius: '8px', padding: '10px' }}>
            <h3 style={{ color: 'white', fontWeight:'bold', fontSize: '20px'}}></h3>
            </div>
            
        </div>
    );
};

export default ProfilePageInfo;
