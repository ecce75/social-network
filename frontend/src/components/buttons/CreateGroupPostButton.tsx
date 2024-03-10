
import CreateGroupPost from "../postcreation/CreateGroupPost";
import {PostProps} from "@/components/postcreation/Post";
import React from "react";
import {CommentProps} from "@/components/comments/Comment";

interface CreatePostButtonGroupProps {
    groupId: string;
    onNewPost?: (newPost: PostProps) => void;
    setComments: React.Dispatch<React.SetStateAction<{[postId: number]: CommentProps[]}>>;

}

const CreateGroupPostButton:React.FC<CreatePostButtonGroupProps> = ({groupId, onNewPost, setComments}) => {
    const openModal = () => {
        const modal = document.getElementById('Modal_Post_Group') as HTMLDialogElement | null;
        if (modal) {
            modal.showModal();
        }
    };

    return (
        <div>
            {/* Open the modal using document.getElementById('ID').showModal() method */}
            <button className="btn btn-primary text-white" onClick={openModal}>Create Post</button>
            <dialog id="Modal_Post_Group" className="modal">
                <div className="modal-box" style={{maxWidth:'none', width: '50%', height: '50%'}}>
                    <h3 className="font-bold text-black text-lg">Group Post Creation</h3>
                    <CreateGroupPost
                    groupId={groupId}
                    onNewPost={onNewPost}
                    setComments={setComments}
                    />
                    
                </div>
                <form method="dialog" className="modal-backdrop">
                    <button>close</button>
                </form>
            </dialog>
        </div>
    )
}

export default CreateGroupPostButton;
