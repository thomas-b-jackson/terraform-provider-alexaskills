## Overview

Options for building Alexa skills resources in Terraform.

## Assumptions

1.	We will built a lex QnA bot per [this AWS ML blog](https://aws.amazon.com/blogs/machine-learning/creating-a-question-and-answer-bot-with-amazon-lex-and-amazon-alexa/)
  * the bot and associated lambda will be built using terraform resources via the AWS provider instead of cloud formation
2.	Per link above, the alexa skill will be built using the bot's interaction model and lambda function
   * Any change to the bot intents will require a rebuild and re-certification of the alexa skill

## Analysis 

Ideally we would leverage an existing terraform provider to build the skill resources. The provider would leverage the [AWS skill management API](https://developer.amazon.com/en-US/docs/alexa/smapi/smapi-overview.html).

Unfortunately, there is no existing terraform provider because:
1. the SMAPI SDK is only available in Node.js, Python, and Java, and
2. the Terraform plugin SDKs (SDKv2 and Terraform  Plugin Framework) only support golang

It seems likely that the SMAPI SDK will eventually be available in golang. In theory, `terraform-provider-aws` could then get updated to support a skills resource. However, there are additional gaps that would also need to get filled. All skills must currently be hosted on `developer.amazon.com`, and authentication to the developer site requires a set of credentials that are different from normal AWS credentials. And unlike most aws resource, skills have a git-based lifecycle/workflow, including a manual `certification` process where Amazon employees review changes to skill intents via a merge request. Amazon will need to migrate skills resources from `developer.amazon.com` to `console.aws.amazon.com` before they become available in `terraform-provider-aws`, but it is unclear if they ever intend to do so.

The SMAPI SDK is also available via a node-based command line utility called `ask`. Our two recommended solutions leverage this utility.

## Recommendation

To work around the current situation, we recommend one of the following two options:
1. make calls to the SMAPI SDK via the `ask cli` in a custom provider (`terraform-provider-aws-skills`), or
2. make calls to the SMAPI SDK via the `ask cli` in a `null_resource`

**NOTE:**
Two additional options were investigated but are NOT recommended:
1. make calls to the SMAPI SDK via the `ask cli` in an extended version of the `terraform-provider-aws`, or
2. build skills using a `aws_cloudformation_stack` resource and certify skills via a `null_resource`

Details on each are provided in the appendix.

## Preferred Option: terraform-provider-aws-skills

Highlights:
* build a provider plugin (using SDKv2) in golang that is hosted as an executable on build clients
* go accesses the SMAPI SDK via calls to the `ask cli` using `exec.Command()`

### Pros

* pure terraform
  * full access to terraform Create/Update/Delete operations
  * skill state accurately reflected in terraform state
* golang-based SDLC
  * re-certification logic in go
  * fine-grained error handling
  * unit testing

### Cons

* more dev effort that option 2
* dependency on external binaries

### Effort

* Dev Effort: 12-18 days
* Pipeline Integration: 2-3 days

Pipeline integration effort includes:
* adding plugin binary, secrets, ask, and ask dependencies (node, etc)

## Second-Best Option: null_resource

Highlights:
* create a terraform module that wraps a bash script in a null_resource
* bash script accesses the SMAPI SDK via calls to the `ask cli`

```
module alexa_skills {
  source = "../scripts/ask.sh"
  depends_on = [module.lex_bot]
  credentials = var.ask_access_token
  recertify = true
}
```

### Pros

* less work than plugin
  
### Cons

* complex logic implemented in bash
  * compounded by poor unit testing options
* no access to terraform Create/Update/Delete operations
* skill state not reflected in terraform state
* certification may need to be module input

### Effort

* Dev Effort: 5-10 days
* Pipeline Integration: 3-5 days

Pipeline integration effort includes:
* adding bash, secrets, ask, and ask dependencies (node, etc)

## Appendix

Options for building Alexa skills resources in Terraform that were explored but are NOT recommended.

### Extended version of terraform-provider-aws

Highlights:
* fork [terraform-provider-aws](https://github.com/hashicorp/terraform-provider-aws) and add a skills resource
* go accesses the SMAPI SDK via calls to the `ask cli` using `exec.Command()`

### Pros

Same pros as `terraform-provider-aws-skills` provider

### Cons

* Likely more effort than the custom provider
  * due to lengthy review process
* Not likely to get approved and merged
  * different credentials
  * calling a node app from provider plugin is a likely non-starter

### aws_cloudformation_stack + null_resource

Highlights:

* build skills using a `aws_cloudformation_stack` resource
* create a terraform `certification` module that wraps a bash script in a null_resource
* bash script accesses the SMAPI SDK via calls to the `ask cli`

### Pros

* somewhat less effort than using null_resource for everything

### Cons

* mix of declarative syntaxes
* infrastructure state split between terraform and cloud formation
* same cons with using null_resource for everything