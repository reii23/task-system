# build stage
FROM node:20-alpine AS builder
WORKDIR /app
RUN apk add --no-cache python3 make g++
COPY . .
RUN npm install
RUN npm run build

# production stage
FROM node:20-alpine AS production
WORKDIR /app
COPY --from=builder /app/.output .
RUN npm install -g serve
EXPOSE 3000
CMD ["serve", "-s", ".", "-l", "3000"]

