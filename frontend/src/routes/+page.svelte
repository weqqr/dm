<script lang="ts">
	import ChatMessage from '$lib/ChatMessage.svelte';
	import { Message } from '$lib/message';

	const messages = [
		new Message('abc', 'sample text', new Date(Date.now())),
		new Message('xyz', 'sample text 2', new Date(Date.now())),
	];
</script>

<div class="flex flex-row h-screen w-screen">
	<div class="bg-dark p-1 select-none">
		<div class="min-w-48">
			<div class="rounded bg-white text-black px-2">test</div>
		</div>
	</div>
	<div class="bg flex flex-col gap-2 flex-1">
		{#each messages as message}
			<ChatMessage {message} />
		{/each}
	</div>
</div>

<style>
	.bg {
		background-color: #1f1f1f;
	}

	.bg-dark {
		background-color: #1a1a1a;
	}
</style>
