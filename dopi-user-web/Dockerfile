FROM node:lts-alpine as build
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

FROM arm64v8/nginx
#FROM nginx
COPY --from=build /app/dist /usr/share/nginx/html/users
COPY nginx.default.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
