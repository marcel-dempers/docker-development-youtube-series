
# Jenkins on Amazon Kubernetes

For running Jenkins on AMAZON, start [here](./amazon-eks/readme.md)

# Jenkins on Local (Docker Windows \ Minikube \ etc)

For running Jenkins on Local Docker for Windows or Minikube <br/>
Watch the [video](https://youtu.be/eRWIJGF3Y2g)

# Setting up Jenkins Agent

After installing `kubernetes-plugin` for Jenkins
* Go to Manage Jenkins | Bottom of Page | Cloud | Kubernetes (Add kubenretes cloud)
* Fill out plugin values
    * Name: kubernetes
    * Kubernetes URL: https://kubernetes.default:443
    * Kubernetes Namespace: jenkins
    * Credentials | Add | Jenkins (Choose Kubernetes service account option & Global + Save)
    * Test Connection | Should be successful! If not, check RBAC permissions and fix it!
    * Jenkins URL: http://jenkins
    * Tunnel : jenkins:50000
    * Apply cap only on alive pods : yes!
    * Add Kubernetes Pod Template
        * Name: jenkins-slave
        * Namespace: jenkins
        * Service Account: jenkins
        * Labels: jenkins-slave (you will need to use this label on all jobs)
        * Containers | Add Template
            * Name: jnlp
            * Docker Image: aimvector/jenkins-slave
            * Command to run : <Make this blank>
            * Arguments to pass to the command: <Make this blank>
            * Allocate pseudo-TTY: yes
            * Add Volume
                * HostPath type
                * HostPath: /var/run/docker.sock
                * Mount Path: /var/run/docker.sock
        * Timeout in seconds for Jenkins connection: 300
* Save

# Test a build

To run docker commands inside a jenkins agent you will need a custom jenkins agent with docker-in-docker working.
Take a look and build the docker file in `./dockerfiles/jenkins-agent`
Push it to a registry and use it instead of above configured `* Docker Image: jenkins/jnlp-slave`
If you do not use the custom image, the below pipeline will not work because default `* Docker Image: jenkins/jnlp-slave` public image does not have docker ability.

* Add a Jenkins Pipeline

```
node('jenkins-slave') {
    
     stage('unit-tests') {
        sh(script: """
            docker run --rm alpine /bin/sh -c "echo hello world"
        """)
    }
}
```
