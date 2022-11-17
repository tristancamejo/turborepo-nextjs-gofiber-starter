import { useEffect, useState } from 'react';

export default function Web() {
	const [text, setText] = useState<string | null>(null);
	useEffect(() => {
		(async () => {
			const response = await fetch(`${process.env.API_URL}/`);

			if (response.ok) {
				const data = await response.text();
				setText(data);
			}
		})();
	}, []);

	return (
		<main>
			<div className="relative px-6">
				<div className="mx-auto max-w-3xl pt-20 pb-32 sm:pt-48 sm:pb-40">
					<div>
						<div>
							<h1 className="text-4xl font-bold tracking-tight sm:text-center sm:text-6xl">{text ?? 'Loading...'}</h1>
							<p className="mt-6 text-lg leading-8 text-gray-600 sm:text-center">this was fetched from the api.</p>
						</div>
					</div>
				</div>
			</div>
		</main>
	);
}
