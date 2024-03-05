import React from 'react';

interface GroupInformationProps {
    title?: string; // New prop for post title
    text?: string;
    pictureUrl?: string;
}

const GroupInformation: React.FC<GroupInformationProps> = ({ title, text, pictureUrl }) => {
    const textColor = 'white'; // Set text color to black
    return (
        <div style={{ display: 'flex', alignItems: 'center' }}>
            {pictureUrl && (
                <div className="avatar" style={{ marginRight: '20px' }}>
                    <div className="w-24 sm:w-10 md:b-15 lg:w-24 rounded-full">
                        <img src={pictureUrl} alt="Uploaded Picture" />
                    </div>
                </div>
            )}
            <div>
                {title && <h2 style={{ fontWeight: 'bold', fontSize: '1.2em', color: textColor }}>{title}</h2>} {/* Render the title if provided */}

                {text && <p style={{ marginBottom: '20px', color: textColor }}>{text}</p>}
            </div>
        </div>
    );
};

export default GroupInformation;
