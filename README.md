# ospath Terraform Provider
This provider allows you to join paths in an OS-independent way. For example:
```terraform
resource "ospath" "my_folder" {
  path = ["src", "my_folder"]
}

outputs "my_folder_path" {
  value = ospath_join.my_folder.result
}
```
The output `my_folder_path` will return `src/my_folder` on Linux and macOS, but will return `src\my_folder` on Windows host systems.

Under the hood, the path list is passed to [Go's path package's Join](https://golang.org/pkg/path/#Join). Hence (from the documentation's examples), here are some behaviors of `join` resource (assuming it is a Unix system):
```terraform
path = ["a", "b", "c"]
# => a/b/c
path = ["a/b", "c"]
# => a/b/c
path = ["a", ""]
# => a
```

## Installation
This project is currently in alpha, minimal testing has been done. To install this plugin, download the binary for your OS from the [releases](/joeltio/terraform-provider-ospath/releases) move it to your [terraform plugin folder](https://www.terraform.io/docs/configuration/providers.html#third-party-plugins).

After moving the binary, remember to run `terraform init` in your project to initialize the plugin.
