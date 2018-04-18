provider "aws" {
  region  = "us-east-1"
  profile = "tfbuild"
}

variable "mysql-user" {}
variable "mysql-password" {}
variable "dbname" {}

resource "aws_security_group" "mysql-sg" {
  name = "mysql-sg"

  // sample, allow any for testing
  ingress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_db_instance" "mysql" {
  allocated_storage   = 10
  engine              = "mysql"
  instance_class      = "db.t2.micro"
  name                = "${var.dbname}"
  identifier          = "mysqlinstance"
  username            = "${var.mysql-user}"
  password            = "${var.mysql-password}"
  publicly_accessible = true
  skip_final_snapshot = true

  vpc_security_group_ids = [
    "${aws_security_group.mysql-sg.id}",
  ]
}

output "mysql-details" {
  value = "${aws_db_instance.mysql.endpoint}"
}
