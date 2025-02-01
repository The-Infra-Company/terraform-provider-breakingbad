<a href="https://terraform.io">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset=".github/terraform_logo_dark.svg">
    <source media="(prefers-color-scheme: light)" srcset=".github/terraform_logo_light.svg">
    <img src=".github/terraform_logo_light.svg" alt="Terraform logo" title="Terraform" align="right" height="50">
  </picture>
</a>

# Terraform Breaking Bad Provider

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 0.12 or [OpenTofu](https://opentofu.org/) >= 1.5
- [Go](https://golang.org/doc/install) >= 1.22 (to build the provider plugin)


## Examples

```terraform
terraform {
  required_version = ">= 1.5.0"

  required_providers {
    breakingbad = {
      source = "The-Infra-Company/breakingbad"
    }
  }
}

provider "breakingbad" {
  api_url = var.api_url
}
```

## Building the Provider

> [!NOTE]
> The following installation uses [Taskfile](https://taskfile.dev/), which can be downloaded by running the following command:
>
> `brew install go-task`

Clone repository:

```sh
git clone https://github.com/The-Infra-Company/terraform-provider-breakingbad.git
```

Enter the provider directory and build the provider:

```sh
cd terraform-provider-breakingbad
task build
```

To use a released provider in your Terraform environment, run [`terraform init`](https://www.terraform.io/docs/commands/init.html) and Terraform will automatically install the provider. To specify a particular provider version when installing released providers, see the [Terraform documentation on provider versioning](https://www.terraform.io/docs/configuration/providers.html#version-provider-versions).

To instead use a custom-built provider in your Terraform environment (e.g. the provider binary from the build instructions above), follow the instructions to [install it as a plugin.](https://www.terraform.io/docs/plugins/basics.html#installing-plugins) After placing the custom-built provider into your plugins directory, run `terraform init` to initialize it.

## Contributing

Check out our [Contributing Docs](./CONTRIBUTING.md) for more information on how to support new resources and data sources, test, and contribute to the provider!

For bug reports & feature requests, please use the [issue tracker](https://github.com/The-Infra-Company/terraform-provider-breakingbad/issues).

PRs are welcome! We follow the typical "fork-and-pull" Git workflow.
 1. **Fork** the repo on GitHub
 2. **Clone** the project to your own machine
 3. **Commit** changes to your own branch
 4. **Push** your work back up to your fork
 5. Submit a **Pull Request** so that we can review your changes

> [!TIP]
> Be sure to merge the latest changes from "upstream" before making a pull request!


