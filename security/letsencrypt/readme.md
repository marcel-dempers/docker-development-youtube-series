

```
docker run -it -v ${PWD}:/work -w /work nginx bash
```

```
docker run -it --rm --name nginx -v ${PWD}/nginx.conf:/etc/nginx/nginx.conf `
-v ${PWD}:/letsencrypt `
-v ${PWD}/certs:/etc/letsencrypt `
-p 80:80 `
-p 443:443 `
nginx

```


```
docker build . -t certbot

docker run -it --rm --name certbot `
-v ${PWD}:/letsencrypt `
-v ${PWD}/certs:/etc/letsencrypt `
certbot bash

certbot certonly --webroot
```


```
IMPORTANT NOTES:
 - Congratulations! Your certificate and chain have been saved at:
   /etc/letsencrypt/live/marcel.guru/fullchain.pem
   Your key file has been saved at:
   /etc/letsencrypt/live/marcel.guru/privkey.pem
   Your cert will expire on 2020-12-03. To obtain a new or tweaked
   version of this certificate in the future, simply run certbot
   again. To non-interactively renew *all* of your certificates, run
   "certbot renew"
 - Your account credentials have been saved in your Certbot
   configuration directory at /etc/letsencrypt. You should make a
   secure backup of this folder now. This configuration directory will
   also contain certificates and private keys obtained by Certbot so
   making regular backups of this folder is ideal.
 - If you like Certbot, please consider supporting our work by:

   Donating to ISRG / Let's Encrypt:   https://letsencrypt.org/donate
   Donating to EFF:                    https://eff.org/donate-le

```


