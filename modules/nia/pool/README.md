# Pool module for Network Infrastructure Automation (NIA)


#### Note: This Terraform module is designed to be used only with **consul-terraform-sync**

## Feature
This module supports the following:
* Scale up and scale down Pool and Pool Members(servers) based on the service(s) configuration in Consul catalog.


## What is consul-terraform-sync?

The **consul-terraform-sync** runs as a daemon that enables a **publisher-subscriber** paradigm between **Consul** and **Controller** to support **Network Infrastructure Automation (NIA)**. 

<p align="left">
<img width="800" src=""> </a>
</p>

* consul-terraform-sync **subscribes to updates from the Consul catalog** and executes one or more automation **"tasks"** with appropriate value of *service variables* based on those updates. **consul-terraform-sync** leverages [Terraform](https://www.terraform.io/) as the underlying automation tool and utilizes the Terraform provider ecosystem to drive relevant change to the network infrastructure. 

* Each task consists of a runbook automation written as a compatible **Terraform module** using resources and data sources for the underlying network infrastructure provider.

Please refer to this [link](https://www.consul.io/docs/nia/installation/install) for getting started with **consul-terraform-sync**



## Requirements

| Name | Version |
|------|---------|
| terraform | \>= 0.13 |
| consul-terraform-sync | \>= 0.1.0 |
| consul | \>= 1.7 |

## Providers

| Name | Version |
|------|---------|
| vmware/avi | \>= 20.1.4 |

## Usage
In order to use this module, you will need to install **consul-terraform-sync**, create a **"task"** with this Terraform module as a source within the task, and run **consul-terraform-sync**.

The users can subscribe to the services in the consul catalog and define the Terraform module which will be executed when there are any updates to the subscribed services using a **"task"**.

**~> Note:** It is recommended to have the [consul-terraform-sync config guide](https://www.consul.io/docs/nia/installation/configuration) for reference.
1. Download the **consul-terraform-sync** on a node which is highly available (prefrably, a node running a consul client)
2. Add **consul-terraform-sync** to the PATH on that node
3. Check the installation
```bash
wget https://releases.hashicorp.com/consul-terraform-sync/<consul-terraform-sync-version>/consul-terraform-sync_<consul-terraform-sync-version>_linux_amd64.zip
unzip consul-terraform-sync_<consul-terraform-sync-version>_linux_amd64.zip
sudo mv consul-terraform-sync /usr/local/bin
consul-terraform-sync --version
```

 4. Create a config file **"tasks.hcl"** for consul-terraform-sync. Please note that this just an example. 
```terraform
log_level = <log_level> # eg. "info"

consul {
  address = "<consul agent address>" # eg. "1.1.1.1:8500"
}

driver "terraform" {
  log = true
  required_providers {
    avi = {
    source =  "vmware/avi"
    version = "20.1.4"
    }
  }
}

terraform_provider "avi" {
  avi_username   = "admin"
  avi_tenant     = "admin"
  avi_password   = "{{ env \"avi_password\" }}"
  avi_controller = "<avi_controller_address>"
  avi_version    = "<api_version>"
  avi_api_timeout    = 50
}

task {
  name = <name of the task (has to be unique)> # eg. "avi-pool-counting"
  description = <description of the task> # eg. "Automatically Scale/Configure AVI Pool Servers"
  source = "vmware/modules/nia/pool" # to be updated
  providers = ["avi"]
  services = ["<list of services you want to subscribe to>"] # eg. ["web", "counting"]
  variable_files = ["<list of files that have user variables for this module (please input full path)>"] # eg. ["/home/user/cts-poc/input_counting.tf"]
  enabled = true
}
```
 5. Start consul-terraform-sync
```
$ consul-terraform-sync -config-file=tasks.hcl
```
**consul-terraform-sync** will create pool for each task and subscribed services from consul catalog will become the servers of pool.

**consul-terraform-sync is now subscribed to the Consul catalog. Any updates to the services identified in the task will result in updating the servers config of pool on avi controller**

**~> Note:** If you are interested in how **consul-terraform-sync** works, please refer to this [section](#how-does-consul-terraform-sync-work).




## Inputs

| Name | Description | Type |  Required |
|------|-------------|------|:--------:|
| name | The name of the pool e.g. pool1 | `string` | yes |
| default\_server\_port | Traffic sent to servers will use this destination server port unless overridden by the server's specific port attribute. The SSL checkbox enables Avi to server encryption. Allowed values are 1-65535. Default value when not specified in API or module is interpreted by ALB Controller as 80. | `number` | no |
| graceful\_disable\_timeout | Used to gracefully disable a server. Virtual service waits for the specified time before terminating the existing connections  to the servers that are disabled. Allowed values are 1-7200. Special values are 0 - 'Immediate', -1 - 'Infinite'. Unit is MIN. Default value when not specified in API or module is interpreted by ALB Controller as 1. | `number` | no |
| connection\_ramp\_duration | Duration for which new connections will be gradually ramped up to a server recently brought online. Useful for LB algorithms that are least connection based. Allowed values are 1-300. Special values are 0 - 'Immediate'. Unit is MIN. Allowed in Basic(Allowed values- 0) edition, Essentials(Allowed values- 0) edition, Enterprise edition. Special default for Basic edition is 0, Essentials edition is 0, Enterprise is 10. Default value when not specified in API or module is interpreted by ALB Controller as 10. | `number` | no |
| max\_concurrent\_connections\_per\_server | The maximum number of concurrent connections allowed to each server within the pool. NOTE  applied value will be no less than the number of service engines that the pool is placed on. If set to 0, no limit is applied. Default value when not specified in API or module is interpreted by ALB Controller as 0. | `number` | no |
| lb\_algorithm | The load balancing algorithm will pick a server within the pool's list of available servers. Values LB_ALGORITHM_NEAREST_SERVER and LB_ALGORITHM_TOPOLOGY are only allowed for GSLB pool. Enum options - LB_ALGORITHM_LEAST_CONNECTIONS, LB_ALGORITHM_ROUND_ROBIN, LB_ALGORITHM_FASTEST_RESPONSE, LB_ALGORITHM_CONSISTENT_HASH, LB_ALGORITHM_LEAST_LOAD, LB_ALGORITHM_FEWEST_SERVERS, LB_ALGORITHM_RANDOM, LB_ALGORITHM_FEWEST_TASKS, LB_ALGORITHM_NEAREST_SERVER, LB_ALGORITHM_CORE_AFFINITY, LB_ALGORITHM_TOPOLOGY. Allowed in Basic(Allowed values- LB_ALGORITHM_LEAST_CONNECTIONS,LB_ALGORITHM_ROUND_ROBIN,LB_ALGORITHM_CONSISTENT_HASH) edition, Essentials(Allowed values- LB_ALGORITHM_LEAST_CONNECTIONS,LB_ALGORITHM_ROUND_ROBIN,LB_ALGORITHM_CONSISTENT_HASH) edition, Enterprise edition. Default value when not specified in API or module is interpreted by ALB Controller as LB_ALGORITHM_LEAST_CONNECTIONS. | `string` | no |
| lb\_algorithm\_hash | Criteria used as a key for determining the hash between the client and  server. Enum options - LB_ALGORITHM_CONSISTENT_HASH_SOURCE_IP_ADDRESS, LB_ALGORITHM_CONSISTENT_HASH_SOURCE_IP_ADDRESS_AND_PORT, LB_ALGORITHM_CONSISTENT_HASH_URI, LB_ALGORITHM_CONSISTENT_HASH_CUSTOM_HEADER, LB_ALGORITHM_CONSISTENT_HASH_CUSTOM_STRING, LB_ALGORITHM_CONSISTENT_HASH_CALLID. Allowed in Basic(Allowed values- LB_ALGORITHM_CONSISTENT_HASH_SOURCE_IP_ADDRESS) edition, Essentials(Allowed values- LB_ALGORITHM_CONSISTENT_HASH_SOURCE_IP_ADDRESS) edition, Enterprise edition. Default value when not specified in API or module is interpreted by ALB Controller as LB_ALGORITHM_CONSISTENT_HASH_SOURCE_IP_ADDRESS. | `string` | no |
| lb\_algorithm\_consistent\_hash\_hdr | HTTP header name to be used for the hash key. | `string` | no |
| application\_persistence\_profile | Persistence will ensure the same user sticks to the same server for a desired duration of time. It is a reference to an object of type ApplicationPersistenceProfile. | `string` | no |
| ssl\_profile | When enabled, Avi re-encrypts traffic to the backend servers. The specific SSL profile defines which ciphers and SSL versions will be supported. It is a reference to an object of type SSLProfile. | `string` | no |
| inline\_health\_monitor | The Passive monitor will monitor client to server connections and requests and adjust traffic load to servers based on successful responses. This may alter the expected behavior of the LB method, such as Round Robin. Default value when not specified in API or module is interpreted by ALB Controller as True. | `bool` | no |
| use\_service\_port | Do not translate the client's destination port when sending the connection to the server. The pool or servers specified service port will still be used for health monitoring. Allowed in Basic(Allowed values- False) edition, Essentials(Allowed values- False) edition, Enterprise edition. Default value when not specified in API or module is interpreted by ALB Controller as False. | `bool` | no |
| capacity\_estimation | Inline estimation of capacity of servers. Allowed in Basic(Allowed values- False) edition, Essentials(Allowed values- False) edition, Enterprise edition. Default value when not specified in API or module is interpreted by ALB Controller as False. | `bool` | no |
| capacity\_estimation\_ttfb\_thresh | The maximum time-to-first-byte of a server. Allowed values are 1-5000. Special values are 0 - 'Automatic'. Unit is MILLISECONDS. Allowed in Basic(Allowed values- 0) edition, Essentials(Allowed values- 0) edition, Enterprise edition. Default value when not specified in API or module is interpreted by ALB Controller as 0. | `number` | no |
| pki\_profile | Avi will validate the SSL certificate present by a server against the selected PKI Profile. It is a reference to an object of type PKIProfile. | `string` | no |
| ssl\_key\_and\_certificate\_uuid | Service Engines will present a client SSL certificate to the server. It is a reference to an object of type SSLKeyAndCertificate. | `string` | no |
| autoscale\_networks | Network Ids for the launch configuration. | `list(string)` | no |
| autoscale\_policy | Reference to Server Autoscale Policy. It is a reference to an object of type ServerAutoScalePolicy. | `string` | no |
| autoscale\_launch\_config | If configured then Avi will trigger orchestration of pool server creation and deletion. It is a reference to an object of type AutoScaleLaunchConfig. | `string` | no |
| vrf | Virtual Routing Context that the pool is bound to. | `string` | no |
| ipaddrgroup | Use list of servers from Ip Address Group. It is a reference to an object of type IpAddrGroup. | `string` | no |
| fewest\_tasks\_feedback\_delay | Periodicity of feedback for fewest tasks server selection algorithm. Allowed values are 1-300. Unit is SEC. Default value when not specified in API or module is interpreted by ALB Controller as 10. | `number` | no |
| enabled | Enable or disable the pool. Disabling will terminate all open connections and pause health monitors. Default value when not specified in API or module is interpreted by ALB Controller as True. | `bool` | no |
| east\_west | Inherited config from VirtualService. | `bool` | no |
| cloud\_config\_cksum | Checksum of cloud configuration for Pool. Internally set by cloud connector. | `string` | no |
| request\_queue\_enabled | Enable request queue when pool is full. Allowed in Basic(Allowed values- False) edition, Essentials(Allowed values- False) edition, Enterprise edition. Default value when not specified in API or module is interpreted by ALB Controller as False. | `bool` | no |
| request\_queue\_depth | Minimum number of requests to be queued when pool is full. Allowed in Basic(Allowed values- 128) edition, Essentials(Allowed values- 128) edition, Enterprise edition. Default value when not specified in API or module is interpreted by ALB Controller as 128. | `number` | no |
| host\_check\_enabled | Enable common name check for server certificate. If enabled and no explicit domain name is specified, Avi will use the incoming host header to do the match. Default value when not specified in API or module is interpreted by ALB Controller as False. | `bool` | no |
| sni\_enabled | Enable TLS SNI for server connections. If disabled, Avi will not send the SNI extension as part of the handshake. Default value when not specified in API or module is interpreted by ALB Controller as True. | `bool` | no |
| server\_name | Fully qualified DNS hostname which will be used in the TLS SNI extension in server connections if SNI is enabled. If no value is specified, Avi will use the incoming host header instead. | `string` | no |
| rewrite\_host\_header\_to\_sni | If SNI server name is specified, rewrite incoming host header to the SNI server name. Default value when not specified in API or module is interpreted by ALB Controller as False. | `bool` | no |
| rewrite\_host\_header\_to\_server\_name | Rewrite incoming Host Header to server name of the server to which the request is proxied. Enabling this feature rewrites Host Header for requests to all servers in the pool. Default value when not specified in API or module is interpreted by ALB Controller as False. | `bool` | no |
| nsx\_securitygroup | A list of NSX Groups where the Servers for the Pool are created . | `list(string)` | no |
| external\_autoscale\_groups | Names of external auto-scale groups for pool servers. Currently available only for AWS and Azure. | `list(string)` | no |
| domain\_name | Comma separated list of domain names which will be used to verify the common names or subject alternative names presented by server certificates. It is performed only when common name check host_check_enabled is enabled. | `list(string)` | no |
| lb\_algorithm\_core\_nonaffinity | Degree of non-affinity for core affinity based server selection. Allowed values are 1-65535. Allowed in Basic(Allowed values- 2) edition, Essentials(Allowed values- 2) edition, Enterprise edition. Default value when not specified in API or module is interpreted by ALB Controller as 2. | `number` | no |
| lookup\_server\_by\_name | Allow server lookup by name. Allowed in Basic(Allowed values- False) edition, Essentials(Allowed values- False) edition, Enterprise edition. Default value when not specified in API or module is interpreted by ALB Controller as False. | `bool` | no |
| analytics\_profile | Specifies settings related to analytics. It is a reference to an object of type AnalyticsProfile. | `string` | no |
| service\_metadata | Metadata pertaining to the service provided by this Pool. In Openshift/Kubernetes environments, app metadata info is stored. Any user input to this field will be overwritten by Avi Vantage. | `string` | no |
| description | A description of the pool. | `string` | no |
| tenant | Tenant. `Default: admin` | 'string' | no |
| cloud | Cloud.` | 'string' | no |
| min\_servers\_up | Minimum number of servers in UP state for marking the pool UP. | `number` | no |
| min\_health\_monitors\_up | Minimum number of health monitors in UP state to mark server UP. Allowed in Basic edition, Essentials edition, Enterprise edition. | `number` | no |
| server\_timeout | Server timeout value specifies the time within which a server connection needs to be established and a request-response exchange completes between AVI and the server. Value of 0 results in using default timeout of 60 minutes. Allowed values are 0-3600000. Unit is MILLISECONDS. Default value when not specified in API or module is interpreted by ALB Controller as 0. | `number` | no |
| delete\_server\_on\_dns\_refresh | Indicates whether existing IPs are disabled(False) or deleted(True) on dns hostname refreshDetail -- On a dns refresh, some IPs set on pool may no longer be returned by the resolver. These IPs are deleted from the pool when this knob is set to True. They are disabled, if the knob is set to False. Allowed in Basic(Allowed values- True) edition, Essentials(Allowed values- True) edition, Enterprise edition. Default value when not specified in API or module is interpreted by ALB Controller as True. | `bool` | no |
| enable\_http2 | Enable HTTP/2 for traffic from VirtualService to all backend servers in this pool. Allowed in Basic(Allowed values- False) edition, Essentials(Allowed values- False) edition, Enterprise edition. Default value when not specified in API or module is interpreted by ALB Controller as False. | `bool` | no |
| ignore\_server\_port | Ignore the server port in building the load balancing state.Applicable only for consistent hash load balancing algorithm or Disable Port translation (use_service_port) use cases. Default value when not specified in API or module is interpreted by ALB Controller as False. | `bool` | no |
| routing\_pool | Enable to do routing when this pool is selected to send traffic. No servers present in routing pool. Default value when not specified in API or module is interpreted by ALB Controller as False. | `bool` | no |
| tier1\_lr | This tier1_lr field should be set same as VirtualService associated for NSX-T. | `string` | no |
| apic\_epg\_name | Synchronize Cisco APIC EPG members with pool servers. | `string` | no |
| health\_monitors | Verify server health by applying one or more health monitors. Active monitors generate synthetic traffic from each Service Engine and mark a server up or down based on the response. The Passive monitor listens only to client to server communication. It raises or lowers the ratio of traffic destined to a server based on successful responses. It is a reference to an object of type HealthMonitor. Maximum of 50 items allowed. | `list(string)` | no |
| gslb\_sp\_enabled | Indicates if the pool is a site-persistence pool. Allowed in Basic edition, Essentials edition, Enterprise edition. | `bool` | no |



## Outputs

| Name | Description |
|------|-------------|
| servers | Servers Config |

## How does consul-terraform-sync work?

There are 2 aspects of consul-terraform-sync.
1. **Updates from Consul catalog:**
In the backend, consul-terraform-sync creates a blocking API query session with the Consul agent indentified in the config to get updates from the Consul catalog.
consul-terraform-sync.
consul-terraform-sync will get an update for the services in the consul catalog when any of the following service attributes are created, updated or deleted. These updates include service creation and deletion as well.
   * service id
   * service name
   * service address
   * service port
   * service meta
   * service tags
   * service namespace
   * service health status
   * node id
   * node address
   * node datacenter
   * node tagged addresses
   * node meta

   
2. **Managing the entire Terraform workflow:**
If a task and is defined, one or more services are associated with the task, provider is declared in the task and a Terraform module is specified using the source field of the task, the following sequence of events will occur:
   1. consul-terraform-sync will install the required version of Terraform.
   2. consul-terraform-sync will install the required version of the Terraform provider defined in the config file and declared in the "task".
   3. A new direstory "nia-tasks" with a sub-directory corresponding to each "task" will be created. This is the reason for having strict guidelines around naming.
   4. Each sub-directory corresponds to a separate Terraform workspace. 
   5. Within each sub-directory corresponding a task, consul-terraform-sync will template a main.tf, variables.tf, terraform.tfvars and terraform.tfvars.tmpl.
      * **main.tf:**
         * This files contains declaration for the required terraform and provider versions based on the task definition. 
         * In addition, this file has the module (identified by the 'source' field in the task) declaration with the input variables
         * Consul K/V is used as the backend state for this Terraform workspace.
      
         Example of generated main.tf by consul-terraform-sync:
          ```terraform
         # This file is generated by Consul Terraform Sync.
            #
            # The HCL blocks, arguments, variables, and values are derived from the
            # operator configuration for Sync. Any manual changes to this file
            # may not be preserved and could be overwritten by a subsequent update.
            #
            # Task: avi-svc-web
            # Description: Automatically Scale AVI Service Redirection Destinations
            
            terraform {
              required_version = ">= 0.13.0, < 0.15"
              required_providers {
                avi = {
                  source  = "vmware/avi"
                  version = "20.1.4"
                }
              }
              backend "consul" {
                address = "x.x.x.x:8500"
                gzip    = true
                path    = "consul-terraform-sync/terraform"
              }
            }
            provider "avi" {
              avi_api_timeout = var.avi.avi_api_timeout
              avi_controller  = var.avi.avi_controller
              avi_password    = var.avi.avi_password
              avi_tenant      = var.avi.avi_tenant
              avi_username    = var.avi.avi_username
              avi_version     = var.avi.avi_version
            }
            
            module "avi-svc-web" {
              source   = "vmware/modules/nia/pool"
              services = var.services
            
              avi_controller = var.avi_controller
              avi_password   = var.avi_password
              avi_username   = var.avi_username
              avi_version    = var.avi_version
              lb_algorithm   = var.lb_algorithm
              pool_name      = var.pool_name
            }
          ```
      * **variables.tf:**
        * This is variables.tf file defined in the module
        
         Example of generated variables.tf by consul-terraform-sync:
         ```terraform
        # This file is generated by Consul Terraform Sync.
        #
        # The HCL blocks, arguments, variables, and values are derived from the
        # operator configuration for Sync. Any manual changes to this file
        # may not be preserved and could be overwritten by a subsequent update.
        #
        # Task: avi-svc-web
        # Description: Automatically Scale AVI Service Redirection Destinations
        
        # Service definition protocol v0
        variable "services" {
          description = "Consul services monitored by Consul Terraform Sync"
          type = map(
            object({
              id        = string
              name      = string
              kind      = string
              address   = string
              port      = number
              meta      = map(string)
              tags      = list(string)
              namespace = string
              status    = string
        
              node                  = string
              node_id               = string
              node_address          = string
              node_datacenter       = string
              node_tagged_addresses = map(string)
              node_meta             = map(string)
        
              cts_user_defined_meta = map(string)
            })
          )
        }
        variable "avi" {
          default     = null
          description = "Configuration object for avi"
          sensitive   = true
          type = object({
            alias           = string
            avi_api_timeout = number
            avi_controller  = string
            avi_password    = string
            avi_tenant      = string
            avi_username    = string
            avi_version     = string
          })
        }
         ```
        Example of generated variables.module.tf by consul-terraform-sync:
        ```terraform
        # This file is generated by Consul Terraform Sync.
        #
        # The HCL blocks, arguments, variables, and values are derived from the
        # operator configuration for Sync. Any manual changes to this file
        # may not be preserved and could be overwritten by a subsequent update.
        #
        # Task: avi-svc-web
        # Description: Automatically Scale AVI Service Redirection Destinations
        
        variable "avi_controller" {
          default = null
          type    = string
        }
        
        variable "avi_password" {
          default = null
          type    = string
        }
        
        variable "avi_username" {
          default = null
          type    = string
        }
        
        variable "avi_version" {
          default = null
          type    = string
        }
        
        variable "lb_algorithm" {
          default = null
          type    = string
        }
        
        variable "pool_name" {
          default = null
          type    = string
        }
        ```
      * **terraform.tfvars:**
         * This is the most important file generated by consul-terraform-sync.
         * This variables file is generated with the most updated values from Consul catalog for all the services identified in the task.
         * consul-terraform-sync updates this file with the latest values when the corresponding service gets updated in Consul catalog.
         
         Example of generated terraform.tfvars by consul-terraform-sync:
         ```terraform
        # This file is generated by Consul Terraform Sync.
        #
        # The HCL blocks, arguments, variables, and values are derived from the
        # operator configuration for Sync. Any manual changes to this file
        # may not be preserved and could be overwritten by a subsequent update.
        #
        # Task: avi-svc-web
        # Description: Automatically Scale AVI Service Redirection Destinations
        
        services = {
          "web.avi-dev.dc1" : {
            id      = "web"
            name    = "web"
            kind    = ""
            address = "x.x.x.x"
            port    = 80
            meta = {
              enabled = "false"
              ratio   = "5"
            }
            tags            = ["rails"]
            namespace       = null
            status          = "passing"
            node            = "ys-dev"
            node_id         = "<id>"
            node_address    = "x.x.x.x"
            node_datacenter = "dc1"
            node_tagged_addresses = {
              lan      = "x.x.x.x"
              lan_ipv4 = "x.x.x.x"
              wan      = "x.x.x.x"
              wan_ipv4 = "x.x.x.x"
            }
            node_meta = {
              consul-network-segment = ""
            }
            cts_user_defined_meta = {}
          },
          "web.mayank-dev.dc1" : {
            id              = "web"
            name            = "web"
            kind            = ""
            address         = "x.x.x.x"
            port            = 80
            meta            = {}
            tags            = ["rails"]
            namespace       = null
            status          = "passing"
            node            = "mayank-dev"
            node_id         = "<id>"
            node_address    = "x.x.x.x"
            node_datacenter = "dc1"
            node_tagged_addresses = {
              lan      = "x.x.x.x"
              lan_ipv4 = "x.x.x.x"
              wan      = "x.x.x.x"
              wan_ipv4 = "x.x.x.x"
            }
            node_meta = {
              consul-network-segment = ""
            }
            cts_user_defined_meta = {}
          },
          "web.ys-avi-dev-blr.dc1" : {
            id              = "web"
            name            = "web"
            kind            = ""
            address         = "x.x.x.x"
            port            = 89
            meta            = {}
            tags            = ["rails"]
            namespace       = null
            status          = "passing"
            node            = "ys-dev"
            node_id         = "<id>"
            node_address    = "x.x.x.x"
            node_datacenter = "dc1"
            node_tagged_addresses = {
              lan      = "x.x.x.x"
              lan_ipv4 = "x.x.x.x"
              wan      = "x.x.x.x"
              wan_ipv4 = "x.x.x.x"
            }
            node_meta = {
              consul-network-segment = ""
            }
            cts_user_defined_meta = {}
          }
        }
         ```
      * **Network Infrastructure Automation (NIA) compatible modules are built to utilize the above service variables**
    6. **consul-terraform-sync** manages the entire Terraform workflow for all the individual workspaces corresponding to the defined     "tasks" based on the updates from the services declared in those tasks.
    
  **In summary, consul-terraform-sync triggers a Terraform workflow based on updates it detects from Consul catalog.**
