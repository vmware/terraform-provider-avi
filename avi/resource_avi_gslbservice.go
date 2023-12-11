// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceGslbServiceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"application_persistence_profile_ref": {
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
		"controller_health_status_enabled": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
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
		"domain_names": {
			Type:     schema.TypeList,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"down_response": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceGslbServiceDownResponseSchema(),
		},
		"enabled": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"groups": {
			Type:     schema.TypeList,
			Required: true,
			Elem:     ResourceGslbPoolSchema(),
		},
		"health_monitor_refs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"health_monitor_scope": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "GSLB_SERVICE_HEALTH_MONITOR_ALL_MEMBERS",
		},
		"hm_off": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateBool,
		},
		"is_federated": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"markers": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceRoleFilterMatchLabelSchema(),
		},
		"min_members": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateInteger,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"num_dns_ip": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateInteger,
		},
		"pki_profile_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"pool_algorithm": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "GSLB_SERVICE_ALGORITHM_PRIORITY",
		},
		"resolve_cname": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"site_persistence_enabled": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"topology_policy_enabled": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"ttl": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateInteger,
		},
		"use_edns_client_subnet": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"wildcard_match": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
	}
}

func resourceAviGslbService() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviGslbServiceCreate,
		Read:   ResourceAviGslbServiceRead,
		Update: resourceAviGslbServiceUpdate,
		Delete: resourceAviGslbServiceDelete,
		Schema: ResourceGslbServiceSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceGslbServiceImporter,
		},
	}
}

func ResourceGslbServiceImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceGslbServiceSchema()
	return ResourceImporter(d, m, "gslbservice", s)
}

func ResourceAviGslbServiceRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceGslbServiceSchema()
	err := APIRead(d, meta, "gslbservice", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviGslbServiceCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceGslbServiceSchema()
	err := APICreateOrUpdate(d, meta, "gslbservice", s)
	if err == nil {
		err = ResourceAviGslbServiceRead(d, meta)
	}
	return err
}

func resourceAviGslbServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceGslbServiceSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "gslbservice", s)
	if err == nil {
		err = ResourceAviGslbServiceRead(d, meta)
	}
	return err
}

func resourceAviGslbServiceDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "gslbservice")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
