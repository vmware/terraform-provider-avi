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

* `app_cookie_persistence_profile` - Specifies the application cookie persistence profile parameters.
* `description` - User defined description for the object.
* `hdr_persistence_profile` - Specifies the custom http header persistence profile parameters.
* `http_cookie_persistence_profile` - Specifies the http cookie persistence profile parameters.
* `ip_persistence_profile` - Specifies the client ip persistence profile parameters.
* `is_federated` - This field describes the object's replication scope.
* `name` - A user-friendly name for the persistence profile.
* `persistence_type` - Method used to persist clients to the same server for a duration of time or a session.
* `server_hm_down_recovery` - Specifies behavior when a persistent server has been marked down by a health monitor.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - Uuid of the persistence profile.

