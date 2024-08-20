import React from 'react';
import { View, Text, StyleSheet } from 'react-native';

interface MessageBubbleProps {
    message: string;
    role: 'user' | 'assistant';
}

const MessageBubble: React.FC<MessageBubbleProps> = ({ message, role }) => {
    return (
        <View
            style={[
                styles.container,
                role === 'user' ? styles.userMessage : styles.assistantMessage,
            ]}
        >
            <Text style={styles.messageText}>{message}</Text>
        </View>
    );
};

const styles = StyleSheet.create({
    container: {
        marginVertical: 8,
        padding: 10,
        borderRadius: 10,
        maxWidth: '80%',
    },
    userMessage: {
        alignSelf: 'flex-end',
        backgroundColor: '#d1f5d3',
    },
    assistantMessage: {
        alignSelf: 'flex-start',
        backgroundColor: '#f0f0f0',
    },
    messageText: {
        fontSize: 16,
    },
});

export default MessageBubble;
