## Overview

Steps to setup a pipeline for building alexa skills using this provider

## Requirements

* ask cli

## Setup

1. Create an account on https://developer.amazon.com/ as a human user
   1. Add additional humans to the account as administrators and developers as appropriate
2. Add a service account to the account as a developer
3. Login to the account as the service account via a browser
4. Run the command, and follow the prompts:
   `ask configure --profile default --no-browser`
5. Create pipeline variables for the resulting token values, including:
   * access_token
   * refresh_token
   * expires_in
   * expires_at
   * vendor_id
6. In the pipeline, render the variables into the ~/.ask/cli_config file
   * see sample template file at examples/.ask/cli_config