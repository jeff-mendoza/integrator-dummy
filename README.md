![golang logo](golang_logo.png)

# Integrator Dummy

## Getting started

**Step 1:** Install Golang
Make sure you have Go 1.16 or higher installed.
https://golang.org/doc/install

**Step 2:** Environment Config
Set-up the standard Go environment variables according to latest guidance (see https://golang.org/doc/install#install).

**Step 3:** Install Dependencies
From the project root, run:
```bash
go mod tidy
```

**Step 4:** Set an environment variable called `SECRET` with the value of your `CLIENT SECRET` that you have in your application created in the Mercado Pago account.

**Step 5:** Run project
```bash
go run main.go
```

## Deploy to Heroku
You can also deploy this app to Heroku:

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

**Step 1:** Create a [Heroku](https://signup.heroku.com/dc) account.

**Step 2:** Install [Heroku CLI](https://devcenter.heroku.com/articles/getting-started-with-go#set-up).

**Step 3:** Use the `heroku login` command to log in to the Heroku CLI.
```bash
heroku login -i
```

**Step 4:** Clone the repository.
```bash
git clone git@github.com:jeff-mendoza/integrator-dummy.git
cd integrator-dummy
```

**Step 5:** Create an app on Heroku, which prepares Heroku to receive your source code.
```bash
heroku create
```

**Step 6:** Now deploy your code
```bash
git push heroku main
```

**Step 7:** Set an environment variable called `SECRET` with the value of your` CLIENT SECRET` that you have in your application created in the Mercado Pago account.
```bash
heroku config:set SECRET=MY_CLIENT_SECRET
```

**Step 8:** As a handy shortcut, you can open the website as follows:
```bash
heroku open
```
