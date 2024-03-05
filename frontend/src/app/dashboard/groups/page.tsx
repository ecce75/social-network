"use client"

import GroupSearchCreateButton from "@/components/buttons/GroupSearchCreateButton";
import GroupTab from "@/components/groups/GroupTab";
import React from "react";




export default function Groups()  {

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
                    {/*<GroupTab/>*/}
                    {/*<GroupTab/>*/}
                </div>

            </section>
        </div>
    </main>

    </div>

    )
}