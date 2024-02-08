import React from 'react';
import GroupSearchCreateButton from '../buttons/GroupSearchCreateButton';
import Group from '../groupcreation/Group';

const GroupsFeed: React.FC = () => {
    return (
        /* Change % for post feed width*/
        <div>
            <div>
                
            </div>
            <section style={{ width: '45%', margin: 'auto', backgroundColor: '#e5e7eb', padding: '20px', height: '100vh', overflowY: 'auto' }}>
                
                <div style={{ marginBottom: '20px' }}>
                <GroupSearchCreateButton/>
                </div>
                <div style={{ display: 'flex', flexDirection: 'column', marginBottom: '20px' }}>
                    <Group/>
                    <Group/>
                </div>

            </section>
        </div>
    );
};

export default GroupsFeed;
