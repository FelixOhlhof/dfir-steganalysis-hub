FROM nginx:alpine

COPY build /usr/share/nginx/html

COPY env.sh /docker-entrypoint.d/

RUN chmod +x /docker-entrypoint.d/env.sh

EXPOSE 80

CMD ["/docker-entrypoint.d/env.sh"]
