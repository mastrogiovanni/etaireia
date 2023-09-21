export function load({ cookies, request }) {
	const forwardAuth = cookies.get('_forward_auth');
	const authUser = request.headers.get("iv-user")
	return {
		authUser,
		forwardAuth
	};
}
