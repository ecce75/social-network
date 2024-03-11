import React, { useState } from 'react';
import { formatDate } from '@/hooks/utils'

interface PostInformationProps {
    title?: string; // New prop for post title
    text?: string;
    pictureUrl?: string;
    createdAt: Date;
    creator: string;
    creator_avatar?: string;
}

const PostInformation: React.FC<PostInformationProps> = ({ title, text, pictureUrl, createdAt, creator, creator_avatar }) => {
    const textColor = 'black'; // Set text color to black
    const createdAtColor = 'linear-gradient(to right, red, gold, green)';
    const createdAtString = formatDate(createdAt?.toString())
    const [isModalOpen, setIsModalOpen] = useState(false);
    return (
        <div>
            <div style={{display: 'flex', justifyContent: 'center', alignItems: 'center'}}>
                <img src={creator_avatar} alt="Avatar" style={{ width: '58px', height: '58px', borderRadius: '50%', marginRight: '10px' }} />
                <span style={{ color: textColor }}><b>{creator}</b></span>
                {createdAt && <p className="createdAtWave" style={{ color: createdAtColor, margin: 'auto', marginRight: '8px' }}>Posted on: {createdAtString}</p>}
            </div>
            {title && <h2 style={{ fontWeight: 'bold', fontSize: '1.2em', color: textColor }}>{title}</h2>} {/* Render the title if provided */}
            {text && <p style={{ marginBottom: '20px', color: textColor }}>{text}</p>}
            {pictureUrl && <img src={pictureUrl} alt="Uploaded Picture" style={{ width: "auto", height: 'auto', cursor: 'pointer' }} onClick={() => setIsModalOpen(true)} />}
            {isModalOpen && (
                <div style={{ position: 'fixed', top: 0, left: 0, width: '100%', height: '100%', backgroundColor: 'rgba(0, 0, 0, 0.5)', display: 'flex', justifyContent: 'center', alignItems: 'center' }} onClick={() => setIsModalOpen(false)}>
                    <img src={pictureUrl} alt="" style={{ maxHeight: '80%', maxWidth: '80%' }} onClick={e => e.stopPropagation()} />
                </div>
            )}
        </div>
        
    );
};



export default PostInformation;
