import { StatusBar } from 'expo-status-bar';
import React, { useState } from 'react';
import { StyleSheet, Text, View, TextInput, Button, Alert } from 'react-native';

export default function App() {
  const [inputValue, setInputValue] = useState('');

  const handleButtonPress = () => {
    Alert.alert('Button Pressed', `You entered: ${inputValue}`);
  };

  return (
    <View style={styles.container}>
      <Text style={styles.header}>Welcome to ButlerAI!</Text>

      <TextInput
        style={styles.input}
        placeholder="Enter something..."
        value={inputValue}
        onChangeText={setInputValue}
      />

      <Button title="Press Me" onPress={handleButtonPress} />

      <StatusBar style="auto" />
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#f0f0f0',
    alignItems: 'center',
    justifyContent: 'center',
    padding: 20,
  },
  header: {
    fontSize: 24,
    fontWeight: 'bold',
    marginBottom: 20,
  },
  input: {
    height: 40,
    borderColor: '#ccc',
    borderWidth: 1,
    paddingHorizontal: 10,
    marginBottom: 20,
    width: '100%',
  },
});
