"use client"

import React, {useEffect, useState} from 'react';

import {useChat} from "@/components/chat/ChatContext";
import "../../../styles/styles.css";
import MessageInput from "@/components/chat/MessageInput";
import MessageList from "@/components/chat/MessageList";

interface ChatBoxProps {
    userID: number;
    userName: string;
    avatar: string;
    onClose: () => void; // Assuming onClose is a function that takes no arguments
}

const ChatBox: React.FC<ChatBoxProps> = ({userID, userName, avatar, onClose}) => {
    const [messages, setMessages] = useState<Array<{id: number, text: string, sender: string, timestamp: string}>>([]); // Example message state
    const [socket, setSocket] = useState<WebSocket | null>(null);

    // Initialize WebSocket connection
    useEffect(() => {
        const newSocket = new WebSocket("ws://localhost:8080/ws");
        setSocket(newSocket);

        newSocket.onmessage = (event) => {
            const message = JSON.parse(event.data);
            // Handle different types of messages here
            if (message.action === "send_message") {
                console.log("Received message:", message);
                setMessages((prevMessages) => [...prevMessages, { id: prevMessages.length + 1, text: message.content, sender: "them", timestamp: message.timestamp}]);
            }
            // Add other message handling logic here
        };

        return () => {
            newSocket.close();
        };
    }, []);


    const handleSend = (text: string) => {
        const newMessage = { id: messages.length + 1, text, sender: "me", timestamp: new Date().toISOString()  }; // Example new message
        setMessages([...messages, newMessage]);
        if (socket && socket.readyState === WebSocket.OPEN) {
            const jsonData = {
                action: "send_message",
                recipientID: userID,
                content: text,
            };
            socket.send(JSON.stringify(jsonData));
        }
    };


    return (
        <div className="chatbox-wrapper bottom-0 right-0 mb-4 mr-4 max-w-xs w-full bg-white shadow-lg rounded-lg flex flex-col">
            <div className="header p-3 border-b bg-primary rounded-t-lg">
                <div className="flex justify-between">
                    <div className="flex justify-start">
                    <div className="avatar">
                        <div className="w-12  rounded-full">
                            <img src={avatar} alt="User avatar"/>
                        </div>
                    </div>

                    <h2 className="font-bold text-xl ml-4 mt-2">{userName}</h2>
                </div>

                    {/*<button onClick={onClose} className="btn relative bottom-6 left-24">X</button>*/}
                    <button onClick={onClose} className="btn btn-xs btn-circle relative bottom-1 bg-primary">
                        <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6 text-white " fill="none" viewBox="0 0 24 24"
                             stroke="currentColor">
                            <path className="bg-primary" strokeLinecap="round" strokeLinejoin="round" strokeWidth="2"
                                  d="M6 18L18 6M6 6l12 12"/>
                        </svg>
                    </button>
                </div>
            </div>
            <MessageList messages={messages} userName={userName}/>
            <MessageInput onSend={handleSend}/>
        </div>
    );
};


export {ChatBox};
