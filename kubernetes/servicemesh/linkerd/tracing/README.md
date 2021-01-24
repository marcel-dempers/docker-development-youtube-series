# Introduction to Linkerd : Distributed Tracing with Jaeger

Deploy the Jaeger All-in-One image for demo purposes

 kubectl apply -f .\kubernetes\servicemesh\linkerd\tracing\jaeger-all-in-one.yaml


 Enable tracing :

 linkerd upgrade config --addon-config kubernetes/servicemesh/linkerd/tracing/config.yaml | kubectl apply -f -
