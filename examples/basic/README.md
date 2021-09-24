### Basic Virtual Service Example
* Using this example we can configure HTTPS Virtual Service on AVI controller with dependent resources including Pool, Poolgroup, SSL Certificates, Health Monitor, VsVip, Application Persistence Profile, Network Profile and datasources including Application Profiles, Tenant, Cloud, Service Engine Group, Network Profile, Analytics Profile, SSL certificates, SSl Profiles, VRF Context. 
* We are reading certificates and key from files app_cert.crt and app_cert.key for SSL key and certificates resource. 
* Update required AVI terraform provider version in versions.tf file.
* Update required variables in variables.tf file or you can set those variables individually while running you terraform plan using -var or define those in .tfvars files.
* [Click here](https://avinetworks.com/docs/20.1/configuration-guide/applications/virtual-services/) for detailed Virtual Service guide.
