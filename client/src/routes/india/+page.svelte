<script>
    import { onMount } from "svelte";
    import { getSliceDesc } from "$lib";

    /** @type {undefined | any}*/
    let data;

    onMount(async () => {
        const response = await fetch("http://localhost:8080/india");
        data = await response.json();
    });
</script>

<svelte:head>
    <title>India News | The Daily Breif</title>
</svelte:head>

<div class="my-8 px-80">
    <h1 class="text-center mb-8">India News</h1>
    {#if data}
        <div data-cy="news" class="flex flex-col gap-12">
            {#each data as d}
                {#each d as doc}
                    <a data-cy="news-link" href={doc.Link}>
                        <div
                            data-cy="news-block"
                            class="flex items-center gap-8"
                        >
                            <img
                                src={doc.Enclosure.Url}
                                alt="news"
                                width="500"
                            />
                            <div class="space-y-8">
                                <h3 data-cy="news-title">
                                    {doc.Title}
                                </h3>
                                <p data-cy="news-desc">
                                    {getSliceDesc(doc.Desc)}
                                </p>
                            </div>
                        </div>
                    </a>
                {/each}
            {/each}
        </div>
    {:else}
        <div>
            <p class="text-center">Loading...</p>
        </div>
    {/if}
</div>
