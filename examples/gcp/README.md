### GCP Controller deployment and AVI Configuration Examples
* These examples deploy AVI controller on GCP, configures AVI controller cluster and cloud of type GCP.
* [Click here](https://github.com/vmware/terraform-provider-avi/tree/eng/examples/gcp/customRoleCreation) to create custom GCP role with required permissions to deploy controller VMs on GCP. [Click here](https://avinetworks.com/docs/20.1/gcp-full-access-roles-and-permissions/) for detailed Roles and Permissions guide. 
* [Click here](https://github.com/vmware/terraform-provider-avi/tree/eng/examples/gcp/controllerCreation) to deploy AVI controller on GCP. [Click here](https://avinetworks.com/docs/20.1/gcp-full-access-deployment-guide/) for GCP deployment guide.
* [Click here](https://github.com/vmware/terraform-provider-avi/blob/eng/examples/gcp/cloudCreation/main.tf) to configure cloud of type CLOUD_GCP on AVI controller. [Click here](https://avinetworks.com/docs/20.1/configuring-gcp-cloud-network/) for detailed cloud configuration guide.
* Update required AVI terraform provider version in versions.tf file.
* Update required variables in variables.tf file or you can set those variables individually while running you terraform plan using -var or define those in .tfvars files.
