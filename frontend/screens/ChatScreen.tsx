import React, { useState } from 'react';
import { View, TextInput, Button, FlatList, StyleSheet } from 'react-native';
import MessageBubble from '../components/MessageBubble';
import { handleConversation } from '../services/mockBackend';
import { Dialogue } from '../types/index';

const ChatScreen: React.FC = () => {
    const [conversation, setConversation] = useState<Dialogue[]>([]);
    const [inputMessage, setInputMessage] = useState('');

    const handleSendMessage = async () => {
        if (inputMessage.trim()) {
            const updatedConversation = await handleConversation(inputMessage, conversation);
            setConversation(updatedConversation);
            setInputMessage('');
        }
    };

    return (
        <View style={styles.container}>
            <FlatList
                data={conversation}
                renderItem={({ item }) => (
                    <MessageBubble message={item.content} role={item.role} />
                )}
                keyExtractor={(item, index) => index.toString()}
            />
            <View style={styles.inputContainer}>
                <TextInput
                    style={styles.input}
                    value={inputMessage}
                    onChangeText={setInputMessage}
                    placeholder="Type a message..."
                />
                <Button title="Send" onPress={handleSendMessage} />
            </View>
        </View>
    );
};

const styles = StyleSheet.create({
    container: {
        flex: 1,
        padding: 10,
    },
    inputContainer: {
        flexDirection: 'row',
        alignItems: 'center',
    },
    input: {
        flex: 1,
        borderColor: '#ccc',
        borderWidth: 1,
        borderRadius: 20,
        padding: 10,
        marginRight: 10,
    },
});

export default ChatScreen;
