// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceLicenseLedgerDetailsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"escrow_infos": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceLicenseInfoSchema(),
		},
		"se_infos": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceLicenseInfoSchema(),
		},
		"tier_usages": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceLicenseTierUsageSchema(),
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviLicenseLedgerDetails() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviLicenseLedgerDetailsCreate,
		Read:   ResourceAviLicenseLedgerDetailsRead,
		Update: resourceAviLicenseLedgerDetailsUpdate,
		Delete: resourceAviLicenseLedgerDetailsDelete,
		Schema: ResourceLicenseLedgerDetailsSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceLicenseLedgerDetailsImporter,
		},
	}
}

func ResourceLicenseLedgerDetailsImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceLicenseLedgerDetailsSchema()
	return ResourceImporter(d, m, "licenseledgerdetails", s)
}

func ResourceAviLicenseLedgerDetailsRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceLicenseLedgerDetailsSchema()
	err := APIRead(d, meta, "licenseledgerdetails", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviLicenseLedgerDetailsCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceLicenseLedgerDetailsSchema()
	err := APICreate(d, meta, "licenseledgerdetails", s)
	if err == nil {
		err = ResourceAviLicenseLedgerDetailsRead(d, meta)
	}
	return err
}

func resourceAviLicenseLedgerDetailsUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceLicenseLedgerDetailsSchema()
	var err error
	err = APIUpdate(d, meta, "licenseledgerdetails", s)
	if err == nil {
		err = ResourceAviLicenseLedgerDetailsRead(d, meta)
	}
	return err
}

func resourceAviLicenseLedgerDetailsDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "licenseledgerdetails")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
