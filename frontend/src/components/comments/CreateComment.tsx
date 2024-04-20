import { useState } from "react";
import { CommentProps } from "./Comment";

interface CreateCommentProps {
    postId: number;
    setComments: React.Dispatch<React.SetStateAction<{ [postId: number]: CommentProps[] }>>;
}

const CreateComment: React.FC<CreateCommentProps> = ({ postId, setComments }) => {
    const BE_PORT = process.env.NEXT_PUBLIC_BACKEND_PORT;
    const FE_URL = process.env.NEXT_PUBLIC_URL;

    const [selectedFile, setSelectedFile] = useState<File | null>(null);
    const [commentContent, setCommentContent] = useState<string>('');

    const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        if (event.target.files && event.target.files.length > 0) {
            setSelectedFile(event.target.files[0]);
        }
    }

    const handleCommentContentChange = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
        setCommentContent(event.target.value);
    }

    const handleSubmit = async (event: React.FormEvent) => {
        // Create a new comment
        event.preventDefault();

        const formData = new FormData();
        formData.append('content', commentContent);
        if (selectedFile) {
            formData.append('image', selectedFile);
        }
        try {
            const response = await fetch(`${FE_URL}:${BE_PORT}/post/${postId}/comment`, {
                method: 'POST',
                body: formData,
                credentials: 'include' // Send cookies with the request
            });

            if (!response.ok) {
                console.error('Error creating comment:', response.statusText);
                return;
            }

            const newComment = await response.json();
            setCommentContent('');
            setSelectedFile(null);
            setComments(prevComments => {
                // Get the current comments for the post or an empty array if none
                const currentCommentsForPost = prevComments[postId] || [];
                // Append the new comment to the array
                const updatedCommentsForPost = [...currentCommentsForPost, newComment];
                // Return the updated comments state with the new comment included for the specific post
                return { ...prevComments, [postId]: updatedCommentsForPost };
            });
        } catch (error) {
            console.log('Error creating comment:', error);
        }
    }

    const fileInputId = `comment-image-${postId}`;

    return (
        <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
            {/* Text field for commenting*/}
            <div className="flex items-center mt-3.5 px-3 py-2 rounded-lg bg-gray-50 dark:bg-gray-700 flex-grow">
                <input type="file"
                    id={fileInputId}
                    className="hidden"
                    accept="image/*"
                    onChange={handleFileChange}
                />
                <label htmlFor={fileInputId} className="inline-flex justify-center p-2 text-gray-500 rounded-lg cursor-pointer hover:text-gray-900 hover:bg-gray-100 dark:text-gray-400 dark:hover:text-white dark:hover:bg-gray-600">
                    <svg className="w-5 h-5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 20 18">
                        <path fill="currentColor" d="M13 5.5a.5.5 0 1 1-1 0 .5.5 0 0 1 1 0ZM7.565 7.423 4.5 14h11.518l-2.516-3.71L11 13 7.565 7.423Z" />
                        <path stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M18 1H2a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1Z" />
                        <path stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M13 5.5a.5.5 0 1 1-1 0 .5.5 0 0 1 1 0ZM7.565 7.423 4.5 14h11.518l-2.516-3.71L11 13 7.565 7.423Z" />
                    </svg>
                    <span className="sr-only">Upload image</span>
                </label>

                <button type="button" className="p-2 text-gray-500 rounded-lg cursor-pointer hover:text-gray-900 hover:bg-gray-100 dark:text-gray-400 dark:hover:text-white dark:hover:bg-gray-600">
                    <svg className="w-5 h-5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 20 20">
                        <path stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M13.408 7.5h.01m-6.876 0h.01M19 10a9 9 0 1 1-18 0 9 9 0 0 1 18 0ZM4.6 11a5.5 5.5 0 0 0 10.81 0H4.6Z" />
                    </svg>
                    <span className="sr-only">Add emoji</span>
                </button>

                <textarea
                    id="chat"
                    rows={1}
                    className="block mx-4 p-2.5 w-full h-full text-sm text-gray-900 bg-white rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-800 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                    placeholder="Your comment..."
                    value={commentContent}
                    onChange={handleCommentContentChange}
                ></textarea>

                {/* Image preview */}
                <div className="uploaded-image-preview relative">
                    {selectedFile && (
                        <>
                            <img
                                src={URL.createObjectURL(selectedFile)}
                                alt="Preview"
                                className="avatar"
                                style={{ width: 150, height: 150 }}
                            />
                            <button
                                onClick={() => setSelectedFile(null)}
                                className="absolute top-0 right-0 bg-red-500 text-white rounded-full p-1 text-xs"
                                style={{ width: '20px', height: '20px', display: 'flex', alignItems: 'center', justifyContent: 'center' }}
                            >
                                x
                            </button>
                        </>
                    )}
                </div>

                <button type="submit" className="inline-flex justify-center p-2 text-black rounded-full cursor-pointer hover:bg-green-100 dark:text-blue-500 dark:hover:bg-gray-600" onClick={handleSubmit}>
                    <svg className="w-5 h-5 rotate-90 rtl:-rotate-90" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 18 20">
                        <path d="m17.914 18.594-8-18a1 1 0 0 0-1.828 0l-8 18a1 1 0 0 0 1.157 1.376L8 18.281V9a1 1 0 0 1 2 0v9.281l6.758 1.689a1 1 0 0 0 1.156-1.376Z" />
                    </svg>
                    <span className="sr-only">Send message</span>
                </button>
            </div>
            {/* Like Button */}
            <button className="btn bg-primary ml-2">
                <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" /></svg>
            </button>
        </div>
    );
}

export default CreateComment;