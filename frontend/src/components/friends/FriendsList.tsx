"use client"
import { useState, useEffect } from "react";
import { useAuth } from "../../util/utils";
import UserTab from "./UserTab";

interface FriendProps {
    id: number;
    firstName: string;
    lastName: string;
    avatar_url: string;
    username: string;
}

export default function FriendsList() {
    const [friends, setFriends] = useState<FriendProps[]>([]);

    useEffect(() => {
        fetch('http://localhost:8080/friends', {
            method: 'GET',
            credentials: 'include' // Send cookies with the request
        })
            .then(response => response.json())
            .then(data => {
                if (data.message === "No friends found") {
                    // TODO: Handle no friends found
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
                        userName={friend.username}
                        avatarUrl={friend.avatar_url}
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
