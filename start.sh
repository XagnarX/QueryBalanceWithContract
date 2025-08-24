#!/bin/bash

# Start all services with PM2
echo "🚀 Starting Wallet Manager services..."

# Stop any existing PM2 processes
pm2 stop ecosystem.config.js 2>/dev/null || true
pm2 delete ecosystem.config.js 2>/dev/null || true

# Start services
pm2 start ecosystem.config.js

# Show status
pm2 status

echo "✅ Services started!"
echo ""
echo "Frontend: http://localhost:5173"
echo "Backend API: http://localhost:3000"
echo ""
echo "Useful commands:"
echo "  pm2 status           - View status of all processes"
echo "  pm2 logs             - View logs for all processes"
echo "  pm2 logs wallet-backend   - View backend logs only"
echo "  pm2 logs wallet-frontend  - View frontend logs only"
echo "  pm2 restart all      - Restart all processes"
echo "  pm2 stop all         - Stop all processes"
echo "  ./stop.sh            - Stop all services"