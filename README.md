# key-value

Current repository is a homework given by a recruiter.

# Programming Task
The following task should be completed using Go and shouldn't take any longer than 60 mins.
## The Task
Please create an HTTP server that exposes a simple in-memory key value store. The key value store should support keys consisting of a string and values consisting of bytes.
## Additional Considerations
- Please design and implement this service as if you were building it for a real production environment.
- As a result, please consider how this service will be maintained, improved or enhanced in future. Are there approaches you can take that will make future alterations to the service easier (such as making the key value store persistent)?
- You may use any third-party packages you feel are required to complete the task, however, the use of any third-party package should be justifiable.

## Run

```commandline
go run cmd/keyvalue -- --port=8081
```

## Usage

### Simple store string

```commandline
curl -X PUT --url http://localhost:8081/store/go -d "Go, go faster!"
curl -X GET --url http://localhost:8081/store/go
```

### Store file

```commandline
curl -X PUT --data-binary @README.md http://localhost:8081/store/readme
curl -X GET http://localhost:8081/store/readme | less
curl -X DELETE http://localhost:8081/store/readme
curl -i -X GET http://localhost:8081/store/readme 
```