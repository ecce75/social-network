// components/MessageList.tsx
import React, {useEffect, useRef} from 'react';

interface MessageListProps {
    messages: Array<{ id: number; text: string; sender: string, timestamp: string}>;
    userName: string;
}

const MessageList = React.forwardRef<HTMLDivElement, MessageListProps>((props, ref) => {
    const { messages, userName } = props;
    const messagesEndRef = useRef<HTMLDivElement | null>(null);

    const scrollToBottom = () => {
        console.log("Scrolling to bottom");
        messagesEndRef.current?.scrollIntoView({ behavior: "smooth" });
    };

    useEffect(scrollToBottom, [messages]);

    // Sort messages in descending order based on id
    const sortedMessages = [...messages].sort((a, b) => a.id - b.id);

    return (
        <div ref={ref} className="messages flex-1 overflow-y-auto p-4 bg-gray-100">
            {sortedMessages.map((message) => (

                <div key={message.id} className={`chat ${message.sender === 'me' ? 'chat-end' : 'chat-start'}`}>
                    <div className="chat-header">
                        <span className="text-gray-500">{message.sender === 'me' ? 'You' : userName}</span>
                        <time
                            className="ml-2 text-black text-xs opacity-50">{new Date(message.timestamp).toLocaleTimeString('en-US', {
                            hour12: false,
                            hour: '2-digit',
                            minute: '2-digit'
                        })}</time>
                    </div>
                    <div
                        className={`chat-bubble ${message.sender === 'me' ? 'bg-primary' : 'bg-secondary'} text-white`}>
                        {message.text}
                    </div>
                </div>
            ))}
            <div ref={messagesEndRef}/>
        </div>
    );
});

MessageList.displayName = 'MessageList';

export default MessageList;
