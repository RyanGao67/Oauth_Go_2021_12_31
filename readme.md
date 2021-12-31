Just a plain box here.
---
Centos7 | a few utils | PasswordAuthentication yes.

# cmd vagrant
```
vagrant up
vagrant ssh
vagrant suspend
vagrant status
vagrant resume
vagrant --help

```

```
ip route show
netstat -ntlp
bin/standalone -b 10.100.196.60
-Djboss.bind.address=10.100.196.60
```

### Realm > Realm settings > general > end points(openid endpoint Configuration) 
What exactly the OpenID connected?   
OpenID Connect is, in fact, an OAuth 2.0 framework with some extra layer on top of it.

In Keycloak, although the OAuth option isn’t available here. OpenID Connect is inheriting the whole functionality from the OAuth 2.0 framework.

You can use Keycloak as an authorization server that implements OAuth 2.0 without having to use a single functionality of OpenID Connect.

```
{
   "issuer":"http://10.100.196.60:8080/auth/realms/learningApp",
   "authorization_endpoint":"http://10.100.196.60:8080/auth/realms/learningApp/protocol/openid-connect/auth",
   "token_endpoint":"http://10.100.196.60:8080/auth/realms/learningApp/protocol/openid-connect/token",
   "introspection_endpoint":"http://10.100.196.60:8080/auth/realms/learningApp/protocol/openid-connect/token/introspect",
   "userinfo_endpoint":"http://10.100.196.60:8080/auth/realms/learningApp/protocol/openid-connect/userinfo",
   "end_session_endpoint":"http://10.100.196.60:8080/auth/realms/learningApp/protocol/openid-connect/logout",
   "frontchannel_logout_session_supported":true,
   "frontchannel_logout_supported":true,
   "jwks_uri":"http://10.100.196.60:8080/auth/realms/learningApp/protocol/openid-connect/certs",
   "check_session_iframe":"http://10.100.196.60:8080/auth/realms/learningApp/protocol/openid-connect/login-status-iframe.html",
   "grant_types_supported":[
      "authorization_code",
      "implicit",
      "refresh_token",
      "password",
      "client_credentials",
      "urn:ietf:params:oauth:grant-type:device_code",
      "urn:openid:params:grant-type:ciba"
   ],
   "response_types_supported":[
      "code",
      "none",
      "id_token",
      "token",
      "id_token token",
      "code id_token",
      "code token",
      "code id_token token"
   ],
   "subject_types_supported":[
      "public",
      "pairwise"
   ],
   "id_token_signing_alg_values_supported":[
      "PS384",
      "ES384",
      "RS384",
      "HS256",
      "HS512",
      "ES256",
      "RS256",
      "HS384",
      "ES512",
      "PS256",
      "PS512",
      "RS512"
   ],
   "id_token_encryption_alg_values_supported":[
      "RSA-OAEP",
      "RSA-OAEP-256",
      "RSA1_5"
   ],
   "id_token_encryption_enc_values_supported":[
      "A256GCM",
      "A192GCM",
      "A128GCM",
      "A128CBC-HS256",
      "A192CBC-HS384",
      "A256CBC-HS512"
   ],
   "userinfo_signing_alg_values_supported":[
      "PS384",
      "ES384",
      "RS384",
      "HS256",
      "HS512",
      "ES256",
      "RS256",
      "HS384",
      "ES512",
      "PS256",
      "PS512",
      "RS512",
      "none"
   ],
   "request_object_signing_alg_values_supported":[
      "PS384",
      "ES384",
      "RS384",
      "HS256",
      "HS512",
      "ES256",
      "RS256",
      "HS384",
      "ES512",
      "PS256",
      "PS512",
      "RS512",
      "none"
   ],
   "request_object_encryption_alg_values_supported":[
      "RSA-OAEP",
      "RSA-OAEP-256",
      "RSA1_5"
   ],
   "request_object_encryption_enc_values_supported":[
      "A256GCM",
      "A192GCM",
      "A128GCM",
      "A128CBC-HS256",
      "A192CBC-HS384",
      "A256CBC-HS512"
   ],
   "response_modes_supported":[
      "query",
      "fragment",
      "form_post",
      "query.jwt",
      "fragment.jwt",
      "form_post.jwt",
      "jwt"
   ],
   "registration_endpoint":"http://10.100.196.60:8080/auth/realms/learningApp/clients-registrations/openid-connect",
   "token_endpoint_auth_methods_supported":[
      "private_key_jwt",
      "client_secret_basic",
      "client_secret_post",
      "tls_client_auth",
      "client_secret_jwt"
   ],
   "token_endpoint_auth_signing_alg_values_supported":[
      "PS384",
      "ES384",
      "RS384",
      "HS256",
      "HS512",
      "ES256",
      "RS256",
      "HS384",
      "ES512",
      "PS256",
      "PS512",
      "RS512"
   ],
   "introspection_endpoint_auth_methods_supported":[
      "private_key_jwt",
      "client_secret_basic",
      "client_secret_post",
      "tls_client_auth",
      "client_secret_jwt"
   ],
   "introspection_endpoint_auth_signing_alg_values_supported":[
      "PS384",
      "ES384",
      "RS384",
      "HS256",
      "HS512",
      "ES256",
      "RS256",
      "HS384",
      "ES512",
      "PS256",
      "PS512",
      "RS512"
   ],
   "authorization_signing_alg_values_supported":[
      "PS384",
      "ES384",
      "RS384",
      "HS256",
      "HS512",
      "ES256",
      "RS256",
      "HS384",
      "ES512",
      "PS256",
      "PS512",
      "RS512"
   ],
   "authorization_encryption_alg_values_supported":[
      "RSA-OAEP",
      "RSA-OAEP-256",
      "RSA1_5"
   ],
   "authorization_encryption_enc_values_supported":[
      "A256GCM",
      "A192GCM",
      "A128GCM",
      "A128CBC-HS256",
      "A192CBC-HS384",
      "A256CBC-HS512"
   ],
   "claims_supported":[
      "aud",
      "sub",
      "iss",
      "auth_time",
      "name",
      "given_name",
      "family_name",
      "preferred_username",
      "email",
      "acr"
   ],
   "claim_types_supported":[
      "normal"
   ],
   "claims_parameter_supported":true,
   "scopes_supported":[
      "openid",
      "email",
      "web-origins",
      "offline_access",
      "phone",
      "roles",
      "microprofile-jwt",
      "profile",
      "address"
   ],
   "request_parameter_supported":true,
   "request_uri_parameter_supported":true,
   "require_request_uri_registration":true,
   "code_challenge_methods_supported":[
      "plain",
      "S256"
   ],
   "tls_client_certificate_bound_access_tokens":true,
   "revocation_endpoint":"http://10.100.196.60:8080/auth/realms/learningApp/protocol/openid-connect/revoke",
   "revocation_endpoint_auth_methods_supported":[
      "private_key_jwt",
      "client_secret_basic",
      "client_secret_post",
      "tls_client_auth",
      "client_secret_jwt"
   ],
   "revocation_endpoint_auth_signing_alg_values_supported":[
      "PS384",
      "ES384",
      "RS384",
      "HS256",
      "HS512",
      "ES256",
      "RS256",
      "HS384",
      "ES512",
      "PS256",
      "PS512",
      "RS512"
   ],
   "backchannel_logout_supported":true,
   "backchannel_logout_session_supported":true,
   "device_authorization_endpoint":"http://10.100.196.60:8080/auth/realms/learningApp/protocol/openid-connect/auth/device",
   "backchannel_token_delivery_modes_supported":[
      "poll",
      "ping"
   ],
   "backchannel_authentication_endpoint":"http://10.100.196.60:8080/auth/realms/learningApp/protocol/openid-connect/ext/ciba/auth",
   "backchannel_authentication_request_signing_alg_values_supported":[
      "PS384",
      "ES384",
      "RS384",
      "ES256",
      "RS256",
      "ES512",
      "PS256",
      "PS512",
      "RS512"
   ],
   "require_pushed_authorization_requests":false,
   "pushed_authorization_request_endpoint":"http://10.100.196.60:8080/auth/realms/learningApp/protocol/openid-connect/ext/par/request",
   "mtls_endpoint_aliases":{
      "token_endpoint":"http://10.100.196.60:8080/auth/realms/learningApp/protocol/openid-connect/token",
      "revocation_endpoint":"http://10.100.196.60:8080/auth/realms/learningApp/protocol/openid-connect/revoke",
      "introspection_endpoint":"http://10.100.196.60:8080/auth/realms/learningApp/protocol/openid-connect/token/introspect",
      "device_authorization_endpoint":"http://10.100.196.60:8080/auth/realms/learningApp/protocol/openid-connect/auth/device",
      "registration_endpoint":"http://10.100.196.60:8080/auth/realms/learningApp/clients-registrations/openid-connect",
      "userinfo_endpoint":"http://10.100.196.60:8080/auth/realms/learningApp/protocol/openid-connect/userinfo",
      "pushed_authorization_request_endpoint":"http://10.100.196.60:8080/auth/realms/learningApp/protocol/openid-connect/ext/par/request",
      "backchannel_authentication_endpoint":"http://10.100.196.60:8080/auth/realms/learningApp/protocol/openid-connect/ext/ciba/auth"
   }
}
```

