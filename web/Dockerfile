FROM node as build-stage-front
WORKDIR /app
COPY client/package*.json ./
RUN npm install --production && npm audit fix
COPY client .
RUN npm run build

FROM nginx:stable-alpine as production-stage
WORKDIR /app
COPY --from=build-stage-front /app/dist /usr/share/nginx/html
COPY client/nginx/nginx.conf /etc/nginx/nginx.conf
COPY client/nginx/default.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

#COPY docker-entrypoint.sh /
#RUN chmod 755 /docker-entrypoint.sh
#ENTRYPOINT ["sh","/docker-entrypoint.sh"]

CMD ["nginx", "-g", "daemon off;"]

