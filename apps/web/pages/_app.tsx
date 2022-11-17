import type { AppProps } from 'next/app';
import '../styles/globals.css';

function DisployIt({ Component, pageProps }: AppProps) {
	return <Component {...pageProps} />;
}

export default DisployIt;
