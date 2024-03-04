"use client"

import GroupsNavFeed from "@/components/feeds/GroupNavFeed"




export default function Groups({
                                   params,
                               }: {
    params: {
        id: string
    }
})  {
    return (       
    
    <div>
    {/* Main Content */}
        <h1>Groups {params.id}</h1>
    <main>
        <section>
            <GroupsNavFeed/>
        </section>
    </main>
    
    </div>
    
    )
  }