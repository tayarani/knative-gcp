package target

const (
	podTemplate = `
apiVersion: v1
kind: Pod
metadata:
  name: target
  labels: 
    app: target
spec:
  containers:
  - name: target
    image: github.com/google/knative-gcp/test/test_images/target
    env:
    - name: TIME
      value: "5m"
`

	serviceTemplate = `
apiVersion: v1
kind: Service
metadata:
  name: target
spec:
  ports:
  - port: 80
    protocol: TCP
  selector:
    app: target
`
)
