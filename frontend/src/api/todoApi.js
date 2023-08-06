import axios from 'axios';

const BASE_URL = "http://localhost:8080";
const api = axios.create({
	baseURL: BASE_URL,
	headers: {
		'Content-Type': 'application/json',
	},
	withCredentials: true,
});

export const getTodoList = async () => {
	try {
		const response = await api.get(`/todos`);
		return response.data;
	} catch (error) {
		throw error.response.data;
	}
};