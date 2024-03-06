import React from 'react';
import GroupInformation from './GroupInformation';
import JoinRequestsButton from '../buttons/JoinRequestsButton';
import InviteGroupButton from '../buttons/InviteGroupButton';
import SendGroupRequestButton from "@/components/buttons/SendGroupRequestButton";

interface GroupPageInfoProps {
    title?: string; // New prop for post title
    text?: string;
    pictureUrl?: string;
    isMember?: boolean;
    id?: string;
}

const GroupPageInfo: React.FC<GroupPageInfoProps> = ({title, text, pictureUrl, isMember, id}) => {
    return (
        <div>
            <div style={{
                border: '2px solid #ccc',
                backgroundColor: '#4F7942',
                borderRadius: '8px',
                padding: '20px',
                marginBottom: '20px'
            }}>
                {/* Group Info*/}
                <GroupInformation
                    title={title} // Pass title prop to GroupContent
                    text={text}
                    pictureUrl={pictureUrl}
                />
            </div>
            <div
                className={`flex flex-col lg:flex-row ${isMember? "justify-between" : "justify-center"} border-2 border-gray-300 bg-primary rounded-lg p-5 mb-5`}>
                {/* Invite People */}
                {isMember ? (
                        <>
                <InviteGroupButton className="mb-5 md:mb-0 md:mr-5"/>

                {/* Requests */}
                <JoinRequestsButton/>
                    </>): (
                    < >
                        <SendGroupRequestButton id={id}/>
                    </>
                    )}
            </div>

            <div style={{border: '2px solid #ccc', backgroundColor: '#4F7942', borderRadius: '8px', padding: '10px'}}>
                <h3 style={{color: 'white', fontWeight: 'bold', fontSize: '20px'}}>People in Group</h3>
            </div>

            {/* People in group list */}
            <div style={{
                border: '2px solid #ccc',
                backgroundColor: '#4F7942',
                borderRadius: '8px',
                height: '50vh',
                padding: '20px',
                marginBottom: '20px',
                overflowY: 'auto'
            }}>
                {/* List */}
                <ul style={{display: 'flex', flexDirection: 'column', marginBottom: '20px'}}>
                    {/* Map through the list of people and render each item */}
                    {/* <UserTab/>
                    <UserTab/> */}

                </ul>
            </div>
        </div>
    );
};

export default GroupPageInfo;
