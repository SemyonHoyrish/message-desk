<script>
    let value = '';
    export let sendMessage;
</script>

<div class="bg-[#DDD1AE] w-full max-w-lg p-4 rounded-xl m-2 flex">
    <input class="text-[#5C4033] outline-none w-full" bind:value type="text" placeholder="your message..." />
    <button class="ml-5 text-[#5C4033] cursor-pointer" on:click={() => {sendMessage(value);value=''}}>send</button>
</div>


