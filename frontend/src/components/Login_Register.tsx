"use client";

import React from "react";
import { Formik, Field, Form, FormikHelpers, useFormikContext, FieldProps } from "formik";
import "../public/styles.css";

interface LoginValues {
	username: string;
	password: string;
}

interface RegisterValues {
	email: string;
	password: string;
	first_name: string;
	last_name: string;
	dob: string;
	avatar_url: string;
	username: string;
	about: string;
}

const handleLogin = (
	values: LoginValues,
	{ setSubmitting }: FormikHelpers<LoginValues>
  ) => {
	console.log(values);
	// Handle form submission logic here
	fetch('http://localhost:8080/api/users/login', {
	  method: 'POST',
	  headers: {
		'Content-Type': 'application/json'
	  },
	  body: JSON.stringify(values)
	})
	.then(response => response.json())
	.then(data => {
	  console.log(data);
	  setSubmitting(false);
	})
	.catch(error => {
	  console.error('Error:', error);
	  setSubmitting(false);
	});
  };

const LoginForm = ({ }) => {
	return (
		<Formik
			initialValues={{
				username: "",
				password: "",
			}}
			onSubmit={handleLogin}
		>
			<Form className="flex items-center">
				<div className="mr-2">
					<Field
						className="rounded-md p-2 border border-black w-32 focus:outline-none"
						id="username"
						name="username"
						placeholder="Username"
						autoComplete="off"
					/>
				</div>

				<div className="mr-2">
					
					<Field
						className="rounded-md p-2 border border-black w-32 focus:outline-none"
						id="password"
						name="password"
						placeholder="Password"
						type="password"
						autoComplete="off"
					/>
				</div>

				<button
					type="submit"
					className="bg-green-500 hover:bg-green-600 text-white px-4 py-2 rounded"
				>
					Login
				</button>
			</Form>
		</Formik>
	);
};

const handleRegister = (
	values: RegisterValues,
	{ setSubmitting }: FormikHelpers<RegisterValues>
  ) => {
	console.log(values);
	values.avatar_url = "values.avatar.name";
	// Handle form submission logic here
	// TODO: change localhost to iriesphere url
	fetch('http://localhost:8080/api/users/register', {
	  method: 'POST',
	  headers: {
		'Content-Type': 'application/json'
	  },
	  body: JSON.stringify(values)
	})
	.then(response => response.json())
	.then(data => {
	  console.log(data);
	  setSubmitting(false);
	})
	.catch(error => {
	  console.error('Error:', error);
	  setSubmitting(false);
	});
  };

  
const RegisterForm = ({ }) => {
	return (
		<Formik
			initialValues={{
				email: "",
				password: "",
				first_name: "",
				last_name: "",
				dob: "",
				avatar_url: "null",
				username: "",
				about: "",
			}}
			onSubmit={handleRegister}
		>
			{({ values, setFieldValue }) => (
				<Form>
					<div style={{backgroundColor: '#F0EAD6'}} className="flex flex-col max-w-md mx-auto mt-10 p-4">
					<h2 className="text-5xl font-rasa font-bold mb-4 text-green-700">Register</h2>
					<div>
						<label className="labelStyle">Email</label>
						<Field type="email" name="email"  className="inputStyle" />
					</div>
						<div className="mb-4">
							<label className="labelStyle">Password</label>
							<Field type="password" name="password"  className="inputStyle" />
						</div>
						<div className="mb-4">
							<label className="labelStyle">First Name</label>
							<Field type="text" name="first_name" className="inputStyle" />
						</div>
						<div className="mb-4">
							<label className="labelStyle">Last Name</label>
							<Field type="text" name="last_name" className="inputStyle" />
						</div>
						<div className="mb-4">
							<label className="labelStyle">Date of Birth</label>
							<Field type="date" name="dob" className="inputStyle" />
						</div>–
						<div className="mb-4">
							<label className="labelStyle">Avatar/Image</label>
							<Field name="avatar">
								{({ field, form }: FieldProps) => (
									<input
										id="avatar"
										name="avatar"
										type="file"
										onChange={(event) => {
											if (event.currentTarget.files) {
											const file = event.currentTarget.files[0];
											form.setFieldValue(field.name, file);
										}
									}}
								/>
							)}
						</Field>
					</div>
					<div>
						<label className="labelStyle">Username</label>
						<Field type="text" name="username" className="inputStyle" />
					</div>
					<div>
						<label className="labelStyle">About Me</label>
						<Field as="textarea" name="about" className="inputStyle" />
					</div>
					<button type="submit" className= "buttonStyle">
						Register
					</button>
				</div>
				</Form>
			)}
		</Formik>
	);
};

export { LoginForm, RegisterForm };
