FROM node:9.4 as build
WORKDIR /app
COPY . .
RUN yarn build


FROM nginx:latest
COPY --from=build /app/build/ /usr/share/nginx/html/
RUN rm /etc/nginx/conf.d/default.conf
COPY nginx/nginx.conf /etc/nginx/conf.d
