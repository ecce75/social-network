import React, { useEffect, useState } from 'react';
import ProfilePageInfo from './ProfilePageInfo';
import { useRouter } from "next/router";
import UserTab from "@/components/friends/UserTab";
import Post, { PostProps } from "../postcreation/Post";

export interface ProfileProps {
    id: number;
    username: string;
    first_name: string;
    last_name: string;
    dob: string;
    avatar_url: string;
    about: string;
    profile_setting: string;
    created_at: string;
}

export interface FriendProps {
    id: number;
    firstName: string;
    lastName: string;
    avatar_url: string;
    username: string;
}

export interface ProfileFeedProps {
    profile: ProfileProps | null;
    friends: FriendProps[];
    userID: any;
}

export const ProfileFeed: React.FC<ProfileFeedProps> = ({ profile, friends, userID }) => {

    const BE_PORT = process.env.NEXT_PUBLIC_BACKEND_PORT;
    const FE_URL = process.env.NEXT_PUBLIC_FRONTEND_URL;
    const [posts, setPosts] = useState<PostProps[]>([]);

    useEffect(() => {
        // Fetch posts
        if (userID !== "me") {
            userID = parseInt(userID);
        }
        fetch(`${FE_URL}:${BE_PORT}/profile/posts/${userID}`, {
            method: 'GET',
            credentials: 'include' // Send cookies with the request
        })
            .then(response => response.json())
            .then(data => {
                if (data === null) {
                    return;
                }
                setPosts(data);
                console.log(data);
            })
            .catch(error => console.error('Error fetching posts:', error));
    }, []);

    return (
        /* Group page with */
        <div style={{ display: 'flex', justifyContent: 'center' }}> {/* Container for both sections */}


            {/* Left section for displaying group information */}
            <div style={{
                flex: '0 0 18%',
                backgroundColor: '#e5e7eb',
                padding: '20px',
                height: '100vh',
                overflowY: 'auto'
            }}>
                <ProfilePageInfo title={profile?.username} pictureUrl={profile?.avatar_url} text={profile?.about} />

                <div style={{ border: '2px solid #ccc', backgroundColor: '#4F7942', borderRadius: '8px', padding: '10px' }}>
                    <h3 style={{ color: 'white', fontWeight: 'bold', fontSize: '20px' }}>Friends</h3>
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
                </div>
            </div>
            {/* Divider */}
            <div style={{ flex: '0 0 5px', backgroundColor: '#B2BEB5', height: '100vh' }}></div>


            {/* Right section for post feed */}
            <section style={{
                flex: '0 0 45%',
                backgroundColor: '#e5e7eb',
                padding: '20px',
                height: '100vh',
                overflowY: 'auto'
            }}>
                <div style={{ marginBottom: '20px', color: 'black' }}>
                    Users past activity
                </div>
                <div style={{ display: 'flex', flexDirection: 'column', marginBottom: '20px' }}>
                    {
                        posts.length > 0 ?
                            posts.map(post =>
                                <Post
                                    key={post.id}
                                    id={post.id}
                                    userId={post.userId}
                                    title={post.title}
                                    content={post.content}
                                    image_url={post.image_url}
                                    privacySetting={post.privacySetting}
                                    created_at={post.created_at}
                                    likes={post.likes}
                                    dislikes={post.dislikes}
                                />
                            )
                            :
                            <div>
                                <p>No posts found</p>
                            </div>
                    }
                </div>
            </section>


        </div>
    );
};

