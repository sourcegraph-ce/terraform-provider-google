# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

output "status_page_public_ip" {
  value = google_compute_instance.project_2_vm.network_interface[0].access_config[0].nat_ip
}
