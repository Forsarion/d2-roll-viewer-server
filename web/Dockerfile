FROM node:8-slim as runner

WORKDIR /app
EXPOSE 80

FROM runner

COPY package.json ./
COPY yarn.lock ./

RUN yarn install --frozen-lockfile && yarn cache clean

COPY . .

RUN npm run build

ARG BUILD_NUMBER=unknown
ARG GIT_COMMIT=unknown
ARG GIT_BRANCH=unknown

LABEL com.jenkins.build=${BUILD_NUMBER}
LABEL com.git.commit=${GIT_COMMIT}
LABEL com.git.branch=${GIT_BRANCH}

ENV BUILD_NUMBER=${BUILD_NUMBER} GIT_COMMIT=${GIT_COMMIT} GIT_BRANCH=${GIT_BRANCH}

CMD [ "npm", "start" ]