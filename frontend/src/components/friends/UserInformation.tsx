import React from 'react';

interface GroupInformationProps {
    userName?: string; // New prop for post title
    pictureUrl?: string;
    
    placeholderuserName?: string; 
    placeholderPictureUrl?: string;
}

const UserInformation: React.FC<GroupInformationProps> = ({ userName, pictureUrl, placeholderuserName, placeholderPictureUrl }) => {
    const textColor = 'white'; // Set text color to black

    return (
        <div style={{ display: 'flex', alignItems: 'center' }}>
            {pictureUrl && (
                <div className="avatar" style={{ marginRight: '10px' }}>
                    <div className="w-16  rounded-full">
                        <img src={pictureUrl} alt="Uploaded Picture" />
                    </div>
                </div>
            )}
            {!pictureUrl && placeholderPictureUrl && (
                <div className="avatar" style={{ marginRight: '10px' }}>
                    <div className="w-12  rounded-full">
                        <img src={placeholderPictureUrl} alt="Loading" />
                    </div>
                </div>
            )}
            <div>
                {userName && <h2 style={{ fontWeight: 'bold', fontSize: '1.2em', color: textColor }}>{userName}</h2>} {/* Render the Name if provided */}
                {!userName && placeholderuserName && <h2 style={{ fontWeight: 'bold', fontSize: '1.2em', color: textColor }}>{placeholderuserName}</h2>} {/* Render the placeholder name if no title provided */}
                
            </div>
        </div>
    );
};

export default UserInformation;
