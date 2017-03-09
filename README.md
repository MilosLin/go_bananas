# go_bananas
golang practice project

## Build binary file with version control

Exec command like this:

    go build -ldflags "-X github.com/MilosLin/go_bananas/cmd.version=0.0.1 -X github.com/MilosLin/go_bananas/cmd.date=`date -u +'%Y-%m-%dT%H:%M:%S'`"

than

    ./go_bananas version
    version=0.0.1, date=2017-03-09T07:32:25

## Configuration

all config are in file `config.json`

## Command

#### show version

./go_bananas version

#### run task

./go_bananas task --name={task name} --argu={task's argu}
