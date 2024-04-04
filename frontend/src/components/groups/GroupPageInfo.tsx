import React, { useEffect } from 'react';
import GroupInformation from './GroupInformation';
import GroupRequestsButton from '../buttons/GroupRequestsButton';
import InviteGroupButton from '../buttons/InviteGroupButton';
import SendGroupRequestButton from "@/components/buttons/SendGroupRequestButton";
import UserTab from "@/components/friends/UserTab";
import AcceptInvitationButton from "@/components/buttons/AcceptinvitationButton";

interface GroupPageInfoProps {
    title?: string; // New prop for post title
    text?: string;
    pictureUrl?: string;
    isMember?: boolean;
    groupId?: string;
    invitationSent?: boolean;
    isCreator?: boolean;
    confirmInvite?: boolean;
}

const GroupPageInfo: React.FC<GroupPageInfoProps> = ({ title, text, pictureUrl, isMember, groupId, invitationSent, isCreator, confirmInvite }) => {
    const BE_PORT = process.env.NEXT_PUBLIC_BACKEND_PORT;
    const FE_URL = process.env.NEXT_PUBLIC_URL;

    const [members, setMembers] = React.useState<{ id: number, username: string, image: string, status: string }[]>([]);
    useEffect(() => {
        try {
            fetch(`${FE_URL}:${BE_PORT}/groups/${groupId}/members`, {
                method: 'GET',
                credentials: 'include'
            })
                .then(response => response.json())
                .then(data => {
                    data.map((user: any) => {
                        const newUser = {
                            id: user.id,
                            username: user.username,
                            image: user.avatar_url,
                            status: user.status
                        }
                        setMembers(prevMembers => [...prevMembers, newUser])
                    }
                    )
                })
        } catch (error) {
            console.log('Error fetching group members:', error)
        }
    }, []);
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
                className={`flex flex-col lg:flex-row ${isMember ? "justify-between" : "justify-center"} border-2 border-gray-300 bg-primary rounded-lg p-5 mb-5`}>
                {/* Invite People */}
                {isMember ? (
                    <>
                        <InviteGroupButton groupID={groupId} />

                        {/* Requests */}
                        {isCreator && (
                            <GroupRequestsButton
                                groupId={groupId} />)}
                    </>) : invitationSent && !confirmInvite ? (
                        < ><p>Join Request sent</p>

                        </>
                    ) : confirmInvite ? (
                        <>
                            <AcceptInvitationButton id={groupId} />
                        </>
                    ) : (
                    < >
                        <SendGroupRequestButton id={groupId} />
                    </>
                )}

            </div>

            <div style={{ border: '2px solid #ccc', backgroundColor: '#4F7942', borderRadius: '8px', padding: '10px' }}>
                <h3 style={{ color: 'white', fontWeight: 'bold', fontSize: '20px' }}>People in Group</h3>
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
                <ul style={{ display: 'flex', flexDirection: 'column', marginBottom: '20px' }}>
                    {members.length > 0 && members.map((user: any) => {
                        return (
                            <UserTab
                                key={user.id}
                                userID={user.id}
                                userName={user.username}
                                avatar={user.image}
                            />
                        )
                    })
                    }

                </ul>
            </div>
        </div>
    );
};

export default GroupPageInfo;
