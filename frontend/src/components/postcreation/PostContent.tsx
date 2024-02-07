import React from 'react';

interface PostContentProps {
    title?: string; // New prop for post title
    text?: string;
    pictureUrl?: string;
    placeholderText?: string;
    placeholderTitle?: string; // New prop for placeholder title
    placeholderPictureUrl?: string;
}

const PostContent: React.FC<PostContentProps> = ({ title, text, pictureUrl, placeholderText, placeholderTitle, placeholderPictureUrl }) => {
    const textColor = 'black'; // Set text color to black

    return (
        <div>
            {title && <h2 style={{ fontWeight: 'bold', fontSize: '1.2em', color: textColor }}>{title}</h2>} {/* Render the title if provided */}
            {text && <p style={{ marginBottom: '20px', color: textColor }}>{text}</p>}
            {pictureUrl && <img src={pictureUrl} alt="Uploaded Picture" style={{ maxWidth: '100%', height: 'auto' }} />}
            {!title && placeholderTitle && <h2 style={{ fontWeight: 'bold', fontSize: '1.2em', color: textColor }}>{placeholderTitle}</h2>} {/* Render the placeholder title if no title provided */}
            {!text && placeholderText && <p style={{ marginBottom: '20px', color: textColor }}>{placeholderText}</p>}
            {!pictureUrl && placeholderPictureUrl && <img src={placeholderPictureUrl} alt="Placeholder Picture" style={{ maxWidth: '100%', height: 'auto' }} />}
        </div>
    );
};

export default PostContent;
