workspace "app" {
  source "github" {
    owner = "mitene"
    repo  = "terrafire"
    path  = "examples/terraform"
    ref   = "web-ui"
  }

  workspace = "dev"
  vars = {
    data = {
      foo = "FOO"
      bar = "BAR"
    }
  }
}
