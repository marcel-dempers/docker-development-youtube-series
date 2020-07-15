resource "kubernetes_manifest" "secret_alertmanager_main" {
  manifest = {
    "apiVersion" = "v1"
    "data" = {}
    "kind" = "Secret"
    "metadata" = {
      "name" = "alertmanager-main"
      "namespace" = "monitoring"
    }
    "stringData" = {
      "alertmanager.yaml" = "\"global\":\n  \"resolve_timeout\": \"5m\"\n\"inhibit_rules\":\n- \"equal\":\n  - \"namespace\"\n  - \"alertname\"\n  \"source_match\":\n    \"severity\": \"critical\"\n  \"target_match_re\":\n    \"severity\": \"warning|info\"\n- \"equal\":\n  - \"namespace\"\n  - \"alertname\"\n  \"source_match\":\n    \"severity\": \"warning\"\n  \"target_match_re\":\n    \"severity\": \"info\"\n\"receivers\":\n- \"name\": \"Default\"\n- \"name\": \"Watchdog\"\n- \"name\": \"Critical\"\n\"route\":\n  \"group_by\":\n  - \"namespace\"\n  \"group_interval\": \"5m\"\n  \"group_wait\": \"30s\"\n  \"receiver\": \"Default\"\n  \"repeat_interval\": \"12h\"\n  \"routes\":\n  - \"match\":\n      \"alertname\": \"Watchdog\"\n    \"receiver\": \"Watchdog\"\n  - \"match\":\n      \"severity\": \"critical\"\n    \"receiver\": \"Critical\""
    }
    "type" = "Opaque"
  }
}
resource "kubernetes_manifest" "serviceaccount_alertmanager_main" {
  manifest = {
    "apiVersion" = "v1"
    "kind" = "ServiceAccount"
    "metadata" = {
      "name" = "alertmanager-main"
      "namespace" = "monitoring"
    }
  }
}
resource "kubernetes_manifest" "service_alertmanager_main" {
  manifest = {
    "apiVersion" = "v1"
    "kind" = "Service"
    "metadata" = {
      "labels" = {
        "alertmanager" = "main"
      }
      "name" = "alertmanager-main"
      "namespace" = "monitoring"
    }
    "spec" = {
      "ports" = [
        {
          "name" = "web"
          "port" = 9093
          "targetPort" = "web"
        },
      ]
      "selector" = {
        "alertmanager" = "main"
        "app" = "alertmanager"
      }
      "sessionAffinity" = "ClientIP"
    }
  }
}
resource "kubernetes_manifest" "servicemonitor_alertmanager" {
  manifest = {
    "apiVersion" = "monitoring.coreos.com/v1"
    "kind" = "ServiceMonitor"
    "metadata" = {
      "labels" = {
        "k8s-app" = "alertmanager"
      }
      "name" = "alertmanager"
      "namespace" = "monitoring"
    }
    "spec" = {
      "endpoints" = [
        {
          "interval" = "30s"
          "port" = "web"
        },
      ]
      "selector" = {
        "matchLabels" = {
          "alertmanager" = "main"
        }
      }
    }
  }
}
resource "kubernetes_manifest" "alertmanager_main" {
  manifest = {
    "apiVersion" = "monitoring.coreos.com/v1"
    "kind" = "Alertmanager"
    "metadata" = {
      "labels" = {
        "alertmanager" = "main"
      }
      "name" = "main"
      "namespace" = "monitoring"
    }
    "spec" = {
      "image" = "quay.io/prometheus/alertmanager:v0.21.0"
      "nodeSelector" = {
        "kubernetes.io/os" = "linux"
      }
      "replicas" = 3
      "securityContext" = {
        "fsGroup" = 2000
        "runAsNonRoot" = true
        "runAsUser" = 1000
      }
      "serviceAccountName" = "alertmanager-main"
      "version" = "v0.21.0"
    }
  }
}
