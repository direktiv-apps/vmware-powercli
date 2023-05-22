
# vmware-powercli 1.0

VMware PowerCLI Environment

---
- #### Categories: build, development
- #### Image: gcr.io/direktiv/functions/vmware-powercli 
- #### License: [Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)
- #### Issue Tracking: https://github.com/direktiv-apps/vmware-powercli/issues
- #### URL: https://github.com/direktiv-apps/vmware-powercli
- #### Maintainer: [direktiv.io](https://www.direktiv.io)
---

## About vmware-powercli

This function has Powershell Core and VMware PowerCLI resources. Built by Nathan Coad originally, modified by Direktiv team for production deployment.
The following versions have been included in the build:
 - VMware-PowerCLI-13.0.0-20829139
 - Microsoft Powershell-7.3

We pass in several environment variables to the powershell script to enable scripts to make use of Direktiv secrets for credentials. These are available as follows: <table> <thead> <tr> <th>Environment Variable</th> <th>Direktiv parameter</th> </tr> </thead> <tbody> <tr> <td>VCENTER_USER</td> <td>username</td> </tr> <tr> <td>VCENTER_PASSWORD</td> <td>password</td> </tr> <tr> <td>VCENTER</td> <td>vcenter</td> </tr> </tbody> </table>

### Example(s)
  #### Function Configuration
```yaml
functions:
- id: vmware-powercli
  image: gcr.io/direktiv/functions/vmware-powercli:1.0
  type: knative-workflow
  size: large
```
   #### Run small script directly
```yaml
- id: vmware-powercli 
  type: action
  action:
    function: vmware-powercli
    secrets: ["vcenterUser", "vCenterPass", "vCenterName"]
    input:
      username: jq(.secrets.vCenterUser)
      password: jq(.secrets.vCenterPass)
      vcenter: jq(.secrets.vCenterName)                 
      files:
      - name: script.ps1
        data: |
          Get-ChildItem . | Select Name | ConvertTo-Json 
      commands:
      - command: pwsh script.ps1
```
   #### Run file
```yaml
- id: vmware-powercli 
  type: action
  action:
    function: vmware-powercli
    secrets: ["vcenterUser", "vCenterPass", "vCenterName"]
    input: 
      username: jq(.secrets.vCenterUser)
      password: jq(.secrets.vCenterPass)
      vcenter: jq(.secrets.vCenterName)
      commands:
      - command: Get-VM -Name jq(.vm) | ConvertTo-Json -Depth 1 -AsArray
```

   ### Secrets


*No secrets required*







### Request



#### Request Attributes
[PostParamsBody](#post-params-body)

### Response
  List of executed Powershell commands.
#### Reponse Types
    
  

[PostOKBody](#post-o-k-body)
#### Example Reponses
    
```json
{
  "result": [
    {
      "Name": "file1.txt"
    },
    {
      "Name": "file2.txt"
    }
  ],
  "success": true
}
```

### Errors
| Type | Description
|------|---------|
| io.direktiv.command.error | Command execution failed |
| io.direktiv.output.error | Template error for output generation of the service |
| io.direktiv.ri.error | Can not create information object from request |


### Types
#### <span id="post-o-k-body"></span> postOKBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| vmware-powercli | [][PostOKBodyVmwarePowercliItems](#post-o-k-body-vmware-powercli-items)| `[]*PostOKBodyVmwarePowercliItems` |  | |  |  |


#### <span id="post-o-k-body-vmware-powercli-items"></span> postOKBodyVmwarePowercliItems

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| result | [interface{}](#interface)| `interface{}` | ✓ | |  |  |
| success | boolean| `bool` | ✓ | |  |  |


#### <span id="post-params-body"></span> postParamsBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| commands | [][PostParamsBodyCommandsItems](#post-params-body-commands-items)| `[]*PostParamsBodyCommandsItems` |  | | Array of commands. |  |
| files | [][DirektivFile](#direktiv-file)| `[]apps.DirektivFile` |  | | File to create before running commands. |  |
| password | string| `string` |  | | Password for authenticating to vCenter | `SecretPassword` |
| username | string| `string` |  | | Username for authenticating to vCenter | `user@domain.com` |
| vcenter | string| `string` |  | | FQDN of the vCenter this workflow will connect to | `vcsa.example.com` |


#### <span id="post-params-body-commands-items"></span> postParamsBodyCommandsItems

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| command | string| `string` |  | | Command to run |  |
| continue | boolean| `bool` |  | `true`| Stops excecution if command fails, otherwise proceeds with next command |  |
| envs | [][PostParamsBodyCommandsItemsEnvsItems](#post-params-body-commands-items-envs-items)| `[]*PostParamsBodyCommandsItemsEnvsItems` |  | | Environment variables set for each command. | `[{"name":"MYVALUE","value":"hello"}]` |
| print | boolean| `bool` |  | `true`| If set to false the command will not print the full command with arguments to logs. |  |
| silent | boolean| `bool` |  | | If set to false the command will not print output to logs. |  |


#### <span id="post-params-body-commands-items-envs-items"></span> postParamsBodyCommandsItemsEnvsItems

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| name | string| `string` |  | | Name of the variable. |  |
| value | string| `string` |  | | Value of the variable. |  |

 
