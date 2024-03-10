import React from 'react';
import { formatDate } from '@/hooks/utils'

interface PostInformationProps {
    title?: string; // New prop for post title
    text?: string;
    pictureUrl?: string;
    createdAt: Date;
}

const PostInformation: React.FC<PostInformationProps> = ({ title, text, pictureUrl, createdAt }) => {
    const textColor = 'black'; // Set text color to black
    const createdAtString = formatDate(createdAt?.toString())
    return (
        <div>
            {title && <h2 style={{ fontWeight: 'bold', fontSize: '1.2em', color: textColor }}>{title}</h2>} {/* Render the title if provided */}
            {text && <p style={{ marginBottom: '20px', color: textColor }}>{text}</p>}
            {pictureUrl && <img src={pictureUrl} alt="Uploaded Picture" style={{width:"auto", height: 'auto'}} />}
            {createdAt && <p style={{ color: textColor }}>Posted on: {createdAtString}</p>}
        </div>
    );
};



export default PostInformation;
