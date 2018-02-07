provider "aws" {
  access_key = "${var.aws_access_key}"

  secret_key = "${var.aws_secret_key}"
  region     = "${var.aws_region}"
}

resource "aws_instance" "avi_controller" {
  ami           = "${var.avi_controller_ami}"
  instance_type = "c4.2xlarge"
  subnet_id     = "${aws_subnet.terraform-subnet.0.id}"

  tags {
    Name = "${var.project_name}-terraform-controller"
  }
}

output "aws_controller_instance" {
  value = "${aws_instance.avi_controller.tags.Name}"
}

output "avi_controller_ip" {
  value = "${aws_instance.avi_controller.private_ip}"
}

resource "aws_subnet" "terraform-subnet" {
  count = "${length(var.aws_subnets)}"

  vpc_id            = "${var.aws_vpc_id}"
  availability_zone = "${var.aws_availability_zones[count.index]}"
  cidr_block        = "${var.aws_subnets[count.index]}/${var.aws_subnet_mask}"

  tags {
    Name    = "${var.project_name}-terraform-subnet-${count.index}"
    Project = "${var.project_name}-terraform-subnets"
  }
}

output "aws_subnets" {
  value = "${aws_subnet.terraform-subnet.*.tags.Name}"
}
