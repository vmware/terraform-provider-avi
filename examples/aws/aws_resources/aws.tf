provider "aws" {
  access_key = "${var.aws_access_key}"

  secret_key = "${var.aws_secret_key}"
  region     = "${var.aws_region}"
}

resource "aws_instance" "avi_controller" {
  ami           = "ami-2c0bbf54"
  instance_type = "c4.2xlarge"
  subnet_id     = "${aws_subnet.grastogi-terraform-subnet.id}"

  tags {
    Name = "grastogi-terraform-controller"
  }
}

output "aws_controller_instance" {
  value = "${aws_instance.avi_controller.tags.Name}"
}

output "avi_controller_ip" {
  value = "${aws_instance.avi_controller.private_ip}"
}

resource "aws_subnet" "grastogi-terraform-subnet" {
  vpc_id     = "${var.aws_vpc_id}"
  cidr_block = "10.144.43.0/24"

  tags {
    Name = "grastogi-terraform-subnet"
  }
}

output "aws_subnet" {
  value = "${aws_subnet.grastogi-terraform-subnet.tags.Name}"
}

resource "aws_instance" "terraform-webserver1" {
  ami           = "ami-05141f7c"
  instance_type = "t2.micro"
  subnet_id     = "${aws_subnet.grastogi-terraform-subnet.id}"

  tags {
    Name = "grastogi-terraform-webserver1"
  }
}

output "aws_webserver1_ip" {
  value = "${aws_instance.terraform-webserver1.private_ip}"
}

resource "aws_instance" "terraform-webserver2" {
  ami           = "ami-05141f7c"
  instance_type = "t2.micro"
  subnet_id     = "${aws_subnet.grastogi-terraform-subnet.id}"

  tags {
    Name = "grastogi-terraform-webserver2"
  }
}

output "aws_webserver2_ip" {
  value = "${aws_instance.terraform-webserver2.private_ip}"
}

resource "aws_instance" "terraform-webserver3" {
  ami           = "ami-05141f7c"
  instance_type = "t2.micro"
  subnet_id     = "${aws_subnet.grastogi-terraform-subnet.id}"

  tags {
    Name = "grastogi-terraform-webserver3"
  }
}

output "aws_webserver3_ip" {
  value = "${aws_instance.terraform-webserver3.private_ip}"
}

resource "aws_instance" "terraform-webserver4" {
  ami           = "ami-05141f7c"
  instance_type = "t2.micro"
  subnet_id     = "${aws_subnet.grastogi-terraform-subnet.id}"

  tags {
    Name = "grastogi-terraform-webserver4"
  }
}

output "aws_webserver4_ip" {
  value = "${aws_instance.terraform-webserver4.private_ip}"
}
