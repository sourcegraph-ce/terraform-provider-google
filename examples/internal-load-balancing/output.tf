# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

output "internal_load_balancer_ip" {
  value = google_compute_forwarding_rule.my-int-lb-forwarding-rule.ip_address
}
