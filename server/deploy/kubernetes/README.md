# Platform Server in Kubernetes

Substitute these parameters in manifest:

- `<TLS-SECRET>` - name of the TLS secret in Kubernetes
- `<PLATFORM-SERVER-DNS>` - DNS name for server
- `<COMMIT-HASH-OR-TAG>` - version of the platform server

Deploy using the commands below:

```
kubectl create ns platform-server
kubectl apply -f platform-server.yaml
```