import React from 'react';
import GroupSearchCreateButton from '../buttons/GroupSearchCreateButton';
import GroupTab from '../groups/GroupTab';

const GroupsNavFeed: React.FC = () => {
    return (
        /* For page groups, Displays groups and create group*/
        <div>
            <div>
                
            </div>
            <section style={{ width: '45%', margin: 'auto', backgroundColor: '#e5e7eb', padding: '20px', height: '100vh', overflowY: 'auto' }}>
                
                <div style={{ marginBottom: '20px' }}>
                <GroupSearchCreateButton/>
                </div>
                <div style={{ display: 'flex', flexDirection: 'column', marginBottom: '20px' }}>
                    <GroupTab/>
                    <GroupTab/>
                </div>

            </section>
        </div>
    );
};

export default GroupsNavFeed;
