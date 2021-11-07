k8s setup
deb 10

# network
cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
EOF
sudo sysctl --system

 modprobe br_netfilter
 
# docker
apt-get update
apt-get -y install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    software-properties-common
	
 curl -fsSL https://download.docker.com/linux/debian/gpg | sudo apt-key add -
 
 add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/debian \
   $(lsb_release -cs) \
   stable"
   
 apt-get update
   
 apt-get install -y --no-install-recommends docker-ce
   
# kubeadm kubectl

sudo apt-get update && sudo apt-get install -y apt-transport-https curl
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
cat <<EOF | sudo tee /etc/apt/sources.list.d/kubernetes.list
deb https://apt.kubernetes.io/ kubernetes-xenial main
EOF
sudo apt-get update
sudo apt-get install -y kubelet kubeadm kubectl
sudo apt-mark hold kubelet kubeadm kubectl
	
# master node

kubeadm init

mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config

# network
kubectl apply -f https://docs.projectcalico.org/v3.14/manifests/calico.yaml


# pods on master node
kubectl taint nodes --all node-role.kubernetes.io/master-


# nginx-ingress bare metal
https://kubernetes.github.io/ingress-nginx/deploy/
host network nicht vergessen





