"use client"

import React, {useEffect, useRef, useState} from 'react';

import "../../../styles/styles.css";
import MessageInput from "@/components/chat/MessageInput";
import MessageList from "@/components/chat/MessageList";

interface ChatBoxProps {
    userID: number;
    userName: string;
    avatar: string;
    onClose: () => void; // Assuming onClose is a function that takes no arguments
    socket: WebSocket;
}

const ChatBox: React.FC<ChatBoxProps> = ({userID, userName, avatar, onClose, socket}) => {
    const [messages, setMessages] = useState<Array<{
        id: number,
        text: string,
        sender: string,
        timestamp: string
    }>>([]); // Example message state
    const [currentPage, setCurrentPage] = useState(1);
    const [allFetched, setAllFetched] = useState(false);
    const messageListRef = useRef<HTMLDivElement>(null); // Ref for the message list container



    // Initialize WebSocket connection
    useEffect(() => {
        fetchMessages(userID, currentPage);

        socket.onmessage = (event) => {
            const message = JSON.parse(event.data);
            // Handle different types of messages here
            if (message.action === "send_message" && message.sender === userID) {
                setMessages((prevMessages) => [...prevMessages, {
                    id: message.id,
                    text: message.content,
                    sender: "them",
                    timestamp: message.timestamp
                }]);
            }
            if (message.action === "chat_history") {
                const fetchedMessages = message.content; // Assuming this is the array of messages

                if (fetchedMessages.length > 0) {
                    const previousScrollHeight = messageListRef.current?.scrollHeight || 0
                    // There are messages, so append them to the current messages
                    setMessages((prevMessages) => [
                        ...prevMessages,
                        ...fetchedMessages.map((msg: any) => ({
                            id: msg.id,
                            text: msg.text,
                            sender: msg.sender === userID ? "them" : "me",
                            timestamp: msg.timestamp,
                        })),
                    ]);
                    // After state update, adjust scroll to maintain position
                    setTimeout(() => {
                        const currentScrollHeight = messageListRef.current?.scrollHeight || 0;
                        const scrollOffset = currentScrollHeight - previousScrollHeight;
                        if (messageListRef.current) {
                            messageListRef.current.scrollTop += scrollOffset;
                        }
                    }, 0);
                } else {
                    // The list is empty, so there are no more messages to fetch
                    setAllFetched(true);
                }
            }
            // Add other message handling logic here
        };

        // }
    }, []);


    const fetchMessages = (userID: number, page: number) => {
        // Construct the request for fetching messages
        const requestData = {
            action: "fetch_chat_history",
            user: userID,
            page: page,
        };
        if (socket.readyState === WebSocket.OPEN) {
            // Send the request via WebSocket
            socket.send(JSON.stringify(requestData));
            // Prepare for the next page
            setCurrentPage(page + 1);
        } else {
            console.log("Socket not ready, cannot fetch messages");
        }
    };

    const loadMoreMessages = () => {
        if (!allFetched && messageListRef.current) {
            const {scrollTop} = messageListRef.current;
            if (scrollTop < 20) { // You might adjust this threshold
                if (socket && socket.readyState === WebSocket.OPEN) {
                    fetchMessages(userID, currentPage);
                }
            }
        }
    };

    useEffect(() => {
        const messageListElement = messageListRef.current;
        if (messageListElement) {
            messageListElement.addEventListener('scroll', loadMoreMessages);

            // Cleanup
            return () => {
                messageListElement.removeEventListener('scroll', loadMoreMessages);
            };
        }
    }, [currentPage, allFetched]);

    const handleSend = (text: string) => {
        const tempID = Date.now();
        const newMessage = {id: tempID, text, sender: "me", timestamp: new Date().toISOString()}; // Example new message
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
        <div
            className="chatbox-wrapper bottom-0 right-0 mb-4 mr-4 max-w-xs w-full bg-white shadow-lg rounded-lg flex flex-col">
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

                    <button onClick={onClose} className="btn btn-xs btn-circle relative bottom-1 bg-primary">
                        <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6 text-white " fill="none"
                             viewBox="0 0 24 24"
                             stroke="currentColor">
                            <path className="bg-primary" strokeLinecap="round" strokeLinejoin="round" strokeWidth="2"
                                  d="M6 18L18 6M6 6l12 12"/>
                        </svg>
                    </button>
                </div>
            </div>
            <MessageList ref={messageListRef} messages={messages} userName={userName}/>
            <MessageInput onSend={handleSend}/>
        </div>
    );
};


export {ChatBox};
