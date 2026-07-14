# --- API Gateway Configuration ---
local_resource('api-gateway', 
               cmd='go run ./services/api-gateway', 
               deps=['services/api-gateway', 'shared'])

# --- Trip Service Configuration ---
local_resource('trip-service', 
               cmd='go run ./services/trip-service/cmd/main.go', 
               deps=['services/trip-service', 'shared'])

# --- Web Frontend Configuration ---
local_resource('web', 
               cmd='npm run dev', 
               dir='services/web', 
               deps=['services/web/src', 'services/web/package.json', 'services/web/vite.config.ts'])