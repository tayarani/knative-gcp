package target

const (
	podTemplate = `
apiVersion: v1
kind: Pod
metadata:
  name: {{.Name}}
  labels: 
    app: {{.Name}}	
spec:
  containers:
    - name: {{.Name}}
      image: {{.Image}}
`

	serviceTemplate = `
apiVersion: v1
kind: Service
metadata:
  name: {{.Name}}
spec:
  ports:
  - port: 80
    protocol: TCP
  selector:
    app: {{.Name}}
`
)
