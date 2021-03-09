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

For "relative paths", use the path named value (see [References to Named Values](https://www.terraform.io/docs/configuration/expressions.html#references-to-named-values)) as part of the list:
```terraform
resource "ospath" "myfile" {
  path = [path.module, "src", "myModuleFile.txt"]
}
```

If you made the mistake of referencing a non-existent file with `ospath` using a file function (e.g. `file` or `filemd5`), use a ternary operator and `fileexists` with your file function:
```terraform
resource "ospath" "myfile" {
  path = [path.module, "src", "myModuleFile.txt"]
}

resource "aws_s3_bucket_object" "my_s3_obj" {
  # truncated...
  etag = fileexists(ospath_join.app_s3_cfn.result) ? filemd5(ospath_join.app_s3_cfn.result) : ""
}
```
This will allow terraform to evaluate the path lazily (see [#10878](https://github.com/hashicorp/terraform/issues/10878) and [`fileexists`](https://www.terraform.io/docs/configuration/functions/fileexists.html)).

Under the hood, the path list is passed to [Go's path package's Join](https://golang.org/pkg/path/#Join). Hence (from the documentation's examples), here are some behaviors of `join` resource (assuming it is a Unix system):
```terraform
path = ["a", "b", "c"]
# => a/b/c
path = ["a/b", "c"]
# => a/b/c
path = ["a", ""]
# => a
# .s are ignored, use the path variable instead (e.g. path.cwd, path.module, path.root)
path = [".", "a"]
# => a
```

## Installation
This project is currently in alpha, minimal testing has been done. To use this plugin, add `joeltio/ospath` as one of your required providers in yuor terraform file:
```
terraform {
  required_providers {
    ospath = {
      source = "joeltio/ospath"
      version = "0.0.3-alpha"
    }
  }
}
```
