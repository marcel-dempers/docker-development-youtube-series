FROM nginx:1.19-alpine

#config
COPY ./nginx.conf /etc/nginx/nginx.conf

#content
COPY ./*.html /usr/share/nginx/html/
COPY ./*.css /usr/share/nginx/html/
