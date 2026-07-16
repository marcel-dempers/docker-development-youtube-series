# 🎬 The Basics of DNS & Domains

## 💡 Preface

This module is part of a course on DevOps. </br>
Checkout the [course introduction](../../../README.md) for more information </br>
This module is part of [chapter 4](../../../chapters/chapter-4-web-and-http/README.md)

This module draws from my extensive experience in the world of Web. Building and managing Web servers, building web sites, web services, monitoring , security, networking etc.  

## What are Domains 

In our module on the overview of Web, we have looked at the HTTP request and mentioned domains and domain names a few times. </br>

An observation we made, when looking at the HTTP request headers, is that we saw the `Host` being a field in the HTTP Request Header and the value being a domain name. </br>

We also touched on the URL part that contains the domain name, I.E `https://<domain-name>/`

Take a look at the following HTTP Request example with `curl`: 

```
curl -v https://marceldempers.dev/course
GET /course HTTP/2
> Host: marceldempers.dev
> user-agent: curl/7.81.0
> accept: */*
```

The domain is part of the address we type in the browser (or `curl`) when visiting a website, such as `marceldempers.dev` </br>

It's a human-readable address used to identify a website or service on the internet. </br>

Think of it as an "alias" to an IP address. </br>
It's much easier to remember a domain name instead of an IP address when you want to visit your favourite web sites </br>

Domain names are also globally unique, so when you purchase one through a domain registrar, the registrar will check if its available for purchase.</br>

<b>When making the HTTP connection:</b>
Under the hood, the HTTP client will do a "domain name lookup" to query or resolve the IP address and will use TCP to create a connection to the destination IP address. So the Domain name is only used to calculate the IP address technically and thereafter is not used by the client in the TCP connection process. </br>

The client does however set a `Host` header value as part of the HTTP request headers as we can see in the above HTTP request. But this "Host" header is not used by the underlying TCP connection</br>
We will see below how the domain name is used to get the IP address. </br>

Domains are something you generally purchase if you want a public domain on the internet. </br>

<b>Note: </b><i>In more advanced platforms, you may even have your own "private" domains that are used by web servers that are not on the public internet and these domains are only addressable within a private network 
</i>

## What are Host Names 

A hostname is the name assigned to a specific device (e.g., a computer, server, or network device) within a network. It uniquely identifies the device within its local network. 

It's important not to confuse hostnames with domain names and host headers in HTTP </br>

Hostnames are typically used to within a local network and are often used as internal identification of servers or devices in a network. <br/>
For example, we may have a virtual server where we installed and configured a web server hosting our website and we may have a hostname for it called `webserver1`

Within the same private network, we could potentially access our site over `http://webserver1` by using its hostname instead of a domain name. </br>

Domain names are resolved to IP addresses using a technology called DNS. Hostnames are not resolved by DNS. </br>

## What is DNS

DNS stands for "Domain Name System" which is just a system (server) that resolves domain names. We pass it a domain name and it returns an IP address. </br>

In a very basic computer network like a home network, there are technically no DNS servers running. Each application on a computer asks the Operating system to resolve DNS. The operating system then relies on the network router to provide DNS information. Network routers would generally be configured to provide DNS using your internet service providers DNS server IP address. 

![diagran](dns.drawio.png) 

### DNS Query

In a typical home network, the **HTTP client itself does not directly perform the DNS lookup to an external DNS server like Google DNS (8.8.8.8).** Instead, the lookup process flows through a series of intermediaries:

**A Typical DNS flow:**

1.  **HTTP Client Initiates Request:** When you type a domain name (e.g., `marceldempers.dev`) into your web browser (the HTTP client), the client needs to know the IP address associated with that domain to connect to the server.

2.  **Client Asks the Operating System (OS):** The HTTP client doesn't have its own built-in DNS resolver. Instead, it delegates the task of resolving the domain name to the underlying **operating system (OS)**. It makes a system call, essentially asking, "Hey OS, what's the IP address for `marceldempers.dev`?"

