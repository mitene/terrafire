terraform {
  backend "local" {
    path = "/tmp/terrafire/terraform.tfstate"
  }
}

variable "data" {
  type = map(string)
}

resource "local_file" "foo" {
  content     = jsonencode(var.data)
  filename = "/tmp/terrafire/foo.json"
}

output "content" {
  value = local_file.foo.content
}
