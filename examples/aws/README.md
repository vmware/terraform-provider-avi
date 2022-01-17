### AWS Controller deployment and AVI Configuration Examples

- These examples deploy AVI controller on AWS, configures AVI controller cluster, cloud of type AWS, Virtual Services and its dependent resources and data sources.
- [Click here] https://avinetworks.com/docs/21.1/installing-avi-controller-in-amazon-web-services/ to deploy AVI controller on AWS using ami and AWS deployment guide.
- To configure AVI controller cluster check our kb at [Avi KB](https://avinetworks.com/docs/21.1/cluster-configuration-in-aws/) for AVI controller cluster configuration guide in AWS.
- To configure cloud of type CLOUD_AWS on AVI controller. [Click here](https://avinetworks.com/docs/21.1/configuring-avi-vantage-for-application-delivery-in-amazon-web-services/) for detailed cloud configuration guide.
- To configure Virtual Service and dependent resources in AWS environment. [Click here](https://avinetworks.com/docs/21.1/configuration-guide/applications/virtual-services/) for detailed guide to configure virtual service.
- Update required AVI terraform provider version in versions.tf file.
- Update required variables in variables.tf file or you can set those variables individually while running your terraform plan using -var or define those in .tfvars files.
- Each step will have a detail README
