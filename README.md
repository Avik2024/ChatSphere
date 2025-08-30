# Distributed Chat Application (TCP/IP)

A simple and high-performance distributed chat application built using **Go**. The application demonstrates real-time communication between clients and a central server using **TCP/IP**. It showcases core concepts of **distributed systems**, **network programming**, **concurrency**, and **client-server architecture**.

---

## 🚀 **Features**

- **Real-time Messaging:** Clients can send and receive messages instantly through the server.
- **Multiple Clients:** The server supports multiple simultaneous client connections.
- **Concurrency:** The server uses goroutines to handle multiple client connections concurrently.
- **Client-Server Communication:** Uses **TCP sockets** for reliable communication between clients and the server.
- **Scalable Architecture:** Designed for easy expansion, such as adding chat rooms or user authentication.
- **Containerized with Docker:** The project is containerized using Docker for easy deployment.

---

## 🛠️ **Technologies Used**

- **Go (Golang):** For building the server and client applications.
- **TCP/IP:** For network communication between the client and server.
- **Docker:** For containerizing the client and server to simplify deployment.
- **Bash:** For automating build, run, and test processes with shell scripts.

---

## 🔧 **Setup Instructions**

### Prerequisites

- **Go (1.19 or higher)** - To build and run the application.
- **Docker** - To run the server and client inside containers (optional but recommended).

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/distributed-chat.git
cd distributed-chat
```

### 2. Install Dependencies

Make sure you have **Go** installed. Run the following command to install any missing dependencies.

```bash
go mod tidy
```

### 3. Build the Application

Run the `build.sh` script to build both the server and client binaries.

```bash
./scripts/build.sh
```

### 4. Run the Application Locally

Run the `run.sh` script to start both the server and client locally.

```bash
./scripts/run.sh
```

This will start the server in the background and open the client interface.

### 5. Run with Docker (Optional)

You can also use Docker to run the server and client as containers. To do so:

1. **Build the Docker images:**

   ```bash
   docker-compose build
   ```

2. **Start the application:**

   ```bash
   docker-compose up
   ```

This will start both the **server** and **client** in separate containers. The server will listen on port `8080`, and the client will connect to the server.

---

## 📈 **Future Enhancements**

- **Authentication:** Add basic authentication (e.g., login, password) for users.
- **Chat Rooms:** Allow clients to join specific chat rooms.
- **Message History:** Store messages in a database for retrieving chat history.
- **Encryption:** Implement SSL/TLS for secure communication between client and server.
- **WebSocket Support:** Extend the application with WebSocket support for real-time communication in a web-based interface.

---

## 📝 **License**

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 💬 **Contributing**

If you want to contribute to this project, feel free to fork the repository, create a branch, and submit a pull request. Contributions are always welcome!

---

## 🏷️ **Contact**

For any questions or suggestions, you can reach out to me at [your.email@example.com](mailto:your.email@example.com).

---

### **Enjoy coding! 🚀**
