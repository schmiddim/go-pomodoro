# go-pomodoro

`go-pomodoro` is a command-line application that helps you manage your time using the Pomodoro Technique. You can set
the duration of your work intervals and the application will notify you when it's time to take a break. It also
integrates with an API to track your accomplishments.

## Installation
### Go install
```shell
 go install github.com/schmiddim/go-pomodoro@latest
```

### Compile
To install `go-pomodoro`, you need to have Go installed on your machine. Once you have Go installed, you can clone this
repository and build the application:

```shell 
git clone https://github.com/schmiddim/go-pomodoro.git 
cd go-pomodoro 
go build -o go-pomodoro
```

## Usage

You can run `go-pomodoro` with the following command:

```shell    
./go-pomodoro --time=<minutes> --habitId=<habitid>
```

Where:

- `<minutes>` is the duration of the timer in minutes.
- `<habitId>` is the ID of the habit you are tracking.

For example:

```shell
./go-pomodoro --time=30 --habitId=16
```

## Environment Variables

`go-pomodoro` uses the following environment variables:

- `POMODORO_API_KEY`: Your API key for the accomplishments tracking API.
- `POMODORO_ENDPOINT`: The endpoint for the accomplishments tracking API.

You can set these environment variables in your shell:
```shell
export POMODORO_API_KEY=your_api_key export POMODORO_ENDPOINT=http://localhost:8080/api/v1
```

## Create a quick Release 
```shell
TAG=0.0.17 ;git commit -am "Version Info"; git tag  $TAG; git push origin $TAG
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)