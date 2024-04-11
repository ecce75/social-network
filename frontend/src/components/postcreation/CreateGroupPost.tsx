import React, { useState } from 'react';
import {PostProps} from "@/components/postcreation/Post";
import {CommentProps} from "@/components/comments/Comment";

interface CreatePostGroupProps {
    groupId: string;
    onNewPost?: (newPost: PostProps) => void;
    setComments: React.Dispatch<React.SetStateAction<{[postId: number]: CommentProps[]}>>;
}

const CreateGroupPost: React.FC<CreatePostGroupProps> = ({ groupId , onNewPost, setComments}) => {
    const [selectedFile, setSelectedFile] = useState<File | null>(null);
    const [title, setTitle] = useState<string>('');
    const [message, setMessage] = useState<string>('');

    const BE_PORT = process.env.NEXT_PUBLIC_BACKEND_PORT;
    const FE_URL = process.env.NEXT_PUBLIC_URL;

    const handleTitleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setTitle(event.target.value);

    };

    const handleMessageChange = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
        setMessage(event.target.value);
    };

    const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        if (event.target.files && event.target.files.length > 0) {
            setSelectedFile(event.target.files[0]);
        }
    };

    const handleSubmit = async (event: React.FormEvent) => {
        event.preventDefault(); // Prevent default form submission behavior


        const formData = new FormData();
        formData.append('title', title);
        formData.append('content', message);
        formData.append('privacy-setting', "private");
        formData.append('group', groupId)

        if (selectedFile) {
            formData.append('image', selectedFile);
        }

        try {
            const response = await fetch(`${FE_URL}:${BE_PORT}/post`, {
                method: 'POST',
                credentials: 'include', // If you're handling sessions
                body: formData, // Send the form data
            });

            if (!response.ok) {
                throw new Error(`Error: ${response.statusText}`);
            }

            const data = await response.json();
            console.log('Group post created:', data.data);
            const newPost: PostProps = {
                id: data.data.id,
                creator: data.data.user_id,
                userId: data.data.user_id,
                groupId: data.data.group_id,
                title: data.data.title,
                content: data.data.content,
                image_url: data.data.image_url,
                privacySetting: data.data.privacy_setting,
                created_at: new Date(data.data.created_at),
                likes: 0,
                dislikes: 0,
                setComments: setComments,
            }
            onNewPost && onNewPost(newPost);
            const modal = document.getElementById('Modal_Post_Group') as HTMLDialogElement | null;
            if (modal) {
                modal.close();
            }
            // Handle success (e.g., clear form, show success message)
        } catch (error) {
            console.error('Error submitting post:', error);
        }
    };


    return (
        <div>
            <div className="flex justify-between">
                    <div>
                    {/* Top message box */}
                    <input type="text" placeholder="Title" className="input mt-2 w-full max-w-sm" onChange={handleTitleChange}/>
                    </div>
                    

                </div>
                {/* Main message box */}
                <div className="relative w-full min-w-[200px] mt-2">
                    <textarea
                        className="peer h-full min-h-[200px] w-full resize-none border-b border-blue-gray-200 bg-transparent pt-4 pb-1.5 font-sans text-lg font-normal text-gray-900 outline outline-0 transition-all placeholder-shown:border-blue-gray-200 focus:border-gray-900 focus:outline-0 disabled:resize-none disabled:border-0 disabled:bg-blue-gray-50"
                        placeholder=""
                        onChange={handleMessageChange}
                    ></textarea>
                    <label
                        className="after:content[' '] pointer-events-none absolute left-0 -top-1.5 flex h-full w-full select-none text-[11px] font-normal leading-tight text-gray-900 transition-all after:absolute after:-bottom-0 after:block after:w-full after:scale-x-0 after:border-b-2 after:border-gray-900 after:transition-transform after:duration-300 peer-placeholder-shown:text-sm peer-placeholder-shown:leading-[4.25] peer-placeholder-shown:text-blue-gray-500 peer-focus:text-[11px] peer-focus:leading-tight peer-focus:text-gray-900 peer-focus:after:scale-x-100 peer-focus:after:border-gray-900 peer-disabled:text-transparent peer-disabled:peer-placeholder-shown:text-blue-gray-500"
                    >
                        Message
                    </label>
                </div>

                {/* Image upload and preview */}
                <div className="mb-4 flex justify-start items-end">
                    <div>
                        <input
                            type="file"
                            id="image-upload"
                            className="hidden"
                            accept="image/*"
                            onChange={handleFileChange}
                        />
                        <label htmlFor="image-upload" className="btn cursor-pointer">
                            Upload Photo
                        </label>
                    </div>
                    <div className="avatar-preview">
                        {selectedFile && (
                            <img
                                src={URL.createObjectURL(selectedFile)}
                                alt="Preview"
                                className="avatar"
                                style={{ width: 100, height: 100 }}
                            />
                        )}
                    </div>

                    {/* Post button*/}
                    <div className="flex-grow" />
                    <button className="btn" onClick={handleSubmit}>Post</button>
                </div>
        </div>
    );
}

export default CreateGroupPost;
