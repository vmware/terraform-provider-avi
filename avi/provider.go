/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"fmt"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/avinetworks/sdk/go/session"
	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"log"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"avi_username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("AVI_USERNAME", nil),
				Description: "Username for Avi Controller.",
			},
			"avi_controller": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("AVI_CONTROLLER", nil),
				Description: "Avi Controller hostname or IP address.",
			},
			"avi_password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("AVI_PASSWORD", nil),
				Description: "Password for Avi Controller.",
			},
			"avi_tenant": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("AVI_TENANT", nil),
				Description: "Avi tenant for Avi Controller.",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"avi_pool":          dataSourceAviPool(),
			"avi_healthmonitor": dataSourceAviHealthMonitor(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"avi_pool":          resourceAviPool(),
			"avi_healthmonitor": resourceAviHealthMonitor(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Credentials{
		Username:   d.Get("avi_username").(string),
		Password:   d.Get("avi_password").(string),
		Controller: d.Get("avi_controller").(string),
		Tenant:     d.Get("avi_tenant").(string),
	}

	log.Println("[INFO] XXXXXXXXXXXXX Avi Client %v", config)

	if err := config.validate(); err != nil {
		return nil, err
	}

	log.Println("[INFO] XXXXXXXXXXXXX Avi Client validated ")

	aviClient, err := clients.NewAviClient(
		config.Controller, config.Username,
		session.SetPassword(config.Password),
		session.SetTenant(config.Tenant),
		session.SetInsecure)

	log.Println("Avi Client created ")

	return aviClient, err
}

type Credentials struct {
	Username   string
	Password   string
	Controller string
	Port       string
	Tenant     string
}

func (c *Credentials) validate() error {
	var err *multierror.Error

	if c.Controller == "" {
		err = multierror.Append(err, fmt.Errorf("Avi Controller must be provided"))
	}

	if c.Username == "" {
		err = multierror.Append(err, fmt.Errorf("Avi Controller username must be provided"))
	}

	if c.Password == "" {
		err = multierror.Append(err, fmt.Errorf("Avi Controller password must be provided"))
	}
	return err.ErrorOrNil()
}
