// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
)

func ResourceSecurityPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
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
	err := APICreateOrUpdate(d, meta, "securitypolicy", s)
	if err == nil {
		err = ResourceAviSecurityPolicyRead(d, meta)
	}
	return err
}

func resourceAviSecurityPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSecurityPolicySchema()
	var err error
	err = APICreateOrUpdate(d, meta, "securitypolicy", s)
	if err == nil {
		err = ResourceAviSecurityPolicyRead(d, meta)
	}
	return err
}

func resourceAviSecurityPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "securitypolicy"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviSecurityPolicyDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
