<script lang="ts">
	import type { Snippet } from 'svelte';
	import '@gleich/ui/styles.css';
	import { siteName } from '@gleich/ui';

	siteName.set('terminal.mattglei.ch');

	const { children }: { children: Snippet } = $props();
</script>

{@render children()}
