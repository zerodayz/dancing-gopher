## Deployment

```bash
oc new-app https://github.com/zerodayz/dancing-gopher --name=dancing-gopher
```

Wait for the deployment to finish
```
$ oc get pods
NAME                     READY   STATUS    RESTARTS   AGE
dancing-gopher-1-build   1/1     Running   0          36s
```

And 
```
$ oc get pods
NAME                             READY   STATUS      RESTARTS   AGE
dancing-gopher-1-build           0/1     Completed   0          65s
dancing-gopher-bf7d97797-fxbtt   1/1     Running     0          25s
```

## Expose Svc

```
oc expose svc/dancing-gopher --port=8080
oc create route passthrough dancing-gopher-tls --service=dancing-gopher --port=8443
```

## Visit HTTP or HTTPS

```
oc get routes
```
