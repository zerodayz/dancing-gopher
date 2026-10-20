## Deployment

```bash
oc new-app https://github.com/zerodayz/dancing-gopher --name=dancing-gopher
oc expose svc/dancing-gopher --port=8080
```
