<script>
    import { Button, Modal } from "sveltestrap";
    import { credential } from '../libs/stores';

    export let isOpen = false;

    function toggle() {
        isOpen = !isOpen
    }
</script>

<Modal body {isOpen} {toggle} header="Firma Digitale">
    <p>Questa è la tua chiave pubblica per esteso.</p>
    <p>Puoi condividerla con chi vuoi, ti identifica.</p>
    <p>
        <code>
        {$credential.publicKey.toUpperCase()}
        </code>
    </p>
    <img
        src="https://static6.depositphotos.com/1110975/606/i/950/depositphotos_6063518-stock-photo-signed-document.jpg"
        alt="Firma Digitale"
        class="img-fluid"
    />
</Modal>
  