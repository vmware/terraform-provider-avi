# Create VPC and Subnets for AVI Controller and SEs
resource "aws_vpc" "avi" {
  count      = var.create_networking ? 1 : 0
  cidr_block = var.avi_cidr_block

  tags = {
    Name = "${var.name_prefix}-avi-vpc"
  }
}

resource "aws_subnet" "avi" {
  for_each                = var.create_networking ? { for idx, az_name in local.az_names : idx => az_name } : {}
  vpc_id                  = aws_vpc.avi[0].id
  cidr_block              = cidrsubnet(aws_vpc.avi[0].cidr_block, 8, 230 + each.key)
  availability_zone       = local.az_names[each.key]
  map_public_ip_on_launch = var.controller_public_address ? true : false
  tags = {
    Name = "${var.name_prefix}-avi-subnet-${local.az_names[each.key]}"
  }
  lifecycle {
    ignore_changes = [tags]
  }
}

resource "aws_internet_gateway" "avi" {
  count  = var.create_networking ? 1 : 0
  vpc_id = aws_vpc.avi[0].id
  tags = {
    Name = "${var.name_prefix}-avi-igw"
  }
}

resource "aws_route" "default_route" {
  count                  = var.create_networking ? 1 : 0
  route_table_id         = aws_vpc.avi[0].main_route_table_id
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = aws_internet_gateway.avi[0].id
}