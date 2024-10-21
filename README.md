
# ButlerAI Project

## Overview
ButlerAI is an AI-powered health and lifestyle assistant that helps users manage tasks like recipe generation, inventory management, health tracking, and more. It integrates modern conversational AI features with real-time data to provide personalized recommendations and an efficient user experience.

This project follows a modular and scalable architecture, leveraging Redis for caching, Squirrel SQL for flexible database queries, and a service-based structure with clear separation of concerns.

## Features
- Conversational AI with peer-to-peer chat support (using GeminiService)
- Recipe and health suggestions based on user preferences and inventory
- Caching layer implemented with Redis for faster performance
- Configurable and scalable architecture using dependency injection
- Custom service layers for database, LLM, chat, and template management
- Custom enum logic in Go to mimic missing native enum functionality
- SQL query handling using Squirrel SQL for more control over queries

## Project Structure
The project follows a clear separation of concerns between packages, services, and the database, allowing for better maintainability and easier future integrations. Key components include:

1. **Services**:
    - `ChatService`: Handles peer-to-peer chat functionality.
    - `GeminiService`: Manages interactions with conversational AI APIs.
    - `TemplateService`: Deals with prompt templates and dynamic conversations.
    - `DBService`: Manages database interactions using Squirrel SQL.
    - `CacheService`: Implements Redis caching for session management and temporary data storage.
    - `UserService`: Handles user-related operations and preferences.

2. **Database**:
    - Switched from GORM ORM to Squirrel SQL for direct SQL query handling, providing better control and flexibility over database operations.

3. **Redis Integration**:
    - Redis is used for caching to enhance performance, especially for session management and temporary data storage.

4. **Dependency Injection**:
    - Dependency injection is implemented across services for better modularity and testability.

5. **Enums**:
    - Custom enum-like behavior is implemented to emulate enums in Go (which lacks native enum support).

6. **Configuration**:
    - The config layer has been enhanced to handle environment-specific settings and additional services.

## Installation and Setup

### Prerequisites
- Go 1.22
- Redis
- PostgreSQL or MySQL for the database (configured via Squirrel SQL)
- Docker (optional, but recommended for easier setup)

### Steps to Set Up Locally

1. **Clone the repository**:
   ```bash
   git clone https://github.com/shrikanthcodes/butler-ai.git
   cd butler-ai
   ```

2. **Set up environment variables**:
   Configure your environment variables for the project. Example `.env` file:
   ```
   REDIS_HOST=localhost
   DB_HOST=localhost
   DB_USER=youruser
   DB_PASSWORD=yourpassword
   ```

3. **Run the project**:
   ```bash
   go run main.go
   ```

4. **(Optional) Using Docker**:
   To set up using Docker, ensure you have Docker installed, and run:
   ```bash
   docker-compose up
   ```

## Usage
Once the project is running, you can interact with ButlerAI through the API endpoints:

- **Chat Service API**: Facilitates peer-to-peer chat and integrates with GeminiService for AI-driven conversations.
- **Template Service API**: Provides dynamic template management for AI prompts.
- **Health and Recipe Suggestions**: Generate personalized recommendations based on user inputs and preferences.

### Example API Request
An example POST request to start a conversation:
```bash
curl -X POST http://localhost:8080/api/chat/start \
  -H 'Content-Type: application/json' \
  -d '{
        "user_id": "123",
        "message": "What should I cook tonight?"
      }'
```

## Contribution Guidelines
We welcome contributions! To get started:

1. Fork the repository.
2. Create a new branch for your feature or bugfix:
   ```bash
   git checkout -b feature/new-feature
   ```
3. Make sure to **rebase** your branch before submitting a pull request to maintain a clean commit history:
   ```bash
   git rebase -i main
   ```
4. Push to your branch and submit a pull request.

## Commit History Highlights
- **CORS support and middleware skeletons**: Implemented CORS to allow requests from ports 3000 and 3001, and added skeletons for basic auth and admin middleware.
- **ChatService and GeminiService**: Added P2P chat architecture, integrated GeminiService for AI-powered conversations, and set up a base architecture for prompt templates.
- **Refactored project structure**: Cleaned up the code by separating concerns across packages, services, and the database. Established a skeleton architecture for future package integrations.
- **Switched from GORM to Squirrel SQL**: Dropped GORM ORM in favor of Squirrel SQL for more flexible and efficient query handling.
- **Dependency injection and Redis**: Implemented DI for better modularity and added Redis caching to optimize performance.

## License
This project is licensed under the MIT License.

---

Future:
APIs: https://github.com/public-apis/public-apis?tab=readme-ov-file#food--drink
