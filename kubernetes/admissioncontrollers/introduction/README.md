# Introduction to Admission controllers

[Admission Webhook](https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/#what-are-admission-webhooks)

<hr/>

## Installation (local)
<hr/>

Create a kind cluster:

```
kind create cluster --name webhook --image kindest/node:v1.20.2
```

## TLS certificate notes for Webhook

<hr/>

In order for our webhook to be invoked by Kubernetes, we need a TLS certificate.<br/>
In this demo I'll be using a self signed cert. <br/>
It's ok for development, but for production I would recommend using a real certificate instead. <br/>

We'll use a very handy CloudFlare SSL tool in a docker container to get this done. <br/>

Follow [Use CFSSL to generate certificates](./tls/ssl_generate_self_signed.md)

After the above, we should have: <br/>
* a Webhook YAML file
* CA Bundle for signing new TLS certificates
* a TLS certificate (Kubernetes secret)
<br/>

## Local Development

<hr/>

We always start with a `dockerfile` since we need a Go dev environment.

```
FROM golang:1.15-alpine as dev-env

WORKDIR /app

```

Build and run the controller

```
# get dev environment: webhook

cd sourcecode
docker build . -t webhook
docker run -it --rm -p 80:80 -v ${PWD}:/app webhook sh

```

We always start with Hello world! <br/>
Let's define our basic main module and a web server

```
go mod init example-webhook
```

New file : `main.go`
```
package main

import (
  "net/http"
	"log"
)

func main() {
  http.HandleFunc("/", HandleRoot)
	http.HandleFunc("/mutate", HandleMutate)
  log.Fatal(http.ListenAndServe(":80", nil))
}

func HandleRoot(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("HandleRoot!"))
}

func HandleMutate(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("HandleMutate!"))
}

```

Build our code and run it

```
export CGO_ENABLED=0
go build -o webhook
./webhook
```

We'll be able to hit the `http://localhost/mutate` endpoint in the browser <br/>

NOTE: In Windows, container networking is not fully supported. Our container exposes port 80, but to access our Kubernetes cluster which runs in another container, we need to enable `--net host` flag. This means exposing port 80 will stop working from here on <br/>

Let's exit the container and start with `--net host` so our container can access our kubernetes `kind` cluster 

```
docker run -it --rm --net host -v ${HOME}/.kube/:/root/.kube/ -v ${PWD}:/app webhook sh
```

We can also test our access to our kubernetes cluster with the config that is mounted in:

```
apk add --no-cache curl
curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl
chmod +x ./kubectl
mv ./kubectl /usr/local/bin/kubectl
```

# Kubernetes 

How do we interact with Kubernetes ? </br>
Kubernetes provides many libraries and we'll interact with some of these today

Since we'll receive webhook events from Kubernetes, we'll need to translate these
requests into objects or structs that we understand.

For this, the serializer is important:

```
"k8s.io/apimachinery/pkg/runtime"
"k8s.io/apimachinery/pkg/runtime/serializer"

var (
	universalDeserializer = serializer.NewCodecFactory(runtime.NewScheme()).UniversalDeserializer()
)

```

To access Kubernetes, we need to define a config and a client using our config. <br/>
We can authenticate with K8s in a number of ways. <br/>

First way is good for local development and thats using a kubeconfig file. <br/>
For production, we'll use a Kubernetes service account with RBAC permissions. <br/>
We'll do both methods today. <br/>

```
# define our config and client
var config *rest.Config
var clientSet *kubernetes.Clientset

# in main()

useKubeConfig := os.Getenv("USE_KUBECONFIG")
kubeConfigFilePath := os.Getenv("KUBECONFIG")

if len(useKubeConfig) == 0 {
		// default to service account in cluster token
		c, err := rest.InClusterConfig()
		if err != nil {
			panic(err.Error())
		}
		config = c
	} else {
		//load from a kube config
		var kubeconfig string

		if kubeConfigFilePath == "" {
			if home := homedir.HomeDir(); home != "" {
				kubeconfig = filepath.Join(home, ".kube", "config")
			} 
		} else {
			kubeconfig = kubeConfigFilePath
		}

    fmt.Println("kubeconfig: " + kubeconfig)

		c, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			panic(err.Error())
		}
		config = c
	}

```

Once we built our kubeconfig, we can instantiate a client to use in our app: 

```
	cs, err := kubernetes.NewForConfig(config)
  if err != nil {
    panic(err.Error())
  }
  clientSet = cs
```

And we'll need to import the dependencies for this:

```
	"os"
	"fmt"
	"path/filepath"
	"k8s.io/client-go/kubernetes"
	rest "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
```

Since we're also using the client-go library, we need to install the same version as 
the other libraries as we can see in the `go.mod` file, we're using `v0.21.0`

```
go get k8s.io/client-go@v0.21.0
```

Rebuild to ensure no errors:

```
go build -o webhook
```

Test with a kubeconfig

```
export USE_KUBECONFIG=true
./webhook
```

To test our access, let's create a `test.go` and return pods from the kube-system namespace

```
#test.go
package main 

import ()

func test(){
}
```

Use our global clientset defined in main() and get all pods 

```
	pods, err := clientSet.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

```

Define dependencies:

```
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"fmt"
```

And finally invoking it in main() calling `test()` <br/>

Run and test our Kubernetes access:

```
bash-5.0# ./webhook
kubeconfig: /root/.kube/config
There are 11 pods in the cluster
```

## Mutating Webhook

Now that we have a working app that can talk to Kubernetes, lets implement our webhook endpoint and deploy it to kubernetes to see what type of message the API server sends us when events happen

Firstly, we need to enable a TLS endpoint </br>
Let's take some parameters where we can set the path to the TLS certificate and port number to run on. </br>

Import flag dependency:

```
"flag"
"strconv"
```

Define our parameters for cert configuration

```
type ServerParameters struct {
	port           int    // webhook server port
	certFile       string // path to the x509 certificate for https
	keyFile        string // path to the x509 private key matching `CertFile`
}

var parameters ServerParameters

# in main()

  flag.IntVar(&parameters.port, "port", 8443, "Webhook server port.")
  flag.StringVar(&parameters.certFile, "tlsCertFile", "/etc/webhook/certs/tls.crt", "File containing the x509 Certificate for HTTPS.")
  flag.StringVar(&parameters.keyFile, "tlsKeyFile", "/etc/webhook/certs/tls.key", "File containing the x509 private key to --tlsCertFile.")
  flag.Parse()

# start our web server exposing TLS endpoint 

	log.Fatal(http.ListenAndServeTLS(":" + strconv.Itoa(parameters.port), parameters.certFile, parameters.keyFile, nil))

```

Let's capture the request coming from Kubernetes and write it to local file for analysis

```
# dependencies
"io/ioutil"

# HandleMutate
body, err := ioutil.ReadAll(r.Body)
err = ioutil.WriteFile("/tmp/request", body, 0644)
if err != nil {
  panic(err.Error())
}
```

## Deployment
<hr/>

Let's built what we have and deploy it to our kubernetes cluster
We will firstly need to add a build step to our `dockerfile` to build the code </br>
And we'll also need to create a smaller runtime layer in our `dockerfile`

Full `dockerfile` :

```
FROM golang:1.15-alpine as dev-env

WORKDIR /app

FROM dev-env as build-env
COPY go.mod /go.sum /app/
RUN go mod download

COPY . /app/

RUN CGO_ENABLED=0 go build -o /webhook

FROM alpine:3.10 as runtime

COPY --from=build-env /webhook /usr/local/bin/webhook
RUN chmod +x /usr/local/bin/webhook

ENTRYPOINT ["webhook"]

```

Let's build the container and push it to a registry:

```
docker build . -t aimvector/example-webhook:v1
docker push aimvector/example-webhook:v1
```

```

# apply generated secret
kubectl -n default apply -f ./tls/example-webhook-tls.yaml


kubectl -n default apply -f rbac.yaml
kubectl -n default apply -f deployment.yaml
kubectl -n default get pods

# ensure above pods are running first

kubectl -n default apply -f webhook.yaml

```

# Deploy a demo that needs mutation

```
kubectl -n default  apply -f ./demo-pod.yaml
```

We should now be able to see an example request from Kubernetes sitting in our `tmp/request` location. This request is called an "AdmissionReview" <br/>

Kubernetes sends us an `AdmissionReview` and expects an AdmissionResponse back. <br/>

We can copy this review locally and use it for development so we dont need to deploy to 
kubernetes constantly. For example:

```
kubectl cp example-webhook-756bcb566b-9kxjp:/tmp/request ./mock-request.json
```

So lets grab the info from the admission request, so we can do something with it


```
  //dependencies 
  "k8s.io/api/admission/v1beta1"
  "errors"

  //HandleMutate()

  var admissionReviewReq v1beta1.AdmissionReview

	if _, _, err := universalDeserializer.Decode(body, nil, &admissionReviewReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Errorf("could not deserialize request: %v", err)
	} else if admissionReviewReq.Request == nil {
		w.WriteHeader(http.StatusBadRequest)
		errors.New("malformed admission review: request is nil")
	}

	fmt.Printf("Type: %v \t Event: %v \t Name: %v \n",
		admissionReviewReq.Request.Kind,
		admissionReviewReq.Request.Operation,
		admissionReviewReq.Request.Name,
	)

```

# Mutation

Firstly we need to grab the Pod object from the admission request


```
//dependencies 
apiv1 "k8s.io/api/core/v1"

var pod apiv1.Pod

err = json.Unmarshal(admissionReviewReq.Request.Object.Raw, &pod)

if err != nil {
  fmt.Errorf("could not unmarshal pod on admission request: %v", err)
}

```

To perform a simple mutation on the object before the Kubernetes API sees the object, we can apply a patch to the operation.

```
//global

type patchOperation struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value,omitempty"`
}

