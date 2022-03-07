apiVersion: v1
kind: Namespace
metadata:
  name: portainer
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: portainer-sa-clusteradmin
  namespace: portainer
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: portainer-crb-clusteradmin
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
  - kind: ServiceAccount
    name: portainer-sa-clusteradmin
    namespace: portainer
# Optional: can be added to expose the agent port 80 to associate an Edge key.
# ---
# apiVersion: v1
# kind: Service
# metadata:
#   name: portainer-agent
#   namespace: portainer
# spec:
#   type: LoadBalancer
#   selector:
#     app: portainer-agent
#   ports:
#     - name: http
#       protocol: TCP
#       port: 80
#       targetPort: 80
---
apiVersion: v1
kind: Service
metadata:
  name: portainer-agent
  namespace: portainer
spec:
  clusterIP: None
  selector:
    app: portainer-agent
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: portainer-agent
  namespace: portainer
spec:
  selector:
    matchLabels:
      app: portainer-agent
  template:
    metadata:
      labels:
        app: portainer-agent
    spec:
      serviceAccountName: portainer-sa-clusteradmin
      containers:
        - name: portainer-agent
          image: portainer/agent:2.11.1
          imagePullPolicy: Always
          env:
            - name: LOG_LEVEL
              value: INFO
            - name: KUBERNETES_POD_IP
              valueFrom:
                fieldRef:
                  fieldPath: status.podIP
            - name: EDGE
              value: "1"
            - name: AGENT_CLUSTER_ADDR
              value: "portainer-agent"
            - name: EDGE_ID
              valueFrom:
                configMapKeyRef:
                  name: portainer-agent-edge
                  key: edge.id
            - name: EDGE_INSECURE_POLL
              valueFrom:
                configMapKeyRef:
                  name: portainer-agent-edge
                  key: edge.insecure_poll
            - name: EDGE_KEY
              valueFrom:
                secretKeyRef:
                  name: portainer-agent-edge-key
                  key: edge.key
          ports:
            - containerPort: 9001
              protocol: TCP
            - containerPort: 80
              protocol: TCP
