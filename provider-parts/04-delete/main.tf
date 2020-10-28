terraform {
  required_providers {
    todo = {
      source  = "terraform.spkane.org/spkane/todo"
      version = "1.1.0"
    }
  }
}

provider "todo" {
  host = "127.0.0.1"
  port = "8080"
  apipath = "/"
  schema = "http"
}

resource "todo" "test1" {
  count = 5
  description = "${count.index}-1 ${var.purpose} todo"
  completed = false
}

