/** @type {import('next').NextConfig} */
const nextConfig = {
  env: {
    NEXT_PUBLIC_API_ADDR: process.env.API_ADDR,
    NEXT_PUBLIC_API_PORT: process.env.API_PORT,
    NEXT_PUBLIC_PORT: process.env.PORT,
  },
  output: "standalone",
  async rewrites() {
    return [
      {
        source: "/api/:path*",
        destination: `http://${process.env.API_ADDR}:${process.env.API_PORT}/api/:path*`, // Proxy to the backend API
      },
    ];
  },
};
export default nextConfig;
