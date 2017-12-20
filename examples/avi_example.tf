provider "avi" {
  avi_username = ""
  avi_tenant = ""
  avi_password = ""
  avi_controller= ""
}

resource "avi_networkprofile" "test_networkprofile" {
  name= "networkprofile-1"
  profile{
    type= "PROTOCOL_TYPE_TCP_PROXY"
  }
}

resource "avi_applicationpersistenceprofile" "test_applicationpersistenceprofile" {
  name = "applicationpersistence-1"
  persistence_type = "PERSISTENCE_TYPE_CLIENT_IP_ADDRESS"
}

resource "avi_vsvip" "test_vsvip" {
  name= "vip-1"
  vip {
    vip_id= "0"
    ip_address {
      type= "V4",
      addr= "10.90.64.88",
    }
  }
}

resource "avi_virtualservice" "test_vs" {
  name= "vs-1"
  pool_ref= "${avi_pool.testpool.id}"
  application_profile_ref= "${avi_applicationprofile.test_applicationprofile.id}"
  network_profile_ref = "${avi_networkprofile.test_networkprofile.id}"
  vsvip_ref = "${avi_vsvip.test_vsvip.id}"
  vip {
    vip_id= "0"
    ip_address {
      type= "V4",
      addr= "10.90.64.88",
    }
  }
  services {
    port= 80
    enable_ssl= true
  }
}

resource "avi_applicationprofile" "test_applicationprofile" {
  name= "applicationprofile-1"
  type= "APPLICATION_PROFILE_TYPE_DNS"
}

resource "avi_healthmonitor" "test_hm_1" {
  name = "healthmonitor-1"
  type = "HEALTH_MONITOR_HTTP"
}


resource "avi_pool" "testpool" {
  name= "pool-5",
  health_monitor_refs= ["${avi_healthmonitor.test_hm_1.id}"]
  application_persistence_profile_ref= "${avi_applicationpersistenceprofile.test_applicationpersistenceprofile.id}"
  servers {
    ip= {
      type= "V4",
      addr= "10.90.64.66",
    }
    port= 8080
  }
  fail_action= {
    type= "FAIL_ACTION_CLOSE_CONN"
  }
}
