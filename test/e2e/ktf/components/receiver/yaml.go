package receiver

const (
	podTemplate = `
apiVersion: v1
kind: Pod
metadata:
  name: receiver
  labels: 
    app: receiver
spec:
  containers:
  - name: receiver
    image: github.com/google/knative-gcp/test/test_images/receiver
    env:
    - name: FIRST_N_ERRS
      value: "0"
`

	serviceTemplate = `
apiVersion: v1
kind: Service
metadata:
  name: receiver
spec:
  ports:
  - port: 80
    protocol: TCP
  selector:
    app: receiver
`
)
