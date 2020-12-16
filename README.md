# saml2aws-refresh

Automatically refresh aws saml session

While using [saml2aws](https://github.com/Versent/saml2aws) to access AWS accounts, the session expires after 1 hour due to max-duration setting on AWS account. `saml2aws-refresh` helps here to automatically refresh the aws sessions for the expected number of times, so that we can run this command once and forget about session expiry.  

