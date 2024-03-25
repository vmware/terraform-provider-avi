// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourcePoolGroupDeploymentPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"auto_disable_old_prod_pools": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
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
		"evaluation_duration": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "300",
			ValidateFunc: validateInteger,
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
		"rules": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourcePGDeploymentRuleSchema(),
		},
		"scheme": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "BLUE_GREEN",
		},
		"target_test_traffic_ratio": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "100",
			ValidateFunc: validateInteger,
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"test_traffic_ratio_rampup": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "100",
			ValidateFunc: validateInteger,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"webhook_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviPoolGroupDeploymentPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviPoolGroupDeploymentPolicyCreate,
		Read:   ResourceAviPoolGroupDeploymentPolicyRead,
		Update: resourceAviPoolGroupDeploymentPolicyUpdate,
		Delete: resourceAviPoolGroupDeploymentPolicyDelete,
		Schema: ResourcePoolGroupDeploymentPolicySchema(),
		Importer: &schema.ResourceImporter{
			State: ResourcePoolGroupDeploymentPolicyImporter,
		},
	}
}

func ResourcePoolGroupDeploymentPolicyImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourcePoolGroupDeploymentPolicySchema()
	return ResourceImporter(d, m, "poolgroupdeploymentpolicy", s)
}

func ResourceAviPoolGroupDeploymentPolicyRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePoolGroupDeploymentPolicySchema()
	err := APIRead(d, meta, "poolgroupdeploymentpolicy", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviPoolGroupDeploymentPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePoolGroupDeploymentPolicySchema()
	err := APICreate(d, meta, "poolgroupdeploymentpolicy", s)
	if err == nil {
		err = ResourceAviPoolGroupDeploymentPolicyRead(d, meta)
	}
	return err
}

func resourceAviPoolGroupDeploymentPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePoolGroupDeploymentPolicySchema()
	var err error
	err = APIUpdate(d, meta, "poolgroupdeploymentpolicy", s)
	if err == nil {
		err = ResourceAviPoolGroupDeploymentPolicyRead(d, meta)
	}
	return err
}

func resourceAviPoolGroupDeploymentPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "poolgroupdeploymentpolicy")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
