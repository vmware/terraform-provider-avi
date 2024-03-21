// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceGslbSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"async_interval": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateInteger,
		},
		"clear_on_max_retries": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "20",
			ValidateFunc: validateInteger,
		},
		"client_ip_addr_group": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceGslbClientIpAddrGroupSchema(),
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
		"dns_configs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceDNSConfigSchema(),
		},
		"enable_config_by_members": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"error_resync_interval": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "300",
			ValidateFunc: validateInteger,
		},
		"fileobject_max_file_versions": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "3",
			ValidateFunc: validateInteger,
		},
		"is_federated": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"leader_cluster_uuid": {
			Type:     schema.TypeString,
			Required: true,
		},
		"maintenance_mode": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"replication_policy": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceReplicationPolicySchema(),
		},
		"send_interval": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "15",
			ValidateFunc: validateInteger,
		},
		"send_interval_prior_to_maintenance_mode": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateInteger,
		},
		"sites": {
			Type:     schema.TypeList,
			Required: true,
			Elem:     ResourceGslbSiteSchema(),
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"tenant_scoped": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"third_party_sites": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceGslbThirdPartySiteSchema(),
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"view_id": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateInteger,
		},
	}
}

func resourceAviGslb() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviGslbCreate,
		Read:   ResourceAviGslbRead,
		Update: resourceAviGslbUpdate,
		Delete: resourceAviGslbDelete,
		Schema: ResourceGslbSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceGslbImporter,
		},
	}
}

func ResourceGslbImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceGslbSchema()
	return ResourceImporter(d, m, "gslb", s)
}

func ResourceAviGslbRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceGslbSchema()
	err := APIRead(d, meta, "gslb", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviGslbCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceGslbSchema()
	err := APICreateOrUpdate(d, meta, "gslb", s)
	if err == nil {
		err = ResourceAviGslbRead(d, meta)
	}
	return err
}

func resourceAviGslbUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceGslbSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "gslb", s)
	if err == nil {
		err = ResourceAviGslbRead(d, meta)
	}
	return err
}

func resourceAviGslbDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "gslb")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
