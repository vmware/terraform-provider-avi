// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceCertJwtStoreSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"jwt": {
			Type:     schema.TypeString,
			Required: true,
		},
		"key": {
			Type:             schema.TypeString,
			Required:         true,
			Sensitive:        true,
			DiffSuppressFunc: suppressSensitiveFieldDiffs,
		},
		"key_passphrase": {
			Type:             schema.TypeString,
			Optional:         true,
			Computed:         true,
			Sensitive:        true,
			DiffSuppressFunc: suppressSensitiveFieldDiffs,
		},
		"kid": {
			Type:     schema.TypeString,
			Required: true,
		},
		"last_rotated_at": {
			Type:     schema.TypeSet,
			Required: true,
			Elem:     ResourceTimeStampSchema(),
		},
		"public_key_algorithm": {
			Type:     schema.TypeString,
			Required: true,
		},
		"type": {
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

func resourceAviCertJwtStore() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviCertJwtStoreCreate,
		Read:   ResourceAviCertJwtStoreRead,
		Update: resourceAviCertJwtStoreUpdate,
		Delete: resourceAviCertJwtStoreDelete,
		Schema: ResourceCertJwtStoreSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceCertJwtStoreImporter,
		},
	}
}

func ResourceCertJwtStoreImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceCertJwtStoreSchema()
	return ResourceImporter(d, m, "certjwtstore", s)
}

func ResourceAviCertJwtStoreRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCertJwtStoreSchema()
	err := APIRead(d, meta, "certjwtstore", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviCertJwtStoreCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCertJwtStoreSchema()
	err := APICreate(d, meta, "certjwtstore", s)
	if err == nil {
		err = ResourceAviCertJwtStoreRead(d, meta)
	}
	return err
}

func resourceAviCertJwtStoreUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCertJwtStoreSchema()
	var err error
	err = APIUpdate(d, meta, "certjwtstore", s)
	if err == nil {
		err = ResourceAviCertJwtStoreRead(d, meta)
	}
	return err
}

func resourceAviCertJwtStoreDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "certjwtstore")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
