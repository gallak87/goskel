# Copied from: https://mherman.org/blog/dockerizing-a-react-app/
FROM node:16.17.0-alpine
LABEL maintainer="George Allakhverdyan <george.a@outlook.com>"
# TODO: parameterize react-script version

WORKDIR /app
ENV PATH /app/node_modules/.bin:$PATH

# NOTE: The docker-compose file at the root adds a volume to the web folder
# effectively allowing hot-reloads to work. Somehow the architecture doesn't matter?
# Able to run both via docker and bare metal and hot reloads and the app just works /shrug
RUN yarn install --silent
EXPOSE 3000

COPY . ./
CMD [ "yarn", "start" ]
