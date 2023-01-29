output "Password_Set" {
  value = var.avi_password == null ? random_string.avi_password_random.result : var.avi_password
}
