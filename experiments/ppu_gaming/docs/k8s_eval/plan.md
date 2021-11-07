# Plan

## Multinode Cluster
- persistence
- netzwerk
- backups von etcd
- dashboard
- ingress
- nodes dazu nodes weg
- wie verhält sich der cluster wenn die main weg ist ?
- wie monitoren ? prometheus und grafana ?
- CRIO Docker oder Containerd ?
- mit hetzner cloud automatisieren also hoch und runter scalen etc
- Wo und wie läuft die DB ? die darf wohl in k8s
- Alerts und Postmortem
- CICD Gitops und jenkins
- rook für persistenz eval

## SingleNode Cluster
- persistenz auf platte
- selbe fragen wie multinode

## MultiNode HA Cluster
- HA Setup bauen

## Automatisierung
- Terraform und Ansible vs selbst in GO
- Operatoren
- Healtch und Ready Check

## SpieleWolke Cluster
- Schnittstelle in den Cluster über CRD und Operatoren (geschrieben in GO)
- Eigenes Image Repo
- Eigene Pipeline zum bauen der Images ?
- Monitoring
- Backup
- Security 
- ChaosEngeneering
- HA
- Deasaster Pläne
- Alerts
- Quotas / ressourcen Planung
- Docker Images für Mine, OpenTDD, Factorio

## SpieleWolke CustomerCenter
- Webapp mit Kundenverwaltung und Abrechnung
- Starten Stoppen der Spiele
- Persistenz oder auch nicht
- Payment

## Operatoren und CRD Sidequest
- Fragen der IA und sync systemstate damit abbilden sowie selfupdate debug etc
