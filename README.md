# gae-spreadsheet-issue

The repository created to demonstrate the issue with access to google spreadsheet API from Google App Engine.

## problem statement

I'm trying to access google spreadsheet via API from the application running on [Google App Engine Go 1.11 Standard Environment](https://cloud.google.com/appengine/docs/standard/go111/).
Unfortunately, the application cannot read [this spreadsheet](https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit).

I'm getting next error on [`Spreadsheets.Values.Get`](https://godoc.org/google.golang.org/api/sheets/v4#SpreadsheetsValuesService.Get) call:

```
googleapi: Error 403: Request had insufficient authentication scopes., forbidden
```

## steps to reproduce

1) deploy app: ``gcloud app deploy``
2) open in a browser (you will get 502): ``gcloud app browse``
3) check logs: ``gcloud app logs read``

```
2018-12-11 21:44:56 default[20181211t134352]  "GET / HTTP/1.1" 502
2018-12-11 21:44:57 default[20181211t134352]  2018/12/11 21:44:57 [DEBUG] metadata scopes: https://www.googleapis.com/auth/appengine.apis
2018-12-11 21:44:57 default[20181211t134352]  https://www.googleapis.com/auth/cloud-platform
2018-12-11 21:44:57 default[20181211t134352]  https://www.googleapis.com/auth/cloud_debugger
2018-12-11 21:44:57 default[20181211t134352]  https://www.googleapis.com/auth/devstorage.full_control
2018-12-11 21:44:57 default[20181211t134352]  https://www.googleapis.com/auth/logging.write
2018-12-11 21:44:57 default[20181211t134352]  https://www.googleapis.com/auth/monitoring.write
2018-12-11 21:44:57 default[20181211t134352]  https://www.googleapis.com/auth/trace.append
2018-12-11 21:44:57 default[20181211t134352]  https://www.googleapis.com/auth/userinfo.email
2018-12-11 21:44:57 default[20181211t134352]  .
2018-12-11 21:44:57 default[20181211t134352]  2018/12/11 21:44:57 Listening on port 8081
2018-12-11 21:44:58 default[20181211t134352]  2018/12/11 21:44:58 Unable to retrieve data from sheet: googleapi: Error 403: Request had insufficient authentication scopes., forbidden
```

Could someone please help to understand how to fix it?

Code represented in the `main.go`. This code is a compilation from next tutorials:
1) [GAE example for Go 1.11 runtime](https://github.com/GoogleCloudPlatform/golang-samples/blob/master/appengine/go11x/helloworld/helloworld.go)
2) [Go quickstart for spreadsheets api v4](https://developers.google.com/sheets/api/quickstart/go)
3) [Setting Up Authentication for Server to Server Production Applications](https://cloud.google.com/docs/authentication/production#auth-cloud-implicit-go)
