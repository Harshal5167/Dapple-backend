# Dapple Backend

<div align="center">
	<img src="./data/logo.png" alt="Dapple Logo" width="200"/>
	<h1>Dapple: one spot at a Time</h1>
	<h4>Dapple is an AI-powered mobile application designed to help neurodiverse individuals, particularly those with autism, improve their communication and social skills. This repository contains the Go-based backend services that power the Dapple platform.</h4>

</div>

## Features

- **Personalized Learning**: Dapple offers AI-driven modules that adapt to each user’s unique learning style, strengths, and challenges. This personalization builds confidence and reduces frustration by allowing users to progress at their own pace in a structured, supportive environment.

- **Real-Time Therapy Sessions**: Integrated virtual therapy sessions provide immediate support and guidance. These one-on-one expert-led sessions help users navigate social scenarios, reinforcing skills and reducing anxiety through real-time practice.

- **Assistive Communication**: Tools such as visual schedules, transition alerts, and "Help Cards" empower non-verbal users to communicate their needs, promoting independence and smoother daily transitions both at home and in community settings.

- **Interactive Learning**: Through engaging methods like multiple-choice questions, scenario-based exercises, and AI-powered role-playing, users can practice real-world interactions. These features help develop conversational skills and foster meaningful connections with peers, family, and colleagues.

## Tech Stack

- **Backend**: Golang (Framework - GoFiber)
- **Database**: Firebase, Redis
- **Authentication**: Firebase Auth

## Getting Started

### Prerequisites

- go1.23.5 or higher
- Docker (for local redis instance)
- Firebase account (for database and authentication)
- Google cloud service account (for Calendar Event Generation)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/Harshal5167/Dapple-backend.git
   cd Dapple-backend
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Configure environment variables:
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

4. Run the server:
   ```bash
   go run main.go
   ```

## Project Structure
```
.
├── config/           # Configuration files
├── data/             # Static data files (e.g., images, JSON files)
├── internal/         # Application code
│   ├── middleware/   # Middleware functions
│   ├── utils/        # Utility functions
│   ├── dto/          # Data Transfer Objects (DTOs)
│   │   ├── request/  # Request DTOs
│   │   └── response/ # Response DTOs
│   ├── clients/      # External API clients
│   │   ├── voiceEvaluation/  # Voice model client
│   │   └── videoEvaluation/  # Video model client
│   ├── interfaces/   # Interfaces for dependency injection 
│   ├── models/       # Data models
│   ├── handlers/     # HTTP request handlers
│   ├── service/      # Business logic
│   └── repository/   # Data access logic
├── .env.example      # Example environment variables file
├── main.go           # Main application entry point
├── go.mod            # Go module file
├── go.sum            # Go module dependencies
├── readme.md         # Project documentation
└── .gitignore        # Git ignore file
```

## License
This project is licensed under the [MIT License](./LICENSE). See the LICENSE file for details.

## Contact

Project Maintainer - Harshal Gainer ([Harshal5167](https://github.com/Harshal5167))