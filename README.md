# Wetherly: The Mostly Harmless Syslog Server

Welcome, intrepid traveler of the digital cosmos, to **Wetherly**, a simple yet delightfully whimsical syslog server. Inspired by the works of Douglas Adams, this project aims to bring a touch of humor and a dash of practicality to the often mundane world of system logging. 

[Read the full blog post about building this project](https://deployharmlessly.dev/building-a-golang-syslog-server-a-journey-through-the-digital-cosmos) for a deeper dive into its creation and design.

## Table of Contents

1. [Introduction](#introduction)
2. [Installation](#installation)
3. [Usage](#usage)
4. [Development](#development)
5. [Contributing](#contributing)
6. [License](#license)

## Introduction

In the vast expanse of the universe, where the improbable is commonplace and the mundane is often overlooked, lies a small, unassuming syslog server known as **Wetherly**. This server, much like the planet Earth, is mostly harmless. It listens patiently on port 6601, eagerly awaiting messages from across the galaxy, which it then displays with the enthusiasm of a Vogon poet at a recital.

## Technologies Used

As you embark on your journey with Wetherly, you'll find it powered by a constellation of technologies, each playing its part in this cosmic symphony:

- **Go**: The language of choice for Wetherly, Go is as efficient as a well-tuned spaceship, offering simplicity and concurrency to handle the vastness of syslog messages with ease.

- **Docker**: Think of Docker as your trusty spaceship, encapsulating Wetherly and its dependencies into a neat container, ensuring it runs smoothly across the galaxy of different environments.

- **Task**: This task runner is your autopilot, automating the myriad of tasks needed to keep Wetherly shipshape. With `Taskfile.yml` as its guide, it ensures everything runs like clockwork.

- **Helm**: For those venturing into the Kubernetes nebula, Helm is your chart, guiding the deployment of Wetherly with precision and ease, managing even the most complex Kubernetes constellations.

- **Netcat (nc)**: Often dubbed the "Swiss Army knife" of networking, Netcat is your communicator, sending test messages to Wetherly with the finesse of a sub-etha signaling device.

- **Golangci-lint**: While optional, this linter is like having a meticulous proofreader on board, ensuring your Go code is as polished as a Vogon constructor fleet.


## Installation

Fear not, for installing Wetherly is as easy as hitching a ride on a passing spaceship.

Before you begin, ensure you have the following tools installed on your system:

- **Go**: [Install Go](https://golang.org/doc/install)
- **Docker**: [Install Docker](https://docs.docker.com/get-docker/)
- **Task**: [Install Task](https://taskfile.dev/#/installation)
- **Helm**: [Install Helm](https://helm.sh/docs/intro/install/)
- **Netcat (nc)**: Available via most package managers (e.g., `apt`, `brew`, `yum`)
- **Golangci-lint**: (Optional) [Install Golangci-lint](https://golangci-lint.run/usage/install/)

Then simply follow these steps:

1. **Clone the Repository:**

    ```bash
    git clone https://github.com/jleski/wetherly.git
    cd wetherly
    ```

2. **Setup Development Environment:**

    ```bash
    task dev:setup
    ```

3. **Build the Docker Image:**

    ```bash
    task docker:build
    ```

4. **Run the Docker Container:**

    ```bash
    docker run -p 6601:6601 jleski/wetherly:latest
    ```

And just like that, you're ready to embark on your syslog journey!

## Usage

Once Wetherly is up and running, you can send it messages using a variety of methods. Here are a few examples to get you started:

* **Send a Test Message:**

    ```bash
    task test:send
    ```

* **Send an RFC5424 Formatted Message:**

    ```bash
    task test:send:rfc5424
    ```

* **Stress Test the Server:**

    ```bash
    task test:send:stress
    ```

## Development

Embark on your development journey with Wetherly, equipped with a comprehensive set of tasks to guide you through the cosmos of code. Here's a quick summary of the tasks available to you:

| Task                  | Description                                                   |
|-----------------------|---------------------------------------------------------------|
| **help**              | Display available tasks.                                      |
| **deps**              | Install development dependencies.                             |
| **fmt**               | Format Go code.                                               |
| **lint**              | Run linters.                                                  |
| **test**              | Run tests.                                                    |
| **build**             | Build the application.                                        |
| **run**               | Run the application locally.                                  |
| **clean**             | Clean build artifacts.                                        |
| **docker:build**      | Build Docker image.                                           |
| **docker:push**       | Push Docker image to registry.                                |
| **helm:lint**         | Lint Helm chart.                                              |
| **helm:template**     | Template Helm chart.                                          |
| **helm:install**      | Install/Upgrade Helm release.                                 |
| **helm:uninstall**    | Uninstall Helm release.                                       |
| **dev:setup**         | Setup development environment.                                |
| **ci**                | Run CI pipeline tasks.                                        |
| **all**               | Run all tasks (format, lint, test, build).                    |
| **test:send**         | Send a test message to the syslog server using netcat.        |
| **test:send:multi**   | Send multiple test messages to the syslog server.             |
| **test:send:logger**  | Send a test message using logger (if available).              |
| **test:send:rfc5424** | Send a properly formatted RFC5424 syslog message.             |
| **test:send:stress**  | Stress test the server with many messages.                    |
| **test:all**          | Run all syslog test messages.                                 |
| **run:test**          | Run the server and send test messages.                        |

Execute these tasks using `task <task_name>`.

With these tasks at your disposal, you're well-equipped to explore and enhance Wetherly's codebase. Remember, the key to great development is a well-prepared task list.

## Contributing

Contributions to Wetherly are as welcome as a towel on a Vogon ship. Feel free to fork the repository, make your changes, and submit a pull request. Remember, the answer to life, the universe, and everything is 42, but good code is a close second.

## License

Wetherly is licensed under the MIT License. Share and enjoy!

So there you have it, the guide to Wetherly, your mostly harmless syslog server. May your logs be ever verbose, your errors be few, and your adventures be plentiful. Don't forget your towel!