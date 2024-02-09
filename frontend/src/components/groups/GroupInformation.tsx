import React from 'react';

interface GroupInformationProps {
    title?: string; // New prop for post title
    text?: string;
    pictureUrl?: string;
    placeholderText?: string;
    placeholderTitle?: string; 
    placeholderPictureUrl?: string;
}

const GroupInformation: React.FC<GroupInformationProps> = ({ title, text, pictureUrl, placeholderText, placeholderTitle, placeholderPictureUrl }) => {
    const textColor = 'white'; // Set text color to black

    return (
        <div style={{ display: 'flex', alignItems: 'center' }}>
            {pictureUrl && (
                <div className="avatar" style={{ marginRight: '20px' }}>
                    <div className="w-24 rounded-full">
                        <img src={pictureUrl} alt="Uploaded Picture" />
                    </div>
                </div>
            )}
            {!pictureUrl && placeholderPictureUrl && (
                <div className="avatar" style={{ marginRight: '20px' }}>
                    <div className="w-24 rounded-full">
                        <img src={placeholderPictureUrl} alt="Loading" />
                    </div>
                </div>
            )}
            <div>
                {title && <h2 style={{ fontWeight: 'bold', fontSize: '1.2em', color: textColor }}>{title}</h2>} {/* Render the title if provided */}
                {text && <p style={{ marginBottom: '20px', color: textColor }}>{text}</p>}
                {!title && placeholderTitle && <h2 style={{ fontWeight: 'bold', fontSize: '1.2em', color: textColor }}>{placeholderTitle}</h2>} {/* Render the placeholder title if no title provided */}
                {!text && placeholderText && <p style={{ marginBottom: '20px', color: textColor }}>{placeholderText}</p>}
            </div>
        </div>
    );
};

export default GroupInformation;
