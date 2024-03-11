// components/MessageInput.tsx
import React, { useState } from 'react';
import InputEmoji from "react-input-emoji";

interface MessageInputProps {
    onSend: (message: string) => void;
}

const MessageInput: React.FC<MessageInputProps> = ({ onSend }) => {
    const [text, setText] = useState('');

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
    
        // Check if the message is not empty before sending
        if (text.trim() !== '') {
          onSend(text);
          setText('');
        }
      };

    return (
        <form onSubmit={handleSubmit} className="input-container p-4 border-t">
            <input
                type="text"
                value={text}
                onChange={(e) => setText(e.target.value)}
                className="input bg-gray-100 p-2 rounded-lg w-full border-gray-300"
                placeholder="Type a message..."
            />
            <button type="submit" className="send-button hidden">Send</button>
        </form>
    );
};

export default MessageInput;
