<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_applicationpersistenceprofile"
sidebar_current: "docs-avi-datasource-applicationpersistenceprofile"
description: |-
  Get information of Avi ApplicationPersistenceProfile.
---

# avi_applicationpersistenceprofile

This data source is used to to get avi_applicationpersistenceprofile objects.

## Example Usage

```hcl
data "avi_applicationpersistenceprofile" "foo_applicationpersistenceprofile" {
    uuid = "applicationpersistenceprofile-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search ApplicationPersistenceProfile by name.
* `uuid` - (Optional) Search ApplicationPersistenceProfile by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `app_cookie_persistence_profile` - Specifies the application cookie persistence profile parameters. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `description` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `hdr_persistence_profile` - Specifies the custom http header persistence profile parameters. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `http_cookie_persistence_profile` - Specifies the http cookie persistence profile parameters. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `ip_persistence_profile` - Specifies the client ip persistence profile parameters. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `is_federated` - This field describes the object's replication scope. If the field is set to false, then the object is visible within the controller-cluster and its associated service-engines. If the field is set to true, then the object is replicated across the federation. Field introduced in 17.1.3. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `markers` - List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `name` - A user-friendly name for the persistence profile. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `persistence_type` - Method used to persist clients to the same server for a duration of time or a session. Enum options - PERSISTENCE_TYPE_CLIENT_IP_ADDRESS, PERSISTENCE_TYPE_HTTP_COOKIE, PERSISTENCE_TYPE_TLS, PERSISTENCE_TYPE_CLIENT_IPV6_ADDRESS, PERSISTENCE_TYPE_CUSTOM_HTTP_HEADER, PERSISTENCE_TYPE_APP_COOKIE, PERSISTENCE_TYPE_GSLB_SITE. Allowed in enterprise edition with any value, essentials edition(allowed values- persistence_type_client_ip_address,persistence_type_http_cookie), basic edition(allowed values- persistence_type_client_ip_address,persistence_type_http_cookie), enterprise with cloud services edition.
* `server_hm_down_recovery` - Specifies behavior when a persistent server has been marked down by a health monitor. Enum options - HM_DOWN_PICK_NEW_SERVER, HM_DOWN_ABORT_CONNECTION, HM_DOWN_CONTINUE_PERSISTENT_SERVER. Allowed in enterprise edition with any value, essentials edition(allowed values- hm_down_pick_new_server), basic edition(allowed values- hm_down_pick_new_server), enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `uuid` - Uuid of the persistence profile. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

