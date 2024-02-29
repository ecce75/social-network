// ChatContext.tsx
import React, { createContext, useContext, useState, ReactNode } from 'react';
import {ChatBox} from "@/components/chat/ChatBox";

interface ChatContextType {
    activeChats: any[];
    openChat: (userInfo: any) => void;
    closeChat: (userName: string) => void;
}

const ChatContext = createContext<ChatContextType | undefined>(undefined);

export const ChatProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
    const [activeChats, setActiveChats] = useState<any[]>([]);

    const openChat = (userInfo: any) => {
        setActiveChats((prev) => [...prev, userInfo]);
    };

    const closeChat = (userName: string) => {
        setActiveChats((prev) => prev.filter((chat) => chat.userName !== userName));
    };

    return (
        <ChatContext.Provider value={{ activeChats, openChat, closeChat }}>
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
    const {activeChats, closeChat} = useChat(); // Assuming useChat is a hook to access your chat context
    console.log(activeChats, closeChat)
    return (
        <div className="chat-manager">
            {activeChats.map((chatBox: { userID: number; userName: string; avatar: string }) => {

                return (
                    <ChatBox
                        key={chatBox.userName}
                        userID={chatBox.userID}
                        userName={chatBox.userName}
                        avatar={chatBox.avatar}
                        onClose={() => closeChat(chatBox.userName)}
                    />
                );
            })}
        </div>
    );
};

