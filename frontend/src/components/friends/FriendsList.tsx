"use client"
import { useState, useEffect } from "react";
import UserTab from "./UserTab";

interface FriendProps {
    id: number;
    firstName: string;
    lastName: string;
    avatar_url: string;
    username: string;
}

export default function FriendsList() {
    const BE_PORT = process.env.NEXT_PUBLIC_BACKEND_PORT;
    const FE_URL = process.env.NEXT_PUBLIC_URL;
    const [friends, setFriends] = useState<FriendProps[]>([]);

    useEffect(() => {
        // fetch(`${process.env.FRONTEND_URL}:${process.env.BACKEND_PORT}/friends`, {
        fetch(`${FE_URL}:${BE_PORT}/friends/me`, {
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
    }, []);


    return (
        <>
            {
                friends.length > 0 ? 
                friends.map(friend =>
                    <UserTab
                        key={friend.id}
                        userID={friend.id}
                        userName={friend.username}
                        friendStatus={'accepted'}
                        avatar={friend.avatar_url}
                    />
                    // <FriendsListContent
                    //     key={friend.id}
                    //     id={friend.id}
                    //     firstName={friend.firstName}
                    //     lastName={friend.lastName}
                    //     avatar={friend.avatar}
                    //     username={friend.username}
                    // />
                )
                :
                //TODO; Add a button to add friends
                <div>
                    <p>No friends found</p>
                </div>
            }
        </>
    );
}

// const FriendsListContent: React.FC<FriendProps> = ({ id, firstName, lastName, avatar_url, username }) => {
//     return (
//         <div>
//             <img src={avatar} alt="Avatar" style={{ maxWidth: '100%', height: 'auto' }} />
//             <h2>{firstName} {lastName}</h2>
//             <p>{username}</p>
//         </div>
//     );
// };
