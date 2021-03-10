# OIDC

## openid connect

OpenID Connect is a simple identity layer built on top of the OAuth 2.0 protocol, which allows clients to verify the identity of an end user based on the authentication performed by an authorization server or identity provider (IdP).

OIDC supprots the following authentication flows:
- The **Implicit Flow** is required for apps that have no “back end” logic on the web server, like a Javascript app.
- The **Authentication (or Basic) Flow** is designed for apps that have a back end that can communicate with the IdP away from prying eyes.
- The **Resource Owner Password Grant** does not have an login UI and is useful when access to a web browser is not possible.
- The **Client Credentials Grant** is useful for machine to machine authorization.

### Authentication Flow
[reference](https://docs.axway.com/bundle/APIGateway_762_OAuthUserGuide_allOS_en_HTML5/page/Content/OAuthGuideTopics/OpenidImport/openid_flow.htm)

![workflow](./images/workflow.png)

When client app redirects the user to the authorization endpointm it should have the following parameters:

| parameters | description |
| --- | --- |
| response_type | must be set to `code` |
| client_id | the client id we registered in the IdP |
| redirect_uri | the location where the authorization code will be sent |
| scope | Optional a space delimited list of scopes, which indicate the access to the resource owner's data requested by the application (openid is the required scope) |
| state | Recommended. Any state the consumer wants reflected back to itself after approval during the callback. This value is used to prevent cross-site request forgery (CSRF) attacks, the consumer delivers this value to the authorization server when making an authorization request. For this reason, the value must be opaque, kept secret by the consumer, and known only to the consumer and the OAuth provider |

for example:
```
http://127.0.0.1:5556/dex/auth?client_id=example-app&redirect_uri=http%3A%2F%2F127.0.0.1%3A5555%2Fcallback&response_type=code&scope=openid+profile+email&state=I+wish+to+wash+my+irish+wristwatch
```

**for details of state, plz refer to: https://tools.ietf.org/html/rfc6819#section-3.6**

After IdP verifies the user's identity, it redirects the user back to the client app with a code that can be exchanged for an ID token.

for example:
```
http://127.0.0.1:5555/callback?code=ewwj54h3zgiwaikpyolcv7nwc&state=I+wish+to+wash+my+irish+wristwatch
```

**no one could understand the code except the IdP**

Then, client app:
1. extract the state and verify the state
2. extract the code
3. use the code to exchange the token
4. extract and verify the id token from the token
5. return id token to the user

Step 3 and step 4 are parts of oauth2 and oidc package, we could use existing ones for each program language.

User once gets an id token, he/she could use it call the api belonged to client-app, client-app would consume [claims](https://auth0.com/docs/scopes/openid-connect-scopes) (which are name/value pairs that contain information about a user) to verify the identify of the user.

for example:
```json
{
  "iss": "http://127.0.0.1:5556/dex", //idp issuer url
  "sub": "CiQxZDkyYzA3Ni1jYjQ4LTQ4ZjEtYTk5Ni03YzhkMGJkZmE4NjMSBWxvY2Fs",
  "aud": "example-app",
  "exp": 1615453229, // token will expire at
  "iat": 1615366829, // token was issued at
  "at_hash": "hWaeSk3_aWCdJ8XJ_wAlPg",
  "c_hash": "LjuZ4Tw49jVcr1dGqQ5vqw",
  "email": "tianqiuhuang@gmail.com",
  "email_verified": true,
  "name": "TQ"
}
```

## WorkShop

setup your own IdP and write a client-app.

Refer to [setup](./setup-idp-and-client-app.md).
