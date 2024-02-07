import React from 'react';

interface PostContentProps {
    text?: string;
    pictureUrl?: string;
    placeholderText?: string;
    placeholderPictureUrl?: string;
    textColor?: string; // New prop for text color
}

const PostContent: React.FC<PostContentProps> = ({ text, pictureUrl, placeholderText, placeholderPictureUrl, textColor }) => {
    return (
        <div>
            {text && <p style={{ marginBottom: '20px', color: textColor }}>{text}</p>}
            {pictureUrl && <img src={pictureUrl} alt="Uploaded Picture" style={{ maxWidth: '100%', height: 'auto' }} />}
            {!text && placeholderText && <p style={{ marginBottom: '20px', color: textColor }}>{placeholderText}</p>}
            {!pictureUrl && placeholderPictureUrl && <img src={placeholderPictureUrl} alt="Placeholder Picture" style={{ maxWidth: '100%', height: 'auto' }} />}
        </div>
    );
};

export default PostContent;
