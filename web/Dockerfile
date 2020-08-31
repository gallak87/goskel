# Copied from: https://mherman.org/blog/dockerizing-a-react-app/
FROM node:12.18.3-alpine
LABEL maintainer="George Allakhverdyan <george.a@outlook.com>"

WORKDIR /app

ENV PATH /app/node_modules.bin:$PATH
COPY package.json yarn.lock ./
# TODO: parameterize react-script version
RUN yarn install --silent
RUN yarn global add react-scripts@3.4.3 --silent

# TODO: use yarn build instead for a production build, dev only atm

EXPOSE 3000

COPY . ./
CMD [ "yarn", "start" ]