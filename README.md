# kubernetes-scaleio-prom

A Kubernetes specific way for getting ScaleIO metrics for Prometheus

## Deployment Considerations

It is assumed that this kubernetes-scaleio-prom container will be deployed to the kube-system namespace for security reasons. It is also assumed that the kube-system namespace is capable of provisioning ScaleIO volumes because you would want to provision Prometheus with persistent volumes for the time-series data. For more information on how to do that, take a look at the following repo: https://github.com/dvonthenen/jop-stack

### Deploying Prometheus
<pre>
# Save the configuration file in Kubernetes for use by Prometheus.
# NOTE: This config.yaml is different from the one found at https://github.com/dvonthenen/jop-stack
# NOTE: This is due to adding in the kubernetes-scaleio-prom metrics endpoint
cd configs
kubectl create configmap prometheus --from-file=config.yaml --namespace=kube-system
cd ..

# We need to open up the Prometheus port to access the UI.
# If you are behind a firewall whether it's on-prem or in your
# favorite cloud like GCE, don't forget to open up the NodePort
# that is allocated!
cd services
kubectl create -f prometheus.yaml
cd ..

# Let's deploy Prometheus
# prometheus-scratch.yaml is for non-persistent deployment of Prometheus
# To deploy Prometheus with persistent storage, visit: https://github.com/dvonthenen/jop-stack
cd deployments
kubectl create -f prometheus-scratch.yaml
cd ..
</pre>

## Deploying kubernetes-scaleio-prom

<pre>
# We need to open up the kubernetes-scaleio-prom port to provide access to the endpoint.
cd services
kubectl create -f kubernetes-scaleio-prom.yaml
cd ..

# Let's deploy kubernetes-scaleio-prom
# NOTE: Before you deploy, open the kubernetes-scaleio-prom.yaml and replace the
#       SCALEIO_ENDPOINT and CLUSTER_NAME and your Kubernetes' ScaleIO Secret name
# NOTE: If you aren't using the Kubernetes' native driver for ScaleIO, you will
#       need to provide values for SCALEIO_USERNAME and SCALEIO_PASSWORD
cd deployments
kubectl create -f kubernetes-scaleio-prom.yaml
cd ..
</pre>
