"use client"

import GroupPageInfo from "@/components/groups/GroupPageInfo";
import CreateGroupPostButton from "@/components/buttons/CreateGroupPostButton";
import GroupEventFeed from "@/components/groups/GroupEventFeed";
import React, { useEffect, useState } from "react";
import Post, { PostProps } from "@/components/postcreation/Post";
import useAuthCheck from "@/hooks/authCheck";
import { CommentProps } from "@/components/comments/Comment";

interface GroupProps {
    id: string
    creator_id: string
    title: string
    description: string
    image: string
    created_at: string
    updated_at: string
    members: GroupMember[]
}

interface GroupMember {
    group_id: string
    user_id: string
    joined_at: string
}


export default function Group({
    params,
}: {
    params: {
        id: string
    }
}) {
    useAuthCheck();

    const [group, setGroup] = React.useState<GroupProps | null>(null);
    const [posts, setPosts] = React.useState<PostProps[]>([]);
    const [isMember, setMember] = React.useState<boolean>(true);
    const [isCreator, setCreator] = React.useState<boolean>(false);
    const [invitationSent, setInvitationSent] = React.useState<boolean>(false);
    const [comments, setComments] = useState<{ [postId: number]: CommentProps[] }>([]);
    const [confirmInvite, setConfirmInvite] = React.useState<boolean>(false);


    const BE_PORT = process.env.NEXT_PUBLIC_BACKEND_PORT;
    const FE_URL = process.env.NEXT_PUBLIC_URL;
    useEffect(() => {
        try {
            fetch(`${FE_URL}:${BE_PORT}/groups/${params.id}`, {
                method: 'GET',
                credentials: 'include'
            })
                .then(response => response.json())
                .then(data => {
                    if (data.is_user_creator) {
                        setCreator(true);
                    }
                    setGroup(data);
                })
        } catch (error) {
            console.error('Error fetching group:', error);
        }
    }, [])

    useEffect(() => {
        try {
            fetch(`${FE_URL}:${BE_PORT}/groups/${params.id}/posts`, {
                method: 'GET',
                credentials: 'include'
            })
                .then(response => response.json())
                .then(data => {
                    if (data.message === 'User not member of group') {
                        setMember(false);
                    } else {
                        setPosts(data);
                    }
                })
        } catch (error) {
            console.error('Error fetching posts:', error);
        }
    }, [])

    useEffect(() => {
        if (!isMember) {
            fetch(`${FE_URL}:${BE_PORT}/invitations/${params.id}`, {
                method: 'GET',
                credentials: 'include'
            })
                .then(response => response.json())
                .then(data => {
                    if (data.message != 'No group invitation found' && data.status == "pending") {
                        if (data.invite_user_id == data.join_user_id) {
                            setInvitationSent(true);
                        } else if (data.invite_user_id != data.join_user_id) {
                            setConfirmInvite(true);
                        }

                    }

                })
        }

    }, [isMember]);

    const handleNewPost = (newPost: PostProps) => {
        setPosts((prevPosts) => [...prevPosts, newPost]);
    };

    return (

        <div>
            {/* Main Content */}
            <main>
                <div style={{ display: 'flex', justifyContent: 'center' }}> {/* Container for both sections */}


                    {/* Left section for displaying group information */}
                    <div style={{
                        flex: '0 0 18%',
                        backgroundColor: '#e5e7eb',
                        padding: '20px',
                        height: '100vh',
                        overflowY: 'auto'
                    }}>
                        <GroupPageInfo
                            title={group?.title}
                            text={group?.description}
                            pictureUrl={group?.image}
                            isMember={isMember}
                            groupId={params.id}
                            invitationSent={invitationSent}
                            isCreator={isCreator}
                            confirmInvite={confirmInvite}
                        />
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
                        {isMember && <div style={{ marginBottom: '20px' }}>
                            <CreateGroupPostButton
                                groupId={params.id}
                                onNewPost={handleNewPost}
                                setComments={setComments}
                            />
                        </div>}
                        <div style={{ display: 'flex', flexDirection: 'column', marginBottom: '20px' }}>
                            {
                                isMember ? (
                                    posts.length > 0 ?
                                        posts.map(post =>
                                            <Post
                                                key={post.id}
                                                {...post}
                                                comments={comments[post.id]}
                                                setComments={setComments}
                                            />
                                        )
                                        :
                                        <div>
                                            <p>No posts found</p>
                                        </div>
                                ) : (
                                    <div>
                                        <p className="text-xl text-green-700 font-semibold mt-2">You're not a member of this group</p>
                                    </div>
                                )
                            }
                        </div>
                    </section>


                    {/* Divider */}
                    <div style={{ flex: '0 0 5px', backgroundColor: '#B2BEB5', height: '100vh' }}></div>


                    {/* Left section for displaying group information */}
                    {isMember && <div style={{
                        flex: '0 0 18%',
                        backgroundColor: '#e5e7eb',
                        padding: '20px',
                        height: '100vh',
                        overflowY: 'auto',
                    }}>
                        <GroupEventFeed groupId={params.id} />
                    </div>}
                </div>
            </main>

        </div>

    )
}