# Shinano

Shinano is a reverse engineering project focused on reproducing parts of the server-side behaviour expected by the original game client. The work involves examining the communication between the client and server, understanding the protobuf messages, and implementing the corresponding game logic.

I initially started the project as a way to learn Go through something more substantial than small exercises. 

## Project approach

The goal is not to recreate the entire original server. The implementation focuses on the systems that are relevant to running the game locally. Some of them have their own logic and persistent state; others use minimal protobuf responses where full functionality is outside the project's scope.

A few things that define the project:

- Designed around a single local commander rather than multiple players or public hosting.
- SQLite is used for persistent game state, with versioned migrations.
- Static game data is kept separate from runtime state and loaded from JSON.
- Protobuf messages are implemented incrementally as the client behaviour becomes understood.
- Minimal responses are used for selected systems that do not need full implementation.

## Architecture

```text
Shinano
├── cfg                 # Static game data and configuration
├── constructor         # Request handlers and protobuf response builders
│   └── zero_proto      # Predefined responses for systems without game logic
├── db                  # SQLite initialization and migrations
├── packets             # Packet parsing, building and dispatch
└── repository          # Game state and database operations
```

