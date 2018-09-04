FROM node

WORKDIR /opt/app
CMD ["npm", "start"]

COPY package.json /opt/app
COPY dist /opt/app