We need to register a client, which is mutually known between the client and the Keycloak.

Oauth specification doesn’t mention about the process of client registration. At the end of the day, a client just needs to be known and recognizable by the Authorization server after all.

If you use Google OAuth there is a process to do so in the Google way. If you use Facebook OAuth's, there is also Facebook’s way of client registration as well, besides, for KeyCloak is just a few clicks away, so let’s do it real quick. Go to clients create name a client anything you like.

Root URL is the root URL of your client’s URL. In other word, the IP or domain name of your client. Client Protocol, there are two options to choose here, The OpenID-Connect is what I have mentioned earlier. Another is SAML which stands for “Security Assertion Markup Language”. We are not going to talk about SAML here, but it is good to know what it is. SAML is similar to OpenID Connect, but instead of JSON format, SAML uses the XML version for sending and receiving documents.










It is http://localhost:8080. Copy that and paste it here. Save them all. Now the brand new client has created.

Consent will be revisited later. Client protocol, yes open-id connect.

Access Type or you can call “client type”. We need to talk more about it here. OAuth defines two client types, based on their ability to authenticate securely with the authorization server. In another word, an ability to maintain the confidentiality of their client credentials. In confidential type. Clients are capable of maintaining the confidentiality of their credentials or capable of secure client authentication using other means. For instance, the client implemented on a secure server with restricted access to the client credentials. In our case, our client is of this type. Because we are going to send client credentials to keycloak without the user’s agent or browser to be able to observe the client credential. You will see a code for that shortly. For public clients. It is incapable of maintaining the confidentiality of their credentials such as a native application on a mobile phone or web browser-based applications. Why is that so? Because it can be inspected by developer tools or by using other means. And hence no longer secret anymore. We will think about this in more detail once our client is fully developed and works with Protected Resources successfully. bearer-only is another form of confidential type.

