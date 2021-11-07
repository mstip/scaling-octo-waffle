#!/usr/bin/env bash
# usage $1 ip of the internal network $2 pod network
# ./install_multi_node_non_ha_master_debian_10.sh 192.168.0.3 10.10.0.0/8
# this script requires an interal network for the nodes
set -x
set -e

echo "internal k8s network for nodes $1"
echo "pod net $2"

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
   
# install kubeadm kubectl
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
cat <<EOF | sudo tee /etc/apt/sources.list.d/kubernetes.list
deb https://apt.kubernetes.io/ kubernetes-xenial main
EOF
sudo apt-get update
sudo apt-get install -y kubelet kubeadm kubectl
sudo apt-mark hold kubelet kubeadm kubectl
	
kubeadm join 192.168.0.2:6443 --token 8tc84w.00ylzxvrt9f6vxcs --discovery-token-ca-cert-hash sha256:0da388e1bba9ce1da909f3cde7f9f6efec0ddf12c41548912733bf7b0f381fb2 --control-plane
