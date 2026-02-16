output "written_file_path" {
  value       = local_file.hello.filename
  description = "The path of the file Terraform wrote."
}

