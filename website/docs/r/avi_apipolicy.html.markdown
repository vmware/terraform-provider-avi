<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_apipolicy"
sidebar_current: "docs-avi-resource-apipolicy"
description: |-
  Creates and manages Avi ApiPolicy.
---

# avi_apipolicy

The ApiPolicy resource allows the creation and management of Avi ApiPolicy

## Example Usage

```hcl
resource "avi_apipolicy" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of this object, unique per tenant. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `active_api_labels` - (Optional) List of labels applied to active api endpoints. An active api is an endpoint whose type is api_active. Endpoints defined in the policy are active by default. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `api_spec_info` - (Optional) Api specification metadata extracted from the associated openapi specification. Automatically populated when a fileobject is associated with this policy. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `description` - (Optional) Description of this api policy. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `file_object_refs` - (Optional) Reference to the uploaded openapi specification file associated with this policy. Only one file is supported at a time. It is a reference to an object of type fileobject. Field introduced in 32.1.4. Maximum of 1 items allowed. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `label_mappings` - (Optional) Mapping of labels to api policy actions. Field introduced in 32.1.4. Maximum of 256 items allowed. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `non_api_url_labels` - (Optional) List of labels applied to non-api url requests. Non-api urls are methods and urls that are outside the scope of the policy. These are usually used to retrieve static information that are not tied to back-end business logic. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `orphan_api_classification_settings` - (Optional) Orphan api classification settings for this api policy. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `orphan_api_labels` - (Optional) List of labels applied to orphan api endpoints. An orphan api is an endpoint that is specified in the api-spec but has not been seen in the datapath for a predefined duration. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `path_refs` - (Optional) List of path specifications. When an oas fileobject is associated to this apipolicy, the paths defined in the oas fileobject will be automatically added to this list. If oas fileobject has a path that is already defined in the list, the existing path in the list will be updated as per the oas fileobject. It is a reference to an object of type apipath. Field introduced in 32.1.4. Maximum of 2000 items allowed. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `routing_info` - (Optional) Optional header-based routing configuration for evh child vs selection. When set, the rules inside are used in addition to server fqdns (host match) and server_info.path_prefix (path match) to determine which child vs handles a request. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `server_info` - (Optional) Server list defining the scope of this api policy. Requests not matching any server url are treated as non-api traffic. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `shadow_api_labels` - (Optional) List of labels applied to shadow api endpoints. A shadow api is an endpoint that is not specified in the api-spec but is inside the scope of this policy (matching the server url and path prefix) and is seen in the datapath. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `tenant_ref` - (Optional) It is a reference to an object of type tenant. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `validation_settings` - (Optional) Validation settings for this api policy. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `zombie_api_classification_settings` - (Optional) Zombie api classification settings for this api policy. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `zombie_api_labels` - (Optional) List of labels applied to zombie api endpoints. A zombie api is an endpoint that is specified in the api-spec but is seen in the datapath only as drip-traffic over a predefined duration. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  The object uuid. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.

