// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
)

func ResourceWafPolicyPSMGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"enable": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"hit_action": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "WAF_ACTION_ALLOW_PARAMETER",
		},
		"is_learning_group": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"locations": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceWafPSMLocationSchema(),
		},
		"markers": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceRoleFilterMatchLabelSchema(),
		},
		"miss_action": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "WAF_ACTION_NO_OP",
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

func resourceAviWafPolicyPSMGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviWafPolicyPSMGroupCreate,
		Read:   ResourceAviWafPolicyPSMGroupRead,
		Update: resourceAviWafPolicyPSMGroupUpdate,
		Delete: resourceAviWafPolicyPSMGroupDelete,
		Schema: ResourceWafPolicyPSMGroupSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceWafPolicyPSMGroupImporter,
		},
	}
}

func ResourceWafPolicyPSMGroupImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceWafPolicyPSMGroupSchema()
	return ResourceImporter(d, m, "wafpolicypsmgroup", s)
}

func ResourceAviWafPolicyPSMGroupRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWafPolicyPSMGroupSchema()
	err := APIRead(d, meta, "wafpolicypsmgroup", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviWafPolicyPSMGroupCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWafPolicyPSMGroupSchema()
	err := APICreateOrUpdate(d, meta, "wafpolicypsmgroup", s)
	if err == nil {
		err = ResourceAviWafPolicyPSMGroupRead(d, meta)
	}
	return err
}

func resourceAviWafPolicyPSMGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWafPolicyPSMGroupSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "wafpolicypsmgroup", s)
	if err == nil {
		err = ResourceAviWafPolicyPSMGroupRead(d, meta)
	}
	return err
}

func resourceAviWafPolicyPSMGroupDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "wafpolicypsmgroup"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviWafPolicyPSMGroupDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
