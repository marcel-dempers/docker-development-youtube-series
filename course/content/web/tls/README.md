# 🎬 The Basics of SSL & TLS

## 💡 Preface

This module is part of a course on DevOps. </br>
Checkout the [course introduction](../../../README.md) for more information </br>
This module is part of [chapter 4](../../../chapters/chapter-4-web-and-http/README.md)

This module draws from my extensive experience in the world of Web. Building and managing Web servers, building web sites, web services, monitoring , security, networking etc.  
In this module, we will cover the basics of SSL and TLS and the importance of it in the DevOps field. </br>
SSL & TLS goes quite deep in terms of complexity and is a very broad topic in the security landscape. </br>
The aim of this module is to highlight important aspects, basic usage and tools for engineers. </br>

As an engineer, you will encounter TLS\SSL when working with HTTPS. 
This may involve TLS certificates expiring, rotation, maintaining and understanding HTTPS, the importance of encryption and troubleshooting HTTPS.

## What is TLS & SSL

In our overview of the web, we talked about client and server and that they both communicate using the HTTP protocol. Much of this chapter covered the HTTP protocol and we mentioned that HTTP generally occurs over TCP and over port 80 and 443. </br>

Generally speaking HTTP requests and responses are unencrypted when using HTTP on port 80. </br>
Which means the data between the client and servers are passed in plain text which is readable by humans and any network devices in between. </br>

That can be problematic especially when an HTTP client needs to pass sensitive information to an HTTP server, like passwords (for authentication) or credit card information (such as on payment webpages)

### What is encryption

Encryption is a very broad and complex topic which is out of scope for this course. </br>
In basic terms encryption is like taking your secret message, and placing it in an unbreakable envelope that only an intended receiver can open for reading. 

Encryption is performed by taking a message, using a "secret key" and performing an algorythm on the message using the "secret key" which generates a scrambled version of our message. This scrambled message that nobody can read is called the "cipher text"
We can then safely send our ciphered message to our receiver and the receiver can unscramble the message using the same "secret key". This secret key is known as an **encryption key**.

---

### TLS & SSL Encryption

Most time in HTTP, we need to pass sensitive information. Things like personal information when using websites, log-in information such as usernames, passwords, API keys and tokens as well as payment pages that accept credit card information. </br> 

Just like we saw in the browser network tools and `curl` , any network device where this traffic passes through, can read the information if its unencrypted. </br>

This is not only concerning for sensitive information like passwords and credit card information, but overall privacy in general. <br/>

This is why we have HTTPS, which is HTTP, but its encrypted. The "S" stands for "Secure"
In order or HTTP traffic to be secure over HTTPS, the server would pass a TLS or SSL "handshake" </br>

SSL stands for `Secure Sockets Layer` </br>
TLS stands for Transport Layer Security</br>

TLS is a protocol used to encrypt communication between applications and systems, such as web browsers and servers, to ensure secure data transmission. </br>
Pretty much does what SSL does, however TLS is a successor to SSL and provides a secure, encrypted channel for data exchange.


SSL and TLS are cryptographic algorythms that allow us to encrypt HTTP payloads. </br>
It's important to know that SSL is an older technology and technically vulnerable and superceeded by TLS. People still often refer to TLS as SSL because the term has been used for so long and the term has become interchangable </br>

From now on throughout this module and course, we will refer to TLS only

For a client and server to encrypt HTTP traffic, a TLS certificate is used and is generally installed on the server side. 
The client and server will perform a "TLS handshake" and during this process the client and server will exchange messages using TCP to verify the TLS certificate
and validate ciphers and establish a secure encrypted channel to communicate over. 

## HTTPS Connections with TLS

### The TLS Handshake

To summarise all the above, lets run a `curl` command to see the TLS handshake:

```
curl -v https://www.google.com
```

The TLS handshake can be see in the `curl` output at the top. It is basically made up of the following steps:

* Client and Server Hello (**Initiation**)
  - The client and server propose TLS versions and Cipher suites to use for the HTTPS connection
* Client and Server Key\certificate verification (**Verification**)
  - process of the client verifying the server certificate
  - client also ensures that the "Host" name in the HTTP Host header, for the HTTP request, matches the Common Name in the certificate
    - For example: HTTP Request Header `Host: www.google.com` matches the certificate Common Name `www.google.com` 
  - client checks if the certificate authority who issued the certificate is one that it trusts
  - client checks if the certificate was issued to `www.google.com`
* Secure Channel Established (**Established**)
  - Both client and server calculates whether everything went well
* Encrypted Data Transfer (**Transfer**)
  - With the secure channel active, the client sends its HTTP request (e.g., GET /) and the server sends its HTTP response, both fully encrypted.

* The HTTP Headers and Body are encrypted during HTTPS communications

### Self-signed TLS certificatte VS CA Certificate

Trusted certificates are signed by a trusted certificate authority also known as "CA"</br>
Let's encrypt is a non-profit CA that provides free certificates. Other commercial CAs (e.g., DigiCert, GlobalSign) offer paid certificates. </br>

Trusted certificates are the ones used widely on the public internet for HTTPS that most browsers and operating systems trust. </br>

Self Signed certificates are certificates that are signed by yourself and are not trusted by any client, device or browser. </br>

The main important points about a self-signed certificate:

* **Who issues it**? 
  - You do! As the name suggests, you create and sign the certificate yourself using your own private key. There's no third party involved.
* **Trust**
  - This is the biggest hurdle. Because you are the "authority" that issued the certificate, no external entity (like web browsers or operating systems) inherently trusts it

* **Use Cases**:
  - Development and testing environments.
  - Internal networks or services where you control all client devices and can manually distribute the certificate.
  - Situations where you need encryption but public trust isn't a concern.

## TLS Tools 

