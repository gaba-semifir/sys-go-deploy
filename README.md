# Deploiement automatique de dockerfile dans un cluste k8s

> Ce script permet dans le détails :  
> - de préciser un chemin vers des dockerfiles  
> - de builder en image ces dockerfiles
> - de les poussers sur un registry (local)
> - de les déployers dans un cluster k8s (k3d)

## Pré-requis

|NOM|
|:--:|
|GO|
|Docker|
|K3d ou k8s|
|Kubectl|