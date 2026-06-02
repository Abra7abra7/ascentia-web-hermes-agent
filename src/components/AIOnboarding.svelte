// src/components/AIOnboarding.svelte
<script>
  let step = 1;
  let userGoal = "";
  let recommendation = "";

  function processGoal(goal) {
    userGoal = goal;
    if (goal === 'automation') {
      recommendation = "Kurz: Automatizácia firiem pomocou AI agentov v roku 2026";
    } else {
      recommendation = "Kurz: Masterclass efektívneho kódovania";
    }
    step = 2;
  }
</script>

<div class="w-full max-w-2xl mx-auto p-8 bg-zinc-900 border border-zinc-800 rounded-2xl shadow-2xl">
  {#if step === 1}
    <h2 class="text-3xl font-black tracking-tight text-zinc-50 mb-6">Aké vedomosti potrebuješ pre svoj vzostup?</h2>
    <div class="space-y-4">
      <button on:click={() => processGoal('automation')} class="w-full text-left p-4 rounded-xl bg-zinc-950 border border-zinc-800 hover:border-violet-500 transition-all text-zinc-300 font-medium">
        🚀 Chcem zautomatizovať svoje podnikanie a procesy
      </button>
      <button on:click={() => processGoal('code')} class="w-full text-left p-4 rounded-xl bg-zinc-950 border border-zinc-800 hover:border-violet-500 transition-all text-zinc-300 font-medium">
        💻 Chcem sa naučiť vyvíjať moderné AI-driven aplikácie
      </button>
    </div>
  {:else if step === 2}
    <div class="animate-fade-in">
      <span class="text-xs font-bold text-violet-500 tracking-widest uppercase">Tvoj osobný plán</span>
      <h3 class="text-2xl font-bold text-zinc-50 mt-2 mb-4">{recommendation}</h3>
      <p class="text-zinc-400 mb-6">Na základe tvojich preferencií sme ti odomkli prístup k úvodnej video lekcii zdarma a možnosti zakúpiť si plný balík s WhatsApp podporou.</p>
      
      <!-- Tu sa napojí Stripe Checkout Embed -->
      <button class="bg-violet-600 hover:bg-violet-700 text-zinc-50 font-bold py-4 px-8 rounded-xl transition-all shadow-lg shadow-violet-600/20">
        Spustiť kurz teraz
      </button>
    </div>
  {/if}
</div>

<style>
  .animate-fade-in {
    animation: fadeIn 0.5s ease-out forwards;
    opacity: 0;
  }
  @keyframes fadeIn {
    to {
      opacity: 1;
    }
  }
</style>