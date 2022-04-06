<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_applicationpersistenceprofile"
sidebar_current: "docs-avi-resource-applicationpersistenceprofile"
description: |-
  Creates and manages Avi ApplicationPersistenceProfile.
---

# avi_applicationpersistenceprofile

The ApplicationPersistenceProfile resource allows the creation and management of Avi ApplicationPersistenceProfile

## Example Usage

```hcl
resource "avi_applicationpersistenceprofile" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) A user-friendly name for the persistence profile. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `persistence_type` - (Required) Method used to persist clients to the same server for a duration of time or a session. Enum options - PERSISTENCE_TYPE_CLIENT_IP_ADDRESS, PERSISTENCE_TYPE_HTTP_COOKIE, PERSISTENCE_TYPE_TLS, PERSISTENCE_TYPE_CLIENT_IPV6_ADDRESS, PERSISTENCE_TYPE_CUSTOM_HTTP_HEADER, PERSISTENCE_TYPE_APP_COOKIE, PERSISTENCE_TYPE_GSLB_SITE. Allowed in enterprise with any value edition, essentials(allowed values- persistence_type_client_ip_address,persistence_type_http_cookie) edition, basic(allowed values- persistence_type_client_ip_address,persistence_type_http_cookie) edition, enterprise with cloud services edition.
* `app_cookie_persistence_profile` - (Optional) Specifies the application cookie persistence profile parameters. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `description` - (Optional) Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `hdr_persistence_profile` - (Optional) Specifies the custom http header persistence profile parameters. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `http_cookie_persistence_profile` - (Optional) Specifies the http cookie persistence profile parameters. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `ip_persistence_profile` - (Optional) Specifies the client ip persistence profile parameters. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `is_federated` - (Optional) This field describes the object's replication scope. If the field is set to false, then the object is visible within the controller-cluster and its associated service-engines. If the field is set to true, then the object is replicated across the federation. Field introduced in 17.1.3. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `markers` - (Optional) List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `server_hm_down_recovery` - (Optional) Specifies behavior when a persistent server has been marked down by a health monitor. Enum options - HM_DOWN_PICK_NEW_SERVER, HM_DOWN_ABORT_CONNECTION, HM_DOWN_CONTINUE_PERSISTENT_SERVER. Allowed in enterprise with any value edition, essentials(allowed values- hm_down_pick_new_server) edition, basic(allowed values- hm_down_pick_new_server) edition, enterprise with cloud services edition.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the persistence profile. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.

