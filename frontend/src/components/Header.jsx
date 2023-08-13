import React, { useEffect } from 'react';
import { Link, useLocation } from "react-router-dom";
import { useAuth } from '../contexts/AuthContext';
import { checkLogin } from '../api/authApi';

function Header() {
	const location = useLocation()
	const { user, setUser } = useAuth();

	useEffect(() => {
		async function doCheckLogin() {
			try {
				const data = await checkLogin();
				if (data) {
					setUser(data);
				}
				else {
					setUser(null);
				}
			} catch (error) {
				console.error("Failed to check login:", error);
			}
		}
		doCheckLogin();
	}, []);

	function handleSignout() {
		setUser(null);
		console.log("Signout successful");
	}

	return (
		<div className="flex justify-between items-center h-14 w-full px-6 border-b min-w-[300px]">
			<Link 
				to="/"
				className="text-2xl text-center" 
			>
				TodoList
			</Link>
			{user!= null ? <div>Hello {user.firstName}!</div> : <></>}
			{((user == null && location.pathname !== '/login' && location.pathname !== '/signup')) ? 
				<div>
					<Link 
						to="/login" 
						className="border rounded-md px-2 py-1 mr-2 bg-blue-500 text-white hover:bg-blue-600"
					>
						Login
					</Link>
					<Link 
						to="/signup"
						className="border rounded-md px-2 py-1 text-blue-500 border-solid border-blue-500 hover:text-white hover:bg-blue-500"
					>
						Sign Up
					</Link>
				</div>
			:
				<></>
			}
			{
				user != null ?
				<button 
					className="border rounded-md px-2 py-1 mr-2 bg-blue-500 text-white hover:bg-blue-600"
					onClick={handleSignout}
				>
					Sign Out
				</button>
			:
				<></>
			}
		</div>
	);
}

export default Header;
