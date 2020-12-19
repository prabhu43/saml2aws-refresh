# saml2aws-refresh

Automatically refresh aws saml session

While using [saml2aws](https://github.com/Versent/saml2aws) to access AWS accounts, the session expires after 1 hour due to max-duration setting on AWS account. `saml2aws-refresh` helps here to automatically refresh the aws sessions for the expected number of times, so that we can run this command once and forget about session expiry.  

### Installation
```
git clone git@github.com:prabhu43/saml2aws-refresh.git
cd saml2aws-refresh
make install
```

### Usage
Login to all aws profiles matching `abc` for 4 times with interval of 59 minutes
```
saml2aws-refresh --profile abc --count 4
```

#### Command Help

```
$ saml2aws --help

NAME:
   saml2aws-refresh - Automatically refresh AWS saml session

USAGE:
   saml2aws-refresh [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --count value    No. of times session has to be refreshed (default: 1)
   --profile value  AWS profile (partial match works if it matches exactly 1 profile)
   --help, -h       show help (default: false)
```
