locals {
  output_file = (startswith(var.output_path, "/") || startswith(var.output_path, "~")) ? pathexpand(var.output_path) : abspath("${path.module}/${var.output_path}")
}

resource "local_file" "hello" {
  filename = local.output_file
  content  = var.content
}

