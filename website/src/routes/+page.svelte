<script>
	import { Copyright, DynamicHead, NavLogo } from '@gleich/ui';
</script>

<DynamicHead title="terminal" description="ssh terminal.mattglei.ch" />

<div class="container">
	<main>
		<div class="title">
			<NavLogo />
			<h1>terminal</h1>
		</div>

		<div class="instructions">
			Run the following command on your system's terminal:
			<div class="command">ssh terminal.mattglei.ch</div>
		</div>

		<Copyright repo="gleich/terminal" />
	</main>
</div>

<style>
	.container {
		display: flex;
		align-items: center;
		justify-content: center;
	}

	main {
		min-height: 100vh;
		display: flex;
		justify-content: center;
		flex-direction: column;
		max-width: 800px;
		width: 100%;
		padding: 10px;
	}

	.title {
		display: flex;
		gap: 20px;
	}

	.instructions {
		margin-top: 50px;
		margin-bottom: 70px;
		display: flex;
		flex-direction: column;
		justify-content: center;
		gap: 10px;
	}

	.command {
		font-family: 'IBM Plex Mono';
		background-color: var(--green-background);
		color: var(--green-foreground);
		padding: 6px 20px;
		border: 1px solid var(--green-border);
		border-radius: var(--border-radius);
		box-shadow: var(--box-shadow);
		display: flex;
		width: fit-content;
		align-items: center;
		gap: 15px;
	}
</style>
