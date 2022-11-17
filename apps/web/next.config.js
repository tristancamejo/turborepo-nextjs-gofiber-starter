/**
 * @type {import('next').NextConfig}
 */
module.exports = {
	reactStrictMode: true,
	experimental: {
		transpilePackages: ['ui'],
	},
	env: {
		API_URL: process.env.API_URL,
	},
};
