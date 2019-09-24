package avi

import (
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

func TestAVIDataSourceSchedulerBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSSchedulerConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_scheduler.testScheduler", "name", "test-Default-Scheduler-abc"),
					resource.TestCheckResourceAttr(
						"avi_scheduler.testScheduler", "enabled", "true"),
					resource.TestCheckResourceAttr(
						"avi_scheduler.testScheduler", "frequency", "1"),
				),
			},
		},
	})

}

const testAccAVIDSSchedulerConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
data "avi_backupconfiguration" "default_backupconfig" {
    name= "Backup-Configuration"
}
resource "avi_scheduler" "testScheduler" {
	start_date_time = "2018-04-06T07:05:32.242307"
	name = "test-Default-Scheduler-abc"
	enabled = true
	tenant_ref = data.avi_tenant.default_tenant.id
	frequency_unit = "SCHEDULER_FREQUENCY_UNIT_DAY"
	backup_config_ref = data.avi_backupconfiguration.default_backupconfig.id
	frequency = "1"
	scheduler_action = "SCHEDULER_ACTION_BACKUP"
	run_mode = "RUN_MODE_PERIODIC"

}

data "avi_scheduler" "testScheduler" {
    name= avi_scheduler.testScheduler.name
}
`
