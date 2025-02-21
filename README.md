# vote-logger
 Lightweight microservice to log votes of various botlists written in go.


# Local Deployent

Fill out the .env.example file!

To start run: (Might only work on unix)
```bash
env $(cat .env | xargs) go run .
```