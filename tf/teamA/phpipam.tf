data "phpipam_subnet" "subnet" {
  subnet_address = "10.0.0.0"
  subnet_mask    = 8
  section_id = 1
}

// Get the first available address
data "phpipam_first_free_subnet" "next_subnet" {
  subnet_id = data.phpipam_subnet.subnet.subnet_id
  subnet_mask = 17
}

resource "phpipam_first_free_subnet" "new_subnet" {
  parent_subnet_id = data.phpipam_subnet.subnet.subnet_id
  subnet_mask = 17
  description = "Managed by Terraform - Team A"
}

