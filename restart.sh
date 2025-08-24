#!/bin/bash

echo "🔄 Restarting Wallet Manager services..."

# Restart all services
pm2 restart ecosystem.config.js

# Show status
pm2 status

echo "✅ Services restarted!"
echo ""
echo "Frontend: http://localhost:5173"
echo "Backend API: http://localhost:3000"