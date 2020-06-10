## Prerequisite 
* Go -- 1.12+ recommended
* Google Cloud Platform CLI (gcloud) -- if you want to deploy to GCP
* Glide -- recommended if you want to deploy to GCP easily

## Try it Online with Google Cloud Platform 
Enjoy! https://davic-helpers-dot-wfchiang-dev.uc.r.appspot.com/davic-helpers 

## Run it Locally 
```
go run main.go
``` 

It will pickup the environment variable **PORT** for the port. 
The default is 8080.

## Deploy to Google Cloud Platform 
#### (1) Gather the Dependenciese Using Glide 
```
glide create
```

Then 

```
glide install 
```

#### (2) Create Go module 
```
go mod init github.com/wfchiang/davic-helpers
```

#### (3) Deploy 
```
gcloud app deploy
```

#### (optional) Cleanup/Reset Glide-installed Dependencies 
```
rm glide.*
```

Then

```
rm -rf vendor
```