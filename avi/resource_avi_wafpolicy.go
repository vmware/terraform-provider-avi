// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"
	"strings"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceWafPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"allow_mode_delegation": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"created_by": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"crs_groups": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceWafRuleGroupSchema(),
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"enable_app_learning": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"failure_mode": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "WAF_FAILURE_MODE_OPEN",
		},
		"mode": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "WAF_MODE_DETECTION_ONLY",
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"paranoia_level": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "WAF_PARANOIA_LEVEL_LOW",
		},
		"positive_security_model": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceWafPositiveSecurityModelSchema(),
		},
		"post_crs_groups": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceWafRuleGroupSchema(),
		},
		"pre_crs_groups": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceWafRuleGroupSchema(),
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
		"waf_crs_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"waf_profile_ref": {
			Type:     schema.TypeString,
			Required: true,
		},
		"whitelist": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceWafPolicyWhitelistSchema(),
		},
	}
}

func resourceAviWafPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviWafPolicyCreate,
		Read:   ResourceAviWafPolicyRead,
		Update: resourceAviWafPolicyUpdate,
		Delete: resourceAviWafPolicyDelete,
		Schema: ResourceWafPolicySchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceWafPolicyImporter,
		},
	}
}

func ResourceWafPolicyImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceWafPolicySchema()
	return ResourceImporter(d, m, "wafpolicy", s)
}

func ResourceAviWafPolicyRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWafPolicySchema()
	err := APIRead(d, meta, "wafpolicy", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviWafPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWafPolicySchema()
	err := APICreateOrUpdate(d, meta, "wafpolicy", s)
	if err == nil {
		err = ResourceAviWafPolicyRead(d, meta)
	}
	return err
}

func resourceAviWafPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWafPolicySchema()
	var err error
	err = APICreateOrUpdate(d, meta, "wafpolicy", s)
	if err == nil {
		err = ResourceAviWafPolicyRead(d, meta)
	}
	return err
}

func resourceAviWafPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "wafpolicy"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviWafPolicyDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
