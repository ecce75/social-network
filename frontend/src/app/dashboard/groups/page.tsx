"use client"

import GroupSearchCreateButton from "@/components/buttons/GroupSearchCreateButton";
import GroupTab, { GroupTabProps } from "@/components/groups/GroupTab";
import React, { useEffect, useState } from "react";
import useAuthCheck from "@/hooks/authCheck";

interface GroupProps {
    id: number
    creator_id: number
    title: string
    description: string
    image: string
    created_at: string
    updated_at: string
    is_user_creator: boolean
    is_user_member: boolean
}

export default function Groups() {
    useAuthCheck();
    const BE_PORT = process.env.NEXT_PUBLIC_BACKEND_PORT;
    const FE_URL = process.env.NEXT_PUBLIC_URL;
    const [groups, setGroups] = useState<GroupProps[]>([]); // Initialize groups state as an empty array
    const [myGroups, setMyGroups] = useState<GroupProps[]>([]); // Initialize myGroups state as an empty array

    useEffect(() => {
        // Fetch groups
        fetch(`${FE_URL}:${BE_PORT}/groups`, {
            method: 'GET',
            credentials: 'include' // Send cookies with the request
        })
            .then(response => response.json())
            .then(data => {
                if (data === null) {
                    return;
                }
                // Filter groups where the user is the creator or member
                setMyGroups(data.filter((group: GroupProps) => group.is_user_creator || group.is_user_member));
                // set groups state for all group where user in not the creator or member
                setGroups(data.filter((group: GroupProps) => !group.is_user_creator && !group.is_user_member));
            })
            .catch(error => console.error('Error fetching groups:', error));
    }, []);

    return (
        <div>
            {/* Main Content */}
            <main>
                <div>
                    <div>
                        <section style={{
                            width: '45%',
                            margin: 'auto',
                            backgroundColor: '#e5e7eb',
                            padding: '20px',
                            height: '100vh',
                            overflowY: 'auto'
                        }}>
                            <div style={{ marginBottom: '20px' }}>
                                <GroupSearchCreateButton />
                            </div>
                            <div style={{ display: 'flex', flexDirection: 'column', marginBottom: '20px' }}>
                                {/* Map through the list of groups and render each item */}
                                {/* My Groups Section */}
                                {myGroups.length > 0 ? <h2 className="createdAtWave" style={{ textAlign: 'center', fontSize: '24px', fontWeight: 'bold' }}>My Groups</h2> : null}
                                {myGroups.length > 0 ?
                                    myGroups.map(group =>
                                        <GroupTab
                                            key={group.id}
                                            id={group.id}
                                            creatorId={group.creator_id}
                                            title={group.title}
                                            description={group.description}
                                            image={group.image}
                                        />
                                    )
                                    :
                                    null
                                }
                                {/* All Groups Section */}
                                {groups.length > 0 ? <h2 className="createdAtWave" style={{ textAlign: 'center', fontSize: '24px', fontWeight: 'bold' }}>{myGroups.length > 0 ? 'Other Groups' : 'All Groups'}</h2> : null}
                                {
                                    groups.length > 0 ?
                                        groups.map(group =>
                                            <GroupTab
                                                key={group.id}
                                                id={group.id}
                                                creatorId={group.creator_id}
                                                title={group.title} // Pass title prop to GroupContent
                                                description={group.description}
                                                image={group.image}
                                            />
                                        )
                                        :
                                        <div>
                                            <p>No groups found</p>
                                        </div>
                                }
                            </div>
                        </section>
                    </div>
                </div>
            </main>
        </div>
    )
}