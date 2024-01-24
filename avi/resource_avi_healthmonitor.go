// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceHealthMonitorSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"allow_duplicate_monitors": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateBool,
		},
		"authentication": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHealthMonitorAuthInfoSchema(),
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
		"disable_quickstart": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateBool,
		},
		"dns_monitor": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHealthMonitorDNSSchema(),
		},
		"external_monitor": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHealthMonitorExternalSchema(),
		},
		"failed_checks": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "2",
			ValidateFunc: validateInteger,
		},
		"ftp_monitor": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHealthMonitorFtpSchema(),
		},
		"ftps_monitor": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHealthMonitorFtpSchema(),
		},
		"http_monitor": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHealthMonitorHttpSchema(),
		},
		"https_monitor": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHealthMonitorHttpSchema(),
		},
		"imap_monitor": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHealthMonitorImapSchema(),
		},
		"imaps_monitor": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHealthMonitorImapSchema(),
		},
		"is_federated": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"ldap_monitor": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHealthMonitorLdapSchema(),
		},
		"ldaps_monitor": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHealthMonitorLdapSchema(),
		},
		"markers": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceRoleFilterMatchLabelSchema(),
		},
		"monitor_port": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateInteger,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"pop3_monitor": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHealthMonitorPop3Schema(),
		},
		"pop3s_monitor": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHealthMonitorPop3Schema(),
		},
		"radius_monitor": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHealthMonitorRadiusSchema(),
		},
		"receive_timeout": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "4",
			ValidateFunc: validateInteger,
		},
		"sctp_monitor": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHealthMonitorSctpSchema(),
		},
		"send_interval": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "10",
			ValidateFunc: validateInteger,
		},
		"sip_monitor": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHealthMonitorSIPSchema(),
		},
		"smtp_monitor": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHealthMonitorSmtpSchema(),
		},
		"smtps_monitor": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHealthMonitorSmtpSchema(),
		},
		"successful_checks": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "2",
			ValidateFunc: validateInteger,
		},
		"tcp_monitor": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHealthMonitorTcpSchema(),
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"type": {
			Type:     schema.TypeString,
			Required: true,
		},
		"udp_monitor": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHealthMonitorUdpSchema(),
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviHealthMonitor() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviHealthMonitorCreate,
		Read:   ResourceAviHealthMonitorRead,
		Update: resourceAviHealthMonitorUpdate,
		Delete: resourceAviHealthMonitorDelete,
		Schema: ResourceHealthMonitorSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceHealthMonitorImporter,
		},
	}
}

func ResourceHealthMonitorImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceHealthMonitorSchema()
	return ResourceImporter(d, m, "healthmonitor", s)
}

func ResourceAviHealthMonitorRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceHealthMonitorSchema()
	err := APIRead(d, meta, "healthmonitor", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviHealthMonitorCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceHealthMonitorSchema()
	err := APICreateOrUpdate(d, meta, "healthmonitor", s)
	if err == nil {
		err = ResourceAviHealthMonitorRead(d, meta)
	}
	return err
}

func resourceAviHealthMonitorUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceHealthMonitorSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "healthmonitor", s)
	if err == nil {
		err = ResourceAviHealthMonitorRead(d, meta)
	}
	return err
}

func resourceAviHealthMonitorDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "healthmonitor")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
