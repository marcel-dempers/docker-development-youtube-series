# 🎬 Creating Web Serves with NGINX

## 💡 Preface

This module is part of a course on DevOps. </br>
Checkout the [course introduction](../../../README.md) for more information </br>
This module is part of [chapter 4](../../../chapters/chapter-4-web-servers/README.md)

This module draws from my extensive experience in the world of Web. Building and managing Web servers, building web sites, web services, monitoring , security, networking etc.  

## Web Server Overview

What is a web server? In Previous modules we covered client and servers in network architecture. </br>
A Server in network terms, is the one that receives a network requests. </br>
In the world of Web and HTTP, a server may be referred to as a "web server" </br>
That is a server that sends an HTTP type response.

![diagram](../html/html.drawio.svg)

<b>Some popular web servers are:</b>

* [NGINX](https://nginx.org/) : Known for its high performance, scalability, and reverse proxy capabilities.
* [Caddy](https://caddyserver.com/) : A modern web server with automatic HTTPS and simple configuration.
* [Microsoft IIS](https://www.iis.net/) (Internet Information Services) : A web server for Windows-based environments
* [Apache HTTP Server](https://httpd.apache.org/) : One of the most widely used web servers, known for its flexibility and extensive module support.

## Installing a Web server: NGINX

In this module, I have given preference to NGINX. The reason for this is that NGINX is not only a web server, but a proxy and load balancer. </br>
This means you can place NGINX in front of two web servers and send 50% of traffic to one, and 50% of traffic to the other. </br>
Load balancers are important for high availability. In a production environment we dont want to run just one web server but at least two. </br>

Therefore we can use NGINX for other things in this course too.

### Install NGINX

To install NGINX for this course, we will be using our virtual server we created in Chapter 2. </br>
We can navigate to the [NGINX](https://nginx.org/) site and to the documentation page where we will find the steps for installing NGINX on our server. </br>

We firstly need to install all pre-requisites for installation of NGINX

We will use 
* `curl` to fetch things from the internet
* `gnupg2`,`ca-certificates`, `lsb-release` and `ubuntu-keyring`  to work with keys to verify the NGINX installation packages for ubuntu

```
sudo apt update
sudo apt install -y curl gnupg2 ca-certificates lsb-release ubuntu-keyring
```

We have learned that the Linux operating system uses package managers to install things. In Ubuntu this is `apt`

To check the NGINX package authenticity we download the NGINX signing keys and let `apt` verify that the packages were signed with those keys. 

<b>Important Security Note </b>

In a world of DevOps and automation, its often tempting to write autonomous scripts that can fetch application binaries and installation files from the internet. </br>
Sometimes the servers where you download things from can be compromised, or your system, or even a system in the middle somewhere could be compromised and the files you are fetching may not be the same files you are expecting.
We should never simply assume that links can be trusted.

The way to work around this is that vendors may use keys to sign their files before placing it on the internet. </br>
When you use crpytographic software, you can take a file, combine it with a private key to produce a digital signature and a public key </br>

In simple terms, the vendor would keep the private key secret, and the digital signature alongside the public key can then be distributed. </br>

This means that anyone with the digital signature and the public key can verify that the original file has been signed by the vendor with a private key and has not been tampered with. </br>

If a malicious actor has tampered with the files you are downloading, you can use the original digital signature and public key and tell that the file is different. </br>

This process is very often used in software distribution to ensure clients who download files get authentic files and protect their systems from malicious files that may have been tampered with. 

We will follow this process with NGINX, so its important to observe and pay attention to the reasons behind these steps as you will come across this in your DevOps journey. 

According to NGINX documentation `Import an official nginx signing key so apt could verify the packages authenticity.`

```
curl https://nginx.org/keys/nginx_signing.key | gpg --dearmor \
    | sudo tee /usr/share/keyrings/nginx-archive-keyring.gpg >/dev/null
```

Verify that the downloaded file contains the proper key:

```
gpg --dry-run --quiet --no-keyring --import --import-options import-show /usr/share/keyrings/nginx-archive-keyring.gpg
```
The output should contain the full fingerprint 573BFD6B3D8FBC641079A6ABABF5BD827BD9BF62 as follows:
```
pub   rsa2048 2011-08-19 [SC] [expires: 2027-05-24]
      573BFD6B3D8FBC641079A6ABABF5BD827BD9BF62
uid                      nginx signing key <signing-key@nginx.com>
```

Note that the output can contain other keys used to sign the packages.

To set up the apt repository for stable nginx packages, run the following command:

```
echo "deb [signed-by=/usr/share/keyrings/nginx-archive-keyring.gpg] \
http://nginx.org/packages/ubuntu `lsb_release -cs` nginx" \
    | sudo tee /etc/apt/sources.list.d/nginx.list
```

Set up repository pinning to prefer our packages over distribution-provided ones:
```
echo -e "Package: *\nPin: origin nginx.org\nPin: release o=nginx\nPin-Priority: 900\n" \
    | sudo tee /etc/apt/preferences.d/99nginx
```

Now we can go ahead and upgrade our `apt` sources and install NGINX from the newly setup sources

```
sudo apt update
sudo apt install -y nginx
```

### Web Server Commands

Once NGINX is installed, we can learn about its command line features, just like every command line program
we have used in previous modules. You can now see why that was so important </br>

```
nginx -h
```

This will provide us with all the potential options or flags we can pass to the command. </br>

Simply just running command `nginx` will start the web server with default configuration values. </br>
We can see a few defaults defined in the help output. </br>
NGINX has helpful signals you can use to restart your web server and even reload a configuration after making a configuration change. NGINX will gracefully make that configuration change without affecting current traffic. </br>

Example:

```
# restarts our web server
sudo nginx -s reopen 

#reloads the configuration file
sudo nginx -s reload

```
We can start our NGINX by just running the `sudo nginx` command. We can then view the NGINX process using `top` and see it running. 

<i>
<b>Important Note:</b> Remember that our virtual server runs on a private network and have a private IP address. We cannot access all ports on our virtual machine because its running in its own network, so we need to setup a port-forward to forward port 80 to our servers port 80
</i>

Once we have a port forward in place -
We can access our site on `http://<IP>` where the IP is the private IP of our server. 

### Automatic Start on Server Boot 

It's impportant to also note that when NGINX installs using `apt`, 

This is because the `apt` package manager, when installing system services like Nginx, typically sets them up to be managed by `systemd` (the modern init system in Ubuntu). Part of this setup usually involves enabling the service, which means it will automatically start at boot.

You can verify this by running: </br>

`systemctl status nginx`

If it says `Loaded: loaded (/lib/systemd/system/nginx.service; enabled; vendor preset: enabled)`, then it's configured to start automatically. The "enabled" status indicates that it will start on boot. </br>

`systemctl` is a command-line utility used to control and inspect the systemd system and service manager in Linux.

#### systemd

`systemd` is the "brain" or the "engine" that manages nearly everything that happens on a modern Linux system after the kernel loads. </br> 
It's the first process that starts (PID 1) and is responsible for booting up the rest of the operating system, managing system services (like web servers, databases, networking, etc.), logging, and shutting down the system. It replaced older "init" systems like System V init.

`systemctl` is the command line tool for interacting with `systemd`


### NGINX Configuration

All web servers generally come with configuration support </br>
Configuration files are what we use to setup the web server, tell it what port to run on, what type of traffic to accept and what to do with that traffic </br>

These will be our intentions.
As a DevOps engineer pay close attention to the intentions rather than the implementation here. </br>
In this module we use NGINX, but the same outcomes could be achieved by using a different web server. 

#### Configuration Structure

NGINX has its own type of configuration structure. It's made up of "simple directives" which are basically names of the configuration options followed by their parameter values separated by spaces and ending in a semicolon. 

For example: `listen 80;` is a simple directive, that tells NGINX what port to listen on when running. </br>
The parameter name is `listen` and its parameter value is `80` and NGINX knows where the directive ends by looking for the `;` semicolon character. </br>

All the NGINX options are defined in simple directives. 

The second type of directive is called the "block directive"  A block directive starts and ends with brackets, and contains other block directives and simple directives

For example: ` <name-of-directive> { <content> }`

An empty `http` directive:
```
http {
}
```
You can have one directive in another, `server` inside `http` 

```
http {
  server {
  }
}
```

Furthermore, you can have simple directives inside block directives

```
http {
  server {
    listen 80;
  }
}
```

It's important to know with NGINX you cannot place any directive inside others.
The outer directives or "main" directives can only be `events` and `http`, like the example above. </br>

`server` is allowed in `http`, and `location` in `server`.

We would generaly have a config like this:

```
http {
  server {
    listen 80;

    location {
        # rules go here !
    }
  }
}
```

The directives are used for the following:

* `http` : Global server settings, holds `server` and `location` blocks
* `server` : Used to define a virtual server block, which essentially acts as a separate web server within the same Nginx instance. It's crucial for handling multiple websites or applications from the same server
* `location` : Is a configuration directive that specifies how the server should handle requests for specific URLs or URL patterns

Example of a complete configuration file:

```
http {
  server {
    listen 80;
    location  / {
        root /webites/my-website;
    }
  }
}
```

So far we've been looking at block directives and nested ones where we place block directives in other block directives as above. We also have "simple directives" above, such as `listen` and `root` 

* `listen` : The port that the virtual server should be listening on 
* `root` : The directory that NGINX will serve request files from. Like the `html`, `.css` , `.js` and other files

#### Multiple Websites

You can see the block directives in action when you start hosting multiple web sites on one NGINX instance.
We can have multiple `server` block directives, each one for a specific website running on different ports 

We can use the `server_name` simple directive to tell NGINX which domain to serve traffic from. So if the domain in the URL matches the server name of a `server` block, then that `server` block will accept the traffic and apply its rules to it.

```
http {
  server {
    listen 80;
    server_name my-website.com;
    location  / {
        root /webites/my-website;
    }
  }

  server {
    listen 81;
    server_name my-other-website.com;
    location  / {
        root /webites/my-other-website;
    }
  }
}
```

### Security Considerations 

As a DevOps, SRE or platform engineer, when running web servers, you need to always be mindful that web servers can be public facing. Having a port open on a server that listens for TCP or HTTP traffic, opens up an attack vector.

Not all traffic coming to your web server will be of good intent. There will be bots, scrapers, especially in the world of AI, bots will frequently crawl your sites and scrape content. Malicious actors may also send requests to your web server in search of vulnerabilities. Dealing with that is a whole other challenge. </br>

We have to take initiative to ensure we apply security best practises when running a web server. We need to ensure our server runs as securely as possible and many engineers overlook security. </br>

During the configuration of our server we will do the following:

* Use `sudo` privileges to install the web server as its needed
* Ensure our web server runs as a non-privilged, non-root user. 
  * Imagine a scenario where there is a vulnerability in your web server that allows certain malicious HTTP requests to contain a payload that can be executed on your web server. An attacker could use this vulnerability to send a "bash" payload and run malicious commands on your web server. If your web server runs as `root`, the "bash" payload can cause a lot more harm on your server. Running as a non-root unprivileged user, means the vulnerability cannot be used to do a lot of damage, so it reduces the attack vector. 
* We'll need to make additional configurations to allow running as an unprivileged user, for example:
  * Create the unprivileged user
  * Provide access to the folder where the website content will be. I.E `/websites`
  * Provide access to allow the user the run an application that listens on port 80 and 443. By default in Linux, all ports lower than 1024 are privileged ports, so we need to add capabilities to allow our unprivileged user to use a port lower than 1024. 

#### Create a user for our webserver

To ensure we run our web server with a non-privileged user, lets go ahead and create a new user for NGINX called `nginx`. Similar to our Linux module:

```
sudo useradd --system --no-create-home --shell /usr/sbin/nologin nginx
```

* `--system` : Creates a system user.
* `--no-create-home` : Prevents creating a home directory for the user.
*`--shell /usr/sbin/nologin` : Disables login for the user.

The directory `/usr/sbin` holds executable programs primarily used by system administrators for tasks like system maintenance and configuration. 

Now this user should already exist as it's created when installing `nginx`, but it's good to learn how to create these types of non privileged users for applications in Linux. </br>

### Configure Our Webserver

Firstly we need to locate the default NGINX configuration. The command `nginx -h` will tell us. </b>
For web servers and software in general, most software have default configurations in default location.
We can either replace the default file with our own, or explicitly set a path to a new file of our own. </br>

In my opinion its best to be explicit, incase default values and default locations change over time.  </br>

Since we would like our website under `/websites/my-website`, we can place our configuration file there and tell `nginx` about it when we start it up with the `-c` option.

For example:

```
sudo nginx -c /websites/my-website/nginx.conf
```

When we run this, NGINX will complain that the file does not exist, so lets go ahead and create it.
Also remember that because we are not `root`, if we want to setup our website in a directory outside our home directory, we need to use `sudo`

```
# create our website folder
sudo mkdir -p /websites/my-website/

# create a blank config file
sudo touch /websites/my-website/nginx.conf

```

Now that we have our website folder, we need to create nessasary permissions for our `nginx` user to allow it can access and serve files in our website directory

```
sudo chown -R nginx:nginx /websites/my-website
```

Let's use `nano` to edit our confinguration and paste the following content into our new configuration file

```
# use nano set out configuration
sudo nano /websites/my-website/nginx.conf
```

```
user nginx;
pid /tmp/nginx.pid;

events {
  worker_connections  1024;
}

error_log  /tmp/error.log;

http {
  access_log /tmp/access.log;
  
  #nginx temporary files
  client_body_temp_path /tmp/client_temp;
  proxy_temp_path       /tmp/proxy_temp_path;
  fastcgi_temp_path     /tmp/fastcgi_temp;
  uwsgi_temp_path       /tmp/uwsgi_temp;
  scgi_temp_path        /tmp/scgi_temp;


  server {
    listen 80;
    server_name my-website.com localhost;
    
    location  / {
        root /websites/my-website;
        index  home.html;
    }
  }
}
```

Above, I introduced a few more simple directives. </br>

* `user` : The user to run the web server as
* `access_log` : The path to a file where the access logs are stored where we will find HTTP request\response details
* `error_log` : The path to a file where NGINX will write error details. For example if a configuration file is invalid, or there is some problem with our server
* `pid` : Specifies the file path where the process ID (PID) of the Nginx master process is stored. Our unprivileged user `nginx` needs access to this file
* `client_body_temp_path`, `proxy_temp_path`,`fastcgi_temp_path`,`uwsgi_temp_path`,`scgi_temp_path` : These are temporary files used by NGINX. Because we are using an unprivileged user, we need to set these temporary file path locations to folders that our `nginx` user will have access to.

We can use the `nginx` command with the `-t` option to test our config before starting NGINX

Take note here that we use `sudo` to run our `nginx` command as the `nginx` user we created. </br>
Remember in our previous chapter on Linx and user security, we talked about `sudo` that that it allows us to run commands as other users. </br>

```
sudo -u nginx nginx -t -c /websites/my-website/nginx.conf
```

Then stop the previous NGINX process and start our server to test it

```
sudo nginx -s stop
sudo -u nginx nginx -c /websites/my-website/nginx.conf
```

Now if we run `top` we can see NGINX running as `nginx` user.

Create a default home page with an HTML file to serve

```
sudo touch /websites/my-website/home.html
sudo touch /websites/my-website/script.js
sudo touch /websites/my-website/style.css
```

We can go ahead and grab the content from a previous module where we learned about [Websites & Web APIs](../html/README.md)

We can copy the content of the above three files from the previous module and simply paste them into these new ones. 

If we refresh the page, we can see that we now have a web server that serves our HTML home page, and that requests two additional files, our JS and CSS files and we now have a prettier looking website. </br>

## DevOps, SRE & Platform Engineering Key takeaways

In this module we covered the following

* <b>Web server installation</b> : We learned how to install a web server such as NGINX. The installation process is something to take away here as many 
  web servers will have a similar installation strategy. You further build on your existing confidence for installing software in Linux. 

* <b>Software dependencies</b> : We learned that software like web servers have dependencies or prerequisites that need to be installed.
* <b>Configuring a web server</b>
  * Software like web servers have configuration options. This is configured by a configuration file
  * Configurations may have default values
* <b>Configuration options</b> : Configuration allows us to set the following:
  * Define each `server` section for hosting multiple websites. A web server generally can host multiple websites
  * Each `server` can listen on a port, host different content
  * A server can run as a user. This allows us to set a non-privileged user for security reasons.
