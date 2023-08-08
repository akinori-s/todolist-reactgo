export const login = (username, password) => {
	// simulate a successful login.
	// wip: replace this with an actual call to backend.
	return new Promise((resolve, reject) => {
		setTimeout(() => {
			if (username === "admin" && password === "password") {
				resolve({ data: { token: "sample_token" } });
			} else {
				reject(new Error("Invalid username or password"));
			}
		}, 1000);
	});
};
