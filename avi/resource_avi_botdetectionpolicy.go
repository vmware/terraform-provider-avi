// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
)

func ResourceBotDetectionPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"allow_list": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceBotAllowListSchema(),
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"ip_location_detector": {
			Type:     schema.TypeSet,
			Required: true,
			Elem:     ResourceBotConfigIPLocationSchema(),
		},
		"ip_reputation_detector": {
			Type:     schema.TypeSet,
			Required: true,
			Elem:     ResourceBotConfigIPReputationSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"system_bot_mapping_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"system_consolidator_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"user_agent_detector": {
			Type:     schema.TypeSet,
			Required: true,
			Elem:     ResourceBotConfigUserAgentSchema(),
		},
		"user_bot_mapping_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"user_consolidator_ref": {
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

func resourceAviBotDetectionPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviBotDetectionPolicyCreate,
		Read:   ResourceAviBotDetectionPolicyRead,
		Update: resourceAviBotDetectionPolicyUpdate,
		Delete: resourceAviBotDetectionPolicyDelete,
		Schema: ResourceBotDetectionPolicySchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceBotDetectionPolicyImporter,
		},
	}
}

func ResourceBotDetectionPolicyImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceBotDetectionPolicySchema()
	return ResourceImporter(d, m, "botdetectionpolicy", s)
}

func ResourceAviBotDetectionPolicyRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceBotDetectionPolicySchema()
	err := APIRead(d, meta, "botdetectionpolicy", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviBotDetectionPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceBotDetectionPolicySchema()
	err := APICreateOrUpdate(d, meta, "botdetectionpolicy", s)
	if err == nil {
		err = ResourceAviBotDetectionPolicyRead(d, meta)
	}
	return err
}

func resourceAviBotDetectionPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceBotDetectionPolicySchema()
	var err error
	err = APICreateOrUpdate(d, meta, "botdetectionpolicy", s)
	if err == nil {
		err = ResourceAviBotDetectionPolicyRead(d, meta)
	}
	return err
}

func resourceAviBotDetectionPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "botdetectionpolicy"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviBotDetectionPolicyDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
