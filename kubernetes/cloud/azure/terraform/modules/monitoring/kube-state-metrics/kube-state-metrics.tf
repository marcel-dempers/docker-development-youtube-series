resource "kubernetes_manifest" "clusterrolebinding_kube_state_metrics" {
  manifest = {
    "apiVersion" = "rbac.authorization.k8s.io/v1"
    "kind" = "ClusterRoleBinding"
    "metadata" = {
      "labels" = {
        "app.kubernetes.io/name" = "kube-state-metrics"
        "app.kubernetes.io/version" = "1.9.5"
      }
      "name" = "kube-state-metrics"
    }
    "roleRef" = {
      "apiGroup" = "rbac.authorization.k8s.io"
      "kind" = "ClusterRole"
      "name" = "kube-state-metrics"
    }
    "subjects" = [
      {
        "kind" = "ServiceAccount"
        "name" = "kube-state-metrics"
        "namespace" = "monitoring"
      },
    ]
  }
}
resource "kubernetes_manifest" "clusterrole_kube_state_metrics" {
  manifest = {
    "apiVersion" = "rbac.authorization.k8s.io/v1"
    "kind" = "ClusterRole"
    "metadata" = {
      "labels" = {
        "app.kubernetes.io/name" = "kube-state-metrics"
        "app.kubernetes.io/version" = "1.9.5"
      }
      "name" = "kube-state-metrics"
    }
    "rules" = [
      {
        "apiGroups" = [
          "",
        ]
        "resources" = [
          "configmaps",
          "secrets",
          "nodes",
          "pods",
          "services",
          "resourcequotas",
          "replicationcontrollers",
          "limitranges",
          "persistentvolumeclaims",
          "persistentvolumes",
          "namespaces",
          "endpoints",
        ]
        "verbs" = [
          "list",
          "watch",
        ]
      },
      {
        "apiGroups" = [
          "extensions",
        ]
        "resources" = [
          "daemonsets",
          "deployments",
          "replicasets",
          "ingresses",
        ]
        "verbs" = [
          "list",
          "watch",
        ]
      },
      {
        "apiGroups" = [
          "apps",
        ]
        "resources" = [
          "statefulsets",
          "daemonsets",
          "deployments",
          "replicasets",
        ]
        "verbs" = [
          "list",
          "watch",
        ]
      },
      {
        "apiGroups" = [
          "batch",
        ]
        "resources" = [
          "cronjobs",
          "jobs",
        ]
        "verbs" = [
          "list",
          "watch",
        ]
      },
      {
        "apiGroups" = [
          "autoscaling",
        ]
        "resources" = [
          "horizontalpodautoscalers",
        ]
        "verbs" = [
          "list",
          "watch",
        ]
      },
      {
        "apiGroups" = [
          "authentication.k8s.io",
        ]
        "resources" = [
          "tokenreviews",
        ]
        "verbs" = [
          "create",
        ]
      },
      {
        "apiGroups" = [
          "authorization.k8s.io",
        ]
        "resources" = [
          "subjectaccessreviews",
        ]
        "verbs" = [
          "create",
        ]
      },
      {
        "apiGroups" = [
          "policy",
        ]
        "resources" = [
          "poddisruptionbudgets",
        ]
        "verbs" = [
          "list",
          "watch",
        ]
      },
      {
        "apiGroups" = [
          "certificates.k8s.io",
        ]
        "resources" = [
          "certificatesigningrequests",
        ]
        "verbs" = [
          "list",
          "watch",
        ]
      },
      {
        "apiGroups" = [
          "storage.k8s.io",
        ]
        "resources" = [
          "storageclasses",
          "volumeattachments",
        ]
        "verbs" = [
          "list",
          "watch",
        ]
      },
      {
        "apiGroups" = [
          "admissionregistration.k8s.io",
        ]
        "resources" = [
          "mutatingwebhookconfigurations",
          "validatingwebhookconfigurations",
        ]
        "verbs" = [
          "list",
          "watch",
        ]
      },
      {
        "apiGroups" = [
          "networking.k8s.io",
        ]
        "resources" = [
          "networkpolicies",
        ]
        "verbs" = [
          "list",
          "watch",
        ]
      },
      {
        "apiGroups" = [
          "coordination.k8s.io",
        ]
        "resources" = [
          "leases",
        ]
        "verbs" = [
          "list",
          "watch",
        ]
      },
    ]
  }
}
resource "kubernetes_manifest" "deployment_kube_state_metrics" {
  manifest = {
    "apiVersion" = "apps/v1"
    "kind" = "Deployment"
    "metadata" = {
      "labels" = {
        "app.kubernetes.io/name" = "kube-state-metrics"
        "app.kubernetes.io/version" = "1.9.5"
      }
      "name" = "kube-state-metrics"
      "namespace" = "monitoring"
    }
    "spec" = {
      "replicas" = 1
      "selector" = {
        "matchLabels" = {
          "app.kubernetes.io/name" = "kube-state-metrics"
        }
      }
      "template" = {
        "metadata" = {
          "labels" = {
            "app.kubernetes.io/name" = "kube-state-metrics"
            "app.kubernetes.io/version" = "1.9.5"
          }
        }
        "spec" = {
          "containers" = [
            {
              "args" = [
                "--host=127.0.0.1",
                "--port=8081",
                "--telemetry-host=127.0.0.1",
                "--telemetry-port=8082",
              ]
              "image" = "quay.io/coreos/kube-state-metrics:v1.9.5"
              "name" = "kube-state-metrics"
              "securityContext" = {
                "runAsUser" = 65534
              }
            },
            {
              "args" = [
                "--logtostderr",
                "--secure-listen-address=:8443",
                "--tls-cipher-suites=TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,TLS_RSA_WITH_AES_128_CBC_SHA256,TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256,TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256",
                "--upstream=http://127.0.0.1:8081/",
              ]
              "image" = "quay.io/coreos/kube-rbac-proxy:v0.4.1"
              "name" = "kube-rbac-proxy-main"
              "ports" = [
                {
                  "containerPort" = 8443
                  "name" = "https-main"
                },
              ]
              "securityContext" = {
                "runAsUser" = 65534
              }
            },
            {
              "args" = [
                "--logtostderr",
                "--secure-listen-address=:9443",
                "--tls-cipher-suites=TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,TLS_RSA_WITH_AES_128_CBC_SHA256,TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256,TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256",
                "--upstream=http://127.0.0.1:8082/",
              ]
              "image" = "quay.io/coreos/kube-rbac-proxy:v0.4.1"
              "name" = "kube-rbac-proxy-self"
              "ports" = [
                {
                  "containerPort" = 9443
                  "name" = "https-self"
                },
              ]
              "securityContext" = {
                "runAsUser" = 65534
              }
            },
          ]
          "nodeSelector" = {
            "kubernetes.io/os" = "linux"
          }
          "serviceAccountName" = "kube-state-metrics"
        }
      }
    }
  }
}
resource "kubernetes_manifest" "serviceaccount_kube_state_metrics" {
  manifest = {
    "apiVersion" = "v1"
    "kind" = "ServiceAccount"
    "metadata" = {
      "labels" = {
        "app.kubernetes.io/name" = "kube-state-metrics"
        "app.kubernetes.io/version" = "1.9.5"
      }
      "name" = "kube-state-metrics"
      "namespace" = "monitoring"
    }
  }
}
resource "kubernetes_manifest" "servicemonitor_kube_state_metrics" {
  manifest = {
    "apiVersion" = "monitoring.coreos.com/v1"
    "kind" = "ServiceMonitor"
    "metadata" = {
      "labels" = {
        "app.kubernetes.io/name" = "kube-state-metrics"
        "app.kubernetes.io/version" = "1.9.5"
        "k8s-app" = "kube-state-metrics"
      }
      "name" = "kube-state-metrics"
      "namespace" = "monitoring"
    }
    "spec" = {
      "endpoints" = [
        {
          "bearerTokenFile" = "/var/run/secrets/kubernetes.io/serviceaccount/token"
          "honorLabels" = true
          "interval" = "30s"
          "port" = "https-main"
          "relabelings" = [
            {
              "action" = "labeldrop"
              "regex" = "(pod|service|endpoint|namespace)"
            },
          ]
          "scheme" = "https"
          "scrapeTimeout" = "30s"
          "tlsConfig" = {
            "insecureSkipVerify" = true
          }
        },
        {
          "bearerTokenFile" = "/var/run/secrets/kubernetes.io/serviceaccount/token"
          "interval" = "30s"
          "port" = "https-self"
          "scheme" = "https"
          "tlsConfig" = {
            "insecureSkipVerify" = true
          }
        },
      ]
      "jobLabel" = "app.kubernetes.io/name"
      "selector" = {
        "matchLabels" = {
          "app.kubernetes.io/name" = "kube-state-metrics"
        }
      }
    }
  }
}
resource "kubernetes_manifest" "service_kube_state_metrics" {
  manifest = {
    "apiVersion" = "v1"
    "kind" = "Service"
    "metadata" = {
      "labels" = {
        "app.kubernetes.io/name" = "kube-state-metrics"
        "app.kubernetes.io/version" = "1.9.5"
      }
      "name" = "kube-state-metrics"
      "namespace" = "monitoring"
    }
    "spec" = {
      "clusterIP" = "None"
      "ports" = [
        {
          "name" = "https-main"
          "port" = 8443
          "targetPort" = "https-main"
        },
        {
          "name" = "https-self"
          "port" = 9443
          "targetPort" = "https-self"
        },
      ]
      "selector" = {
        "app.kubernetes.io/name" = "kube-state-metrics"
      }
    }
  }
}
