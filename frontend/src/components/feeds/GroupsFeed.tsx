import React from 'react';
import Post from '../postcreation/Post';
import CreatePostButtonGroups from '../buttons/CreatePostButtonGroup';
import GroupSearchCreateButton from '../buttons/GroupSearchCreateButton';


const GroupsFeed: React.FC = () => {
    return (
        /* Change % for post feed width*/
        <section style={{ width: '45%', margin: 'auto', backgroundColor: '#e5e7eb', padding: '20px', maxHeight: '110vh', overflowY: 'auto' }}>
            <div style={{ display: 'flex', flexDirection: 'column' }}>
                <GroupSearchCreateButton/>
                {/* Post Creation Form */}
                <div style={{ marginBottom: '20px' }}>
                    
                </div>
                {/* Posts */}
                <div style={{ marginBottom: '20px' }}>
                    {/* Display Groups*/}
                    
                </div>
            </div>
        </section>
    );
};

export default GroupsFeed;