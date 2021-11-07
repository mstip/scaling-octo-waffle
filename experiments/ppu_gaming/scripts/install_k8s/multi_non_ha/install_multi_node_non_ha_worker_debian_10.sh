#!/usr/bin/env bash
set -x
set -e

echo "master ip $1"
echo "token $2"
echo "cert $3"

# check if run as root
if [[ $EUID -ne 0 ]]; then
   echo "This script must be run as root"
   exit 1
fi

# network
cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
EOF
sysctl --system
modprobe br_netfilter
 
# install docker
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
   
# install kubeadm 
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
cat <<EOF | sudo tee /etc/apt/sources.list.d/kubernetes.list
deb https://apt.kubernetes.io/ kubernetes-xenial main
EOF
sudo apt-get update
sudo apt-get install -y kubeadm kubelet
sudo apt-mark hold kubeadm kubelet
	
# join the cluster
kubeadm join $1:6443 --token $2 --discovery-token-ca-cert-hash $3
