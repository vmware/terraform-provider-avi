output "aws_controller_instance" {
  value = aws_instance.avi_controller[*].tags.Name
}

output "avi_controller_ip" {
  value = aws_instance.avi_controller[*].private_ip
}

output "avi_controller_public_ip_leader" {
  value = aws_eip_association.ctrl_eip_assoc.public_ip
}

output "avi_controller_public_dns" {
  value = aws_eip.ctrl_eip.public_dns
}

output "avi_controller_iam_role" {
  value = data.aws_iam_instance_profile.avi_controller_iam.name
}

output "aws_subnets" {
  value = aws_subnet.terraform-subnet[*].tags.Name
}

output "subnets" {
  value = aws_subnet.terraform-subnet[*].cidr_block
}
