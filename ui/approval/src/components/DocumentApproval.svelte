<script lang="ts">
  import { approveCredential, rejectCredential } from "$lib/api";

  import { Modal, Table, Checkbox, TableBody, TableBodyRow, TableBodyCell, Gallery, Button, type ImgType, Img } from "flowbite-svelte";

  export let personToValidate: any;
  export let open = false
  export let command = () => {}

  let selectedImage: ImgType | undefined = undefined

  let images: ImgType[] = []
  
  $: {
    if (personToValidate) {
      images = [
        { alt: 'front', src: '/signer/api/v1/document/signer/front/' + personToValidate?.id },
        { alt: 'back', src: '/signer/api/v1/document/signer/back/' + personToValidate.id },
      ]
      selectedImage = images[0]
    }
  }

  let checkFiscalCode = false;
  let checkDocument = false;
  let checkPlaceOfBirth = false;
  let checkDateOfBirth = false;
  let checkDocumentExpireTime = false;

  let approvalDisabled = false

  $: {
    approvalDisabled = !(checkFiscalCode && checkDocument && checkPlaceOfBirth && checkDateOfBirth && checkDocumentExpireTime)
  }

  function selectImage(image: { src: any; alt: any; } | undefined) {
      selectedImage = image
  }

  function getDocumentType(docType: string) {
    if (docType === "CI") {
      return "Carta d'Identità"
    }
    if (docType === "PT") {
      return "Patente"
    }
    if (docType === "PS") {
      return "Passaporto"
    }
    return docType
  }

  async function _approve() {
    let resp = await approveCredential(personToValidate.hexPublicKey)
    console.log(resp)
    open = false
    command()
  }

  async function _reject() {
    let resp = await rejectCredential(personToValidate.hexPublicKey)
    console.log(resp)
    open = false
    command()
  }

</script>

<Modal title="Approvazione Utente: {personToValidate?.name} {personToValidate?.surname}" bind:open={open} autoclose size="lg">

  <p class="text-base leading-relaxed text-gray-500 dark:text-gray-400">
    Verifica che i dati sulle foto della carta siano uguali ai dati qui sotto.<br/>
    Puoi approvare solo dopo aver apposto tutte le check.
  </p>

  <div class="m-12 mb-2 mt-2">

  <Table hoverable={true} shadow color="green">
    <TableBody tableBodyClass="divide-y">
      <TableBodyRow>
        <TableBodyCell class="!p-4">
          <Checkbox bind:checked={checkFiscalCode} />
        </TableBodyCell>
        <TableBodyCell>Codice Fiscale</TableBodyCell>
        <TableBodyCell>{personToValidate.fiscalCode}</TableBodyCell>
      </TableBodyRow>
      <TableBodyRow>
        <TableBodyCell class="!p-4">
          <Checkbox bind:checked={checkDocument} />
        </TableBodyCell>
        <TableBodyCell>Documento</TableBodyCell>
        <TableBodyCell>{getDocumentType(personToValidate.documentType)}: {personToValidate.documentId}</TableBodyCell>
      </TableBodyRow>
      <TableBodyRow>
        <TableBodyCell class="!p-4">
          <Checkbox bind:checked={checkPlaceOfBirth} />
        </TableBodyCell>
        <TableBodyCell>Luogo Nascita</TableBodyCell>
        <TableBodyCell>{personToValidate.placeOfBirth}</TableBodyCell>
      </TableBodyRow>
      <TableBodyRow>
        <TableBodyCell class="!p-4">
          <Checkbox bind:checked={checkDateOfBirth} />
        </TableBodyCell>
        <TableBodyCell>Data Nascita</TableBodyCell>
        <TableBodyCell>{personToValidate.dateOfBirth}</TableBodyCell>
      </TableBodyRow>
      <TableBodyRow>
        <TableBodyCell class="!p-4">
          <Checkbox bind:checked={checkDocumentExpireTime} />
        </TableBodyCell>
        <TableBodyCell>Data Scadenza Documento</TableBodyCell>
        <TableBodyCell>{personToValidate.documentExpireTime}</TableBodyCell>
      </TableBodyRow>
    </TableBody>
  </Table>

  </div>

  <Gallery class="gap-4">
      {#if selectedImage}
        <img src={selectedImage.src} alt={selectedImage.alt} class="h-auto max-w-full rounded-lg" />
      {/if}
      <Gallery class="grid-cols-5" items={images} let:item>
          <div class="p-1" on:click={() => selectImage(item) }>
              <img src={item?.src} alt={item?.alt} class="h-auto max-w-full" />
          </div>
      </Gallery>
  </Gallery>

  <svelte:fragment slot="footer">
    <Button color="green" disabled={approvalDisabled} on:click={_approve}>Approva Credenziale</Button>
    <Button color="red" on:click={_reject}>Rigetta</Button>
    <Button color="alternative">Cancella</Button>
  </svelte:fragment>

</Modal>
