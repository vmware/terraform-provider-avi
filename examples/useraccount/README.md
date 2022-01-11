### Change controller password using useraccount resource

* Initially configure the current password in both the places
  * provider initialization(var.admin_old_password)
  * old_password field of useraccount resource.
```
provider "avi" {
  avi_username   = var.admin_username
  avi_password   = var.admin_old_password
  avi_controller = var.controller_ip
  avi_tenant     = "admin"
  avi_version = var.avi_version
}
```

```
resource "avi_useraccount" "avi_user" {
  username     = "admin"
  old_password = var.admin_old_password
  password     = var.admin_new_password
}
```
* Also configure depends_on in all other resources. 
Terraform runs resources and datasources concurrently. 
Useraccount password reset functionality will invalidate the session at some point which invalidates
the controller session and there is possibility that other resources will harm due to this behaviour.
Example
```
data "avi_cloud" "default_cloud" {
    name= "Default-Cloud"
	depends_on = [avi_useraccount.avi_user]
}
```
* Then before next run set the new password for both old_password(var.admin_old_password) and
new_password(var.admin_new_password) variables. 
* **(Optional)** Now you can assign var.admin_new_password instead of var.admin_old_password
to the provider password variable.
```
provider "avi" {
  avi_username   = var.admin_username
  avi_password   = var.admin_new_password
  avi_controller = var.controller_ip
  avi_tenant     = "admin"
  avi_version = var.avi_version
}
```
