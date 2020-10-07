package sender

const (
	podTemplate = `
apiVersion: v1
kind: Pod
metadata:
  name: sender
  labels: 
    app: sender
spec:
  containers:
  - name: sender
    image: github.com/google/knative-gcp/test/test_images/sender
    #env:
    #- name: FIRST_N_ERRS
    #  value: "0"
`

	serviceTemplate = `
apiVersion: v1
kind: Service
metadata:
  name: sender
spec:
  ports:
  - port: 80
    protocol: TCP
  selector:
    app: sender
`
)
