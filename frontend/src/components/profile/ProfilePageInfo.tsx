import React from 'react';

import ProfileInformation from './ProfileInformation';
import UserTab from "@/components/friends/UserTab";

interface ProfilePageInfoProps {
    title?: string; // New prop for post title
    text?: string;
    pictureUrl?: string;
}

const ProfilePageInfo: React.FC<ProfilePageInfoProps> = ({ title, text, pictureUrl}) => {
    return (
        <div>
            <div style={{ border: '2px solid #ccc', backgroundColor: '#4F7942', borderRadius: '8px', padding: '20px', marginBottom: '20px' }}>
                {/* Group Info*/}
                <div style={{display: 'flex', alignItems: 'center'}}>
                    {pictureUrl && (
                        <div className="avatar" style={{marginRight: '20px'}}>
                            <div className="w-24 rounded-full">
                                <img src={pictureUrl} alt="Uploaded Picture"/>
                            </div>
                        </div>
                    )}
                    <div>
                        {title && <h2 style={{
                            fontWeight: 'bold',
                            fontSize: '1.2em',
                            color: "white"
                        }}>{title}</h2>} {/* Render the title if provided */}
                        {text && <p style={{marginBottom: '20px', color: "white"}}>{text}</p>}
                    </div>
                </div>
            </div>



        </div>
    );
};

export default ProfilePageInfo;
