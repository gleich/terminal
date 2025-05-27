<script>
	import { DynamicHead } from '@gleich/ui';

	const command = 'ssh terminal.mattglei.ch';
	let copied = $state(false);

	function copyCommandToClipboard() {
		navigator.clipboard.writeText(command).then(() => {
			copied = true;
			setTimeout(() => (copied = false), 2_000);
		});
	}
</script>

<DynamicHead title="terminal" description="ssh terminal.mattglei.ch" />

<main>
	<div class="command-container">
		<div class="command">{command}</div>
		<div
			class="clipboard"
			role="button"
			tabindex="0"
			onclick={copyCommandToClipboard}
			onkeypress={(e) => (e.key === 'Enter' || e.key === ' ') && copyCommandToClipboard()}
			aria-label="Copy SSH command to clipboard"
			title="Copy SSH command to clipboard"
		>
			{#if copied}
				<svg
					class="check-icon"
					xmlns="http://www.w3.org/2000/svg"
					viewBox="0 0 24 24"
					fill="none"
					stroke="currentColor"
					stroke-width="2"
					stroke-linecap="round"
					stroke-linejoin="round"><polyline points="20 6 9 17 4 12"></polyline></svg
				>
			{:else}
				<svg
					class="clipboard-icon"
					xmlns="http://www.w3.org/2000/svg"
					viewBox="0 0 24 24"
					fill="none"
					stroke="currentColor"
					stroke-width="2"
					stroke-linecap="round"
					stroke-linejoin="round"
					><path d="M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2"
					></path><rect x="8" y="2" width="8" height="4" rx="1" ry="1"></rect></svg
				>
			{/if}
		</div>
	</div>
	<p class="instruction">
		{#if copied}
			Command copied to clipboard
		{:else}
			Run in your system's terminal emulator
		{/if}
	</p>
</main>

<style>
	main {
		min-height: 100vh;
		display: flex;
		align-items: center;
		justify-content: center;
		flex-direction: column;
		gap: 10px;
	}

	.command-container {
		display: flex;
	}

	.command {
		color: var(--green-foreground);
		background-color: var(--green-background);
		padding: 10px 15px;
		border-radius: var(--border-radius);
		font-family: 'IBM Plex Mono';
		border-top-right-radius: 0;
		border-bottom-right-radius: 0;
		border: 1px solid var(--green-border);
	}

	.clipboard {
		display: flex;
		align-items: center;
		padding: 0 12px;
		background-color: var(--background);
		border: 1px solid var(--border);
		border-left: 0;
		border-radius: var(--border-radius);
		border-top-left-radius: 0;
		border-bottom-left-radius: 0;
		cursor: pointer;
	}

	.clipboard-icon,
	.check-icon {
		height: 20px;
		width: auto;
	}

	.check-icon {
		color: var(--green-foreground);
	}

	.instruction {
		color: grey;
		font-family: 'IBM Plex Mono';
		font-size: 14px;
	}
</style>
