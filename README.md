# Social-Network

## Table of Contents
- ### [General Information](#general-information)
- ### [Technologies](#technologies)
- ### [Setup](#setup)
- ### [Usage](#usage)
- ### [Authors](#authors)

## General Information

This project involves the creation of a feature-rich social networking platform, inspired by popular social networks like Facebook. The platform will encompass a wide range of functionalities, including user registration and login, the ability to create posts, engage in commenting, send private messages, manage followers, and interact within groups. Key technologies utilized in this project include SQLite for data storage, Golang for backend development, JavaScript for frontend interactivity, HTML for structuring web pages, and CSS for styling. A single-page application approach will be adopted for a seamless user experience, with real-time features powered by WebSockets.

The project contains the following features:

- Friends
- Profile
- Posts
- Groups
- Notifications
- Chats

## Technologies
- ### [TypeScript](https://www.typescriptlang.org/)
- ### [Nextjs](https://nextjs.org/)
- ### [Golang](https://go.dev/)
- ### [HTML](https://www.w3.org/html/)
- ### [CSS](https://developer.mozilla.org/en-US/docs/Web/CSS)
- ### [SQLite](https://sqlite.org/index.html)

## Setup
<!-- TODO -->
Clone the repository
```
git clone https://github.com/ecce75/social-network.git
```

## Usage
1. Run backend:
```
cd backend/
```
```
go run .
```

2. Open a new terminal and move to frontend directory:
```
cd frontend/
```
3. Install the required modules:
```
npm install
```
wait a minute or two for all the dependencies to install

4. Run frontend:
```
npm run dev
```
5. Test it from the [http://localhost:3000/](http://localhost:3000/)

Users for testing:
- mark
- marek
- karlo
- james

Password is same as username for all.

For the docker, run:
```
make docker-run
```
To stop the docker, use:
```
make docker-remove
```

## Authors
- Andrei Tuhkru
- Robin Rattasepp
- 
