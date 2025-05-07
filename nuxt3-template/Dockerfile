# build
FROM node:20-alpine AS builder
WORKDIR /app
RUN apk add --no-cache python3 make g++
COPY . .
RUN npm install
RUN npm run build

# prod
FROM node:20-alpine AS production
WORKDIR /app
COPY --from=builder /app/.output .output
COPY --from=builder /app/package.json .
COPY --from=builder /app/node_modules ./node_modules
EXPOSE 3000
CMD ["node", ".output/server/index.mjs"]

