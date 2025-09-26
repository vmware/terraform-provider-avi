// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceGslbCRMRuntimeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cluster_uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"events": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceEventInfoSchema(),
		},
		"fds_info": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceFdsInfoSchema(),
		},
		"local_info": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceLocalInfoSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"obj_uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"remote_info": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceRemoteInfoSchema(),
		},
		"replication_policy": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceReplicationPolicySchema(),
		},
		"site_name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"status_info": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceOperationalStatusSchema(),
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

func resourceAviGslbCRMRuntime() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviGslbCRMRuntimeCreate,
		Read:   ResourceAviGslbCRMRuntimeRead,
		Update: resourceAviGslbCRMRuntimeUpdate,
		Delete: resourceAviGslbCRMRuntimeDelete,
		Schema: ResourceGslbCRMRuntimeSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceGslbCRMRuntimeImporter,
		},
	}
}

func ResourceGslbCRMRuntimeImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceGslbCRMRuntimeSchema()
	return ResourceImporter(d, m, "gslbcrmruntime", s)
}

func ResourceAviGslbCRMRuntimeRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceGslbCRMRuntimeSchema()
	err := APIRead(d, meta, "gslbcrmruntime", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviGslbCRMRuntimeCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceGslbCRMRuntimeSchema()
	err := APICreate(d, meta, "gslbcrmruntime", s)
	if err == nil {
		err = ResourceAviGslbCRMRuntimeRead(d, meta)
	}
	return err
}

func resourceAviGslbCRMRuntimeUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceGslbCRMRuntimeSchema()
	var err error
	err = APIUpdate(d, meta, "gslbcrmruntime", s)
	if err == nil {
		err = ResourceAviGslbCRMRuntimeRead(d, meta)
	}
	return err
}

func resourceAviGslbCRMRuntimeDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "gslbcrmruntime")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
