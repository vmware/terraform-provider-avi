resource "aws_security_group" "avi_controller_sg" {
  name        = "${var.name_prefix}-avi-controller-sg"
  description = "Allow Traffic for AVI Controller"
  vpc_id      = var.create_networking ? aws_vpc.avi[0].id : var.custom_vpc_id

  ingress {
    description = "SSH from Internet"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  ingress {
    description = "HTTPS from Internet"
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  ingress {
    description = "Secure Channel from VPC"
    from_port   = 8443
    to_port     = 8443
    protocol    = "tcp"
    cidr_blocks = [var.avi_cidr_block]
  }
  ingress {
    description = "AVI_CLI from Internet"
    from_port   = 5054
    to_port     = 5054
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  ingress {
    description = "ICMP to Controller"
    from_port   = -1
    to_port     = -1
    protocol    = "icmp"
    cidr_blocks = [var.avi_cidr_block]
  }
  egress {
    description = "Allow traffic from VPC"
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "${var.name_prefix}-avi-controller-sg"
  }
}
resource "aws_security_group" "avi_se_mgmt_sg" {
  name        = "${var.name_prefix}-avi-se-mgmt-sg"
  description = "Allow traffic for AVI SE MGMT NICs"
  vpc_id      = var.create_networking ? aws_vpc.avi[0].id : var.custom_vpc_id

  ingress {
    description = "SSH to SE"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = [var.avi_cidr_block]
  }
  ingress {
    description = "ICMP to SE"
    from_port   = -1
    to_port     = -1
    protocol    = "icmp"
    cidr_blocks = [var.avi_cidr_block]
  }
  ingress {
    description = "ETHERIP to SE"
    from_port   = -1
    to_port     = -1
    protocol    = "97"
    cidr_blocks = [var.avi_cidr_block]
  }
  ingress {
    description = "CPHB to SE"
    from_port   = -1
    to_port     = -1
    protocol    = "73"
    cidr_blocks = [var.avi_cidr_block]
  }
  ingress {
    description = "Proto63 to SE"
    from_port   = -1
    to_port     = -1
    protocol    = "63"
    cidr_blocks = [var.avi_cidr_block]
  }
  egress {
    description = "Allow SSH to Controller"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = [var.avi_cidr_block]
  }
  egress {
    description = "Allow 8443 to Controller"
    from_port   = 8443
    to_port     = 8443
    protocol    = "tcp"
    cidr_blocks = [var.avi_cidr_block]
  }
  egress {
    description = "Allow NTP to Controller"
    from_port   = 123
    to_port     = 123
    protocol    = "udp"
    cidr_blocks = [var.avi_cidr_block]
  }
  tags = {
    Name = "${var.name_prefix}-avi-se-mgmt-sg"
  }
}
resource "aws_security_group" "avi_data_sg" {
  name        = "${var.name_prefix}-avi-data-sg"
  description = "Allow traffic for Avi SE Data NICs"
  vpc_id      = var.create_networking ? aws_vpc.avi[0].id : var.custom_vpc_id

  ingress {
    description = "Allow Traffic from VPC"
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = [var.avi_cidr_block]
  }
  ingress {
    description = "HTTPS Inbound"
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  ingress {
    description = "HTTP Inbound"
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  egress {
    description = "Allow Traffic Outbound"
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
  tags = {
    Name = "${var.name_prefix}-avi-data-sg"
  }
}