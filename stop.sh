#!/bin/bash

echo "🛑 Stopping Wallet Manager services..."

# Stop all PM2 processes for this project
pm2 stop ecosystem.config.js

# Show status
pm2 status

echo "✅ All services stopped!"
echo ""
echo "To completely remove processes from PM2:"
echo "  pm2 delete ecosystem.config.js"
echo ""
echo "To start services again:"
echo "  ./start.sh"