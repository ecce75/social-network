"use client";

import React, { useState } from "react";
import { Formik, Field, Form, FormikHelpers, useFormikContext, FieldProps } from "formik";
import "../../styles/styles.css";
import { useRouter} from 'next/navigation';

interface LoginValues {
	username: string;
	password: string;
}

interface RegisterValues {
	[key: string]: any;
	email: string;
	password: string;
	first_name: string;
	last_name: string;
	dob: string;
	avatar: File | null;
	username: string;
	about: string;
}


const handleLogin = (
	values: LoginValues,
	formikHelpers: FormikHelpers<LoginValues>,
	router: any
  ) => {

	// Form submission logic
	fetch('http://localhost:8080/api/users/login', {
	  method: 'POST',
	  headers: {
		'Content-Type': 'application/json'
	  },
	  body: JSON.stringify(values),
	  credentials: 'include' // Send cookies with the request
	})
	.then(response => {
		if (!response.ok) {
			return response.text().then(text => {
				throw new Error(text);
			  });
		  }
		return response.json()
	})
	.then(data => {
	  console.log(data);
	  formikHelpers.setSubmitting(false);
	  router.push('/');
	})
	.catch(error => {
	  formikHelpers.setSubmitting(false);
	  console.error('Catch Error:', error);
	  alert("Invalid username or password")
	});
  };
 

const LoginForm = (({}) => {
	const router = useRouter();
	return ( 
		<Formik
			initialValues={{
				username: "",
				password: "",
			}}
			onSubmit={(values, formikHelpers) => handleLogin(values, formikHelpers, router)}
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
});

const handleRegister = (
	values: RegisterValues,
	formikHelpers: FormikHelpers<RegisterValues>,
	router: any
  ) => {
	const formData = new FormData();

	// Append all form fields to formData
	Object.keys(values).forEach((key) => {
	formData.append(key, values[key]);
	});

	// Form submission logic
	// TODO: change localhost to iriesphere url
	fetch('http://localhost:8080/api/users/register', {
	  method: 'POST',
	  body: formData,
	  credentials: 'include' // Send cookies with the request
	})
	.then(response => {
		if (!response.ok) {
			return response.text().then(text => {
				throw new Error(text);
			  });
		  }
		return response.json()
	})
	.then(data => {
	  console.log(data);
	  formikHelpers.setSubmitting(false);
	  router.push('/');
	})
	.catch(error => {
	  formikHelpers.setSubmitting(false);
	  console.error('Catch Error:', error);
	  alert("Invalid username or password")
	});
  };

  
const RegisterForm = ({ }) => {
	const router = useRouter();
	return (
		<Formik
			initialValues={{
				email: "",
				password: "",
				first_name: "",
				last_name: "",
				dob: "",
				avatar: null as File | null,
				username: "",
				about: "",
			}}
			onSubmit={(values, formikHelpers) => handleRegister(values, formikHelpers, router)}
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
						</div>
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
											setFieldValue("avatar", file);
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
	



