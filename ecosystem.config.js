module.exports = {
  apps: [
    // Backend Go service
    {
      name: 'wallet-backend',
      script: 'go',
      args: 'run cmd/main.go',
      cwd: './backend',
      instances: 1,
      autorestart: true,
      watch: false, // Set to true if you want to watch for file changes
      max_memory_restart: '1G',
      env: {
        NODE_ENV: 'development',
        GO_ENV: 'development',
        PORT: '3000'
      },
      env_production: {
        NODE_ENV: 'production',
        GO_ENV: 'production',
        PORT: '3000'
      },
      error_file: './logs/backend-error.log',
      out_file: './logs/backend-out.log',
      log_file: './logs/backend-combined.log',
      time: true
    },
    // Frontend Vue.js service
    {
      name: 'wallet-frontend',
      script: 'npm',
      args: 'run dev',
      cwd: './frontend',
      instances: 1,
      autorestart: true,
      watch: false,
      max_memory_restart: '500M',
      env: {
        NODE_ENV: 'development',
        PORT: '3000'
      },
      env_production: {
        NODE_ENV: 'production',
        PORT: '3000'
      },
      error_file: './logs/frontend-error.log',
      out_file: './logs/frontend-out.log',
      log_file: './logs/frontend-combined.log',
      time: true
    }
  ],

  // Deployment configuration (optional)
  deploy: {
    production: {
      user: 'agnarsong',
      host: 'localhost',
      ref: 'origin/main',
      repo: 'git@github.com:username/wallet-manager.git',
      path: '/var/www/wallet-manager',
      'pre-deploy-local': '',
      'post-deploy': 'npm install && pm2 reload ecosystem.config.js --env production',
      'pre-setup': ''
    }
  }
};