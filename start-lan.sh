#!/bin/bash

# Wallet Manager LAN Access Startup Script
# This script starts the services accessible from LAN

LOCAL_IP=$(ifconfig | grep "inet " | grep -v 127.0.0.1 | awk '{print $2}' | head -1)

echo "======================================"
echo "  Wallet Manager - LAN Mode"
echo "======================================"
echo ""
echo "🌐 Services will be accessible from:"
echo "   Local IP: $LOCAL_IP"
echo ""
echo "📱 Access URLs:"
echo "   Frontend: http://$LOCAL_IP:3000"
echo "   Backend:  http://$LOCAL_IP:8888"
echo ""
echo "📝 Other devices on your network can access these URLs"
echo "======================================"
echo ""

# Stop any existing PM2 processes
echo "🛑 Stopping existing services..."
pm2 stop ecosystem.config.js 2>/dev/null || true
pm2 delete ecosystem.config.js 2>/dev/null || true

# Start services
echo "🚀 Starting services..."
pm2 start ecosystem.config.js

# Show status
pm2 status

echo ""
echo "✅ Services started successfully!"
echo ""
echo "📱 Share these URLs with other devices on your network:"
echo "   Frontend: http://$LOCAL_IP:3000"
echo "   Backend:  http://$LOCAL_IP:8888"
echo ""
echo "📝 Useful commands:"
echo "  pm2 logs             - View logs"
echo "  pm2 restart all      - Restart services"
echo "  ./stop.sh            - Stop all services"
echo ""
