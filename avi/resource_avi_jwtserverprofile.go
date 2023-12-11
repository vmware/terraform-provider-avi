// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceJWTServerProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"controller_internal_auth": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceControllerInternalAuthSchema(),
		},
		"is_federated": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"issuer": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"jwks_keys": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"jwt_profile_type": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "CLIENT_AUTH",
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

func resourceAviJWTServerProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviJWTServerProfileCreate,
		Read:   ResourceAviJWTServerProfileRead,
		Update: resourceAviJWTServerProfileUpdate,
		Delete: resourceAviJWTServerProfileDelete,
		Schema: ResourceJWTServerProfileSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceJWTServerProfileImporter,
		},
	}
}

func ResourceJWTServerProfileImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceJWTServerProfileSchema()
	return ResourceImporter(d, m, "jwtserverprofile", s)
}

func ResourceAviJWTServerProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceJWTServerProfileSchema()
	err := APIRead(d, meta, "jwtserverprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviJWTServerProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceJWTServerProfileSchema()
	err := APICreateOrUpdate(d, meta, "jwtserverprofile", s)
	if err == nil {
		err = ResourceAviJWTServerProfileRead(d, meta)
	}
	return err
}

func resourceAviJWTServerProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceJWTServerProfileSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "jwtserverprofile", s)
	if err == nil {
		err = ResourceAviJWTServerProfileRead(d, meta)
	}
	return err
}

func resourceAviJWTServerProfileDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "jwtserverprofile")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
