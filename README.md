# terraspec
Terraform testing framework


## Usage

```go
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
```

```bash
Key =  engine_version | Value =  "5.7"
Key =  instance_class | Value =  "db.t3.micro"
Key =  name | Value =  "simple_aws_db_instance"
Key =  username | Value =  "username"
Key =  password | Value =  "password"
Key =  skip_final_snapshot | Value =  true
Key =  allocated_storage | Value =  5
Key =  engine | Value =  "mysql"
```