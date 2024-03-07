"use client"

import GroupSearchCreateButton from "@/components/buttons/GroupSearchCreateButton";
import GroupTab, { GroupTabProps } from "@/components/groups/GroupTab";
import React, { useEffect, useState } from "react";
import useAuthCheck from "@/hooks/authCheck";




export default function Groups()  {
    useAuthCheck();
    const BE_PORT = process.env.NEXT_PUBLIC_BACKEND_PORT;
    const FE_URL = process.env.NEXT_PUBLIC_FRONTEND_URL;
    const [groups, setGroups] = useState<GroupTabProps[]>([]); // Initialize groups state as an empty array

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
                setGroups(data);

            })
            .catch(error => console.error('Error fetching groups:', error));
    }, [BE_PORT, FE_URL]);

    return (       
    
    <div>
    {/* Main Content */}
    <main>
        <div>
            <div>

            </div>
            <section style={{
                width: '45%',
                margin: 'auto',
                backgroundColor: '#e5e7eb',
                padding: '20px',
                height: '100vh',
                overflowY: 'auto'
            }}>

                <div style={{marginBottom: '20px'}}>
                    <GroupSearchCreateButton/>
                </div>
                <div style={{display: 'flex', flexDirection: 'column', marginBottom: '20px'}}>
                    {/* Map through the list of groups and render each item */}
                    {
                        groups.length > 0 ?
                            groups.map(group =>
                                <GroupTab
                                    key={group.id}
                                    id={group.id}
                                    creatorId={group.creatorId}
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
    </main>

    </div>

    )
}