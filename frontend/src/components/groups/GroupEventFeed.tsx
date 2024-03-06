import React from 'react';
import GroupInformation from './GroupInformation';
import JoinRequestsButton from '../buttons/GroupRequestsButton';
import InviteGroupButton from '../buttons/InviteGroupButton';
import CreateEventButton from '../buttons/CreateEventBtn';
import Post from '../postcreation/Post';


interface GroupEventFeedProps {
    title?: string; // New prop for post title
    text?: string;
    pictureUrl?: string;
}

const GroupEventFeed: React.FC<GroupEventFeedProps> = ({ title, text, pictureUrl }) => {
    return (
        <div>
                
                <CreateEventButton/>
            
                <div style={{ border: '2px solid #ccc', backgroundColor: '#4F7942', borderRadius: '8px', padding: '10px', marginTop:'10px' }}>
                <h3 style={{ color: 'white', fontWeight:'bold', fontSize: '20px'}}>Events</h3>
                </div>
            
                {/* Events list */}
                <div style={{ border: '2px solid #ccc', backgroundColor: '#4F7942', borderRadius: '8px', height: '50vh', padding: '20px', marginBottom: '20px', overflowY: 'auto' }}>
                    {/* List */}
                    <ul style={{ display: 'flex', flexDirection: 'column', marginBottom: '20px' }}>
                        {/* Map through the list of events and render each item */}


                    </ul>
                </div>
        </div>
    );
};

export default GroupEventFeed;
