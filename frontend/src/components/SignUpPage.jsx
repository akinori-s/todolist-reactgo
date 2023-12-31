import React, { useState } from 'react';
import { Link, useNavigate } from "react-router-dom";
import { signup } from '../api/authApi';
import useForm from '../hooks/useForm';
import { validateEmail } from '../utils/utils';
import { useAuth } from '../contexts/AuthContext';

function SignUpPage() {
	const [loginInfo, handleChange, resetForm] = useForm({ firstName: '', lastName: '', email: '', password: '', passwordConfirmation: '' });
	const [error, setError] = useState(null);
	const navigate = useNavigate();
	const { setUser } = useAuth();

	const handleSubmit = async () => {
		setError(null);
		console.log("Logging in with", loginInfo); // wip: remove later

		if (loginInfo.firstName.length === 0 || 
			loginInfo.lastName.length === 0 || 
			loginInfo.email.length === 0 || 
			loginInfo.password.length === 0) {
			setError("Please fill out all fields!");
			return;
		}
		if (!validateEmail(loginInfo.email)) {
			setError("Please enter a valid email address!");
			return;
		}
		if (loginInfo.password.length < 8) {
			setError("Password must be at least 8 characters long!");
			return;
		}
		if (loginInfo.password !== loginInfo.passwordConfirmation) {
			setError("Passwords do not match!");
			return;
		}
		try {
			const response = await signup(loginInfo);
			setUser(response);
		} catch (err) {
			setError(err.message);
			return;
		}
		resetForm();
		navigate("/");
	};

	return (
		<div className="flex flex-col items-center h-screen w-full py-20 space-y-8 min-w-[300px]">
		<h1 className="text-2xl text-center">Sign Up</h1>
		<div className="flex flex-col w-1/2 min-w-[300px] max-w-[500px] space-y-4">
			<p style={{ color: 'red' }}>{error}</p>
			<input 
				type="text" 
				name="firstName"
				value={loginInfo.firstName}
				onChange={handleChange}
				className="border rounded-md px-2 py-1 w-full" 
				placeholder="First Name"
			/>
			<input 
				type="text" 
				name="lastName" 
				value={loginInfo.lastName} 
				onChange={handleChange} 
				className="border rounded-md px-2 py-1 w-full" 
				placeholder="Last Name"
			/>
			<input 
				type="email" 
				name="email" 
				value={loginInfo.email} 
				onChange={handleChange} 
				className="border rounded-md px-2 py-1 w-full" 
				placeholder="Email Address"
			/>
			<input 
				type="password" 
				name="password" 
				value={loginInfo.password} 
				onChange={handleChange} 
				className="border rounded-md px-2 py-1 w-full" 
				placeholder="Password"
			/>
			<input 
				type="password" 
				name="passwordConfirmation" 
				value={loginInfo.passwordConfirmation} 
				onChange={handleChange} 
				className="border rounded-md px-2 py-1 w-full" 
				placeholder="Confirm Password"
			/>
			<button 
				className="border rounded-md px-4 py-2 bg-blue-500 text-white hover:bg-blue-600 w-full"
				onClick={handleSubmit}
			>
				Sign Up
			</button>
		</div>
		<div className="flex flex-row items-center space-x-4">
			<p className="text-gray-500">Already have an account?</p>
			<Link 
				to="/login" 
				className="border rounded-md px-4 py-2 bg-blue-500 text-white hover:bg-blue-600" 
			>
				Login
			</Link>
		</div>
	</div>	
	);
}

export default SignUpPage;
