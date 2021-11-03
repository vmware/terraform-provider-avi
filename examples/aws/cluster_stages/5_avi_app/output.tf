output "aws_webserver_ips" {
  value = aws_instance.terraform-webserver.*.private_ip
}

output "aws_webserver_tags" {
  value = aws_instance.terraform-webserver.*.tags
}