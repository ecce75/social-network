
import React, { useState } from 'react';
import { useRouter } from "next/navigation";

function CreateGroup() {
    const router = useRouter();
    const BE_PORT = process.env.NEXT_PUBLIC_BACKEND_PORT;
    const FE_URL = process.env.NEXT_PUBLIC_URL;

    const [selectedFile, setSelectedFile] = useState<File | null>(null);
    const [groupName, setGroupName] = useState<string>('');
    const [groupDescription, setGroupDescription] = useState<string>('');


    const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        if (event.target.files && event.target.files.length > 0) {
            setSelectedFile(event.target.files[0]);
        }
    };
    const handleGroupNameChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setGroupName(event.target.value);
    }
    const handleGroupDescriptionChange = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
        setGroupDescription(event.target.value);
    }

    const handleSubmit = async (event: React.FormEvent) => {
        event.preventDefault();

        const formData = new FormData();
        formData.append('title', groupName);
        formData.append('description', groupDescription);
        if (selectedFile) {
            formData.append('image', selectedFile);
        }
        try {
            const response = await fetch(`${FE_URL}:${BE_PORT}/groups`, {
                method: 'POST',
                body: formData,
                credentials: 'include' // Send cookies with the request
            });

            if (!response.ok) {
                console.error('Error creating group:', response.statusText);
                return;
            }

            const data = await response.json();
            console.log('Group created:', data.id);

            await router.push(`/dashboard/groups/${data.id}`);

        } catch (error) {
            console.log('Error creating group:', error);
        }
    };

    return (
        <div>
            <div className="flex justify-between">
                <div>
                    {/* Title message box */}
                    <input type="text" placeholder="Name your group" className="input mt-2 w-full max-w-sm" onChange={handleGroupNameChange} />
                </div>

            </div>
            {/* Main message box */}
            <div className="relative w-full min-w-[200px] mt-2">
                <textarea
                    className="peer h-full min-h-[100px] w-full resize-none border-b border-blue-gray-200 bg-transparent pt-4 pb-1.5 font-sans text-lg font-normal text-gray-900 outline outline-0 transition-all placeholder-shown:border-blue-gray-200 focus:border-gray-900 focus:outline-0 disabled:resize-none disabled:border-0 disabled:bg-blue-gray-50"
                    placeholder=""
                    onChange={handleGroupDescriptionChange}
                ></textarea>
                <label
                    className="after:content[' '] pointer-events-none absolute left-0 -top-1.5 flex h-full w-full select-none text-[11px] font-normal leading-tight text-gray-900 transition-all after:absolute after:-bottom-0 after:block after:w-full after:scale-x-0 after:border-b-2 after:border-gray-900 after:transition-transform after:duration-300 peer-placeholder-shown:text-sm peer-placeholder-shown:leading-[4.25] peer-placeholder-shown:text-blue-gray-500 peer-focus:text-[11px] peer-focus:leading-tight peer-focus:text-gray-900 peer-focus:after:scale-x-100 peer-focus:after:border-gray-900 peer-disabled:text-transparent peer-disabled:peer-placeholder-shown:text-blue-gray-500"
                >
                    Description
                </label>
            </div>

            {/* Image upload and preview */}
            <div className="mb-4 flex justify-start items-end">

                <div>
                    <div className="avatar-preview">
                        {selectedFile && (
                            <img
                                src={URL.createObjectURL(selectedFile)}
                                alt="Preview"
                                className="avatar"
                                style={{ width: 150, height: 150 }}
                            />
                        )}
                    </div>
                    <input
                        type="file"
                        id="image-upload"
                        className="hidden"
                        accept="image/*"
                        onChange={handleFileChange}
                    />
                    <label htmlFor="image-upload" className="btn cursor-pointer mt-12">
                        Group Picture
                    </label>
                </div>

                {/* Post button*/}
                <div className="flex-grow" />
                <button className="btn" onClick={handleSubmit}>Create</button>
            </div>
        </div>
    );
}

export default CreateGroup;
