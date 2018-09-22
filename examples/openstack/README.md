# OpenStack terraform example

## Instruction to lanch an instance 

To run OpenStack example you should know the following things of OpenStack
```
 - Openstack credentials
 - Network name
 - Subnet 
 - You should have image file 
 - Controller default password
```

## Commands to run terraform scripts 
```
terraform init
terraform plan
terraform apply
```

To launch an instance just go inside the launch_instance directory and run following commands while running just pass the required arguments that it prompts
```
terraform init
terraform plan
terraform apply
```

To pass default values to terraform script then create terraform.tfvars file in the directory and add values into that file in the following format

```
openstack_url = "http://11.5.26.67:/v3"
openstack_username = "username"
openstack_password = "something"
avi_controller = "48.48.48.48"
avi_password = "something"
avi_tenant = "admin"
openstack_subnet = "subnet"
network_name = "network"
op_subnet_ip = "10.23.4.55"
``` 

To set password to the controller go inside the avi_user_account and run same commands

To set up OpenStack cloud on controller go inside the setup_cloud directory and run the same commands

To create an application on controller go inside the web_app directory and run the same commands 


