FROM ubuntu:24.04

# -------------------------------------------------------
# Same as Chapter 4: install prerequisites
# -------------------------------------------------------
RUN apt update && apt install -y \
    curl \
    gnupg2 \
    ca-certificates \
    lsb-release \
    ubuntu-keyring \
    libcap2-bin

# -------------------------------------------------------
# Same as Chapter 4: add the official NGINX signing key
# so apt can verify the package authenticity
# -------------------------------------------------------
RUN curl https://nginx.org/keys/nginx_signing.key | gpg --dearmor \
    | tee /usr/share/keyrings/nginx-archive-keyring.gpg >/dev/null

# -------------------------------------------------------
# Same as Chapter 4: add the NGINX apt repository
# -------------------------------------------------------
RUN echo "deb [signed-by=/usr/share/keyrings/nginx-archive-keyring.gpg] \
http://nginx.org/packages/ubuntu $(lsb_release -cs) nginx" \
    | tee /etc/apt/sources.list.d/nginx.list

# -------------------------------------------------------
# Same as Chapter 4: pin NGINX packages to prefer the
# official nginx.org packages over the Ubuntu defaults
# -------------------------------------------------------
RUN printf "Package: *\nPin: origin nginx.org\nPin: release o=nginx\nPin-Priority: 900\n" \
    | tee /etc/apt/preferences.d/99nginx

# -------------------------------------------------------
# Same as Chapter 4: install NGINX
# (apt also creates the unprivileged 'nginx' system user)
# -------------------------------------------------------
RUN apt update && apt install -y nginx

# -------------------------------------------------------
# Same as Chapter 4: grant the nginx binary the capability
# to bind to privileged ports (< 1024) as a non-root user
# -------------------------------------------------------
RUN setcap 'cap_net_bind_service=+ep' /usr/sbin/nginx

# -------------------------------------------------------
# Same as Chapter 4: create the website directory and
# hand ownership to the nginx user
# -------------------------------------------------------
RUN mkdir -p /websites/my-website \
    && chown -R nginx:nginx /websites/my-website

# -------------------------------------------------------
# Copy our NGINX config and website files into the image
# -------------------------------------------------------
COPY --chown=nginx:nginx nginx.conf  /websites/my-website/nginx.conf
COPY --chown=nginx:nginx my-website/home.html   /websites/my-website/home.html
COPY --chown=nginx:nginx my-website/style.css   /websites/my-website/style.css
COPY --chown=nginx:nginx my-website/script.js   /websites/my-website/script.js

# -------------------------------------------------------
# Same as Chapter 4: run as the unprivileged nginx user
# mirrors: sudo -u nginx nginx -c /websites/my-website/nginx.conf
# -------------------------------------------------------
USER nginx

EXPOSE 80

# 'daemon off' keeps nginx in the foreground.
# Containers have no systemd, so the process must not
# daemonize or the container will immediately exit.
CMD ["nginx", "-c", "/websites/my-website/nginx.conf", "-g", "daemon off;"]