3.  **Operating System's DNS Resolver:** The OS has its own DNS client (resolver) built-in. Before it sends out a query, it will check:
    * **Local DNS Cache:** The OS maintains a local cache of recently resolved domain names. If it finds the IP address for `marceldempers.dev` in its cache, it will immediately return it to the HTTP client, skipping further steps.
    * **Configured DNS Servers:** If the IP address isn't in its cache, the OS looks at its network configuration to find the IP addresses of the **DNS servers it should use**. In a typical home network, these DNS server addresses are usually obtained automatically from the router via DHCP.

4.  **Query Sent to the Home Router:** The OS will send the DNS query (e.g., "What's the IP for `marceldempers.dev`?") to the DNS server address it has configured. In most home networks, this address will be the **IP address of your home router** (e.g., `192.168.1.1`).

5.  **Home Router's DNS Role:** Your home router acts as a **DNS relay or proxy**. It receives the DNS query from your computer. What the router does next depends on its own configuration:
    * **Default Behavior (Often):** The router itself typically doesn't have the entire internet's DNS records. Instead, it's configured to forward DNS queries to **upstream DNS servers**. These upstream servers are often those provided by your Internet Service Provider (ISP).
    * **User-Configured DNS:**  Many users, for reasons of speed, privacy, or content filtering, will manually configure their router to use public DNS servers like Google DNS (8.8.8.8 and 8.8.4.4), Cloudflare DNS (1.1.1.1), or OpenDNS. If your router is configured this way, it will forward the query to Google DNS.

6.  **Query Sent to Upstream DNS Server (e.g., Google DNS):** The router forwards the DNS query to the configured upstream DNS server (e.g., 8.8.8.8).

7.  **Upstream DNS Server Resolves:** The upstream DNS server (e.g., Google DNS) then performs the actual recursive lookup process through the internet's DNS hierarchy (root servers, TLD servers, authoritative servers) to find the authoritative IP address for `marceldempers.dev`.

8.  **Response Back Through the Chain:** Once the upstream DNS server finds the IP address, it sends the response back:
    * To your home router.
    * The router then sends the response back to your OS.
    * Your OS then sends the IP address back to the HTTP client.

9.  **HTTP Client Connects:** Finally, with the IP address in hand (e.g., `142.250.190.132`), the HTTP client can establish a TCP connection to the web server hosting `marceldempers.dev` and proceed with the HTTP request.

## DNS commandline tools

Two very popular command line tools to troubleshoot and test DNS are `nslookup` and `dig`. 

`nslookup` is a network administration command-line tool available for many computer operating systems for querying the Domain Name System (DNS) to obtain domain name or IP address mapping or other DNS records.

**`nslookup` Basic Usage:**

To perform a basic lookup of the A record (IP address) for a domain, simply type `nslookup` followed by the domain name.

```
nslookup google.com
```

**Example output:**

```
Server:     192.168.1.1
Address:    192.168.1.1#53

Non-authoritative answer:
Name:   google.com
Address: 172.67.147.168
Name:   google.com
Address: 104.21.73.18
```

**Description of Output:**

* `Server:` and `Address`: 
  * These lines indicate the DNS server that `nslookup` used to perform the query. In this example, it's a local router or a DNS server configured on your system.

* `Non-authoritative answer`: 
  * This means the answer was provided by a caching DNS server, not directly by the authoritative DNS server for `google.com`.
* `Name: google.com`: The domain name you queried.
* `Address: [IP Address]`: 
  * The IP address(es) associated with the domain. In this case, `google.com` has multiple A records, indicating it can be reached at different IP addresses (often for load balancing or redundancy).


`dig` (Domain Information Groper) is a more advanced and flexible command-line tool for interrogating DNS name servers. It is generally preferred by network administrators for troubleshooting DNS issues.

**`dig` Basic Usage:**

To perform a basic lookup of the A record for a domain, simply type `dig` followed by the domain name.

```
dig google.com
```

**Example Partial Output:**

