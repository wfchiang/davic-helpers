
## Installation 
I assume that you already have **Davic** installed. If not, please install Davic first. 
In addition to Davic, you will also need **Gorilla/mux** framework. 
Executing the python script of this directory should help installing it: 
```
python myinstall.py
```

## Run
Assume that you are playing this sample app locally... 

#### Start the Server 
In command-line:  
```
go run main.go
```

## Deploy to Google Cloud Platform 
#### (1) Create Go module 
```
go mod init github.com/wfchiang/davic-helpers
```

#### (2) Deploy 
```
gcloud app deploy
```

## Glide Usage
#### Install
```
glide create
```
```
glide install
```

#### Cleanup
```
rm glide.*
```
```
rm -rf vendor
```