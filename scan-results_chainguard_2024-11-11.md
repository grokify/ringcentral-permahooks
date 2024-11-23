# Chainguard and Alpine Comparison using Grype

## TL;DR

This is an example comparison of using Alpine and Chainguard based images performed on Nov 11, 2024 for
the Chainguard course "Painless Vulnerability Management" Final Project.

The results show that Alpine has 2 fixed medium vulnerabilities and Chainuard as 0 vulerabilites.

1. `golang:1.23-alpine` (2 fixed medium vulnerabilites)

```
 ✔ Scanned for vulnerabilities     [2 vulnerability matches]  
   ├── by severity: 0 critical, 0 high, 2 medium, 0 low, 0 negligible
   └── by status:   2 fixed, 0 not-fixed, 0 ignored 
```

2. `FROM cgr.dev/saviynt.com/go:latest` (1.23.3) (0 vulnerabilites0)

```
 ✔ Scanned for vulnerabilities     [0 vulnerability matches]  
   ├── by severity: 0 critical, 0 high, 0 medium, 0 low, 0 negligible
   └── by status:   0 fixed, 0 not-fixed, 0 ignored 
```

## Select Application (Step 1: Set Up a Demo Application)

Use the [`ringcentral-permahooks`](https://github.com/grokify/ringcentral-permahooks) application, also
posted on [Docker Hub](https://hub.docker.com/r/grokify/ringcentral-permahooks).

## Using Alpine Image

### Dockerfile with Alpine (Step 2: Create a “traditional” Dockerfile)

Define `Dockerfile` with:

`FROM golang:1.23-alpine`

### Docker build with Alpine (Step 3: Build the Image)

```
% docker build -t ringcentral-permahooks:latest .
[+] Building 1.6s (16/16) FINISHED                                                                                                     
 => [internal] load build definition from Dockerfile                                                                              0.0s
 => => transferring dockerfile: 120B                                                                                              0.0s
 => [internal] load .dockerignore                                                                                                 0.0s
 => => transferring context: 2B                                                                                                   0.0s
 => resolve image config for docker.io/docker/dockerfile:1                                                                        0.6s
 => CACHED docker-image://docker.io/docker/dockerfile:1@sha256:865e5dd094beca432e8c0a1d5e1c465db5f998dca4e439981029b3b81fb39ed5   0.0s
 => [internal] load build definition from Dockerfile                                                                              0.0s
 => [internal] load metadata for docker.io/library/golang:1.23-alpine                                                             0.6s
 => [internal] load .dockerignore                                                                                                 0.0s
 => [1/7] FROM docker.io/library/golang:1.23-alpine@sha256:09742590377387b931261cbeb72ce56da1b0d750a27379f7385245b2b058b63a       0.0s
 => [internal] load build context                                                                                                 0.0s
 => => transferring context: 82B                                                                                                  0.0s
 => CACHED [2/7] WORKDIR /app                                                                                                     0.0s
 => CACHED [3/7] COPY go.mod ./                                                                                                   0.0s
 => CACHED [4/7] COPY go.sum ./                                                                                                   0.0s
 => CACHED [5/7] RUN go mod download                                                                                              0.0s
 => CACHED [6/7] COPY *.go ./                                                                                                     0.0s
 => CACHED [7/7] RUN go build -o /ringcentral-permahooks                                                                          0.0s
 => exporting to image                                                                                                            0.0s
 => => exporting layers                                                                                                           0.0s
 => => writing image sha256:dd69fa369a35c66a0eb974e56a8b8282410e5ac8fb158b84d77b3daf4fde3e5b                                      0.0s
 => => naming to docker.io/library/ringcentral-permahooks:latest                                                                  0.0s
```

### Grype Vulnerabilities with Alpine (Step 4: Scan official image with Grype)

```
% grype ringcentral-permahooks:latest            
 ✔ Loaded image                                                                                         ringcentral-permahooks:latest
 ✔ Parsed image                                               sha256:dd69fa369a35c66a0eb974e56a8b8282410e5ac8fb158b84d77b3daf4fde3e5b
 ✔ Cataloged contents                                                69fae45fead9cf1cbb303828ef679be8510df9a5a575c421ff6d42c01faee39c
   ├── ✔ Packages                        [109 packages]  
   ├── ✔ File digests                    [251 files]  
   ├── ✔ File metadata                   [251 locations]  
   └── ✔ Executables                     [66 executables]  
 ✔ Scanned for vulnerabilities     [2 vulnerability matches]  
   ├── by severity: 0 critical, 0 high, 2 medium, 0 low, 0 negligible
   └── by status:   2 fixed, 0 not-fixed, 0 ignored 
NAME        INSTALLED  FIXED-IN  TYPE  VULNERABILITY  SEVERITY 
libcrypto3  3.3.2-r0   3.3.2-r1  apk   CVE-2024-9143  Medium    
libssl3     3.3.2-r0   3.3.2-r1  apk   CVE-2024-9143  Medium
```

## Using Chainguard Image

### Dockerfile with Chainguard

Define `Dockerfile` with:

`FROM cgr.dev/saviynt.com/go:latest` (1.23.3)

### Docker build with Chainguard (Step 5: Build the Chainguard Image)

```
% docker build -t ringcentral-permahooks-wolfi:latest .
[+] Building 1.4s (16/16) FINISHED                                                                                                     
 => [internal] load build definition from Dockerfile                                                                              0.0s
 => => transferring dockerfile: 120B                                                                                              0.0s
 => [internal] load .dockerignore                                                                                                 0.0s
 => => transferring context: 2B                                                                                                   0.0s
 => resolve image config for docker.io/docker/dockerfile:1                                                                        0.6s
 => CACHED docker-image://docker.io/docker/dockerfile:1@sha256:865e5dd094beca432e8c0a1d5e1c465db5f998dca4e439981029b3b81fb39ed5   0.0s
 => [internal] load build definition from Dockerfile                                                                              0.0s
 => [internal] load metadata for cgr.dev/chainguard/go:latest                                                                     0.4s
 => [internal] load .dockerignore                                                                                                 0.0s
 => [internal] load build context                                                                                                 0.0s
 => => transferring context: 82B                                                                                                  0.0s
 => [1/7] FROM cgr.dev/chainguard/go:latest@sha256:266990d29364a6d64611877e07e12effe5468a3421dbbd3d61c93067ad5581b0               0.0s
 => CACHED [2/7] WORKDIR /app                                                                                                     0.0s
 => CACHED [3/7] COPY go.mod ./                                                                                                   0.0s
 => CACHED [4/7] COPY go.sum ./                                                                                                   0.0s
 => CACHED [5/7] RUN go mod download                                                                                              0.0s
 => CACHED [6/7] COPY *.go ./                                                                                                     0.0s
 => CACHED [7/7] RUN go build -o /ringcentral-permahooks                                                                          0.0s
 => exporting to image                                                                                                            0.0s
 => => exporting layers                                                                                                           0.0s
 => => writing image sha256:812caec16a6a7d71b54b9c599b72e0846aec717dd130483a480ac9cdc7bf7faf                                      0.0s
 => => naming to docker.io/library/ringcentral-permahooks-wolfi:latest                                                            0.0s
```

### Grype Vulnerabilities with Chainguard (Step 6: Scan the Chainguard Image)

```
% grype ringcentral-permahooks-wolfi:latest  
 ✔ Loaded image                                                                                   ringcentral-permahooks-wolfi:latest
 ✔ Parsed image                                               sha256:812caec16a6a7d71b54b9c599b72e0846aec717dd130483a480ac9cdc7bf7faf
 ✔ Cataloged contents                                                ef5066c34f3d3a2a1745ff9c940d23b686bd6ab0328b5d653c234ab4168b48c1
   ├── ✔ Packages                        [160 packages]  
   ├── ✔ File digests                    [9,116 files]  
   ├── ✔ File metadata                   [9,116 locations]  
   └── ✔ Executables                     [218 executables]  
 ✔ Scanned for vulnerabilities     [0 vulnerability matches]  
   ├── by severity: 0 critical, 0 high, 0 medium, 0 low, 0 negligible
   └── by status:   0 fixed, 0 not-fixed, 0 ignored 
No vulnerabilities found
```
