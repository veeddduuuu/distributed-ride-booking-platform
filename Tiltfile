# --- API Gateway Configuration ---
docker_build('api-gateway', '.', dockerfile='services/api-gateway/Dockerfile')
k8s_yaml('k8s/api-gateway.yaml')
k8s_resource('api-gateway', port_forwards=8080)

# --- Trip Service Configuration ---
docker_build('trip-service', '.', dockerfile='services/trip-service/Dockerfile')
k8s_yaml('k8s/trip-service.yaml')
k8s_resource('trip-service', port_forwards=8083)

# --- Web Frontend Configuration ---
docker_build('web', 'services/web', dockerfile='services/web/Dockerfile')
k8s_yaml('k8s/web.yaml')
k8s_resource('web', port_forwards='3000:80')