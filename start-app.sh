#!/bin/bash

echo "Starting Blinket Gap Filler application..."

# Start the backend server
echo "Starting Go backend server..."
(cd backend && go run main.go) &

# Start the frontend server
echo "Starting React frontend server..."
(cd frontend && npm start) &

echo "Both servers are now running!"
echo "- Backend: http://localhost:8080"
echo "- Frontend: http://localhost:3000"

# Wait for both processes to finish
wait