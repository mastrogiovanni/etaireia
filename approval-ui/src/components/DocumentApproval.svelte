<script>

    import { Card, Modal, Table, TableHead, TableHeadCell, Checkbox, TableBody, TableBodyRow, TableBodyCell, Gallery, Button } from "flowbite-svelte";
    import { onMount } from "svelte";

    export let personToValidate = new Object()
    export let open = false

    let selectedImage = undefined

    let images = []
    
    // onMount(() => {
    //   images = personToValidate.images
    //   images[0]
    // })

    $: {
      if (personToValidate) {
        images = personToValidate.images
        if (images && images.length > 0) {
          selectedImage = images[0]
        }
      }
    }

    function selectImage(image) {
        selectedImage = image
    }

</script>

<Modal title="Approvazione Utente: {personToValidate.name} {personToValidate.surname}" bind:open={open} autoclose size="lg">

  <p class="text-base leading-relaxed text-gray-500 dark:text-gray-400">
    Verifica che i dati sulle foto della carta siano uguali ai dati qui sotto.<br/>
    Puoi approvare solo dopo aver apposto tutte le check.
  </p>

  <div class="m-12 mb-2 mt-2">

  <Table hoverable={true} shadow color="green">
    <TableBody class="divide-y">
      <TableBodyRow>
        <TableBodyCell class="!p-4">
          <Checkbox />
        </TableBodyCell>
        <TableBodyCell>Codice Fiscale</TableBodyCell>
        <TableBodyCell>{personToValidate.fiscalCode}</TableBodyCell>
      </TableBodyRow>
      <TableBodyRow>
        <TableBodyCell class="!p-4">
          <Checkbox />
        </TableBodyCell>
        <TableBodyCell>Luogo Nascita</TableBodyCell>
        <TableBodyCell>{personToValidate.birthPlace}</TableBodyCell>
      </TableBodyRow>
      <TableBodyRow>
        <TableBodyCell class="!p-4">
          <Checkbox />
        </TableBodyCell>
        <TableBodyCell>Data Nascita</TableBodyCell>
        <TableBodyCell>{personToValidate.birthDate}</TableBodyCell>
      </TableBodyRow>
      <TableBodyRow>
        <TableBodyCell class="!p-4">
          <Checkbox />
        </TableBodyCell>
        <TableBodyCell>Data Emissione Documento</TableBodyCell>
        <TableBodyCell>{personToValidate.cardEmission}</TableBodyCell>
      </TableBodyRow>
      <TableBodyRow>
        <TableBodyCell class="!p-4">
          <Checkbox />
        </TableBodyCell>
        <TableBodyCell>Data Scadenza Documento</TableBodyCell>
        <TableBodyCell>{personToValidate.cardExpire}</TableBodyCell>
      </TableBodyRow>
    </TableBody>
  </Table>

  </div>

  {#if selectedImage}
  <Gallery class="gap-4">
      <img src={selectedImage.src} alt={selectedImage.alt} class="h-auto max-w-full rounded-lg" />
      <Gallery class="grid-cols-5" items={images} let:item>
          <div class="p-1" on:click={() => selectImage(item) }>
              <img src={item.src} alt={item.alt} class="h-auto max-w-full" />
          </div>
      </Gallery>
  </Gallery>
  {/if}

  <svelte:fragment slot="footer">
    <Button on:click={() => alert('Handle "success"')}>Approve</Button>
    <Button color="alternative">Cancella</Button>
  </svelte:fragment>


</Modal>