Standard flow, of course enable it. Save it. That is it. Reload the client.



# access token

```
{
   "access_token":"eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJwcV9XYnhtYW92WVFCanFmWjFaTUc3UWVTTFpQdUk5NkpKdnduNnFSb1g0In0.eyJleHAiOjE2NDA4MzQxNTMsImlhdCI6MTY0MDgzMzg1MywiYXV0aF90aW1lIjoxNjQwODMyMTUxLCJqdGkiOiJhOTU4NmU5ZS0xZGRjLTRkNmQtYTk1ZC03YzIwOGNkNTZjNDciLCJpc3MiOiJodHRwOi8vMTAuMTAwLjE5Ni42MDo4MDgwL2F1dGgvcmVhbG1zL2xlYXJuaW5nQXBwIiwiYXVkIjoiYWNjb3VudCIsInN1YiI6IjYxODFkMDc2LWE4NzctNDk4Ni04MGZkLWQ1NzQ1ODY0MGMwNiIsInR5cCI6IkJlYXJlciIsImF6cCI6ImJpbGxpbmdBcHAiLCJzZXNzaW9uX3N0YXRlIjoiZTA1ZTIxMjYtMjJlMC00ZWMyLTkzMTktNWFiMjFjNjYwYWEwIiwiYWNyIjoiMCIsImFsbG93ZWQtb3JpZ2lucyI6WyJodHRwOi8vMTI3LjAuMC4xOjgwODAiXSwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbIm9mZmxpbmVfYWNjZXNzIiwiZGVmYXVsdC1yb2xlcy1sZWFybmluZ2FwcCIsInVtYV9hdXRob3JpemF0aW9uIl19LCJyZXNvdXJjZV9hY2Nlc3MiOnsiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwidmlldy1wcm9maWxlIl19fSwic2NvcGUiOiJlbWFpbCBwcm9maWxlIiwic2lkIjoiZTA1ZTIxMjYtMjJlMC00ZWMyLTkzMTktNWFiMjFjNjYwYWEwIiwiZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJuYW1lIjoiYm9iIGJvYiIsInByZWZlcnJlZF91c2VybmFtZSI6ImJvYiIsImdpdmVuX25hbWUiOiJib2IiLCJmYW1pbHlfbmFtZSI6ImJvYiIsImVtYWlsIjoiYm9iQGdtYWlsLmNvbSJ9.V2Jd30BYXHaiAZvXBt7eYKIwi2O18KNdWGT7wSsYb0tf6PC_v0G9GJ9N2LkP4ts6rWcDIuUVYBiMzQPnAIPFnJAAvUN_F7eFScK_1TLaqj7Nlu7czuqNNjETRY0lkEo5zEJsVgZn3wTCaWvFwLqKwqackU-_ygxXvJ1J2lrwAZoZIcIhmMaNID2T4K8XoZbrylMbPV6AWBikqm54IxtsNSFx11oRz9J-F1GBNG3N5DJnviwbRDWZP6_CAMgGJQpGDvdLKCV0XWHxAdFY93YnuQNP8tIzsuZ1wSu8DZTvP9DRbOvxLMrxpwQXpshZScMepFqcAtJvV8LH0iAuVrHngg",
   "expires_in":300,
   "refresh_expires_in":1800,
   "refresh_token":"eyJhbGciOiJIUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICI5OGRlZjM2Mi1iMjQzLTQ5MDMtOTViYi0xODE1Mzk1ZDg3ZWIifQ.eyJleHAiOjE2NDA4MzU2NTMsImlhdCI6MTY0MDgzMzg1MywianRpIjoiOGM0NjA3YzAtODg0OS00ZWMyLWFkNjctMmNjMmMwY2JlZTIyIiwiaXNzIjoiaHR0cDovLzEwLjEwMC4xOTYuNjA6ODA4MC9hdXRoL3JlYWxtcy9sZWFybmluZ0FwcCIsImF1ZCI6Imh0dHA6Ly8xMC4xMDAuMTk2LjYwOjgwODAvYXV0aC9yZWFsbXMvbGVhcm5pbmdBcHAiLCJzdWIiOiI2MTgxZDA3Ni1hODc3LTQ5ODYtODBmZC1kNTc0NTg2NDBjMDYiLCJ0eXAiOiJSZWZyZXNoIiwiYXpwIjoiYmlsbGluZ0FwcCIsInNlc3Npb25fc3RhdGUiOiJlMDVlMjEyNi0yMmUwLTRlYzItOTMxOS01YWIyMWM2NjBhYTAiLCJzY29wZSI6ImVtYWlsIHByb2ZpbGUiLCJzaWQiOiJlMDVlMjEyNi0yMmUwLTRlYzItOTMxOS01YWIyMWM2NjBhYTAifQ.90Ew2TxqJbUkhRGqhDJiUpHSGRdvPIyK-5BlARXBR-g",
   "token_type":"Bearer",
   "not-before-policy":1640817515,
   "session_state":"e05e2126-22e0-4ec2-9319-5ab21c660aa0",
   "scope":"email profile"
}
```






