# HTTP Authentication

Learn how to use HTTP Basic Authentication and Digest Authentication.

## HTTP Basic Authentication

**Request**

```
Authorization: Basic BASE64(name:password)
```

**Response for non-authenticated users**

```
HTTP/1.1 401 Authorization Required
WWW-Authenticate: Basic realm="Secure Area"
```

## HTTP Digest Authentication

**Request**

```
Authorization: Digest username="Mufasa",
  realm="testrealm@host.com",
  nonce="dcd98b7102dd2f0e8b11d0f600bfb0c093",
  uri="/dir/index.html",
  qop=auth,
  nc=00000001,
  cnonce="0a4f113b",
  response="6629fae49393a05397450978507c4ef1",
  opaque="5ccc069c403ebaf9f0171e9517f40e41"
```

**Response for non-authenticated users**

```
WWW-Authenticate: Digest realm="testrealm@host.com",
  qop="auth",
  nonce="dcd98b7102dd2f0e8b11d0f600bfb0c093",
  opaque="5ccc069c403ebaf9f0171e9517f40e41"
```

**Digest**

```
HA1 = MD5(username:realm:password)
HA2 = MD5(method:digestURI)
response = MD5(HA1:nonce:nonceCount:clientNonce:qop:HA2)
```
