data "breakingbad_quote" "this" {}

locals {
  formatted_quote = "${data.breakingbad_quote.this.quote} - ${data.breakingbad_quote.this.author}"
}

output "quote" {
  value = local.formatted_quote
}
