# IrieSphere Backend

## Table of Contents

- [Backend Structure](#backend-structure)
  - [pkg](#pkg)
    - [db](#db)
    - [model](#model)
    - [repository](#repository)
    - [handler](#handler)
  - [api](#api)
  - [util](#util)
- [Functionalities](#functionalities)
  - [Session](#session)
  - [Posts](#posts)
  - [Comments](#comments)

## Backend Structure

### pkg

This directory contains the core logic of your application and is where most of your code will reside.

#### db

Contains database setup, connection logic, and migrations. Keeping migrations close to your database logic makes sense since they are tightly coupled.

#### model

Defines the data structures used by your application. These are often representations of your database tables but can also include other domain-specific logic.

#### repository

Acts as the data access layer. Repositories use models to interact with the database, encapsulating the logic needed to query or update the database.

#### handler

Contains the business logic of your application. Handlers call into repositories to fetch and store data, perform operations, and implement application-specific rules. Handlers should also validate and sanitize the input data.

### api

This is where you define your HTTP handlers and routing. It's the layer that interacts with the outside world, translating HTTP requests into actions on your services and ultimately your models and database.

### util

A place for utility functions that don't naturally fit elsewhere. These might include helper functions for formatting, validation, or small tasks used across multiple parts of the application.

## Functionalities

### Session

- **User Registration**: Endpoint `/api/users/register` (POST)
- **User Logout**: Endpoint `/api/users/logout` (POST)
- **User Login**: Endpoint `/api/users/login` (POST)
- **Check User Authentication**: Endpoint `/api/users/check-auth` (GET)

### Posts

- **Create Post**: Endpoint `/post` (POST)
- **Delete Post**: Endpoint `/post/{id}` (DELETE)
- **Update Post**: Endpoint `/post/{id}` (PUT)

### Comments

- **Get Comments**: Endpoint `/post/{id}/comments` (GET)
- **Create Comment**: Endpoint `/comment` (POST)
- **Delete Comment**: Endpoint `/comment/{id}` (DELETE)
