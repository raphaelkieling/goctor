# Goctor

> THIS PROJECT IS UNDER **DEVELOPMENT**. THIS IS **NOT READY** TO BE USED.

<div style="width:100%; display: flex; justify-content: center;">
    <img src="./temp-logo.png" width="150px">
</div>

The purposal is standarize how we setup the environment for any kinda of project. Why?

- Facilitate a new dev inside the project.
- Facilitate give tech support for the project.
- Avoid the necessity to onboarding session just to configure the project.
- Avoid put texts and texts inside the readme only for the setup.

## Setup

Install using:

```sh
...
```

Create a yml file inside your project:

```yml
version: "1"
banner:
    title: Setup Project
exams:
    - name: Application
      description: Necessary to start the application
      steps:
        - name: NodeJS (version 16.0.0)
          run: exit 1
    - name: Postgres
      description: To start the app database
      steps:
        - name: Docker
          run: exit 1
        - name: Docker compose
          run: exit 0
```

Run: `goctor -f ./example/goctor.yml`

![](./run-example-setup.png)

## Inpiration

The mainly inspiration: https://docs.flutter.dev/get-started/install/windows#run-flutter-doctor
I only used Flutter one time and i felt in love with the flutter doctor.

## How use

Today we have tree layers in the cli.

```sh
[] - Exam
  - Step 1
    - Possibility 1
    - Possibility 2
  - Step 2
    - Possibility 1
  - Step 3
```

The ideia is ever give the step by step for the new developers. Example of a real output:

```sh
[✔️] - Application
  [✔️] Nodejs
  [✔️] NPM or Yarn
[❕] - Postgres
  [✔️] Docker
  [❕] Docker Compose
    - Try run the `brew install docker-compose`
  [❕] Datbase runnin
    - Try run the `docker-compose up`
```