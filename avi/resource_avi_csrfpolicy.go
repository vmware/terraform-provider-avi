// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
)

func ResourceCSRFPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"cookie_name": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "X-CSRF-TOKEN",
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"rules": {
			Type:     schema.TypeList,
			Required: true,
			Elem:     ResourceCSRFRuleSchema(),
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"token_validity_time_min": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "360",
			ValidateFunc: validateInteger,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviCSRFPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviCSRFPolicyCreate,
		Read:   ResourceAviCSRFPolicyRead,
		Update: resourceAviCSRFPolicyUpdate,
		Delete: resourceAviCSRFPolicyDelete,
		Schema: ResourceCSRFPolicySchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceCSRFPolicyImporter,
		},
	}
}

func ResourceCSRFPolicyImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceCSRFPolicySchema()
	return ResourceImporter(d, m, "csrfpolicy", s)
}

func ResourceAviCSRFPolicyRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCSRFPolicySchema()
	err := APIRead(d, meta, "csrfpolicy", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviCSRFPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCSRFPolicySchema()
	err := APICreateOrUpdate(d, meta, "csrfpolicy", s)
	if err == nil {
		err = ResourceAviCSRFPolicyRead(d, meta)
	}
	return err
}

func resourceAviCSRFPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCSRFPolicySchema()
	var err error
	err = APICreateOrUpdate(d, meta, "csrfpolicy", s)
	if err == nil {
		err = ResourceAviCSRFPolicyRead(d, meta)
	}
	return err
}

func resourceAviCSRFPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "csrfpolicy"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviCSRFPolicyDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
