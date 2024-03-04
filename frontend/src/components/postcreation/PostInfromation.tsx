import React from 'react';

interface PostInformationProps {
    title?: string; // New prop for post title
    text?: string;
    pictureUrl?: string;
}

const PostInformation: React.FC<PostInformationProps> = ({ title, text, pictureUrl }) => {
    const textColor = 'black'; // Set text color to black

    return (
        <div>
            {title && <h2 style={{ fontWeight: 'bold', fontSize: '1.2em', color: textColor }}>{title}</h2>} {/* Render the title if provided */}
            {text && <p style={{ marginBottom: '20px', color: textColor }}>{text}</p>}
            {pictureUrl && <img src={pictureUrl} alt="Uploaded Picture" style={{ maxWidth: '100%', height: 'auto' }} />}
        </div>
    );
};

export default PostInformation;
