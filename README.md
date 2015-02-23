# Quill http://filepreviews.io/ Client

## Purpose
Provides a simple go client for File Previews (http://filepreviews.io/).

## Install

```
go get github.com/quillcontent/go-filepreview
```

```
svc, err := url.Parse("https://api.filepreviews.io/v1/")
f := "http://linktofiletopreview..."

c := filepreview.ServiceConfig{}
c.EndPoint = *svc
c.ApiKey = "yourkey"

res, err := GeneratePreview(f, c)
```
