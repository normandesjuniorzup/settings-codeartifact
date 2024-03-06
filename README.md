# settings.json generator with AWS Credentials

## Requirements

- AWS CLI v2 already installed
- AWS Profile already configured

### AWS Profile

Add the following code snippet to the ~/.aws/config file

```
[profile CodeArtifactClaro]
sso_start_url = https://clarocorp.awsapps.com/start
sso_region = us-east-1
sso_account_id = 319569500149
sso_role_name = ViewOnlyAccess
region = us-east-1
output = json
```
