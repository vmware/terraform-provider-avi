<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_apipolicy"
sidebar_current: "docs-avi-datasource-apipolicy"
description: |-
  Get information of Avi ApiPolicy.
---

# avi_apipolicy

This data source is used to to get avi_apipolicy objects.

## Example Usage

```hcl
data "avi_apipolicy" "foo_apipolicy" {
    uuid = "apipolicy-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search ApiPolicy by name.
* `uuid` - (Optional) Search ApiPolicy by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `active_api_labels` - List of labels applied to active api endpoints. An active api is an endpoint whose type is api_active. Endpoints defined in the policy are active by default. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `api_spec_info` - Api specification metadata extracted from the associated openapi specification. Automatically populated when a fileobject is associated with this policy. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `description` - Description of this api policy. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `file_object_refs` - Reference to the uploaded openapi specification file associated with this policy. Only one file is supported at a time. It is a reference to an object of type fileobject. Field introduced in 32.1.4. Maximum of 1 items allowed. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `label_mappings` - Mapping of labels to api policy actions. Field introduced in 32.1.4. Maximum of 256 items allowed. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `name` - Name of this object, unique per tenant. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `non_api_url_labels` - List of labels applied to non-api url requests. Non-api urls are methods and urls that are outside the scope of the policy. These are usually used to retrieve static information that are not tied to back-end business logic. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `orphan_api_classification_settings` - Orphan api classification settings for this api policy. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `orphan_api_labels` - List of labels applied to orphan api endpoints. An orphan api is an endpoint that is specified in the api-spec but has not been seen in the datapath for a predefined duration. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `path_refs` - List of path specifications. When an oas fileobject is associated to this apipolicy, the paths defined in the oas fileobject will be automatically added to this list. If oas fileobject has a path that is already defined in the list, the existing path in the list will be updated as per the oas fileobject. It is a reference to an object of type apipath. Field introduced in 32.1.4. Maximum of 2000 items allowed. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `routing_info` - Optional header-based routing configuration for evh child vs selection. When set, the rules inside are used in addition to server fqdns (host match) and server_info.path_prefix (path match) to determine which child vs handles a request. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `server_info` - Server list defining the scope of this api policy. Requests not matching any server url are treated as non-api traffic. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `shadow_api_labels` - List of labels applied to shadow api endpoints. A shadow api is an endpoint that is not specified in the api-spec but is inside the scope of this policy (matching the server url and path prefix) and is seen in the datapath. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `uuid` - The object uuid. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `validation_settings` - Validation settings for this api policy. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `zombie_api_classification_settings` - Zombie api classification settings for this api policy. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `zombie_api_labels` - List of labels applied to zombie api endpoints. A zombie api is an endpoint that is specified in the api-spec but is seen in the datapath only as drip-traffic over a predefined duration. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.

