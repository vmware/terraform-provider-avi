output "aws_webserver_ips" {
  value = aws_instance.terraform-webserver.*.private_ip
}

output "aws_webserver_tags" {
  value = aws_instance.terraform-webserver.*.tags
}
output "aws_vs_vip_0" {
  value = avi_vsvip.terraform-vip.vip[0]
}

output "aws_vs_vip" {
  value = avi_virtualservice.terraform-virtualservice.vip
}


