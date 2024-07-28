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
```

### Delete Alert
`DELETE /alerts/delete`
```json
{
    "user_id": "1",
    "coin": "BTC"
}
```

### Fetch Alerts
`GET /alerts/fetch?user_id=1`
```json
{
    "user_id": "1"
}
```

## Setup
1. Clone the repository.
2. Create a `.env` file in the root directory and add the following environment variables:
```bash
POSTGRES_USER: user
POSTGRES_PASSWORD: password
POSTGRES_DB: price_alert
```

3. Run the following command to start the application:
```bash
docker-compose up
```

4. The application will be running on `http://localhost:8000`.

## Technologies
- Golang
- PostgreSQL
- Redis
- Docker
- Binance WebSocket
- JWT

## Thank You
Made with ❤️ by [Mohammad Shaad](https://github.com/mohammadshaad).