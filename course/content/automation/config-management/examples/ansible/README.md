# 🎬 Configuration Management Example: Automate our Infrastructure (Ansible)

## 💡 Preface

This module is part of a course on DevOps. </br>
Checkout the [course introduction](../../../../../README.md) for more information </br>
This module is part of [chapter 5](../../../../../chapters/chapter-5-ci-cd/README.md)

This module draws from my extensive experience in Configuration management, Infrastructure as Code & Automation. </br> When I started as a Software Engineer, I had to look after the server infrastructure where I had to deploy the code I wrote, and make sure these servers were always up and running. </br>

When compiling code, packaging it for deployment, copying it to servers, extracting and setting up the environment - I very quickly faced the pitfalls of doing things manually and got excited when I realised I could script and automate all these steps. </br>
Ever since, I have been deeply interested in automation. </br>

Configuration Management is all about "setting up servers" and ensuring they are all configured the same. If we setup a new server, we want to ensure that server matches the others and the setup process is repeatable. 

When creating our Web server in Chapter 4, we noticed we had to install many dependencies with `apt-get install`, and make a lot of directory and file changes. </br>
Configuration management tools aim to automate these tasks. <br>

## Challenges with Idempotence

In our guides on scripting, we experienced challenges with ensuring a desired outcome. We learned that scripts do not guarantee idempotence. Scripts change the state of a system on every line that executes, and can also fail at any line, so there is a lot of complexity in error handling and retry logic to ensure outcomes are deterministic. </br>

## Value in Declarative automation

Previously, we've taken a look at desired state using Vagrant to provision our server as well as concepts such as declarative patterns. Instead of writing a script which focuses on the "How", we write our desired state which focuses on the "What". We state exactly "what" we want, instead of "how" to do it. Vagrant then makes it happen. </br>

## Configuration Management

Vagrant and other infrastructure as code tools mostly focus on infrastructure, like servers. </br>
But what if our server is already running ? What about the software, configuration, dependencies and files required **on the server** ? </br>

This is where configuration management tools come in. Once we have a server, we use a configuration management tool to provision all the software we need on that server. </br>
These tools can be declarative and they focus on solving for idempotence. </br>
They're also repeatable, so we can run them more than once and they aim to guarantee a repeatable consistent desired outcome. </br>

## Introduction to Ansible

Ansible is a very popular Configuration management tool </br>
Let's walk through some of the [Official Documentation](https://docs.ansible.com/ansible/latest/getting_started/introduction.html)

### Run our Server

In this module, we'll continue where we left off in the previous module, however we will use a new Vagrant file so each module's work is isolated. </br>
This way each course module is isolated and we can start each module with clean progression. </br>

1. Let's review the [Vagrant file](.test/Vagrantfile) contents

2. Open a Terminal and `cd` to the `Vagrantfile` folder, perform `ls` so you can ensure you are in the correct folder and can see the `Vagrantfile`

3. Configure our settings (Full variable list) </br>
   We can create a script that sets our variables which will be more convenient than pasting it every time we want to run our server. </br>
   <i>**NOTE:** i've set this script to be ignored by GIT (using `.gitignore`) so that it cannot be checked in to source control because it contains sensitive variables</i>
  
For Windows👉🏽 create a file by running `New-Item -ItemType File -Path "set_env.ps1"`
Paste the following in the file, edit it, and save it.
```powershell
$ENV:VAGRANT_HOME = "C:\temp\vms\vagrant"
$ENV:VAGRANT_DOTFILE_PATH = "C:\temp\vms\vagrant\.vagrant"
$ENV:SERVER_SHAREDFOLDER = "C:\gitrepos"
$ENV:SERVER_USERNAME = "devopsguy"
$ENV:SERVER_PASSWORD = "devopsguy"
$ENV:SERVER_NAME = "my-website-02"

# ci settings
$ENV:PAT_TOKEN = ""
$ENV:GITHUB_ORG = "<your-github-username>"
$ENV:GITHUB_REPO = "my-website"
$ENV:RUNNER_NAME = "test-runner-01"
```

