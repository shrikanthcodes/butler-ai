## Steps on how to run frontend

Step 1: 
# Install nodejs and npm:
- sudo apt-get install nodejs
- sudo apt-get install npm

Step 2:
# Install expo-cli
- npm install -g expo-cli

Step 3: 
- npx expo install react-native-web react-dom @expo/metro-runtime -- --legacy-peer-deps

Step 4: 
- cd frontend
- npm run android
- npm run ios # you need to use macOS to build the iOS project - use the Expo app if you need to do iOS development without a Mac
- npm run web





## How to create frontend

- npx create-expo-app --template

- choose template : plain with typescript

- name of project: frontend
