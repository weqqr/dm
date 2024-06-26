<script lang="ts">
	import type { Message } from '$lib/message';

	export let message: Message;

	function getDate() {
		const date = new Date(Date.now());
		return date.toLocaleDateString(window.navigator.languages, {
			day: 'numeric',
			month: 'long',
			year: 'numeric'
		});
	}
</script>

<div class="flex flex-row gap-2">
	<div style="width: 48px; height: 48px" class="rounded-full bg-gray-700"></div>
	<div class="flex flex-col">
		<div>
			<a href="#" class="font-semibold">{message.nickname}</a>
			<span class="text-sm">{message.created_at}</span>
		</div>
		<span class="whitespace-pre">{message.text}</span>
	</div>
</div>
