# Simple CLI Messaging App

A lightweight command-line messaging application written in Go. This project consists of a server component and a client-side executable that communicates over WebSockets.

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Important Notice](#important-notice)
- [Installation](#installation)
  - [Server Setup](#server-setup)
  - [Client Setup](#client-setup)
- [Usage](#usage)
- [Deployment Options](#deployment-options)
- [Final Note](#final-note)
- [License](#license)

## Overview

This project provides a simple way to send and receive messages directly from your terminal. The server is built in Go and can be run locally or deployed to a web service like [Render.com](https://render.com/). The client, distributed as a compiled executable, connects to the server using WebSockets.

## Features

- **CLI-Based**: Operate entirely from the command line.
- **WebSocket Connection**: Real-time messaging over WebSockets.
- **Flexible Deployment**: Run the server locally with port forwarding or deploy on a hosted service.
- **Easy Setup**: Download the client executable from the GitHub releases section.

## Important Notice

- **No Encryption**: The messages sent through this application are not encrypted. Use with caution if transmitting sensitive information.
- **Ephemeral Messaging**: Messages are not saved or stored. They are delivered in real time and:
  - If there are no users available to receive the message at the time of sending, the message will not be stored.
  - Once all clients that received a message disconnect, the message is gone permanently.

## Installation

### Server Setup

1. **Clone the Repository**:
   
   ```sh
   git clone https://github.com/Squid-Logic/GoLang-Chat
   cd your-project
   ```

2. **Build the Server**:
   
   Ensure you have Go installed. Then, build the server:
   
   ```sh
   go build -o server ./server
   ```

3. **Run the Server Locally**:
   
   You can run the server locally. If you're behind NAT, consider using port forwarding:
   
   ```sh
   ./server
   ```
   
   The server listens on a specified port for WebSocket connections.

### Client Setup

1. **Download the Client**:
   
   - Navigate to the Releases section of the GitHub repository.
   - Download the client executable for your operating system.

2. **Run the Client**:
   
   - Open the executable.
   - When prompted, enter your connection URL. For example, if your server is deployed on Render:
   
   ```sh
   wss://your-project.onrender.com
   ```

## Usage

1. **Start the Server**:
   
   - Run the server either locally or on your chosen web hosting service.

2. **Launch the Client**:
   
   - Run the client executable.
   - Enter the WebSocket URL when prompted.
   - Begin sending and receiving messages.

## Deployment Options

### Local Deployment

- Run the server on your machine with appropriate port forwarding to allow external access.

### Hosted Deployment

- Deploy the server on any web hosting service. Render.com is one example where you can host your Go server.
- Adjust the WebSocket URL in your client accordingly (e.g., `wss://your-project.onrender.com`).

## Final Note

This project was a short, fun experiment to learn more about Go and WebSockets. It is not intended for production use, and I will not be actively maintaining it. Feel free to fork the repository, review the code, and build your own more robust version if you find this concept interesting. Happy coding and learning!
