#  EXAMPLE OF USING ARGOCD
 - convenient because it's all in one docker container
 - app-of-apps approach
 - example of adding a custom application
 - prometheus monitoring included

[![Build and Push Docker Image](https://github.com/dzhunli/argocd_tests/actions/workflows/build-push.yml/badge.svg)](https://github.com/dzhunli/argocd_tests/actions/workflows/build-push.yml)


---


## Requirements
1. You must have Linux
2. Нou need to install docker if it is not already installed on your system
docker --> please visit https://docs.docker.com/engine/install/ (you don`t need docker desktop ONLY docker engine)
## Install (Preparing)
- Run k3s server in docker container 
```bash
    docker run -d --name k3s-server --privileged --restart=unless-stopped --network=host rancher/k3s:v1.26.0-k3s1 server
```

- Setup some envs 
 ```bash
    cat /etc/rancher/k3s/k3s.yaml > k3s.yaml && export KUBECONFIG=$PWD/k3s.yaml
```

- Create namespace for argocd
 ```bash
   kubectl create namespace argocd
```

-  Apply argocd masnifest in created namespace
 ```bash
   kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```

-  Wait needed condition for next steps 
 ```bash
kubectl wait --for=condition=available --timeout=300s -n argocd deployment/argocd-server
```

-  Set up NodePort for web 
 ```bash
NODE_PORT=$(kubectl get svc argocd-server -n argocd -o jsonpath='{.spec.ports[0].nodePort}')
```

-  Create admin password 
 ```bash
ARGOCD_PASSWORD=$(kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d)
```

-  Getting the port and password from the web site
 ```bash
echo -e  " Port: $NODE_PORT \n User: admin \n Pass: $ARGOCD_PASSWORD \n" 
```

**-  Go to 127.0.0.1:port in your browser and try to log in**

