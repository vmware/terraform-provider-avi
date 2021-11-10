### AWS Controller deployment and AVI Configuration Examples
** Each section does have the appropriate README. Make sure to read it <br />
** You will start with: <br /> <br />

*** aws_resources – Builds 3 controllers <br/> 
*** avi_user_account – Sets the password for the admin user in order to login from the UI <br/> 
*** avi_cluster – Creates the cluster for the 3 controllers <br/> 
*** avi_infra – Creates the subnets, se groups, sets the flavors and set the demo license <br/>
*** avi_app – This will deploy the web servers backend and configure your virtual servers <br/>

NOTE: Check the single cluster option as well. 


*** DELETE 
To delete your env you must start from the bottom up so reverse order. Keep in mind that if Service Engines have been deployed the terraform destroy may fail. 
