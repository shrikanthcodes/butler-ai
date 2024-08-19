import { Dialogue } from '../types';

export const handleConversation = async (userMessage: string, conversation: Dialogue[]): Promise<Dialogue[]> => {
    // Add the user message to the conversation
    conversation.push({ role: 'user', content: userMessage });

    // Simulate AI response (you can replace this with actual logic from your Go backend later)
    const aiResponse = simulateAIResponse(userMessage);
    conversation.push({ role: 'assistant', content: aiResponse });

    return conversation;
};

// Simulate an AI response (replace with real backend logic later)
const simulateAIResponse = (userMessage: string): string => {
    // Example simple logic; you can expand this
    return `AI Response to "${userMessage}"`;
};
