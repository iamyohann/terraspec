package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/iamyohann/terraspec/pkg/project"
)

func main() {
	wrkdir, _ := os.Getwd()
	folder := "examples/terraform/simple-db"
	dir := filepath.Join(wrkdir, folder)

	project, _ := project.FromDirectory(dir)
	databases := project.FindResources("aws_db_instance")

	for _, resource := range databases {
		for key, value := range resource.Attributes {
			fmt.Println("Key = ", key, "| Value = ", value)
		}
	}

}
