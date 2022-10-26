### Azure Controller deployment and AVI Configuration Examples

* These examples deploy AVI controller on Azure, configures AVI controller cluster, cloud of type Azure, Virtual Services and its dependent resources and data sources.
* Under the folder cluster_stages to deploy AVI controller on Azure. [Click here](https://avinetworks.com/docs/21.1/avi-deployment-guide-for-microsoft-azure/) for Azure deployment guide.
* Step 3 will allow you to configure AVI controller cluster. [Click here](https://avinetworks.com/docs/21.1/azure-cluster-ip/) for AVI controller cluster configuration guide in Azure.
* Step 4 will allow you to configure cloud of type CLOUD_AZURE on AVI controller. [Click here](https://avinetworks.com/docs/21.1/configuring-avi-vantage-for-application-delivery-in-microsoft-azure/) for detailed cloud configuration guide.
* Step 6 will configure Virtual Service and dependent resources in Azure environment. [Click here](https://avinetworks.com/docs/21.1/configuration-guide/applications/virtual-services/) for detailed guide to configure virtual service. 
* Each step contains the provider file to set the version need for the AVI terraform provider.
* Use the terraform.tfvars.examples on each folder to set sensetive values. Default .gitignore automatically will not push that file into git.
* avi_saml_auth, is from the old example, I personally have not tested.
