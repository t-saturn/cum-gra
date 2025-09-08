import type { NextConfig } from 'next';

const nextConfig: NextConfig = {
  images: {
    remotePatterns: [
      {
        protocol: 'https',
        hostname: 'raw.githubusercontent.com',
        pathname: '/t-saturn/resources/**', // <-- más seguro, solo tu repo
      },
      {
        protocol: 'https',
        hostname: 'cdn.jsdelivr.net',
      },
    ],
  },
};

export default nextConfig;
