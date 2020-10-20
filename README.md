![Main](https://github.com/dhenkel92/pod-helper/workflows/Main/badge.svg)
![Release](https://github.com/dhenkel92/pod-helper/workflows/Release/badge.svg)


<!-- PROJECT LOGO -->
<br />
<p align="center">
  <h3 align="center">Kubernetes Pod Helper</h3>

  <p align="center">
    A tool to easily operate on multiple pods at the same time!
    <br />
    <br />
    <a href="https://github.com/dhenkel92/pod-helper/issues">Report Bug</a>
    ·
    <a href="https://github.com/dhenkel92/pod-helper/issues">Request Feature</a>
  </p>
</p>



<!-- TABLE OF CONTENTS -->
## Table of Contents

* [About the Project](#about-the-project)
* [Prerequisites](#prerequisites)
* [Installation](#installation)
* [Usage](#usage)
  * [Run](#run)
  * [Logs](#logs)
* [Contributing](#contributing)
* [License](#license)
* [Contact](#contact)
* [Acknowledgements](#acknowledgements)



<!-- ABOUT THE PROJECT -->
## About The Project

![Product Name Screen Shot](images/pod_helper_example.gif)

In my day to day job i often need to debug pods in a kubernetes cluster or need to get logs from all of them. However, with the default tools, provided by kubernetes, it's super hard and time consuming to do it.
This is why I decided to build a small helper tool in order to keep us productive and less annoyed.

I've tried to solve the following problems:
* Get Logs from multiple containers in one stream
* Execute a command in multiple pods
* Follow Logs of multiple pods (WIP)
* Exec into multiple pods with an interactive shell (WIP)

## Prerequisites

### Runtime

* A functional Kubernetes cluster (tested with >= 1.17)

### Development

* [Golang](https://golang.org/) >= 1.15
* [pre-commit](https://pre-commit.com/) >= 2.0

## Installation

### MacOS

You can install the tool via Homebrew and the tap repository can be found [here.](https://github.com/dhenkel92/homebrew-tap)
```
brew install dhenkel92/tap/pod-helper
```

In order to get a newer version, just upgrade via Homebrew
```
brew upgrade dhenkel92/tap/pod-helper
```


### Other distributions

See the [Releases page](https://github.com/dhenkel92/pod-helper/releases) for a list of Peek packages for various distributions.


## Usage


### General

Use different .kubeconfig file than `~/.kube/config`
```shell script
pod-helper --kubeconfig /kube/config logs
or short
pod-helper -config /kube/config logs
```

### Run

Run `ls -al` in all pods of the default namespace
```shell script
pod-helper --namespace default run --command "ls -al"
or short
pod-helper -n default run -c "ls -al"
```

Run `ls -al` in all the pods
```shell script
pod-helper run --comand "ls -al"
or short
pod-helper run -c "ls -al"
```

Override the entrypoint for the command execution (default `/bin/sh -c`)
```shell script
pod-helper run --entrypoint "/bin/bash -c" --comand "ls -al"
or short
pod-helper run -e "/bin/bash -c" -c "ls -al"
```

Select pod by label selector in default namespace
```shell script
pod-helper --namespace default --label app=nginx run --command "ls -al"
or short
pod-helper -n default -l app=nginx run --command "ls -al"
```

Run a command on all the pods in the namespace default and kube-system
```shell script
pod-helper --namespace default --namespace kube-system --label app=nginx run --command "ls -al"
or short
pod-helper -n default -n kube-system -l app=nginx run --command "ls -al"
```

### Logs

Get all logs from all the pods in the cluster
```shell script
pod-helper logs
or short
pod-helper logs
```

Get all logs from all pods in the kube-system namespace
```shell script
pod-helper --namespace kube-system logs
or short
pod-helper -n kube-system logs
```

Get the last ten log entries of all the pods
```shell script
pod-helper logs --tail 10
or short
pod-helper logs -t 10
```

Get logs of the first container in all pods in the cluster
```shell script
pod-helper --container-index 0 logs
or short
pod-helper -ci 0 logs
```

Get logs of containers mit name `nginx`.
```shell script
pod-helper --container nginx logs
or short
pod-helper -con nginx logs
```

## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/amazing_feature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/amazing_feature`)
5. Open a Pull Request


## License

Distributed under the MIT License. See `LICENSE` for more information.


## Contact

Daniel Henkel - daniel@henkel.tech

Project Link: [https://github.com/dhenkel92/pod-helper](https://github.com/dhenkel92/pod-helper)


## Acknowledgements
* [pre-commit](https://pre-commit.com/)
* [Goreleaser](https://goreleaser.com/)
* [Cli](https://github.com/urfave/cli)
* [Aurora](https://github.com/logrusorgru/aurora)
* [Choose an Open Source License](https://choosealicense.com)
* [README template](https://github.com/othneildrew/Best-README-Template)
