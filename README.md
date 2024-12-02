# Template

This template meant to be my personal use. However, you can use it if you desired. I structure this project based on the concept of **Clean Architecture** and **Domain Driven Development (DDD)**. Note that this template is still improving, changes are to be made in the future. 

## Installation

- Clone this repository by running  
  
  ```shell
  git clone git@github.com:lutffmn/template.git
  ```

- Set up your apps configuration in ``config.yaml`` file



## Initialize Project

In the root directory, run:

```shell
make init
```

Then, input your project name.

**Note: This action will also installing some dependencies.**



## Build

In the root directory, run:

```shell
make build
```

This action will build the program, and put the compiled binary in ``bin/`` directory.



## Run

There's 2 options to run the project:

- Run the compiled binary

- Run the program

### Run the Compiled Binary

In the root directory, run:

```shell
make run
```

### Run the Program

In the root directory, run:

```shell
make dev
```



## Catch a bug?

If you spot a bug or errors in this template, please make an **Issues**. I will check it, and solve it. If you want to contribute to this project, please make a **Pull Request** and mention the **Issues** related.
