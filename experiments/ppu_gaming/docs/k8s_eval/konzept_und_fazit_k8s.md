# Pay Per Use Gaming

Die Idee ist das man sich spiele server starten und stoppen kann wie man will und man nur für die genutze zeit bezahlt. eigentlich wie mit vms bei den cloud providern.  
Mit Webbasiertem oneclick starten bissl server konfiguieren mit dazubuchbarer persistenz oder wenn man will einfach savegames runter und hochladen.  
Spiele für den beginn: MC, factorio, opentdd  
names idee: SpieleWolke  

# Technik
Technisch wird das ein k8s cluster mit persistenz. Jeder Server ist ein docker container der eine perstitenz mounted die wir dannach löschen oder auch nicht.  
Dieser Cluster wird vollautomatisiert. Monitoring / backup / security / alerts nicht vergessen.  
Vollautomatisierung auch mit self heal etc.  
Das Webportal dienst für nutzer anmeldung server kontrolieren etc. soll auch einen "admin" bereich mit dashboards etc haben.  
Auch abbrechnung implementieren  
Das Projekt soll voll in go geschrieben werden mit mysql 8 als db  
MonoRepo  

# Sideeffect
Learning auch für die rechnungs wolke

# Steps
1. K8s Cluster aufbauen
2. Container erstellen
3. Container mal mit persistenz deployen und ausmachen und auch an quotas denken
4. "controller" service erstellen der container spawnen ausschalten und alles andere versteuert (agent artig)
5. Monitoring / backup / security / selfheal
6. ci cd für controller
7. web dran bauen. User login spiele starten stoppen configs ändern save games downloaden.
8. ci cd für web
9. admin views
10. abbrechnung


# Fazit Eval k8s cluster
der k8s cluster ist zu kompliziert und auserdem ist nicht klar wie er sich auf die performance auswirkt.
Deshalb wird einfach ein cloud vm provider wie zb. hetzner oder digtial ocean benutzt und per api gesteuert.
ist einfacher und sicher perfomanter
