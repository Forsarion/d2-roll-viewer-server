FROM node:8-slim as environment

WORKDIR /app
RUN yarn config set cache-folder /yarn_cache

FROM environment as builder

COPY package.json ./
COPY yarn.lock ./

RUN yarn install --frozen-lockfile && yarn cache clean

COPY . .

RUN yarn build

FROM zzswang/docker-nginx-react as production

ENV DEBUG=off

COPY --from=builder /app/build /app