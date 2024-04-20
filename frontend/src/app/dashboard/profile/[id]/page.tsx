"use client"

import { FriendProps, ProfileFeed, ProfileProps } from "@/components/profile/ProfilePage";
import { useEffect, useState } from "react";
import useAuthCheck from "@/hooks/authCheck";


export default function Profile({
    params,
}: {
    params: {
        id: string
    }
}) {
    useAuthCheck();

    const BE_PORT = process.env.NEXT_PUBLIC_BACKEND_PORT;
    const FE_URL = process.env.NEXT_PUBLIC_URL;
    const [friends, setFriends] = useState<FriendProps[]>([]);
    const [profileData, setProfileData] = useState<ProfileProps | null>(null);

    useEffect(() => {
        const fetchData = async () => {
            const url = `${FE_URL}:${BE_PORT}/profile/users/${params.id}`;
            const response = await fetch(url, {
                method: 'GET',
                credentials: 'include',
            });

            if (response.ok) {
                const data: ProfileProps = await response.json();
                setProfileData(data);
            } else {
                console.error('Failed to fetch profile data');
                // You might want to handle this error, maybe set some error state
            }
        };

        fetchData();
    }, [params.id]);

    useEffect(() => {
        fetch(`${FE_URL}:${BE_PORT}/friends/${params.id}`, {
            method: 'GET',
            credentials: 'include' // Send cookies with the request
        })
            .then(response => response.json())
            .then(data => {
                if (data === null) {
                    return;
                }
                setFriends(data)
            })
            .catch(error => console.error('Error fetching friends:', error));
    }, [params.id]);


    return (

        <div>
            <ProfileFeed profile={profileData} friends={friends} userID={params.id} />

        </div>
    );
}
