# DOAH AI in GO

## HEROKU
see `https://godoahai-d846d75541aa.herokuapp.com/`
GET https://godoahai-d846d75541aa.herokuapp.com/status

curl -X POST https://godoahai-d846d75541aa.herokuapp.com/process \
  -H "Content-Type: application/json" \
  -d '{"data": "some important payload", "id": 12345}'


Deploy: `git push heroku main`


## Terminal notes
501  git add .
502  git commit -m "Initial commit: Go API server with /status and /process endpoints"
503  git branch -M main
504  git remote add origin https://github.com/mattkc7/go_doah_ai.git
505  git push -u origin main
506  heroku git:remote -a godoahai
507  heroku git:remote -a godoahai
508  git push heroku main


// GET http://localhost:8080/status

// curl -X POST http://localhost:8080/process \
//   -H "Content-Type: application/json" \
//   -d '{"data": "some important payload", "id": 12345}'
