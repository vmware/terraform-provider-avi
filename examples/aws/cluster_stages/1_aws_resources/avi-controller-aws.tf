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
  ami                  = var.avi_controller_ami
  count                = var.controller_counts
  instance_type        = "c4.2xlarge"
  subnet_id            = aws_subnet.terraform-subnet[0].id
  iam_instance_profile = data.aws_iam_instance_profile.avi_controller_iam.name
  key_name             = aws_key_pair.generated.key_name

  tags = {
    Name = "${var.project_name}-terraform-controller-${count.index + 1}"
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
