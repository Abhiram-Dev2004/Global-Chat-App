Distributed Global Real-Time Chat System – Engineered a fault-tolerant, scalable chat platform using Golang, Kafka, Redis, Postgres, WebSockets, and Next.js, showcasing distributed event-driven architecture. Implemented user authentication, and HTTP-publish/WebSocket-subscribe real-time delivery for geographically distributed clients, designed to handle thousands of concurrent users and hundreds of messages/sec.

Message flow (data path)
HTTP publish
→ Kafka producer
→ Kafka topic
→ (parallel)
     → Kafka → DB
     → Kafka → Redis
→ Redis → server
→ server → WebSocket clients    


create a folder inside the apps/server, name it certs. place your ca.pem of you aiven postgres
create a .env folder, place a your postgres dsn-> dsn = "example1324"
obtain service.cert and servie.key form the aiven kakfa and place it in apps/server
obtain ca.pem from the aiven kafka and place it in apps/server

how to run?
1) go to apps/server
2) run go mod tidy
3) go mod download
4) go to root dir and run npm install

If you encounter this error: "Error TS6053: File '@repo/typescript-config/nextjs.json' not found."
just run : "npm install" in the root dir
