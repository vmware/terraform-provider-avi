resource "aws_security_group" "avi_sg" {
  name = var.avi_sg_name

  description = "Allow inbound SSH traffic from my IP"
  vpc_id      = var.aws_vpc_id

  ingress {
    from_port   = "443"
    to_port     = "443"
    protocol    = "TCP"
    cidr_blocks = var.trusted_cidr
    #cidr_blocks = ["0.0.0.0/0"]
  }
  ingress {
    from_port   = "80"
    to_port     = "80"
    protocol    = "TCP"
    cidr_blocks = var.trusted_cidr
    #cidr_blocks = ["0.0.0.0/0"]
  }
  ingress {
    from_port   = "123"
    to_port     = "123"
    protocol    = "TCP"
    cidr_blocks = ["0.0.0.0/0"]
  }
  ingress {
    from_port   = "8443"
    to_port     = "8443"
    protocol    = "TCP"
    cidr_blocks = var.trusted_cidr
    #cidr_blocks = ["0.0.0.0/0"]
  }
  ingress {
    from_port   = "22"
    to_port     = "22"
    protocol    = "TCP"
    cidr_blocks = var.trusted_cidr
    #cidr_blocks = ["0.0.0.0/0"]
  }
  ingress {
    from_port   = "-1"
    to_port     = "-1"
    protocol    = "ICMP"
    cidr_blocks = ["0.0.0.0/0"]
  }
  egress {
    from_port   = "0"
    to_port     = "0"
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}
