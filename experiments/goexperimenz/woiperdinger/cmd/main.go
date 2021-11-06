package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

type ResourceField struct {
	Name string
	Type string
}

type Resource struct {
	Name   string
	Fields []ResourceField
}

type Endpoint struct {
	Name     string
	Resource string
}

type AppDefinition struct {
	Name      string
	Resources []Resource
	Endpoints []Endpoint
}

func (a *AppDefinition) PPrint() {
	fmt.Println("Name: " + a.Name)
	fmt.Println("Resources:")
	for _, resource := range a.Resources {
		fmt.Println("    Name: " + resource.Name)
		fmt.Println("    Fields: ")
		for _, field := range resource.Fields {
			fmt.Println("        Name: " + field.Name)
			fmt.Println("        Type: " + field.Type)
		}
	}
	fmt.Println("Endpoints:")
	for _, endpoint := range a.Endpoints {
		fmt.Println("    Name: " + endpoint.Name)
		fmt.Println("    Resource: " + endpoint.Resource)
	}
}

func main() {
	tmpl := template.Must(template.ParseFiles("templates/api.go.tmpl"))

	rawAppDef, err := ioutil.ReadFile("example/todo.yaml")
	var appDef AppDefinition
	err = yaml.Unmarshal(rawAppDef, &appDef)
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}
	//appDef.PPrint()

	f, err := os.Create("example/example.go")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(f, appDef)
	if err != nil {
		log.Fatal(err)
	}
}
