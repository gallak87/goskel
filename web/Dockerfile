# Copied from: https://mherman.org/blog/dockerizing-a-react-app/
FROM node:16.17.0-alpine
LABEL maintainer="George Allakhverdyan <george.a@outlook.com>"

### NOTE: to support hot reloads, use a volume
# issue right now is i can't get volumes: - ./web:/app to work in docker-compose!
#VOLUME [ "/app" ]
WORKDIR /app
ENV PATH /app/node_modules/.bin:$PATH

# TODO: parameterize react-script version
COPY package.json yarn.lock ./
RUN yarn install --silent

RUN echo HELLO > test.txt

# TODO: use yarn build instead for a production build, dev only atm

EXPOSE 3000

COPY . ./
CMD [ "yarn", "start" ]