There are some great tools available to 
* Generate self-signed TLS certificates that we can use for testing and development
* Test TLS certificates for HTTPS </br>
  - We can test a website address to validate its certificate. We may want to validate a website and it's certificate before we send traffic to it.

### Creating a self-signed TLS certificate 

In this module we will generate a self-signed certificate which is something that may come in handy for you throughout your DevOps career when its required. </br>

In a previous module, we created a web server and were able to access our website over HTTP.
Let's do this over HTTPS to see what it takes to configure a web server with a TLS certificate and get an idea of how clients handle self signed certificates

Let's fire up our server from Chapter 2, and run the following command to create a self-signed certificate and private key for the certificate

```
openssl req -x509 \
-nodes \
-days 365 \
-newkey rsa:2048 \
-keyout server.key \
-out server.crt \
-subj "/CN=localhost"

```

<i>**Tip:** For the above command I introduced something new to your command line knowledge. Notice the `\` character on the end of each line. This is very handy for long commands and helps you break long one-liners into multiple lines so it becomes a little more readable  </i>

* `openssl req`: Command to create and process certificate requests.
* `-x509`: Output a self-signed certificate instead of a certificate request.
* `-nodes`: Do not encrypt the private key (no passphrase).
* `-days 365`: Make the certificate valid for 365 days.
* `-newkey rsa:2048`: Generate a new RSA private key of 2048 bits.
* `-keyout server.key`: Output the private key to server.key.
* `-out server.crt`: Output the certificate to server.crt.
* `-subj "/CN=localhost"`: Automatically set the Common Name to localhost. This is important for matching the domain you'll connect to.

You should now have two files in your current directory after running this command.
* `server.key`: your private key
* `server.crt`: your self-signed certificate

## Setting up TLS in Web servers for HTTPS

To setup TLS connections, you will need a valid and trusted TLS certificate </br>
For learning purposes, we have introduced a self-signed certificate.

### Move certificate and key to a secure location 

Move our certificates to the NGINX folder and give the `nginx` user correct permissions to the certificate and ensuring `root` has ownership:

```
sudo mkdir -p /etc/nginx/ssl
sudo mv server.crt /etc/nginx/ssl/
sudo mv server.key /etc/nginx/ssl/

sudo chmod 600 /etc/nginx/ssl/server.key
sudo chown root:root /etc/nginx/ssl/server.key
sudo chown root:root /etc/nginx/ssl/server.crt
```

### Edit web server configuration file

We will need to enable TLS within our web server configuration. </br>
This generally involves setting our port to `443` , enabling `TLS\SSL` and providing a key and certificate file. </br>

Let's edit our NGINX configuration file. Remember we used `systemd` to configure our web server so it manages `nginx`. </br>
We used a custom configuration file. </br>
We can locate it, using `sudo systemctl cat nginx.service` and we can see it tells us where our configuration is pointing to. 

We can edit this using `nano` 

```
sudo nano /websites/my-website/nginx.conf
```

We will need to add the following configuration items in the `server` block

```
server {
    listen 443 ssl;                                # Listen on port 443 for SSL/TLS traffic
    listen [::]:443 ssl;                           # Listen on IPv6 as well

    server_name localhost;                         # Important: Matches the CN in your certificate

    ssl_certificate /etc/nginx/ssl/server.crt;     # Path to your certificate
    ssl_certificate_key /etc/nginx/ssl/server.key; # Path to your private key

    # Optional: Stronger SSL/TLS settings
    ssl_protocols TLSv1.2 TLSv1.3;

}
```

Next, we can test our configuration to ensure there are no errors 
We should see output indicating a success

```
sudo nginx -t -c /websites/my-website/nginx.conf
```

Restart our NGINX web server so it takes the new configuration changes

```
sudo systemctl restart nginx
```

### Test our HTTPS server with `curl`

Now that Nginx is serving HTTPS traffic with your self-signed certificate, you can test it with `curl`. </br>
Since the certificate is self-signed (not issued by a publicly trusted Certificate Authority), `curl` won't inherently trust it.  </br>

If you run `curl https://localhost` you will see the `curl` throws an error to output:

```
curl: (60) SSL certificate problem: self signed certificate
More details here: https://curl.haxx.se/docs/sslcerts.html

curl failed to verify the legitimacy of the server and therefore could not
establish a secure connection to it. To learn more about this situation and
how to fix it, please visit the web page mentioned above.
```

This is because by default, operating systems, devices and clients like `curl` and web browsers will not pass the TLS handshake where self-signed certificates are used. </br>

We have two main options:

**Option A**: Trust the specific certificate (`--cacert`)
This is more secure as it explicitly trusts only your server.crt.


```
curl --cacert /etc/nginx/ssl/server.crt https://localhost/
```

`--cacert /etc/nginx/ssl/server.crt` Tells `curl` to trust the certificate located at this path as a Certificate Authority for this connection.

**Option B**:  Insecure connection (bypassing certificate validation `-k` or `--insecure`)
This is less secure and should only be used for testing or in environments where you explicitly understand the risk.

```
curl -k https://localhost
```

Remember if we wanted to access our web server from our browser (outside the VM) we will need to forward ports.
We use port `8080` forwarded to port `80` for HTTP </br>
We can forward port `8081` to port `443` for HTTPS </br>

# Key Takeaways 

* TLS & Encryption is a very broad topic under security and cryptography. Although DevOps will be dealing with Web, HTTPS, TLS and Web server configuration, we will often encounter it. </br>
  - It would become an expectation to have a high level understanding about everything covered in this module and a deeper understanding will always come with experience.
* Understand that TLS handshakes exist and its overhead
* Understand whats involved in making TLS work with a client and a web server as demonstrated in this module
* Know the difference between a self-signed created certificate and a CA-signed trusted certificate. 
