// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceHTTPPolicySetSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cloud_config_cksum": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"created_by": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"geo_db_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"http_request_policy": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHTTPRequestPolicySchema(),
		},
		"http_response_policy": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHTTPResponsePolicySchema(),
		},
		"http_security_policy": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHTTPSecurityPolicySchema(),
		},
		"ip_reputation_db_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"is_internal_policy": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"markers": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceRoleFilterMatchLabelSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviHTTPPolicySet() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviHTTPPolicySetCreate,
		Read:   ResourceAviHTTPPolicySetRead,
		Update: resourceAviHTTPPolicySetUpdate,
		Delete: resourceAviHTTPPolicySetDelete,
		Schema: ResourceHTTPPolicySetSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceHTTPPolicySetImporter,
		},
	}
}

func ResourceHTTPPolicySetImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceHTTPPolicySetSchema()
	return ResourceImporter(d, m, "httppolicyset", s)
}

func ResourceAviHTTPPolicySetRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceHTTPPolicySetSchema()
	err := APIRead(d, meta, "httppolicyset", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviHTTPPolicySetCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceHTTPPolicySetSchema()
	err := APICreateOrUpdate(d, meta, "httppolicyset", s)
	if err == nil {
		err = ResourceAviHTTPPolicySetRead(d, meta)
	}
	return err
}

func resourceAviHTTPPolicySetUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceHTTPPolicySetSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "httppolicyset", s)
	if err == nil {
		err = ResourceAviHTTPPolicySetRead(d, meta)
	}
	return err
}

func resourceAviHTTPPolicySetDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "httppolicyset")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
