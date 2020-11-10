FROM node:12
WORKDIR /app
COPY package.json package.json
COPY package-lock.json package-lock.json
RUN npm ci
COPY . . 
RUN npm run build
CMD npx serve dist -l tcp://0.0.0.0:3000
