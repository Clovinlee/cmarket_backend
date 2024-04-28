# CMarket Backend With Go Gin

Backend of ECommerce Catalogue Page Mock Up <a href="https://github.com/Clovinlee/cmarket">[CMarket]</a> using Go Gin.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)

## Introduction

This project is backend alternative for my <a href="https://github.com/Clovinlee/cmarket">[CMarket]</a> using Go and Gin (another version of the <a href="https://github.com/Clovinlee/cmarket_backend_spring">[backend]</a> is using Java Springboot).

## Features

- REST Api using Go, Gin
- Advanced queries, filters, & pagination 

## Getting Started

To get start, follow these steps below.

### Prerequisites

List any software or dependencies required to run the project. Include links or installation instructions if necessary.

- Go 1.22.2 (or higher)
- git

### Installation

Here are the instructions to install and configure the project.

1. Clone the repository:

   ```sh
   git clone https://github.com/Clovinlee/cmarket_backend

2. Init project dependency using gradle
   ```sh
   go mod download

3. Change the *.env* accordingly

4. Run the project with either of these command
   ```
   go run .
   ``` 
   or
   ```
   CompileDaemon -command="./cmarket"
   ```
**Note: Currently, I dont provide seed / migration in this backend, you can use the raw_db.sql for base data**