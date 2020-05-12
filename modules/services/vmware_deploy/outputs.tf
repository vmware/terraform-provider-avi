output "vsphere_virtual_machine_vm1" {
  value = vsphere_virtual_machine.vm[0].default_ip_address
}

output "vsphere_virtual_machine_vm2" {
  value = vsphere_virtual_machine.vm[1].default_ip_address
}

output "vsphere_virtual_machine_vm3" {
  value = vsphere_virtual_machine.vm[2].default_ip_address
}

output "avi_cluster_output" {
  value = avi_cluster.vmware_cluster
}

