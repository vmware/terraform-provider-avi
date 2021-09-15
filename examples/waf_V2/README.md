### WAF Example
* Using this example we can configure WAF Policy on AVI controller with resources including Waf Profile, WAF Policy, Virtual Service, Pool, Pool Group, SSL profile, VsVip, Network Profile and datasources including WAF Application Signature Provider, Application Profiles, Tenant, Cloud, SSL certificates, SSl Profiles, VRF Context.
* Update required AVI terraform provider version in versions.tf file.
* Update required variables in variables.tf file or you can set those variables individually while running you terraform plan using -var or define those in .tfvars files.
* [Click here](https://avinetworks.com/docs/21.1/waf-policy/) for detailed WAF Policy guide.
