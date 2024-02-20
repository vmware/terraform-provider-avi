// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceSecurityPolicySchema() map[string]*schema.Schema {
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
		"dns_amplification_denyports": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourcePortMatchGenericSchema(),
		},
		"dns_attacks": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceDnsAttacksSchema(),
		},
		"dns_policy_index": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
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
		"network_security_policy_index": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateInteger,
		},
		"oper_mode": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "DETECTION",
		},
		"tcp_attacks": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceTcpAttacksSchema(),
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"udp_attacks": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceUdpAttacksSchema(),
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviSecurityPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviSecurityPolicyCreate,
		Read:   ResourceAviSecurityPolicyRead,
		Update: resourceAviSecurityPolicyUpdate,
		Delete: resourceAviSecurityPolicyDelete,
		Schema: ResourceSecurityPolicySchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceSecurityPolicyImporter,
		},
	}
}

func ResourceSecurityPolicyImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceSecurityPolicySchema()
	return ResourceImporter(d, m, "securitypolicy", s)
}

func ResourceAviSecurityPolicyRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSecurityPolicySchema()
	err := APIRead(d, meta, "securitypolicy", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviSecurityPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSecurityPolicySchema()
	err := APICreate(d, meta, "securitypolicy", s)
	if err == nil {
		err = ResourceAviSecurityPolicyRead(d, meta)
	}
	return err
}

func resourceAviSecurityPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSecurityPolicySchema()
	var err error
	err = APIUpdate(d, meta, "securitypolicy", s)
	if err == nil {
		err = ResourceAviSecurityPolicyRead(d, meta)
	}
	return err
}

func resourceAviSecurityPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "securitypolicy")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
