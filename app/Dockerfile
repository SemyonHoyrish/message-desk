FROM node:23.11.0-alpine3.21 AS builder

WORKDIR /build

COPY package.json package-lock.json .

RUN npm i

COPY . .

RUN npm run build


FROM nginx:1.28-alpine3.21 AS runner

RUN rm -rf /usr/share/nginx/html/*

COPY --from=builder /build/build /usr/share/nginx/html

