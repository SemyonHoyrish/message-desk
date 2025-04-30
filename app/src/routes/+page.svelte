<script>
    import { PUBLIC_SERVICE_NAME } from '$env/static/public';

    import MessageCard from '$lib/components/MessageCard.svelte';
    import MessageInput from '$lib/components/MessageInput.svelte';

    import { onMount } from 'svelte';
    let data = $state([]);
    onMount(() => {
        const json = localStorage.getItem('data') || '[]';
        data = JSON.parse(json);
    });

    const sendMessage = (value) => {
        if (value.replaceAll(' ', '') === '') return;

        data.push(value);
        localStorage.setItem('data', JSON.stringify(data));
    };
</script>

<header class="bg-[#B6A38B] p-6 text-center mb-5">
    <h1 class="text-2xl text-[#3B2F2F]">Welcome to <span class="font-bold">{PUBLIC_SERVICE_NAME}</span>'s message desk</h1>
</header>

<div class="h-dvh w-dvw flex flex-col items-center">
<MessageInput sendMessage={sendMessage} />

{#each data.toReversed() as msg}
    <MessageCard text={msg} />
{/each}
<MessageCard text="hello, world!" />
<MessageCard text="Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum" />

</div>

