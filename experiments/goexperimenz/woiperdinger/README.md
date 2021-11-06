# woiperdinger

## Konzept
code generator für web apps  
generiert aus yaml files go web apps  
cruds apis grpc, rest etc  
mit allem was dazu gehört  
mit "hooks" für custom code  
mit eigener web oberfläche die die yaml file generieren kann  
extensions  

1. restful web apps aus yaml
2. weboberfläche für restful web apps
3. komfort -> auth, swagger, docker 

```
woiper new myapi

myapi 
- woiper.yaml
- dockerfile
- README.md
- pkg
-- generated (contains generated code)
-- custom (contains hooks)
- cmd
-- web/main
-- db/ ?
```

## TODO 
- api tmpl
-- dry code
-- middleware for api key auth
-- relations
-- validation
-- multiple dbs
-- openapi
- cmd: turn it into actual usable with flags etc

