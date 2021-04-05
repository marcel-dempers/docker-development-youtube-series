package main

import (
		"io/ioutil"
    "fmt"
		"encoding/json"
		"errors"
		"context"
		"flag"
		"time"
		"path/filepath"
		"os"
		"net/http"
		"strconv"
		log "github.com/sirupsen/logrus"
		"k8s.io/apimachinery/pkg/runtime"
		"k8s.io/apimachinery/pkg/runtime/serializer"
		"k8s.io/api/admission/v1beta1"
	  "k8s.io/client-go/kubernetes"
		rest "k8s.io/client-go/rest"
		"k8s.io/client-go/tools/clientcmd"
		"k8s.io/client-go/util/homedir"
)

var (
	universalDeserializer = serializer.NewCodecFactory(runtime.NewScheme()).UniversalDeserializer()
)

var config *rest.Config
var clientSet *kubernetes.Clientset

type ServerParameters struct {
	port           int    // webhook server port
	certFile       string // path to the x509 certificate for https
	keyFile        string // path to the x509 private key matching `CertFile`
}

func HandleMutate(w http.ResponseWriter, r *http.Request){

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Errorf("invalid method %s, only POST requests are allowed", r.Method)
	}

	body, err := ioutil.ReadAll(r.Body)
	//fmt.Println(string(body))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Errorf("could not read request body: %v", err)
	}

	if contentType := r.Header.Get("Content-Type"); contentType != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Errorf("unsupported content type %s, only %s is supported", contentType, "application/json")
	}

	var admissionReviewReq v1beta1.AdmissionReview

	if _, _, err := universalDeserializer.Decode(body, nil, &admissionReviewReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Errorf("could not deserialize request: %v", err)
	} else if admissionReviewReq.Request == nil {
		w.WriteHeader(http.StatusBadRequest)
		errors.New("malformed admission review: request is nil")
	}

	admissionReviewResponse, err := Mutate(admissionReviewReq)
	
	// Return the AdmissionReview with a response as JSON.
	bytes, err := json.Marshal(&admissionReviewResponse)
	if err != nil {
		fmt.Errorf("marshaling response: %v", err)
	}

	w.Write(bytes)
	
}

func HandleRoot(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("ok"))
}

var parameters ServerParameters

func main() {
	
		log.SetFormatter(&log.JSONFormatter{})
		useKubeConfig := os.Getenv("USE_KUBECONFIG")
		kubeConfigFilePath := os.Getenv("KUBECONFIG")
		runTestCode := os.Getenv("RUN_TESTCODE")

		ctx, cancel := context.WithTimeout(context.Background(), 600*time.Second)
		defer cancel()

		flag.IntVar(&parameters.port, "port", 8443, "Webhook server port.")
		flag.StringVar(&parameters.certFile, "tlsCertFile", "/etc/webhook/certs/tls.crt", "File containing the x509 Certificate for HTTPS.")
		flag.StringVar(&parameters.keyFile, "tlsKeyFile", "/etc/webhook/certs/tls.key", "File containing the x509 private key to --tlsCertFile.")
		flag.Parse()

		http.HandleFunc("/", HandleRoot)
		http.HandleFunc("/mutate", HandleMutate)

		if len(useKubeConfig) == 0 {
			// default to service account in cluster token
			c, err := rest.InClusterConfig()
			if err != nil {
				panic(err.Error())
			}
			config = c
		} else {
		  //load from a kube config
			var kubeconfig *string

			if kubeConfigFilePath == "" {
				if home := homedir.HomeDir(); home != "" {
					kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
				} else {
					kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
				}
			} else {
			  fmt.Println("test: " + kubeConfigFilePath)
				kubeconfig = flag.String("kubeconfig", kubeConfigFilePath, "")
			}

			flag.Parse()

			// use the current context in kubeconfig
			c, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
			if err != nil {
				panic(err.Error())
			}
			config = c

		}
		
		cs, err := kubernetes.NewForConfig(config)
		if err != nil {
			panic(err.Error())
		}
		clientSet = cs

		//testing area
		if len(runTestCode) != 0 {
		}
		
		err = http.ListenAndServeTLS(":" + strconv.Itoa(parameters.port), parameters.certFile, parameters.keyFile, nil)
    if err != nil {
        panic(err.Error())
    }

    fmt.Println("Starting spot scheduler")
 
}