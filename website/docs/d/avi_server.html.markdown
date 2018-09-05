---
layout: "avi"
page_title: "AVI: avi_server"
sidebar_current: "docs-avi-datasource-server"
description: |-
Get information of Avi Server.
---

# avi_server

This data source is used to to get avi_server objects.

## Example Usage

```hcl
data "avi_server" "foo_Server" {
    uuid = "server-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    pool_ref = "/api/pool/pool-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    ip='10.0.0.3'
  }
```

## Argument Reference

* `pool_ref` - (Optional) Search Server by pool_ref.
* `uuid` - (Optional) Search Server by uuid.
* `ip` - (Optional) Search Server by ip.
  
## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` - (Optional) Search server object by uuid.
* `pool_ref` - The pool is an object that contains destination servers and related attributes such as load-balancing and persistence.
* `ip` - IP address of a destination servers. 
* `port` - Port of a destination servers.
* `type` - Type of ip address (V4)
* `autoscaling_group_name` - Name of autoscaling group this server belongs to.
* `description` - A description of the server.
* `enabled` - Enable or disable the server.
* `external_orchestration_id` - UID of server in external orchestration systems.
* `external_uuid` - UUID identifying VM in OpenStack and other external compute.
* `hostname` - DNS resolvable name of the server. May be used in place of the IP address.
* `location` - Geographic location of the server.Currently only for internal usage.
* `nw_ref` - This field is used internally by Avi, not editable by the user. It is a reference to an object of type VIMgrNWRuntime. 
* `prst_hdr_val` - Header value for custom header persistence. 
* `rewrite_host_header` - Rewrite incoming Host Header to server name.
* `vm_ref` - This field is used internally by Avi, not editable by the user. It is a reference to an object of type VIMgrVMRuntime.
