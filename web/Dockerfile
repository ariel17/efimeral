FROM node:12.6.0-buster as build
WORKDIR /app
COPY . .
ENV REACT_APP_API_URL=http://localhost:8080
ENV REACT_APP_DOCS_URL=http://localhost:9090
ENV PUBLIC_URL=
RUN echo "REACT_APP_API_URL=${REACT_APP_API_URL}" > .env
RUN echo "REACT_APP_DOCS_URL=${REACT_APP_DOCS_URL}" > .env
RUN echo "PUBLIC_URL=${PUBLIC_URL}" >> .env
RUN yarn install
RUN yarn build


FROM nginx:latest
COPY --from=build /app/build/ /usr/share/nginx/html/
RUN rm /etc/nginx/conf.d/default.conf
COPY nginx/nginx.conf /etc/nginx/conf.d
