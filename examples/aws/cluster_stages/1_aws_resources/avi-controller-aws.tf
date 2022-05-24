data "aws_ami" "latest-avi_controller" {
  most_recent = true
  owners      = ["aws-marketplace"]

  filter {
    name   = "name"
    values = ["Avi*Controller-*"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }
}

data "aws_iam_instance_profile" "avi_controller_iam" {
  name = var.avi_controller_iam
}
resource "aws_eip" "ctrl_eip" {
  vpc = true
}

resource "aws_eip_association" "ctrl_eip_assoc" {
  instance_id   = aws_instance.avi_controller[0].id
  allocation_id = aws_eip.ctrl_eip.id

}

resource "aws_instance" "avi_controller" {
  ami                  = data.aws_ami.latest-avi_controller.id
  count                = var.controller_counts
  instance_type        = "c4.2xlarge"
  subnet_id            = aws_subnet.terraform-subnet[0].id
  iam_instance_profile = data.aws_iam_instance_profile.avi_controller_iam.name
  vpc_security_group_ids = [
    aws_security_group.avi_sg.id,
  ]
  key_name = aws_key_pair.generated.key_name

  tags              = {
    Name            = "${var.project_name}-terraform-controller-${count.index + 1}"
    shutdown_policy = var.shutdown_policy
    department      = var.department
  }
}
resource "aws_subnet" "terraform-subnet" {
  count = length(var.aws_subnets)

  vpc_id            = var.aws_vpc_id
  availability_zone = var.aws_availability_zones[count.index]
  cidr_block        = "${var.aws_subnets[count.index]}/${var.aws_subnet_mask}"
  tags = {
    Name    = "${var.project_name}-terraform-subnet-${count.index}"
    Project = "${var.project_name}-terraform-subnets"
  }
}

resource "null_resource" "wait_for_controller" {
  count = length(aws_instance.avi_controller)
  provisioner "local-exec" {
    command = "./wait-for-controller.sh"
    environment = {
      CONTROLLER_ADDRESS = aws_eip_association.ctrl_eip_assoc.public_ip
      POLL_INTERVAL      = 45
    }
  }
  depends_on = [
    aws_instance.avi_controller
  ]
}
