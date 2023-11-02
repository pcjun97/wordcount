# wordcount

## Description

`wordcount` takes a file and prints a word count of its contents.
The output consists of a line for each word,
with the number of its occurrences in the file.
It is also sorted by the number of occurrences starting with the most frequent word.

## Installation

```sh
git clone github.com/pcjun97/wordcount
cd wordcount
go build ./cmd/wordcount -o wordcount
```

## Usage

Calling `wordcount` with a filename will give the expected outputs:

```shellsession
$ cat test.txt
foo bar bar

$ ./wordcount test.txt
bar: 2
foo: 1
```

Calling `wordcount` with a `-i` flag makes `wordcount` ignore punctuation when counting the words:

```shellsession
$ cat test.txt
foo bar bar,

$ ./wordcount test.txt
bar: 2
foo: 1
```

Calling `wordcount` with `-s` starts a HTTP server.
The port can be configured with the `-p` flag, default to `3000`:

```shellsession
$ ./wordcount -s

# in another shell
$ curl localhost:3000 -X POST -d "foo bar bar"
bar: 2
foo: 1
```

## Deploying with Kubernetes

The Docker image for `wordcount` can be found [here](https://hub.docker.com/r/pcjun97/wordcount).

### Running As a Job

The default behaviour of `wordcount` accepts a file, output the word count, and exit.
As a demonstration, this can be deploy as a Kubernetes [job](https://kubernetes.io/docs/concepts/workloads/controllers/job/).

```shellsession
$ kubectl apply -f ./deploy/kubernetes/job

$ kubectl logs job.batch/wordcount
et: 5
id: 4
in: 4
arcu: 3
at: 3
eget: 3
nunc: 3
Eget: 2
ac: 2
adipiscing: 2
...
```

### Running As a Server

Running in a server mode, `wordcount` can be deployed as a Kubernetes [deployment](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/).

```shellsession
$ kubectl apply -f ./deploy/kubernetes/deployment

# from a pod in cluster
$ curl wordcount.default -X POST -d "foo bar bar"
bar: 2
foo
```
