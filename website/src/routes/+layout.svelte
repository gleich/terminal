<script lang="ts">
	import type { Snippet } from 'svelte';
	import '@gleich/ui/styles.css';

	const { children }: { children: Snippet } = $props();
</script>

{@render children()}
