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
	firstName: string;
	lastName: string;
	dob: string;
	avatar: null;
	nickname: string;
	aboutMe: string;
}

const LoginForm = ({ }) => {
	return (
		<Formik
			initialValues={{
				username: "",
				password: "",
			}}
			onSubmit={(
				values: LoginValues,
				{ setSubmitting }: FormikHelpers<LoginValues>
			) => {
				setTimeout(() => {
					alert(JSON.stringify(values, null, 2));
					setSubmitting(false);
				}, 500);
			}}
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

const RegisterForm = ({ }) => {
	return (
		
		<Formik
			initialValues={{
				email: "",
				password: "",
				firstName: "",
				lastName: "",
				dob: "",
				avatar: null,
				nickname: "",
				aboutMe: "",
			}}
			onSubmit={(
				values: RegisterValues,
				{ setSubmitting }: FormikHelpers<RegisterValues>
			) => {
				// Handle form submission logic here
				console.log(values);
				setSubmitting(false);
			}}
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
							<Field type="text" name="firstName" className="inputStyle" />
						</div>
						<div className="mb-4">
							<label className="labelStyle">Last Name</label>
							<Field type="text" name="lastName" className="inputStyle" />
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
											form.setFieldValue(field.name, file);
										}
									}}
								/>
							)}
						</Field>
					</div>
					<div>
						<label className="labelStyle">Nickname</label>
						<Field type="text" name="nickname" className="inputStyle" />
					</div>
					<div>
						<label className="labelStyle">About Me</label>
						<Field as="textarea" name="aboutMe" className="inputStyle" />
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
