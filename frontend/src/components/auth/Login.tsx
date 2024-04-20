"use client";

import {Field, Form, Formik, FormikHelpers} from "formik";
import {useRouter} from "next/navigation";
import React from "react";

interface LoginValues {
username: string;
password: string;
}

const handleLogin = (
    values: LoginValues,
    formikHelpers: FormikHelpers<LoginValues>,
    router: any
) => {
    const BE_PORT = process.env.NEXT_PUBLIC_BACKEND_PORT;
    const FE_URL = process.env.NEXT_PUBLIC_URL;
    // Form submission logic
    fetch(`${FE_URL}:${BE_PORT}/api/users/login`, {
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
            formikHelpers.setSubmitting(false);
            router.push('/');
        })
        .catch(error => {
            formikHelpers.setSubmitting(false);
            if (JSON.parse(error.message).error === "Incorrect login credentials.") {
                alert("Invalid username or password")
            }
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
                    className="bg-green-500 hover:bg-primary text-white px-4 py-2 rounded"
                >
                    Login
                </button>
            </Form>
        </Formik>
    );
});

export default LoginForm;