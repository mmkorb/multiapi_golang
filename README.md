Go Backend with RabbitMQ, REST API, and PostgreSQL
This is a Go backend that supports sending and receiving messages via RabbitMQ, exposes a REST API documented with Swagger, and uses PostgreSQL as the database.

📁 Project Structure

📦 backend-project
├── 📂 cmd                  # Main applications
│   ├── 📂 rabbitmq_consumer # RabbitMQ message consumer worker
│   └── 📂 rest_api          # REST API server
├── 📂 configs              # Project configurations
│   ├── config.go           # Configuration file
├── 📂 internal             # Application domain logic
│   ├── 📂 domain           # Entity definitions
│   ├── 📂 repository       # Repository implementations
│   └── 📂 service          # Business rules and services
├── 📂 pkg                  # Reusable packages
│   ├── 📂 datastore        # Data storage connection
│   ├── 📂 db               # Database setup
│   └── 📂 rabbitmq         # RabbitMQ connection implementation
├── go.mod                  # Dependency management
├── go.sum                  # Dependency checksums
├── README.md               # Project documentation

2️⃣ Database Setup
The project uses PostgreSQL. To start the database using Docker, run:
'''docker-compose up -d'''

3️⃣ Running the Application
📌 Starting the REST API and the RABBITMQ Worker
'''go run cmd/main.go'''

🛠️ API Endpoints
The Swagger documentation is automatically generated and can be accessed at:
📌 http://localhost:8080/swagger/index.html

🏗️ Project Folder Structure Explained
cmd/ → Contains the main executables (REST API and RabbitMQ consumer).
configs/ → Manages global configurations.
internal/ → Contains business logic, repositories, and services.
pkg/ → Reusable packages such as database and messaging connections.

📝 License
This project is licensed under the MIT license. Feel free to use and contribute!
