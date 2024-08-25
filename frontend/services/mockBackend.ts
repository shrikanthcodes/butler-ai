import { Dialogue } from '../types';

export const handleConversation = async (userMessage: string, conversation: Dialogue[]): Promise<Dialogue[]> => {
    // Add the user message to the conversation
    conversation.push({ role: 'user', content: userMessage });

    // Send the conversation to the backend
    const response = await fetch('http://localhost:8080/chat/conversation', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            conversation_id: 'test',
            messages: conversation,
        }),
    });

    if (!response.ok) {
        throw new Error('Failed to handle conversation');
    }

    const data = await response.json();
    return data.messages;
};
