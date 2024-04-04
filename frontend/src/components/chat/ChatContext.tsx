// ChatContext.tsx
import React, {createContext, useContext, useState, ReactNode, useEffect} from 'react';
import {ChatBox} from "@/components/chat/ChatBox";

interface ChatContextType {
    activeChats: any[];
    openChat: (userInfo: any) => void;
    closeChat: (userName: string) => void;
    socket: WebSocket | null;
}

const ChatContext = createContext<ChatContextType | undefined>(undefined);

export const ChatProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
    const [activeChats, setActiveChats] = useState<any[]>([]);
    const [socket, setSocket] = useState<WebSocket | null>(null); // State to hold the WebSocket connection

    const BE_PORT = process.env.NEXT_PUBLIC_BACKEND_PORT;
    // TODO: use the correct URL in prod
    const WS_URL = "ws://iriesphere";

    useEffect(() => {
        // Initialize WebSocket connection
        const newSocket = new WebSocket(`ws://localhost:${BE_PORT}/ws`);
        setSocket(newSocket);

        // Clean up on unmount
        return () => {
            newSocket.close();
        };
    }, []);

    const openChat = (chat: any) => {
        // Check if the chat is already open
        const isChatOpen = activeChats.some(activeChat => activeChat.userID === chat.userID);

        if (!isChatOpen) {
            // If the chat is not open, open it
            setActiveChats(prevChats => [...prevChats, chat]);
        } else {
            // If the chat is already open, bring it to the front
            setActiveChats(prevChats => {
                const otherChats = prevChats.filter(activeChat => activeChat.userID !== chat.userID);
                return [...otherChats, chat];
            });
        }
    };

    const closeChat = (userName: string) => {
        setActiveChats((prev) => prev.filter((chat) => chat.userName !== userName));
    };

    return (
        <ChatContext.Provider value={{ activeChats, openChat, closeChat, socket }}>
            {children}
        </ChatContext.Provider>
    );
};

export const useChat = () => {
    const context = useContext(ChatContext);
    if (context === undefined) {
        throw new Error('useChat must be used within a ChatProvider');
    }
    return context;
};

export const ChatManager = () => {
    const {activeChats, closeChat, socket} = useChat(); // Assuming useChat is a hook to access your chat context
    return (
        <div className="chat-manager">
            {activeChats.map((chatBox: { userID: number; userName: string; avatar: string }) => {
                if (socket !== null) {
                    return (
                        <ChatBox
                            key={chatBox.userName}
                            userID={chatBox.userID}
                            userName={chatBox.userName}
                            avatar={chatBox.avatar}
                            onClose={() => closeChat(chatBox.userName)}
                            socket={socket}
                        />
                    );
                }
            })}
        </div>
    );
};

