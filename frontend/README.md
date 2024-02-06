## Prerequisites

Make sure you have Node.js and npm (Node Package Manager) installed on your computer. If not, you can download and install them from https://nodejs.org/.
Instructions

1. Clone the Repository:

`git clone https://github.com/ecce75/social-network.git`


2. Navigate to the Frontend Directory:

`cd frontend>`

3. Install Dependencies:

`npm install`

This command installs all the necessary dependencies for your Next.js app.

4. Run the Development Server:

`    npm run dev`

This command starts the development server. You can also use yarn dev, pnpm dev, or bun dev based on your preferred package manager.

5. Access the Application:

Open your web browser and navigate to http://localhost:3000 to view your Next.js application.

Note:

    The instructions assume that the backend server (Go server) is already running at http://localhost:8080. If it's not, you'll need to start the Go server before running the Next.js app.

## Components

### LoginForm

**Purpose**: Handles user login.

**Dependencies**: Formik for form management.

**Key Features**:
    Two input fields: username and password.
    Submits a POST request to http://localhost:8080/api/users/login upon form submission.

### RegisterForm

**Purpose**: Manages user registration.

**Dependencies**: Formik for form management.

**Key Features**:

    Input fields for email, password, name, date of birth, avatar/image (with file upload), username, and "about me".
    Submits a POST request to http://localhost:8080/api/users/register upon form submission.