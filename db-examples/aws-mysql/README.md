# aws-mysql

### A simple deployment and connection to AWS MySQL with Terraform and Go

1. Create a terraform.tfvars file within the terraform directory, adding in the following variables:
```
mysql-user = "username that you want"

mysql-password = "your desired password"

dbname = "name of your db to create"
```
2. Create a *.env* file in project root
3. `terraform init`
4. `terraform apply`
5. Copy the `mysql-details` output from above and place in .env, along with db user name, and password.
```
SQL_STRING=sampleendpointhost.....us-east-1.rds.amazonaws.com:3306
SQL_USER=sampleuser
SQL_PASSWORD=samplepass
```
6. `go run main.go` will select your database and create a simple table. 