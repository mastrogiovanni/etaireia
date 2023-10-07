<script lang="ts">
	import { Button, Helper, Input, Label } from "flowbite-svelte";
	import UserData from "../../../components/UserData.svelte";
	import { recoverCredential } from "$lib/credential";

	let password: string;
	let name: string;
	let surname: string;
	let email: string;
	let fiscalCode: string;
	let dateOfBirth: Date;
	let placeOfBirth: string;
	let documentType: string;
	let documentId: string;
	let documentExpireTime: string;
	let mnemonic: string;

	async function _recoverCredential() {
		let credentials = await recoverCredential(mnemonic, {
			name,
			surname
		}, password)
        console.log(credentials)
	}

</script>

<h1>Importa una credenziale</h1>

<div class="mb-6">
    <Label for="mnemonic" class="block mb-2">Codice Mnemonico</Label>
    <Input id="mnemonic" placeholder="Mnemonic" />
    <Helper class="text-sm mt-2">
        Serie di parole che identifica una credenziale: (fanfara malgrado nevrosi...)
    </Helper>
</div>
<UserData 
    bind:name={name} 
    bind:surname={surname}
    bind:email={email}
    bind:fiscalCode={fiscalCode}
    bind:dateOfBirth={dateOfBirth}
    bind:placeOfBirth={placeOfBirth}
    bind:documentType={documentType}
    bind:documentId={documentId}
    bind:documentExpireTime={documentExpireTime}>
    </UserData>
<div class="mb-6">
    <Label for="password" class="block mb-2">Password</Label>
    <Input id="password" type="password" />
    <Helper class="text-sm mt-2">
        Viene utilizzata per decifrare la tua credenziale, assicurati che sia la stessa usata quando hai generato la credenziale
    </Helper>
</div>
<Button color="green" pill on:click={_recoverCredential}>Importa</Button>