For Linux👉🏽 create a file by running `touch set_env.sh`
Paste the following in the file, edit it, and save it.

```bash
export VAGRANT_HOME="~/vagrant"
export VAGRANT_DOTFILE_PATH="~/vagrant/.vagrant"
export SERVER_SHAREDFOLDER=~/gitrepos
export SERVER_USERNAME="devopsguy"
export SERVER_PASSWORD="devopsguy"
export SERVER_NAME="my-website-02"

# ci settings
export PAT_TOKEN=""
export GITHUB_ORG="<your-github-username>"
export GITHUB_REPO="my-website"
export RUNNER_NAME="test-runner-01"

```

4. Let's start from a clean slate. So we'll destroy our server, and let's comment out some steps in our [Vagrantfile](./.test/Vagrantfile) so we start with a clean server </br>
  Let's comment out all the "shell" provisioners, but keep the "file" ones that copy important files to the server.

5. Start our server from scratch  
```
#linux
. ./set_env.sh
#windows 
. .\set_env.ps1

vagrant destroy
vagrant up --provision
```

6. SSH terminal access to our new server

```
ssh devopsguy@localhost -p 2222
```

### Navigating the Documentation

I always start with the [documentation](https://docs.ansible.com/) </br>

Four main steps I want to cover in this module are:
* Introduction (high level)
* Installation 
* Inventory (your servers)
* Playbooks (the automation)

We'll walk through the Introduction using the official documentation and talk through a couple of important points which also touches on some of the topics we've pointed out in previous modules. </br>

### Installing Ansible

Now according to the documentation, we could install Ansible with:
```
pip install ansible
```

It's important to know that `pip` is not found on our Linux Server. </br>
`pip` is a package manager for Python. Just like `apt` is the package manager for Ubuntu Linux. </br>

So to install `pip`, we would need to install `python`. </br>
This would be done through command: `sudo apt install python3-pip`. </br>

Now as our server is set up to be provisioned by code, we will need to add a provisioner for Ansible, so we would not install it manually. 

There is an [Ansible Provisioner](https://developer.hashicorp.com/vagrant/docs/provisioning/ansible) for Vagrant.

Our example starting point:
```vagrantfile
Vagrant.configure("2") do |config|

  #
  # Run Ansible from the Vagrant Host
  #
  config.vm.provision "file", source: "playbook.yaml", destination: "/tmp/playbook.yaml"
  config.vm.provision "ansible_local" do |ansible|
    ansible.playbook = "playbook.yaml"
    ansible.provisioning_path = "/tmp"
    ansible.extra_vars = {
      SERVER_USERNAME: "#{SERVER_USERNAME}",
      SERVER_PASSWORD: "#{SERVER_PASSWORD}",
      PAT_TOKEN: "#{PAT_TOKEN}",
      GITHUB_ORG: "#{GITHUB_ORG}",
      GITHUB_REPO: "#{GITHUB_REPO}",
      RUNNER_NAME: "#{RUNNER_NAME}"
    }
  end

end

```

### Ansible Playbooks

You may have noticed that there is a field in the `Vagrantfile` called `ansible.playbook`.

Firstly, notice there is no scripting. No `sudo apt install -y python3-pip && sudo apt install -y ansible` commands. There are no `if` statements and control flows like scripts, which means it's easy to read and simple. </br>

In this case, Vagrant will install Ansible in a declarative manner </br>
If Ansible is already installed, it will not re-install it, so its more idempotent as well </br>
It will also run an Ansible Playbook. </br>

Ansible Playbooks are sets of plays that can run on our servers </br>
A play may be something like `Provision Web server`, or `Provision CI` </br>
A play may have a list of tasks to accomplish its goal. </br>
A task could be something like "Install packages XYZ", or "Configure Web Server" or "Copy Files from A to B" </br>
A task may execute modules like `ansible.builtin.apt` to install packages. </br>
Or a task may execute modules like `ansible.builtin.copy` to copy files </br>

For example of the Declarative structure:

```
- [PLAY = PROVISION WEB SERVER] - GOAL 
  TASKS
    - [INSTALL PACKAGES XYZ]
      - execute module - "ansible.builtin.apt"
    - [COPY FILES FROM A to B]
      - execute module - "ansible.builtin.copy"
```

Modules generally follow desired state and idempotency: </br>
```
"Most Ansible modules check whether the desired final state has already been achieved and exit without performing any actions if that state has been achieved. Repeating the task does not change the final state. Modules that behave this way are ‘idempotent’. Whether you run a playbook once or multiple times, the outcome should be the same. However, not all playbooks and not all modules behave this way. If you are unsure, test your playbooks in a sandbox environment before running them multiple times in production."
```

We can read about playbooks in the official documentation and discuss how it works. </br>
Playbooks are written in YAML format </br>

1. Create our first playbook. I've created our **empty** [playbook.yaml](./.test/playbook.yaml) which we will walk-through in steps. You can reference the [playbook-final.yaml](./.test/playbook-final.yaml) if you need to compare your progress with the final outcome. </br>
  The first step is giving it a `name`, `hosts` and `remote_user` and we can walk through some of the playbook keywords.

Taking a look at our vagrant file, the first thing we provision are user accounts. That's pretty simple so let's start with that. 

```yaml
- name: Provision User accounts
  hosts: all
  remote_user: root
```

2. Add our list of tasks. </br>
   The tasks below are just a list of actions we want to take. This gives us a moment to take the time to workout what steps we actually need. We can take a look at our [provision_user.sh](./.test/provision_user.sh)

```yaml
- name: Provision User accounts
  hosts: all
  remote_user: root
  tasks:
  - name: create admin user account
  - name: set admin user password
  - name: add admin user to sudo group
```

3. Add variables for our inputs. </br>
   Ansible allows us to inject inputs using environment variables. If we look at our user provisioning script, it accepts a username and password, so let's get that set.

```yaml
- name: Provision User accounts
  hosts: all
  remote_user: root
  vars:
    SERVER_USERNAME: "{{ lookup('env', 'SERVER_USERNAME') }}"
    SERVER_PASSWORD: "{{ lookup('env', 'SERVER_PASSWORD') }}"
  tasks:
  - name: create admin user account
  - name: set admin user password
  - name: add admin user to sudo group

```
  
4. Add modules to our tasks. </br>
   Tasks use modules to achieve their goals. In order to have our task execute a module, we will need to find an appropriate module from the [Ansible collections](https://docs.ansible.com/ansible/latest/collections/index.html#list-of-collections) page.
   There is a module for managing users that we can use in our task.

```yaml
- name: Provision User accounts
  hosts: all
  remote_user: root
  vars:
    SERVER_USERNAME: "{{ lookup('env', 'SERVER_USERNAME') }}"
    SERVER_PASSWORD: "{{ lookup('env', 'SERVER_PASSWORD') }}"
  tasks:
  - name: create admin user account
    become: true
    ansible.builtin.user:
      name: "{{ SERVER_USERNAME }}"
      shell: /bin/bash
```

The above will create a user account if it does not already exist.

5. We can use the `ansible.builtin.shell` module to set the admin accounts password 

```yaml
- name: set admin user password
  become: true
  ansible.builtin.shell: echo "{{ SERVER_USERNAME }}:{{ SERVER_PASSWORD }}" | sudo chpasswd
```

6. Finally, we can add the admin user to the `sudo` group

```yaml
- name: add admin user to sudo group
  become: true
  ansible.builtin.user:
    name: "{{ SERVER_USERNAME }}"
    groups: sudo
    append: yes
```

## Migrating our Script Provisioners to Ansible playbook

Now you should get a clear idea of the benefits of Ansible and the power that the declarative configuration has over an imperative scripting approach. We can now go ahead and add all the other provisioners as Ansible plays to our playbook. </br>

### Create a web server play

Let's start with a new Play in our existing Playbook. </br>
We'll give it a name, a task list, and variables for known bits that we will be using repeatedly. </br>

```yaml
- name: Provision Web server
  hosts: all 
  remote_user: root
  vars:
    WEBSITE_PATH: /websites/my-website
  tasks:
  - name: Install NGINX dependencies
  - name: Create NGINX group for user
  - name: Create NGINX user
  - name: Add NGINX repository get signing key
  - name: Add NGINX repository process signing key
  - name: Add NGINX repository
  - name: Add NGINX repository priority
  - name: Install NGINX
  - name: Configure NGINX service directory
  - name: Configure NGINX systemd override
  - name: Create website directory
  - name: Copy website content
  - name: Set website file permissions 
  - name: Start NGINX service
```

Let's fill in our tasks. </br>

1. Let's start by installing our NGINX dependencies </br>
   We use the `ansible.builtin.apt` module to do the equivalent of `apt-get install`

```yaml
- name: Install NGINX dependencies
  become: true
  ansible.builtin.apt:
    name: gnupg2
    state: present
    update_cache: yes
```

2. Create our NGINX unprivileged user and user group </br>

```yaml
- name: Create NGINX group for user
  become: true
  ansible.builtin.group:
    name: nginx
    state: present
- name: Create NGINX user
  become: true
  ansible.builtin.user:
    name: nginx
    group: nginx
    shell: /sbin/nologin
    state: present
```

3. In our script we set up an APT repository as per the NGINX installation instructions so that we can download and install the official NGINX package. <br>
  To do this in Ansible, we will use a module `ansible.builtin.get_url` to make a web request to grab the signing key, just like we did in our script. </br>
  Then we will run `gpg` to dearmor the key and use another Ansible module to add the repository.
  Finally we set the priority on the repository so Linux will install NGINX from the repo we added.

```yaml
- name: Add NGINX repository get signing key
  become: true
  ansible.builtin.get_url:
    url: https://nginx.org/keys/nginx_signing.key
    dest: /tmp/nginx_signing.key
- name: Add NGINX repository process signing key
  become: true
  ansible.builtin.shell: |
    gpg --dearmor < /tmp/nginx_signing.key | sudo tee /usr/share/keyrings/nginx-archive-keyring.gpg > /dev/null
- name: Add NGINX repository
  become: true
  ansible.builtin.apt_repository:
    repo: "deb [signed-by=/usr/share/keyrings/nginx-archive-keyring.gpg] http://nginx.org/packages/ubuntu {{ ansible_distribution_release }} nginx"
    state: present
    filename: nginx
- name: Add NGINX repository priority
  become: true
  ansible.builtin.copy:
    content: |
      Package: *
      Pin: origin nginx.org
      Pin: release o=nginx
      Pin-Priority: 900
    dest: /etc/apt/preferences.d/99nginx
    mode: '0644'
```

4. Now we can install NGINX with the `apt` module, the standard Linux way:

```yaml
- name: Install NGINX
  become: true
  ansible.builtin.apt:
    name: nginx
    state: present
    update_cache: yes
```

5. In our script the next step was to set up the web server `systemd` service so that NGINX runs as a service and is automatically started when our server starts up. </br>
We also indicate that we'd like the NGINX process to start with a custom configuration file
We can use the Ansible module `ansible.builtin.file` to manage `systemd` files. 

```yaml
- name: Configure NGINX service directory
  become: true
  ansible.builtin.file:
    path: "/etc/systemd/system/nginx.service.d"
    state: directory
    mode: '0755'
```

6. Now that the `systemd` override directory exists for our NGINX, we can create the override file that tells Linux how to run our NGINX process:

```yaml
- name: Configure NGINX systemd override
  become: true
  ansible.builtin.copy:
    content: |
      [Service]
      ExecStart=
      ExecStart=/usr/sbin/nginx -c ${CONFFILE}
      Environment="CONFFILE={{ WEBSITE_PATH }}/nginx.conf"
    dest: "/etc/systemd/system/nginx.service.d/override.conf"
    mode: '0644'
  notify:
  - Reload Systemd
  - Restart NGINX
```

Notice that we have a `notify` signal here, this tells Ansible to execute a handler once this task is done. In our case we want to reload `systemd` so that our changes are affected. </br>
We'll also need one to restart NGINX </br>
Let's go ahead and define this handler which is its own field within the play. </br>

```yaml
handlers:
- name: Reload Systemd
  become: true
  ansible.builtin.systemd:
    daemon_reload: yes
- name: Restart NGINX
  ansible.builtin.systemd:
    name: nginx
    state: restarted
```

7. Now we're ready to set up our website directory. We've used Vagrant to copy the files to our server on a temporary path. We'll create a directory for our website, move the files there and set the permissions for the web server user to be able to serve the files:

```yaml
- name: Create website directory
  become: true
  ansible.builtin.file:
    path: "{{ WEBSITE_PATH }}"
    state: directory
    owner: nginx
    group: nginx
    mode: '0755'
- name: Copy website content
  become: true
  ansible.builtin.copy:
    src: /home/vagrant/tmp/my-website/
    dest: "{{ WEBSITE_PATH }}"
    owner: nginx
    group: nginx
    mode: '0644'
    remote_src: yes
- name: Set website file permissions
  become: true
  ansible.builtin.file:
    path: "{{ WEBSITE_PATH }}"
    state: directory
    owner: nginx
    group: nginx
    mode: '0755'
```

8. Last but not least, we can now start our NGINX service

```yaml
- name: Start NGINX service
  become: true
  ansible.builtin.systemd:
    name: nginx
    state: started
    enabled: yes
```

### Create a CI/CD runner play

Once you start with Ansible, and perform all these common types of Linux tasks, you will quickly start gaining momentum as you'll be re-using a lot of the same modules and patterns. </br>
We'll create user accounts, directories, files, run commands, download and install stuff. </br>

Let's start with a new Play to provision our Github Actions Runner and add our basic task list

```
- name: Provision Github CI/CD runner
  hosts: all 
  remote_user: root
  vars:
    PAT_TOKEN: "{{ lookup('env', 'PAT_TOKEN') }}"
    GITHUB_ORG: "{{ lookup('env', 'GITHUB_ORG') }}"
    GITHUB_REPO: "{{ lookup('env', 'GITHUB_REPO') }}"
    RUNNER_NAME: "{{ lookup('env', 'RUNNER_NAME') }}"
    WEBSITE_PATH: /websites/my-website
    REPO_URL: "https://github.com/{{ GITHUB_ORG }}/{{ GITHUB_REPO }}"
  tasks:
  - name: Create Github user
  - name: Install Github dependencies
  - name: Set permissions on website folder
  - name: Set permissions on website reload
  - name: Github Runner - Create runner folder
  - name: Github Runner - Get latest runner version
  - name: Github Runner - Set runner version
  - name: Github Runner - Download
  - name: Github Runner - Extract 
  - name: Install runner dependencies
  - name: Install runner dependencies marker file
  - name: Github Runner - Get registration token
  - name: Github Runner - Install
  - name: Github Runner - Marker file
  - name: Github Service - Install
  - name: Github Service - Marker file
```

1. Create Github user
```yaml
- name: Create Github user
  become: true
  ansible.builtin.user:
    name: "github"
    shell: /bin/bash
```

2. Install Github dependencies
```yaml
- name: Install Github dependencies
  become: true
  ansible.builtin.apt:
    name: acl
    state: present
    update_cache: yes
```

3. Set permissions on website folder
```yaml
- name: Set permissions on website folder
  become: true
  ansible.builtin.shell: setfacl -m u:github:rwx {{ WEBSITE_PATH }}
```
4. Set permissions on website reload
```yaml
- name: Set permissions on website reload
  become: true
  ansible.builtin.copy:
    content: |
      # Permissions for the GitHub Actions runner user
      github ALL=(root) NOPASSWD: /usr/bin/sed, /usr/bin/grep, /usr/sbin/nginx -s reload
    dest: /etc/sudoers.d/github
    mode: '0440'
    validate: 'visudo -cf %s'
```
5. Github Runner - Create runner folder
```yaml
- name: Github Runner - Create runner folder
  become: true
  ansible.builtin.file:
    path: /home/github/{{ RUNNER_NAME }}
    state: directory
    owner: github
    group: github
    mode: '0755'
```
6. Github Runner - Get latest runner version
```yaml
- name: Github Runner - Get latest runner version
  ansible.builtin.uri:
    url: "https://api.github.com/repos/actions/runner/releases/latest"
    return_content: yes
  register: latest_release
```
7. Github Runner - Set runner version
```yaml
- name: Github Runner - Set runner version
  ansible.builtin.set_fact:
    runner_version: "{{ latest_release.json.tag_name | replace('v','') }}"
    runner_file: "actions-runner-linux-x64-{{ latest_release.json.tag_name | replace('v','') }}.tar.gz"
```
8. Github Runner - Download
```yaml

- name: Github Runner - Download
  become: true
  ansible.builtin.get_url:
    url: "https://github.com/actions/runner/releases/download/{{ latest_release.json.tag_name }}/{{ runner_file }}"
    dest: /home/github/{{ RUNNER_NAME }}
    owner: "github"
    group: "github"
    mode: '0644'
```
9. Github Runner - Extract 
```yaml
- name: Github Runner - Extract 
  become: true
  ansible.builtin.unarchive:
    src: "/home/github/{{ RUNNER_NAME }}/{{ runner_file }}"
    dest: "/home/github/{{ RUNNER_NAME }}"
    owner: "github"
    group: "github"
    remote_src: yes
```
10. Install runner dependencies
```yaml
- name: Install runner dependencies
  become: true
  ansible.builtin.command: |
    /home/github/{{ RUNNER_NAME }}/bin/installdependencies.sh
  args:
    chdir: /home/github/{{ RUNNER_NAME }}
    creates: /home/github/{{ RUNNER_NAME }}/.installed_dependencies
```
11. Install runner dependencies marker file
```yaml
- name: Install runner dependencies marker file
  become: true
  ansible.builtin.file:
    path: /home/github/{{ RUNNER_NAME }}/.installed_dependencies
    state: touch
    owner: github
    group: github
```
12. Github Runner - Get registration token
```yaml
- name: Github Runner - Get registration token
  register: token_result
  changed_when: true
  no_log: true
  ansible.builtin.uri:
    url: "https://api.github.com/repos/{{ GITHUB_ORG }}/{{ GITHUB_REPO }}/actions/runners/registration-token"
    method: POST
    headers:
      Authorization: "token {{ PAT_TOKEN }}"
      Accept: "application/vnd.github+json"
    body_format: json
    status_code: 201
```
13. Github Runner - Install
```yaml
- name: Github Runner - Install
  become: true
  become_user: github
  ansible.builtin.command: |
    ./config.sh --unattended \
      --token "{{ token_result.json.token }}" \
      --url "{{ REPO_URL }}" \
      --name "{{ RUNNER_NAME }}" \
      --labels "self-hosted" "linux" \
      --work "_work" \
      --replace \
      --runasservice \
      --runnergroup "Default"
  args:
    chdir: /home/github/{{ RUNNER_NAME }}
    creates: /home/github/{{ RUNNER_NAME }}/.installed_runner
```
14. Github Runner - Marker file
```yaml
- name: Github Runner - Marker file
  become: true
  ansible.builtin.file:
    path: /home/github/{{ RUNNER_NAME }}/.installed_runner
    state: touch
    owner: github
    group: github
```
15. Github Service - Install
```yaml
- name: Github Service - Install
  become: true
  ansible.builtin.shell: |
    ./svc.sh install github
    ./svc.sh start
  args:
    chdir: /home/github/{{ RUNNER_NAME }}
    creates: /home/github/{{ RUNNER_NAME }}/.installed_runner_service  
```
16. Github Service - Marker file
```yaml
- name: Github Service - Marker file
  become: true
  ansible.builtin.file:
    path: /home/github/{{ RUNNER_NAME }}/.installed_runner_service
    state: touch
    owner: github
    group: github
```
