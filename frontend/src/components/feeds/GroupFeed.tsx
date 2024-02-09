import React from 'react';
import Post from '../postcreation/Post';

import GroupPageInfo from '../groups/GroupPageInfo';
import GroupEventFeed from '../groups/GroupEventFeed';
import CreatePostButtonGroup from '../buttons/CreatePostButtonGroup';

const GroupFeed: React.FC = () => {
    return (
        /* Group page with */
        <div style={{ display: 'flex', justifyContent: 'center' }}> {/* Container for both sections */}


            {/* Left section for displaying group information */}
            <div style={{ flex: '0 0 18%', backgroundColor: '#e5e7eb', padding: '20px', height: '100vh', overflowY: 'auto' }}>
                <GroupPageInfo />
            </div>


            {/* Divider */}
            <div style={{ flex: '0 0 5px', backgroundColor: '#B2BEB5', height: '100vh' }}></div>


            {/* Right section for post feed */}
            <section style={{ flex: '0 0 45%', backgroundColor: '#e5e7eb', padding: '20px', height: '100vh', overflowY: 'auto' }}>
                <div style={{ marginBottom: '20px' }}>
                    <CreatePostButtonGroup/>
                </div>
                <div style={{ display: 'flex', flexDirection: 'column', marginBottom: '20px' }}>
                    <Post/>
                    <Post/>
                </div>
            </section>


            {/* Divider */}
            <div style={{ flex: '0 0 5px', backgroundColor: '#B2BEB5', height: '100vh' }}></div>


            {/* Left section for displaying group information */}
            <div style={{ flex: '0 0 14%', backgroundColor: '#e5e7eb', padding: '20px', height: '100vh', overflowY: 'auto' }}>
                <GroupEventFeed/>
            </div>
        </div>
    );
};

export default GroupFeed;
