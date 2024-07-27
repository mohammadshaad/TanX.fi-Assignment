# Price Alert Application

## Overview
This application allows users to set price alerts for Bitcoin (BTC). When the target price is reached, an email notification is sent to the user.

## Features
- Create, delete, and fetch price alerts.
- Real-time price updates using Binance WebSocket.
- User authentication with JWT tokens.
- Email notifications using Gmail SMTP.
- Caching with Redis.
- Dockerized setup.

## Endpoints
### Create Alert
`POST /alerts/create`
```json
{
    "user_id": "1",
    "coin": "BTC",
    "target_price": 33000.00
}
Delete Alert
DELETE /alerts/delete/{id}

Fetch Alerts
GET /alerts/fetch?status=triggered&page=1&limit=10

Running the Project
Clone the repository.
Create a .env file with your environment variables.
Run docker-compose up to start the application.
Sending Alerts
The application listens for price updates and triggers alerts when the target price is reached. It uses RabbitMQ to manage email notifications.

This is a high-level overview to get you started. Each section will need further elaboration and error handling to be production-ready.





