### Azure Controller deployment and AVI Configuration Examples

* These examples deploy AVI controller on Azure, configures AVI controller cluster, cloud of type Azure, Virtual Services and its dependent resources and data sources.
* [Click here](https://github.com/vmware/terraform-provider-avi/tree/eng/examples/azure/azure_resources) to deploy AVI controller on Azure. [Click here](https://avinetworks.com/docs/21.1/avi-deployment-guide-for-microsoft-azure/) for Azure deployment guide.
* [Click here](https://github.com/vmware/terraform-provider-avi/tree/eng/examples/azure/avi_cluster) to configure AVI controller cluster. [Click here](https://avinetworks.com/docs/21.1/azure-cluster-ip/) for AVI controller cluster configuration guide in Azure.
* [Click here](https://github.com/vmware/terraform-provider-avi/tree/eng/examples/aws/avi_infra) to configure cloud of type CLOUD_AZURE on AVI controller. [Click here](https://avinetworks.com/docs/21.1/configuring-avi-vantage-for-application-delivery-in-microsoft-azure/) for detailed cloud configuration guide.
* [Click here](https://github.com/vmware/terraform-provider-avi/tree/eng/examples/azure/avi_azure_app) to configure Virtual Service and dependent resources in Azure enviornment. [Click here](https://avinetworks.com/docs/21.1/configuration-guide/applications/virtual-services/) for detailed guide to configure virtual service. 
* Update required AVI terraform provider version in versions.tf file.
* Update required variables in variables.tf file or you can set those variables individually while running you terraform plan using -var or define those in .tfvars files.
