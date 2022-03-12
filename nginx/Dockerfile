FROM nginx:alpine

COPY ./default.conf /etc/nginx/conf.d/default.conf
ADD ./cert/ /etc/nginx/certs/
RUN chown -R nginx:nginx /etc/nginx/certs/
RUN chown -R nginx:nginx /etc/nginx/
