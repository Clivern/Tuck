variable "output_path" {
  type        = string
  description = "Where Terraform should write the file."
  default     = "output.txt"
}

variable "content" {
  type        = string
  description = "Content to write into the file."
  default     = "Hello from Terraform!\n"
}
