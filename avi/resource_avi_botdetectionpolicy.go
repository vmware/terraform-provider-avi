// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceBotDetectionPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"allow_list": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceBotAllowListSchema(),
		},
		"client_behavior_detector": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceBotConfigClientBehaviorSchema(),
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
	err := APICreate(d, meta, "botdetectionpolicy", s)
	if err == nil {
		err = ResourceAviBotDetectionPolicyRead(d, meta)
	}
	return err
}

func resourceAviBotDetectionPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceBotDetectionPolicySchema()
	var err error
	err = APIUpdate(d, meta, "botdetectionpolicy", s)
	if err == nil {
		err = ResourceAviBotDetectionPolicyRead(d, meta)
	}
	return err
}

func resourceAviBotDetectionPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "botdetectionpolicy")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
