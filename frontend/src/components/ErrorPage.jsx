import { useRouteError } from "react-router-dom";

export default function ErrorPage() {
	const error = useRouteError();
	console.error(error);

	return (
	<div className="flex flex-col items-center h-screen w-full py-20 space-y-8 min-w-[300px]">
		<h1 className="text-2xl text-center">Wow! Much empty...</h1>
		<p className="text-center">Looks like you've reached a page that doesn't exist.</p>
	</div>
	);
}