```
; <<>> DiG 9.16.1-Ubuntu <<>> google.com
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 36780
;; flags: qr rd ra; QUERY: 1, ANSWER: 2, AUTHORITY: 0, ADDITIONAL: 1

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 1232
;; QUESTION SECTION:
;google.com.        IN    A

;; ANSWER SECTION:
google.com.    300    IN    A    104.21.73.18
google.com.    300    IN    A    172.67.147.168

;; Query time: 1 msec
;; SERVER: 127.0.0.53#53(127.0.0.53)
;; WHEN: Fri Jun 14 20:54:51 AEST 2025
;; MSG SIZE  rcvd: 91
```

**Description of Output (Key Sections):**

* **`;; HEADER`**: Provides high-level information about the DNS query and response, including the `status: NOERROR` which indicates a successful query.
* **`;; QUESTION SECTION:`**: Shows the query that was made, in this case, for an `A` record for `google.com`.
* **`;; ANSWER SECTION:`**: This is the most important part, providing the DNS records found.
    * `google.com.` : The queried domain.
    * `300`: The Time-To-Live (TTL) in seconds, indicating how long this record can be cached by other DNS servers.
    * `IN`: Internet class.
    * `A`: The record type, indicating an IPv4 address.
    * `104.21.73.18`, `172.67.147.168`: The IP addresses associated with the domain.
* **`;; SERVER:`**: The DNS server that `dig` used to perform the query.
* **`;; WHEN:`**: The timestamp of when the query was performed.

### DNS Record Types

There are a few types of DNS records, but two very important onces to learn about in my experience. </br>

* **A Record (Address Record)** 
  * An A record directly maps a domain name to an IPv4 address. <br/>
    It's the most basic way to tell a web browser or other application which IP address corresponds to a human-readable domain name.<br/>
    So basically, for the website `marceldempers.dev` you are trying to reach, go to `104.21.73.18`. <br/>
* **CNAME Record (Canonical Name Record)**
  * A CNAME record maps one domain name (an alias) to another domain name (its canonical, or "true," name). <br/>
  It never points directly to an IP address. <br/>
  Instead, it creates an alias, telling a DNS resolver, "If you're looking for this name, go look for that other name instead." <br/>

  Benefits of using `CNAME` records. 

  * Management of Dynamic IP addresses
    * Cloud providers (AWS, Azure, Google Cloud, etc.) often assign dynamic IP addresses to resources like load balancers, CDNs, or managed services. <br/>
    If you used an A record, you'd constantly have to update your DNS records every time that IP address changed.
    So its better to map your domain name to a CNAME alias, so the IP address on the target resource can change and not affect your clients using your domain name. 
  * Reduce manual DNS updates
    * If you used a cloud load balancer like one from AWS, you would point your domain `marceldempers.dev` to the cloud provider's hostname (e.g., `your-load-balancer-id.us-east-1.elb.amazonaws.com`).  </br>
    If the IP address of that load balancer changes internally within the cloud provider's network, you don't need to touch your DNS records. The cloud provider handles the resolution of their hostname to the correct IP.

# Key Takeaways 

* **Key usages about DNS**
  * AS a DevOps engineer, in my experience, DNS plays a large role in infrastructure, cloud and the world of Web. Therefore understanding the basics in this module is crucial. DNS is a strange technology, as you will not always be directly involved in working on DNS, DNS configuration, settings, or servers, and most of the time, we expect DNS to work and just "be there" for our use. The cloud mostly takes care of this for us.
  * DNS is often the root cause of major outages and issues, and least expected. There is a meme stating "It's always DNS" which points to the irony of DNS causing outages that are the least expected.
* **Understand the basics**
  * Therefore its crucial to **know the basics**, and be **able to test**, **dedug**, **troubleshoot** DNS frequently **using the command-line**.
* **Understand DNS queries**
  * You want to make sure you have a grip on how DNS resolution works as shown in our diagram. How DNS is resolving a host name to IP and how it got there. 
