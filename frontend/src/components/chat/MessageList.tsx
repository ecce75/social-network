// components/MessageList.tsx
import React from 'react';

interface MessageListProps {
    messages: Array<{ id: number; text: string; sender: string, timestamp: string}>;
    userName: string;
}

const MessageList: React.FC<MessageListProps> = ({ messages, userName }) => {
    return (
        <div className="messages flex-1 overflow-y-auto p-4">
            {messages.map((message) => (

                <div key={message.id} className={`chat ${message.sender === 'me' ? 'chat-end' : 'chat-start'}`}>
                    <div className="chat-header">
                        <span className="text-gray-500">{message.sender === 'me' ? 'You' : userName}</span>
                        <time className="ml-2 text-black text-xs opacity-50">{new Date(message.timestamp).toLocaleTimeString('en-US', { hour12: false, hour: '2-digit', minute: '2-digit' })}</time>
                    </div>
                    <div className={`chat-bubble ${message.sender === 'me' ? 'bg-primary' : 'bg-secondary'} text-white`}>
                        {message.text}
                    </div>
                </div>
            ))}
        </div>
    );
};

export default MessageList;
