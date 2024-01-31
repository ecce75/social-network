import React from "react";
import { Formik, Field, Form, FormikHelpers } from 'formik';


interface Values {
    username: string;
    password: string;
}

const LoginForm= ({ }) => {
    return (
        <Formik
            initialValues={{
                username: '',
                password: '',
            }}

            onSubmit={(
                values: Values,
                { setSubmitting }: FormikHelpers<Values>
            ) => {
                setTimeout(() => {
                    alert(JSON.stringify(values, null, 2));
                    setSubmitting(false);
                }, 500);
            }}
            >
            <Form className="flex items-center">

                <div className="mb-3">
                    <Field className="form-control" id="username" name="username" placeholder="Username"
                           aria-describedby="usernameHelp"/>
                </div>

                <div className="mb-3">
                    <Field className="form-control" id="password" name="password" placeholder="Password"
                           type="password"/>
                </div>

                <button type="submit" className="btn btn-primary">Login</button>
            </Form>
        </Formik>
    );
};

export default LoginForm;
