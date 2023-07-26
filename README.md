# settings.json generator with AWS Credentials

## Requirements

- AWS CLI v2 already installed
- AWS Profile already configured

### AWS Profile

Add the following code snippet to the ~/.aws/config file

```
[profile CodeArtifact]
sso_start_url = https://zup.awsapps.com/start
sso_region = us-east-1
sso_account_id = 546045978864
sso_role_name = ZupCodeArtifactUsers
region = us-east-1
output = json
```
