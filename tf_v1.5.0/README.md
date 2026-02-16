# tf_v1.5.0

Very small Terraform example that writes a local file using the `hashicorp/local` provider.

## Run

```bash
cd tf_v1.5.0
terraform init
terraform apply
```

After apply, you should see `output.txt` created in this directory.

## Customize

```bash
terraform apply -var='content=Custom content\n' -var='output_path=./myfile.txt'
```
