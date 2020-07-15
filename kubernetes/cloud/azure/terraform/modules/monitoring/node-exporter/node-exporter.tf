resource "kubernetes_manifest" "clusterrolebinding_node_exporter" {
  manifest = {
    "apiVersion" = "rbac.authorization.k8s.io/v1"
    "kind" = "ClusterRoleBinding"
    "metadata" = {
      "name" = "node-exporter"
    }
    "roleRef" = {
      "apiGroup" = "rbac.authorization.k8s.io"
      "kind" = "ClusterRole"
      "name" = "node-exporter"
    }
    "subjects" = [
      {
        "kind" = "ServiceAccount"
        "name" = "node-exporter"
        "namespace" = "monitoring"
      },
    ]
  }
}
resource "kubernetes_manifest" "clusterrole_node_exporter" {
  manifest = {
    "apiVersion" = "rbac.authorization.k8s.io/v1"
    "kind" = "ClusterRole"
    "metadata" = {
      "name" = "node-exporter"
    }
    "rules" = [
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
    ]
  }
}
resource "kubernetes_manifest" "daemonset_node_exporter" {
  manifest = {
    "apiVersion" = "apps/v1"
    "kind" = "DaemonSet"
    "metadata" = {
      "labels" = {
        "app.kubernetes.io/name" = "node-exporter"
        "app.kubernetes.io/version" = "v0.18.1"
      }
      "name" = "node-exporter"
      "namespace" = "monitoring"
    }
    "spec" = {
      "selector" = {
        "matchLabels" = {
          "app.kubernetes.io/name" = "node-exporter"
        }
      }
      "template" = {
        "metadata" = {
          "labels" = {
            "app.kubernetes.io/name" = "node-exporter"
            "app.kubernetes.io/version" = "v0.18.1"
          }
        }
        "spec" = {
          "containers" = [
            {
              "args" = [
                "--web.listen-address=127.0.0.1:9100",
                "--path.procfs=/host/proc",
                "--path.sysfs=/host/sys",
                "--path.rootfs=/host/root",
                "--no-collector.wifi",
                "--no-collector.hwmon",
                "--collector.filesystem.ignored-mount-points=^/(dev|proc|sys|var/lib/docker/.+|var/lib/kubelet/pods/.+)($|/)",
              ]
              "image" = "quay.io/prometheus/node-exporter:v0.18.1"
              "name" = "node-exporter"
              "resources" = {
                "limits" = {
                  "cpu" = "250m"
                  "memory" = "180Mi"
                }
                "requests" = {
                  "cpu" = "102m"
                  "memory" = "180Mi"
                }
              }
              "volumeMounts" = [
                {
                  "mountPath" = "/host/proc"
                  "name" = "proc"
                  "readOnly" = false
                },
                {
                  "mountPath" = "/host/sys"
                  "name" = "sys"
                  "readOnly" = false
                },
                {
                  "mountPath" = "/host/root"
                  "mountPropagation" = "HostToContainer"
                  "name" = "root"
                  "readOnly" = true
                },
              ]
            },
            {
              "args" = [
                "--logtostderr",
                "--secure-listen-address=[$(IP)]:9100",
                "--tls-cipher-suites=TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,TLS_RSA_WITH_AES_128_CBC_SHA256,TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256,TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256",
                "--upstream=http://127.0.0.1:9100/",
              ]
              "env" = [
                {
                  "name" = "IP"
                  "valueFrom" = {
                    "fieldRef" = {
                      "fieldPath" = "status.podIP"
                    }
                  }
                },
              ]
              "image" = "quay.io/coreos/kube-rbac-proxy:v0.4.1"
              "name" = "kube-rbac-proxy"
              "ports" = [
                {
                  "containerPort" = 9100
                  "hostPort" = 9100
                  "name" = "https"
                },
              ]
              "resources" = {
                "limits" = {
                  "cpu" = "20m"
                  "memory" = "40Mi"
                }
                "requests" = {
                  "cpu" = "10m"
                  "memory" = "20Mi"
                }
              }
            },
          ]
          "hostNetwork" = true
          "hostPID" = true
          "nodeSelector" = {
            "kubernetes.io/os" = "linux"
          }
          "securityContext" = {
            "runAsNonRoot" = true
            "runAsUser" = 65534
          }
          "serviceAccountName" = "node-exporter"
          "tolerations" = [
            {
              "operator" = "Exists"
            },
          ]
          "volumes" = [
            {
              "hostPath" = {
                "path" = "/proc"
              }
              "name" = "proc"
            },
            {
              "hostPath" = {
                "path" = "/sys"
              }
              "name" = "sys"
            },
            {
              "hostPath" = {
                "path" = "/"
              }
              "name" = "root"
            },
          ]
        }
      }
    }
  }
}
resource "kubernetes_manifest" "serviceaccount_node_exporter" {
  manifest = {
    "apiVersion" = "v1"
    "kind" = "ServiceAccount"
    "metadata" = {
      "name" = "node-exporter"
      "namespace" = "monitoring"
    }
  }
}
resource "kubernetes_manifest" "servicemonitor_node_exporter" {
  manifest = {
    "apiVersion" = "monitoring.coreos.com/v1"
    "kind" = "ServiceMonitor"
    "metadata" = {
      "labels" = {
        "app.kubernetes.io/name" = "node-exporter"
        "app.kubernetes.io/version" = "v0.18.1"
        "k8s-app" = "node-exporter"
      }
      "name" = "node-exporter"
      "namespace" = "monitoring"
    }
    "spec" = {
      "endpoints" = [
        {
          "bearerTokenFile" = "/var/run/secrets/kubernetes.io/serviceaccount/token"
          "interval" = "15s"
          "port" = "https"
          "relabelings" = [
            {
              "action" = "replace"
              "regex" = "(.*)"
              "replacement" = "$1"
              "sourceLabels" = [
                "__meta_kubernetes_pod_node_name",
              ]
              "targetLabel" = "instance"
            },
          ]
          "scheme" = "https"
          "tlsConfig" = {
            "insecureSkipVerify" = true
          }
        },
      ]
      "jobLabel" = "app.kubernetes.io/name"
      "selector" = {
        "matchLabels" = {
          "app.kubernetes.io/name" = "node-exporter"
        }
      }
    }
  }
}
resource "kubernetes_manifest" "service_node_exporter" {
  manifest = {
    "apiVersion" = "v1"
    "kind" = "Service"
    "metadata" = {
      "labels" = {
        "app.kubernetes.io/name" = "node-exporter"
        "app.kubernetes.io/version" = "v0.18.1"
      }
      "name" = "node-exporter"
      "namespace" = "monitoring"
    }
    "spec" = {
      "clusterIP" = "None"
      "ports" = [
        {
          "name" = "https"
          "port" = 9100
          "targetPort" = "https"
        },
      ]
      "selector" = {
        "app.kubernetes.io/name" = "node-exporter"
      }
    }
  }
}
