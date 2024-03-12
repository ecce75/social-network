// components/MessageInput.tsx
import React, { useState, useRef, useEffect } from 'react';
import { FaJoint } from 'react-icons/fa';
import data from '@emoji-mart/data';
import Picker from '@emoji-mart/react';

interface MessageInputProps {
  onSend: (message: string) => void;
}

const MessageInput: React.FC<MessageInputProps> = ({ onSend }) => {
  const [text, setText] = useState('');
  const [showEmoji, setShowEmoji] = useState(false);
  const emojiButtonRef = useRef<HTMLButtonElement>(null);
  const emojiPickerRef = useRef<HTMLDivElement>(null);

  // add emoji
  const addEmoji = (e: { unified: string }) => {
    const sym = e.unified.split('_');
    const codeArray = sym.map((el) => parseInt(el, 16));
    const emoji = String.fromCodePoint(...codeArray);
    setText((prevText) => prevText + emoji);
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();

    // Check if the message is not empty before sending
    if (text.trim() !== '') {
      onSend(text);
      setText('');
    }
  };

  const handleDocumentClick = (e: MouseEvent) => {
    // Close the emoji picker if the click is outside the emoji button and picker
    if (
      emojiButtonRef.current &&
      !emojiButtonRef.current.contains(e.target as Node) &&
      emojiPickerRef.current &&
      !emojiPickerRef.current.contains(e.target as Node)
    ) {
      setShowEmoji(false);
    }
  };

  useEffect(() => {
    document.addEventListener('click', handleDocumentClick);

    return () => {
      document.removeEventListener('click', handleDocumentClick);
    };
  }, []);

  return (
    <form onSubmit={handleSubmit} className="input-container p-2 border-t">
      <input
        type="text"
        value={text}
        onChange={(e) => setText(e.target.value)}
        className="input bg-gray-100 p-2 rounded-lg w-full border-gray-300"
        placeholder="Type a message..."
      />
      <div style={{ display: 'flex' }}>
        <button
          type="submit"
          className="send-button"
          style={{
            backgroundColor: 'darkgreen',
            borderRadius: 10,
            padding: 0,
            marginBlockStart: 4,
            width: 200,
          }}
        >
          <strong>Send</strong>
        </button>
        <button
          type="button"
          onClick={() => setShowEmoji(!showEmoji)}
          className="emoji"
          style={{
            color: 'darkgreen',
            backgroundColor: 'lightgreen',
            borderRadius: 25,
            width: 48,
            height: 30,
            marginInlineStart: 30,
            marginBlockStart: 4,
            paddingInlineStart: 10,
            fontSize: 25,
          }}
          ref={emojiButtonRef}
        >
          <FaJoint />
        </button>

        {showEmoji && (
          <div
            className="absolute bottom-[4%] right-80"
            ref={emojiPickerRef}
          >
            <Picker
              data={data}
              emojiSize={20}
              emojiButtonSize={30}
              onEmojiSelect={addEmoji}
              maxFrequentRows={0}
            />
          </div>
        )}
      </div>
    </form>
  );
};

export default MessageInput;

