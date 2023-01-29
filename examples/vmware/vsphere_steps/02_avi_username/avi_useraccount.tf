resource "avi_useraccount" "avi_user" {
  username     = var.avi_tenant
  old_password = var.avi_current_password
  password     = var.avi_password == null ? random_string.avi_password_random.result : var.avi_password
}

resource "random_string" "avi_password_random" {
  length           = 8
  special          = true
  min_lower        = 2
  min_upper        = 2
  min_numeric      = 2
  min_special      = 2
  override_special = "_"
}

resource "local_file" "output_passwd_file_random" {
  count   = var.avi_password == null ? 1 : 0
  content     = "{\"avi_password\": ${jsonencode(random_string.avi_password_random.result)}}"
  filename = "../.password.json"
}

resource "local_file" "output_passwd_file_static" {
  count   = var.avi_password == null ? 0 : 1
  content     = "{\"avi_password\": ${jsonencode(var.avi_password)}}"
  filename = "../.password.json"
}
