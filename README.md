# Benthos Vader

[Benthos](https://benthos.dev) with sentiment analysis support based on VADER
(Valence Aware Dictionary and sEntiment Reasoner).

This project adds a `sentiment()`
[bloblang method](https://www.benthos.dev/docs/guides/bloblang/methods) to
Benthos which can be called on any string field within a bloblang mapping.

Wraps the [github.com/jonreiter/govader](https://github.com/jonreiter/govader)
package. See its [docs](https://pkg.go.dev/github.com/jonreiter/govader) for
more information.

## Build instructions

```shell
> docker build -t mihaitodor/benthos-vader:latest .
```

## Run instructions

Create a file called `config.yaml` with the following contents:

```yaml
input:
  stdin: {}

pipeline:
  processors:
    - bloblang: root = content().sentiment()

output:
  stdout: {}
```

Run the `mihaitodor/benthos-vader` Docker container using this config:

```shell
> docker run --rm -it -v $(pwd)/config.yaml:/etc/benthos/config.yaml mihaitodor/benthos-vader:latest -c /etc/benthos/config.yaml
```

Type any sentence and hit enter.
