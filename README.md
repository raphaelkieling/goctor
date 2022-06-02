# Goctor

> THIS PROJECT IS UNDER **DEVELOPMENT**. THIS IS **NOT READY** TO BE USED.

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
version: '1'
analysis:
    - name: Check if the environment is populated
      run_sh: |
        if [ -z "MY_VAR" ]
        then
            echo "goctor:WITHOUT_MY_VAR"
        if
      possibilites:
        - code: WITHOUT_MY_VAR
          message: Try run 'source ./init.sh'
        - message: Perfect!
```

Run: `goctor`

## Inpiration
The mainly inspiration: https://docs.flutter.dev/get-started/install/windows#run-flutter-doctor
I only used Flutter one time and i felt in love with the flutter doctor.
