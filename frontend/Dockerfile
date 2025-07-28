# Build stage
FROM node:20-alpine AS build
WORKDIR /app
COPY frontend/package*.json ./
RUN npm install
COPY frontend/ ./
RUN npm run build -- --configuration production

# Serve stage
FROM nginx:alpine
COPY --from=build /app/dist/car-rental /usr/share/nginx/html
COPY frontend/nginx.conf /etc/nginx/conf.d/default.conf