//HandleMutate()
var patches []patchOperation
```

Add a label that we can inject on the pod
We have to craft the kubernetes object we want to patch. <br/>
For example, a label is part of the Metadata API on the Pod spec

https://pkg.go.dev/k8s.io/api/core/v1#Pod

```
// Get existing Metadata labels

labels := pod.ObjectMeta.Labels
labels["example-webhook"] = "it-worked"

patches = append(patches, patchOperation{
				Op:    "add",
				Path:  "/metadata/labels",
				Value: labels,
		})
```

Once you have completed all your patching, convert the patches to byte slice:

```
	patchBytes, err := json.Marshal(patches)
	if err != nil {
		fmt.Errorf("could not marshal JSON patch: %v", err)
	}
```

Add it to the admission response

```
  admissionReviewResponse := v1beta1.AdmissionReview{
      Response: &v1beta1.AdmissionResponse{
        UID: admissionReviewReq.Request.UID,
        Allowed: true,
      },
    }

  admissionReviewResponse.Response.Patch = patchBytes

  bytes, err := json.Marshal(&admissionReviewResponse)
    if err != nil {
      fmt.Errorf("marshaling response: %v", err)
    }
  
  w.Write(bytes)

  //dependencies 
  "encoding/json"
```

# Build and push the updates

```
docker build . -t aimvector/example-webhook:v1
docker push aimvector/example-webhook:v1
```

# Delete all pods to get latest image

```
kubectl delete pods --all
```

# Redeploy our demo pod and see the mutations

```
kubectl -n default  apply -f ./demo-pod.yaml
```

See the injected label

```
kubectl get pods --show-labels
```