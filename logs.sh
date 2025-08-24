#!/bin/bash

# Function to show help
show_help() {
    echo "📋 Log viewer for Wallet Manager"
    echo ""
    echo "Usage: ./logs.sh [option]"
    echo ""
    echo "Options:"
    echo "  all       - Show logs for all services (default)"
    echo "  backend   - Show backend logs only"
    echo "  frontend  - Show frontend logs only"
    echo "  follow    - Follow live logs for all services"
    echo "  follow-backend  - Follow live backend logs"
    echo "  follow-frontend - Follow live frontend logs"
    echo "  clear     - Clear all logs"
    echo "  help      - Show this help message"
    echo ""
    echo "Examples:"
    echo "  ./logs.sh"
    echo "  ./logs.sh backend"
    echo "  ./logs.sh follow"
}

case "${1:-all}" in
    "all")
        echo "📜 Showing logs for all services..."
        pm2 logs --lines 50
        ;;
    "backend")
        echo "📜 Showing backend logs..."
        pm2 logs wallet-backend --lines 50
        ;;
    "frontend")
        echo "📜 Showing frontend logs..."
        pm2 logs wallet-frontend --lines 50
        ;;
    "follow")
        echo "📜 Following logs for all services... (Press Ctrl+C to exit)"
        pm2 logs --lines 0
        ;;
    "follow-backend")
        echo "📜 Following backend logs... (Press Ctrl+C to exit)"
        pm2 logs wallet-backend --lines 0
        ;;
    "follow-frontend")
        echo "📜 Following frontend logs... (Press Ctrl+C to exit)"
        pm2 logs wallet-frontend --lines 0
        ;;
    "clear")
        echo "🗑️ Clearing all logs..."
        pm2 flush
        echo "✅ Logs cleared!"
        ;;
    "help")
        show_help
        ;;
    *)
        echo "❌ Unknown option: $1"
        show_help
        exit 1
        ;;
esac